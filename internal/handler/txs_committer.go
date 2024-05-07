package handler

import (
	"encoding/json"
	"fmt"
	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/internal/types"
	"github.com/b2network/b2committer/pkg/contract/op"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"time"
)

func SubmitNextTxsProposal(ctx *svc.ServiceContext, lastProposal op.OpProposalTxsRootProposal, latestProposalID uint64) (*ethTypes.Transaction, *types.TxsRootProposal, error) {
	log.Infof("this proposal has been successful or just beginning : %d", latestProposalID)
	// submit new proposal
	newTxsRootProposal, err := constructNextProposal(ctx, lastProposal)
	if err != nil {
		return nil, nil, fmt.Errorf("[SubmitNextTxsProposal][GetBlobsAndCommitProposal] Try to construct new proposal: %s", err.Error())
	}

	trans, err := ctx.OpCommitterClient.SubmitTxsRoot(newTxsRootProposal)
	if err != nil {
		return nil, nil, fmt.Errorf("[SubmitNextTxsProposal][GetBlobsAndCommitProposal] Try to submit new proposal: %s", err.Error())
	}
	voteAddress := ctx.B2NodeConfig.Address
	err = confirmSubmitTxsProposal(ctx, newTxsRootProposal.ProposalID, voteAddress)
	if err != nil {
		return nil, nil, fmt.Errorf("[SubmitNextTxsProposal][confirmSubmitTxsProposal] Try to confirm submitTxProposal failed: %s", err.Error())
	}
	log.Infof("[SubmitNextTxsProposal] submit new txs proposal: %s", newTxsRootProposal.ProposalID)
	return trans, newTxsRootProposal, nil
}

func ProcessTxsVotingAndTimeoutState(ctx *svc.ServiceContext, lastProposal op.OpProposalTxsRootProposal) error {
	voteAddress := ctx.B2NodeConfig.Address
	// check address voted or not
	phase, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnTxsRootProposalPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
	if err != nil {
		return fmt.Errorf("[ProcessVotingAndTimeoutState] Try to find address voted or not in Voting and Timeout phase: %s", err)
	}
	if phase {
		return fmt.Errorf("[ProcessVotingAndTimeoutState] address already voted in voting status: %s", voteAddress)
	}
	var voteProposalStartTimestamp uint64
	var voteProposalEndTimestamp uint64
	if lastProposal.ProposalID == 1 {
		voteProposalStartTimestamp = lastProposal.StartTimestamp
		voteProposalEndTimestamp = voteProposalStartTimestamp + ctx.Config.BlobIntervalTime
	} else {
		beforeLastProposal, err := ctx.OpCommitterClient.ProposalManager.GetTxsRootProposal(&bind.CallOpts{}, lastProposal.ProposalID-1)
		if err != nil {
			return fmt.Errorf("[ProcessVotingAndTimeoutState] Try to get before last proposal: %s", err.Error())
		}
		voteProposalStartTimestamp = beforeLastProposal.EndTimestamp + 1
		voteProposalEndTimestamp = voteProposalStartTimestamp + ctx.Config.BlobIntervalTime
	}

	tsp, err := constructTxsRootProposal(ctx, lastProposal.ProposalID, voteProposalStartTimestamp, voteProposalEndTimestamp)
	if err != nil {
		return fmt.Errorf("[ProcessVotingAndTimeoutState] Try to construct new proposal to vote: %s", err.Error())
	}
	txs, err := ctx.OpCommitterClient.SubmitTxsRoot(tsp)
	if err != nil {
		return fmt.Errorf("[ProcessVotingAndTimeoutState] Try to submit new proposal to vote: %s", err.Error())
	}
	err = confirmSubmitTxsProposal(ctx, tsp.ProposalID, voteAddress)
	if err != nil {
		return fmt.Errorf("[ProcessVotingAndTimeoutState][confirmSubmitTxsProposal] Try to confirm submitTxProposal failed: %s", err.Error())
	}
	log.Infof("[ProcessVotingAndTimeoutState] vote txs proposal %s, %s, transaction: %s ", tsp.ProposalID, voteAddress, txs.Hash())
	return nil
}

