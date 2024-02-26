package handler

import (
	"time"

	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// SyncProposal sync proposal and process voting status
func SyncProposal(ctx *svc.ServiceContext) {
	time.Sleep(10 * time.Second)
	proposalID := ctx.Config.InitProposalID
	for {
		lastProposalID, _, err := ctx.NodeClient.QueryLastProposalID()
		if err != nil {
			log.Errorf("[Handler.Committer][QueryLastProposalID] error info: %s", errors.WithStack(err).Error())
			time.Sleep(3 * time.Second)
			continue
		}
		if lastProposalID < proposalID {
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
		if dbProposal.ProposalID != 0 {
			log.Infof("[Handler.SyncProposal] already voted :", ctx.B2NodeConfig.Address)
			proposalID++
			continue
		}

		if proposal.Status == schema.VotingStatus {
			// voting
			verifyBatchInfo, err := GetVerifyBatchInfoByLastBatchNum(ctx, proposal.StartIndex)
			if err != nil {
				log.Errorf("[Handler.SyncProposal] GetVerifyBatchInfoByLastBatchNum error info:", errors.WithStack(err))
				time.Sleep(3 * time.Second)
				continue
			}

			_, err = ctx.NodeClient.SubmitProof(proposal.Id, ctx.B2NodeConfig.Address, verifyBatchInfo.proofRootHash, verifyBatchInfo.stateRootHash,
				verifyBatchInfo.startBatchNum, verifyBatchInfo.endBatchNum)
			if err != nil {
				log.Errorf("[Handler.SyncProposal] vote proposal error info", errors.WithStack(err))
				time.Sleep(3 * time.Second)
				continue
			}

			dbProposal := &schema.Proposal{
				Base: schema.Base{
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				EndBatchNum:   verifyBatchInfo.endBatchNum,
				ProposalID:    proposal.Id,
				Proposer:      proposal.Proposer,
				Status:        schema.VotingStatus,
				StateRootHash: verifyBatchInfo.stateRootHash,
				ProofRootHash: verifyBatchInfo.proofRootHash,
				StartBatchNum: verifyBatchInfo.startBatchNum,
			}

			// store db
			err = ctx.DB.Save(dbProposal).Error
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
