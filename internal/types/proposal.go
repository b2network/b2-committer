package types

import (
	"encoding/json"
	"github.com/b2network/b2committer/internal/schema"
	"strconv"
	"strings"
)

type TxsRootProposal struct {
	ProposalID       uint64
	StartTimestamp   uint64
	EndTimestamp     uint64
	StartBlockNumber uint64
	EndBlockNumber   uint64
	TxsRoot          string
	BlockList        []uint64
	BlockListStr     string
}

func convertBlockListToString(blockList []uint64) string {
	strArr := make([]string, 0)
	for _, block := range blockList {
		number := strconv.FormatUint(block, 10)
		if !stringExistsInSlice(strArr, number) {
			strArr = append(strArr, number)
		}
	}
	return strings.Join(strArr, ",")
}

func stringExistsInSlice(slice []string, target string) bool {
	for _, element := range slice {
		if element == target {
			return true
		}
	}
	return false
}

func uintExistsInSlice(slice []uint64, target uint64) bool {
	for _, element := range slice {
		if element == target {
			return true
		}
	}
	return false
}

func NewTxsRootProposal(proposalID uint64, txsRoot string, blobs []schema.BlobInfo) *TxsRootProposal {
	proposal := &TxsRootProposal{
		ProposalID:       proposalID,
		StartBlockNumber: uint64(blobs[0].BlockNumber),
		EndBlockNumber:   uint64(blobs[len(blobs)-1].BlockNumber),
		StartTimestamp:   blobs[0].BlockTime,
		EndTimestamp:     blobs[len(blobs)-1].BlockTime,
		TxsRoot:          txsRoot,
	}
	fillBlockListAndBlockInfo(blobs, proposal)
	proposal.BlockListStr = convertBlockListToString(proposal.BlockList)
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
		if !uintExistsInSlice(blockList, uint64(blob.BlockNumber)) {
			blockList = append(blockList, uint64(blob.BlockNumber))
		}
	}
	proposal.BlockList = blockList
}

type StateRootProposal struct {
	ProposalID         uint64
	OutputRoot         string
	StartL1Timestamp   uint64
	EndL1Timestamp     uint64
	StartL2BlockNumber uint64
	EndL2BlockNumber   uint64
	OutputStartIndex   uint64
	OutputEndIndex     uint64
}

func NewStateRootProposal(proposalID uint64, output string, events []schema.SyncEvent) (*StateRootProposal, error) {
	proposal := &StateRootProposal{
		ProposalID: proposalID,
		OutputRoot: output,
	}
	err := fillStateRootProposal(events, proposal)
	if err != nil {
		return nil, err
	}
	return proposal, nil
}

func fillStateRootProposal(events []schema.SyncEvent, proposal *StateRootProposal) error {
	var data OutputEvent
	for _, event := range events {
		err := json.Unmarshal([]byte(event.Data), &data)
		if err != nil {
			return err
		}
		if proposal.StartL1Timestamp == 0 && proposal.StartL2BlockNumber == 0 && proposal.OutputStartIndex == 0 {
			proposal.StartL1Timestamp = data.L1Timestamp
			proposal.StartL2BlockNumber = data.L2blockNumber
			proposal.OutputStartIndex = data.L2OutputIndex
		}
		if proposal.StartL1Timestamp > data.L1Timestamp {
			proposal.StartL1Timestamp = data.L1Timestamp
		}
		if proposal.EndL1Timestamp < data.L1Timestamp {
			proposal.EndL1Timestamp = data.L1Timestamp
		}
		if proposal.StartL2BlockNumber > data.L2blockNumber {
			proposal.StartL2BlockNumber = data.L2blockNumber
		}
		if proposal.EndL2BlockNumber < data.L2blockNumber {
			proposal.EndL2BlockNumber = data.L2blockNumber
		}
		if proposal.OutputStartIndex > data.L2OutputIndex {
			proposal.OutputStartIndex = data.L2OutputIndex
		}
		if proposal.OutputEndIndex < data.L2OutputIndex {
			proposal.OutputEndIndex = data.L2OutputIndex
		}
	}
	return nil
}

type OutputEvent struct {
	OutputRoot    string `json:"outputRoot"`
	L1Timestamp   uint64 `json:"l1Timestamp"`
	L2blockNumber uint64 `json:"l2blockNumber"`
	L2OutputIndex uint64 `json:"l2OutputIndex"`
}

func ConvertEventDataToOutputEvent(events []schema.SyncEvent) ([]OutputEvent, error) {
	outputs := make([]OutputEvent, 0)
	for _, event := range events {
		var data OutputEvent
		err := json.Unmarshal([]byte(event.Data), &data)
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, data)
	}
	return outputs, nil
}

func ConvertOutputsToEventData(outputs []OutputEvent) ([]schema.SyncEvent, error) {
	events := make([]schema.SyncEvent, 0)
	for _, output := range outputs {
		var event schema.SyncEvent
		data, err := json.Marshal(output)
		if err != nil {
			return nil, err
		}
		event.Data = string(data)
		events = append(events, event)
	}
	return events, nil
}
