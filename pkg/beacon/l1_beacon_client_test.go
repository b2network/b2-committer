package beacon

import (
	"context"
	"fmt"
	"github.com/b2network/b2committer/pkg/client"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

const (
	URL                  = "https://quaint-white-season.ethereum-sepolia.quiknode.pro/b5c30cbb548d8743f08dd175fe50e3e923259d30/"
	chainID              = 11155111
	SepoliaBatcherSender = "0x8F23BB38F531600e5d8FDDaAEC41F13FaB46E98c"
	SepoliaBatcherInbox  = "0xff00000000000000000000000000000011155420"
)

func TestBlockByNumber(t *testing.T) {
	l1Signer := types.NewCancunSigner(big.NewInt(chainID))
	fmt.Println(l1Signer.ChainID())
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()

	block, _ := conn.BlockByNumber(context.Background(), big.NewInt(5679002))
	fmt.Println(block.Number())
	tx1, _, _ := conn.TransactionByHash(context.Background(), common.HexToHash("0x3f343bbaa8dd0d76365fad90c87ef6df1fb66fa0342ace6f545aa3e2c8cf7c95"))
	fmt.Println(tx1.Hash())
	from, err := l1Signer.Sender(tx1)
	fmt.Println(from.Hex())
	blobIndex := 0 // index of each blob in the block's blob sidecar
	var hashes []eth.IndexedBlobHash
	for _, tx := range block.Transactions() {
		if !isValidBatchTx(tx, l1Signer, common.HexToAddress(SepoliaBatcherInbox), common.HexToAddress(SepoliaBatcherSender)) {
			blobIndex += len(tx.BlobHashes())
			continue
		}
		if tx.Type() != types.BlobTxType {
			continue
		}
		for _, h := range tx.BlobHashes() {
			idh := eth.IndexedBlobHash{
				Index: uint64(blobIndex),
				Hash:  h,
			}
			hashes = append(hashes, idh)
			blobIndex += 1
		}
	}
	fmt.Println(hashes)
}

func TestNodeVersion(t *testing.T) {
	log.Info("test beacon node...")
	l1Beacon := sources.NewBeaconHTTPClient(client.NewBasicHTTPClient(URL))
	version, _ := l1Beacon.NodeVersion(context.Background())
	fmt.Println(version)
}

func TestConfigSpec(t *testing.T) {
	l1Beacon := sources.NewBeaconHTTPClient(client.NewBasicHTTPClient(URL))
	config, _ := l1Beacon.ConfigSpec(context.Background())
	fmt.Println(config.Data.SecondsPerSlot)
}

func TestGetGenesis(t *testing.T) {
	//logger := testlog.Logger(t, 0).New("component", "beaconClient")
	l1Beacon := sources.NewBeaconHTTPClient(client.NewBasicHTTPClient(URL))
	config, _ := l1Beacon.BeaconGenesis(context.Background())
	fmt.Println(config.Data.GenesisTime)
}

func TestBeaconBlobSidecars(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()
	block, _ := conn.BlockByNumber(context.Background(), big.NewInt(5679002))
	fmt.Println(block.Number())
	txs := block.Transactions()
	l1Beacon := sources.NewBeaconHTTPClient(client.NewBasicHTTPClient(URL))
	l1BlobFetcher := sources.NewL1BeaconClient(l1Beacon, sources.L1BeaconClientConfig{FetchAllSidecars: false})
	l1Signer := types.NewCancunSigner(big.NewInt(chainID))
	hashes := DataAndHashesFromTxs(txs, l1Signer, common.HexToAddress(SepoliaBatcherInbox), common.HexToAddress(SepoliaBatcherSender))
	l1Block := eth.L1BlockRef{
		Hash:       block.Hash(),
		Number:     block.Number().Uint64(),
		ParentHash: block.ParentHash(),
		Time:       block.Time(),
	}
	for _, bHash := range hashes {
		fmt.Println(bHash)
	}
	blobs, err := l1BlobFetcher.GetBlobs(context.Background(), l1Block, hashes)
	fmt.Println(blobs)
}

func TestGetBlobByBlockNum(t *testing.T) {
	conn, err := ethclient.Dial(URL)
	require.NoError(t, err)
	defer conn.Close()
	l1Signer := types.NewCancunSigner(big.NewInt(chainID))
	l1Beacon := sources.NewBeaconHTTPClient(client.NewBasicHTTPClient(URL))
	l1BlobFetcher := sources.NewL1BeaconClient(l1Beacon, sources.L1BeaconClientConfig{FetchAllSidecars: false})
	bds := NewBlobDataSource(l1Signer, common.HexToAddress(SepoliaBatcherInbox), common.HexToAddress(SepoliaBatcherSender), l1BlobFetcher, conn)
	blockBlobInfo, err := bds.GetBlobByBlockNum(context.Background(), big.NewInt(5687502))
	fmt.Println(blockBlobInfo[0].Hash.Index)
	fmt.Println(blockBlobInfo[0].Hash.Hash)
	fmt.Println(blockBlobInfo[0].BlobSidecar.KZGCommitment)
}

func TestGetBlobByBlockNumMainNet(t *testing.T) {
	conn, err := ethclient.Dial("https://ethereum-rpc.publicnode.com")
	require.NoError(t, err)
	defer conn.Close()
	l1Signer := types.NewCancunSigner(big.NewInt(1))
	l1Beacon := sources.NewBeaconHTTPClient(client.NewBasicHTTPClient("https://ethereum-beacon-api.publicnode.com"))
	l1BlobFetcher := sources.NewL1BeaconClient(l1Beacon, sources.L1BeaconClientConfig{FetchAllSidecars: false})
	bds := NewBlobDataSource(l1Signer, common.HexToAddress("0xFF00000000000000000000000000000000000010"), common.HexToAddress("0x6887246668a3b87F54DeB3b94Ba47a6f63F32985"), l1BlobFetcher, conn)
	blockBlobInfo, err := bds.GetBlobByBlockNum(context.Background(), big.NewInt(19645654))
	require.NoError(t, err)
	if len(blockBlobInfo) != 0 {
		fmt.Println(blockBlobInfo[0].Hash.Index)
		fmt.Println(blockBlobInfo[0].Hash.Hash)
		fmt.Println(blockBlobInfo[0].BlobSidecar.KZGCommitment)
		fmt.Println(blockBlobInfo[0].BlobSidecar.Index)
	}
}
