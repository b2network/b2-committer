package handler

import (
	"time"

	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/pkg/errors"
)

// CheckStatus check proposal vote status
func CheckStatus(ctx *svc.ServiceContext) {
	for {
		var dbProposal schema.Proposal
		err := ctx.DB.Where("status=?", schema.VotingStatus).Order("end_batch_num asc").First(&dbProposal).Error
		if err != nil {
			log.Errorf("[Handler.CheckStatus] find Voting proposal err: %s\n", errors.WithStack(err))
			time.Sleep(5 * time.Second)
			continue
		}
		proposal, err := ctx.NodeClient.QueryProposalByID(dbProposal.ProposalID)
		if err != nil {
			log.Errorf("[Handler.CheckStatus] QueryProposalByID err: %s\n", errors.WithStack(err))
			continue
		}
		if proposal.Status != schema.VotingStatus && proposal.Winner != "" {
			dbProposal.Winner = proposal.Winner
			dbProposal.Status = proposal.Status
			dbProposal.BlockHeight = proposal.BlockHight
			ctx.DB.Save(dbProposal)
		}
		time.Sleep(3 * time.Second)
	}
}

func CheckStatusTimeOut(ctx *svc.ServiceContext) {
	for {
		time.Sleep(30 * time.Second)
		var dbProposal schema.Proposal
		err := ctx.DB.Where("status!=?", schema.SucceedStatus).Order("end_batch_num asc").First(&dbProposal).Error
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
		if proposal.BitcoinTxHash == "" && proposal.Status == schema.PendingStatus && proposal.Winner != ctx.B2NodeConfig.Address {
			num := uint64(ctx.LatestBlockNumber) - proposal.BlockHight
			if num > 10000 {
				err := ctx.NodeClient.TimeoutProposal(proposal.Id)
				if err != nil {
					log.Errorf("[Handler.CheckStatusTimeOut] TimeoutProposal err: %s\n", errors.WithStack(err))
					continue
				}
			}
			time.Sleep(2 * time.Second)
		}
	}
}
