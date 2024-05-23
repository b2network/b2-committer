package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/internal/types"
	"github.com/b2network/b2committer/pkg/btcapi"
	"github.com/b2network/b2committer/pkg/contract/op"
	"github.com/b2network/b2committer/pkg/inscribe"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func SubmitNextStateRootProposal(ctx *svc.ServiceContext, lastProposal op.OpProposalStateRootProposal, latestProposalID uint64) (*ethTypes.Transaction, *types.StateRootProposal, error) {
	log.Infof("[SubmitNextStateRootProposal] this proposal has been successful or just beginning : %d", latestProposalID)
	// submit new proposal
	newStateRootProposal, err := constructNextStateRootProposal(ctx, lastProposal)
	if err != nil {
		return nil, nil, fmt.Errorf("[SubmitNextStateRootProposal] Try to construct new state root proposal: %s", err.Error())
	}
	tx, err := ctx.OpCommitterClient.SubmitStateRoot(newStateRootProposal)
	if err != nil {
		return nil, nil, fmt.Errorf("[SubmitNextStateRootProposal]Try to submit new state root proposal: %s, proposalID: %d", err.Error(), newStateRootProposal.ProposalID)
	}
	voteAddress := ctx.B2NodeConfig.Address
	err = confirmSubmitStateRootProposal(ctx, newStateRootProposal.ProposalID, voteAddress)
	if err != nil {
		return nil, nil, fmt.Errorf("[ProcessStateRootVotingAndTimeoutState] Try to confirm SubmitStateRoot failed: %s", err.Error())
	}
	log.Infof("[SubmitNextStateRootProposal] submit new state root proposal: %s", newStateRootProposal.ProposalID)
	return tx, newStateRootProposal, nil
}

func ProcessStateRootVotingAndTimeoutState(ctx *svc.ServiceContext, lastProposal op.OpProposalStateRootProposal) error {
	voteAddress := ctx.B2NodeConfig.Address
	// check address voted or not
	phase, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnStateRootProposalPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
	if err != nil {
		return fmt.Errorf("[ProcessStateRootVotingAndTimeoutState] Try to find address voted or not: %s", err)
	}
	if phase {
		return fmt.Errorf("[ProcessStateRootVotingAndTimeoutState] address already voted in voting status: %s", voteAddress)
	}
	var voteProposalStartL1Timestamp uint64
	var voteProposalEndL1Timestamp uint64
	if lastProposal.ProposalID == 1 {
		voteProposalStartL1Timestamp = lastProposal.StartL1Timestamp
		voteProposalEndL1Timestamp = voteProposalStartL1Timestamp + ctx.Config.OutputIntervalTime
	} else {
		beforeLastProposal, err := ctx.OpCommitterClient.ProposalManager.GetStateRootProposal(&bind.CallOpts{}, lastProposal.ProposalID-1)
		if err != nil {
			return fmt.Errorf("[ProcessStateRootVotingAndTimeoutState]Try to get before last proposal: %s", err.Error())
		}
		voteProposalStartL1Timestamp = beforeLastProposal.EndL1Timestamp + 1
		voteProposalEndL1Timestamp = voteProposalStartL1Timestamp + ctx.Config.OutputIntervalTime
	}
	tsp, err := constructStateRootProposal(ctx, lastProposal.ProposalID, voteProposalStartL1Timestamp, voteProposalEndL1Timestamp)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootVotingAndTimeoutState] Try to construct new proposal to vote: %s", err.Error())
	}
	_, err = ctx.OpCommitterClient.SubmitStateRoot(tsp)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootVotingAndTimeoutState] Try to submit new proposal to vote: %s", err.Error())
	}
	err = confirmSubmitStateRootProposal(ctx, tsp.ProposalID, voteAddress)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootVotingAndTimeoutState] Try to confirm SubmitStateRoot failed: %s", err.Error())
	}
	log.Infof("[ProcessStateRootVotingAndTimeoutState] vote txs proposal %s, %s ", tsp.ProposalID, voteAddress)
	return nil
}

