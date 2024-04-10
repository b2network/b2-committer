package zkevm

import (
	"encoding/json"

	"github.com/b2network/b2committer/pkg/event"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	VerifyBatchesName = "verifyBatchesTrustedAggregator"

	VerifyBatchesHash = crypto.Keccak256([]byte("VerifyBatchesTrustedAggregator(uint32,uint64,bytes32,bytes32,address)"))
)

type VerifyBatches struct {
	RollupID   uint32 `json:"rollupID"`
	BatchNum   uint64 `json:"numBatch"`
	StateRoot  string `json:"stateRoot"`
	ExitRoot   string `json:"exitRoot"`
	Aggregator string `json:"aggregator"`
}

func (*VerifyBatches) Name() string {
	return VerifyBatchesName
}

func (*VerifyBatches) EventHash() common.Hash {
	return common.BytesToHash(VerifyBatchesHash)
}

func (t *VerifyBatches) ToObj(data string) error {
	err := json.Unmarshal([]byte(data), &t)
	if err != nil {
		return err
	}
	return nil
}

func (*VerifyBatches) Data(log types.Log) (string, error) {

	transfer := &VerifyBatches{
		RollupID:   uint32(event.TopicToInt64(log, 1)),
		Aggregator: event.TopicToAddress(log, 2).Hex(),
		BatchNum:   uint64(event.DataToInt64(log, 0)),
		StateRoot:  event.DataToHash(log, 1).Hex(),
		ExitRoot:   event.DataToHash(log, 2).Hex(),
	}
	data, err := event.ToJSON(transfer)
	if err != nil {
		return "", err
	}
	return data, nil
}
