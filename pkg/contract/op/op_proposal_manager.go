// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package op

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// OpProposalStateRootProposal is an auto generated low-level Go binding around an user-defined struct.
type OpProposalStateRootProposal struct {
	ProposalID         uint64
	OutputRoot         string
	StartL1Timestamp   uint64
	EndL1Timestamp     uint64
	StartL2BlockNumber uint64
	EndL2BlockNumber   uint64
	OutputStartIndex   uint64
	OutputEndIndex     uint64
	Timeout            *big.Int
	Status             uint8
	DsType             uint8
	DsTxHash           string
	BitcoinTxHash      string
	Winner             common.Address
}

// OpProposalTxsRootProposal is an auto generated low-level Go binding around an user-defined struct.
type OpProposalTxsRootProposal struct {
	ProposalID       uint64
	StartTimestamp   uint64
	EndTimestamp     uint64
	StartBlockNumber uint64
	EndBlockNumber   uint64
	TxsRoot          string
	Timeout          *big.Int
	Status           uint8
	DsType           uint8
	DsTxHash         string
	Winner           common.Address
	BlockList        []uint64
}

// OpProposalManagerMetaData contains all meta data concerning the OpProposalManager contract.
var OpProposalManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"}],\"name\":\"allBitcoinTxVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalType\",\"type\":\"uint64\"}],\"name\":\"allDSTxVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"proposalType\",\"type\":\"uint8\"}],\"name\":\"allProposalVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"proposalType\",\"type\":\"uint8\"}],\"name\":\"cleanVotes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastStateRootProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"outputRoot\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"startL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputStartIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputEndIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dsType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"dsTxHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bitcoinTxHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"internalType\":\"structOpProposal.StateRootProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastTxsRootProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"txsRoot\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dsType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"dsTxHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"blockList\",\"type\":\"uint64[]\"}],\"internalType\":\"structOpProposal.TxsRootProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"}],\"name\":\"getStateRootProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"outputRoot\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"startL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputStartIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputEndIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dsType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"dsTxHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bitcoinTxHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"internalType\":\"structOpProposal.StateRootProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"}],\"name\":\"getTxsRootProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"txsRoot\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dsType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"dsTxHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"blockList\",\"type\":\"uint64[]\"}],\"internalType\":\"structOpProposal.TxsRootProposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVotedOnStateRootDSTxPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVotedOnStateRootProposalPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVotedOnSubmitBitcoinTxPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVotedOnTxsRootProposalPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVotedOntxsRootDSTxPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStateRootProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"outputRoot\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"startL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputStartIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputEndIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dsType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"dsTxHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bitcoinTxHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTxsRootProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"txsRoot\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dsType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"dsTxHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"opCommitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"committer\",\"type\":\"address\"}],\"name\":\"setCommitter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"outputRoot\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"startL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputStartIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputEndIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dsType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"dsTxHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bitcoinTxHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"internalType\":\"structOpProposal.StateRootProposal\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"setLastStateRootProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"txsRoot\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dsType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"dsTxHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"blockList\",\"type\":\"uint64[]\"}],\"internalType\":\"structOpProposal.TxsRootProposal\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"setLastTxsRootProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"outputRoot\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"startL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputStartIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputEndIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dsType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"dsTxHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bitcoinTxHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"internalType\":\"structOpProposal.StateRootProposal\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"setStateRootProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"txsRoot\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dsType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"dsTxHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"blockList\",\"type\":\"uint64[]\"}],\"internalType\":\"structOpProposal.TxsRootProposal\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"name\":\"setTxsRootProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stateBitcoinTxVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stateDSTxVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stateRootProposalVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"stateRootProposals\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"outputRoot\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"startL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputStartIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputEndIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dsType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"dsTxHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bitcoinTxHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"txsDSTxVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"txsRootProposalVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"txsRootProposals\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"txsRoot\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"dsType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"dsTxHash\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"proposers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"voteAndUpdateStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OpProposalManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use OpProposalManagerMetaData.ABI instead.
var OpProposalManagerABI = OpProposalManagerMetaData.ABI

// OpProposalManager is an auto generated Go binding around an Ethereum contract.
type OpProposalManager struct {
	OpProposalManagerCaller     // Read-only binding to the contract
	OpProposalManagerTransactor // Write-only binding to the contract
	OpProposalManagerFilterer   // Log filterer for contract events
}

// OpProposalManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpProposalManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpProposalManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpProposalManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpProposalManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpProposalManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpProposalManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpProposalManagerSession struct {
	Contract     *OpProposalManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OpProposalManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpProposalManagerCallerSession struct {
	Contract *OpProposalManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OpProposalManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpProposalManagerTransactorSession struct {
	Contract     *OpProposalManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OpProposalManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpProposalManagerRaw struct {
	Contract *OpProposalManager // Generic contract binding to access the raw methods on
}

// OpProposalManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpProposalManagerCallerRaw struct {
	Contract *OpProposalManagerCaller // Generic read-only contract binding to access the raw methods on
}

// OpProposalManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpProposalManagerTransactorRaw struct {
	Contract *OpProposalManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpProposalManager creates a new instance of OpProposalManager, bound to a specific deployed contract.
func NewOpProposalManager(address common.Address, backend bind.ContractBackend) (*OpProposalManager, error) {
	contract, err := bindOpProposalManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpProposalManager{OpProposalManagerCaller: OpProposalManagerCaller{contract: contract}, OpProposalManagerTransactor: OpProposalManagerTransactor{contract: contract}, OpProposalManagerFilterer: OpProposalManagerFilterer{contract: contract}}, nil
}

// NewOpProposalManagerCaller creates a new read-only instance of OpProposalManager, bound to a specific deployed contract.
func NewOpProposalManagerCaller(address common.Address, caller bind.ContractCaller) (*OpProposalManagerCaller, error) {
	contract, err := bindOpProposalManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpProposalManagerCaller{contract: contract}, nil
}

// NewOpProposalManagerTransactor creates a new write-only instance of OpProposalManager, bound to a specific deployed contract.
func NewOpProposalManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*OpProposalManagerTransactor, error) {
	contract, err := bindOpProposalManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpProposalManagerTransactor{contract: contract}, nil
}

// NewOpProposalManagerFilterer creates a new log filterer instance of OpProposalManager, bound to a specific deployed contract.
func NewOpProposalManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*OpProposalManagerFilterer, error) {
	contract, err := bindOpProposalManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpProposalManagerFilterer{contract: contract}, nil
}

