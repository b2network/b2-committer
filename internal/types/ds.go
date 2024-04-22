package types

import (
	"encoding/json"
	"github.com/b2network/b2committer/internal/schema"
)

type DsTxsProposal struct {
	ChainID    int64
	ProposalID uint64
	TxsRoot    string
	Blobs      *[]DsBlob
}

type DsStateRootProposal struct {
	ChainID      int64
	ProposalID   uint64
	OutputRoot   string
	OutputEvents []OutputEvent
}

type DsBlob struct {
	BlockID int64
	Blob    string
}

func NewDsTxsProposal(chainID int64, proposalID uint64, txsRoot string, blobs []schema.BlobInfo) *DsTxsProposal {
	return &DsTxsProposal{
		ChainID:    chainID,
		ProposalID: proposalID,
		TxsRoot:    txsRoot,
		Blobs:      convertBlobToDsBlob(blobs),
	}
}

func NewDsStateRootProposal(chainID int64, proposalID uint64, outputRoot string, events []schema.SyncEvent) (*DsStateRootProposal, error) {
	outputEvents, err := ConvertEventDataToOutputEvent(events)
	if err != nil {
		return nil, err
	}
	return &DsStateRootProposal{
		ChainID:      chainID,
		ProposalID:   proposalID,
		OutputRoot:   outputRoot,
		OutputEvents: outputEvents,
	}, nil
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

func (b *DsTxsProposal) GetDBBlobInfos() ([]schema.BlobInfo, error) {
	var dbBlobs []schema.BlobInfo
	for _, blob := range *b.Blobs {
		dbBlobs = append(dbBlobs, schema.BlobInfo{
			BlockNumber: blob.BlockID,
			Blob:        blob.Blob,
		})
	}
	return dbBlobs, nil
}

func (b *DsTxsProposal) MarshalJson() ([]byte, error) {
	marshal, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (s *DsStateRootProposal) MarshalJson() ([]byte, error) {
	marshal, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

type BtcStateRootProposal struct {
	Proposal *StateRootProposal
	ChainID  int64
}
