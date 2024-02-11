package inscribe

import (
	"encoding/hex"
	"fmt"

	"github.com/b2network/b2committer/pkg/log"

	"github.com/b2network/b2committer/pkg/btcapi"
	btcmempool "github.com/b2network/b2committer/pkg/btcapi/mempool"
	"github.com/btcsuite/btcd/mempool"
	"github.com/btcsuite/btcd/rpcclient"

	"github.com/btcsuite/btcd/chaincfg/chainhash"

	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/pkg/errors"
)

// InscriptionData is the data of an inscription
type InscriptionData struct {
	Body        []byte // The body of the inscription
	Destination string // The destination of the inscription
}

// InscriptionRequest is the request of an inscription
type InscriptionRequest struct {
	CommitTxOutPointList   []*wire.OutPoint
	CommitTxPrivateKeyList []*btcec.PrivateKey // If used without RPC,
	// a local signature is required for committing the commit tx.
	// Currently, CommitTxPrivateKeyList[i] sign CommitTxOutPointList[i]
	CommitFeeRate      int64
	FeeRate            int64
	DataList           []InscriptionData
	SingleRevealTxOnly bool // Currently, the official Ordinal parser can only parse a single NFT per transaction.
	// When the official Ordinal parser supports parsing multiple NFTs in the future, we can consider using a single reveal transaction.
	RevealOutValue int64
}

// inscriptionTxCtxData is the context of an inscription
type inscriptionTxCtxData struct {
	privateKey              *btcec.PrivateKey
	inscriptionScript       []byte
	commitTxAddressPkScript []byte
	controlBlockWitness     []byte
	recoveryPrivateKeyWIF   string
	revealTxPrevOutput      *wire.TxOut
}

type blockchainClient struct {
	rpcClient    *rpcclient.Client
	btcAPIClient btcapi.Client
}

// InscriptionTool is the tool of an inscription
type InscriptionTool struct {
	net                       *chaincfg.Params
	client                    *blockchainClient
	commitTxPrevOutputFetcher *txscript.MultiPrevOutFetcher
	commitTxPrivateKeyList    []*btcec.PrivateKey
	txCtxDataList             []*inscriptionTxCtxData
	revealTxPrevOutputFetcher *txscript.MultiPrevOutFetcher
	revealTx                  []*wire.MsgTx
	commitTx                  *wire.MsgTx
}

const (
	defaultSequenceNum    = wire.MaxTxInSequenceNum - 10
	defaultRevealOutValue = int64(500) // 500 sat, ord default 10000
	MaxStandardTxWeight   = blockchain.MaxBlockWeight / 10
)

type Result struct {
	CommitTxHash     *chainhash.Hash
	RevealTxHashList []*chainhash.Hash
	Inscriptions     []string
	Fees             int64
}

func Inscribe(privateKeyHex string, stateRootHash string, proofRootHash string, destination string, net *chaincfg.Params) (*Result, error) {
	btcAPIClient := btcmempool.NewClient(net)
	dataList := make([]InscriptionData, 0)
	utxoPrivateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("decode private key:  %s error: %w", privateKeyHex, err)
	}
	utxoPrivateKey, _ := btcec.PrivKeyFromBytes(utxoPrivateKeyBytes)
	utxoTaprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(utxoPrivateKey.PubKey())), net)
	if err != nil {
		return nil, fmt.Errorf("private key generate address error: %w", err)
	}
	dataList = append(dataList, InscriptionData{
		Body:        []byte(stateRootHash + proofRootHash),
		Destination: destination,
	})
	fmt.Println(utxoTaprootAddress.EncodeAddress())
	req, err := NewInscriptionRequest(btcAPIClient, utxoTaprootAddress, utxoPrivateKey, dataList)
	if err != nil {
		return nil, fmt.Errorf("create inscription request error: %w", err)
	}
	tool, err := NewInscriptionToolWithBtcAPIClient(net, btcAPIClient, req)
	if err != nil {
		return nil, fmt.Errorf("NewInscriptionToolWithBtcAPIClient error: %w", err)
	}
	commitTxHash, revealTxHashList, inscriptions, fees, err := tool.Inscribe()
	return &Result{
		CommitTxHash:     commitTxHash,
		RevealTxHashList: revealTxHashList,
		Inscriptions:     inscriptions,
		Fees:             fees,
	}, err
}

