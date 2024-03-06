package b2node

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/b2network/b2committer/pkg/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type NodeClient struct {
	PrivateKey      *ecdsa.PrivateKey
	ContractAddress string
	Address         string
	ChainID         int64
	Conn            *ethclient.Client
	Committer       *contract.Committer
	Auth            *bind.TransactOpts
}

func NewNodeClient(privateKeyStr string, chainID int64, address string, contractAddress string, conn *ethclient.Client) *NodeClient {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if nil != err {
		panic(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	committer, err := contract.NewCommitter(common.HexToAddress(contractAddress), conn)
	return &NodeClient{
		PrivateKey: privateKey,
		Address:    address,
		ChainID:    chainID,
		Conn:       conn,
		Committer:  committer,
		Auth:       auth,
	}
}

func (n NodeClient) AddCommitter(address string) (string, error) {
	tx, err := n.Committer.AddProposer(&bind.TransactOpts{
		From:   n.Auth.From,
		Signer: n.Auth.Signer,
	}, common.HexToAddress(address))
	if err != nil {
		return "", fmt.Errorf("[AddCommitter] err: %s", err)
	}
	return tx.Hash().String(), nil
}

func (n NodeClient) QueryAllProposers() ([]common.Address, error) {
	res, err := n.Committer.AllProposers(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("[QueryAllProposers] err: %s", err)
	}
	return res, nil
}

func (n NodeClient) RemoveProposers(address common.Address) (*types.Transaction, error) {
	tx, err := n.Committer.RemoveProposer(&bind.TransactOpts{
		From:   n.Auth.From,
		Signer: n.Auth.Signer,
	}, address)
	if err != nil {
		return nil, fmt.Errorf("[RemoveProposers] err: %s", err)
	}
	return tx, nil
}

func (n NodeClient) QueryLastProposal() (*contract.CommitterProposal, error) {
	proposal, err := n.Committer.GetLastProposal(&bind.CallOpts{
		From: common.HexToAddress(n.Address),
	})
	if err != nil {
		return nil, fmt.Errorf("[QueryLastProposalID] err: %s", err)
	}
	return &proposal, nil
}

func (n NodeClient) QueryProposalByID(id uint64) (*contract.CommitterProposal, error) {
	proposal, err := n.Committer.Proposal(&bind.CallOpts{
		From: common.HexToAddress(n.Address),
	}, id)
	if err != nil {
		return nil, fmt.Errorf("[QueryProposalByID] err: %s", err)
	}
	return &proposal, nil
}

func (n NodeClient) SubmitProof(id uint64, proofHash string, stateRootHash string,
	startIndex uint64, endIndex uint64,
) (*types.Transaction, error) {
	tx, err := n.Committer.SubmitProof(&bind.TransactOpts{
		From:   n.Auth.From,
		Signer: n.Auth.Signer,
	}, id, proofHash, stateRootHash, startIndex, endIndex)
	if err != nil {
		return nil, fmt.Errorf("[SubmitProof] err: %s", err)
	}
	return tx, nil
}

func (n NodeClient) BitcoinTxHash(id uint64, txHash string) (*types.Transaction, error) {
	tx, err := n.Committer.BitcoinTx(&bind.TransactOpts{
		From:   n.Auth.From,
		Signer: n.Auth.Signer,
	}, id, txHash)
	if err != nil {
		return nil, fmt.Errorf("[BitcoinTxHash] err: %s", err)
	}
	return tx, nil
}

func (n NodeClient) TimeoutProposal(id uint64) (*types.Transaction, error) {
	tx, err := n.Committer.TimeoutProposal(&bind.TransactOpts{
		From:   n.Auth.From,
		Signer: n.Auth.Signer,
	}, id)
	if err != nil {
		return nil, fmt.Errorf("[TimeoutProposal] err: %s", err)
	}
	return tx, nil
}

//
//func (n NodeClient) GetAccountInfo(address string) (*eTypes.EthAccount, error) {
//	authClient := authTypes.NewQueryClient(n.GrpcConn)
//	res, err := authClient.Account(context.Background(), &authTypes.QueryAccountRequest{Address: address})
//	if err != nil {
//		return nil, fmt.Errorf("[NodeClient] GetAccountInfo err: %s", err)
//	}
//	ethAccount := &eTypes.EthAccount{}
//	err = ethAccount.Unmarshal(res.GetAccount().GetValue())
//	if err != nil {
//		return nil, fmt.Errorf("[NodeClient][ethAccount.Unmarshal] err: %s", err)
//	}
//	return ethAccount, nil
//}
//
//func (n NodeClient) SubmitProof(id uint64, from string, proofHash string, stateRootHash string,
//	startIndex uint64, endIndex uint64,
//) (uint64, error) {
//	msg := committerTypes.NewMsgSubmitProof(id, from, proofHash, stateRootHash, startIndex, endIndex)
//	msgResponse, err := n.broadcastTx(msg)
//	if err != nil {
//		return 0, fmt.Errorf("[SubmitProof] err: %s", err)
//	}
//	code := msgResponse.TxResponse.Code
//	rawLog := msgResponse.TxResponse.RawLog
//	if code != 0 {
//		return 0, fmt.Errorf("[SubmitProof][msgResponse.TxResponse.Code] err: %s", rawLog)
//	}
//	hexData := msgResponse.TxResponse.Data
//	byteData, err := hex.DecodeString(hexData)
//	if err != nil {
//		return 0, fmt.Errorf("[SubmitProof][hex.DecodeString] err: %s", err)
//	}
//	pbMsg := &sdk.TxMsgData{}
//	err = pbMsg.Unmarshal(byteData)
//	if err != nil {
//		return 0, fmt.Errorf("[SubmitProof][pbMsg.Unmarshal] err: %s", err)
//	}
//
//	resMsgRes := &committerTypes.MsgSubmitProofResponse{}
//	err = resMsgRes.Unmarshal(pbMsg.MsgResponses[0].GetValue())
//	if err != nil {
//		return 0, fmt.Errorf("[SubmitProof][resMsgRes.Unmarshal] err: %s", err)
//	}
//	return resMsgRes.Id, err
//}
//
//func (n NodeClient) BitcoinTx(proposalID uint64, from string, bitcoinTxHash string) (uint64, error) {
//	msg := committerTypes.NewMsgBitcoinTx(proposalID, from, bitcoinTxHash)
//	msgResponse, err := n.broadcastTx(msg)
//	if err != nil {
//		return 0, fmt.Errorf("[BitcoinTx] err: %s", err)
//	}
//	code := msgResponse.TxResponse.Code
//	rawLog := msgResponse.TxResponse.RawLog
//	if code != 0 {
//		return 0, fmt.Errorf("[BitcoinTx][msgResponse.TxResponse.Code] err: %s", rawLog)
//	}
//	hexData := msgResponse.TxResponse.Data
//	byteData, err := hex.DecodeString(hexData)
//	if err != nil {
//		return 0, fmt.Errorf("[BitcoinTx][hex.DecodeString] err: %s", err)
//	}
//	pbMsg := &sdk.TxMsgData{}
//	err = pbMsg.Unmarshal(byteData)
//	if err != nil {
//		return 0, fmt.Errorf("[BitcoinTx][pbMsg.Unmarshal] err: %s", err)
//	}
//
//	resMsgRes := &committerTypes.MsgSubmitProofResponse{}
//	err = resMsgRes.Unmarshal(pbMsg.MsgResponses[0].GetValue())
//	if err != nil {
//		return 0, fmt.Errorf("[BitcoinTx][resMsgRes.Unmarshal] err: %s", err)
//	}
//	return resMsgRes.Id, err
//}
//
//func (n NodeClient) GetEthGasPrice() (uint64, error) {
//	gasPriceByte, err := rpc.HTTPPostJSON("", n.RPCUrl, `{"jsonrpc":"2.0","method":"eth_gasPrice","params":[],"id":73}`)
//	if err != nil {
//		return 0, fmt.Errorf("[GetEthGasPrice] err: %s", err)
//	}
//	var g GasPriceRsp
//	if err := json.Unmarshal(gasPriceByte, &g); err != nil {
//		return 0, fmt.Errorf("[GetEthGasPrice.json.Unmarshal] err: %s", err)
//	}
//	parseUint, err := strconv.ParseUint(g.Result, 0, 64)
//	if err != nil {
//		return 0, fmt.Errorf("[GetEthGasPrice.strconv.ParseUint] err: %s", err)
//	}
//	return parseUint, nil
//}
//
//func (n NodeClient) GetGasPrice() (uint64, error) {
//	queryClient := feeTypes.NewQueryClient(n.GrpcConn)
//	res, err := queryClient.Params(context.Background(), &feeTypes.QueryParamsRequest{})
//	if err != nil {
//		return 0, fmt.Errorf("[GetGasPrice] err: %s", err)
//	}
//	return res.Params.BaseFee.Uint64(), nil
//}
//
//func (n NodeClient) broadcastTx(msgs ...sdk.Msg) (*tx.BroadcastTxResponse, error) {
//	gasPrice, err := n.GetGasPrice()
//	if err != nil {
//		return nil, fmt.Errorf("[broadcastTx][GetGasPrice] err: %s", err)
//	}
//	txBytes, err := n.buildSimTx(gasPrice, msgs...)
//	if err != nil {
//		return nil, fmt.Errorf("[SubmitProof] err: %s", err)
//	}
//	txClient := tx.NewServiceClient(n.GrpcConn)
//	res, err := txClient.BroadcastTx(context.Background(), &tx.BroadcastTxRequest{
//		Mode:    tx.BroadcastMode_BROADCAST_MODE_BLOCK,
//		TxBytes: txBytes,
//	})
//	if err != nil {
//		return nil, fmt.Errorf("[SubmitProof][BroadcastTx] err: %s", err)
//	}
//	return res, err
//}
//
//func (n NodeClient) buildSimTx(gasPrice uint64, msgs ...sdk.Msg) ([]byte, error) {
//	encCfg := simapp.MakeTestEncodingConfig()
//	txBuilder := encCfg.TxConfig.NewTxBuilder()
//	err := txBuilder.SetMsgs(msgs...)
//	if err != nil {
//		return nil, fmt.Errorf("[BuildSimTx][SetMsgs] err: %s", err)
//	}
//	ethAccount, err := n.GetAccountInfo(n.Address)
//	if nil != err {
//		return nil, fmt.Errorf("[BuildSimTx][GetAccountInfo]err: %s", err)
//	}
//	signV2 := signing.SignatureV2{
//		PubKey: n.PrivateKey.PubKey(),
//		Data: &signing.SingleSignatureData{
//			SignMode: encCfg.TxConfig.SignModeHandler().DefaultMode(),
//		},
//		Sequence: ethAccount.BaseAccount.Sequence,
//	}
//	err = txBuilder.SetSignatures(signV2)
//	if err != nil {
//		return nil, fmt.Errorf("[BuildSimTx][SetSignatures 1]err: %s", err)
//	}
//	txBuilder.SetGasLimit(DefaultBaseGasPrice)
//	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.Coin{
//		Denom:  n.Denom,
//		Amount: sdkmath.NewIntFromUint64(gasPrice * DefaultBaseGasPrice),
//	}))
//
//	signerData := xauthsigning.SignerData{
//		ChainID:       n.ChainID,
//		AccountNumber: ethAccount.BaseAccount.AccountNumber,
//		Sequence:      ethAccount.BaseAccount.Sequence,
//	}
//
//	sigV2, err := clientTx.SignWithPrivKey(
//		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
//		txBuilder, &n.PrivateKey, encCfg.TxConfig, ethAccount.BaseAccount.Sequence)
//	if err != nil {
//		return nil, fmt.Errorf("[BuildSimTx][SignWithPrivKey] err: %s", err)
//	}
//
//	err = txBuilder.SetSignatures(sigV2)
//	if err != nil {
//		return nil, fmt.Errorf("[BuildSimTx][SetSignatures 2] err: %s", err)
//	}
//	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
//	if err != nil {
//		return nil, fmt.Errorf("[BuildSimTx][GetTx] err: %s", err)
//	}
//	return txBytes, err
//}
//
//func (n NodeClient) QueryLastProposalID() (uint64, uint64, error) {
//	queryClient := committerTypes.NewQueryClient(n.GrpcConn)
//	res, err := queryClient.LastProposalID(context.Background(), &committerTypes.QueryLastProposalIdRequest{})
//	if err != nil {
//		return 0, 0, fmt.Errorf("[QueryLastProposalID] err: %s", err)
//	}
//	return res.LastProposalId, res.EndIndex, nil
//}
//
//func (n NodeClient) QueryProposalByID(id uint64) (*committerTypes.Proposal, error) {
//	queryClient := committerTypes.NewQueryClient(n.GrpcConn)
//	res, err := queryClient.Proposal(context.Background(), &committerTypes.QueryProposalRequest{ProposalId: id})
//	if err != nil {
//		return nil, fmt.Errorf("[QueryProposalByID] err: %s", err)
//	}
//	return res.Proposal, nil
//}
//
//func (n NodeClient) CommitterBitcoinTx(msg *committerTypes.MsgBitcoinTx) (*tx.BroadcastTxResponse, error) {
//	gasPrice, err := n.GetGasPrice()
//	if err != nil {
//		return nil, fmt.Errorf("[CommitterBitcoinTx][GetGasPrice] err: %s", err)
//	}
//	txBytes, err := n.buildSimTx(gasPrice, msg)
//	if err != nil {
//		return nil, fmt.Errorf("[SubmitProof] err: %s", err)
//	}
//	txClient := tx.NewServiceClient(n.GrpcConn)
//	res, err := txClient.BroadcastTx(context.Background(), &tx.BroadcastTxRequest{
//		Mode:    tx.BroadcastMode_BROADCAST_MODE_BLOCK,
//		TxBytes: txBytes,
//	})
//	if err != nil {
//		return nil, fmt.Errorf("[SubmitProof][BroadcastTx] err: %s", err)
//	}
//	return res, err
//}
//
//func (n NodeClient) TimeoutProposal(id uint64) error {
//	msg := &committerTypes.MsgTimeoutProposal{Id: id, From: n.Address}
//	_, err := n.broadcastTx(msg)
//	if err != nil {
//		return fmt.Errorf("[TimeoutProposal] err: %s", err)
//	}
//	return nil
//}