func ProcessTxsPendingStates(ctx *svc.ServiceContext, lastProposal op.OpProposalTxsRootProposal) error {
	voteAddress := ctx.B2NodeConfig.Address
	phase, err := ctx.OpCommitterClient.ProposalManager.IsVotedOntxsRootDSTxPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
	if err != nil {
		return fmt.Errorf("[ProcessPendingStates][IsVotedOntxsRootDSTxPhase] is failed : %s", err)
	}
	if phase {
		return fmt.Errorf("[Handler.GetBlobsAndCommitProposal] address already voted in pending status: %s", voteAddress)
	}
	if lastProposal.Winner == common.HexToAddress(ctx.B2NodeConfig.Address) {
		err = ProcessTxsPendingWinner(ctx, lastProposal, voteAddress)
		if err != nil {
			return fmt.Errorf("[ProcessPendingStates][ProcessPendingWinner] Try to process pending winner: %s", err.Error())
		}
	} else {
		err = ProcessTxsPendingVoter(ctx, lastProposal, voteAddress)
		if err != nil {
			return fmt.Errorf("[ProcessPendingStates][ProcessPendingVoter] Try to process pending voter: %s", err.Error())
		}
	}
	return nil
}

func ProcessTxsPendingVoter(ctx *svc.ServiceContext, lastProposal op.OpProposalTxsRootProposal, voteAddress string) error {
	blobs, err := ctx.DecentralizedStore.QueryDetailsByTxID(lastProposal.DsTxHash)
	if err != nil {
		return fmt.Errorf("[ProcessPendingVoter] Try to get blobs from ds: %s", err.Error())
	}
	var dsProposal types.DsTxsProposal
	err = json.Unmarshal(blobs, &dsProposal)
	if err != nil {
		return fmt.Errorf("[ProcessPendingVoter] Try to unmarshal ds proposal: %s", err.Error())
	}
	dbBlobs, err := dsProposal.GetDBBlobInfos()
	if err != nil {
		return fmt.Errorf("[ProcessPendingVoter] Try to get blobs from ds: %s", err.Error())
	}
	verifyTxsRootHash, err := GetBlobsMerkleRoot(dbBlobs)

	if verifyTxsRootHash != lastProposal.TxsRoot {
		return fmt.Errorf("[ProcessPendingVoter] Try to verify blobs from ds: %s", err.Error())
	}
	_, err = ctx.OpCommitterClient.DsHash(lastProposal.ProposalID, schema.ProposalTypeTxsRoot, schema.DsTypeArWeave, lastProposal.DsTxHash)
	if err != nil {
		return fmt.Errorf("[ProcessPendingVoter] Try to send ds proposal: %s", err.Error())
	}
	err = confirmTxsDSTxPhase(ctx, lastProposal, voteAddress)
	if err != nil {
		return fmt.Errorf("[ProcessPendingVoter][confirmDSTxPhase] Try to confirm ds tx phase: %s", err.Error())
	}
	log.Infof("[ProcessPendingVoter] success verify and vote submit txs from ds: %s, dsHash: %s", lastProposal.ProposalID, lastProposal.DsTxHash)
	return nil
}

