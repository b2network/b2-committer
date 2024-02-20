package handler

import (
	"encoding/json"
	"time"

	"github.com/b2network/b2committer/pkg/btcapi"
	btcmempool "github.com/b2network/b2committer/pkg/btcapi/mempool"

	"github.com/b2network/b2committer/pkg/inscribe"

	"github.com/b2network/b2committer/pkg/log"

	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/pkg/errors"
)

// Inscribe check proposal statues. process pending proposal.
func Inscribe(ctx *svc.ServiceContext) {
	time.Sleep(30 * time.Second)
	for {
		var dbProposal schema.Proposal
		err := ctx.DB.Where("status=?", schema.PendingStatus).Order("end_batch_num asc").First(&dbProposal).Error
		if err != nil {
			log.Errorf("[Handler.Inscribe] Pending and timeout proposal err: %s\n", errors.WithStack(err))
			time.Sleep(5 * time.Second)
			continue
		}
		proposal, err := ctx.NodeClient.QueryProposalByID(dbProposal.ProposalID)
		if err != nil {
			log.Errorf("[CheckProposalPending] QueryProposalByID err: %s\n", errors.WithStack(err))
			continue
		}
		if proposal.Status == schema.SucceedStatus {
			dbProposal.Status = proposal.Status
			dbProposal.Winner = proposal.Winner
			dbProposal.BtcRevealTxHash = proposal.BitcoinTxHash
			ctx.DB.Save(dbProposal)
		}
		if proposal.Status == schema.PendingStatus &&
			proposal.Winner == ctx.B2NodeConfig.Address && proposal.BitcoinTxHash == "" {
			rs, err := inscribe.Inscribe(ctx.BTCConfig.PrivateKey, proposal.StateRootHash, proposal.ProofHash, ctx.BTCConfig.DestinationAddress, btcapi.ChainParams(ctx.BTCConfig.NetworkName))
			if err != nil {
				log.Errorf("[Handler.Inscribe] Inscribe err: %s\n", errors.WithStack(err))
				continue
			}
			str, err := json.Marshal(rs)
			if err != nil {
				log.Errorf("[Handler.Inscribe] Marshal result err: %s\n", errors.WithStack(err))
				continue
			}
			log.Infof("[Handler.Inscribe] inscribe result: %s", str)
			bitcoinTxHash := rs.RevealTxHashList[0].String()

			_, err = ctx.NodeClient.BitcoinTx(proposal.Id, proposal.Winner, bitcoinTxHash)
			if err != nil {
				log.Errorf("[Handler.Inscribe] BitcoinTx err: %s\n", errors.WithStack(err))
				continue
			}
			dbProposal.BtcRevealTxHash = bitcoinTxHash
			dbProposal.BtcCommitTxHash = rs.CommitTxHash.String()

			ctx.DB.Save(dbProposal)
		}
		if proposal.Status == schema.PendingStatus && proposal.BitcoinTxHash != "" && proposal.Winner != ctx.B2NodeConfig.Address {
			// 拿到bitcoin 去btc network query一下 确认状态 对比大于6个高度后 确认后就提交提案
			btcAPIClient := btcmempool.NewClient(btcapi.ChainParams(ctx.BTCConfig.NetworkName))
			transaction, err := btcAPIClient.GetTransactionByID(proposal.BitcoinTxHash)
			if err != nil {
				log.Errorf("[Handler.Inscribe] GetTransactionByID err: %s\n", errors.WithStack(err))
				continue
			}
			if transaction.Status.Confirmed && (ctx.LatestBTCBlockNumber-transaction.Status.BlockHeight) >= 6 {
				_, err = ctx.NodeClient.BitcoinTx(proposal.Id, proposal.Proposer, proposal.BitcoinTxHash)
				if err != nil {
					log.Errorf("[Handler.Inscribe] BitcoinTx err: %s\n", errors.WithStack(err))
					continue
				}
			}
		}
		time.Sleep(3 * time.Second)
	}
}
