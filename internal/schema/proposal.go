package schema

type Proposal struct {
	Base
	ProposalID          uint64 `json:"proposal_id"`
	StateRootHash       string `json:"state_root_hash"`
	ProofRootHash       string `json:"proof_root_hash"`
	StartBatchNum       uint64 `json:"start_batch_num"`
	EndBatchNum         uint64 `json:"end_batch_num"`
	BtcTxHash           string `json:"btc_tx_hash"`
	Winner              string `json:"winner"`
	Status              uint64 `json:"status"`
	GenerateDetailsFile bool   `json:"generate_details_file"`
	ArTxHash            string `json:"ar_tx_hash"`
}

func (Proposal) TableName() string {
	return "proposal"
}
