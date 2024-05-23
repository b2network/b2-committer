package handler

import (
	"encoding/hex"
	"encoding/json"
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
		if lastProposal.Status == schema.ProposalSucceedStatus || lastProposal.ProposalID == 0 {
			tx, newStateRootProposal, err := SubmitNextStateRootProposal(ctx, lastProposal, latestProposalID)
			if err != nil {
				log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to submit proposal: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] submit new proposal success. proposalID: %s, transaction: %s", newStateRootProposal.ProposalID, tx.Hash())
		}

		if lastProposal.Status == schema.ProposalVotingStatus || lastProposal.Status == schema.ProposalTimeoutStatus {
			err = ProcessStateRootVotingAndTimeoutState(ctx, lastProposal)
			if err != nil {
				log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal][ProcessStateRootVotingAndTimeoutState] Try to process voting and timeout state: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
		}

		if lastProposal.Status == schema.ProposalPendingStatus {
			err = ProcessStateRootPendingStates(ctx, lastProposal)
			if err != nil {
				log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal][ProcessStateRootPendingStates] Try to process pending state: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
		}

		if lastProposal.Status == schema.ProposalCommitting {
			err = ProcessStateRootCommittingStates(ctx, lastProposal)
			if err != nil {
				log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal][ProcessStateRootCommittingStates] Try to process committing state: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
		}
		time.Sleep(15 * time.Second)
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
	var events []schema.SyncEvent
	for {
		event = schema.SyncEvent{}
		err := ctx.DB.Where("block_time > ?", endTimestamp).Order("block_number").First(&event).Error
		if err != nil {
			return nil, fmt.Errorf("sync blob blocks is not completed: %s", errors.WithStack(err))
		}
		err = ctx.DB.Where("block_time between ? and ?", startTimestamp, endTimestamp).Order("block_number").Find(&events).Error
		if err != nil {
			return nil, fmt.Errorf("collecting the blob blocks of proposal is failed. err : %s", errors.WithStack(err))
		}
		if len(events) == 0 {
			endTimestamp += ctx.Config.OutputIntervalTime
		} else {
			break
		}
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