// bindOpProposalManager binds a generic wrapper to an already deployed contract.
func bindOpProposalManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OpProposalManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpProposalManager *OpProposalManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpProposalManager.Contract.OpProposalManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpProposalManager *OpProposalManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpProposalManager.Contract.OpProposalManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpProposalManager *OpProposalManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpProposalManager.Contract.OpProposalManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpProposalManager *OpProposalManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpProposalManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpProposalManager *OpProposalManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpProposalManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpProposalManager *OpProposalManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpProposalManager.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_OpProposalManager *OpProposalManagerCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_OpProposalManager *OpProposalManagerSession) ADMINROLE() ([32]byte, error) {
	return _OpProposalManager.Contract.ADMINROLE(&_OpProposalManager.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_OpProposalManager *OpProposalManagerCallerSession) ADMINROLE() ([32]byte, error) {
	return _OpProposalManager.Contract.ADMINROLE(&_OpProposalManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OpProposalManager *OpProposalManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OpProposalManager *OpProposalManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OpProposalManager.Contract.DEFAULTADMINROLE(&_OpProposalManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OpProposalManager *OpProposalManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OpProposalManager.Contract.DEFAULTADMINROLE(&_OpProposalManager.CallOpts)
}

// AllBitcoinTxVotes is a free data retrieval call binding the contract method 0x4f8bcb3d.
//
// Solidity: function allBitcoinTxVotes(uint64 proposalID) view returns(address[])
func (_OpProposalManager *OpProposalManagerCaller) AllBitcoinTxVotes(opts *bind.CallOpts, proposalID uint64) ([]common.Address, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "allBitcoinTxVotes", proposalID)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllBitcoinTxVotes is a free data retrieval call binding the contract method 0x4f8bcb3d.
//
// Solidity: function allBitcoinTxVotes(uint64 proposalID) view returns(address[])
func (_OpProposalManager *OpProposalManagerSession) AllBitcoinTxVotes(proposalID uint64) ([]common.Address, error) {
	return _OpProposalManager.Contract.AllBitcoinTxVotes(&_OpProposalManager.CallOpts, proposalID)
}

// AllBitcoinTxVotes is a free data retrieval call binding the contract method 0x4f8bcb3d.
//
// Solidity: function allBitcoinTxVotes(uint64 proposalID) view returns(address[])
func (_OpProposalManager *OpProposalManagerCallerSession) AllBitcoinTxVotes(proposalID uint64) ([]common.Address, error) {
	return _OpProposalManager.Contract.AllBitcoinTxVotes(&_OpProposalManager.CallOpts, proposalID)
}

// AllDSTxVotes is a free data retrieval call binding the contract method 0x44d2e4e0.
//
// Solidity: function allDSTxVotes(uint64 proposalID, uint64 proposalType) view returns(address[])
func (_OpProposalManager *OpProposalManagerCaller) AllDSTxVotes(opts *bind.CallOpts, proposalID uint64, proposalType uint64) ([]common.Address, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "allDSTxVotes", proposalID, proposalType)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllDSTxVotes is a free data retrieval call binding the contract method 0x44d2e4e0.
//
// Solidity: function allDSTxVotes(uint64 proposalID, uint64 proposalType) view returns(address[])
func (_OpProposalManager *OpProposalManagerSession) AllDSTxVotes(proposalID uint64, proposalType uint64) ([]common.Address, error) {
	return _OpProposalManager.Contract.AllDSTxVotes(&_OpProposalManager.CallOpts, proposalID, proposalType)
}

// AllDSTxVotes is a free data retrieval call binding the contract method 0x44d2e4e0.
//
// Solidity: function allDSTxVotes(uint64 proposalID, uint64 proposalType) view returns(address[])
func (_OpProposalManager *OpProposalManagerCallerSession) AllDSTxVotes(proposalID uint64, proposalType uint64) ([]common.Address, error) {
	return _OpProposalManager.Contract.AllDSTxVotes(&_OpProposalManager.CallOpts, proposalID, proposalType)
}

// AllProposalVotes is a free data retrieval call binding the contract method 0x49b97431.
//
// Solidity: function allProposalVotes(uint64 proposalID, uint8 proposalType) view returns(address[])
func (_OpProposalManager *OpProposalManagerCaller) AllProposalVotes(opts *bind.CallOpts, proposalID uint64, proposalType uint8) ([]common.Address, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "allProposalVotes", proposalID, proposalType)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllProposalVotes is a free data retrieval call binding the contract method 0x49b97431.
//
// Solidity: function allProposalVotes(uint64 proposalID, uint8 proposalType) view returns(address[])
func (_OpProposalManager *OpProposalManagerSession) AllProposalVotes(proposalID uint64, proposalType uint8) ([]common.Address, error) {
	return _OpProposalManager.Contract.AllProposalVotes(&_OpProposalManager.CallOpts, proposalID, proposalType)
}

// AllProposalVotes is a free data retrieval call binding the contract method 0x49b97431.
//
// Solidity: function allProposalVotes(uint64 proposalID, uint8 proposalType) view returns(address[])
func (_OpProposalManager *OpProposalManagerCallerSession) AllProposalVotes(proposalID uint64, proposalType uint8) ([]common.Address, error) {
	return _OpProposalManager.Contract.AllProposalVotes(&_OpProposalManager.CallOpts, proposalID, proposalType)
}

// GetLastStateRootProposal is a free data retrieval call binding the contract method 0x418f1657.
//
// Solidity: function getLastStateRootProposal() view returns((uint64,string,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint8,uint8,string,string,address))
func (_OpProposalManager *OpProposalManagerCaller) GetLastStateRootProposal(opts *bind.CallOpts) (OpProposalStateRootProposal, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "getLastStateRootProposal")

	if err != nil {
		return *new(OpProposalStateRootProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(OpProposalStateRootProposal)).(*OpProposalStateRootProposal)

	return out0, err

}

// GetLastStateRootProposal is a free data retrieval call binding the contract method 0x418f1657.
//
// Solidity: function getLastStateRootProposal() view returns((uint64,string,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint8,uint8,string,string,address))
func (_OpProposalManager *OpProposalManagerSession) GetLastStateRootProposal() (OpProposalStateRootProposal, error) {
	return _OpProposalManager.Contract.GetLastStateRootProposal(&_OpProposalManager.CallOpts)
}

// GetLastStateRootProposal is a free data retrieval call binding the contract method 0x418f1657.
//
// Solidity: function getLastStateRootProposal() view returns((uint64,string,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint8,uint8,string,string,address))
func (_OpProposalManager *OpProposalManagerCallerSession) GetLastStateRootProposal() (OpProposalStateRootProposal, error) {
	return _OpProposalManager.Contract.GetLastStateRootProposal(&_OpProposalManager.CallOpts)
}

// GetLastTxsRootProposal is a free data retrieval call binding the contract method 0x182a2cdb.
//
// Solidity: function getLastTxsRootProposal() view returns((uint64,uint64,uint64,uint64,uint64,string,uint256,uint8,uint8,string,address,uint64[]))
func (_OpProposalManager *OpProposalManagerCaller) GetLastTxsRootProposal(opts *bind.CallOpts) (OpProposalTxsRootProposal, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "getLastTxsRootProposal")

	if err != nil {
		return *new(OpProposalTxsRootProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(OpProposalTxsRootProposal)).(*OpProposalTxsRootProposal)

	return out0, err

}

// GetLastTxsRootProposal is a free data retrieval call binding the contract method 0x182a2cdb.
//
// Solidity: function getLastTxsRootProposal() view returns((uint64,uint64,uint64,uint64,uint64,string,uint256,uint8,uint8,string,address,uint64[]))
func (_OpProposalManager *OpProposalManagerSession) GetLastTxsRootProposal() (OpProposalTxsRootProposal, error) {
	return _OpProposalManager.Contract.GetLastTxsRootProposal(&_OpProposalManager.CallOpts)
}

// GetLastTxsRootProposal is a free data retrieval call binding the contract method 0x182a2cdb.
//
// Solidity: function getLastTxsRootProposal() view returns((uint64,uint64,uint64,uint64,uint64,string,uint256,uint8,uint8,string,address,uint64[]))
func (_OpProposalManager *OpProposalManagerCallerSession) GetLastTxsRootProposal() (OpProposalTxsRootProposal, error) {
	return _OpProposalManager.Contract.GetLastTxsRootProposal(&_OpProposalManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OpProposalManager *OpProposalManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OpProposalManager *OpProposalManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OpProposalManager.Contract.GetRoleAdmin(&_OpProposalManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OpProposalManager *OpProposalManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OpProposalManager.Contract.GetRoleAdmin(&_OpProposalManager.CallOpts, role)
}

// GetStateRootProposal is a free data retrieval call binding the contract method 0xacb340b8.
//
// Solidity: function getStateRootProposal(uint64 proposalID) view returns((uint64,string,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint8,uint8,string,string,address))
func (_OpProposalManager *OpProposalManagerCaller) GetStateRootProposal(opts *bind.CallOpts, proposalID uint64) (OpProposalStateRootProposal, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "getStateRootProposal", proposalID)

	if err != nil {
		return *new(OpProposalStateRootProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(OpProposalStateRootProposal)).(*OpProposalStateRootProposal)

	return out0, err

}

// GetStateRootProposal is a free data retrieval call binding the contract method 0xacb340b8.
//
// Solidity: function getStateRootProposal(uint64 proposalID) view returns((uint64,string,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint8,uint8,string,string,address))
func (_OpProposalManager *OpProposalManagerSession) GetStateRootProposal(proposalID uint64) (OpProposalStateRootProposal, error) {
	return _OpProposalManager.Contract.GetStateRootProposal(&_OpProposalManager.CallOpts, proposalID)
}

// GetStateRootProposal is a free data retrieval call binding the contract method 0xacb340b8.
//
// Solidity: function getStateRootProposal(uint64 proposalID) view returns((uint64,string,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint8,uint8,string,string,address))
func (_OpProposalManager *OpProposalManagerCallerSession) GetStateRootProposal(proposalID uint64) (OpProposalStateRootProposal, error) {
	return _OpProposalManager.Contract.GetStateRootProposal(&_OpProposalManager.CallOpts, proposalID)
}

// GetTxsRootProposal is a free data retrieval call binding the contract method 0xac536038.
//
// Solidity: function getTxsRootProposal(uint64 proposalID) view returns((uint64,uint64,uint64,uint64,uint64,string,uint256,uint8,uint8,string,address,uint64[]))
func (_OpProposalManager *OpProposalManagerCaller) GetTxsRootProposal(opts *bind.CallOpts, proposalID uint64) (OpProposalTxsRootProposal, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "getTxsRootProposal", proposalID)

	if err != nil {
		return *new(OpProposalTxsRootProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(OpProposalTxsRootProposal)).(*OpProposalTxsRootProposal)

	return out0, err

}

// GetTxsRootProposal is a free data retrieval call binding the contract method 0xac536038.
//
// Solidity: function getTxsRootProposal(uint64 proposalID) view returns((uint64,uint64,uint64,uint64,uint64,string,uint256,uint8,uint8,string,address,uint64[]))
func (_OpProposalManager *OpProposalManagerSession) GetTxsRootProposal(proposalID uint64) (OpProposalTxsRootProposal, error) {
	return _OpProposalManager.Contract.GetTxsRootProposal(&_OpProposalManager.CallOpts, proposalID)
}

// GetTxsRootProposal is a free data retrieval call binding the contract method 0xac536038.
//
// Solidity: function getTxsRootProposal(uint64 proposalID) view returns((uint64,uint64,uint64,uint64,uint64,string,uint256,uint8,uint8,string,address,uint64[]))
func (_OpProposalManager *OpProposalManagerCallerSession) GetTxsRootProposal(proposalID uint64) (OpProposalTxsRootProposal, error) {
	return _OpProposalManager.Contract.GetTxsRootProposal(&_OpProposalManager.CallOpts, proposalID)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OpProposalManager *OpProposalManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OpProposalManager *OpProposalManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OpProposalManager.Contract.HasRole(&_OpProposalManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OpProposalManager *OpProposalManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OpProposalManager.Contract.HasRole(&_OpProposalManager.CallOpts, role, account)
}

// IsVotedOnStateRootDSTxPhase is a free data retrieval call binding the contract method 0xb48ec5f1.
//
// Solidity: function isVotedOnStateRootDSTxPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerCaller) IsVotedOnStateRootDSTxPhase(opts *bind.CallOpts, proposalID uint64, voter common.Address) (bool, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "isVotedOnStateRootDSTxPhase", proposalID, voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVotedOnStateRootDSTxPhase is a free data retrieval call binding the contract method 0xb48ec5f1.
//
// Solidity: function isVotedOnStateRootDSTxPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerSession) IsVotedOnStateRootDSTxPhase(proposalID uint64, voter common.Address) (bool, error) {
	return _OpProposalManager.Contract.IsVotedOnStateRootDSTxPhase(&_OpProposalManager.CallOpts, proposalID, voter)
}

// IsVotedOnStateRootDSTxPhase is a free data retrieval call binding the contract method 0xb48ec5f1.
//
// Solidity: function isVotedOnStateRootDSTxPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerCallerSession) IsVotedOnStateRootDSTxPhase(proposalID uint64, voter common.Address) (bool, error) {
	return _OpProposalManager.Contract.IsVotedOnStateRootDSTxPhase(&_OpProposalManager.CallOpts, proposalID, voter)
}

// IsVotedOnStateRootProposalPhase is a free data retrieval call binding the contract method 0x641d1f67.
//
// Solidity: function isVotedOnStateRootProposalPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerCaller) IsVotedOnStateRootProposalPhase(opts *bind.CallOpts, proposalID uint64, voter common.Address) (bool, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "isVotedOnStateRootProposalPhase", proposalID, voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVotedOnStateRootProposalPhase is a free data retrieval call binding the contract method 0x641d1f67.
//
// Solidity: function isVotedOnStateRootProposalPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerSession) IsVotedOnStateRootProposalPhase(proposalID uint64, voter common.Address) (bool, error) {
	return _OpProposalManager.Contract.IsVotedOnStateRootProposalPhase(&_OpProposalManager.CallOpts, proposalID, voter)
}

// IsVotedOnStateRootProposalPhase is a free data retrieval call binding the contract method 0x641d1f67.
//
// Solidity: function isVotedOnStateRootProposalPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerCallerSession) IsVotedOnStateRootProposalPhase(proposalID uint64, voter common.Address) (bool, error) {
	return _OpProposalManager.Contract.IsVotedOnStateRootProposalPhase(&_OpProposalManager.CallOpts, proposalID, voter)
}

// IsVotedOnSubmitBitcoinTxPhase is a free data retrieval call binding the contract method 0xfab0c232.
//
// Solidity: function isVotedOnSubmitBitcoinTxPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerCaller) IsVotedOnSubmitBitcoinTxPhase(opts *bind.CallOpts, proposalID uint64, voter common.Address) (bool, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "isVotedOnSubmitBitcoinTxPhase", proposalID, voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVotedOnSubmitBitcoinTxPhase is a free data retrieval call binding the contract method 0xfab0c232.
//
// Solidity: function isVotedOnSubmitBitcoinTxPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerSession) IsVotedOnSubmitBitcoinTxPhase(proposalID uint64, voter common.Address) (bool, error) {
	return _OpProposalManager.Contract.IsVotedOnSubmitBitcoinTxPhase(&_OpProposalManager.CallOpts, proposalID, voter)
}

// IsVotedOnSubmitBitcoinTxPhase is a free data retrieval call binding the contract method 0xfab0c232.
//
// Solidity: function isVotedOnSubmitBitcoinTxPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerCallerSession) IsVotedOnSubmitBitcoinTxPhase(proposalID uint64, voter common.Address) (bool, error) {
	return _OpProposalManager.Contract.IsVotedOnSubmitBitcoinTxPhase(&_OpProposalManager.CallOpts, proposalID, voter)
}

