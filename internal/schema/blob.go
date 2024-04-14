package schema

type BlobInfo struct {
	Base
	BlockNumber           int64  `json:"block_number"`
	BlockHashHex          string `json:"block_hash_hex"`
	BlockTime             uint64 `json:"block_time"`
	BlobVersionedHash     string `json:"blob_versioned_hash"`
	BlobHashesIndex       uint64 `json:"blob_hashes_index"`
	BlobSideCarIndex      uint64 `json:"blob_side_car_index"`
	BlobSideCarCommitment string `json:"blob_side_car_commitment"`
	Blob                  string `json:"blob"`
}

func (BlobInfo) TableName() string {
	return "blob_info"
}
