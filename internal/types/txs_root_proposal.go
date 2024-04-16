package types

import "github.com/b2network/b2committer/internal/schema"

type TxsRootProposal struct {
	ProposalID       uint64
	StartTimestamp   uint64
	EndTimestamp     uint64
	StartBlockNumber uint64
	EndBlockNumber   uint64
	TxsRoot          string
	BlockList        []uint64
}

func NewTxsRootProposal(proposalID uint64, txsRoot string, blobs []schema.BlobInfo) *TxsRootProposal {
	proposal := &TxsRootProposal{
		ProposalID: proposalID,
		TxsRoot:    txsRoot,
	}
	fillBlockListAndBlockInfo(blobs, proposal)
	return proposal
}

func fillBlockListAndBlockInfo(blobs []schema.BlobInfo, proposal *TxsRootProposal) {
	blockList := make([]uint64, 0)

	for _, blob := range blobs {
		if proposal.StartBlockNumber > uint64(blob.BlockNumber) {
			proposal.StartBlockNumber = uint64(blob.BlockNumber)
		}
		if proposal.EndBlockNumber < uint64(blob.BlockNumber) {
			proposal.EndBlockNumber = uint64(blob.BlockNumber)
		}
		if proposal.StartTimestamp > blob.BlockTime {
			proposal.StartTimestamp = blob.BlockTime
		}
		if proposal.EndTimestamp < blob.BlockTime {
			proposal.EndTimestamp = blob.BlockTime
		}
		blockList = append(blockList, uint64(blob.BlockNumber))
	}
	proposal.BlockList = blockList
}