func ProcessStateRootPendingStates(ctx *svc.ServiceContext, lastProposal op.OpProposalStateRootProposal) error {
	voteAddress := ctx.B2NodeConfig.Address
	phase, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnStateRootDSTxPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingStates][IsVotedOntxsRootDSTxPhase] is failed : %s", err)
	}
	if phase {
		return fmt.Errorf("[ProcessStateRootPendingStates][IsVotedOntxsRootDSTxPhase] address already voted in pending status: %s", voteAddress)
	}
	if lastProposal.Winner == common.HexToAddress(ctx.B2NodeConfig.Address) {
		err = ProcessStateRootPendingWinner(ctx, lastProposal, voteAddress)
		if err != nil {
			return fmt.Errorf("[ProcessStateRootPendingStates][ProcessStateRootPendingWinner] Try to process pending winner: %s", err.Error())
		}
	} else {
		err = ProcessStateRootPendingVoter(ctx, lastProposal, voteAddress)
		if err != nil {
			return fmt.Errorf("[ProcessStateRootPendingStates][ProcessStateRootPendingVoter] Try to process pending voter: %s", err.Error())
		}
	}
	return nil
}

func ProcessStateRootPendingWinner(ctx *svc.ServiceContext, lastProposal op.OpProposalStateRootProposal, voteAddress string) error {
	var events []schema.SyncEvent
	err := ctx.DB.Where("block_time between ? and ?", lastProposal.StartL1Timestamp, lastProposal.EndL1Timestamp).Order("block_number").Find(&events).Error
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingWinner]collecting the state root blocks of proposal is failed. err : %s", errors.WithStack(err))
	}
	stateRoots, err := types.NewDsStateRootProposal(ctx.B2NodeConfig.ChainID, lastProposal.ProposalID, lastProposal.OutputRoot, events)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingWinner] constructing state root for ds is failed. err : %s", errors.WithStack(err))
	}
	dsJSON, err := stateRoots.ToJSONBytes()
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingWinner] Try to marshal ds proposal: %s", err.Error())
	}

	dsTxID, err := ctx.DecentralizedStore.StoreDetailsOnChain(dsJSON, ctx.B2NodeConfig.ChainID, lastProposal.ProposalID)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingWinner] Try to store ds proposal: %s", err.Error())
	}
	err = confirmDSTransaction(ctx, dsTxID)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingWinner] Try to confirm ds tx: %s", err.Error())
	}
	log.Infof("[ProcessStateRootPendingWinner] proposal %s, success data to ds %s", lastProposal.ProposalID, dsTxID)
	_, err = ctx.OpCommitterClient.DsHash(lastProposal.ProposalID, schema.ProposalTypeStateRoot, schema.DsTypeArWeave, dsTxID)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingWinner] Try to send ds proposal by winner: %s", err.Error())
	}
	err = confirmStateRootDSTxPhase(ctx, lastProposal, voteAddress)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingWinner][confirmStateRootDSTxPhase] Try to confirm ds tx phase: %s", err.Error())
	}
	log.Infof("[ProcessStateRootPendingWinner] success submit txs to ds: %s, dsHash: %s", lastProposal.ProposalID, dsTxID)
	return nil
}

func ProcessStateRootPendingVoter(ctx *svc.ServiceContext, lastProposal op.OpProposalStateRootProposal, voteAddress string) error {
	outputs, err := ctx.DecentralizedStore.QueryDetailsByTxID(lastProposal.DsTxHash)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingVoter] Try to get outputs from ds: %s", err.Error())
	}
	var dsProposal types.DsStateRootProposal
	err = json.Unmarshal(outputs, &dsProposal)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingVoter] Try to unmarshal ds proposal: %s", err.Error())
	}

	events, err := types.ConvertOutputsToEventData(dsProposal.OutputEvents)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingVoter] Try to convert outputs to event data: %s", err.Error())
	}
	verifyOutputRoots, err := GetOutputRootMerkleRoot(events)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingVoter] Try to get outputs merkle root: %s", err.Error())
	}

	if verifyOutputRoots != lastProposal.OutputRoot {
		return fmt.Errorf("[ProcessStateRootPendingVoter] Try to verify output from ds: %s", err.Error())
	}
	_, err = ctx.OpCommitterClient.DsHash(lastProposal.ProposalID, schema.ProposalTypeStateRoot, schema.DsTypeArWeave, lastProposal.DsTxHash)
	if err != nil {
		return fmt.Errorf("[Handler.GetStateRootAndCommitStateRootProposal] Try to send ds proposal by voter: %s", err.Error())
	}
	err = confirmStateRootDSTxPhase(ctx, lastProposal, voteAddress)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootPendingVoter][confirmDSTxPhase] Try to confirm ds tx phase: %s", err.Error())
	}
	log.Infof("[ProcessStateRootPendingVoter] success verify and vote submit output from ds: %s, dsHash: %s", lastProposal.ProposalID, lastProposal.DsTxHash)
	return nil
}

