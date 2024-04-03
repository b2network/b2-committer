package handler

import (
	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"github.com/b2network/b2committer/internal/svc"
)

// SyncProposal sync proposal and process voting status
func SyncProposal(ctx *svc.ServiceContext) {
	time.Sleep(10 * time.Second)
	proposalID := ctx.Config.InitProposalID
	for {
		lastProposal, err := ctx.NodeClient.QueryLastProposal()
		if err != nil {
			log.Errorf("[Handler.Committer][QueryLastProposalID] error info: %s", errors.WithStack(err).Error())
			time.Sleep(3 * time.Second)
			continue
		}
		if lastProposal.Id < proposalID {
			log.Infof("[Handler.SyncProposal] Current proposalId is latest")
			time.Sleep(30 * time.Second)
			continue
		}

		proposal, err := ctx.NodeClient.QueryProposalByID(proposalID)
		if err != nil {
			log.Errorf("[Handler.SyncProposal] proposal QueryProposalByID error info:", errors.WithStack(err))
			time.Sleep(3 * time.Second)
			continue
		}
		if proposal == nil {
			log.Infof("[Handler.SyncProposal] proposal is nil", proposalID)
			proposalID++
			continue
		}
		var dbProposal schema.Proposal
		err = ctx.DB.Where("proposal_id = ?", proposalID).Order("proposal_id desc").First(&dbProposal).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("[Handler.SyncProposal] db query error info:", errors.WithStack(err))
			time.Sleep(3 * time.Second)
			continue
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			dbProposal = schema.Proposal{
				Base: schema.Base{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				ProposalID:    proposal.Id,
				StateRootHash: proposal.StateRootHash,
				ProofRootHash: proposal.ProofHash,
				StartBatchNum: proposal.StartIndex,
				EndBatchNum:   proposal.EndIndex,
				BtcTxHash:     proposal.BtcTxHash,
				Winner:        proposal.Winner.String(),
				Status:        uint64(proposal.Status),
				ArTxHash:      proposal.ArweaveTxHash,
			}
			err = ctx.DB.Create(&dbProposal).Error
			if err != nil {
				log.Errorf("[Handler.SyncProposal] db create error info: %s", errors.WithStack(err).Error())
			}
		}

		if dbProposal.ProposalID != 0 && dbProposal.Status != uint64(schema.ProposalVotingStatus) {
			log.Infof("[Handler.SyncProposal] already voted :", ctx.B2NodeConfig.Address)
			proposalID++
			continue
		}

		if proposal.Status == schema.ProposalVotingStatus {
			// voting
			verifyBatchInfo, err := GetVerifyBatchInfoByLastBatchNum(ctx, proposal.StartIndex)
			if err != nil {
				log.Errorf("[Handler.SyncProposal] GetVerifyBatchInfoByLastBatchNum error info:", errors.WithStack(err))
				time.Sleep(3 * time.Second)
				continue
			}

			_, err = ctx.NodeClient.SubmitProof(proposal.Id, verifyBatchInfo.proofRootHash, verifyBatchInfo.stateRootHash,
				verifyBatchInfo.startBatchNum, verifyBatchInfo.endBatchNum)
			if err != nil {
				log.Errorf("[Handler.SyncProposal] vote proposal error info", errors.WithStack(err))
				time.Sleep(3 * time.Second)
				continue
			}
			proposalID++
			continue
		}
		time.Sleep(10 * time.Second)
	}
}
