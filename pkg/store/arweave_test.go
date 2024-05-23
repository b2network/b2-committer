package store

import (
	"fmt"
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestQueryTxsByTxID(t *testing.T) {
	arClient := goar.NewClient("https://arweave.net")
	txs, _ := arClient.GetTransactionData("vKYQjEJIl2n8r6u_cmM5YMev8KAJ3DAwOMhsNI717Ws", "json")
	//require.NoError(t, err)
	fmt.Println(string(txs))

	//state, err := arClient.GetTransactionStatus("tB0zc2qLAGsw8x2kNUUPqzOBQuRV1C5CFWEN5tfY2CY")
	//require.NoError(t, err)
	//fmt.Println(state)
}

func TestSendData(t *testing.T) {
	tags := []types.Tag{
		{Name: "Content-Type", Value: "application/json"},
		{Name: "title", Value: "b2-batch"},
		{Name: "chainID", Value: strconv.FormatInt(123, 10)},
	}
	arNode := "https://arweave.net"
	wallet := "../../wallet/account.json"
	w, err := goar.NewWalletFromPath(wallet, arNode)
	require.NoError(t, err)
	arTx, err := w.SendData([]byte{}, tags)
	fmt.Println(arTx)
}
