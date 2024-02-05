package handler

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/b2network/b2committer/pkg/log"

	"github.com/b2network/b2committer/internal/schema"
	"github.com/b2network/b2committer/internal/svc"
	"github.com/b2network/b2committer/pkg/event/zkevm"
	"github.com/b2network/b2committer/pkg/merkle"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	mt "github.com/txaty/go-merkletree"
)

type VerifyBatchesTrustedAggregatorParams struct {
	pendingStateNum  uint64
	initNumBatch     uint64
	finalNewBatch    uint64
	newLocalExitRoot string
	newStateRoot     string
	proof            string
}

type VerifyBatchesAndTxHash struct {
	verifyBatches *zkevm.VerifyBatches
	txHash        string
}

type VerifyRangBatchInfo struct {
	startBatchNum uint64
	endBatchNum   uint64
	proofRootHash string
	stateRootHash string
}

// Committer find verifyBatchesTrustedAggregator event and commit stateRoot proof to b2node
func Committer(ctx *svc.ServiceContext) {
	for {
		var proposals []schema.Proposal
		lastProposalID, lastFinalBatchNum, err := ctx.NodeClient.QueryLastProposalID()
		if err != nil {
			log.Errorf("[Handler.Committer][QueryLastProposalID] error info:", errors.WithStack(err))
			time.Sleep(10 * time.Second)
			continue
		}
		err = ctx.DB.Where("end_batch_num > ?", lastFinalBatchNum).Find(&proposals).Error
		if err != nil {
			log.Errorf("[Handler.Committer][DB] error info:", errors.WithStack(err))
			time.Sleep(10 * time.Second)
			continue
		}
		if len(proposals) > 0 {
			log.Errorf("[Handler.Committer] proposal already is existed, lastFinalBatchNum: %s", lastFinalBatchNum)
			continue
		}
		verifyBatchInfo, err := GetVerifyBatchInfoByLastBatchNum(ctx, lastFinalBatchNum)
		if err != nil {
			log.Errorf("[Handler.Committer] error info:", errors.WithStack(err))
			time.Sleep(10 * time.Second)
			continue
		}
		err = committerProposal(ctx, verifyBatchInfo, lastProposalID)
		if err != nil {
			log.Errorf("[Handler.Committer] error info:", errors.WithStack(err))
			time.Sleep(10 * time.Second)
			continue
		}
		time.Sleep(10 * time.Second)
	}
}

func GetVerifyBatchInfoByLastBatchNum(ctx *svc.ServiceContext, lastFinalBatchNum uint64) (*VerifyRangBatchInfo, error) {
	verifyBatchesAndTxHashs, err := GetVerifyBatchesFromStartBatchNum(ctx, lastFinalBatchNum, ctx.Config.LimitNum)
	if err != nil || len(verifyBatchesAndTxHashs) != ctx.Config.LimitNum {
		return nil, fmt.Errorf("[GetVerifyBatchInfo] error info: %s", errors.WithStack(err))
	}
	verifyBatchesParams := make([]*VerifyBatchesTrustedAggregatorParams, 0, ctx.Config.LimitNum)
	for _, verifyBatch := range verifyBatchesAndTxHashs {
		verifyBatchParam, err := GetVerifyBatchesParamsByTxHash(ctx, common.HexToHash(verifyBatch.txHash))
		if err != nil {
			return nil, fmt.Errorf("[GetVerifyBatchInfo] error info: %s", errors.WithStack(err))
		}
		verifyBatchesParams = append(verifyBatchesParams, verifyBatchParam)
	}
	verifyBatchInfo, err := GetMerkleStateRootsAndProofs(verifyBatchesParams)
	if err != nil {
		return nil, fmt.Errorf("[GetVerifyBatchInfo] error info: %s", errors.WithStack(err))
	}
	return verifyBatchInfo, nil
}

