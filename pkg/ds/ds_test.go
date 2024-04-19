package ds

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
	txs, _ := arClient.GetTransactionData("5m16jdjzJpWAOrCd83CM9izkn8md1IYKcsOK2SMESDU", "json")
	fmt.Println(string(txs))
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
