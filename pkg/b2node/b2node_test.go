package b2node

import (
	"encoding/hex"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/b2network/b2committer/internal/types"
	xcommitterTypes "github.com/evmos/ethermint/x/committer/types"
)

func TestGetAccountInfo(t *testing.T) {
	privateKeHex := "234def6d66bb45c288521bbbd20b7b3cd542e5b6e26386dcec24f8e918251f5a"
	chainID := "ethermint_9000-1"
	address := "ethm1jvqt5echmshc8gjsqdzk9unclt8qkx4knxcjdj"
	rpcUrl := "http://localhost:8545"
	grpcConn, err := types.GetClientConnection("127.0.0.1", types.WithClientPortOption(9090))
	if err != nil {
		panic(err)
	}
	nodeClient := NewNodeClient(privateKeHex, chainID, address, grpcConn, rpcUrl)
	addInfo, err := nodeClient.GetAccountInfo(address)
	//
	fmt.Println(addInfo.CodeHash)
	fmt.Println(addInfo.BaseAccount.Sequence)
	fmt.Println(addInfo.BaseAccount.Address)
}

func TestSubmitProof(t *testing.T) {
	privateKeHex := "234def6d66bb45c288521bbbd20b7b3cd542e5b6e26386dcec24f8e918251f5a"
	chainID := "ethermint_9000-1"
	address := "ethm1jvqt5echmshc8gjsqdzk9unclt8qkx4knxcjdj"
	rpcUrl := "http://localhost:8545"
	grpcConn, err := types.GetClientConnection("127.0.0.1", types.WithClientPortOption(9090))
	if err != nil {
		panic(err)
	}
	nodeClient := NewNodeClient(privateKeHex, chainID, address, grpcConn, rpcUrl)
	proposalID, err := nodeClient.SubmitProof(0, address, "proof7", "stateRoot", 1, 70)
	require.NoError(t, err)
	fmt.Println(proposalID)
}

func TestDecodeTxResponseData(t *testing.T) {
	byteData, _ := hex.DecodeString("12370A312F65766D6F732E65746865726D696E742E636F6D6D69747465722E4D73675375626D697450726F6F66526573706F6E736512020808")
	var pbMsg = &sdk.TxMsgData{}
	//proto.Marshal(&sdk.TxMsgData{MsgResponses: msgResponses})
	pbMsg.Unmarshal(byteData)
	fmt.Println(pbMsg.MsgResponses[0].TypeUrl)
	resMsgRes := &xcommitterTypes.MsgSubmitProofResponse{}
	resMsgRes.Unmarshal(pbMsg.MsgResponses[0].GetValue())
	fmt.Println(resMsgRes.Id)
}

func TestQueryLastProposalID(t *testing.T) {
	privateKeHex := "53da55319c649af5dec2d9ff11c0476698b27cf3bf8dfbce55fd29ab78caadf0"
	chainID := "ethermint_9000-1"
	address := "ethm17ezey9h6zw0yzaxq00w3gmt0rdet063v3vfmee"
	rpcUrl := "http://localhost:8545"
	grpcConn, err := types.GetClientConnection("127.0.0.1", types.WithClientPortOption(9090))
	if err != nil {
		panic(err)
	}
	nodeClient := NewNodeClient(privateKeHex, chainID, address, grpcConn, rpcUrl)
	lastID, endIndex, err := nodeClient.QueryLastProposalID()
	if err != nil {
		panic(err)
	}
	fmt.Println("lastID:", lastID)
	fmt.Println("index:", endIndex)
}

func TestQueryProposalByID(t *testing.T) {
	privateKeHex := "53da55319c649af5dec2d9ff11c0476698b27cf3bf8dfbce55fd29ab78caadf0"
	chainID := "ethermint_9000-1"
	address := "ethm17ezey9h6zw0yzaxq00w3gmt0rdet063v3vfmee"
	rpcUrl := "http://localhost:8545"
	grpcConn, err := types.GetClientConnection("127.0.0.1", types.WithClientPortOption(9090))
	if err != nil {
		panic(err)
	}
	nodeClient := NewNodeClient(privateKeHex, chainID, address, grpcConn, rpcUrl)
	proposal, err := nodeClient.QueryProposalByID(6)
	fmt.Println("id:", proposal.Id)
	fmt.Println("proposer:", proposal.Proposer)
	fmt.Println("status:", proposal.Status)
	fmt.Println("stateRootHash:", proposal.StateRootHash)
	fmt.Println("winner:", proposal.Winner)
	fmt.Println("voteList:", proposal.VotedListPhaseCommit)
	fmt.Println("start index:", proposal.StartIndex)
	fmt.Println("end index:", proposal.EndIndex)
	fmt.Println("bitcoinTx:", proposal.BitcoinTxHash)
}

func TestCommitterBitcoinTx(t *testing.T) {
	privateKeHex := "53da55319c649af5dec2d9ff11c0476698b27cf3bf8dfbce55fd29ab78caadf0"
	chainID := "ethermint_9000-1"
	address := "ethm17ezey9h6zw0yzaxq00w3gmt0rdet063v3vfmee"
	rpcUrl := "http://localhost:8545"
	grpcConn, err := types.GetClientConnection("127.0.0.1", types.WithClientPortOption(9090))
	if err != nil {
		panic(err)
	}
	nodeClient := NewNodeClient(privateKeHex, chainID, address, grpcConn, rpcUrl)
	res, err := nodeClient.CommitterBitcoinTx(&xcommitterTypes.MsgBitcoinTx{Id: 1, From: "ethm10ky5utnz5ddlmus5t2mm5ftxal3u0u6rsnx5nl", BitcoinTxHash: "1234567890"})
	require.NoError(t, err)
	fmt.Println(res)
}

func TestGetGasPrice(t *testing.T) {
	privateKeHex := "0c993419ff40521f20370c45721c92626c2f1fd35267258fb3d093ed0826b611"
	chainID := "ethermint_9000-1"
	address := "ethm1mffw0yzmusgm9fwd40jaal3vwustuhhx8rh03q"
	rpcUrl := "http://localhost:8545"
	grpcConn, err := types.GetClientConnection("127.0.0.1", types.WithClientPortOption(9090))
	if err != nil {
		panic(err)
	}
	nodeClient := NewNodeClient(privateKeHex, chainID, address, grpcConn, rpcUrl)
	gasprice, err := nodeClient.GetEthGasPrice()
	require.NoError(t, err)
	fmt.Println(gasprice)
}

func TestAddCommitter(t *testing.T) {
	privateKeHex := "0bca05e42968d16e52c1ea996068fa3cfa3a08a8b6afdf506f19c46b56caea04"
	chainID := "ethermint_9000-1"
	address := "ethm1jcf9exvr970jjc4efzsdh4y4pa698mmkpn5y6m"
	rpcUrl := "http://localhost:8545"
	grpcConn, err := types.GetClientConnection("127.0.0.1", types.WithClientPortOption(9090))
	if err != nil {
		panic(err)
	}
	nodeClient := NewNodeClient(privateKeHex, chainID, address, grpcConn, rpcUrl)
	add, err := nodeClient.AddCommitter("ethm1yz4g5svztygcvsen5whhlsa40uccwwtg3c9pdk")
	require.NoError(t, err)
	fmt.Println(add)
}