// IsVotedOnTxsRootProposalPhase is a free data retrieval call binding the contract method 0xb698fcea.
//
// Solidity: function isVotedOnTxsRootProposalPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerCaller) IsVotedOnTxsRootProposalPhase(opts *bind.CallOpts, proposalID uint64, voter common.Address) (bool, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "isVotedOnTxsRootProposalPhase", proposalID, voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVotedOnTxsRootProposalPhase is a free data retrieval call binding the contract method 0xb698fcea.
//
// Solidity: function isVotedOnTxsRootProposalPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerSession) IsVotedOnTxsRootProposalPhase(proposalID uint64, voter common.Address) (bool, error) {
	return _OpProposalManager.Contract.IsVotedOnTxsRootProposalPhase(&_OpProposalManager.CallOpts, proposalID, voter)
}

// IsVotedOnTxsRootProposalPhase is a free data retrieval call binding the contract method 0xb698fcea.
//
// Solidity: function isVotedOnTxsRootProposalPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerCallerSession) IsVotedOnTxsRootProposalPhase(proposalID uint64, voter common.Address) (bool, error) {
	return _OpProposalManager.Contract.IsVotedOnTxsRootProposalPhase(&_OpProposalManager.CallOpts, proposalID, voter)
}

// IsVotedOntxsRootDSTxPhase is a free data retrieval call binding the contract method 0x34c8f27e.
//
// Solidity: function isVotedOntxsRootDSTxPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerCaller) IsVotedOntxsRootDSTxPhase(opts *bind.CallOpts, proposalID uint64, voter common.Address) (bool, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "isVotedOntxsRootDSTxPhase", proposalID, voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVotedOntxsRootDSTxPhase is a free data retrieval call binding the contract method 0x34c8f27e.
//
// Solidity: function isVotedOntxsRootDSTxPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerSession) IsVotedOntxsRootDSTxPhase(proposalID uint64, voter common.Address) (bool, error) {
	return _OpProposalManager.Contract.IsVotedOntxsRootDSTxPhase(&_OpProposalManager.CallOpts, proposalID, voter)
}

