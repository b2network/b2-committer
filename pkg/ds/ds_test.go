package ds

import (
	"fmt"
	"github.com/everFinance/goar"
	"testing"
)

func TestQueryTxsByTxID(t *testing.T) {
	arClient := goar.NewClient("https://arweave.net")
	txs, _ := arClient.GetTransactionData("5m16jdjzJpWAOrCd83CM9izkn8md1IYKcsOK2SMESDU", "json")
	fmt.Println(string(txs))
}