// CommitterProposal committer transaction to b2-node
func committerProposal(ctx *svc.ServiceContext, verifyBatchInfo *VerifyRangBatchInfo, lastProposalID uint64) error {
	proposalID, err := ctx.NodeClient.SubmitProof(lastProposalID+1, ctx.B2NodeConfig.Address, verifyBatchInfo.proofRootHash, verifyBatchInfo.stateRootHash,
		verifyBatchInfo.startBatchNum, verifyBatchInfo.endBatchNum)
	if err != nil {
		return fmt.Errorf("[committerProposal] submit proof error info: %s", errors.WithStack(err))
	}

	dbProposal := &schema.Proposal{
		Base: schema.Base{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		EndBatchNum:   verifyBatchInfo.endBatchNum,
		ProposalID:    proposalID,
		Proposer:      ctx.B2NodeConfig.Address,
		Status:        schema.VotingStatus,
		StateRootHash: verifyBatchInfo.stateRootHash,
		ProofRootHash: verifyBatchInfo.proofRootHash,
		StartBatchNum: verifyBatchInfo.startBatchNum,
	}

	// store db
	err = ctx.DB.Save(dbProposal).Error
	if err != nil {
		return fmt.Errorf("[committerProposal] save proposal error info: %s", errors.WithStack(err))
	}
	return nil
}

func GetMerkleStateRootsAndProofs(params []*VerifyBatchesTrustedAggregatorParams) (*VerifyRangBatchInfo, error) {
	stateRoots := make([]string, 0, 10)
	proofs := make([]string, 0, 10)
	var startBatchNum uint64
	var endBatchNum uint64
	for index, param := range params {
		if index == 0 {
			startBatchNum = param.initNumBatch
			endBatchNum = param.finalNewBatch
		}
		if startBatchNum > param.initNumBatch {
			startBatchNum = param.initNumBatch
		}
		if endBatchNum <= param.finalNewBatch {
			endBatchNum = param.finalNewBatch
		}
		stateRoots = append(stateRoots, param.newStateRoot)
		proofs = append(proofs, param.proof)
	}
	if startBatchNum == 0 {
		startBatchNum = 1
	}
	stateBlocks := merkle.GenerateBlocks(stateRoots)
	proofBlocks := merkle.GenerateBlocks(proofs)

	stateTree, err := mt.New(nil, stateBlocks)
	if err != nil {
		return nil, fmt.Errorf("[GetMerkleStateRootsAndProofs] generate state tree err: %s", err)
	}
	proofTree, err := mt.New(nil, proofBlocks)
	if err != nil {
		return nil, fmt.Errorf("[GetMerkleStateRootsAndProofs] generate proof tree err: %s", err)
	}
	return &VerifyRangBatchInfo{
		startBatchNum: startBatchNum,
		endBatchNum:   endBatchNum,
		proofRootHash: hex.EncodeToString(proofTree.Root),
		stateRootHash: hex.EncodeToString(stateTree.Root),
	}, nil
}

func GetVerifyBatchesFromStartBatchNum(ctx *svc.ServiceContext, startBatchNum uint64, limit int) ([]*VerifyBatchesAndTxHash, error) {
	events := make([]schema.SyncEvent, 0, limit)
	err := ctx.DB.Table("sync_events").Select("*, JSON_EXTRACT(data, '$.numBatch') as numBatch").
		Where("JSON_EXTRACT(data, '$.numBatch') >= ?", startBatchNum).Order("numBatch").Limit(limit).Find(&events).Error
	if err != nil {
		return nil, fmt.Errorf("[GetVerifyBatchesFromStartBatchNum] dbbase err: %s", err)
	}
	if len(events) != 10 {
		return nil, fmt.Errorf("[GetVerifyBatchesFromStartBatchNum] sync_events find event is not enough %s", err)
	}
	verifyBatchesAndTxHashs := make([]*VerifyBatchesAndTxHash, 0, limit)
	for _, event := range events {
		verifyBatch := &zkevm.VerifyBatches{}
		err = verifyBatch.ToObj(event.Data)
		if err != nil {
			return nil, fmt.Errorf("[GetVerifyBatchesFromStartBatchNum] parse event data error: %s", errors.WithStack(err))
		}
		verifyBatchesAndTxHash := &VerifyBatchesAndTxHash{
			verifyBatches: verifyBatch,
			txHash:        event.TxHash,
		}
		verifyBatchesAndTxHashs = append(verifyBatchesAndTxHashs, verifyBatchesAndTxHash)
	}
	return verifyBatchesAndTxHashs, nil
}

func DecodeTransactionInputData(contractABI abi.ABI, data []byte) (map[string]interface{}, string) {
	methodSigData := data[:4]
	inputsSigData := data[4:]

	method, err := contractABI.MethodById(methodSigData)
	if err != nil {
		log.Errorf("[DecodeTransactionInputData] parse abi error: %s\n", errors.WithStack(err))
	}

	inputsMap := make(map[string]interface{})

	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		log.Errorf("[DecodeTransactionInputData] parse abi error: %s\n", errors.WithStack(err))
	}
	return inputsMap, method.Name
}

func GetVerifyBatchesParamsByTxHash(ctx *svc.ServiceContext, txHash common.Hash) (*VerifyBatchesTrustedAggregatorParams, error) {
	abiObject, err := abi.JSON(strings.NewReader(ZkEVMMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("[GetVerifyBatchesParamsByTxHash] parse abi error: %s", errors.WithStack(err))
	}
	// method := "VerifyBatchesTrustedAggregator"
	tx, _, err := ctx.RPC.TransactionByHash(context.Background(), txHash)

	inputsMap, methodName := DecodeTransactionInputData(abiObject, tx.Data())

	if methodName != "verifyBatchesTrustedAggregator" {
		return nil, fmt.Errorf("[GetVerifyBatchesParamsByTxHash] methodName is :  %s parse method error: %s", methodName, errors.WithStack(err))
	}

	e := inputsMap["newLocalExitRoot"].([32]byte)
	f := inputsMap["newStateRoot"].([32]byte)
	g := inputsMap["proof"].([24][32]byte)

	var result string
	for _, arr := range g {
		result += hex.EncodeToString(arr[:])
	}

	log.Debugf("newStateRoot outputs: %v\n", hex.EncodeToString(f[:]))
	log.Debugf("proof outputs: %v\n", common.HexToHash(result))

	return &VerifyBatchesTrustedAggregatorParams{
		pendingStateNum:  inputsMap["pendingStateNum"].(uint64),
		initNumBatch:     inputsMap["initNumBatch"].(uint64),
		finalNewBatch:    inputsMap["finalNewBatch"].(uint64),
		newLocalExitRoot: hex.EncodeToString(e[:]),
		newStateRoot:     hex.EncodeToString(f[:]),
		proof:            common.HexToHash(result).String(),
	}, err
}