func ProcessStateRootCommittingStates(ctx *svc.ServiceContext, lastProposal op.OpProposalStateRootProposal) error {
	voteAddress := ctx.B2NodeConfig.Address
	isVotedBtcTx, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnSubmitBitcoinTxPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
	if err != nil {
		return fmt.Errorf("[ProcessStateRootCommittingStates][IsVotedOnSubmitBitcoinTxPhase] is failed : %s", err)
	}
	if isVotedBtcTx {
		return fmt.Errorf("[ProcessStateRootCommittingStates] address already voted btc tx in committing status: %s, proposalID: %d", voteAddress, lastProposal.ProposalID)
	}
	if lastProposal.Winner == common.HexToAddress(ctx.B2NodeConfig.Address) {
		err = ProcessStateRootCommittingWinner(ctx, lastProposal, voteAddress)
		if err != nil {
			return fmt.Errorf("[ProcessStateRootPendingStates][ProcessStateRootPendingWinner] Try to process pending winner: %s", err.Error())
		}
	} else {
		err = ProcessStateRootCommittingVoter(ctx, lastProposal, voteAddress)
		if err != nil {
			return fmt.Errorf("[ProcessStateRootPendingStates][ProcessStateRootPendingVoter] Try to process pending voter: %s", err.Error())
		}
	}
	return nil
}

func ProcessStateRootCommittingWinner(ctx *svc.ServiceContext, lastProposal op.OpProposalStateRootProposal, voteAddress string) error {
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
		return fmt.Errorf("[ProcessStateRootCommittingWinner] Pending btcHash Try to marshal state root proposal: %s", err.Error())
	}
	rs, err := inscribe.Inscribe(ctx.BTCConfig.PrivateKey, content,
		ctx.BTCConfig.DestinationAddress, btcapi.ChainParams(ctx.BTCConfig.NetworkName))
	if err != nil {
		return fmt.Errorf("[ProcessStateRootCommittingWinner] Inscribe state root err: %s", errors.WithStack(err).Error())
	}
	str, err := json.Marshal(rs)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootCommittingWinner] Marshal result err: %s", errors.WithStack(err).Error())
	}
	log.Infof("[ProcessStateRootCommittingWinner] inscribe result: %s", str)
	bitcoinTxHash := rs.RevealTxHashList[0].String()
	_, err = ctx.OpCommitterClient.BitcoinTxHash(lastProposal.ProposalID, bitcoinTxHash)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootCommittingWinner] Try to send bitcoin tx hash: %s, winner:%s", err.Error(), lastProposal.Winner.String())
	}
	err = confirmBitcoinTxHashPhase(ctx, lastProposal, voteAddress)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootCommittingWinner][confirmBitcoinTxHashPhase] Try to confirm btc tx hash phase: %s", err.Error())
	}
	log.Infof("[Handler.GetStateRootAndCommitStateRootProposal] success submit content to btc network. proposalID: %s, btcTxHash: %s", lastProposal.ProposalID, bitcoinTxHash)
	return nil
}

