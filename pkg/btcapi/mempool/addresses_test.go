package mempool

import (
	"testing"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/require"
)

func TestListUnspent(t *testing.T) {
	// https://mempool.space/signet/api/address/tb1phjwg5zqzgf593vkc96fhtamp7v6wt6qjjuek2gygd8m8jqfnextsnpnsrr/utxo
	netParams := &chaincfg.SigNetParams
	client := NewClient(netParams)
	address, _ := btcutil.DecodeAddress("tb1phjwg5zqzgf593vkc96fhtamp7v6wt6qjjuek2gygd8m8jqfnextsnpnsrr", netParams)
	unspentList, err := client.ListUnspent(address)
	require.NoError(t, err)
	t.Log(len(unspentList))
	for _, output := range unspentList {
		t.Log(output.Outpoint.Hash.String(), "    ", output.Outpoint.Index)
	}
}
