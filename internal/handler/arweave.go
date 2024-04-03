package handler

import (
	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"time"
)

func BatchDetailsToAr(ctx *svc.ServiceContext) {
	for {
		var dbProposal schema.Proposal
		err := ctx.DB.Where("status = ? and generate_details_file = ? and winner = ? and ar_tx_hash=''",
			schema.ProposalCommitting, true, ctx.B2NodeConfig.Address).Order("proposal_id asc").First(&dbProposal).Error
		if err != nil {
			log.Errorf("[Handler.BatchDetailsToAr] error info: %s", errors.WithStack(err).Error())
			time.Sleep(10 * time.Second)
			continue
		}
		proposal, err := ctx.NodeClient.QueryProposalByID(dbProposal.ProposalID)
		if proposal.Status == schema.ProposalSucceedStatus || proposal.ArweaveTxHash != "" {
			log.Errorf("[Handler.BatchDetailsToAr] batch details already upload")
			continue
		}

		arTx, err := uploadDetailToAR(ctx, dbProposal.StartBatchNum, dbProposal.EndBatchNum)
		if err != nil {
			log.Errorf("[Handler.BatchDetailsToAr] upload detail to ar error: %s", errors.WithStack(err).Error())
			continue
		}
		_, err = ctx.NodeClient.ArweaveTx(dbProposal.ProposalID, arTx)
		if err != nil {
			log.Errorf("[Handler.BatchDetailsToAr] get ar tx error: %s", errors.WithStack(err).Error())
			continue
		}
		err = ctx.DB.Model(&dbProposal).Update("ar_tx_hash", arTx).Error
		if err != nil {
			log.Errorf("[Handler.BatchDetailsToAr] update proposal error: %s", errors.WithStack(err).Error())
			continue
		}
		time.Sleep(3 * time.Second)
	}
}

func uploadDetailToAR(ctx *svc.ServiceContext, startBatchNum uint64, endBatchNum uint64) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Errorf("[WriteFile] get current path error: %s", errors.WithStack(err))
		return "", err
	}

	arNode := ctx.Config.ArweaveRPC
	wallet := ctx.Config.ArweaveWallet
	fileName := strconv.FormatUint(startBatchNum, 10) + "-" + strconv.FormatUint(endBatchNum, 10) + ".json"
	filePath := path + "/" + ctx.Config.BatchPath
	content, err := os.ReadFile(filePath + "/" + fileName)
	w, err := goar.NewWalletFromPath(wallet, arNode)
	if err != nil {
		log.Errorf("[WriteFile] new wallet error: %s", errors.WithStack(err))
		return "", err
	}
	tags := []types.Tag{
		{Name: "Content-Type", Value: "application/json"},
		{Name: "title", Value: "b2-batch"},
	}
	tx, err := w.SendData(content, tags)
	if err != nil {
		log.Errorf("[WriteFile] send data error: %s", errors.WithStack(err))
		return "", err
	}
	return tx.ID, err
}
