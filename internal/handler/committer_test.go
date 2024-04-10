package handler

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/b2network/b2committer/pkg/contract"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestGetVerifyBatchesParamsByTxHash(t *testing.T) {
	//txHash := "0x05a4cf1c23eb2fea36a7d23dc0d419f35e27474f8217bf371ff8299502d99224"
	txHash := "0x442edca80b4882ef2cd3482e5cb95570a81d19fd8a03ed1e302a711fd7e994ed"
	// method := "VerifyBatchesTrustedAggregator"
	rpc, err := ethclient.Dial("https://ethereum-goerli-rpc.publicnode.com")
	tx, _, err := rpc.TransactionByHash(context.Background(), common.HexToHash(txHash))

	//inputsMap, methodName := DecodeTransactionInputData(abiObject, tx.Data())

	methodSigData := tx.Data()[:4]
	inputsSigData := tx.Data()[4:]

	abiObject, err := abi.JSON(strings.NewReader(contract.VerifyMetaData.ABI))
	require.NoError(t, err)
	method, err := abiObject.MethodById(methodSigData)
	require.NoError(t, err)

	inputsMap := make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(inputsMap, inputsSigData)
	require.NoError(t, err)
	initNumBatch := inputsMap["initNumBatch"].(uint64)
	finalNewBatch := inputsMap["finalNewBatch"].(uint64)
	e := inputsMap["newLocalExitRoot"].([32]byte)
	f := inputsMap["newStateRoot"].([32]byte)
	g := inputsMap["proof"].([24][32]byte)

	var result string
	for _, arr := range g {
		result += hex.EncodeToString(arr[:])
	}
	fmt.Println("initNumBatch", initNumBatch)
	fmt.Println("finalNewBatch", finalNewBatch)
	fmt.Println("newLocalExitRoot", hex.EncodeToString(e[:]))
	fmt.Println("newStateRoot outputs:", hex.EncodeToString(f[:]))
	fmt.Println("proof outputs: ", common.HexToHash(result))
}
