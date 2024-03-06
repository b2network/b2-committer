package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/pkg/contract"
	"github.com/b2network/b2committer/pkg/event/zkevm"
	"github.com/b2network/b2committer/pkg/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"os"
	"strings"
	"time"
)

type CollectionSequenceBatches struct {
	IsCompleted     bool
	SequenceBatches []*SequenceBatchesAndTxHash
	StartBatchNum   uint64
	EndBatchNum     uint64
}

type SequenceBatchesAndTxHash struct {
	NumBatch uint64
	TxHash   string
}

func SequenceBatches(ctx *svc.ServiceContext) {
	for {
		var dbProposal schema.Proposal
		err := ctx.DB.Where("status = ?", schema.ProposalSucceedStatus).Order("proposal_id asc").First(&dbProposal).Error
		//dbProposal.StartBatchNum
		//dbProposal.EndBatchNum
		//collectionSequenceBatches, err := GetSequenceBatchesFromStartBatchNum(ctx, 718506, 719063)
		collectionSequenceBatches, err := GetSequenceBatchesFromStartBatchNum(ctx, dbProposal.StartBatchNum, dbProposal.EndBatchNum)
		if err != nil {
			log.Errorf("[Handler.SequenceBatches][GetSequenceBatchesFromStartBatchNum] error info: %s", errors.WithStack(err).Error())
			time.Sleep(10 * time.Second)
			continue
		}
		if !collectionSequenceBatches.IsCompleted {
			log.Errorf("[Handler.SequenceBatches] sync batches not completed")
			time.Sleep(10 * time.Second)
			continue
		}
		sequenceBatchesMap, err := GetSequenceBatchesDetails(ctx, collectionSequenceBatches.SequenceBatches)
		if err != nil {
			log.Errorf("[Handler.SequenceBatches][GetSequenceBatchesDetails] error info: %s", errors.WithStack(err).Error())
			time.Sleep(3 * time.Second)
			continue
		}
		WriteFile(ctx, sequenceBatchesMap)
	}
}

func GetSequenceBatchesFromStartBatchNum(ctx *svc.ServiceContext, startBatchNum uint64, endBatchNum uint64) (*CollectionSequenceBatches, error) {
	isCompleted := false
	events := make([]schema.SyncEvent, 0)
	err := ctx.DB.Table("sync_events").Select("*, JSON_EXTRACT(data, '$.numBatch') as numBatch ").
		Where(" event_name = ? and JSON_EXTRACT(data, '$.numBatch') between ? and ?", "sequenceBatches",
			startBatchNum, endBatchNum).Order("numBatch").Find(&events).Error
	if err != nil {
		return nil, fmt.Errorf("[GetSequenceBatchesFromStartBatchNum] dbbase err: %s", err)
	}

	sequenceBatchesAndTxHashes := make([]*SequenceBatchesAndTxHash, 0)
	for _, event := range events {
		sequenceBatch := &zkevm.SequenceBatches{}
		err = sequenceBatch.ToObj(event.Data)
		if err != nil {
			return nil, fmt.Errorf("[GetSequenceBatchesFromStartBatchNum] parse event data error: %s", errors.WithStack(err))
		}
		SequenceBatchesAndTxHash := &SequenceBatchesAndTxHash{
			NumBatch: sequenceBatch.BatchNum,
			TxHash:   event.TxHash,
		}
		sequenceBatchesAndTxHashes = append(sequenceBatchesAndTxHashes, SequenceBatchesAndTxHash)
		if sequenceBatch.BatchNum == endBatchNum {
			isCompleted = true
		}
	}

	return &CollectionSequenceBatches{
		StartBatchNum:   startBatchNum,
		EndBatchNum:     endBatchNum,
		SequenceBatches: sequenceBatchesAndTxHashes,
		IsCompleted:     isCompleted,
	}, nil
}

func GetSequenceBatchesDetails(ctx *svc.ServiceContext, sequenceBatches []*SequenceBatchesAndTxHash) (map[uint64]map[string]interface{}, error) {
	abiObject, err := abi.JSON(strings.NewReader(contract.ZkEVMMetaData.ABI))

	if err != nil {
		return nil, fmt.Errorf("[GetSequenceBatchesDetails] parse abi error: %s", errors.WithStack(err))
	}
	sequenceBatchesMap := make(map[uint64]map[string]interface{})
	for _, sequenceBatch := range sequenceBatches {
		txHash := sequenceBatch.TxHash
		tx, _, err := ctx.RPC.TransactionByHash(context.Background(), common.HexToHash(txHash))
		if err != nil {
			return nil, fmt.Errorf("[GetSequenceBatchesDetails] get tx error: %s", errors.WithStack(err))
		}
		inputsMap, methodName := DecodeTransactionInputData(abiObject, tx.Data())

		if methodName != "sequenceBatches" {
			return nil, fmt.Errorf("[GetSequenceBatchesDetails] methodName is :  %s parse method error: %s", methodName, errors.WithStack(err))
		}
		sequenceBatchesMap[sequenceBatch.NumBatch] = inputsMap
	}
	return sequenceBatchesMap, nil
}

func WriteFile(ctx *svc.ServiceContext, sequenceBatchesMap map[uint64]map[string]interface{}) {
	jsonData, err := json.Marshal(sequenceBatchesMap)
	if err != nil {
		log.Errorf("[WriteFile] json marshal error: %s", errors.WithStack(err))
		return
	}

	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("create file error:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("write file error:", err)
		return
	}

	fmt.Println("map write file output.json success")
}
