package handler

import (
	"time"

	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/pkg/errors"
)

// CheckStatusVoting check proposal vote status
func CheckStatusVoting(ctx *svc.ServiceContext) {
	for {
		var dbProposal schema.Proposal
		err := ctx.DB.Where("status=?", schema.ProposalVotingStatus).Order("end_batch_num asc").First(&dbProposal).Error
		if err != nil {
			log.Errorf("[Handler.CheckStatus] find Voting proposal err: %s\n", errors.WithStack(err).Error())
			time.Sleep(5 * time.Second)
			continue
		}
		proposal, err := ctx.NodeClient.QueryProposalByID(dbProposal.ProposalID)
		if err != nil {
			log.Errorf("[Handler.CheckStatus] QueryProposalByID err: %s\n", errors.WithStack(err).Error())
			continue
		}
		if proposal.Status != schema.ProposalVotingStatus && proposal.Winner.String() != "" {
			dbProposal.Winner = proposal.Winner.String()
			dbProposal.Status = uint64(proposal.Status)
			ctx.DB.Save(dbProposal)
		}
		time.Sleep(3 * time.Second)
	}
}

func CheckStatusPending(ctx *svc.ServiceContext) {
	for {
		var dbProposal schema.Proposal
		err := ctx.DB.Where("status = ?", schema.ProposalPendingStatus).Order("end_batch_num asc").First(&dbProposal).Error
		if err != nil {
			log.Errorf("[Handler.CheckStatusTimeOut] find Voting proposal err: %s\n", errors.WithStack(err))
			time.Sleep(5 * time.Second)
			continue
		}
		proposal, err := ctx.NodeClient.QueryProposalByID(dbProposal.ProposalID)
		if err != nil {
			log.Errorf("[Handler.CheckStatusTimeOut] QueryProposalByID err: %s\n", errors.WithStack(err))
			continue
		}
		if proposal.Status == schema.ProposalCommitting {
			dbProposal.Status = uint64(proposal.Status)
			dbProposal.Winner = proposal.Winner.String()
			dbProposal.BtcTxHash = proposal.BtcTxHash
			ctx.DB.Save(dbProposal)
			continue
		}
		time.Sleep(2 * time.Second)
	}
}

func CheckStatusCommitting(ctx *svc.ServiceContext) {
	for {
		var dbProposal schema.Proposal
		err := ctx.DB.Where("status = ?", schema.ProposalCommitting).Order("end_batch_num asc").First(&dbProposal).Error
		if err != nil {
			log.Errorf("[Handler.CheckStatusTimeOut] find Voting proposal err: %s\n", errors.WithStack(err))
			time.Sleep(5 * time.Second)
			continue
		}
		proposal, err := ctx.NodeClient.QueryProposalByID(dbProposal.ProposalID)
		if err != nil {
			log.Errorf("[Handler.CheckStatusTimeOut] QueryProposalByID err: %s\n", errors.WithStack(err))
			continue
		}
		if proposal.Status == schema.ProposalSucceedStatus {
			dbProposal.Status = uint64(proposal.Status)
			dbProposal.Winner = proposal.Winner.String()
			dbProposal.ArTxHash = proposal.ArweaveTxHash
			ctx.DB.Save(dbProposal)
			continue
		}
		time.Sleep(2 * time.Second)
	}
}

func CheckStatusPendingTimeOut(ctx *svc.ServiceContext) {
	for {
		time.Sleep(30 * time.Second)
		var dbProposal schema.Proposal
		err := ctx.DB.Where("status = ? or status = ?", schema.ProposalPendingStatus,
			schema.ProposalCommitting).Order("end_batch_num asc").First(&dbProposal).Error
		if err != nil {
			log.Errorf("[Handler.CheckStatusTimeOut] find Voting proposal and Committing proposal err: %s\n", errors.WithStack(err).Error())
			time.Sleep(5 * time.Second)
			continue
		}
		proposal, err := ctx.NodeClient.QueryProposalByID(dbProposal.ProposalID)
		if err != nil {
			log.Errorf("[Handler.CheckStatusTimeOut] QueryProposalByID err: %s\n", errors.WithStack(err))
			continue
		}
		if (proposal.BtcTxHash == "" || proposal.ArweaveTxHash == "") && proposal.Winner.String() != ctx.B2NodeConfig.Address {
			res, err := ctx.NodeClient.IsProposalTimeout(proposal.Id)
			if err != nil {
				log.Errorf("[Handler.CheckStatusTimeOut] TimeoutProposal err: %s\n", errors.WithStack(err))
				continue
			}
			if res {
				dbProposal.Status = schema.ProposalTimeoutStatus
				ctx.DB.Save(dbProposal)
			}
			time.Sleep(2 * time.Second)
		}
	}
}
