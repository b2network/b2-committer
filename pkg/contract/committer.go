// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// CommitterProposal is an auto generated low-level Go binding around an user-defined struct.
type CommitterProposal struct {
	Id            uint64
	StartIndex    uint64
	EndIndex      uint64
	Status        uint8
	Timeout       *big.Int
	Winner        common.Address
	ProofHash     string
	StateRootHash string
	BtcTxHash     string
	ArweaveTxHash string
}

// CommitterMetaData contains all meta data concerning the Committer contract.
var CommitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposalIsNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proofHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"stateRootHash\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"ProposalSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"}],\"name\":\"ProposalTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"}],\"name\":\"VoteProposal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"addChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"addProposer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allChains\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allProposers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"}],\"name\":\"allSubmitBitcoinTxVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"}],\"name\":\"allSubmitProofVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"name\":\"arweaveTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"name\":\"bitcoinTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"chains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chainsList\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compareStrings\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"getLastProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proofHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"stateRootHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"btcTxHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"arweaveTxHash\",\"type\":\"string\"}],\"internalType\":\"structCommitter.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"}],\"name\":\"isProposalTimeout\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isProposer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"isSupportedChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVotedOnBitcoinTxPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVotedOnSubmitProofPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"lastProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proofHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"stateRootHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"btcTxHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"arweaveTxHash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"}],\"name\":\"proposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proofHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"stateRootHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"btcTxHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"arweaveTxHash\",\"type\":\"string\"}],\"internalType\":\"structCommitter.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"proofHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"stateRootHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"btcTxHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"arweaveTxHash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proposers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposersList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"removeChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"removeProposer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setTimeoutPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"submitBitcoinTxPhaseVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"proofHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"stateRootHash\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"startIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endIndex\",\"type\":\"uint64\"}],\"name\":\"submitProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"submitProofPhaseVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeoutPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"}],\"name\":\"timeoutProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CommitterABI is the input ABI used to generate the binding from.
// Deprecated: Use CommitterMetaData.ABI instead.
var CommitterABI = CommitterMetaData.ABI

// Committer is an auto generated Go binding around an Ethereum contract.
type Committer struct {
	CommitterCaller     // Read-only binding to the contract
	CommitterTransactor // Write-only binding to the contract
	CommitterFilterer   // Log filterer for contract events
}

// CommitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommitterSession struct {
	Contract     *Committer        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommitterCallerSession struct {
	Contract *CommitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CommitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommitterTransactorSession struct {
	Contract     *CommitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CommitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommitterRaw struct {
	Contract *Committer // Generic contract binding to access the raw methods on
}

// CommitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommitterCallerRaw struct {
	Contract *CommitterCaller // Generic read-only contract binding to access the raw methods on
}

// CommitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommitterTransactorRaw struct {
	Contract *CommitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommitter creates a new instance of Committer, bound to a specific deployed contract.
func NewCommitter(address common.Address, backend bind.ContractBackend) (*Committer, error) {
	contract, err := bindCommitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Committer{CommitterCaller: CommitterCaller{contract: contract}, CommitterTransactor: CommitterTransactor{contract: contract}, CommitterFilterer: CommitterFilterer{contract: contract}}, nil
}

// NewCommitterCaller creates a new read-only instance of Committer, bound to a specific deployed contract.
func NewCommitterCaller(address common.Address, caller bind.ContractCaller) (*CommitterCaller, error) {
	contract, err := bindCommitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitterCaller{contract: contract}, nil
}

// NewCommitterTransactor creates a new write-only instance of Committer, bound to a specific deployed contract.
func NewCommitterTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitterTransactor, error) {
	contract, err := bindCommitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitterTransactor{contract: contract}, nil
}

// NewCommitterFilterer creates a new log filterer instance of Committer, bound to a specific deployed contract.
func NewCommitterFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitterFilterer, error) {
	contract, err := bindCommitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitterFilterer{contract: contract}, nil
}