// NewInscriptionRequest creates a new InscriptionRequest
func NewInscriptionRequest(btcAPIClient btcapi.Client, address btcutil.Address, privateKey *btcec.PrivateKey, dataList []InscriptionData) (*InscriptionRequest, error) {
	unspentList, err := btcAPIClient.ListUnspent(address)
	if err != nil {
		return nil, err
	}
	commitTxOutPointList := make([]*wire.OutPoint, 0)
	commitTxPrivateKeyList := make([]*btcec.PrivateKey, 0)
	for i := range unspentList {
		commitTxOutPointList = append(commitTxOutPointList, unspentList[i].Outpoint)
		commitTxPrivateKeyList = append(commitTxPrivateKeyList, privateKey)
	}
	recommendedFees, err := btcAPIClient.GetRecommendedFees()
	if err != nil {
		return nil, err
	}

	return &InscriptionRequest{
		CommitTxOutPointList:   commitTxOutPointList,
		CommitTxPrivateKeyList: commitTxPrivateKeyList,
		CommitFeeRate:          recommendedFees.FastestFee + 1,
		FeeRate:                recommendedFees.FastestFee + 2,
		DataList:               dataList,
		SingleRevealTxOnly:     false,
	}, nil
}

func NewInscriptionToolWithBtcAPIClient(net *chaincfg.Params, btcAPIClient btcapi.Client, request *InscriptionRequest) (*InscriptionTool, error) {
	if len(request.CommitTxPrivateKeyList) != len(request.CommitTxOutPointList) {
		return nil, errors.New("the length of CommitTxPrivateKeyList and CommitTxOutPointList should be the same")
	}
	tool := &InscriptionTool{
		net: net,
		client: &blockchainClient{
			btcAPIClient: btcAPIClient,
		},
		commitTxPrevOutputFetcher: txscript.NewMultiPrevOutFetcher(nil),
		commitTxPrivateKeyList:    request.CommitTxPrivateKeyList,
		revealTxPrevOutputFetcher: txscript.NewMultiPrevOutFetcher(nil),
	}
	return tool, tool._initTool(net, request)
}

// _initTool initializes the InscriptionTool
func (tool *InscriptionTool) _initTool(net *chaincfg.Params, request *InscriptionRequest) error {
	revealOutValue := defaultRevealOutValue
	if request.RevealOutValue > 0 {
		revealOutValue = request.RevealOutValue
	}
	tool.txCtxDataList = make([]*inscriptionTxCtxData, len(request.DataList))
	destinations := make([]string, len(request.DataList))
	for i := 0; i < len(request.DataList); i++ {
		txCtxData, err := createInscriptionTxCtxData(net, request.DataList[i])
		if err != nil {
			return err
		}
		tool.txCtxDataList[i] = txCtxData
		destinations[i] = request.DataList[i].Destination
	}
	totalRevealPrevOutput, err := tool.buildEmptyRevealTx(destinations, revealOutValue, request.FeeRate)
	if err != nil {
		return err
	}
	err = tool.buildCommitTx(request.CommitTxOutPointList, totalRevealPrevOutput, request.CommitFeeRate)
	if err != nil {
		return err
	}
	err = tool.completeRevealTx()
	if err != nil {
		return err
	}
	err = tool.signCommitTx()
	if err != nil {
		return errors.Wrap(err, "sign commit tx error")
	}
	return err
}

