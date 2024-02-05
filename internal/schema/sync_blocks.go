package schema

const (
	BlockPending  = "pending"
	BlockValid    = "valid"
	BlockRollback = "rollback"
	BlockInvalid  = "invalid"
)

type SyncBlock struct {
	Base
	Blockchain  string `json:"blockchain"`
	Miner       string `json:"miner"`
	BlockTime   int64  `json:"block_time"`
	BlockNumber int64  `json:"block_number"`
	BlockHash   string `json:"block_hash"`
	TxCount     int64  `json:"tx_count"`
	EventCount  int64  `json:"event_count"`
	ParentHash  string `json:"parent_hash"`
	Status      string `json:"status"`
	CheckCount  int64  `json:"check_count"`
}

func (SyncBlock) TableName() string {
	return "`sync_blocks`"
}

type SyncBlockHistory SyncBlock

func (SyncBlockHistory) TableName() string {
	return "`sync_blocks_history`"
}
