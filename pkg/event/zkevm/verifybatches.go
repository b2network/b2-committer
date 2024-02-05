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

	VerifyBatchesHash = crypto.Keccak256([]byte("VerifyBatchesTrustedAggregator(uint64,bytes32,address)"))
)

type VerifyBatches struct {
	BatchNum   uint64 `json:"numBatch"`
	StateRoot  string `json:"stateRoot"`
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
		BatchNum:   uint64(event.TopicToInt64(log, 1)),
		StateRoot:  event.DataToHash(log, 0).Hex(),
		Aggregator: event.TopicToHash(log, 2).Hex(),
	}
	data, err := event.ToJSON(transfer)
	if err != nil {
		return "", err
	}
	return data, nil
}
