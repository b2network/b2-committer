package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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
	"time"
)

func GetBlobsAndCommitProposal(ctx *svc.ServiceContext) {
	// check address
	res, err := ctx.OpCommitterClient.Proposer.IsProposer(&bind.CallOpts{}, common.HexToAddress(ctx.B2NodeConfig.Address))
	if err != nil || !res {
		panic(ctx.B2NodeConfig.Address + " has no right to processing, please contact admin")
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
			log.Infof("this proposal has been successful or just beginning : %d", latestProposalID)
			latestProposalID = latestProposalID + 1
			// submit new proposal
			newTxsRootProposal, err := constructNewProposal(ctx, lastProposal)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to construct new proposal: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}

			_, err = ctx.OpCommitterClient.SubmitTxsRoot(newTxsRootProposal)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to submit new proposal: %s", err.Error())
			}
		}

		if lastProposal.Status == schema.ProposalVotingStatus {
			voteAddress := ctx.B2NodeConfig.Address
			// check address voted or not
			phase, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnTxsRootProposalPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(ctx.B2NodeConfig.Address))
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to find address voted or not: %s", err)
				time.Sleep(3 * time.Second)
				continue
			}
			if phase {
				log.Infof("[Handler.GetBlobsAndCommitProposal] address already voted: %s", voteAddress)
				continue
			}
			// vote proposal return
			_, err = verifyAndVotingProposal(ctx, lastProposal)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to find address voted or not: %s", err)
			}
		}

		if lastProposal.Status == schema.ProposalTimeoutStatus {
			proposal, err := resubmitTimeoutProposal(ctx, lastProposal)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to construct new proposal: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}

			_, err = ctx.OpCommitterClient.SubmitTxsRoot(proposal)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to submit new proposal: %s", err.Error())
			}
		}
		time.Sleep(30 * time.Second)
	}

}

func constructNewProposal(ctx *svc.ServiceContext, lastProposal op.OpProposalTxsRootProposal) (*types.TxsRootProposal, error) {
	var blob schema.BlobInfo
	var latestEndTimestamp uint64
	var latestStartTimestamp uint64
	if lastProposal.EndBlockNumber == 0 {
		// contract has no one proposal
		err := ctx.DB.Order("block_number").First(&blob).Error
		if err != nil {
			return nil, fmt.Errorf("find original blob block error: %s", errors.WithStack(err))
		}
		latestEndTimestamp = blob.BlockTime + ctx.Config.BlobIntervalTime
		latestStartTimestamp = blob.BlockTime
	} else {
		latestStartTimestamp = lastProposal.EndTimestamp + 1 // plus 1 to exclude the last proposal end blob block
		latestEndTimestamp = lastProposal.EndTimestamp + ctx.Config.BlobIntervalTime
	}

	err := ctx.DB.Where("block_time > ?", latestEndTimestamp).Order("block_number").First(&blob).Error
	if err != nil {
		return nil, fmt.Errorf("sync blob blocks is not completed: %s", errors.WithStack(err))
	}
	var blobs []schema.BlobInfo
	err = ctx.DB.Where("block_time between ? and ?", latestStartTimestamp, latestEndTimestamp).Order("block_number").Find(&blobs).Error
	if err != nil {
		return nil, fmt.Errorf("collecting the blob blocks of proposal is failed. err : %s", errors.WithStack(err))
	}
	blobMerkleRoot, err := GetBlobsMerkleRoot(blobs)
	return types.NewTxsRootProposal(lastProposal.ProposalID+1, blobMerkleRoot, blobs), nil
}

func verifyAndVotingProposal(ctx *svc.ServiceContext, lastProposal op.OpProposalTxsRootProposal) (bool, error) {
	var blob schema.BlobInfo
	err := ctx.DB.Where("block_time >= ?", lastProposal.EndTimestamp).Order("block_number").First(&blob).Error
	if err != nil {
		return false, fmt.Errorf("[verifyAndVotingProposal]sync blob blocks is not completed: %s", errors.WithStack(err))
	}
	var blobs []schema.BlobInfo
	err = ctx.DB.Where("block_time between ? and ?", lastProposal.StartTimestamp, lastProposal.EndTimestamp).Order("block_number").Find(&blobs).Error
	if err != nil {
		return false, fmt.Errorf("[verifyAndVotingProposal]collecting the blob blocks of proposal is failed. err : %s", errors.WithStack(err))
	}
	blobMerkleRoot, err := GetBlobsMerkleRoot(blobs)
	if blobMerkleRoot != lastProposal.TxsRoot {
		return false, fmt.Errorf("[verifyAndVotingProposal]blobMerkleRoot verify failed")
	}
	voteProposal := &types.TxsRootProposal{
		ProposalID:       lastProposal.ProposalID,
		StartBlockNumber: lastProposal.StartBlockNumber,
		EndBlockNumber:   lastProposal.EndBlockNumber,
		StartTimestamp:   lastProposal.StartTimestamp,
		EndTimestamp:     lastProposal.EndTimestamp,
		TxsRoot:          blobMerkleRoot,
		BlockList:        lastProposal.BlockList,
	}

	_, err = ctx.OpCommitterClient.SubmitTxsRoot(voteProposal)
	if err != nil {
		return false, fmt.Errorf("[verifyAndVotingProposal]vote proposal error: %s", err.Error())
	}
	return true, nil
}

func resubmitTimeoutProposal(ctx *svc.ServiceContext, lastProposal op.OpProposalTxsRootProposal) (*types.TxsRootProposal, error) {
	beforeLastProposal, err := ctx.OpCommitterClient.ProposalManager.GetTxsRootProposal(&bind.CallOpts{}, lastProposal.ProposalID-1)
	if err != nil {
		return nil, fmt.Errorf("[resubmitTimeoutProposal] Try to get last proposal from contract: %s", err.Error())
	}
	proposal, err := constructNewProposal(ctx, beforeLastProposal)
	if err != nil {
		return nil, fmt.Errorf("[resubmitTimeoutProposal] construct new proposal err: %s", err.Error())
	}
	return proposal, nil
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
