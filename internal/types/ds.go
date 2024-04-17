package types

import (
	"encoding/json"
	"github.com/b2network/b2committer/internal/schema"
)

type DsProposal struct {
	ChainID    int64
	ProposalID uint64
	TxsRoot    string
	Blobs      *[]DsBlob
}

type DsBlob struct {
	BlockID int64
	Blob    string
}

func NewDsProposal(chainID int64, proposalID uint64, txsRoot string, blobs []schema.BlobInfo) *DsProposal {
	return &DsProposal{
		ChainID:    chainID,
		ProposalID: proposalID,
		TxsRoot:    txsRoot,
		Blobs:      convertBlobToDsBlob(blobs),
	}
}

func convertBlobToDsBlob(blobs []schema.BlobInfo) *[]DsBlob {
	var dsBlobs []DsBlob
	for _, blob := range blobs {
		dsBlobs = append(dsBlobs, DsBlob{
			BlockID: blob.BlockNumber,
			Blob:    blob.Blob,
		})
	}
	return &dsBlobs
}

func (b *DsProposal) GetDBBlobInfos() ([]schema.BlobInfo, error) {
	var dbBlobs []schema.BlobInfo
	for _, blob := range *b.Blobs {
		dbBlobs = append(dbBlobs, schema.BlobInfo{
			BlockNumber: blob.BlockID,
			Blob:        blob.Blob,
		})
	}
	return dbBlobs, nil
}

func (b *DsProposal) MarshalJson() ([]byte, error) {
	marshal, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
