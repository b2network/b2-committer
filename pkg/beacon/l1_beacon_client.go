package beacon

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/b2network/b2committer/pkg/errcode"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlobDataSource struct {
	L1Signer           types.Signer
	BatchInboxAddress  common.Address
	BatchSenderAddress common.Address
	//nolint:unused
	data         []BlockBlobInfo
	BlobsFetcher *sources.L1BeaconClient
	EthRPC       *ethclient.Client
}

func NewBlobDataSource(l1Signer types.Signer, batchInboxAddress common.Address, batchSenderAddress common.Address,
	blobsFetcher *sources.L1BeaconClient, ethRPC *ethclient.Client,
) *BlobDataSource {
	return &BlobDataSource{
		L1Signer:           l1Signer,
		BatchInboxAddress:  batchInboxAddress,
		BatchSenderAddress: batchSenderAddress,
		BlobsFetcher:       blobsFetcher,
		EthRPC:             ethRPC,
	}
}

func (bds *BlobDataSource) GetBlobByBlockNum(ctx context.Context, blockNum *big.Int) ([]BlockBlobInfo, error) {
	block, err := bds.EthRPC.BlockByNumber(ctx, blockNum)
	if err != nil {
		return nil, fmt.Errorf("failed to get block %d: %v", blockNum, err)
	}
	txs := block.Transactions()
	hashes := DataAndHashesFromTxs(txs, bds.L1Signer, bds.BatchInboxAddress, bds.BatchSenderAddress)
	if len(hashes) == 0 {
		// there are no blobs to fetch so we can return immediately

		return nil, errcode.ErrNoBlobFoundInBlock
	}
	l1Block := eth.L1BlockRef{
		Hash:       block.Hash(),
		Number:     block.Number().Uint64(),
		ParentHash: block.ParentHash(),
		Time:       block.Time(),
	}

	blobSidecars, err := bds.BlobsFetcher.GetBlobSidecars(context.Background(), l1Block, hashes)

	if errors.Is(err, ethereum.NotFound) {
		// If the L1 block was available, then the blobs should be available too. The only
		// exception is if the blob retention window has expired, which we will ultimately handle
		// by failing over to a blob archival service.
		return nil, fmt.Errorf("failed to fetch blobs: %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("failed to fetch blobs: %w", err)
	}
	//nolint:prealloc
	var data []BlockBlobInfo
	for i, hash := range hashes {
		if blobSidecars[i] == nil {
			log.Errorf("blob %d not found in block %d", hash.Index, blockNum)
			continue
		}
		//nolint:gosec,exportloopref
		data = append(data, BlockBlobInfo{
			BlobSidecar: blobSidecars[i],
			Hash:        &hash,
		})
	}
	return data, nil
}

// dataAndHashesFromTxs extracts calldata and datahashes from the input transactions and returns them. It
// creates a placeholder blobOrCalldata element for each returned blob hash that must be populated
// by fillBlobPointers after blob bodies are retrieved.
func DataAndHashesFromTxs(txs types.Transactions, l1Signer types.Signer, batchInboxAddress common.Address, batcherAddr common.Address) []eth.IndexedBlobHash {
	var hashes []eth.IndexedBlobHash
	blobIndex := 0 // index of each blob in the block's blob sidecar
	for _, tx := range txs {
		// skip any non-batcher transactions
		if !isValidBatchTx(tx, l1Signer, batchInboxAddress, batcherAddr) {
			continue
		}
		// handle non-blob batcher transactions by extracting their calldata
		if tx.Type() != types.BlobTxType {
			continue
		}
		// handle blob batcher transactions by extracting their blob hashes, ignoring any calldata.
		if len(tx.Data()) > 0 {
			log.Errorf("blob tx has calldata, which will be ignored %s, %s", "txhash", tx.Hash())
		}
		for _, h := range tx.BlobHashes() {
			idh := eth.IndexedBlobHash{
				Index: uint64(blobIndex),
				Hash:  h,
			}
			hashes = append(hashes, idh)
			blobIndex++
		}
	}
	return hashes
}

// isValidBatchTx returns true if:
//  1. the transaction has a To() address that matches the batch inbox address, and
//  2. the transaction has a valid signature from the batcher address
func isValidBatchTx(tx *types.Transaction, l1Signer types.Signer, batchInboxAddr, batcherAddr common.Address) bool {
	to := tx.To()
	if to == nil || *to != batchInboxAddr {
		return false
	}
	seqDataSubmitter, err := l1Signer.Sender(tx) // optimization: only derive sender if To is correct
	if err != nil {
		return false
	}
	// some random L1 user might have sent a transaction to our batch inbox, ignore them
	if seqDataSubmitter != batcherAddr {
		return false
	}
	return true
}

type BlockBlobInfo struct {
	BlobSidecar *eth.BlobSidecar
	Hash        *eth.IndexedBlobHash
}