// createInscriptionTxCtxData creates a new inscriptionTxCtxData
func createInscriptionTxCtxData(net *chaincfg.Params, data InscriptionData) (*inscriptionTxCtxData, error) {
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, err
	}
	inscriptionBuilder := txscript.NewScriptBuilder().
		AddData(schnorr.SerializePubKey(privateKey.PubKey())).
		AddOp(txscript.OP_CHECKSIG).
		AddOp(txscript.OP_FALSE).
		AddOp(txscript.OP_IF)
	maxChunkSize := 520
	bodySize := len(data.Body)
	for i := 0; i < bodySize; i += maxChunkSize {
		end := i + maxChunkSize
		if end > bodySize {
			end = bodySize
		}
		// to skip txscript.MaxScriptSize 10000
		inscriptionBuilder.AddFullData(data.Body[i:end])
	}
	inscriptionScript, err := inscriptionBuilder.Script()
	if err != nil {
		return nil, err
	}
	// to skip txscript.MaxScriptSize 10000
	inscriptionScript = append(inscriptionScript, txscript.OP_ENDIF)

	leafNode := txscript.NewBaseTapLeaf(inscriptionScript)
	proof := &txscript.TapscriptProof{
		TapLeaf:  leafNode,
		RootNode: leafNode,
	}

	controlBlock := proof.ToControlBlock(privateKey.PubKey())
	controlBlockWitness, err := controlBlock.ToBytes()
	if err != nil {
		return nil, err
	}

	tapHash := proof.RootNode.TapHash()
	commitTxAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootOutputKey(privateKey.PubKey(), tapHash[:])), net)
	if err != nil {
		return nil, err
	}
	commitTxAddressPkScript, err := txscript.PayToAddrScript(commitTxAddress)
	if err != nil {
		return nil, err
	}

	recoveryPrivateKeyWIF, err := btcutil.NewWIF(txscript.TweakTaprootPrivKey(*privateKey, tapHash[:]), net, true)
	if err != nil {
		return nil, err
	}

	return &inscriptionTxCtxData{
		privateKey:              privateKey,
		inscriptionScript:       inscriptionScript,
		commitTxAddressPkScript: commitTxAddressPkScript,
		controlBlockWitness:     controlBlockWitness,
		recoveryPrivateKeyWIF:   recoveryPrivateKeyWIF.String(),
	}, nil
}

// buildEmptyRevealTx builds an empty reveal transaction, and calculate the revealTx input value, it means the total value of the reveal transaction
func (tool *InscriptionTool) buildEmptyRevealTx(destination []string, revealOutValue, feeRate int64) (int64, error) {
	var revealTx []*wire.MsgTx
	totalPrevOutput := int64(0)
	total := len(tool.txCtxDataList)
	addTxInTxOutIntoRevealTx := func(tx *wire.MsgTx, index int) error {
		in := wire.NewTxIn(&wire.OutPoint{Index: uint32(index)}, nil, nil) // #nosec 701
		in.Sequence = defaultSequenceNum
		tx.AddTxIn(in)
		receiver, err := btcutil.DecodeAddress(destination[index], tool.net)
		if err != nil {
			return err
		}
		scriptPubKey, err := txscript.PayToAddrScript(receiver)
		if err != nil {
			return err
		}
		out := wire.NewTxOut(revealOutValue, scriptPubKey)
		tx.AddTxOut(out)
		return nil
	}

	revealTx = make([]*wire.MsgTx, total)
	for i := 0; i < total; i++ {
		tx := wire.NewMsgTx(wire.TxVersion)
		err := addTxInTxOutIntoRevealTx(tx, i)
		if err != nil {
			return 0, err
		}
		prevOutput := revealOutValue + int64(tx.SerializeSize())*feeRate
		{
			emptySignature := make([]byte, 64)
			emptyControlBlockWitness := make([]byte, 33)
			fee := (int64(wire.TxWitness{emptySignature, tool.txCtxDataList[i].inscriptionScript, emptyControlBlockWitness}.SerializeSize()+2+3) / 4) * feeRate
			prevOutput += fee
			tool.txCtxDataList[i].revealTxPrevOutput = &wire.TxOut{
				PkScript: tool.txCtxDataList[i].commitTxAddressPkScript,
				Value:    prevOutput,
			}
		}
		totalPrevOutput += prevOutput
		revealTx[i] = tx
	}

	tool.revealTx = revealTx
	return totalPrevOutput, nil
}

