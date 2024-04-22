package handler

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/internal/types"
	"github.com/b2network/b2committer/pkg/btcapi"
	"github.com/b2network/b2committer/pkg/contract/op"
	"github.com/b2network/b2committer/pkg/inscribe"
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
		voteAddress := ctx.B2NodeConfig.Address
		if lastProposal.Status == schema.ProposalSucceedStatus || lastProposal.ProposalID == 0 {
			log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] this proposal has been successful or just beginning : %d", latestProposalID)
			// submit new proposal
			newStateRootProposal, err := constructNextStateRootProposal(ctx, lastProposal)
			if err != nil {
				log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to construct new state root proposal: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			_, err = ctx.OpCommitterClient.SubmitStateRoot(newStateRootProposal)
			if err != nil {
				log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal]Try to submit new state root proposal: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] submit new state root proposal: %s", newStateRootProposal.ProposalID)
			time.Sleep(10 * time.Second)
			continue
		}

		if lastProposal.Status == schema.ProposalVotingStatus || lastProposal.Status == schema.ProposalTimeoutStatus {
			// check address voted or not
			phase, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnStateRootProposalPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
			if err != nil {
				log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to find address voted or not: %s", err)
				time.Sleep(3 * time.Second)
				continue
			}
			if phase {
				log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] address already voted in voting status: %s", voteAddress)
				continue
			}
			var voteProposalStartL1Timestamp uint64
			var voteProposalEndL1Timestamp uint64
			if lastProposal.ProposalID == 1 {
				voteProposalStartL1Timestamp = lastProposal.StartL1Timestamp
				voteProposalEndL1Timestamp = voteProposalStartL1Timestamp + ctx.Config.BlobIntervalTime
			} else {
				beforeLastProposal, err := ctx.OpCommitterClient.ProposalManager.GetTxsRootProposal(&bind.CallOpts{}, lastProposal.ProposalID-1)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal]Try to get before last proposal: %s", err.Error())
					time.Sleep(3 * time.Second)
					continue
				}
				voteProposalStartL1Timestamp = beforeLastProposal.EndTimestamp + 1
				voteProposalEndL1Timestamp = voteProposalStartL1Timestamp + ctx.Config.BlobIntervalTime
			}
			tsp, err := constructStateRootProposal(ctx, lastProposal.ProposalID, voteProposalStartL1Timestamp, voteProposalEndL1Timestamp)
			if err != nil {
				log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to construct new proposal to vote: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			_, err = ctx.OpCommitterClient.SubmitStateRoot(tsp)
			if err != nil {
				log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to submit new proposal to vote: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] vote txs proposal %s, %s ", tsp.ProposalID, voteAddress)
			continue
		}

		if lastProposal.Status == schema.ProposalPendingStatus {
			phase, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnStateRootDSTxPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
			if err != nil {
				log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal][IsVotedOnStateRootDSTxPhase] is failed : %s", err)
				time.Sleep(3 * time.Second)
				continue
			}
			if phase {
				log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] address already voted in pending status: %s", voteAddress)
				continue
			}
			if lastProposal.Winner == common.HexToAddress(ctx.B2NodeConfig.Address) {
				var events []schema.SyncEvent
				err = ctx.DB.Where("block_time between ? and ?", lastProposal.StartL1Timestamp, lastProposal.EndL1Timestamp).Order("block_number").Find(&events).Error
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal]collecting the state root blocks of proposal is failed. err : %s", errors.WithStack(err))
					continue
				}
				stateRoots, err := types.NewDsStateRootProposal(ctx.B2NodeConfig.ChainID, lastProposal.ProposalID, lastProposal.OutputRoot, events)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] constructing state root for ds is failed. err : %s", errors.WithStack(err))
				}
				dsJson, err := stateRoots.MarshalJson()
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to marshal ds proposal: %s", err.Error())
					time.Sleep(3 * time.Second)
					continue
				}

				dsTxID, err := ctx.DecentralizedStore.StoreDetailsOnChain(dsJson, ctx.B2NodeConfig.ChainID, lastProposal.ProposalID)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to store ds proposal: %s", err.Error())
					continue
				}
				_, err = ctx.OpCommitterClient.DsHash(lastProposal.ProposalID, schema.ProposalTypeStateRoot, schema.DsTypeArWeave, dsTxID)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to send ds proposal: %s", err.Error())
					continue
				}
				log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] success submit txs to ds: %s, dsHash: %s", lastProposal.ProposalID, lastProposal.DsTxHash)

			}
			if lastProposal.Winner != common.HexToAddress(voteAddress) {
				outputs, err := ctx.DecentralizedStore.QueryDetailsByTxID(lastProposal.DsTxHash)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to get outputs from ds: %s", err.Error())
					time.Sleep(3 * time.Second)
					continue
				}
				var dsProposal types.DsStateRootProposal
				err = json.Unmarshal(outputs, &dsProposal)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to unmarshal ds proposal: %s", err.Error())
					continue
				}

				events, err := types.ConvertOutputsToEventData(dsProposal.OutputEvents)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to convert outputs to event data: %s", err.Error())
					continue
				}
				verifyOutputRoots, err := GetOutputRootMerkleRoot(events)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to get outputs merkle root: %s", err.Error())
					continue
				}

				if verifyOutputRoots != lastProposal.OutputRoot {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to verify output from ds: %s", err.Error())
					continue
				}
				_, err = ctx.OpCommitterClient.DsHash(lastProposal.ProposalID, schema.ProposalTypeStateRoot, schema.DsTypeArWeave, lastProposal.DsTxHash)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to send ds proposal: %s", err.Error())
					continue
				}
				log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] success verify and vote submit output from ds: %s, dsHash: %s", lastProposal.ProposalID, lastProposal.DsTxHash)
			}

		}

		if lastProposal.Status == schema.ProposalCommitting {
			isVotedBtcTx, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnSubmitBitcoinTxPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
			if err != nil {
				log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal][IsVotedOnSubmitBitcoinTxPhase] is failed : %s", err)
				time.Sleep(3 * time.Second)
				continue
			}
			if isVotedBtcTx {
				log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] address already voted btc tx in committing status: %s", voteAddress)
				continue
			}
			if lastProposal.Winner == common.HexToAddress(voteAddress) {
				stateRoot := &types.StateRootProposal{
					ProposalID:         lastProposal.ProposalID,
					OutputRoot:         lastProposal.OutputRoot,
					StartL1Timestamp:   lastProposal.StartL1Timestamp,
					EndL1Timestamp:     lastProposal.EndL1Timestamp,
					StartL2BlockNumber: lastProposal.StartL2BlockNumber,
					EndL2BlockNumber:   lastProposal.EndL2BlockNumber,
					OutputStartIndex:   lastProposal.OutputStartIndex,
					OutputEndIndex:     lastProposal.OutputEndIndex,
				}
				btcStateRoot := &types.BtcStateRootProposal{
					Proposal: stateRoot,
					ChainID:  ctx.B2NodeConfig.ChainID,
				}
				content, err := json.Marshal(btcStateRoot)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Pending btcHash Try to marshal state root proposal: %s", err.Error())
					continue
				}
				rs, err := inscribe.Inscribe(ctx.BTCConfig.PrivateKey, content,
					ctx.BTCConfig.DestinationAddress, btcapi.ChainParams(ctx.BTCConfig.NetworkName))
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Inscribe state root err: %s\n", errors.WithStack(err).Error())
					continue
				}
				str, err := json.Marshal(rs)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Marshal result err: %s\n", errors.WithStack(err).Error())
					continue
				}
				log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] inscribe result: %s", str)
				bitcoinTxHash := rs.RevealTxHashList[0].String()
				_, err = ctx.OpCommitterClient.BitcoinTxHash(lastProposal.ProposalID, bitcoinTxHash)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to send bitcoin tx hash: %s", err.Error())
					continue
				}
				log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] success submit content to btc network. proposalID: %s, btcTxHash: %s", lastProposal.ProposalID, bitcoinTxHash)
			} else {
				outs, err := ctx.UnisatHTTPClient.QueryAPIBTCTxOutputsByTxID(context.Background(), lastProposal.BitcoinTxHash)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to get outputs from btc: %s", err.Error())
					time.Sleep(3 * time.Second)
					continue
				}
				if len(outs.Data) <= 0 {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to get outputs from btc: no data")
					time.Sleep(3 * time.Second)
					continue
				}
				if len(outs.Data[0].Inscriptions) <= 0 {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to get outputs from btc: no inscription")
					time.Sleep(3 * time.Second)
					continue
				}
				blockHeight := outs.Data[0].Height + 6
				if uint64(ctx.LatestBTCBlockNumber) < blockHeight {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to get outputs from btc: block height is too low")
					time.Sleep(3 * time.Second)
					continue
				}

				insID := outs.Data[0].Inscriptions[0].InscriptionID
				btcStateRootProposal, err := ctx.UnisatHTTPClient.QueryStateRootProposalByInsID(context.Background(), insID)
				if lastProposal.ProposalID != btcStateRootProposal.Proposal.ProposalID ||
					btcStateRootProposal.Proposal.OutputRoot != lastProposal.OutputRoot ||
					btcStateRootProposal.Proposal.StartL1Timestamp != lastProposal.StartL1Timestamp ||
					btcStateRootProposal.Proposal.EndL1Timestamp != lastProposal.EndL1Timestamp ||
					btcStateRootProposal.Proposal.StartL2BlockNumber != lastProposal.StartL2BlockNumber ||
					btcStateRootProposal.Proposal.EndL2BlockNumber != lastProposal.EndL2BlockNumber ||
					btcStateRootProposal.Proposal.OutputStartIndex != lastProposal.OutputStartIndex ||
					btcStateRootProposal.Proposal.OutputEndIndex != lastProposal.OutputEndIndex ||
					btcStateRootProposal.ChainID != ctx.B2NodeConfig.ChainID {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to verify btc state root proposal: %s", err.Error())
					time.Sleep(3 * time.Second)
					continue
				}

				_, err = ctx.OpCommitterClient.BitcoinTxHash(lastProposal.ProposalID, lastProposal.BitcoinTxHash)
				if err != nil {
					log.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to vote bitcoin tx hash: %s", err.Error())
					continue
				}
				log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] success verify and vote submit output from btc: %s, btcTxHash: %s", lastProposal.ProposalID, lastProposal.BitcoinTxHash)
			}

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
