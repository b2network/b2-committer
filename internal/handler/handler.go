package handler

import (
	"github.com/b2network/b2committer/internal/svc"
)

func Run(ctx *svc.ServiceContext) {

	// query last block number
	go LatestBlackNumber(ctx)
	// sync blocks
	go SyncBlock(ctx)
	// sync events
	go SyncEvent(ctx)
	// query blob and store in local
	go QueryBlobOnChainAndStoreInLocal(ctx)
	// execute committer
	//go Committer(ctx)
	//// check and inscribe
	//go Inscribe(ctx)
	//// check status
	//go CheckStatusByContract(ctx)
	//// sync proposal
	//go SyncProposal(ctx)
	//// sequence batches
	//go SequenceBatches(ctx)
}
