package op

import (
	"encoding/json"
	"math/big"

	"github.com/b2network/b2committer/pkg/event"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	OutputProposedName = "OutputProposed"

	OutputProposedHash = crypto.Keccak256([]byte("OutputProposed(bytes32,uint256,uint256,uint256)"))
)

type OutputProposed struct {
	OutputRoot    string   `json:"outputRoot"`
	L2OutputIndex *big.Int `json:"l2OutputIndex"`
	L2BlockNumber *big.Int `json:"l2BlockNumber"`
	L1Timestamp   *big.Int `json:"l1Timestamp"`
}

func (*OutputProposed) Name() string {
	return OutputProposedName
}

func (*OutputProposed) EventHash() common.Hash {
	return common.BytesToHash(OutputProposedHash)
}

func (t *OutputProposed) ToObj(data string) error {
	err := json.Unmarshal([]byte(data), &t)
	if err != nil {
		return err
	}
	return nil
}

func (*OutputProposed) Data(log types.Log) (string, error) {
	transfer := &OutputProposed{
		OutputRoot:    event.TopicToHash(log, 1).Hex(),
		L2OutputIndex: big.NewInt(event.TopicToInt64(log, 2)),
		L2BlockNumber: big.NewInt(event.TopicToInt64(log, 3)),
		L1Timestamp:   big.NewInt(event.DataToInt64(log, 0)),
	}
	data, err := event.ToJSON(transfer)
	if err != nil {
		return "", err
	}
	return data, nil
}
