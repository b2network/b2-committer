package b2node

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/b2network/b2committer/pkg/rpc"

	sdkmath "cosmossdk.io/math"

	clientTx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	eTypes "github.com/evmos/ethermint/types"
	committerTypes "github.com/evmos/ethermint/x/committer/types"
	feeTypes "github.com/evmos/ethermint/x/feemarket/types"
	"google.golang.org/grpc"
)

const (
	DefaultBaseGasPrice = 10_000_000
)

type NodeClient struct {
	PrivateKey ethsecp256k1.PrivKey
	Address    string
	ChainID    string
	GrpcConn   *grpc.ClientConn
	RPCUrl     string
	Denom      string
}

type GasPriceRsp struct {
	ID      int64  `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

func NewNodeClient(privateKeyHex string, chainID string, address string, grpcConn *grpc.ClientConn,
	rpcURL string, denom string,
) *NodeClient {
	privatekeyBytes, err := hex.DecodeString(privateKeyHex)
	if nil != err {
		panic(err)
	}
	return &NodeClient{
		PrivateKey: ethsecp256k1.PrivKey{
			Key: privatekeyBytes,
		},
		Address:  address,
		ChainID:  chainID,
		GrpcConn: grpcConn,
		RPCUrl:   rpcURL,
		Denom:    denom,
	}
}

func (n NodeClient) AddCommitter(address string) (string, error) {
	msg := committerTypes.NewMsgAddCommitter(n.Address, address)
	msgResponse, err := n.broadcastTx(msg)
	if err != nil {
		return "", fmt.Errorf("[AddCommitter] err: %s", err)
	}
	code := msgResponse.TxResponse.Code
	rawLog := msgResponse.TxResponse.RawLog
	if code != 0 {
		return "", fmt.Errorf("[AddCommitter][msgResponse.TxResponse.Code] err: %s", rawLog)
	}
	hexData := msgResponse.TxResponse.Data
	byteData, err := hex.DecodeString(hexData)
	if err != nil {
		return "", fmt.Errorf("[AddCommitter][hex.DecodeString] err: %s", err)
	}
	pbMsg := &sdk.TxMsgData{}
	err = pbMsg.Unmarshal(byteData)
	if err != nil {
		return "", fmt.Errorf("[AddCommitter][pbMsg.Unmarshal] err: %s", err)
	}
	resMsgRes := &committerTypes.MsgAddCommitterResponse{}
	err = resMsgRes.Unmarshal(pbMsg.MsgResponses[0].GetValue())
	if err != nil {
		return "", fmt.Errorf("[AddCommitter][resMsgRes.Unmarshal] err: %s", err)
	}
	return resMsgRes.Committer, err
}

func (n NodeClient) GetAccountInfo(address string) (*eTypes.EthAccount, error) {
	authClient := authTypes.NewQueryClient(n.GrpcConn)
	res, err := authClient.Account(context.Background(), &authTypes.QueryAccountRequest{Address: address})
	if err != nil {
		return nil, fmt.Errorf("[NodeClient] GetAccountInfo err: %s", err)
	}
	ethAccount := &eTypes.EthAccount{}
	err = ethAccount.Unmarshal(res.GetAccount().GetValue())
	if err != nil {
		return nil, fmt.Errorf("[NodeClient][ethAccount.Unmarshal] err: %s", err)
	}
	return ethAccount, nil
}

func (n NodeClient) SubmitProof(id uint64, from string, proofHash string, stateRootHash string,
	startIndex uint64, endIndex uint64,
) (uint64, error) {
	msg := committerTypes.NewMsgSubmitProof(id, from, proofHash, stateRootHash, startIndex, endIndex)
	msgResponse, err := n.broadcastTx(msg)
	if err != nil {
		return 0, fmt.Errorf("[SubmitProof] err: %s", err)
	}
	code := msgResponse.TxResponse.Code
	rawLog := msgResponse.TxResponse.RawLog
	if code != 0 {
		return 0, fmt.Errorf("[SubmitProof][msgResponse.TxResponse.Code] err: %s", rawLog)
	}
	hexData := msgResponse.TxResponse.Data
	byteData, err := hex.DecodeString(hexData)
	if err != nil {
		return 0, fmt.Errorf("[SubmitProof][hex.DecodeString] err: %s", err)
	}
	pbMsg := &sdk.TxMsgData{}
	err = pbMsg.Unmarshal(byteData)
	if err != nil {
		return 0, fmt.Errorf("[SubmitProof][pbMsg.Unmarshal] err: %s", err)
	}

	resMsgRes := &committerTypes.MsgSubmitProofResponse{}
	err = resMsgRes.Unmarshal(pbMsg.MsgResponses[0].GetValue())
	if err != nil {
		return 0, fmt.Errorf("[SubmitProof][resMsgRes.Unmarshal] err: %s", err)
	}
	return resMsgRes.Id, err
}

func (n NodeClient) BitcoinTx(proposalID uint64, from string, bitcoinTxHash string) (uint64, error) {
	msg := committerTypes.NewMsgBitcoinTx(proposalID, from, bitcoinTxHash)
	msgResponse, err := n.broadcastTx(msg)
	if err != nil {
		return 0, fmt.Errorf("[BitcoinTx] err: %s", err)
	}
	code := msgResponse.TxResponse.Code
	rawLog := msgResponse.TxResponse.RawLog
	if code != 0 {
		return 0, fmt.Errorf("[BitcoinTx][msgResponse.TxResponse.Code] err: %s", rawLog)
	}
	hexData := msgResponse.TxResponse.Data
	byteData, err := hex.DecodeString(hexData)
	if err != nil {
		return 0, fmt.Errorf("[BitcoinTx][hex.DecodeString] err: %s", err)
	}
	pbMsg := &sdk.TxMsgData{}
	err = pbMsg.Unmarshal(byteData)
	if err != nil {
		return 0, fmt.Errorf("[BitcoinTx][pbMsg.Unmarshal] err: %s", err)
	}

	resMsgRes := &committerTypes.MsgSubmitProofResponse{}
	err = resMsgRes.Unmarshal(pbMsg.MsgResponses[0].GetValue())
	if err != nil {
		return 0, fmt.Errorf("[BitcoinTx][resMsgRes.Unmarshal] err: %s", err)
	}
	return resMsgRes.Id, err
}

func (n NodeClient) GetEthGasPrice() (uint64, error) {
	gasPriceByte, err := rpc.HTTPPostJSON("", n.RPCUrl, `{"jsonrpc":"2.0","method":"eth_gasPrice","params":[],"id":73}`)
	if err != nil {
		return 0, fmt.Errorf("[GetEthGasPrice] err: %s", err)
	}
	var g GasPriceRsp
	if err := json.Unmarshal(gasPriceByte, &g); err != nil {
		return 0, fmt.Errorf("[GetEthGasPrice.json.Unmarshal] err: %s", err)
	}
	parseUint, err := strconv.ParseUint(g.Result, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("[GetEthGasPrice.strconv.ParseUint] err: %s", err)
	}
	return parseUint, nil
}

func (n NodeClient) GetGasPrice() (uint64, error) {
	queryClient := feeTypes.NewQueryClient(n.GrpcConn)
	res, err := queryClient.Params(context.Background(), &feeTypes.QueryParamsRequest{})
	if err != nil {
		return 0, fmt.Errorf("[GetGasPrice] err: %s", err)
	}
	return res.Params.BaseFee.Uint64(), nil
}

func (n NodeClient) broadcastTx(msgs ...sdk.Msg) (*tx.BroadcastTxResponse, error) {
	gasPrice, err := n.GetGasPrice()
	if err != nil {
		return nil, fmt.Errorf("[broadcastTx][GetGasPrice] err: %s", err)
	}
	txBytes, err := n.buildSimTx(gasPrice, msgs...)
	if err != nil {
		return nil, fmt.Errorf("[SubmitProof] err: %s", err)
	}
	txClient := tx.NewServiceClient(n.GrpcConn)
	res, err := txClient.BroadcastTx(context.Background(), &tx.BroadcastTxRequest{
		Mode:    tx.BroadcastMode_BROADCAST_MODE_BLOCK,
		TxBytes: txBytes,
	})
	if err != nil {
		return nil, fmt.Errorf("[SubmitProof][BroadcastTx] err: %s", err)
	}
	return res, err
}

func (n NodeClient) buildSimTx(gasPrice uint64, msgs ...sdk.Msg) ([]byte, error) {
	encCfg := simapp.MakeTestEncodingConfig()
	txBuilder := encCfg.TxConfig.NewTxBuilder()
	err := txBuilder.SetMsgs(msgs...)
	if err != nil {
		return nil, fmt.Errorf("[BuildSimTx][SetMsgs] err: %s", err)
	}
	ethAccount, err := n.GetAccountInfo(n.Address)
	if nil != err {
		return nil, fmt.Errorf("[BuildSimTx][GetAccountInfo]err: %s", err)
	}
	signV2 := signing.SignatureV2{
		PubKey: n.PrivateKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode: encCfg.TxConfig.SignModeHandler().DefaultMode(),
		},
		Sequence: ethAccount.BaseAccount.Sequence,
	}
	err = txBuilder.SetSignatures(signV2)
	if err != nil {
		return nil, fmt.Errorf("[BuildSimTx][SetSignatures 1]err: %s", err)
	}
	txBuilder.SetGasLimit(DefaultBaseGasPrice)
	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.Coin{
		Denom:  n.Denom,
		Amount: sdkmath.NewIntFromUint64(gasPrice * DefaultBaseGasPrice),
	}))

	signerData := xauthsigning.SignerData{
		ChainID:       n.ChainID,
		AccountNumber: ethAccount.BaseAccount.AccountNumber,
		Sequence:      ethAccount.BaseAccount.Sequence,
	}

	sigV2, err := clientTx.SignWithPrivKey(
		encCfg.TxConfig.SignModeHandler().DefaultMode(), signerData,
		txBuilder, &n.PrivateKey, encCfg.TxConfig, ethAccount.BaseAccount.Sequence)
	if err != nil {
		return nil, fmt.Errorf("[BuildSimTx][SignWithPrivKey] err: %s", err)
	}

	err = txBuilder.SetSignatures(sigV2)
	if err != nil {
		return nil, fmt.Errorf("[BuildSimTx][SetSignatures 2] err: %s", err)
	}
	txBytes, err := encCfg.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, fmt.Errorf("[BuildSimTx][GetTx] err: %s", err)
	}
	return txBytes, err
}

func (n NodeClient) QueryLastProposalID() (uint64, uint64, error) {
	queryClient := committerTypes.NewQueryClient(n.GrpcConn)
	res, err := queryClient.LastProposalID(context.Background(), &committerTypes.QueryLastProposalIdRequest{})
	if err != nil {
		return 0, 0, fmt.Errorf("[QueryLastProposalID] err: %s", err)
	}
	return res.LastProposalId, res.EndIndex, nil
}

func (n NodeClient) QueryProposalByID(id uint64) (*committerTypes.Proposal, error) {
	queryClient := committerTypes.NewQueryClient(n.GrpcConn)
	res, err := queryClient.Proposal(context.Background(), &committerTypes.QueryProposalRequest{ProposalId: id})
	if err != nil {
		return nil, fmt.Errorf("[QueryProposalByID] err: %s", err)
	}
	return res.Proposal, nil
}

func (n NodeClient) CommitterBitcoinTx(msg *committerTypes.MsgBitcoinTx) (*tx.BroadcastTxResponse, error) {
	gasPrice, err := n.GetGasPrice()
	if err != nil {
		return nil, fmt.Errorf("[CommitterBitcoinTx][GetGasPrice] err: %s", err)
	}
	txBytes, err := n.buildSimTx(gasPrice, msg)
	if err != nil {
		return nil, fmt.Errorf("[SubmitProof] err: %s", err)
	}
	txClient := tx.NewServiceClient(n.GrpcConn)
	res, err := txClient.BroadcastTx(context.Background(), &tx.BroadcastTxRequest{
		Mode:    tx.BroadcastMode_BROADCAST_MODE_BLOCK,
		TxBytes: txBytes,
	})
	if err != nil {
		return nil, fmt.Errorf("[SubmitProof][BroadcastTx] err: %s", err)
	}
	return res, err
}

func (n NodeClient) TimeoutProposal(id uint64) error {
	msg := &committerTypes.MsgTimeoutProposal{Id: id, From: n.Address}
	_, err := n.broadcastTx(msg)
	if err != nil {
		return fmt.Errorf("[TimeoutProposal] err: %s", err)
	}
	return nil
}
