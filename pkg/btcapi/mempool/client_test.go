package mempool

import (
	"fmt"
	"log"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/stretchr/testify/require"
)

func TestGetRawTransaction(t *testing.T) {
	// https://mempool.space/signet/tx/2dd04235d653ed0cd59bb8aad741c55e0d171edd7a604a5bc62abfbc1fc3744c
	client := NewClient(&chaincfg.SigNetParams)
	txId, _ := chainhash.NewHashFromStr("2dd04235d653ed0cd59bb8aad741c55e0d171edd7a604a5bc62abfbc1fc3744c")
	log.Printf("result:" + txId.String())
	transaction, err := client.GetRawTransaction(txId)
	require.NoError(t, err)
	require.Equal(t, transaction.TxHash().String(), "2dd04235d653ed0cd59bb8aad741c55e0d171edd7a604a5bc62abfbc1fc3744c")
}

func TestGetTransactionByID(t *testing.T) {
	client := NewClient(&chaincfg.SigNetParams)
	txId := "0911087694a6001f8bf0c9bb0d2cd99ac975ef28d99ac1fe0c3746f834c74d8c"
	transaction, err := client.GetTransactionByID(txId)
	require.NoError(t, err)
	require.Equal(t, transaction.TxID, txId)
	require.Equal(t, transaction.Status.BlockHeight, int64(180550))
	require.Equal(t, transaction.Status.BlockHash, "00000053d10d115e512144bbeb005d408e7b2d4639487be5ea9cdbcf8cb5c992")
}

func TestGetCurrentBlockHash(t *testing.T) {
	client := NewClient(&chaincfg.SigNetParams)
	hash, err := client.GetCurrentBlockHash()
	require.NoError(t, err)
	fmt.Print(hash)
}