// bindCommitter binds a generic wrapper to an already deployed contract.
func bindCommitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CommitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Committer *CommitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Committer.Contract.CommitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Committer *CommitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Committer.Contract.CommitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Committer *CommitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Committer.Contract.CommitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Committer *CommitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Committer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Committer *CommitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Committer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Committer *CommitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Committer.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Committer *CommitterCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Committer *CommitterSession) ADMINROLE() ([32]byte, error) {
	return _Committer.Contract.ADMINROLE(&_Committer.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Committer *CommitterCallerSession) ADMINROLE() ([32]byte, error) {
	return _Committer.Contract.ADMINROLE(&_Committer.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Committer *CommitterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Committer *CommitterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Committer.Contract.DEFAULTADMINROLE(&_Committer.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Committer *CommitterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Committer.Contract.DEFAULTADMINROLE(&_Committer.CallOpts)
}

// AllChains is a free data retrieval call binding the contract method 0x7f3f94b1.
//
// Solidity: function allChains() view returns(uint64[])
func (_Committer *CommitterCaller) AllChains(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "allChains")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// AllChains is a free data retrieval call binding the contract method 0x7f3f94b1.
//
// Solidity: function allChains() view returns(uint64[])
func (_Committer *CommitterSession) AllChains() ([]uint64, error) {
	return _Committer.Contract.AllChains(&_Committer.CallOpts)
}

// AllChains is a free data retrieval call binding the contract method 0x7f3f94b1.
//
// Solidity: function allChains() view returns(uint64[])
func (_Committer *CommitterCallerSession) AllChains() ([]uint64, error) {
	return _Committer.Contract.AllChains(&_Committer.CallOpts)
}

// AllProposers is a free data retrieval call binding the contract method 0x239e1e3d.
//
// Solidity: function allProposers() view returns(address[])
func (_Committer *CommitterCaller) AllProposers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "allProposers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllProposers is a free data retrieval call binding the contract method 0x239e1e3d.
//
// Solidity: function allProposers() view returns(address[])
func (_Committer *CommitterSession) AllProposers() ([]common.Address, error) {
	return _Committer.Contract.AllProposers(&_Committer.CallOpts)
}

// AllProposers is a free data retrieval call binding the contract method 0x239e1e3d.
//
// Solidity: function allProposers() view returns(address[])
func (_Committer *CommitterCallerSession) AllProposers() ([]common.Address, error) {
	return _Committer.Contract.AllProposers(&_Committer.CallOpts)
}

// AllSubmitBitcoinTxVotes is a free data retrieval call binding the contract method 0xbedc7561.
//
// Solidity: function allSubmitBitcoinTxVotes(uint64 chainID, uint64 proposalID) view returns(address[])
func (_Committer *CommitterCaller) AllSubmitBitcoinTxVotes(opts *bind.CallOpts, chainID uint64, proposalID uint64) ([]common.Address, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "allSubmitBitcoinTxVotes", chainID, proposalID)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllSubmitBitcoinTxVotes is a free data retrieval call binding the contract method 0xbedc7561.
//
// Solidity: function allSubmitBitcoinTxVotes(uint64 chainID, uint64 proposalID) view returns(address[])
func (_Committer *CommitterSession) AllSubmitBitcoinTxVotes(chainID uint64, proposalID uint64) ([]common.Address, error) {
	return _Committer.Contract.AllSubmitBitcoinTxVotes(&_Committer.CallOpts, chainID, proposalID)
}

// AllSubmitBitcoinTxVotes is a free data retrieval call binding the contract method 0xbedc7561.
//
// Solidity: function allSubmitBitcoinTxVotes(uint64 chainID, uint64 proposalID) view returns(address[])
func (_Committer *CommitterCallerSession) AllSubmitBitcoinTxVotes(chainID uint64, proposalID uint64) ([]common.Address, error) {
	return _Committer.Contract.AllSubmitBitcoinTxVotes(&_Committer.CallOpts, chainID, proposalID)
}

// AllSubmitProofVotes is a free data retrieval call binding the contract method 0xea7d929a.
//
// Solidity: function allSubmitProofVotes(uint64 chainID, uint64 proposalID) view returns(address[])
func (_Committer *CommitterCaller) AllSubmitProofVotes(opts *bind.CallOpts, chainID uint64, proposalID uint64) ([]common.Address, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "allSubmitProofVotes", chainID, proposalID)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllSubmitProofVotes is a free data retrieval call binding the contract method 0xea7d929a.
//
// Solidity: function allSubmitProofVotes(uint64 chainID, uint64 proposalID) view returns(address[])
func (_Committer *CommitterSession) AllSubmitProofVotes(chainID uint64, proposalID uint64) ([]common.Address, error) {
	return _Committer.Contract.AllSubmitProofVotes(&_Committer.CallOpts, chainID, proposalID)
}

// AllSubmitProofVotes is a free data retrieval call binding the contract method 0xea7d929a.
//
// Solidity: function allSubmitProofVotes(uint64 chainID, uint64 proposalID) view returns(address[])
func (_Committer *CommitterCallerSession) AllSubmitProofVotes(chainID uint64, proposalID uint64) ([]common.Address, error) {
	return _Committer.Contract.AllSubmitProofVotes(&_Committer.CallOpts, chainID, proposalID)
}

// Chains is a free data retrieval call binding the contract method 0xada8bcdc.
//
// Solidity: function chains(uint64 ) view returns(bool)
func (_Committer *CommitterCaller) Chains(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "chains", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Chains is a free data retrieval call binding the contract method 0xada8bcdc.
//
// Solidity: function chains(uint64 ) view returns(bool)
func (_Committer *CommitterSession) Chains(arg0 uint64) (bool, error) {
	return _Committer.Contract.Chains(&_Committer.CallOpts, arg0)
}

// Chains is a free data retrieval call binding the contract method 0xada8bcdc.
//
// Solidity: function chains(uint64 ) view returns(bool)
func (_Committer *CommitterCallerSession) Chains(arg0 uint64) (bool, error) {
	return _Committer.Contract.Chains(&_Committer.CallOpts, arg0)
}

// ChainsList is a free data retrieval call binding the contract method 0x2bc3db0a.
//
// Solidity: function chainsList(uint256 ) view returns(uint64)
func (_Committer *CommitterCaller) ChainsList(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "chainsList", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainsList is a free data retrieval call binding the contract method 0x2bc3db0a.
//
// Solidity: function chainsList(uint256 ) view returns(uint64)
func (_Committer *CommitterSession) ChainsList(arg0 *big.Int) (uint64, error) {
	return _Committer.Contract.ChainsList(&_Committer.CallOpts, arg0)
}

// ChainsList is a free data retrieval call binding the contract method 0x2bc3db0a.
//
// Solidity: function chainsList(uint256 ) view returns(uint64)
func (_Committer *CommitterCallerSession) ChainsList(arg0 *big.Int) (uint64, error) {
	return _Committer.Contract.ChainsList(&_Committer.CallOpts, arg0)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) pure returns(bool)
func (_Committer *CommitterCaller) CompareStrings(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "compareStrings", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) pure returns(bool)
func (_Committer *CommitterSession) CompareStrings(a string, b string) (bool, error) {
	return _Committer.Contract.CompareStrings(&_Committer.CallOpts, a, b)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string a, string b) pure returns(bool)
func (_Committer *CommitterCallerSession) CompareStrings(a string, b string) (bool, error) {
	return _Committer.Contract.CompareStrings(&_Committer.CallOpts, a, b)
}

// GetLastProposal is a free data retrieval call binding the contract method 0xa2c732fd.
//
// Solidity: function getLastProposal(uint64 chainID) view returns((uint64,uint64,uint64,uint8,uint256,address,string,string,string,string))
func (_Committer *CommitterCaller) GetLastProposal(opts *bind.CallOpts, chainID uint64) (CommitterProposal, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "getLastProposal", chainID)

	if err != nil {
		return *new(CommitterProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitterProposal)).(*CommitterProposal)

	return out0, err

}

// GetLastProposal is a free data retrieval call binding the contract method 0xa2c732fd.
//
// Solidity: function getLastProposal(uint64 chainID) view returns((uint64,uint64,uint64,uint8,uint256,address,string,string,string,string))
func (_Committer *CommitterSession) GetLastProposal(chainID uint64) (CommitterProposal, error) {
	return _Committer.Contract.GetLastProposal(&_Committer.CallOpts, chainID)
}

// GetLastProposal is a free data retrieval call binding the contract method 0xa2c732fd.
//
// Solidity: function getLastProposal(uint64 chainID) view returns((uint64,uint64,uint64,uint8,uint256,address,string,string,string,string))
func (_Committer *CommitterCallerSession) GetLastProposal(chainID uint64) (CommitterProposal, error) {
	return _Committer.Contract.GetLastProposal(&_Committer.CallOpts, chainID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Committer *CommitterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Committer *CommitterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Committer.Contract.GetRoleAdmin(&_Committer.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Committer *CommitterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Committer.Contract.GetRoleAdmin(&_Committer.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Committer *CommitterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Committer *CommitterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Committer.Contract.HasRole(&_Committer.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Committer *CommitterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Committer.Contract.HasRole(&_Committer.CallOpts, role, account)
}

// IsProposalTimeout is a free data retrieval call binding the contract method 0x1b77a930.
//
// Solidity: function isProposalTimeout(uint64 chainID, uint64 proposalID) view returns(bool)
func (_Committer *CommitterCaller) IsProposalTimeout(opts *bind.CallOpts, chainID uint64, proposalID uint64) (bool, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "isProposalTimeout", chainID, proposalID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProposalTimeout is a free data retrieval call binding the contract method 0x1b77a930.
//
// Solidity: function isProposalTimeout(uint64 chainID, uint64 proposalID) view returns(bool)
func (_Committer *CommitterSession) IsProposalTimeout(chainID uint64, proposalID uint64) (bool, error) {
	return _Committer.Contract.IsProposalTimeout(&_Committer.CallOpts, chainID, proposalID)
}

// IsProposalTimeout is a free data retrieval call binding the contract method 0x1b77a930.
//
// Solidity: function isProposalTimeout(uint64 chainID, uint64 proposalID) view returns(bool)
func (_Committer *CommitterCallerSession) IsProposalTimeout(chainID uint64, proposalID uint64) (bool, error) {
	return _Committer.Contract.IsProposalTimeout(&_Committer.CallOpts, chainID, proposalID)
}

// IsProposer is a free data retrieval call binding the contract method 0x74ec29a0.
//
// Solidity: function isProposer(address sender) view returns(bool)
func (_Committer *CommitterCaller) IsProposer(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "isProposer", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProposer is a free data retrieval call binding the contract method 0x74ec29a0.
//
// Solidity: function isProposer(address sender) view returns(bool)
func (_Committer *CommitterSession) IsProposer(sender common.Address) (bool, error) {
	return _Committer.Contract.IsProposer(&_Committer.CallOpts, sender)
}

// IsProposer is a free data retrieval call binding the contract method 0x74ec29a0.
//
// Solidity: function isProposer(address sender) view returns(bool)
func (_Committer *CommitterCallerSession) IsProposer(sender common.Address) (bool, error) {
	return _Committer.Contract.IsProposer(&_Committer.CallOpts, sender)
}

// IsSupportedChain is a free data retrieval call binding the contract method 0x8926f54f.
//
// Solidity: function isSupportedChain(uint64 chainID) view returns(bool)
func (_Committer *CommitterCaller) IsSupportedChain(opts *bind.CallOpts, chainID uint64) (bool, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "isSupportedChain", chainID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedChain is a free data retrieval call binding the contract method 0x8926f54f.
//
// Solidity: function isSupportedChain(uint64 chainID) view returns(bool)
func (_Committer *CommitterSession) IsSupportedChain(chainID uint64) (bool, error) {
	return _Committer.Contract.IsSupportedChain(&_Committer.CallOpts, chainID)
}

// IsSupportedChain is a free data retrieval call binding the contract method 0x8926f54f.
//
// Solidity: function isSupportedChain(uint64 chainID) view returns(bool)
func (_Committer *CommitterCallerSession) IsSupportedChain(chainID uint64) (bool, error) {
	return _Committer.Contract.IsSupportedChain(&_Committer.CallOpts, chainID)
}

// IsVotedOnBitcoinTxPhase is a free data retrieval call binding the contract method 0x87c8eabd.
//
// Solidity: function isVotedOnBitcoinTxPhase(uint64 chainID, uint64 proposalID, address voter) view returns(bool)
func (_Committer *CommitterCaller) IsVotedOnBitcoinTxPhase(opts *bind.CallOpts, chainID uint64, proposalID uint64, voter common.Address) (bool, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "isVotedOnBitcoinTxPhase", chainID, proposalID, voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVotedOnBitcoinTxPhase is a free data retrieval call binding the contract method 0x87c8eabd.
//
// Solidity: function isVotedOnBitcoinTxPhase(uint64 chainID, uint64 proposalID, address voter) view returns(bool)
func (_Committer *CommitterSession) IsVotedOnBitcoinTxPhase(chainID uint64, proposalID uint64, voter common.Address) (bool, error) {
	return _Committer.Contract.IsVotedOnBitcoinTxPhase(&_Committer.CallOpts, chainID, proposalID, voter)
}

// IsVotedOnBitcoinTxPhase is a free data retrieval call binding the contract method 0x87c8eabd.
//
// Solidity: function isVotedOnBitcoinTxPhase(uint64 chainID, uint64 proposalID, address voter) view returns(bool)
func (_Committer *CommitterCallerSession) IsVotedOnBitcoinTxPhase(chainID uint64, proposalID uint64, voter common.Address) (bool, error) {
	return _Committer.Contract.IsVotedOnBitcoinTxPhase(&_Committer.CallOpts, chainID, proposalID, voter)
}

// IsVotedOnSubmitProofPhase is a free data retrieval call binding the contract method 0x975a4ce7.
//
// Solidity: function isVotedOnSubmitProofPhase(uint64 chainID, uint64 proposalID, address voter) view returns(bool)
func (_Committer *CommitterCaller) IsVotedOnSubmitProofPhase(opts *bind.CallOpts, chainID uint64, proposalID uint64, voter common.Address) (bool, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "isVotedOnSubmitProofPhase", chainID, proposalID, voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVotedOnSubmitProofPhase is a free data retrieval call binding the contract method 0x975a4ce7.
//
// Solidity: function isVotedOnSubmitProofPhase(uint64 chainID, uint64 proposalID, address voter) view returns(bool)
func (_Committer *CommitterSession) IsVotedOnSubmitProofPhase(chainID uint64, proposalID uint64, voter common.Address) (bool, error) {
	return _Committer.Contract.IsVotedOnSubmitProofPhase(&_Committer.CallOpts, chainID, proposalID, voter)
}

// IsVotedOnSubmitProofPhase is a free data retrieval call binding the contract method 0x975a4ce7.
//
// Solidity: function isVotedOnSubmitProofPhase(uint64 chainID, uint64 proposalID, address voter) view returns(bool)
func (_Committer *CommitterCallerSession) IsVotedOnSubmitProofPhase(chainID uint64, proposalID uint64, voter common.Address) (bool, error) {
	return _Committer.Contract.IsVotedOnSubmitProofPhase(&_Committer.CallOpts, chainID, proposalID, voter)
}

// LastProposal is a free data retrieval call binding the contract method 0x36a31476.
//
// Solidity: function lastProposal(uint64 ) view returns(uint64 id, uint64 startIndex, uint64 endIndex, uint8 status, uint256 timeout, address winner, string proofHash, string stateRootHash, string btcTxHash, string arweaveTxHash)
func (_Committer *CommitterCaller) LastProposal(opts *bind.CallOpts, arg0 uint64) (struct {
	Id            uint64
	StartIndex    uint64
	EndIndex      uint64
	Status        uint8
	Timeout       *big.Int
	Winner        common.Address
	ProofHash     string
	StateRootHash string
	BtcTxHash     string
	ArweaveTxHash string
}, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "lastProposal", arg0)

	outstruct := new(struct {
		Id            uint64
		StartIndex    uint64
		EndIndex      uint64
		Status        uint8
		Timeout       *big.Int
		Winner        common.Address
		ProofHash     string
		StateRootHash string
		BtcTxHash     string
		ArweaveTxHash string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.StartIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.EndIndex = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Timeout = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Winner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.ProofHash = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.StateRootHash = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.BtcTxHash = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.ArweaveTxHash = *abi.ConvertType(out[9], new(string)).(*string)

	return *outstruct, err

}

// LastProposal is a free data retrieval call binding the contract method 0x36a31476.
//
// Solidity: function lastProposal(uint64 ) view returns(uint64 id, uint64 startIndex, uint64 endIndex, uint8 status, uint256 timeout, address winner, string proofHash, string stateRootHash, string btcTxHash, string arweaveTxHash)
func (_Committer *CommitterSession) LastProposal(arg0 uint64) (struct {
	Id            uint64
	StartIndex    uint64
	EndIndex      uint64
	Status        uint8
	Timeout       *big.Int
	Winner        common.Address
	ProofHash     string
	StateRootHash string
	BtcTxHash     string
	ArweaveTxHash string
}, error) {
	return _Committer.Contract.LastProposal(&_Committer.CallOpts, arg0)
}

// LastProposal is a free data retrieval call binding the contract method 0x36a31476.
//
// Solidity: function lastProposal(uint64 ) view returns(uint64 id, uint64 startIndex, uint64 endIndex, uint8 status, uint256 timeout, address winner, string proofHash, string stateRootHash, string btcTxHash, string arweaveTxHash)
func (_Committer *CommitterCallerSession) LastProposal(arg0 uint64) (struct {
	Id            uint64
	StartIndex    uint64
	EndIndex      uint64
	Status        uint8
	Timeout       *big.Int
	Winner        common.Address
	ProofHash     string
	StateRootHash string
	BtcTxHash     string
	ArweaveTxHash string
}, error) {
	return _Committer.Contract.LastProposal(&_Committer.CallOpts, arg0)
}

// Proposal is a free data retrieval call binding the contract method 0x0f93fa52.
//
// Solidity: function proposal(uint64 chainID, uint64 proposalID) view returns((uint64,uint64,uint64,uint8,uint256,address,string,string,string,string))
func (_Committer *CommitterCaller) Proposal(opts *bind.CallOpts, chainID uint64, proposalID uint64) (CommitterProposal, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "proposal", chainID, proposalID)

	if err != nil {
		return *new(CommitterProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitterProposal)).(*CommitterProposal)

	return out0, err

}

// Proposal is a free data retrieval call binding the contract method 0x0f93fa52.
//
// Solidity: function proposal(uint64 chainID, uint64 proposalID) view returns((uint64,uint64,uint64,uint8,uint256,address,string,string,string,string))
func (_Committer *CommitterSession) Proposal(chainID uint64, proposalID uint64) (CommitterProposal, error) {
	return _Committer.Contract.Proposal(&_Committer.CallOpts, chainID, proposalID)
}

// Proposal is a free data retrieval call binding the contract method 0x0f93fa52.
//
// Solidity: function proposal(uint64 chainID, uint64 proposalID) view returns((uint64,uint64,uint64,uint8,uint256,address,string,string,string,string))
func (_Committer *CommitterCallerSession) Proposal(chainID uint64, proposalID uint64) (CommitterProposal, error) {
	return _Committer.Contract.Proposal(&_Committer.CallOpts, chainID, proposalID)
}

// Proposals is a free data retrieval call binding the contract method 0x7176d85d.
//
// Solidity: function proposals(uint64 , uint64 ) view returns(uint64 id, uint64 startIndex, uint64 endIndex, uint8 status, uint256 timeout, address winner, string proofHash, string stateRootHash, string btcTxHash, string arweaveTxHash)
func (_Committer *CommitterCaller) Proposals(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (struct {
	Id            uint64
	StartIndex    uint64
	EndIndex      uint64
	Status        uint8
	Timeout       *big.Int
	Winner        common.Address
	ProofHash     string
	StateRootHash string
	BtcTxHash     string
	ArweaveTxHash string
}, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "proposals", arg0, arg1)

	outstruct := new(struct {
		Id            uint64
		StartIndex    uint64
		EndIndex      uint64
		Status        uint8
		Timeout       *big.Int
		Winner        common.Address
		ProofHash     string
		StateRootHash string
		BtcTxHash     string
		ArweaveTxHash string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.StartIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.EndIndex = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Timeout = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Winner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.ProofHash = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.StateRootHash = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.BtcTxHash = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.ArweaveTxHash = *abi.ConvertType(out[9], new(string)).(*string)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x7176d85d.
//
// Solidity: function proposals(uint64 , uint64 ) view returns(uint64 id, uint64 startIndex, uint64 endIndex, uint8 status, uint256 timeout, address winner, string proofHash, string stateRootHash, string btcTxHash, string arweaveTxHash)
func (_Committer *CommitterSession) Proposals(arg0 uint64, arg1 uint64) (struct {
	Id            uint64
	StartIndex    uint64
	EndIndex      uint64
	Status        uint8
	Timeout       *big.Int
	Winner        common.Address
	ProofHash     string
	StateRootHash string
	BtcTxHash     string
	ArweaveTxHash string
}, error) {
	return _Committer.Contract.Proposals(&_Committer.CallOpts, arg0, arg1)
}

// Proposals is a free data retrieval call binding the contract method 0x7176d85d.
//
// Solidity: function proposals(uint64 , uint64 ) view returns(uint64 id, uint64 startIndex, uint64 endIndex, uint8 status, uint256 timeout, address winner, string proofHash, string stateRootHash, string btcTxHash, string arweaveTxHash)
func (_Committer *CommitterCallerSession) Proposals(arg0 uint64, arg1 uint64) (struct {
	Id            uint64
	StartIndex    uint64
	EndIndex      uint64
	Status        uint8
	Timeout       *big.Int
	Winner        common.Address
	ProofHash     string
	StateRootHash string
	BtcTxHash     string
	ArweaveTxHash string
}, error) {
	return _Committer.Contract.Proposals(&_Committer.CallOpts, arg0, arg1)
}

// Proposers is a free data retrieval call binding the contract method 0x18177497.
//
// Solidity: function proposers(address ) view returns(bool)
func (_Committer *CommitterCaller) Proposers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "proposers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Proposers is a free data retrieval call binding the contract method 0x18177497.
//
// Solidity: function proposers(address ) view returns(bool)
func (_Committer *CommitterSession) Proposers(arg0 common.Address) (bool, error) {
	return _Committer.Contract.Proposers(&_Committer.CallOpts, arg0)
}

// Proposers is a free data retrieval call binding the contract method 0x18177497.
//
// Solidity: function proposers(address ) view returns(bool)
func (_Committer *CommitterCallerSession) Proposers(arg0 common.Address) (bool, error) {
	return _Committer.Contract.Proposers(&_Committer.CallOpts, arg0)
}

// ProposersList is a free data retrieval call binding the contract method 0x5ad97dd8.
//
// Solidity: function proposersList(uint256 ) view returns(address)
func (_Committer *CommitterCaller) ProposersList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "proposersList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposersList is a free data retrieval call binding the contract method 0x5ad97dd8.
//
// Solidity: function proposersList(uint256 ) view returns(address)
func (_Committer *CommitterSession) ProposersList(arg0 *big.Int) (common.Address, error) {
	return _Committer.Contract.ProposersList(&_Committer.CallOpts, arg0)
}

// ProposersList is a free data retrieval call binding the contract method 0x5ad97dd8.
//
// Solidity: function proposersList(uint256 ) view returns(address)
func (_Committer *CommitterCallerSession) ProposersList(arg0 *big.Int) (common.Address, error) {
	return _Committer.Contract.ProposersList(&_Committer.CallOpts, arg0)
}

// SubmitBitcoinTxPhaseVotes is a free data retrieval call binding the contract method 0x9bfa86d3.
//
// Solidity: function submitBitcoinTxPhaseVotes(uint64 , uint64 , uint256 ) view returns(address)
func (_Committer *CommitterCaller) SubmitBitcoinTxPhaseVotes(opts *bind.CallOpts, arg0 uint64, arg1 uint64, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "submitBitcoinTxPhaseVotes", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SubmitBitcoinTxPhaseVotes is a free data retrieval call binding the contract method 0x9bfa86d3.
//
// Solidity: function submitBitcoinTxPhaseVotes(uint64 , uint64 , uint256 ) view returns(address)
func (_Committer *CommitterSession) SubmitBitcoinTxPhaseVotes(arg0 uint64, arg1 uint64, arg2 *big.Int) (common.Address, error) {
	return _Committer.Contract.SubmitBitcoinTxPhaseVotes(&_Committer.CallOpts, arg0, arg1, arg2)
}

// SubmitBitcoinTxPhaseVotes is a free data retrieval call binding the contract method 0x9bfa86d3.
//
// Solidity: function submitBitcoinTxPhaseVotes(uint64 , uint64 , uint256 ) view returns(address)
func (_Committer *CommitterCallerSession) SubmitBitcoinTxPhaseVotes(arg0 uint64, arg1 uint64, arg2 *big.Int) (common.Address, error) {
	return _Committer.Contract.SubmitBitcoinTxPhaseVotes(&_Committer.CallOpts, arg0, arg1, arg2)
}

// SubmitProofPhaseVotes is a free data retrieval call binding the contract method 0x24ff16b3.
//
// Solidity: function submitProofPhaseVotes(uint64 , uint64 , uint256 ) view returns(address)
func (_Committer *CommitterCaller) SubmitProofPhaseVotes(opts *bind.CallOpts, arg0 uint64, arg1 uint64, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "submitProofPhaseVotes", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SubmitProofPhaseVotes is a free data retrieval call binding the contract method 0x24ff16b3.
//
// Solidity: function submitProofPhaseVotes(uint64 , uint64 , uint256 ) view returns(address)
func (_Committer *CommitterSession) SubmitProofPhaseVotes(arg0 uint64, arg1 uint64, arg2 *big.Int) (common.Address, error) {
	return _Committer.Contract.SubmitProofPhaseVotes(&_Committer.CallOpts, arg0, arg1, arg2)
}

// SubmitProofPhaseVotes is a free data retrieval call binding the contract method 0x24ff16b3.
//
// Solidity: function submitProofPhaseVotes(uint64 , uint64 , uint256 ) view returns(address)
func (_Committer *CommitterCallerSession) SubmitProofPhaseVotes(arg0 uint64, arg1 uint64, arg2 *big.Int) (common.Address, error) {
	return _Committer.Contract.SubmitProofPhaseVotes(&_Committer.CallOpts, arg0, arg1, arg2)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Committer *CommitterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Committer *CommitterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Committer.Contract.SupportsInterface(&_Committer.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Committer *CommitterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Committer.Contract.SupportsInterface(&_Committer.CallOpts, interfaceId)
}

// TimeoutPeriod is a free data retrieval call binding the contract method 0x6cea6b82.
//
// Solidity: function timeoutPeriod() view returns(uint256)
func (_Committer *CommitterCaller) TimeoutPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Committer.contract.Call(opts, &out, "timeoutPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeoutPeriod is a free data retrieval call binding the contract method 0x6cea6b82.
//
// Solidity: function timeoutPeriod() view returns(uint256)
func (_Committer *CommitterSession) TimeoutPeriod() (*big.Int, error) {
	return _Committer.Contract.TimeoutPeriod(&_Committer.CallOpts)
}

// TimeoutPeriod is a free data retrieval call binding the contract method 0x6cea6b82.
//
// Solidity: function timeoutPeriod() view returns(uint256)
func (_Committer *CommitterCallerSession) TimeoutPeriod() (*big.Int, error) {
	return _Committer.Contract.TimeoutPeriod(&_Committer.CallOpts)
}

// AddChain is a paid mutator transaction binding the contract method 0xe2947e54.
//
// Solidity: function addChain(uint64 chainID) returns()
func (_Committer *CommitterTransactor) AddChain(opts *bind.TransactOpts, chainID uint64) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "addChain", chainID)
}

// AddChain is a paid mutator transaction binding the contract method 0xe2947e54.
//
// Solidity: function addChain(uint64 chainID) returns()
func (_Committer *CommitterSession) AddChain(chainID uint64) (*types.Transaction, error) {
	return _Committer.Contract.AddChain(&_Committer.TransactOpts, chainID)
}

// AddChain is a paid mutator transaction binding the contract method 0xe2947e54.
//
// Solidity: function addChain(uint64 chainID) returns()
func (_Committer *CommitterTransactorSession) AddChain(chainID uint64) (*types.Transaction, error) {
	return _Committer.Contract.AddChain(&_Committer.TransactOpts, chainID)
}

// AddProposer is a paid mutator transaction binding the contract method 0xb03cd418.
//
// Solidity: function addProposer(address proposer) returns()
func (_Committer *CommitterTransactor) AddProposer(opts *bind.TransactOpts, proposer common.Address) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "addProposer", proposer)
}

// AddProposer is a paid mutator transaction binding the contract method 0xb03cd418.
//
// Solidity: function addProposer(address proposer) returns()
func (_Committer *CommitterSession) AddProposer(proposer common.Address) (*types.Transaction, error) {
	return _Committer.Contract.AddProposer(&_Committer.TransactOpts, proposer)
}

// AddProposer is a paid mutator transaction binding the contract method 0xb03cd418.
//
// Solidity: function addProposer(address proposer) returns()
func (_Committer *CommitterTransactorSession) AddProposer(proposer common.Address) (*types.Transaction, error) {
	return _Committer.Contract.AddProposer(&_Committer.TransactOpts, proposer)
}

// ArweaveTx is a paid mutator transaction binding the contract method 0xe0a9daf6.
//
// Solidity: function arweaveTx(uint64 chainID, uint64 proposalID, string txHash) returns()
func (_Committer *CommitterTransactor) ArweaveTx(opts *bind.TransactOpts, chainID uint64, proposalID uint64, txHash string) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "arweaveTx", chainID, proposalID, txHash)
}

// ArweaveTx is a paid mutator transaction binding the contract method 0xe0a9daf6.
//
// Solidity: function arweaveTx(uint64 chainID, uint64 proposalID, string txHash) returns()
func (_Committer *CommitterSession) ArweaveTx(chainID uint64, proposalID uint64, txHash string) (*types.Transaction, error) {
	return _Committer.Contract.ArweaveTx(&_Committer.TransactOpts, chainID, proposalID, txHash)
}

// ArweaveTx is a paid mutator transaction binding the contract method 0xe0a9daf6.
//
// Solidity: function arweaveTx(uint64 chainID, uint64 proposalID, string txHash) returns()
func (_Committer *CommitterTransactorSession) ArweaveTx(chainID uint64, proposalID uint64, txHash string) (*types.Transaction, error) {
	return _Committer.Contract.ArweaveTx(&_Committer.TransactOpts, chainID, proposalID, txHash)
}

// BitcoinTx is a paid mutator transaction binding the contract method 0x54cab80e.
//
// Solidity: function bitcoinTx(uint64 chainID, uint64 proposalID, string txHash) returns()
func (_Committer *CommitterTransactor) BitcoinTx(opts *bind.TransactOpts, chainID uint64, proposalID uint64, txHash string) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "bitcoinTx", chainID, proposalID, txHash)
}

// BitcoinTx is a paid mutator transaction binding the contract method 0x54cab80e.
//
// Solidity: function bitcoinTx(uint64 chainID, uint64 proposalID, string txHash) returns()
func (_Committer *CommitterSession) BitcoinTx(chainID uint64, proposalID uint64, txHash string) (*types.Transaction, error) {
	return _Committer.Contract.BitcoinTx(&_Committer.TransactOpts, chainID, proposalID, txHash)
}

// BitcoinTx is a paid mutator transaction binding the contract method 0x54cab80e.
//
// Solidity: function bitcoinTx(uint64 chainID, uint64 proposalID, string txHash) returns()
func (_Committer *CommitterTransactorSession) BitcoinTx(chainID uint64, proposalID uint64, txHash string) (*types.Transaction, error) {
	return _Committer.Contract.BitcoinTx(&_Committer.TransactOpts, chainID, proposalID, txHash)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Committer *CommitterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Committer *CommitterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Committer.Contract.GrantRole(&_Committer.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Committer *CommitterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Committer.Contract.GrantRole(&_Committer.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Committer *CommitterTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Committer *CommitterSession) Initialize() (*types.Transaction, error) {
	return _Committer.Contract.Initialize(&_Committer.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Committer *CommitterTransactorSession) Initialize() (*types.Transaction, error) {
	return _Committer.Contract.Initialize(&_Committer.TransactOpts)
}

// RemoveChain is a paid mutator transaction binding the contract method 0x39aa1335.
//
// Solidity: function removeChain(uint64 chainID) returns()
func (_Committer *CommitterTransactor) RemoveChain(opts *bind.TransactOpts, chainID uint64) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "removeChain", chainID)
}

// RemoveChain is a paid mutator transaction binding the contract method 0x39aa1335.
//
// Solidity: function removeChain(uint64 chainID) returns()
func (_Committer *CommitterSession) RemoveChain(chainID uint64) (*types.Transaction, error) {
	return _Committer.Contract.RemoveChain(&_Committer.TransactOpts, chainID)
}

// RemoveChain is a paid mutator transaction binding the contract method 0x39aa1335.
//
// Solidity: function removeChain(uint64 chainID) returns()
func (_Committer *CommitterTransactorSession) RemoveChain(chainID uint64) (*types.Transaction, error) {
	return _Committer.Contract.RemoveChain(&_Committer.TransactOpts, chainID)
}

// RemoveProposer is a paid mutator transaction binding the contract method 0x09d632d3.
//
// Solidity: function removeProposer(address proposer) returns()
func (_Committer *CommitterTransactor) RemoveProposer(opts *bind.TransactOpts, proposer common.Address) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "removeProposer", proposer)
}

// RemoveProposer is a paid mutator transaction binding the contract method 0x09d632d3.
//
// Solidity: function removeProposer(address proposer) returns()
func (_Committer *CommitterSession) RemoveProposer(proposer common.Address) (*types.Transaction, error) {
	return _Committer.Contract.RemoveProposer(&_Committer.TransactOpts, proposer)
}

// RemoveProposer is a paid mutator transaction binding the contract method 0x09d632d3.
//
// Solidity: function removeProposer(address proposer) returns()
func (_Committer *CommitterTransactorSession) RemoveProposer(proposer common.Address) (*types.Transaction, error) {
	return _Committer.Contract.RemoveProposer(&_Committer.TransactOpts, proposer)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Committer *CommitterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Committer *CommitterSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Committer.Contract.RenounceRole(&_Committer.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Committer *CommitterTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Committer.Contract.RenounceRole(&_Committer.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Committer *CommitterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Committer *CommitterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Committer.Contract.RevokeRole(&_Committer.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Committer *CommitterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Committer.Contract.RevokeRole(&_Committer.TransactOpts, role, account)
}

// SetTimeoutPeriod is a paid mutator transaction binding the contract method 0x227c6dfa.
//
// Solidity: function setTimeoutPeriod(uint256 period) returns()
func (_Committer *CommitterTransactor) SetTimeoutPeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "setTimeoutPeriod", period)
}

// SetTimeoutPeriod is a paid mutator transaction binding the contract method 0x227c6dfa.
//
// Solidity: function setTimeoutPeriod(uint256 period) returns()
func (_Committer *CommitterSession) SetTimeoutPeriod(period *big.Int) (*types.Transaction, error) {
	return _Committer.Contract.SetTimeoutPeriod(&_Committer.TransactOpts, period)
}

// SetTimeoutPeriod is a paid mutator transaction binding the contract method 0x227c6dfa.
//
// Solidity: function setTimeoutPeriod(uint256 period) returns()
func (_Committer *CommitterTransactorSession) SetTimeoutPeriod(period *big.Int) (*types.Transaction, error) {
	return _Committer.Contract.SetTimeoutPeriod(&_Committer.TransactOpts, period)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x1f981c36.
//
// Solidity: function submitProof(uint64 chainID, uint64 proposalID, string proofHash, string stateRootHash, uint64 startIndex, uint64 endIndex) returns()
func (_Committer *CommitterTransactor) SubmitProof(opts *bind.TransactOpts, chainID uint64, proposalID uint64, proofHash string, stateRootHash string, startIndex uint64, endIndex uint64) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "submitProof", chainID, proposalID, proofHash, stateRootHash, startIndex, endIndex)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x1f981c36.
//
// Solidity: function submitProof(uint64 chainID, uint64 proposalID, string proofHash, string stateRootHash, uint64 startIndex, uint64 endIndex) returns()
func (_Committer *CommitterSession) SubmitProof(chainID uint64, proposalID uint64, proofHash string, stateRootHash string, startIndex uint64, endIndex uint64) (*types.Transaction, error) {
	return _Committer.Contract.SubmitProof(&_Committer.TransactOpts, chainID, proposalID, proofHash, stateRootHash, startIndex, endIndex)
}

// SubmitProof is a paid mutator transaction binding the contract method 0x1f981c36.
//
// Solidity: function submitProof(uint64 chainID, uint64 proposalID, string proofHash, string stateRootHash, uint64 startIndex, uint64 endIndex) returns()
func (_Committer *CommitterTransactorSession) SubmitProof(chainID uint64, proposalID uint64, proofHash string, stateRootHash string, startIndex uint64, endIndex uint64) (*types.Transaction, error) {
	return _Committer.Contract.SubmitProof(&_Committer.TransactOpts, chainID, proposalID, proofHash, stateRootHash, startIndex, endIndex)
}

// TimeoutProposal is a paid mutator transaction binding the contract method 0x04956a75.
//
// Solidity: function timeoutProposal(uint64 chainID, uint64 proposalID) returns()
func (_Committer *CommitterTransactor) TimeoutProposal(opts *bind.TransactOpts, chainID uint64, proposalID uint64) (*types.Transaction, error) {
	return _Committer.contract.Transact(opts, "timeoutProposal", chainID, proposalID)
}

// TimeoutProposal is a paid mutator transaction binding the contract method 0x04956a75.
//
// Solidity: function timeoutProposal(uint64 chainID, uint64 proposalID) returns()
func (_Committer *CommitterSession) TimeoutProposal(chainID uint64, proposalID uint64) (*types.Transaction, error) {
	return _Committer.Contract.TimeoutProposal(&_Committer.TransactOpts, chainID, proposalID)
}

// TimeoutProposal is a paid mutator transaction binding the contract method 0x04956a75.
//
// Solidity: function timeoutProposal(uint64 chainID, uint64 proposalID) returns()
func (_Committer *CommitterTransactorSession) TimeoutProposal(chainID uint64, proposalID uint64) (*types.Transaction, error) {
	return _Committer.Contract.TimeoutProposal(&_Committer.TransactOpts, chainID, proposalID)
}

// CommitterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Committer contract.
type CommitterInitializedIterator struct {
	Event *CommitterInitialized // Event containing the contract specifics and raw log

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
func (it *CommitterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitterInitialized)
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
		it.Event = new(CommitterInitialized)
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
func (it *CommitterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitterInitialized represents a Initialized event raised by the Committer contract.
type CommitterInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Committer *CommitterFilterer) FilterInitialized(opts *bind.FilterOpts) (*CommitterInitializedIterator, error) {

	logs, sub, err := _Committer.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CommitterInitializedIterator{contract: _Committer.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Committer *CommitterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CommitterInitialized) (event.Subscription, error) {

	logs, sub, err := _Committer.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitterInitialized)
				if err := _Committer.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Committer *CommitterFilterer) ParseInitialized(log types.Log) (*CommitterInitialized, error) {
	event := new(CommitterInitialized)
	if err := _Committer.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommitterProposalSubmittedIterator is returned from FilterProposalSubmitted and is used to iterate over the raw logs and unpacked data for ProposalSubmitted events raised by the Committer contract.
type CommitterProposalSubmittedIterator struct {
	Event *CommitterProposalSubmitted // Event containing the contract specifics and raw log

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
func (it *CommitterProposalSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitterProposalSubmitted)
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
		it.Event = new(CommitterProposalSubmitted)
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
func (it *CommitterProposalSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitterProposalSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitterProposalSubmitted represents a ProposalSubmitted event raised by the Committer contract.
type CommitterProposalSubmitted struct {
	ChainID       uint64
	ProposalID    uint64
	Proposer      common.Address
	ProofHash     string
	StateRootHash string
	StartIndex    *big.Int
	EndIndex      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProposalSubmitted is a free log retrieval operation binding the contract event 0x08db083a8474187b3aaed9068af476148cc1c18a272491f43b8a93c5ab91606a.
//
// Solidity: event ProposalSubmitted(uint64 chainID, uint64 indexed proposalID, address indexed proposer, string proofHash, string stateRootHash, uint256 indexed startIndex, uint256 endIndex)
func (_Committer *CommitterFilterer) FilterProposalSubmitted(opts *bind.FilterOpts, proposalID []uint64, proposer []common.Address, startIndex []*big.Int) (*CommitterProposalSubmittedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	var startIndexRule []interface{}
	for _, startIndexItem := range startIndex {
		startIndexRule = append(startIndexRule, startIndexItem)
	}

	logs, sub, err := _Committer.contract.FilterLogs(opts, "ProposalSubmitted", proposalIDRule, proposerRule, startIndexRule)
	if err != nil {
		return nil, err
	}
	return &CommitterProposalSubmittedIterator{contract: _Committer.contract, event: "ProposalSubmitted", logs: logs, sub: sub}, nil
}

// WatchProposalSubmitted is a free log subscription operation binding the contract event 0x08db083a8474187b3aaed9068af476148cc1c18a272491f43b8a93c5ab91606a.
//
// Solidity: event ProposalSubmitted(uint64 chainID, uint64 indexed proposalID, address indexed proposer, string proofHash, string stateRootHash, uint256 indexed startIndex, uint256 endIndex)
func (_Committer *CommitterFilterer) WatchProposalSubmitted(opts *bind.WatchOpts, sink chan<- *CommitterProposalSubmitted, proposalID []uint64, proposer []common.Address, startIndex []*big.Int) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	var startIndexRule []interface{}
	for _, startIndexItem := range startIndex {
		startIndexRule = append(startIndexRule, startIndexItem)
	}

	logs, sub, err := _Committer.contract.WatchLogs(opts, "ProposalSubmitted", proposalIDRule, proposerRule, startIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitterProposalSubmitted)
				if err := _Committer.contract.UnpackLog(event, "ProposalSubmitted", log); err != nil {
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

// ParseProposalSubmitted is a log parse operation binding the contract event 0x08db083a8474187b3aaed9068af476148cc1c18a272491f43b8a93c5ab91606a.
//
// Solidity: event ProposalSubmitted(uint64 chainID, uint64 indexed proposalID, address indexed proposer, string proofHash, string stateRootHash, uint256 indexed startIndex, uint256 endIndex)
func (_Committer *CommitterFilterer) ParseProposalSubmitted(log types.Log) (*CommitterProposalSubmitted, error) {
	event := new(CommitterProposalSubmitted)
	if err := _Committer.contract.UnpackLog(event, "ProposalSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommitterProposalTimedOutIterator is returned from FilterProposalTimedOut and is used to iterate over the raw logs and unpacked data for ProposalTimedOut events raised by the Committer contract.
type CommitterProposalTimedOutIterator struct {
	Event *CommitterProposalTimedOut // Event containing the contract specifics and raw log

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
func (it *CommitterProposalTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitterProposalTimedOut)
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
		it.Event = new(CommitterProposalTimedOut)
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
func (it *CommitterProposalTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitterProposalTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitterProposalTimedOut represents a ProposalTimedOut event raised by the Committer contract.
type CommitterProposalTimedOut struct {
	ChainID    uint64
	ProposalID uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalTimedOut is a free log retrieval operation binding the contract event 0xe81e592c2f9a76eb49a3b40a0a0f6955323a7ec31c08805d2c04a4eea8b558fd.
//
// Solidity: event ProposalTimedOut(uint64 chainID, uint64 proposalID)
func (_Committer *CommitterFilterer) FilterProposalTimedOut(opts *bind.FilterOpts) (*CommitterProposalTimedOutIterator, error) {

	logs, sub, err := _Committer.contract.FilterLogs(opts, "ProposalTimedOut")
	if err != nil {
		return nil, err
	}
	return &CommitterProposalTimedOutIterator{contract: _Committer.contract, event: "ProposalTimedOut", logs: logs, sub: sub}, nil
}

// WatchProposalTimedOut is a free log subscription operation binding the contract event 0xe81e592c2f9a76eb49a3b40a0a0f6955323a7ec31c08805d2c04a4eea8b558fd.
//
// Solidity: event ProposalTimedOut(uint64 chainID, uint64 proposalID)
func (_Committer *CommitterFilterer) WatchProposalTimedOut(opts *bind.WatchOpts, sink chan<- *CommitterProposalTimedOut) (event.Subscription, error) {

	logs, sub, err := _Committer.contract.WatchLogs(opts, "ProposalTimedOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitterProposalTimedOut)
				if err := _Committer.contract.UnpackLog(event, "ProposalTimedOut", log); err != nil {
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

// ParseProposalTimedOut is a log parse operation binding the contract event 0xe81e592c2f9a76eb49a3b40a0a0f6955323a7ec31c08805d2c04a4eea8b558fd.
//
// Solidity: event ProposalTimedOut(uint64 chainID, uint64 proposalID)
func (_Committer *CommitterFilterer) ParseProposalTimedOut(log types.Log) (*CommitterProposalTimedOut, error) {
	event := new(CommitterProposalTimedOut)
	if err := _Committer.contract.UnpackLog(event, "ProposalTimedOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommitterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Committer contract.
type CommitterRoleAdminChangedIterator struct {
	Event *CommitterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CommitterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitterRoleAdminChanged)
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
		it.Event = new(CommitterRoleAdminChanged)
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
func (it *CommitterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitterRoleAdminChanged represents a RoleAdminChanged event raised by the Committer contract.
type CommitterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Committer *CommitterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CommitterRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Committer.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CommitterRoleAdminChangedIterator{contract: _Committer.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Committer *CommitterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CommitterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Committer.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitterRoleAdminChanged)
				if err := _Committer.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Committer *CommitterFilterer) ParseRoleAdminChanged(log types.Log) (*CommitterRoleAdminChanged, error) {
	event := new(CommitterRoleAdminChanged)
	if err := _Committer.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommitterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Committer contract.
type CommitterRoleGrantedIterator struct {
	Event *CommitterRoleGranted // Event containing the contract specifics and raw log

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
func (it *CommitterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitterRoleGranted)
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
		it.Event = new(CommitterRoleGranted)
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
func (it *CommitterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitterRoleGranted represents a RoleGranted event raised by the Committer contract.
type CommitterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Committer *CommitterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CommitterRoleGrantedIterator, error) {

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

	logs, sub, err := _Committer.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CommitterRoleGrantedIterator{contract: _Committer.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Committer *CommitterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CommitterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Committer.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitterRoleGranted)
				if err := _Committer.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Committer *CommitterFilterer) ParseRoleGranted(log types.Log) (*CommitterRoleGranted, error) {
	event := new(CommitterRoleGranted)
	if err := _Committer.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommitterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Committer contract.
type CommitterRoleRevokedIterator struct {
	Event *CommitterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CommitterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitterRoleRevoked)
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
		it.Event = new(CommitterRoleRevoked)
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
func (it *CommitterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitterRoleRevoked represents a RoleRevoked event raised by the Committer contract.
type CommitterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Committer *CommitterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CommitterRoleRevokedIterator, error) {

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

	logs, sub, err := _Committer.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CommitterRoleRevokedIterator{contract: _Committer.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Committer *CommitterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CommitterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Committer.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitterRoleRevoked)
				if err := _Committer.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Committer *CommitterFilterer) ParseRoleRevoked(log types.Log) (*CommitterRoleRevoked, error) {
	event := new(CommitterRoleRevoked)
	if err := _Committer.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CommitterVoteProposalIterator is returned from FilterVoteProposal and is used to iterate over the raw logs and unpacked data for VoteProposal events raised by the Committer contract.
type CommitterVoteProposalIterator struct {
	Event *CommitterVoteProposal // Event containing the contract specifics and raw log

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
func (it *CommitterVoteProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommitterVoteProposal)
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
		it.Event = new(CommitterVoteProposal)
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
func (it *CommitterVoteProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommitterVoteProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommitterVoteProposal represents a VoteProposal event raised by the Committer contract.
type CommitterVoteProposal struct {
	ChainID    uint64
	ProposalID uint64
	Voter      common.Address
	Phase      uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteProposal is a free log retrieval operation binding the contract event 0xc488e7e4772e55cf4eb67cd1cc5a8a0169f7b08a7565742e419b88aac6744116.
//
// Solidity: event VoteProposal(uint64 chainID, uint64 proposalID, address voter, uint8 phase)
func (_Committer *CommitterFilterer) FilterVoteProposal(opts *bind.FilterOpts) (*CommitterVoteProposalIterator, error) {

	logs, sub, err := _Committer.contract.FilterLogs(opts, "VoteProposal")
	if err != nil {
		return nil, err
	}
	return &CommitterVoteProposalIterator{contract: _Committer.contract, event: "VoteProposal", logs: logs, sub: sub}, nil
}

// WatchVoteProposal is a free log subscription operation binding the contract event 0xc488e7e4772e55cf4eb67cd1cc5a8a0169f7b08a7565742e419b88aac6744116.
//
// Solidity: event VoteProposal(uint64 chainID, uint64 proposalID, address voter, uint8 phase)
func (_Committer *CommitterFilterer) WatchVoteProposal(opts *bind.WatchOpts, sink chan<- *CommitterVoteProposal) (event.Subscription, error) {

	logs, sub, err := _Committer.contract.WatchLogs(opts, "VoteProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommitterVoteProposal)
				if err := _Committer.contract.UnpackLog(event, "VoteProposal", log); err != nil {
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

// ParseVoteProposal is a log parse operation binding the contract event 0xc488e7e4772e55cf4eb67cd1cc5a8a0169f7b08a7565742e419b88aac6744116.
//
// Solidity: event VoteProposal(uint64 chainID, uint64 proposalID, address voter, uint8 phase)
func (_Committer *CommitterFilterer) ParseVoteProposal(log types.Log) (*CommitterVoteProposal, error) {
	event := new(CommitterVoteProposal)
	if err := _Committer.contract.UnpackLog(event, "VoteProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
