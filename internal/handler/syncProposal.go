package handler

import (
	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/pkg/errors"
	"time"
)

// SyncProposal sync proposal and process voting status
func SyncProposal(ctx *svc.ServiceContext) {
	time.Sleep(10 * time.Second)
	proposalId := uint64(1) //自定义
	for {
		lastProposalID, _, err := ctx.NodeClient.QueryLastProposalID()
		if err != nil {
			log.Errorf("[Handler.Committer][QueryLastProposalID] error info:", errors.WithStack(err))
			time.Sleep(3 * time.Second)
			continue
		}
		if lastProposalID < proposalId {
			log.Infof("[Handler.SyncProposal] Current proposalId is latest")
			time.Sleep(30 * time.Second)
			continue
		}

		proposal, err := ctx.NodeClient.QueryProposalByID(proposalId)
		if err != nil {
			log.Errorf("[Handler.SyncProposal] proposal QueryProposalByID error info:", errors.WithStack(err))
			time.Sleep(3 * time.Second)
			continue
		}
		if proposal == nil {
			log.Infof("[Handler.SyncProposal] proposal is nil", proposalId)
			proposalId++
			continue
		}
		var dbProposal schema.Proposal
		err = ctx.DB.Where("proposal_id = ?", proposalId).Order("proposal_id desc").First(&dbProposal).Error
		if err != nil {
			log.Errorf("[Handler.SyncProposal] db query error info:", errors.WithStack(err))
			time.Sleep(3 * time.Second)
			continue
		}
		if dbProposal.ProposalID != 0 {
			log.Infof("[Handler.SyncProposal] already voted :", ctx.B2NodeConfig.Address)
			proposalId++
			continue
		}

		if proposal.Status == schema.VotingStatus {
			// voting
			verifyBatchInfo, err := GetVerifyBatchInfoByLastBatchNum(ctx, proposal.EndIndex)
			if err != nil {
				log.Errorf("[Handler.SyncProposal] GetVerifyBatchInfoByLastBatchNum error info:", errors.WithStack(err))
				time.Sleep(3 * time.Second)
				continue
			}

			_, err = ctx.NodeClient.SubmitProof(lastProposalID+1, ctx.B2NodeConfig.Address, verifyBatchInfo.proofRootHash, verifyBatchInfo.stateRootHash,
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
			proposalId++
			continue
		}

	}

}
