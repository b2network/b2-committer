package handler

import (
	"time"

	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/pkg/errors"
)

// CheckStatusByContract sync check proposal status
func CheckStatusByContract(ctx *svc.ServiceContext) {
	for {
		proposal, err := ctx.NodeClient.QueryLastProposal()
		if err != nil {
			log.Errorf("[Handler.CheckStatusByContract][QueryLastProposalID] error info: %s", errors.WithStack(err).Error())
			continue
		}
		proposalContract, err := ctx.NodeClient.QueryProposalByID(proposal.Id)
		if err != nil {
			log.Errorf("[Handler.CheckStatusByContract][QueryProposalByID] error info: %s", errors.WithStack(err).Error())
			continue
		}
		var dbProposal schema.Proposal
		err = ctx.DB.Where("proposal_id=?", proposal.Id).First(&dbProposal).Error
		if err != nil {
			log.Errorf("[Handler.CheckStatusByContract] find proposal err: %s\n", errors.WithStack(err).Error())
			continue
		}
		dbProposal.BtcTxHash = proposalContract.BtcTxHash
		dbProposal.ArTxHash = proposalContract.ArweaveTxHash
		dbProposal.Winner = proposalContract.Winner.String()
		dbProposal.Status = uint64(proposalContract.Status)
		ctx.DB.Save(&dbProposal)
		time.Sleep(20 * time.Second)
	}

}
