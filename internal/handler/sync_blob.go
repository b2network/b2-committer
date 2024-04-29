package handler

import (
	"context"
	"math/big"
	"time"

	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/pkg/errcode"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func QueryBlobOnChainAndStoreInLocal(ctx *svc.ServiceContext) {
	time.Sleep(10 * time.Second)
	var dbBlob schema.BlobInfo
	var count int64
	err := ctx.DB.Order("block_number desc").First(&dbBlob).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		panic(err)
	}

	switch err {
	case gorm.ErrRecordNotFound:
		ctx.SyncedBlobBlockNumber = ctx.Config.InitBlobBlockNumber
		ctx.SyncedBlobBlockHash = common.HexToHash(ctx.Config.InitBlobBlockHash)
	default:
		ctx.SyncedBlobBlockNumber = dbBlob.BlockNumber
		ctx.SyncedBlobBlockHash = common.HexToHash(dbBlob.BlockHashHex)
	}

	log.Infof("[Handler.QueryBlobOnChainAndStoreInLocal]SyncedBlobBlockNumber: %d", ctx.SyncedBlobBlockNumber)
	log.Infof("[Handler.QueryBlobOnChainAndStoreInLocal]SyncedBlobBlockHash:%s", ctx.SyncedBlobBlockHash.String())

	for {
		syncingBlobBlockNumber := ctx.SyncedBlobBlockNumber + 1
		log.Infof("[Handler.QueryBlobOnChainAndStoreInLocal] Try to sync block number: %d\n", syncingBlobBlockNumber)
		if syncingBlobBlockNumber == 23373 {
			log.Infof("debug")
		}

		// if syncingBlobBlockNumber > ctx.LatestBlockNumber {
		//	time.Sleep(3 * time.Second)
		//	continue
		//}

		blockOnChain, err := ctx.RPC.BlockByNumber(context.Background(), big.NewInt(syncingBlobBlockNumber))
		if err != nil {
			log.Errorf("[Handler.QueryBlobOnChainAndStoreInLocal] Get block by number error: %s", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		if blockOnChain.ParentHash() != ctx.SyncedBlobBlockHash {
			log.Errorf("[Handler.QueryBlobOnChainAndStoreInLocal] ParentHash of the block being synchronized is inconsistent: %s \n", ctx.SyncedBlobBlockHash)
			continue
		}

		blobInfos, err1 := ctx.BlobDataSource.GetBlobByBlockNum(context.Background(), blockOnChain.Number())
		if errors.Is(err1, errcode.ErrNoBlobFoundInBlock) {
			log.Errorf("[Handler.QueryBlobOnChainAndStoreInLocal] %s", blockOnChain.Number(), err1.Error())
			ctx.SyncedBlobBlockNumber = blockOnChain.Number().Int64()
			ctx.SyncedBlobBlockHash = blockOnChain.Hash()
			continue
		}
		if err1 != nil {
			log.Errorf("[Handler.QueryBlobOnChainAndStoreInLocal] Get blob by block number error: %s", err1.Error())
			continue
		}

		err1 = ctx.DB.Where("block_number=? ", syncingBlobBlockNumber).Order("block_number desc").Find(&dbBlob).Count(&count).Error
		if err1 != nil && err != gorm.ErrRecordNotFound {
			log.Errorf("[Handler.QueryBlobOnChainAndStoreInLocal] DB Count error: %s", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}

		if count != 0 {
			if int(count) == len(blobInfos) {
				err := ctx.DB.Where("block_number=? ", syncingBlobBlockNumber).Order("block_number desc").First(&dbBlob).Error
				if err != nil {
					log.Errorf("[Handler.QueryBlobOnChainAndStoreInLocal] DB First error: %s", err.Error())
					time.Sleep(3 * time.Second)
					continue
				}
				ctx.SyncedBlobBlockNumber = blockOnChain.Number().Int64()
				ctx.SyncedBlobBlockHash = blockOnChain.Hash()
			} else {
				err := ctx.DB.Delete(&dbBlob, "block_number=?", syncingBlobBlockNumber).Error
				if err != nil {
					log.Errorf("[Handler.QueryBlobOnChainAndStoreInLocal] DB Delete error: %s", err.Error())
					time.Sleep(3 * time.Second)
					continue
				}
			}
			continue
		}

		for _, bif := range blobInfos {
			blob := &schema.BlobInfo{
				BlockNumber:           blockOnChain.Number().Int64(),
				BlockHashHex:          blockOnChain.Hash().String(),
				BlockTime:             blockOnChain.Time(),
				BlobVersionedHash:     bif.Hash.Hash.String(),
				BlobHashesIndex:       bif.Hash.Index,
				BlobSideCarIndex:      uint64(bif.BlobSidecar.Index),
				BlobSideCarCommitment: bif.BlobSidecar.KZGCommitment.String(),
				Blob:                  bif.BlobSidecar.Blob.String(),
			}

			err := ctx.DB.Save(blob).Error
			if err != nil {
				log.Errorf("[Handler.SyncBlock] DB Create SyncBlock error: %s\n", errors.WithStack(err))
				time.Sleep(1 * time.Second)
				continue
			}
		}
		ctx.SyncedBlobBlockNumber = blockOnChain.Number().Int64()
		ctx.SyncedBlobBlockHash = blockOnChain.Hash()
	}
}