// buildCommitTx builds a commit transaction
func (tool *InscriptionTool) buildCommitTx(commitTxOutPointList []*wire.OutPoint, totalRevealPrevOutput, commitFeeRate int64) error {
	totalSenderAmount := btcutil.Amount(0)
	tx := wire.NewMsgTx(wire.TxVersion)
	var changePkScript *[]byte
	for i := range commitTxOutPointList {
		txOut, err := tool.getTxOutByOutPoint(commitTxOutPointList[i])
		if err != nil {
			return err
		}
		if changePkScript == nil { // first sender as change address
			changePkScript = &txOut.PkScript
		}
		in := wire.NewTxIn(commitTxOutPointList[i], nil, nil)
		in.Sequence = defaultSequenceNum
		tx.AddTxIn(in)

		totalSenderAmount += btcutil.Amount(txOut.Value)
	}
	for i := range tool.txCtxDataList {
		tx.AddTxOut(tool.txCtxDataList[i].revealTxPrevOutput)
	}

	tx.AddTxOut(wire.NewTxOut(0, *changePkScript))
	fee := btcutil.Amount(mempool.GetTxVirtualSize(btcutil.NewTx(tx))) * btcutil.Amount(commitFeeRate)
	changeAmount := totalSenderAmount - btcutil.Amount(totalRevealPrevOutput) - fee
	if changeAmount > 0 {
		tx.TxOut[len(tx.TxOut)-1].Value = int64(changeAmount)
	} else {
		tx.TxOut = tx.TxOut[:len(tx.TxOut)-1]
		if changeAmount < 0 {
			feeWithoutChange := btcutil.Amount(mempool.GetTxVirtualSize(btcutil.NewTx(tx))) * btcutil.Amount(commitFeeRate)
			if totalSenderAmount-btcutil.Amount(totalRevealPrevOutput)-feeWithoutChange < 0 {
				return errors.New("insufficient balance")
			}
		}
	}
	tool.commitTx = tx
	return nil
}

// getTxOutByOutPoint gets the txOut by outPoint
func (tool *InscriptionTool) getTxOutByOutPoint(outPoint *wire.OutPoint) (*wire.TxOut, error) {
	var txOut *wire.TxOut
	if tool.client.rpcClient != nil {
		tx, err := tool.client.rpcClient.GetRawTransactionVerbose(&outPoint.Hash)
		if err != nil {
			return nil, err
		}
		if int(outPoint.Index) >= len(tx.Vout) {
			return nil, errors.New("err out point")
		}
		vout := tx.Vout[outPoint.Index]
		pkScript, err := hex.DecodeString(vout.ScriptPubKey.Hex)
		if err != nil {
			return nil, err
		}
		amount, err := btcutil.NewAmount(vout.Value)
		if err != nil {
			return nil, err
		}
		txOut = wire.NewTxOut(int64(amount), pkScript)
	} else {
		tx, err := tool.client.btcAPIClient.GetRawTransaction(&outPoint.Hash)
		if err != nil {
			return nil, err
		}
		if int(outPoint.Index) >= len(tx.TxOut) {
			return nil, errors.New("err out point")
		}
		txOut = tx.TxOut[outPoint.Index]
	}
	tool.commitTxPrevOutputFetcher.AddPrevOut(*outPoint, txOut)
	return txOut, nil
}

// completeRevealTx completes the reveal transaction
func (tool *InscriptionTool) completeRevealTx() error {
	for i := range tool.txCtxDataList {
		tool.revealTxPrevOutputFetcher.AddPrevOut(wire.OutPoint{
			Hash:  tool.commitTx.TxHash(),
			Index: uint32(i), // #nosec 701
		}, tool.txCtxDataList[i].revealTxPrevOutput)
		if len(tool.revealTx) == 1 {
			tool.revealTx[0].TxIn[i].PreviousOutPoint.Hash = tool.commitTx.TxHash()
		} else {
			tool.revealTx[i].TxIn[0].PreviousOutPoint.Hash = tool.commitTx.TxHash()
		}
	}
	witnessList := make([]wire.TxWitness, len(tool.txCtxDataList))
	for i := range tool.txCtxDataList {
		revealTx := tool.revealTx[0]
		idx := i
		if len(tool.revealTx) != 1 {
			revealTx = tool.revealTx[i]
			idx = 0
		}
		witnessArray, err := txscript.CalcTapscriptSignaturehash(txscript.NewTxSigHashes(revealTx, tool.revealTxPrevOutputFetcher),
			txscript.SigHashDefault, revealTx, idx, tool.revealTxPrevOutputFetcher, txscript.NewBaseTapLeaf(tool.txCtxDataList[i].inscriptionScript))
		if err != nil {
			return err
		}
		signature, err := schnorr.Sign(tool.txCtxDataList[i].privateKey, witnessArray)
		if err != nil {
			return err
		}
		witnessList[i] = wire.TxWitness{signature.Serialize(), tool.txCtxDataList[i].inscriptionScript, tool.txCtxDataList[i].controlBlockWitness}
	}
	for i := range witnessList {
		if len(tool.revealTx) == 1 {
			tool.revealTx[0].TxIn[i].Witness = witnessList[i]
		} else {
			tool.revealTx[i].TxIn[0].Witness = witnessList[i]
		}
	}
	// check tx max tx wight
	for i, tx := range tool.revealTx {
		revealWeight := blockchain.GetTransactionWeight(btcutil.NewTx(tx))
		if revealWeight > MaxStandardTxWeight {
			return fmt.Errorf("reveal(index %d) transaction weight greater "+
				"than %d (MAX_STANDARD_TX_WEIGHT): %d", i, MaxStandardTxWeight, revealWeight)
		}
	}
	return nil
}

