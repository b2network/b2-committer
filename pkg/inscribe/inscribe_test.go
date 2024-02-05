package inscribe

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/stretchr/testify/require"
)

// stateRootHash: %s  a021b898b4c842021f7c5d2c25fbaeca94b9a9192ccc295ab519c212fd5c2200
// proofRootHash: %s  1492e66e89e186840231850712161255d203b5bbf48d21242f0b51519b5eb3d4

// f5d9aaa2c5ae280e153efcca934d9a022ad21dfd9ead2b203664f977dc90ec80
// tb1phjwg5zqzgf593vkc96fhtamp7v6wt6qjjuek2gygd8m8jqfnextsnpnsrr

// 522d135b91d13d0eba1b002f49a19acb51ec115fe9c8713319ceff9f351cd4f1
// tb1pydyjgpezqp2q98ux0dc283d3wvt6c4mluvwxfxrsqcle2jz8vwysep55n5
func TestInscribe(t *testing.T) {
	privateKeyHex := "f5d9aaa2c5ae280e153efcca934d9a022ad21dfd9ead2b203664f977dc90ec80"
	stateRootHash := "a021b898b4c842021f7c5d2c25fbaeca94b9a9192ccc295ab519c212fd5c2200"
	proofRootHash := "1492e66e89e186840231850712161255d203b5bbf48d21242f0b51519b5eb3d4"

	destinationPrivateKeyHex := "522d135b91d13d0eba1b002f49a19acb51ec115fe9c8713319ceff9f351cd4f1"
	destinationPrivateKeyBytes, _ := hex.DecodeString(destinationPrivateKeyHex)
	destinationPrivateKey, _ := btcec.PrivKeyFromBytes(destinationPrivateKeyBytes)
	destinationTaprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(destinationPrivateKey.PubKey())), &chaincfg.SigNetParams)
	rs, err := Inscribe(privateKeyHex, stateRootHash, proofRootHash, destinationTaprootAddress.EncodeAddress(), &chaincfg.SigNetParams)
	require.NoError(t, err)
	fmt.Println("destination address:" + destinationTaprootAddress.EncodeAddress())
	require.Equal(t, "tb1pydyjgpezqp2q98ux0dc283d3wvt6c4mluvwxfxrsqcle2jz8vwysep55n5", destinationTaprootAddress.EncodeAddress())

	fmt.Println("commitTxHash," + rs.CommitTxHash.String())
	for i := range rs.RevealTxHashList {
		fmt.Println("revealTxHash, " + rs.RevealTxHashList[i].String())
	}
	for i := range rs.Inscriptions {
		fmt.Println("inscription, " + rs.Inscriptions[i])
	}
	fmt.Println("fees: ", rs.Fees)
}
