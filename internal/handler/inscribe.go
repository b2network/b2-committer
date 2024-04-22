package handler

import (
	"encoding/json"
	"github.com/b2network/b2committer/pkg/btcapi"
	btcmempool "github.com/b2network/b2committer/pkg/btcapi/mempool"
	"github.com/b2network/b2committer/pkg/inscribe"
	"time"

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
		err := ctx.DB.Where("status=? AND btc_tx_hash=''", schema.ProposalPendingStatus).Order("end_batch_num asc").First(&dbProposal).Error
		if err != nil {
			log.Errorf("[Handler.Inscribe] no proposal wait for inscribe. err: %s\n", errors.WithStack(err).Error())
			time.Sleep(5 * time.Second)
			continue
		}
		proposal, err := ctx.NodeClient.QueryProposalByID(dbProposal.ProposalID)
		if err != nil {
			log.Errorf("[CheckProposalPending] QueryProposalByID err: %s\n", errors.WithStack(err).Error())
			continue
		}

		if proposal.Status == schema.ProposalPendingStatus &&
			proposal.Winner.String() == ctx.B2NodeConfig.Address && proposal.BtcTxHash == "" {
			rs, err := inscribe.Inscribe(ctx.BTCConfig.PrivateKey, []byte(proposal.StateRootHash+proposal.ProofHash),
				ctx.BTCConfig.DestinationAddress, btcapi.ChainParams(ctx.BTCConfig.NetworkName))
			if err != nil {
				log.Errorf("[Handler.Inscribe] Inscribe err: %s\n", errors.WithStack(err).Error())
				continue
			}
			str, err := json.Marshal(rs)
			if err != nil {
				log.Errorf("[Handler.Inscribe] Marshal result err: %s\n", errors.WithStack(err).Error())
				continue
			}
			log.Infof("[Handler.Inscribe] inscribe result: %s", str)
			bitcoinTxHash := rs.RevealTxHashList[0].String()

			_, err = ctx.NodeClient.BitcoinTxHash(proposal.Id, bitcoinTxHash)
			if err != nil {
				log.Errorf("[Handler.Inscribe] BitcoinTx err: %s\n", errors.WithStack(err).Error())
				continue
			}
			dbProposal.BtcTxHash = bitcoinTxHash
			ctx.DB.Save(dbProposal)
		}
		if proposal.Status == schema.ProposalPendingStatus && proposal.BtcTxHash != "" && proposal.Winner.String() != ctx.B2NodeConfig.Address {
			// Get bitcoin txHash and query on btc network  confirm status If the comparison is greater than 6 heights, submit the proposal after confirmation
			btcAPIClient := btcmempool.NewClient(btcapi.ChainParams(ctx.BTCConfig.NetworkName))
			transaction, err := btcAPIClient.GetTransactionByID(proposal.BtcTxHash)
			if err != nil {
				log.Errorf("[Handler.Inscribe] GetTransactionByID err: %s\n", errors.WithStack(err).Error())
				continue
			}
			if transaction.Status.Confirmed && (ctx.LatestBTCBlockNumber-transaction.Status.BlockHeight) >= 6 {
				_, err = ctx.NodeClient.BitcoinTxHash(proposal.Id, proposal.BtcTxHash)
				if err != nil {
					log.Errorf("[Handler.Inscribe] BitcoinTx err: %s\n", errors.WithStack(err).Error())
					continue
				}
			}
		}
		time.Sleep(3 * time.Second)
	}
}
