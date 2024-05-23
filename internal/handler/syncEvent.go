package handler

import (
	"sync"
	"time"

	"github.com/b2network/b2committer/pkg/log"

	"github.com/b2network/b2committer/internal/blockchain"
	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func SyncEvent(ctx *svc.ServiceContext) {
	for {
		var blocks []schema.SyncBlock
		err := ctx.DB.Where("status=? OR status=?", schema.BlockPending, schema.BlockRollback).Order("block_number").Limit(50).Find(&blocks).Error
		if err != nil {
			log.Errorf("[Handler.SyncEvent] Pending and rollback blocks err: %s\n", errors.WithStack(err))
			time.Sleep(5 * time.Second)
			continue
		}
		if len(blocks) == 0 {
			log.Infof("[Handler.SyncEvent] Pending blocks count is 0\n")
			time.Sleep(2 * time.Second)
			continue
		}

		var wg sync.WaitGroup
		for _, block := range blocks {
			wg.Add(1)
			go func(_wg *sync.WaitGroup, ctx *svc.ServiceContext, block schema.SyncBlock) {
				defer _wg.Done()
				if block.Status == schema.BlockPending {
					// add events & block.status= valid
					err = HandlePendingBlock(ctx, block)
					if err != nil {
						log.Errorf("[Handler.SyncEvent] HandlePendingBlock err: %s\n", errors.WithStack(err))
						time.Sleep(500 * time.Millisecond)
					}
				} else if block.Status == schema.BlockRollback {
					// event.status=rollback & block.status=invalid
					err = HandleRollbackBlock(ctx, block)
					if err != nil {
						log.Errorf("[Handler.SyncEvent] HandleRollbackBlock err: %s\n", errors.WithStack(err))
						time.Sleep(500 * time.Millisecond)
					}
				}
			}(&wg, ctx, block)
		}
		wg.Wait()
	}
}

func HandlePendingBlock(ctx *svc.ServiceContext, block schema.SyncBlock) error {
	log.Infof("[Handler.SyncEvent.PendingBlock]Start: %d, %s \n", block.BlockNumber, block.BlockHash)
	log.Infof("[Handler.SyncEvent.PendingBlock]GetContracts: %v\n", blockchain.GetContracts())
	log.Infof("[Handler.SyncEvent.PendingBlock]GetEvents: %v\n", blockchain.GetEvents())
	events, err := LogFilter(ctx, block, blockchain.GetContracts(), [][]common.Hash{blockchain.GetEvents()})
	log.Infof("[Handler.SyncEvent.PendingBlock] block %d, events number is %d:", block.BlockNumber, len(events))
	if err != nil {
		log.Errorf("[Handler.SyncEvent.PendingBlock] Log filter err: %s\n", err)
		return errors.WithStack(err)
	}
	eventCount := len(events)
	if eventCount > 0 && events[0].BlockHash != block.BlockHash {
		log.Infof("[Handler.SyncEvent.PendingBlock]Don't match block hash\n")
		return nil
	} else if eventCount > 0 && events[0].BlockHash == block.BlockHash {
		BatchEvents := make([]*schema.SyncEvent, 0)
		for _, event := range events {
			var one schema.SyncEvent
			log.Infof("[Handler.SyncEvent.PendingBlock]BlockLogIndexed %d ,TxHash %s,EventHash %s", event.BlockLogIndexed, event.TxHash, event.EventHash)
			err = ctx.DB.Select("id").Where("sync_block_id=? AND block_log_indexed=? AND tx_hash=? AND event_hash=? ",
				block.ID, event.BlockLogIndexed, event.TxHash, event.EventHash).First(&one).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				log.Errorf("[Handler.SyncEvent.PendingBlock]Query SyncEvent err: %s\n ", err)
				return errors.WithStack(err)
			} else if err == gorm.ErrRecordNotFound {
				BatchEvents = append(BatchEvents, event)
				log.Infof("[Handler.SyncEvent.PendingBlock]block %d, BatchEvents len is %d:", block.BlockNumber, len(BatchEvents))
			}
		}
		if len(BatchEvents) > 0 {
			err = ctx.DB.Transaction(func(tx *gorm.DB) error {
				err = tx.CreateInBatches(&BatchEvents, 200).Error
				if err != nil {
					log.Errorf("[Handler.SyncEvent.PendingBlock]CreateInBatches err: %s\n ", err)
					return errors.WithStack(err)
				}
				block.Status = schema.BlockValid
				block.EventCount = int64(eventCount)
				err = tx.Save(&block).Error
				if err != nil {
					log.Errorf("[Handler.SyncEvent.PendingBlock]Batch Events Update SyncBlock Status err: %s\n ", err)
					return errors.WithStack(err)
				}
				return nil
			})
			if err != nil {
				return err
			}
			return nil
		}
	}
	block.Status = schema.BlockValid
	block.EventCount = int64(eventCount)
	err = ctx.DB.Save(&block).Error
	if err != nil {
		log.Errorf("[Handler.PendingBlock]Update SyncBlock Status err: %s\n ", err)
		return errors.WithStack(err)
	}
	log.Infof("[Handler.SyncEvent.PendingBlock]End: %d, %s \n", block.BlockNumber, block.BlockHash)
	return nil
}

func HandleRollbackBlock(ctx *svc.ServiceContext, block schema.SyncBlock) error {
	log.Infof("[Handler.RollbackBlock] Start: %d, %s\n", block.BlockNumber, block.BlockHash)
	err := ctx.DB.Transaction(func(tx *gorm.DB) error {
		now := time.Now()
		// event.status=rollback by syncBlockId
		err := tx.Model(schema.SyncEvent{}).Where("sync_block_id=?", block.ID).
			Updates(map[string]interface{}{"status": schema.EventRollback, "updated_at": now}).Error
		if err != nil {
			log.Errorf("[Handler.RollbackBlock]Query SyncBlock Status err: %s ,id : %d \n", err, block.ID)
			return errors.WithStack(err)
		}
		block.Status = schema.BlockInvalid
		err = tx.Save(&block).Error
		if err != nil {
			log.Errorf("[Handler.RollbackBlock]Save SyncBlock Status err: %s\n ", err)
			return errors.WithStack(err)
		}
		return nil
	})
	if err != nil {
		log.Errorf("[Handler.RollbackBlock] err: %s\n ", err)
		return errors.WithStack(err)
	}
	return nil
}
