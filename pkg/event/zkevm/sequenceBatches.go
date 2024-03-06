package zkevm

import (
	"encoding/json"
	"github.com/b2network/b2committer/pkg/event"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	SequenceBatchesName = "sequenceBatches"

	SequenceBatchesNameHash = crypto.Keccak256([]byte("SequenceBatches(uint64)"))
)

type SequenceBatches struct {
	BatchNum uint64 `json:"numBatch"`
}

func (*SequenceBatches) Name() string {
	return SequenceBatchesName
}

func (*SequenceBatches) EventHash() common.Hash {
	return common.BytesToHash(SequenceBatchesNameHash)
}

func (t *SequenceBatches) ToObj(data string) error {
	err := json.Unmarshal([]byte(data), &t)
	if err != nil {
		return err
	}
	return nil
}

func (*SequenceBatches) Data(log types.Log) (string, error) {
	transfer := &SequenceBatches{
		BatchNum: uint64(event.TopicToInt64(log, 1)),
	}
	data, err := event.ToJSON(transfer)
	if err != nil {
		return "", err
	}
	return data, nil
}
