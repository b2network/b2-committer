package handler

import (
	"encoding/hex"
	"encoding/json"
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

func GetStateRootAndCommitStateRootProposal(ctx *svc.ServiceContext) {
	// check address
	res, err := ctx.OpCommitterClient.Proposer.IsProposer(&bind.CallOpts{}, common.HexToAddress(ctx.B2NodeConfig.Address))
	if err != nil || !res {
		panic(err)
	}
	for {
		lastProposal, err := ctx.OpCommitterClient.ProposalManager.GetLastStateRootProposal(&bind.CallOpts{})
		if err != nil {
			log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to get last state root proposal from contract: %s", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		latestProposalID := lastProposal.ProposalID
		//voteAddress := ctx.B2NodeConfig.Address
		if lastProposal.Status == schema.ProposalSucceedStatus || lastProposal.ProposalID == 0 {
			log.Infof("this proposal has been successful or just beginning : %d", latestProposalID)
			// submit new proposal
			newStateRootProposal, err := constructNextStateRootProposal(ctx, lastProposal)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to construct new state root proposal: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			_, err = ctx.OpCommitterClient.SubmitStateRoot(newStateRootProposal)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to submit new state root proposal: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			log.Infof("[Handler.GetBlobsAndCommitProposal] submit new state root proposal: %s", newStateRootProposal.ProposalID)
			time.Sleep(10 * time.Second)
			continue
		}

		if lastProposal.Status == schema.ProposalVotingStatus || lastProposal.Status == schema.ProposalTimeoutStatus {

		}

	}

}

func constructNextStateRootProposal(ctx *svc.ServiceContext, lastProposal op.OpProposalStateRootProposal) (*types.StateRootProposal, error) {
	var event schema.SyncEvent
	var latestEndTimestamp uint64
	var latestStartTimestamp uint64
	if lastProposal.EndL1Timestamp == 0 {
		// contract has no one proposal
		err := ctx.DB.Where("event_name=?", "OutputProposed").Order("block_number").First(&event).Error
		if err != nil {
			return nil, fmt.Errorf("find original blob block error: %s", errors.WithStack(err))
		}
		latestStartTimestamp = uint64(event.BlockTime)
		latestEndTimestamp = uint64(event.BlockTime) + ctx.Config.OutputIntervalTime
	} else {
		latestStartTimestamp = lastProposal.EndL1Timestamp + 1 // plus 1 to exclude the last proposal end blob block
		latestEndTimestamp = lastProposal.EndL1Timestamp + ctx.Config.BlobIntervalTime
	}
	tsp, err := constructStateRootProposal(ctx, lastProposal.ProposalID+1, latestStartTimestamp, latestEndTimestamp)
	if err != nil {
		return nil, fmt.Errorf("construct txs root proposal failed, err: %s", errors.WithStack(err))
	}
	return tsp, nil
}

func constructStateRootProposal(ctx *svc.ServiceContext, proposalID uint64, startTimestamp uint64, endTimestamp uint64) (*types.StateRootProposal, error) {
	var event schema.SyncEvent
	err := ctx.DB.Where("block_time > ?", endTimestamp).Order("block_number").First(&event).Error
	if err != nil {
		return nil, fmt.Errorf("sync blob blocks is not completed: %s", errors.WithStack(err))
	}
	var events []schema.SyncEvent
	err = ctx.DB.Where("block_time between ? and ?", startTimestamp, endTimestamp).Order("block_number").Find(&events).Error
	if err != nil {
		return nil, fmt.Errorf("collecting the blob blocks of proposal is failed. err : %s", errors.WithStack(err))
	}
	outputRoots, err := GetOutputRootMerkleRoot(events)
	if err != nil {
		return nil, fmt.Errorf("getting the output merkle root is failed. err : %s", errors.WithStack(err))
	}
	stateRootProposal, err := types.NewStateRootProposal(proposalID, outputRoots, events)
	if err != nil {
		return nil, fmt.Errorf("constructing the state root proposal is failed. err : %s", errors.WithStack(err))
	}
	return stateRootProposal, nil
}

func GetOutputRootMerkleRoot(events []schema.SyncEvent) (string, error) {
	if len(events) == 0 {
		return "", fmt.Errorf("no output root")
	}
	var data types.OutputEvent
	if len(events) == 1 {
		err := json.Unmarshal([]byte(events[0].Data), &data)
		if err != nil {
			return "", fmt.Errorf("[GetOutputRootMerkleRoot] unmarshal output event failed %s", err.Error())
		}
		return data.OutputRoot, nil
	}
	newOutputRoots := make([]string, 0)
	for _, event := range events {
		err := json.Unmarshal([]byte(event.Data), &data)
		if err != nil {
			return "", fmt.Errorf("[GetOutputRootMerkleRoot] unmarshal output event failed %s", err.Error())
		}
		newOutputRoots = append(newOutputRoots, data.OutputRoot)
	}
	outputs := merkle.GenerateBlocks(newOutputRoots)
	outputsTree, _ := mt.New(nil, outputs)
	return hex.EncodeToString(outputsTree.Root), nil
}
