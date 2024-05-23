package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/internal/types"
	"github.com/b2network/b2committer/pkg/contract/op"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/b2network/b2committer/pkg/merkle"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	mt "github.com/txaty/go-merkletree"
)

func GetBlobsAndCommitTxsProposal(ctx *svc.ServiceContext) {
	// check address
	res, err := ctx.OpCommitterClient.Proposer.IsProposer(&bind.CallOpts{}, common.HexToAddress(ctx.B2NodeConfig.Address))
	if err != nil || !res {
		panic(err)
	}
	for {
		lastProposal, err := ctx.OpCommitterClient.ProposalManager.GetLastTxsRootProposal(&bind.CallOpts{})
		if err != nil {
			log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to get last proposal from contract: %s", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		latestProposalID := lastProposal.ProposalID
		if lastProposal.Status == schema.ProposalSucceedStatus || lastProposal.ProposalID == 0 {
			tx, newProposal, err := SubmitNextTxsProposal(ctx, lastProposal, latestProposalID)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal][SubmitNextTxsProposal] Try to submit proposal: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			log.Infof("[Handler.GetBlobsAndCommitProposal][SubmitNextTxsProposal] submit new proposal success. proposalID: %s, transaction: %s", newProposal.ProposalID, tx.Hash())
		}
		if lastProposal.Status == schema.ProposalVotingStatus || lastProposal.Status == schema.ProposalTimeoutStatus {
			err = ProcessTxsVotingAndTimeoutState(ctx, lastProposal)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal][ProcessVotingAndTimeoutState] Try to process voting and timeout state: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
		}
		if lastProposal.Status == schema.ProposalPendingStatus {
			err = ProcessTxsPendingStates(ctx, lastProposal)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal][ProcessPendingStates] Try to process pending state: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			log.Infof("[Handler.GetBlobsAndCommitProposal] success verify and vote submit txs from ds: %s, dsHash: %s", lastProposal.ProposalID, lastProposal.DsTxHash)
		}
		time.Sleep(15 * time.Second)
	}
}

func GetBlobsByBlockListFromDB(ctx *svc.ServiceContext, blockList []uint64) ([]schema.BlobInfo, error) {
	where := map[string]interface{}{}
	where["block_number"] = blockList

	var blobs []schema.BlobInfo
	err := ctx.DB.Where(where).Find(&blobs).Error
	if err != nil {
		return nil, fmt.Errorf("[GetBlobsByBlockListFromDB] Try to get blobs from db: %s", err.Error())
	}
	return blobs, nil
}

func constructTxsRootProposal(ctx *svc.ServiceContext, proposalID uint64, startTimestamp uint64, endTimestamp uint64) (*types.TxsRootProposal, error) {
	var blob schema.BlobInfo

	var blobs []schema.BlobInfo
	for {
		blob = schema.BlobInfo{}
		err := ctx.DB.Where("block_time > ?", endTimestamp).Order("block_number").First(&blob).Error
		if err != nil {
			return nil, fmt.Errorf("sync blob blocks is not completed: %s", errors.WithStack(err))
		}
		err = ctx.DB.Where("block_time between ? and ?", startTimestamp, endTimestamp).Order("block_number").Find(&blobs).Error
		if err != nil {
			return nil, fmt.Errorf("collecting the blob blocks of proposal is failed. err : %s", errors.WithStack(err))
		}
		if len(blobs) == 0 {
			endTimestamp += ctx.Config.BlobIntervalTime
		} else {
			break
		}
	}
	blobMerkleRoot, err := GetBlobsMerkleRoot(blobs)
	if err != nil {
		return nil, fmt.Errorf("getting the blob merkle root is failed. err : %s", errors.WithStack(err))
	}
	return types.NewTxsRootProposal(proposalID, blobMerkleRoot, blobs), nil
}

func constructNextProposal(ctx *svc.ServiceContext, lastProposal op.OpProposalTxsRootProposal) (*types.TxsRootProposal, error) {
	var blob schema.BlobInfo
	var latestEndTimestamp uint64
	var latestStartTimestamp uint64
	if lastProposal.EndBlockNumber == 0 {
		// contract has no one proposal
		err := ctx.DB.Order("block_number").First(&blob).Error
		if err != nil {
			return nil, fmt.Errorf("find original blob block error: %s", errors.WithStack(err))
		}
		latestStartTimestamp = blob.BlockTime
		latestEndTimestamp = blob.BlockTime + ctx.Config.BlobIntervalTime
	} else {
		latestStartTimestamp = lastProposal.EndTimestamp + 1 // plus 1 to exclude the last proposal end blob block
		latestEndTimestamp = lastProposal.EndTimestamp + ctx.Config.BlobIntervalTime
	}
	tsp, err := constructTxsRootProposal(ctx, lastProposal.ProposalID+1, latestStartTimestamp, latestEndTimestamp)
	if err != nil {
		return nil, fmt.Errorf("construct txs root proposal failed, err: %s", errors.WithStack(err))
	}
	return tsp, nil
}

func GetBlobsMerkleRoot(blobs []schema.BlobInfo) (string, error) {
	if len(blobs) == 0 {
		return "", fmt.Errorf("no blob data")
	}
	if len(blobs) == 1 {
		hash := sha256.Sum256([]byte(blobs[0].Blob))
		return hex.EncodeToString(hash[:]), nil
	}
	newBlobRoots := make([]string, 0)
	for _, blob := range blobs {
		blobHash := sha256.Sum256([]byte(blob.Blob))
		newBlobRoots = append(newBlobRoots, hex.EncodeToString(blobHash[:]))
	}
	blobBlocks := merkle.GenerateBlocks(newBlobRoots)
	blobTree, _ := mt.New(nil, blobBlocks)
	return hex.EncodeToString(blobTree.Root), nil
}
