package store

type DecentralizedStore interface {
	StoreDetailsOnChain(txs []byte, chainID int64, proposalID uint64) (string, error)
	QueryDetailsByTxID(txID string) ([]byte, error)
}
