package handler

import (
	"context"
	"time"

	"github.com/b2network/b2committer/pkg/btcapi"
	btcmempool "github.com/b2network/b2committer/pkg/btcapi/mempool"

	"github.com/b2network/b2committer/pkg/log"

	"github.com/b2network/b2committer/internal/svc"
)

func LatestBlackNumber(ctx *svc.ServiceContext) {
	for {
		latest, err := ctx.RPC.BlockNumber(context.Background())
		if err != nil {
			log.Errorf("[Handle.LatestBlackNumber]Syncing latest block number error: %s\n", err)
			time.Sleep(3 * time.Second)
			continue
		}
		ctx.LatestBlockNumber = int64(latest)
		log.Infof("[Handle.LatestBlackNumber] Syncing latest block number: %d \n", latest)

		btcAPIClient := btcmempool.NewClient(btcapi.ChainParams(ctx.BTCConfig.NetworkName))
		btcLatest, err := btcAPIClient.GetCurrentBlockHash()
		if err != nil {
			log.Errorf("[Handle.LatestBTCBlackNumber]Syncing btc network latest block number error: %s\n", err)
			time.Sleep(3 * time.Second)
			continue
		}
		ctx.LatestBTCBlockNumber = btcLatest
		log.Infof("[Handle.LatestBTCBlackNumber] Syncing btc network latest block number: %d \n", btcLatest)
		time.Sleep(3 * time.Second)
	}
}