// IsVotedOntxsRootDSTxPhase is a free data retrieval call binding the contract method 0x34c8f27e.
//
// Solidity: function isVotedOntxsRootDSTxPhase(uint64 proposalID, address voter) view returns(bool)
func (_OpProposalManager *OpProposalManagerCallerSession) IsVotedOntxsRootDSTxPhase(proposalID uint64, voter common.Address) (bool, error) {
	return _OpProposalManager.Contract.IsVotedOntxsRootDSTxPhase(&_OpProposalManager.CallOpts, proposalID, voter)
}

// LastStateRootProposal is a free data retrieval call binding the contract method 0x50613301.
//
// Solidity: function lastStateRootProposal() view returns(uint64 proposalID, string outputRoot, uint64 startL1Timestamp, uint64 endL1Timestamp, uint64 startL2BlockNumber, uint64 endL2BlockNumber, uint64 outputStartIndex, uint64 outputEndIndex, uint256 timeout, uint8 status, uint8 dsType, string dsTxHash, string bitcoinTxHash, address winner)
func (_OpProposalManager *OpProposalManagerCaller) LastStateRootProposal(opts *bind.CallOpts) (struct {
	ProposalID         uint64
	OutputRoot         string
	StartL1Timestamp   uint64
	EndL1Timestamp     uint64
	StartL2BlockNumber uint64
	EndL2BlockNumber   uint64
	OutputStartIndex   uint64
	OutputEndIndex     uint64
	Timeout            *big.Int
	Status             uint8
	DsType             uint8
	DsTxHash           string
	BitcoinTxHash      string
	Winner             common.Address
}, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "lastStateRootProposal")

	outstruct := new(struct {
		ProposalID         uint64
		OutputRoot         string
		StartL1Timestamp   uint64
		EndL1Timestamp     uint64
		StartL2BlockNumber uint64
		EndL2BlockNumber   uint64
		OutputStartIndex   uint64
		OutputEndIndex     uint64
		Timeout            *big.Int
		Status             uint8
		DsType             uint8
		DsTxHash           string
		BitcoinTxHash      string
		Winner             common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposalID = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.OutputRoot = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.StartL1Timestamp = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.EndL1Timestamp = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.StartL2BlockNumber = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.EndL2BlockNumber = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.OutputStartIndex = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.OutputEndIndex = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.Timeout = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[9], new(uint8)).(*uint8)
	outstruct.DsType = *abi.ConvertType(out[10], new(uint8)).(*uint8)
	outstruct.DsTxHash = *abi.ConvertType(out[11], new(string)).(*string)
	outstruct.BitcoinTxHash = *abi.ConvertType(out[12], new(string)).(*string)
	outstruct.Winner = *abi.ConvertType(out[13], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// LastStateRootProposal is a free data retrieval call binding the contract method 0x50613301.
//
// Solidity: function lastStateRootProposal() view returns(uint64 proposalID, string outputRoot, uint64 startL1Timestamp, uint64 endL1Timestamp, uint64 startL2BlockNumber, uint64 endL2BlockNumber, uint64 outputStartIndex, uint64 outputEndIndex, uint256 timeout, uint8 status, uint8 dsType, string dsTxHash, string bitcoinTxHash, address winner)
func (_OpProposalManager *OpProposalManagerSession) LastStateRootProposal() (struct {
	ProposalID         uint64
	OutputRoot         string
	StartL1Timestamp   uint64
	EndL1Timestamp     uint64
	StartL2BlockNumber uint64
	EndL2BlockNumber   uint64
	OutputStartIndex   uint64
	OutputEndIndex     uint64
	Timeout            *big.Int
	Status             uint8
	DsType             uint8
	DsTxHash           string
	BitcoinTxHash      string
	Winner             common.Address
}, error) {
	return _OpProposalManager.Contract.LastStateRootProposal(&_OpProposalManager.CallOpts)
}

// LastStateRootProposal is a free data retrieval call binding the contract method 0x50613301.
//
// Solidity: function lastStateRootProposal() view returns(uint64 proposalID, string outputRoot, uint64 startL1Timestamp, uint64 endL1Timestamp, uint64 startL2BlockNumber, uint64 endL2BlockNumber, uint64 outputStartIndex, uint64 outputEndIndex, uint256 timeout, uint8 status, uint8 dsType, string dsTxHash, string bitcoinTxHash, address winner)
func (_OpProposalManager *OpProposalManagerCallerSession) LastStateRootProposal() (struct {
	ProposalID         uint64
	OutputRoot         string
	StartL1Timestamp   uint64
	EndL1Timestamp     uint64
	StartL2BlockNumber uint64
	EndL2BlockNumber   uint64
	OutputStartIndex   uint64
	OutputEndIndex     uint64
	Timeout            *big.Int
	Status             uint8
	DsType             uint8
	DsTxHash           string
	BitcoinTxHash      string
	Winner             common.Address
}, error) {
	return _OpProposalManager.Contract.LastStateRootProposal(&_OpProposalManager.CallOpts)
}

// LastTxsRootProposal is a free data retrieval call binding the contract method 0xdba274ab.
//
// Solidity: function lastTxsRootProposal() view returns(uint64 proposalID, uint64 startTimestamp, uint64 endTimestamp, uint64 startBlockNumber, uint64 endBlockNumber, string txsRoot, uint256 timeout, uint8 status, uint8 dsType, string dsTxHash, address winner)
func (_OpProposalManager *OpProposalManagerCaller) LastTxsRootProposal(opts *bind.CallOpts) (struct {
	ProposalID       uint64
	StartTimestamp   uint64
	EndTimestamp     uint64
	StartBlockNumber uint64
	EndBlockNumber   uint64
	TxsRoot          string
	Timeout          *big.Int
	Status           uint8
	DsType           uint8
	DsTxHash         string
	Winner           common.Address
}, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "lastTxsRootProposal")

	outstruct := new(struct {
		ProposalID       uint64
		StartTimestamp   uint64
		EndTimestamp     uint64
		StartBlockNumber uint64
		EndBlockNumber   uint64
		TxsRoot          string
		Timeout          *big.Int
		Status           uint8
		DsType           uint8
		DsTxHash         string
		Winner           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposalID = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.StartTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.EndTimestamp = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.StartBlockNumber = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.EndBlockNumber = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.TxsRoot = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Timeout = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.DsType = *abi.ConvertType(out[8], new(uint8)).(*uint8)
	outstruct.DsTxHash = *abi.ConvertType(out[9], new(string)).(*string)
	outstruct.Winner = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// LastTxsRootProposal is a free data retrieval call binding the contract method 0xdba274ab.
//
// Solidity: function lastTxsRootProposal() view returns(uint64 proposalID, uint64 startTimestamp, uint64 endTimestamp, uint64 startBlockNumber, uint64 endBlockNumber, string txsRoot, uint256 timeout, uint8 status, uint8 dsType, string dsTxHash, address winner)
func (_OpProposalManager *OpProposalManagerSession) LastTxsRootProposal() (struct {
	ProposalID       uint64
	StartTimestamp   uint64
	EndTimestamp     uint64
	StartBlockNumber uint64
	EndBlockNumber   uint64
	TxsRoot          string
	Timeout          *big.Int
	Status           uint8
	DsType           uint8
	DsTxHash         string
	Winner           common.Address
}, error) {
	return _OpProposalManager.Contract.LastTxsRootProposal(&_OpProposalManager.CallOpts)
}

// LastTxsRootProposal is a free data retrieval call binding the contract method 0xdba274ab.
//
// Solidity: function lastTxsRootProposal() view returns(uint64 proposalID, uint64 startTimestamp, uint64 endTimestamp, uint64 startBlockNumber, uint64 endBlockNumber, string txsRoot, uint256 timeout, uint8 status, uint8 dsType, string dsTxHash, address winner)
func (_OpProposalManager *OpProposalManagerCallerSession) LastTxsRootProposal() (struct {
	ProposalID       uint64
	StartTimestamp   uint64
	EndTimestamp     uint64
	StartBlockNumber uint64
	EndBlockNumber   uint64
	TxsRoot          string
	Timeout          *big.Int
	Status           uint8
	DsType           uint8
	DsTxHash         string
	Winner           common.Address
}, error) {
	return _OpProposalManager.Contract.LastTxsRootProposal(&_OpProposalManager.CallOpts)
}

// OpCommitter is a free data retrieval call binding the contract method 0x960a5fce.
//
// Solidity: function opCommitter() view returns(address)
func (_OpProposalManager *OpProposalManagerCaller) OpCommitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "opCommitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OpCommitter is a free data retrieval call binding the contract method 0x960a5fce.
//
// Solidity: function opCommitter() view returns(address)
func (_OpProposalManager *OpProposalManagerSession) OpCommitter() (common.Address, error) {
	return _OpProposalManager.Contract.OpCommitter(&_OpProposalManager.CallOpts)
}

// OpCommitter is a free data retrieval call binding the contract method 0x960a5fce.
//
// Solidity: function opCommitter() view returns(address)
func (_OpProposalManager *OpProposalManagerCallerSession) OpCommitter() (common.Address, error) {
	return _OpProposalManager.Contract.OpCommitter(&_OpProposalManager.CallOpts)
}

// StateBitcoinTxVotes is a free data retrieval call binding the contract method 0xf8ebaea9.
//
// Solidity: function stateBitcoinTxVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerCaller) StateBitcoinTxVotes(opts *bind.CallOpts, arg0 uint64, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "stateBitcoinTxVotes", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateBitcoinTxVotes is a free data retrieval call binding the contract method 0xf8ebaea9.
//
// Solidity: function stateBitcoinTxVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerSession) StateBitcoinTxVotes(arg0 uint64, arg1 *big.Int) (common.Address, error) {
	return _OpProposalManager.Contract.StateBitcoinTxVotes(&_OpProposalManager.CallOpts, arg0, arg1)
}

// StateBitcoinTxVotes is a free data retrieval call binding the contract method 0xf8ebaea9.
//
// Solidity: function stateBitcoinTxVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerCallerSession) StateBitcoinTxVotes(arg0 uint64, arg1 *big.Int) (common.Address, error) {
	return _OpProposalManager.Contract.StateBitcoinTxVotes(&_OpProposalManager.CallOpts, arg0, arg1)
}

// StateDSTxVotes is a free data retrieval call binding the contract method 0x5914f79b.
//
// Solidity: function stateDSTxVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerCaller) StateDSTxVotes(opts *bind.CallOpts, arg0 uint64, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "stateDSTxVotes", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateDSTxVotes is a free data retrieval call binding the contract method 0x5914f79b.
//
// Solidity: function stateDSTxVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerSession) StateDSTxVotes(arg0 uint64, arg1 *big.Int) (common.Address, error) {
	return _OpProposalManager.Contract.StateDSTxVotes(&_OpProposalManager.CallOpts, arg0, arg1)
}

// StateDSTxVotes is a free data retrieval call binding the contract method 0x5914f79b.
//
// Solidity: function stateDSTxVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerCallerSession) StateDSTxVotes(arg0 uint64, arg1 *big.Int) (common.Address, error) {
	return _OpProposalManager.Contract.StateDSTxVotes(&_OpProposalManager.CallOpts, arg0, arg1)
}

