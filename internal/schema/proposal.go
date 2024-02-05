package schema

const (
	VotingStatus  = 0
	PendingStatus = 1
	SucceedStatus = 2
	TimeoutStatus = 3
)

type Proposal struct {
	Base
	ProposalID      uint64 `json:"proposal_id"`
	Proposer        string `json:"proposer"`
	StateRootHash   string `json:"state_root_hash"`
	ProofRootHash   string `json:"proof_root_hash"`
	StartBatchNum   uint64 `json:"start_batch_num"`
	EndBatchNum     uint64 `json:"end_batch_num"`
	BtcCommitTxHash string `json:"btc_commit_tx_hash"`
	BtcRevealTxHash string `json:"btc_reveal_tx_hash"`
	BlockHeight     uint64 `json:"block_height"`
	Winner          string `json:"winner"`
	Status          uint64 `json:"status"`
}

func (Proposal) TableName() string {
	return "`proposal`"
}