func ProcessTxsPendingWinner(ctx *svc.ServiceContext, lastProposal op.OpProposalTxsRootProposal, voteAddress string) error {
	blobs, err := GetBlobsByBlockListFromDB(ctx, lastProposal.BlockList)
	if err != nil {
		return fmt.Errorf("[ProcessPendingWinner][GetBlobsByBlockListFromDB] Try to get blobs from db: %s", err.Error())

	}
	blobMerkleRoot, err := GetBlobsMerkleRoot(blobs)
	if err != nil {
		return fmt.Errorf("[ProcessPendingWinner][GetBlobsMerkleRoot] Try to get blobs merkle root: %s", err.Error())
	}
	dsProposal := types.NewDsTxsProposal(ctx.B2NodeConfig.ChainID, lastProposal.ProposalID, blobMerkleRoot, blobs)
	dsJSON, err := dsProposal.ToJSONBytes()
	if err != nil {
		return fmt.Errorf("[ProcessPendingWinner] Try to marshal ds proposal: %s", err.Error())
	}
	// 重试机制
	dsTxID, err := ctx.DecentralizedStore.StoreDetailsOnChain(dsJSON, ctx.B2NodeConfig.ChainID, lastProposal.ProposalID)
	if err != nil {
		return fmt.Errorf("[StoreDetailsOnChain] Try to store ds proposal on decentralized store: %s", err.Error())
	}
	err = confirmDSTransaction(ctx, dsTxID)
	if err != nil {
		return fmt.Errorf("[confirmDSTransaction] Try to confirm ds tx: %s", err.Error())
	}
	log.Infof("[Handler.GetBlobsAndCommitProposal] proposal %s, success data to ds %s", lastProposal.ProposalID, dsTxID)
	_, err = ctx.OpCommitterClient.DsHash(lastProposal.ProposalID, schema.ProposalTypeTxsRoot, schema.DsTypeArWeave, dsTxID)
	if err != nil {
		return fmt.Errorf("[OpCommitterClient.DsHash] Try to send ds proposal: %s", err.Error())

	}
	err = confirmTxsDSTxPhase(ctx, lastProposal, voteAddress)
	if err != nil {
		return fmt.Errorf("[ProcessPendingWinner][confirmDSTxPhase] confirm ds tx phase %s", err.Error())
	}
	log.Infof("[Handler.GetBlobsAndCommitProposal] success submit txs to ds: %s, dsHash: %s", lastProposal.ProposalID, lastProposal.DsTxHash)
	return nil
}

func confirmSubmitTxsProposal(ctx *svc.ServiceContext, proposalID uint64, voteAddress string) error {
	var times uint64 = 0
	for {
		if times > 60 {
			return fmt.Errorf("[confirmSubmitTxsProposal] confirm submitTxsRoot fail")
		}
		phase, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnTxsRootProposalPhase(&bind.CallOpts{}, proposalID, common.HexToAddress(voteAddress))
		if err != nil {
			fmt.Printf("[confirmSubmitTxsProposal] Try to check submitTxProposal status fail: %s\n", err)
			time.Sleep(60 * time.Second)
			continue
		}
		if phase {
			log.Infof("[confirmSubmitTxsProposal] proposalID: %s voteAddress: %s,  IsVotedOnTxsRootProposalPhase", proposalID, voteAddress)
			break
		}
		times++
	}
	return nil
}

func confirmTxsDSTxPhase(ctx *svc.ServiceContext, lastProposal op.OpProposalTxsRootProposal, voteAddress string) error {
	var times uint64 = 0
	for {
		if times > 60 {
			return fmt.Errorf("[confirmDSTxPhase] confirmDSTxPhase fail, proposalID: %s, voteAddress: %s", lastProposal.ProposalID, voteAddress)
		}
		res, err := ctx.OpCommitterClient.ProposalManager.IsVotedOntxsRootDSTxPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
		if err != nil {
			fmt.Printf("[confirmDSTxPhase] is failed : %s\n, proposalID: %d\n, voteAddress: %s\n\n", err, lastProposal.ProposalID, voteAddress)
			time.Sleep(60 * time.Second)
			times++
			continue
		}
		if res {
			log.Infof("[confirmDSTxPhase] confirm decentralized store has been processed : %s, proposal: %s", voteAddress, lastProposal.ProposalID)
			break
		}
		times++
	}
	return nil
}

func confirmDSTransaction(ctx *svc.ServiceContext, dsTxID string) error {
	var times uint64 = 0
	for {
		if times > 60 {
			return fmt.Errorf("[confirmDSTransaction] confirmDSTransaction fail, dsTxID: %s", dsTxID)
		}
		bytes, err := ctx.DecentralizedStore.QueryDetailsByTxID(dsTxID)
		if err != nil {
			fmt.Printf("[confirmDSTransaction] Try to query details by txID: %s\n", errors.WithStack(err).Error())
			time.Sleep(60 * time.Second)
			times++
			continue
		}
		if len(bytes) != 0 {
			break
		}
		times++
	}
	return nil
}