// StateRootProposalVotes is a free data retrieval call binding the contract method 0x2820260d.
//
// Solidity: function stateRootProposalVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerCaller) StateRootProposalVotes(opts *bind.CallOpts, arg0 uint64, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "stateRootProposalVotes", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateRootProposalVotes is a free data retrieval call binding the contract method 0x2820260d.
//
// Solidity: function stateRootProposalVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerSession) StateRootProposalVotes(arg0 uint64, arg1 *big.Int) (common.Address, error) {
	return _OpProposalManager.Contract.StateRootProposalVotes(&_OpProposalManager.CallOpts, arg0, arg1)
}

// StateRootProposalVotes is a free data retrieval call binding the contract method 0x2820260d.
//
// Solidity: function stateRootProposalVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerCallerSession) StateRootProposalVotes(arg0 uint64, arg1 *big.Int) (common.Address, error) {
	return _OpProposalManager.Contract.StateRootProposalVotes(&_OpProposalManager.CallOpts, arg0, arg1)
}

// StateRootProposals is a free data retrieval call binding the contract method 0x09a99a33.
//
// Solidity: function stateRootProposals(uint64 ) view returns(uint64 proposalID, string outputRoot, uint64 startL1Timestamp, uint64 endL1Timestamp, uint64 startL2BlockNumber, uint64 endL2BlockNumber, uint64 outputStartIndex, uint64 outputEndIndex, uint256 timeout, uint8 status, uint8 dsType, string dsTxHash, string bitcoinTxHash, address winner)
func (_OpProposalManager *OpProposalManagerCaller) StateRootProposals(opts *bind.CallOpts, arg0 uint64) (struct {
	ProposalID         uint64
	OutputRoot         string
	StartL1Timestamp   uint64
	EndL1Timestamp     uint64
	StartL2BlockNumber uint64
	EndL2BlockNumber   uint64
	OutputStartIndex   uint64
	OutputEndIndex     uint64
	Timeout            *big.Int
	Status             uint8
	DsType             uint8
	DsTxHash           string
	BitcoinTxHash      string
	Winner             common.Address
}, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "stateRootProposals", arg0)

	outstruct := new(struct {
		ProposalID         uint64
		OutputRoot         string
		StartL1Timestamp   uint64
		EndL1Timestamp     uint64
		StartL2BlockNumber uint64
		EndL2BlockNumber   uint64
		OutputStartIndex   uint64
		OutputEndIndex     uint64
		Timeout            *big.Int
		Status             uint8
		DsType             uint8
		DsTxHash           string
		BitcoinTxHash      string
		Winner             common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposalID = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.OutputRoot = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.StartL1Timestamp = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.EndL1Timestamp = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.StartL2BlockNumber = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.EndL2BlockNumber = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.OutputStartIndex = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.OutputEndIndex = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.Timeout = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[9], new(uint8)).(*uint8)
	outstruct.DsType = *abi.ConvertType(out[10], new(uint8)).(*uint8)
	outstruct.DsTxHash = *abi.ConvertType(out[11], new(string)).(*string)
	outstruct.BitcoinTxHash = *abi.ConvertType(out[12], new(string)).(*string)
	outstruct.Winner = *abi.ConvertType(out[13], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// StateRootProposals is a free data retrieval call binding the contract method 0x09a99a33.
//
// Solidity: function stateRootProposals(uint64 ) view returns(uint64 proposalID, string outputRoot, uint64 startL1Timestamp, uint64 endL1Timestamp, uint64 startL2BlockNumber, uint64 endL2BlockNumber, uint64 outputStartIndex, uint64 outputEndIndex, uint256 timeout, uint8 status, uint8 dsType, string dsTxHash, string bitcoinTxHash, address winner)
func (_OpProposalManager *OpProposalManagerSession) StateRootProposals(arg0 uint64) (struct {
	ProposalID         uint64
	OutputRoot         string
	StartL1Timestamp   uint64
	EndL1Timestamp     uint64
	StartL2BlockNumber uint64
	EndL2BlockNumber   uint64
	OutputStartIndex   uint64
	OutputEndIndex     uint64
	Timeout            *big.Int
	Status             uint8
	DsType             uint8
	DsTxHash           string
	BitcoinTxHash      string
	Winner             common.Address
}, error) {
	return _OpProposalManager.Contract.StateRootProposals(&_OpProposalManager.CallOpts, arg0)
}

// StateRootProposals is a free data retrieval call binding the contract method 0x09a99a33.
//
// Solidity: function stateRootProposals(uint64 ) view returns(uint64 proposalID, string outputRoot, uint64 startL1Timestamp, uint64 endL1Timestamp, uint64 startL2BlockNumber, uint64 endL2BlockNumber, uint64 outputStartIndex, uint64 outputEndIndex, uint256 timeout, uint8 status, uint8 dsType, string dsTxHash, string bitcoinTxHash, address winner)
func (_OpProposalManager *OpProposalManagerCallerSession) StateRootProposals(arg0 uint64) (struct {
	ProposalID         uint64
	OutputRoot         string
	StartL1Timestamp   uint64
	EndL1Timestamp     uint64
	StartL2BlockNumber uint64
	EndL2BlockNumber   uint64
	OutputStartIndex   uint64
	OutputEndIndex     uint64
	Timeout            *big.Int
	Status             uint8
	DsType             uint8
	DsTxHash           string
	BitcoinTxHash      string
	Winner             common.Address
}, error) {
	return _OpProposalManager.Contract.StateRootProposals(&_OpProposalManager.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OpProposalManager *OpProposalManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OpProposalManager *OpProposalManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OpProposalManager.Contract.SupportsInterface(&_OpProposalManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OpProposalManager *OpProposalManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OpProposalManager.Contract.SupportsInterface(&_OpProposalManager.CallOpts, interfaceId)
}

// TxsDSTxVotes is a free data retrieval call binding the contract method 0x7aef9daa.
//
// Solidity: function txsDSTxVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerCaller) TxsDSTxVotes(opts *bind.CallOpts, arg0 uint64, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "txsDSTxVotes", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TxsDSTxVotes is a free data retrieval call binding the contract method 0x7aef9daa.
//
// Solidity: function txsDSTxVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerSession) TxsDSTxVotes(arg0 uint64, arg1 *big.Int) (common.Address, error) {
	return _OpProposalManager.Contract.TxsDSTxVotes(&_OpProposalManager.CallOpts, arg0, arg1)
}

// TxsDSTxVotes is a free data retrieval call binding the contract method 0x7aef9daa.
//
// Solidity: function txsDSTxVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerCallerSession) TxsDSTxVotes(arg0 uint64, arg1 *big.Int) (common.Address, error) {
	return _OpProposalManager.Contract.TxsDSTxVotes(&_OpProposalManager.CallOpts, arg0, arg1)
}

// TxsRootProposalVotes is a free data retrieval call binding the contract method 0x79d4f0bc.
//
// Solidity: function txsRootProposalVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerCaller) TxsRootProposalVotes(opts *bind.CallOpts, arg0 uint64, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "txsRootProposalVotes", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TxsRootProposalVotes is a free data retrieval call binding the contract method 0x79d4f0bc.
//
// Solidity: function txsRootProposalVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerSession) TxsRootProposalVotes(arg0 uint64, arg1 *big.Int) (common.Address, error) {
	return _OpProposalManager.Contract.TxsRootProposalVotes(&_OpProposalManager.CallOpts, arg0, arg1)
}

// TxsRootProposalVotes is a free data retrieval call binding the contract method 0x79d4f0bc.
//
// Solidity: function txsRootProposalVotes(uint64 , uint256 ) view returns(address)
func (_OpProposalManager *OpProposalManagerCallerSession) TxsRootProposalVotes(arg0 uint64, arg1 *big.Int) (common.Address, error) {
	return _OpProposalManager.Contract.TxsRootProposalVotes(&_OpProposalManager.CallOpts, arg0, arg1)
}

// TxsRootProposals is a free data retrieval call binding the contract method 0x4a568b6a.
//
// Solidity: function txsRootProposals(uint64 ) view returns(uint64 proposalID, uint64 startTimestamp, uint64 endTimestamp, uint64 startBlockNumber, uint64 endBlockNumber, string txsRoot, uint256 timeout, uint8 status, uint8 dsType, string dsTxHash, address winner)
func (_OpProposalManager *OpProposalManagerCaller) TxsRootProposals(opts *bind.CallOpts, arg0 uint64) (struct {
	ProposalID       uint64
	StartTimestamp   uint64
	EndTimestamp     uint64
	StartBlockNumber uint64
	EndBlockNumber   uint64
	TxsRoot          string
	Timeout          *big.Int
	Status           uint8
	DsType           uint8
	DsTxHash         string
	Winner           common.Address
}, error) {
	var out []interface{}
	err := _OpProposalManager.contract.Call(opts, &out, "txsRootProposals", arg0)

	outstruct := new(struct {
		ProposalID       uint64
		StartTimestamp   uint64
		EndTimestamp     uint64
		StartBlockNumber uint64
		EndBlockNumber   uint64
		TxsRoot          string
		Timeout          *big.Int
		Status           uint8
		DsType           uint8
		DsTxHash         string
		Winner           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposalID = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.StartTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.EndTimestamp = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.StartBlockNumber = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.EndBlockNumber = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.TxsRoot = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Timeout = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.DsType = *abi.ConvertType(out[8], new(uint8)).(*uint8)
	outstruct.DsTxHash = *abi.ConvertType(out[9], new(string)).(*string)
	outstruct.Winner = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// TxsRootProposals is a free data retrieval call binding the contract method 0x4a568b6a.
//
// Solidity: function txsRootProposals(uint64 ) view returns(uint64 proposalID, uint64 startTimestamp, uint64 endTimestamp, uint64 startBlockNumber, uint64 endBlockNumber, string txsRoot, uint256 timeout, uint8 status, uint8 dsType, string dsTxHash, address winner)
func (_OpProposalManager *OpProposalManagerSession) TxsRootProposals(arg0 uint64) (struct {
	ProposalID       uint64
	StartTimestamp   uint64
	EndTimestamp     uint64
	StartBlockNumber uint64
	EndBlockNumber   uint64
	TxsRoot          string
	Timeout          *big.Int
	Status           uint8
	DsType           uint8
	DsTxHash         string
	Winner           common.Address
}, error) {
	return _OpProposalManager.Contract.TxsRootProposals(&_OpProposalManager.CallOpts, arg0)
}

// TxsRootProposals is a free data retrieval call binding the contract method 0x4a568b6a.
//
// Solidity: function txsRootProposals(uint64 ) view returns(uint64 proposalID, uint64 startTimestamp, uint64 endTimestamp, uint64 startBlockNumber, uint64 endBlockNumber, string txsRoot, uint256 timeout, uint8 status, uint8 dsType, string dsTxHash, address winner)
func (_OpProposalManager *OpProposalManagerCallerSession) TxsRootProposals(arg0 uint64) (struct {
	ProposalID       uint64
	StartTimestamp   uint64
	EndTimestamp     uint64
	StartBlockNumber uint64
	EndBlockNumber   uint64
	TxsRoot          string
	Timeout          *big.Int
	Status           uint8
	DsType           uint8
	DsTxHash         string
	Winner           common.Address
}, error) {
	return _OpProposalManager.Contract.TxsRootProposals(&_OpProposalManager.CallOpts, arg0)
}

// CleanVotes is a paid mutator transaction binding the contract method 0x09fa1b23.
//
// Solidity: function cleanVotes(uint64 proposalID, uint8 proposalType) returns()
func (_OpProposalManager *OpProposalManagerTransactor) CleanVotes(opts *bind.TransactOpts, proposalID uint64, proposalType uint8) (*types.Transaction, error) {
	return _OpProposalManager.contract.Transact(opts, "cleanVotes", proposalID, proposalType)
}

// CleanVotes is a paid mutator transaction binding the contract method 0x09fa1b23.
//
// Solidity: function cleanVotes(uint64 proposalID, uint8 proposalType) returns()
func (_OpProposalManager *OpProposalManagerSession) CleanVotes(proposalID uint64, proposalType uint8) (*types.Transaction, error) {
	return _OpProposalManager.Contract.CleanVotes(&_OpProposalManager.TransactOpts, proposalID, proposalType)
}

// CleanVotes is a paid mutator transaction binding the contract method 0x09fa1b23.
//
// Solidity: function cleanVotes(uint64 proposalID, uint8 proposalType) returns()
func (_OpProposalManager *OpProposalManagerTransactorSession) CleanVotes(proposalID uint64, proposalType uint8) (*types.Transaction, error) {
	return _OpProposalManager.Contract.CleanVotes(&_OpProposalManager.TransactOpts, proposalID, proposalType)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OpProposalManager *OpProposalManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpProposalManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OpProposalManager *OpProposalManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpProposalManager.Contract.GrantRole(&_OpProposalManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OpProposalManager *OpProposalManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpProposalManager.Contract.GrantRole(&_OpProposalManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OpProposalManager *OpProposalManagerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpProposalManager.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OpProposalManager *OpProposalManagerSession) Initialize() (*types.Transaction, error) {
	return _OpProposalManager.Contract.Initialize(&_OpProposalManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OpProposalManager *OpProposalManagerTransactorSession) Initialize() (*types.Transaction, error) {
	return _OpProposalManager.Contract.Initialize(&_OpProposalManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OpProposalManager *OpProposalManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OpProposalManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OpProposalManager *OpProposalManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OpProposalManager.Contract.RenounceRole(&_OpProposalManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OpProposalManager *OpProposalManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OpProposalManager.Contract.RenounceRole(&_OpProposalManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OpProposalManager *OpProposalManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpProposalManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OpProposalManager *OpProposalManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpProposalManager.Contract.RevokeRole(&_OpProposalManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OpProposalManager *OpProposalManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpProposalManager.Contract.RevokeRole(&_OpProposalManager.TransactOpts, role, account)
}

// SetCommitter is a paid mutator transaction binding the contract method 0xdd51ce22.
//
// Solidity: function setCommitter(address committer) returns()
func (_OpProposalManager *OpProposalManagerTransactor) SetCommitter(opts *bind.TransactOpts, committer common.Address) (*types.Transaction, error) {
	return _OpProposalManager.contract.Transact(opts, "setCommitter", committer)
}

// SetCommitter is a paid mutator transaction binding the contract method 0xdd51ce22.
//
// Solidity: function setCommitter(address committer) returns()
func (_OpProposalManager *OpProposalManagerSession) SetCommitter(committer common.Address) (*types.Transaction, error) {
	return _OpProposalManager.Contract.SetCommitter(&_OpProposalManager.TransactOpts, committer)
}

// SetCommitter is a paid mutator transaction binding the contract method 0xdd51ce22.
//
// Solidity: function setCommitter(address committer) returns()
func (_OpProposalManager *OpProposalManagerTransactorSession) SetCommitter(committer common.Address) (*types.Transaction, error) {
	return _OpProposalManager.Contract.SetCommitter(&_OpProposalManager.TransactOpts, committer)
}

// SetLastStateRootProposal is a paid mutator transaction binding the contract method 0x9cb06586.
//
// Solidity: function setLastStateRootProposal((uint64,string,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint8,uint8,string,string,address) proposal) returns()
func (_OpProposalManager *OpProposalManagerTransactor) SetLastStateRootProposal(opts *bind.TransactOpts, proposal OpProposalStateRootProposal) (*types.Transaction, error) {
	return _OpProposalManager.contract.Transact(opts, "setLastStateRootProposal", proposal)
}

// SetLastStateRootProposal is a paid mutator transaction binding the contract method 0x9cb06586.
//
// Solidity: function setLastStateRootProposal((uint64,string,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint8,uint8,string,string,address) proposal) returns()
func (_OpProposalManager *OpProposalManagerSession) SetLastStateRootProposal(proposal OpProposalStateRootProposal) (*types.Transaction, error) {
	return _OpProposalManager.Contract.SetLastStateRootProposal(&_OpProposalManager.TransactOpts, proposal)
}

// SetLastStateRootProposal is a paid mutator transaction binding the contract method 0x9cb06586.
//
// Solidity: function setLastStateRootProposal((uint64,string,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint8,uint8,string,string,address) proposal) returns()
func (_OpProposalManager *OpProposalManagerTransactorSession) SetLastStateRootProposal(proposal OpProposalStateRootProposal) (*types.Transaction, error) {
	return _OpProposalManager.Contract.SetLastStateRootProposal(&_OpProposalManager.TransactOpts, proposal)
}

// SetLastTxsRootProposal is a paid mutator transaction binding the contract method 0x4d377509.
//
// Solidity: function setLastTxsRootProposal((uint64,uint64,uint64,uint64,uint64,string,uint256,uint8,uint8,string,address,uint64[]) proposal) returns()
func (_OpProposalManager *OpProposalManagerTransactor) SetLastTxsRootProposal(opts *bind.TransactOpts, proposal OpProposalTxsRootProposal) (*types.Transaction, error) {
	return _OpProposalManager.contract.Transact(opts, "setLastTxsRootProposal", proposal)
}

// SetLastTxsRootProposal is a paid mutator transaction binding the contract method 0x4d377509.
//
// Solidity: function setLastTxsRootProposal((uint64,uint64,uint64,uint64,uint64,string,uint256,uint8,uint8,string,address,uint64[]) proposal) returns()
func (_OpProposalManager *OpProposalManagerSession) SetLastTxsRootProposal(proposal OpProposalTxsRootProposal) (*types.Transaction, error) {
	return _OpProposalManager.Contract.SetLastTxsRootProposal(&_OpProposalManager.TransactOpts, proposal)
}

// SetLastTxsRootProposal is a paid mutator transaction binding the contract method 0x4d377509.
//
// Solidity: function setLastTxsRootProposal((uint64,uint64,uint64,uint64,uint64,string,uint256,uint8,uint8,string,address,uint64[]) proposal) returns()
func (_OpProposalManager *OpProposalManagerTransactorSession) SetLastTxsRootProposal(proposal OpProposalTxsRootProposal) (*types.Transaction, error) {
	return _OpProposalManager.Contract.SetLastTxsRootProposal(&_OpProposalManager.TransactOpts, proposal)
}

// SetStateRootProposal is a paid mutator transaction binding the contract method 0x7aec53cb.
//
// Solidity: function setStateRootProposal(uint64 proposalID, (uint64,string,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint8,uint8,string,string,address) proposal) returns()
func (_OpProposalManager *OpProposalManagerTransactor) SetStateRootProposal(opts *bind.TransactOpts, proposalID uint64, proposal OpProposalStateRootProposal) (*types.Transaction, error) {
	return _OpProposalManager.contract.Transact(opts, "setStateRootProposal", proposalID, proposal)
}

// SetStateRootProposal is a paid mutator transaction binding the contract method 0x7aec53cb.
//
// Solidity: function setStateRootProposal(uint64 proposalID, (uint64,string,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint8,uint8,string,string,address) proposal) returns()
func (_OpProposalManager *OpProposalManagerSession) SetStateRootProposal(proposalID uint64, proposal OpProposalStateRootProposal) (*types.Transaction, error) {
	return _OpProposalManager.Contract.SetStateRootProposal(&_OpProposalManager.TransactOpts, proposalID, proposal)
}

// SetStateRootProposal is a paid mutator transaction binding the contract method 0x7aec53cb.
//
// Solidity: function setStateRootProposal(uint64 proposalID, (uint64,string,uint64,uint64,uint64,uint64,uint64,uint64,uint256,uint8,uint8,string,string,address) proposal) returns()
func (_OpProposalManager *OpProposalManagerTransactorSession) SetStateRootProposal(proposalID uint64, proposal OpProposalStateRootProposal) (*types.Transaction, error) {
	return _OpProposalManager.Contract.SetStateRootProposal(&_OpProposalManager.TransactOpts, proposalID, proposal)
}

// SetTxsRootProposal is a paid mutator transaction binding the contract method 0xf5a40448.
//
// Solidity: function setTxsRootProposal(uint64 proposalID, (uint64,uint64,uint64,uint64,uint64,string,uint256,uint8,uint8,string,address,uint64[]) proposal) returns()
func (_OpProposalManager *OpProposalManagerTransactor) SetTxsRootProposal(opts *bind.TransactOpts, proposalID uint64, proposal OpProposalTxsRootProposal) (*types.Transaction, error) {
	return _OpProposalManager.contract.Transact(opts, "setTxsRootProposal", proposalID, proposal)
}

// SetTxsRootProposal is a paid mutator transaction binding the contract method 0xf5a40448.
//
// Solidity: function setTxsRootProposal(uint64 proposalID, (uint64,uint64,uint64,uint64,uint64,string,uint256,uint8,uint8,string,address,uint64[]) proposal) returns()
func (_OpProposalManager *OpProposalManagerSession) SetTxsRootProposal(proposalID uint64, proposal OpProposalTxsRootProposal) (*types.Transaction, error) {
	return _OpProposalManager.Contract.SetTxsRootProposal(&_OpProposalManager.TransactOpts, proposalID, proposal)
}

// SetTxsRootProposal is a paid mutator transaction binding the contract method 0xf5a40448.
//
// Solidity: function setTxsRootProposal(uint64 proposalID, (uint64,uint64,uint64,uint64,uint64,string,uint256,uint8,uint8,string,address,uint64[]) proposal) returns()
func (_OpProposalManager *OpProposalManagerTransactorSession) SetTxsRootProposal(proposalID uint64, proposal OpProposalTxsRootProposal) (*types.Transaction, error) {
	return _OpProposalManager.Contract.SetTxsRootProposal(&_OpProposalManager.TransactOpts, proposalID, proposal)
}

// VoteAndUpdateStatus is a paid mutator transaction binding the contract method 0xe4f5c160.
//
// Solidity: function voteAndUpdateStatus(uint64 proposalID, uint8 proposalType, address voter, address[] proposers, uint8 phase, uint8 status) returns()
func (_OpProposalManager *OpProposalManagerTransactor) VoteAndUpdateStatus(opts *bind.TransactOpts, proposalID uint64, proposalType uint8, voter common.Address, proposers []common.Address, phase uint8, status uint8) (*types.Transaction, error) {
	return _OpProposalManager.contract.Transact(opts, "voteAndUpdateStatus", proposalID, proposalType, voter, proposers, phase, status)
}

// VoteAndUpdateStatus is a paid mutator transaction binding the contract method 0xe4f5c160.
//
// Solidity: function voteAndUpdateStatus(uint64 proposalID, uint8 proposalType, address voter, address[] proposers, uint8 phase, uint8 status) returns()
func (_OpProposalManager *OpProposalManagerSession) VoteAndUpdateStatus(proposalID uint64, proposalType uint8, voter common.Address, proposers []common.Address, phase uint8, status uint8) (*types.Transaction, error) {
	return _OpProposalManager.Contract.VoteAndUpdateStatus(&_OpProposalManager.TransactOpts, proposalID, proposalType, voter, proposers, phase, status)
}

// VoteAndUpdateStatus is a paid mutator transaction binding the contract method 0xe4f5c160.
//
// Solidity: function voteAndUpdateStatus(uint64 proposalID, uint8 proposalType, address voter, address[] proposers, uint8 phase, uint8 status) returns()
func (_OpProposalManager *OpProposalManagerTransactorSession) VoteAndUpdateStatus(proposalID uint64, proposalType uint8, voter common.Address, proposers []common.Address, phase uint8, status uint8) (*types.Transaction, error) {
	return _OpProposalManager.Contract.VoteAndUpdateStatus(&_OpProposalManager.TransactOpts, proposalID, proposalType, voter, proposers, phase, status)
}

// OpProposalManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OpProposalManager contract.
type OpProposalManagerInitializedIterator struct {
	Event *OpProposalManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OpProposalManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpProposalManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OpProposalManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OpProposalManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpProposalManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpProposalManagerInitialized represents a Initialized event raised by the OpProposalManager contract.
type OpProposalManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OpProposalManager *OpProposalManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*OpProposalManagerInitializedIterator, error) {

	logs, sub, err := _OpProposalManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OpProposalManagerInitializedIterator{contract: _OpProposalManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OpProposalManager *OpProposalManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OpProposalManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _OpProposalManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpProposalManagerInitialized)
				if err := _OpProposalManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OpProposalManager *OpProposalManagerFilterer) ParseInitialized(log types.Log) (*OpProposalManagerInitialized, error) {
	event := new(OpProposalManagerInitialized)
	if err := _OpProposalManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpProposalManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the OpProposalManager contract.
type OpProposalManagerRoleAdminChangedIterator struct {
	Event *OpProposalManagerRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OpProposalManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpProposalManagerRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OpProposalManagerRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OpProposalManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpProposalManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpProposalManagerRoleAdminChanged represents a RoleAdminChanged event raised by the OpProposalManager contract.
type OpProposalManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OpProposalManager *OpProposalManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OpProposalManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _OpProposalManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OpProposalManagerRoleAdminChangedIterator{contract: _OpProposalManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OpProposalManager *OpProposalManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OpProposalManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _OpProposalManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpProposalManagerRoleAdminChanged)
				if err := _OpProposalManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OpProposalManager *OpProposalManagerFilterer) ParseRoleAdminChanged(log types.Log) (*OpProposalManagerRoleAdminChanged, error) {
	event := new(OpProposalManagerRoleAdminChanged)
	if err := _OpProposalManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpProposalManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the OpProposalManager contract.
type OpProposalManagerRoleGrantedIterator struct {
	Event *OpProposalManagerRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OpProposalManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpProposalManagerRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OpProposalManagerRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OpProposalManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpProposalManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpProposalManagerRoleGranted represents a RoleGranted event raised by the OpProposalManager contract.
type OpProposalManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpProposalManager *OpProposalManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OpProposalManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OpProposalManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OpProposalManagerRoleGrantedIterator{contract: _OpProposalManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpProposalManager *OpProposalManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OpProposalManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OpProposalManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpProposalManagerRoleGranted)
				if err := _OpProposalManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpProposalManager *OpProposalManagerFilterer) ParseRoleGranted(log types.Log) (*OpProposalManagerRoleGranted, error) {
	event := new(OpProposalManagerRoleGranted)
	if err := _OpProposalManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpProposalManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the OpProposalManager contract.
type OpProposalManagerRoleRevokedIterator struct {
	Event *OpProposalManagerRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OpProposalManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpProposalManagerRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OpProposalManagerRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OpProposalManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpProposalManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpProposalManagerRoleRevoked represents a RoleRevoked event raised by the OpProposalManager contract.
type OpProposalManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpProposalManager *OpProposalManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OpProposalManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OpProposalManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OpProposalManagerRoleRevokedIterator{contract: _OpProposalManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpProposalManager *OpProposalManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OpProposalManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _OpProposalManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpProposalManagerRoleRevoked)
				if err := _OpProposalManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpProposalManager *OpProposalManagerFilterer) ParseRoleRevoked(log types.Log) (*OpProposalManagerRoleRevoked, error) {
	event := new(OpProposalManagerRoleRevoked)
	if err := _OpProposalManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
