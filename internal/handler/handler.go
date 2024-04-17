package handler

import (
	"github.com/b2network/b2committer/internal/svc"
)

func Run(ctx *svc.ServiceContext) {

	// query last block number
	//go LatestBlackNumber(ctx)
	//// sync blocks
	//go SyncBlock(ctx)
	//// sync events
	//go SyncEvent(ctx)
	// query blob and store in local
	go QueryBlobOnChainAndStoreInLocal(ctx)
	// commit and vote txs proposal
	//go GetBlobsAndCommitProposal(ctx)

}
