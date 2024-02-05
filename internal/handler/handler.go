package handler

import (
	"github.com/b2network/b2committer/internal/svc"
)

func Run(ctx *svc.ServiceContext) {
	//// 最新高度
	//go LatestBlackNumber(ctx)
	//// 同步区块
	//go SyncBlock(ctx)
	//// 同步事件
	//go SyncEvent(ctx)
	//// 执行committer
	//go Committer(ctx)
	//// 检查vote状态
	//go CheckStatus(ctx)
	//// 检查并铭刻
	//go Inscribe(ctx)
	//// check time out
	//go CheckStatusTimeOut(ctx)

	go SyncProposal(ctx)

	//// 检查
	// go CheckBlock(ctx)
	//// 迁移Block
	// go MigrateBlock(ctx)
	//// 迁移Event
	// go MigrateEvent(ctx)
	//// 处理syncTask
	// go SyncTask(ctx)
}