func ProcessStateRootCommittingVoter(ctx *svc.ServiceContext, lastProposal op.OpProposalStateRootProposal, voteAddress string) error {
	outs, err := ctx.UnisatHTTPClient.QueryAPIBTCTxOutputsByTxID(context.Background(), lastProposal.BitcoinTxHash)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootCommittingVoter] Try to get outputs from btc: %s", err.Error())
	}
	if len(outs.Data) == 0 {
		return fmt.Errorf("[ProcessStateRootCommittingVoter] Try to get outputs from btc: no data")
	}
	if len(outs.Data[0].Inscriptions) == 0 {
		return fmt.Errorf("[ProcessStateRootCommittingVoter] Try to get outputs from btc: no inscription")
	}
	blockHeight := outs.Data[0].Height + 6
	if uint64(ctx.LatestBTCBlockNumber) < blockHeight {
		return fmt.Errorf("[ProcessStateRootCommittingVoter] Try to get outputs from btc: block height is too low")
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
		return fmt.Errorf("[ProcessStateRootCommittingVoter] Try to verify btc state root proposal: %s", err.Error())
	}

	_, err = ctx.OpCommitterClient.BitcoinTxHash(lastProposal.ProposalID, lastProposal.BitcoinTxHash)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootCommittingVoter] Try to vote bitcoin tx hash: %s", err.Error())
	}
	err = confirmBitcoinTxHashPhase(ctx, lastProposal, voteAddress)
	if err != nil {
		return fmt.Errorf("[ProcessStateRootCommittingVoter][confirmBitcoinTxHashPhase] Try to confirm btc tx hash phase: %s", err.Error())
	}
	log.Infof("[ProcessStateRootCommittingVoter] success verify and vote submit output from btc: %s, btcTxHash: %s", lastProposal.ProposalID, lastProposal.BitcoinTxHash)
	return nil
}

func confirmBitcoinTxHashPhase(ctx *svc.ServiceContext, lastProposal op.OpProposalStateRootProposal, voteAddress string) error {
	var times uint64
	for {
		if times > 60 {
			return fmt.Errorf("[confirmBitcoinTxHashPhase] confirmBitcoinTxHashPhase fail, proposalID: %d, voteAddress: %s", lastProposal.ProposalID, voteAddress)
		}
		confirmBTCTx, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnSubmitBitcoinTxPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
		if err != nil {
			fmt.Printf("[confirmBitcoinTxHashPhase] is failed : %s\n, proposalID: %d\n, voteAddress: %s\n\n", err, lastProposal.ProposalID, voteAddress)
			time.Sleep(60 * time.Second)
			times++
			continue
		}
		if confirmBTCTx {
			log.Infof("[confirmBitcoinTxHashPhase] confirm bitcoin tx hash has been processed : %s, proposal: %s", voteAddress, lastProposal.ProposalID)
			break
		}
		times++
	}
	return nil
}

func confirmStateRootDSTxPhase(ctx *svc.ServiceContext, lastProposal op.OpProposalStateRootProposal, voteAddress string) error {
	var times uint64
	for {
		if times > 60 {
			return fmt.Errorf("[confirmDSTxPhase] confirmDSTxPhase fail, proposalID: %d, voteAddress: %s", lastProposal.ProposalID, voteAddress)
		}
		res, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnStateRootDSTxPhase(&bind.CallOpts{}, lastProposal.ProposalID, common.HexToAddress(voteAddress))
		if err != nil {
			fmt.Printf("[confirmStateRootDSTxPhase] is failed : %s\n, proposalID: %d\n, voteAddress: %s\n\n", err, lastProposal.ProposalID, voteAddress)
			time.Sleep(60 * time.Second)
			times++
			continue
		}
		if res {
			log.Infof("[confirmStateRootDSTxPhase] confirm decentralized store has been processed : %s, proposal: %s", voteAddress, lastProposal.ProposalID)
			break
		}
		times++
	}
	return nil
}

func confirmSubmitStateRootProposal(ctx *svc.ServiceContext, proposalID uint64, voteAddress string) error {
	var times uint64
	for {
		if times > 60 {
			return fmt.Errorf("[confirmSubmitStateRootProposal] confirm SubmitStateRoot fail")
		}
		phase, err := ctx.OpCommitterClient.ProposalManager.IsVotedOnStateRootProposalPhase(&bind.CallOpts{}, proposalID, common.HexToAddress(voteAddress))
		if err != nil {
			fmt.Printf("[confirmSubmitStateRootProposal] Try to check submitTxProposal status fail: %s\n", err)
			time.Sleep(60 * time.Second)
			continue
		}
		if phase {
			log.Infof("[confirmSubmitStateRootProposal] proposalID: %s voteAddress: %s,  IsVotedOnTxsRootProposalPhase", proposalID, voteAddress)
			break
		}
		times++
	}
	return nil
}
