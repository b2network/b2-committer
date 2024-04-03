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
	// execute committer
	go Committer(ctx)
	// check and inscribe
	go Inscribe(ctx)
	// check status
	go CheckStatusVoting(ctx)
	// check pending
	go CheckStatusPending(ctx)
	// check time out
	go CheckStatusPendingTimeOut(ctx)
	// sync proposal
	go SyncProposal(ctx)
	// sequence batches
	go SequenceBatches(ctx)
	// upload batch detail to ar
	go BatchDetailsToAr(ctx)
}