// signCommitTx signs the commit transaction
func (tool *InscriptionTool) signCommitTx() error {
	if len(tool.commitTxPrivateKeyList) == 0 {
		commitSignTransaction, isSignComplete, err := tool.client.rpcClient.SignRawTransactionWithWallet(tool.commitTx)
		if err != nil {
			log.Errorf("sign commit tx error, %v", err)
			return err
		}
		if !isSignComplete {
			return errors.New("sign commit tx error")
		}
		tool.commitTx = commitSignTransaction
	} else {
		witnessList := make([]wire.TxWitness, len(tool.commitTx.TxIn))
		for i := range tool.commitTx.TxIn {
			txOut := tool.commitTxPrevOutputFetcher.FetchPrevOutput(tool.commitTx.TxIn[i].PreviousOutPoint)
			witness, err := txscript.TaprootWitnessSignature(tool.commitTx, txscript.NewTxSigHashes(tool.commitTx, tool.commitTxPrevOutputFetcher),
				i, txOut.Value, txOut.PkScript, txscript.SigHashDefault, tool.commitTxPrivateKeyList[i])
			if err != nil {
				return err
			}
			witnessList[i] = witness
		}
		for i := range witnessList {
			tool.commitTx.TxIn[i].Witness = witnessList[i]
		}
	}
	return nil
}

func (tool *InscriptionTool) sendRawTransaction(tx *wire.MsgTx) (*chainhash.Hash, error) {
	if tool.client.rpcClient != nil {
		return tool.client.rpcClient.SendRawTransaction(tx, false)
	}
	return tool.client.btcAPIClient.BroadcastTx(tx)
}

// Inscribe inscribes the data
func (tool *InscriptionTool) Inscribe() (commitTxHash *chainhash.Hash, revealTxHashList []*chainhash.Hash, inscriptions []string, fees int64, err error) {
	fees = tool.calculateFee()
	commitTxHash, err = tool.sendRawTransaction(tool.commitTx)
	if err != nil {
		return nil, nil, nil, fees, errors.Wrap(err, "send commit tx error")
	}
	revealTxHashList = make([]*chainhash.Hash, len(tool.revealTx))
	inscriptions = make([]string, len(tool.txCtxDataList))
	for i := range tool.revealTx {
		_revealTxHash, err := tool.sendRawTransaction(tool.revealTx[i])
		if err != nil {
			return commitTxHash, revealTxHashList, nil, fees, errors.Wrap(err, fmt.Sprintf("send reveal tx error, %d。", i))
		}
		revealTxHashList[i] = _revealTxHash
		if len(tool.revealTx) == len(tool.txCtxDataList) {
			inscriptions[i] = fmt.Sprintf("%si0", _revealTxHash)
		} else {
			inscriptions[i] = fmt.Sprintf("%si", _revealTxHash)
		}
	}
	if len(tool.revealTx) != len(tool.txCtxDataList) {
		for i := len(inscriptions) - 1; i > 0; i-- {
			inscriptions[i] = fmt.Sprintf("%s%d", inscriptions[0], i)
		}
	}
	return commitTxHash, revealTxHashList, inscriptions, fees, nil
}

// calculateFee calculates the fee
func (tool *InscriptionTool) calculateFee() int64 {
	fees := int64(0)
	for _, in := range tool.commitTx.TxIn {
		fees += tool.commitTxPrevOutputFetcher.FetchPrevOutput(in.PreviousOutPoint).Value
	}
	for _, out := range tool.commitTx.TxOut {
		fees -= out.Value
	}
	for _, tx := range tool.revealTx {
		for _, in := range tx.TxIn {
			fees += tool.revealTxPrevOutputFetcher.FetchPrevOutput(in.PreviousOutPoint).Value
		}
		for _, out := range tx.TxOut {
			fees -= out.Value
		}
	}
	return fees
}
