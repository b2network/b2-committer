package handler

import (
	"crypto/sha256"
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
		voteAddress := ctx.B2NodeConfig.Address
		if lastProposal.Status == schema.ProposalSucceedStatus || lastProposal.ProposalID == 0 {
			time.Sleep(30 * time.Second)
			log.Infof("this proposal has been successful or just beginning : %d", latestProposalID)
			// submit new proposal
			newTxsRootProposal, err := constructNextProposal(ctx, lastProposal)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to construct new proposal: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}

			_, err = ctx.OpCommitterClient.SubmitTxsRoot(newTxsRootProposal)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to submit new proposal: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			log.Infof("[Handler.GetBlobsAndCommitProposal] submit new txs proposal: %s", newTxsRootProposal.ProposalID)
			time.Sleep(10 * time.Second)
			continue
		}

		//nolint: dupl
		if lastProposal.Status == schema.ProposalVotingStatus || lastProposal.Status == schema.ProposalTimeoutStatus {
			// check address voted or not
			phase, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnTxsRootProposalPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to find address voted or not: %s", err)
				time.Sleep(3 * time.Second)
				continue
			}
			if phase {
				log.Infof("[Handler.GetBlobsAndCommitProposal] address already voted in voting status: %s", voteAddress)
				continue
			}
			var voteProposalStartTimestamp uint64
			var voteProposalEndTimestamp uint64
			if lastProposal.ProposalID == 1 {
				voteProposalStartTimestamp = lastProposal.StartTimestamp
				voteProposalEndTimestamp = voteProposalStartTimestamp + ctx.Config.BlobIntervalTime
			} else {
				beforeLastProposal, err := ctx.OpCommitterClient.ProposalManager.GetTxsRootProposal(&bind.CallOpts{}, lastProposal.ProposalID-1)
				if err != nil {
					log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to get before last proposal: %s", err.Error())
					time.Sleep(3 * time.Second)
					continue
				}
				voteProposalStartTimestamp = beforeLastProposal.EndTimestamp + 1
				voteProposalEndTimestamp = voteProposalStartTimestamp + ctx.Config.BlobIntervalTime
			}

			tsp, err := constructTxsRootProposal(ctx, lastProposal.ProposalID, voteProposalStartTimestamp, voteProposalEndTimestamp)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to construct new proposal to vote: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			_, err = ctx.OpCommitterClient.SubmitTxsRoot(tsp)
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to submit new proposal to vote: %s", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			log.Infof("[Handler.GetBlobsAndCommitProposal] vote txs proposal %s, %s ", tsp.ProposalID, voteAddress)
			time.Sleep(30 * time.Second)
		}

		if lastProposal.Status == schema.ProposalPendingStatus {
			phase, err := ctx.OpCommitterClient.ProposalManager.IsVotedOntxsRootDSTxPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
			if err != nil {
				log.Errorf("[Handler.GetBlobsAndCommitProposal][IsVotedOntxsRootDSTxPhase] is failed : %s", err)
				time.Sleep(3 * time.Second)
				continue
			}
			if phase {
				log.Infof("[Handler.GetBlobsAndCommitProposal] address already voted in pending status: %s", voteAddress)
				continue
			}
			if lastProposal.Winner == common.HexToAddress(ctx.B2NodeConfig.Address) {
				blobs, err := GetBlobsByBlockListFromDB(ctx, lastProposal.BlockList)
				if err != nil {
					log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to get blobs from db: %s", err.Error())
					time.Sleep(3 * time.Second)
					continue
				}
				blobMerkleRoot, err := GetBlobsMerkleRoot(blobs)
				if err != nil {
					log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to get blobs merkle root: %s", err.Error())
					time.Sleep(3 * time.Second)
					continue
				}
				dsProposal := types.NewDsTxsProposal(ctx.B2NodeConfig.ChainID, lastProposal.ProposalID, blobMerkleRoot, blobs)
				dsJSON, err := dsProposal.ToJSONBytes()
				if err != nil {
					log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to marshal ds proposal: %s", err.Error())
					time.Sleep(3 * time.Second)
					continue
				}
				dsTxID, err := ctx.DecentralizedStore.StoreDetailsOnChain(dsJSON, ctx.B2NodeConfig.ChainID, lastProposal.ProposalID)
				if err != nil {
					log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to store ds proposal: %s", err.Error())
					continue
				}
				log.Infof("[Handler.GetBlobsAndCommitProposal] proposal %s, success data to ds %s", lastProposal.ProposalID, dsTxID)
				_, err = ctx.OpCommitterClient.DsHash(lastProposal.ProposalID, schema.ProposalTypeTxsRoot, schema.DsTypeArWeave, dsTxID)
				if err != nil {
					log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to send ds proposal: %s", err.Error())
					continue
				}
				log.Infof("[Handler.GetBlobsAndCommitProposal] success submit txs to ds: %s, dsHash: %s", lastProposal.ProposalID, lastProposal.DsTxHash)
				time.Sleep(30 * time.Second)
			}
			if lastProposal.Winner != common.HexToAddress(ctx.B2NodeConfig.Address) {
				blobs, err := ctx.DecentralizedStore.QueryDetailsByTxID(lastProposal.DsTxHash)
				if err != nil {
					log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to get blobs from ds: %s", err.Error())
					time.Sleep(3 * time.Second)
					continue
				}
				var dsProposal types.DsTxsProposal
				err = json.Unmarshal(blobs, &dsProposal)
				if err != nil {
					log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to unmarshal ds proposal: %s", err.Error())
					continue
				}
				dbBlobs, err := dsProposal.GetDBBlobInfos()
				if err != nil {
					log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to get blobs from ds: %s", err.Error())
					continue
				}
				verifyTxsRootHash, err := GetBlobsMerkleRoot(dbBlobs)

				if verifyTxsRootHash != lastProposal.TxsRoot {
					log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to verify blobs from ds: %s", err.Error())
					continue
				}
				_, err = ctx.OpCommitterClient.DsHash(lastProposal.ProposalID, schema.ProposalTypeTxsRoot, schema.DsTypeArWeave, lastProposal.DsTxHash)
				if err != nil {
					log.Errorf("[Handler.GetBlobsAndCommitProposal] Try to send ds proposal: %s", err.Error())
					continue
				}
				log.Infof("[Handler.GetBlobsAndCommitProposal] success verify and vote submit txs from ds: %s, dsHash: %s", lastProposal.ProposalID, lastProposal.DsTxHash)
				time.Sleep(30 * time.Second)
			}
		}
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
