package schema

import (
	"github.com/ethereum/go-ethereum/common"
)

type TxsProposal struct {
	Base
	ProposalID       uint64
	StartTimestamp   uint64
	EndTimestamp     uint64
	StartBlockNumber uint64
	EndBlockNumber   uint64
	TxsRoot          string
	Status           uint8
	DsType           uint8
	DsTxHash         string
	Winner           common.Address
	BlockList        []uint64
}

func (TxsProposal) TableName() string {
	return "txs_proposal"
}
