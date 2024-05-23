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

// OpCommitterMetaData contains all meta data concerning the OpCommitter contract.
var OpCommitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"outputRoot\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startL1Timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endL1Timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startL2BlockNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endL2BlockNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"outputStartIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"outputEndIndex\",\"type\":\"uint64\"}],\"name\":\"StateRootSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txsRoot\",\"type\":\"string\"}],\"name\":\"TxsRootSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"name\":\"bitcoinTxHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalType\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dsType\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"name\":\"dsTxHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proposalManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalManager\",\"outputs\":[{\"internalType\":\"contractIOpProposalManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposer\",\"outputs\":[{\"internalType\":\"contractIProposer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proposalManager\",\"type\":\"address\"}],\"name\":\"setProposalManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"}],\"name\":\"setProposer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setStateRootTimeoutPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setTxsRootTimeoutPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateRootTimeoutPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"outputRoot\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"startL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL1Timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endL2BlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputStartIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"outputEndIndex\",\"type\":\"uint64\"}],\"name\":\"submitStateRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"endBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"txsRoot\",\"type\":\"string\"},{\"internalType\":\"uint64[]\",\"name\":\"blockList\",\"type\":\"uint64[]\"}],\"name\":\"submitTxsRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposalType\",\"type\":\"uint64\"}],\"name\":\"timeoutProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"txsRootTimeoutPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OpCommitterABI is the input ABI used to generate the binding from.
// Deprecated: Use OpCommitterMetaData.ABI instead.
var OpCommitterABI = OpCommitterMetaData.ABI

// OpCommitter is an auto generated Go binding around an Ethereum contract.
type OpCommitter struct {
	OpCommitterCaller     // Read-only binding to the contract
	OpCommitterTransactor // Write-only binding to the contract
	OpCommitterFilterer   // Log filterer for contract events
}

// OpCommitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpCommitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpCommitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpCommitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpCommitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpCommitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpCommitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpCommitterSession struct {
	Contract     *OpCommitter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpCommitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpCommitterCallerSession struct {
	Contract *OpCommitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OpCommitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpCommitterTransactorSession struct {
	Contract     *OpCommitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OpCommitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpCommitterRaw struct {
	Contract *OpCommitter // Generic contract binding to access the raw methods on
}

// OpCommitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpCommitterCallerRaw struct {
	Contract *OpCommitterCaller // Generic read-only contract binding to access the raw methods on
}

// OpCommitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpCommitterTransactorRaw struct {
	Contract *OpCommitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpCommitter creates a new instance of OpCommitter, bound to a specific deployed contract.
func NewOpCommitter(address common.Address, backend bind.ContractBackend) (*OpCommitter, error) {
	contract, err := bindOpCommitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpCommitter{OpCommitterCaller: OpCommitterCaller{contract: contract}, OpCommitterTransactor: OpCommitterTransactor{contract: contract}, OpCommitterFilterer: OpCommitterFilterer{contract: contract}}, nil
}

// NewOpCommitterCaller creates a new read-only instance of OpCommitter, bound to a specific deployed contract.
func NewOpCommitterCaller(address common.Address, caller bind.ContractCaller) (*OpCommitterCaller, error) {
	contract, err := bindOpCommitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpCommitterCaller{contract: contract}, nil
}

// NewOpCommitterTransactor creates a new write-only instance of OpCommitter, bound to a specific deployed contract.
func NewOpCommitterTransactor(address common.Address, transactor bind.ContractTransactor) (*OpCommitterTransactor, error) {
	contract, err := bindOpCommitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpCommitterTransactor{contract: contract}, nil
}

// NewOpCommitterFilterer creates a new log filterer instance of OpCommitter, bound to a specific deployed contract.
func NewOpCommitterFilterer(address common.Address, filterer bind.ContractFilterer) (*OpCommitterFilterer, error) {
	contract, err := bindOpCommitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpCommitterFilterer{contract: contract}, nil
}

// bindOpCommitter binds a generic wrapper to an already deployed contract.
func bindOpCommitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OpCommitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpCommitter *OpCommitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpCommitter.Contract.OpCommitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpCommitter *OpCommitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpCommitter.Contract.OpCommitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpCommitter *OpCommitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpCommitter.Contract.OpCommitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpCommitter *OpCommitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpCommitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpCommitter *OpCommitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpCommitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpCommitter *OpCommitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpCommitter.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_OpCommitter *OpCommitterCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OpCommitter.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_OpCommitter *OpCommitterSession) ADMINROLE() ([32]byte, error) {
	return _OpCommitter.Contract.ADMINROLE(&_OpCommitter.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_OpCommitter *OpCommitterCallerSession) ADMINROLE() ([32]byte, error) {
	return _OpCommitter.Contract.ADMINROLE(&_OpCommitter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OpCommitter *OpCommitterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OpCommitter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OpCommitter *OpCommitterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OpCommitter.Contract.DEFAULTADMINROLE(&_OpCommitter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_OpCommitter *OpCommitterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _OpCommitter.Contract.DEFAULTADMINROLE(&_OpCommitter.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OpCommitter *OpCommitterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _OpCommitter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OpCommitter *OpCommitterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OpCommitter.Contract.GetRoleAdmin(&_OpCommitter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_OpCommitter *OpCommitterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _OpCommitter.Contract.GetRoleAdmin(&_OpCommitter.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OpCommitter *OpCommitterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _OpCommitter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OpCommitter *OpCommitterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OpCommitter.Contract.HasRole(&_OpCommitter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_OpCommitter *OpCommitterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _OpCommitter.Contract.HasRole(&_OpCommitter.CallOpts, role, account)
}

// ProposalManager is a free data retrieval call binding the contract method 0x02f89be2.
//
// Solidity: function proposalManager() view returns(address)
func (_OpCommitter *OpCommitterCaller) ProposalManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpCommitter.contract.Call(opts, &out, "proposalManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalManager is a free data retrieval call binding the contract method 0x02f89be2.
//
// Solidity: function proposalManager() view returns(address)
func (_OpCommitter *OpCommitterSession) ProposalManager() (common.Address, error) {
	return _OpCommitter.Contract.ProposalManager(&_OpCommitter.CallOpts)
}

// ProposalManager is a free data retrieval call binding the contract method 0x02f89be2.
//
// Solidity: function proposalManager() view returns(address)
func (_OpCommitter *OpCommitterCallerSession) ProposalManager() (common.Address, error) {
	return _OpCommitter.Contract.ProposalManager(&_OpCommitter.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_OpCommitter *OpCommitterCaller) Proposer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OpCommitter.contract.Call(opts, &out, "proposer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_OpCommitter *OpCommitterSession) Proposer() (common.Address, error) {
	return _OpCommitter.Contract.Proposer(&_OpCommitter.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_OpCommitter *OpCommitterCallerSession) Proposer() (common.Address, error) {
	return _OpCommitter.Contract.Proposer(&_OpCommitter.CallOpts)
}

// StateRootTimeoutPeriod is a free data retrieval call binding the contract method 0x77361d85.
//
// Solidity: function stateRootTimeoutPeriod() view returns(uint256)
func (_OpCommitter *OpCommitterCaller) StateRootTimeoutPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OpCommitter.contract.Call(opts, &out, "stateRootTimeoutPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateRootTimeoutPeriod is a free data retrieval call binding the contract method 0x77361d85.
//
// Solidity: function stateRootTimeoutPeriod() view returns(uint256)
func (_OpCommitter *OpCommitterSession) StateRootTimeoutPeriod() (*big.Int, error) {
	return _OpCommitter.Contract.StateRootTimeoutPeriod(&_OpCommitter.CallOpts)
}

// StateRootTimeoutPeriod is a free data retrieval call binding the contract method 0x77361d85.
//
// Solidity: function stateRootTimeoutPeriod() view returns(uint256)
func (_OpCommitter *OpCommitterCallerSession) StateRootTimeoutPeriod() (*big.Int, error) {
	return _OpCommitter.Contract.StateRootTimeoutPeriod(&_OpCommitter.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OpCommitter *OpCommitterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _OpCommitter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OpCommitter *OpCommitterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OpCommitter.Contract.SupportsInterface(&_OpCommitter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_OpCommitter *OpCommitterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _OpCommitter.Contract.SupportsInterface(&_OpCommitter.CallOpts, interfaceId)
}

// TxsRootTimeoutPeriod is a free data retrieval call binding the contract method 0x15c4441f.
//
// Solidity: function txsRootTimeoutPeriod() view returns(uint256)
func (_OpCommitter *OpCommitterCaller) TxsRootTimeoutPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OpCommitter.contract.Call(opts, &out, "txsRootTimeoutPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TxsRootTimeoutPeriod is a free data retrieval call binding the contract method 0x15c4441f.
//
// Solidity: function txsRootTimeoutPeriod() view returns(uint256)
func (_OpCommitter *OpCommitterSession) TxsRootTimeoutPeriod() (*big.Int, error) {
	return _OpCommitter.Contract.TxsRootTimeoutPeriod(&_OpCommitter.CallOpts)
}

// TxsRootTimeoutPeriod is a free data retrieval call binding the contract method 0x15c4441f.
//
// Solidity: function txsRootTimeoutPeriod() view returns(uint256)
func (_OpCommitter *OpCommitterCallerSession) TxsRootTimeoutPeriod() (*big.Int, error) {
	return _OpCommitter.Contract.TxsRootTimeoutPeriod(&_OpCommitter.CallOpts)
}

// BitcoinTxHash is a paid mutator transaction binding the contract method 0xa5772e5e.
//
// Solidity: function bitcoinTxHash(uint64 proposalID, string txHash) returns()
func (_OpCommitter *OpCommitterTransactor) BitcoinTxHash(opts *bind.TransactOpts, proposalID uint64, txHash string) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "bitcoinTxHash", proposalID, txHash)
}

// BitcoinTxHash is a paid mutator transaction binding the contract method 0xa5772e5e.
//
// Solidity: function bitcoinTxHash(uint64 proposalID, string txHash) returns()
func (_OpCommitter *OpCommitterSession) BitcoinTxHash(proposalID uint64, txHash string) (*types.Transaction, error) {
	return _OpCommitter.Contract.BitcoinTxHash(&_OpCommitter.TransactOpts, proposalID, txHash)
}

// BitcoinTxHash is a paid mutator transaction binding the contract method 0xa5772e5e.
//
// Solidity: function bitcoinTxHash(uint64 proposalID, string txHash) returns()
func (_OpCommitter *OpCommitterTransactorSession) BitcoinTxHash(proposalID uint64, txHash string) (*types.Transaction, error) {
	return _OpCommitter.Contract.BitcoinTxHash(&_OpCommitter.TransactOpts, proposalID, txHash)
}

// DsTxHash is a paid mutator transaction binding the contract method 0xe61bd62f.
//
// Solidity: function dsTxHash(uint64 proposalID, uint64 proposalType, uint64 dsType, string txHash) returns()
func (_OpCommitter *OpCommitterTransactor) DsTxHash(opts *bind.TransactOpts, proposalID uint64, proposalType uint64, dsType uint64, txHash string) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "dsTxHash", proposalID, proposalType, dsType, txHash)
}

// DsTxHash is a paid mutator transaction binding the contract method 0xe61bd62f.
//
// Solidity: function dsTxHash(uint64 proposalID, uint64 proposalType, uint64 dsType, string txHash) returns()
func (_OpCommitter *OpCommitterSession) DsTxHash(proposalID uint64, proposalType uint64, dsType uint64, txHash string) (*types.Transaction, error) {
	return _OpCommitter.Contract.DsTxHash(&_OpCommitter.TransactOpts, proposalID, proposalType, dsType, txHash)
}

// DsTxHash is a paid mutator transaction binding the contract method 0xe61bd62f.
//
// Solidity: function dsTxHash(uint64 proposalID, uint64 proposalType, uint64 dsType, string txHash) returns()
func (_OpCommitter *OpCommitterTransactorSession) DsTxHash(proposalID uint64, proposalType uint64, dsType uint64, txHash string) (*types.Transaction, error) {
	return _OpCommitter.Contract.DsTxHash(&_OpCommitter.TransactOpts, proposalID, proposalType, dsType, txHash)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OpCommitter *OpCommitterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OpCommitter *OpCommitterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpCommitter.Contract.GrantRole(&_OpCommitter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_OpCommitter *OpCommitterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpCommitter.Contract.GrantRole(&_OpCommitter.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _proposer, address _proposalManager) returns()
func (_OpCommitter *OpCommitterTransactor) Initialize(opts *bind.TransactOpts, _proposer common.Address, _proposalManager common.Address) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "initialize", _proposer, _proposalManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _proposer, address _proposalManager) returns()
func (_OpCommitter *OpCommitterSession) Initialize(_proposer common.Address, _proposalManager common.Address) (*types.Transaction, error) {
	return _OpCommitter.Contract.Initialize(&_OpCommitter.TransactOpts, _proposer, _proposalManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _proposer, address _proposalManager) returns()
func (_OpCommitter *OpCommitterTransactorSession) Initialize(_proposer common.Address, _proposalManager common.Address) (*types.Transaction, error) {
	return _OpCommitter.Contract.Initialize(&_OpCommitter.TransactOpts, _proposer, _proposalManager)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OpCommitter *OpCommitterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OpCommitter *OpCommitterSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OpCommitter.Contract.RenounceRole(&_OpCommitter.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_OpCommitter *OpCommitterTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _OpCommitter.Contract.RenounceRole(&_OpCommitter.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OpCommitter *OpCommitterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OpCommitter *OpCommitterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpCommitter.Contract.RevokeRole(&_OpCommitter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_OpCommitter *OpCommitterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _OpCommitter.Contract.RevokeRole(&_OpCommitter.TransactOpts, role, account)
}

// SetProposalManager is a paid mutator transaction binding the contract method 0x25023f0b.
//
// Solidity: function setProposalManager(address _proposalManager) returns()
func (_OpCommitter *OpCommitterTransactor) SetProposalManager(opts *bind.TransactOpts, _proposalManager common.Address) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "setProposalManager", _proposalManager)
}

// SetProposalManager is a paid mutator transaction binding the contract method 0x25023f0b.
//
// Solidity: function setProposalManager(address _proposalManager) returns()
func (_OpCommitter *OpCommitterSession) SetProposalManager(_proposalManager common.Address) (*types.Transaction, error) {
	return _OpCommitter.Contract.SetProposalManager(&_OpCommitter.TransactOpts, _proposalManager)
}

// SetProposalManager is a paid mutator transaction binding the contract method 0x25023f0b.
//
// Solidity: function setProposalManager(address _proposalManager) returns()
func (_OpCommitter *OpCommitterTransactorSession) SetProposalManager(_proposalManager common.Address) (*types.Transaction, error) {
	return _OpCommitter.Contract.SetProposalManager(&_OpCommitter.TransactOpts, _proposalManager)
}

// SetProposer is a paid mutator transaction binding the contract method 0x1fb4a228.
//
// Solidity: function setProposer(address _proposer) returns()
func (_OpCommitter *OpCommitterTransactor) SetProposer(opts *bind.TransactOpts, _proposer common.Address) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "setProposer", _proposer)
}

// SetProposer is a paid mutator transaction binding the contract method 0x1fb4a228.
//
// Solidity: function setProposer(address _proposer) returns()
func (_OpCommitter *OpCommitterSession) SetProposer(_proposer common.Address) (*types.Transaction, error) {
	return _OpCommitter.Contract.SetProposer(&_OpCommitter.TransactOpts, _proposer)
}

// SetProposer is a paid mutator transaction binding the contract method 0x1fb4a228.
//
// Solidity: function setProposer(address _proposer) returns()
func (_OpCommitter *OpCommitterTransactorSession) SetProposer(_proposer common.Address) (*types.Transaction, error) {
	return _OpCommitter.Contract.SetProposer(&_OpCommitter.TransactOpts, _proposer)
}

// SetStateRootTimeoutPeriod is a paid mutator transaction binding the contract method 0x6ed9808a.
//
// Solidity: function setStateRootTimeoutPeriod(uint256 period) returns()
func (_OpCommitter *OpCommitterTransactor) SetStateRootTimeoutPeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "setStateRootTimeoutPeriod", period)
}

// SetStateRootTimeoutPeriod is a paid mutator transaction binding the contract method 0x6ed9808a.
//
// Solidity: function setStateRootTimeoutPeriod(uint256 period) returns()
func (_OpCommitter *OpCommitterSession) SetStateRootTimeoutPeriod(period *big.Int) (*types.Transaction, error) {
	return _OpCommitter.Contract.SetStateRootTimeoutPeriod(&_OpCommitter.TransactOpts, period)
}

// SetStateRootTimeoutPeriod is a paid mutator transaction binding the contract method 0x6ed9808a.
//
// Solidity: function setStateRootTimeoutPeriod(uint256 period) returns()
func (_OpCommitter *OpCommitterTransactorSession) SetStateRootTimeoutPeriod(period *big.Int) (*types.Transaction, error) {
	return _OpCommitter.Contract.SetStateRootTimeoutPeriod(&_OpCommitter.TransactOpts, period)
}

// SetTxsRootTimeoutPeriod is a paid mutator transaction binding the contract method 0x23f7743d.
//
// Solidity: function setTxsRootTimeoutPeriod(uint256 period) returns()
func (_OpCommitter *OpCommitterTransactor) SetTxsRootTimeoutPeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "setTxsRootTimeoutPeriod", period)
}

// SetTxsRootTimeoutPeriod is a paid mutator transaction binding the contract method 0x23f7743d.
//
// Solidity: function setTxsRootTimeoutPeriod(uint256 period) returns()
func (_OpCommitter *OpCommitterSession) SetTxsRootTimeoutPeriod(period *big.Int) (*types.Transaction, error) {
	return _OpCommitter.Contract.SetTxsRootTimeoutPeriod(&_OpCommitter.TransactOpts, period)
}

// SetTxsRootTimeoutPeriod is a paid mutator transaction binding the contract method 0x23f7743d.
//
// Solidity: function setTxsRootTimeoutPeriod(uint256 period) returns()
func (_OpCommitter *OpCommitterTransactorSession) SetTxsRootTimeoutPeriod(period *big.Int) (*types.Transaction, error) {
	return _OpCommitter.Contract.SetTxsRootTimeoutPeriod(&_OpCommitter.TransactOpts, period)
}

// SubmitStateRoot is a paid mutator transaction binding the contract method 0xc7b41985.
//
// Solidity: function submitStateRoot(uint64 proposalID, string outputRoot, uint64 startL1Timestamp, uint64 endL1Timestamp, uint64 startL2BlockNumber, uint64 endL2BlockNumber, uint64 outputStartIndex, uint64 outputEndIndex) returns()
func (_OpCommitter *OpCommitterTransactor) SubmitStateRoot(opts *bind.TransactOpts, proposalID uint64, outputRoot string, startL1Timestamp uint64, endL1Timestamp uint64, startL2BlockNumber uint64, endL2BlockNumber uint64, outputStartIndex uint64, outputEndIndex uint64) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "submitStateRoot", proposalID, outputRoot, startL1Timestamp, endL1Timestamp, startL2BlockNumber, endL2BlockNumber, outputStartIndex, outputEndIndex)
}

// SubmitStateRoot is a paid mutator transaction binding the contract method 0xc7b41985.
//
// Solidity: function submitStateRoot(uint64 proposalID, string outputRoot, uint64 startL1Timestamp, uint64 endL1Timestamp, uint64 startL2BlockNumber, uint64 endL2BlockNumber, uint64 outputStartIndex, uint64 outputEndIndex) returns()
func (_OpCommitter *OpCommitterSession) SubmitStateRoot(proposalID uint64, outputRoot string, startL1Timestamp uint64, endL1Timestamp uint64, startL2BlockNumber uint64, endL2BlockNumber uint64, outputStartIndex uint64, outputEndIndex uint64) (*types.Transaction, error) {
	return _OpCommitter.Contract.SubmitStateRoot(&_OpCommitter.TransactOpts, proposalID, outputRoot, startL1Timestamp, endL1Timestamp, startL2BlockNumber, endL2BlockNumber, outputStartIndex, outputEndIndex)
}

// SubmitStateRoot is a paid mutator transaction binding the contract method 0xc7b41985.
//
// Solidity: function submitStateRoot(uint64 proposalID, string outputRoot, uint64 startL1Timestamp, uint64 endL1Timestamp, uint64 startL2BlockNumber, uint64 endL2BlockNumber, uint64 outputStartIndex, uint64 outputEndIndex) returns()
func (_OpCommitter *OpCommitterTransactorSession) SubmitStateRoot(proposalID uint64, outputRoot string, startL1Timestamp uint64, endL1Timestamp uint64, startL2BlockNumber uint64, endL2BlockNumber uint64, outputStartIndex uint64, outputEndIndex uint64) (*types.Transaction, error) {
	return _OpCommitter.Contract.SubmitStateRoot(&_OpCommitter.TransactOpts, proposalID, outputRoot, startL1Timestamp, endL1Timestamp, startL2BlockNumber, endL2BlockNumber, outputStartIndex, outputEndIndex)
}

// SubmitTxsRoot is a paid mutator transaction binding the contract method 0xc4bd5e19.
//
// Solidity: function submitTxsRoot(uint64 proposalID, uint64 startTimestamp, uint64 endTimestamp, uint64 startBlockNumber, uint64 endBlockNumber, string txsRoot, uint64[] blockList) returns()
func (_OpCommitter *OpCommitterTransactor) SubmitTxsRoot(opts *bind.TransactOpts, proposalID uint64, startTimestamp uint64, endTimestamp uint64, startBlockNumber uint64, endBlockNumber uint64, txsRoot string, blockList []uint64) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "submitTxsRoot", proposalID, startTimestamp, endTimestamp, startBlockNumber, endBlockNumber, txsRoot, blockList)
}

// SubmitTxsRoot is a paid mutator transaction binding the contract method 0xc4bd5e19.
//
// Solidity: function submitTxsRoot(uint64 proposalID, uint64 startTimestamp, uint64 endTimestamp, uint64 startBlockNumber, uint64 endBlockNumber, string txsRoot, uint64[] blockList) returns()
func (_OpCommitter *OpCommitterSession) SubmitTxsRoot(proposalID uint64, startTimestamp uint64, endTimestamp uint64, startBlockNumber uint64, endBlockNumber uint64, txsRoot string, blockList []uint64) (*types.Transaction, error) {
	return _OpCommitter.Contract.SubmitTxsRoot(&_OpCommitter.TransactOpts, proposalID, startTimestamp, endTimestamp, startBlockNumber, endBlockNumber, txsRoot, blockList)
}

// SubmitTxsRoot is a paid mutator transaction binding the contract method 0xc4bd5e19.
//
// Solidity: function submitTxsRoot(uint64 proposalID, uint64 startTimestamp, uint64 endTimestamp, uint64 startBlockNumber, uint64 endBlockNumber, string txsRoot, uint64[] blockList) returns()
func (_OpCommitter *OpCommitterTransactorSession) SubmitTxsRoot(proposalID uint64, startTimestamp uint64, endTimestamp uint64, startBlockNumber uint64, endBlockNumber uint64, txsRoot string, blockList []uint64) (*types.Transaction, error) {
	return _OpCommitter.Contract.SubmitTxsRoot(&_OpCommitter.TransactOpts, proposalID, startTimestamp, endTimestamp, startBlockNumber, endBlockNumber, txsRoot, blockList)
}

// TimeoutProposal is a paid mutator transaction binding the contract method 0x04956a75.
//
// Solidity: function timeoutProposal(uint64 proposalID, uint64 proposalType) returns()
func (_OpCommitter *OpCommitterTransactor) TimeoutProposal(opts *bind.TransactOpts, proposalID uint64, proposalType uint64) (*types.Transaction, error) {
	return _OpCommitter.contract.Transact(opts, "timeoutProposal", proposalID, proposalType)
}

// TimeoutProposal is a paid mutator transaction binding the contract method 0x04956a75.
//
// Solidity: function timeoutProposal(uint64 proposalID, uint64 proposalType) returns()
func (_OpCommitter *OpCommitterSession) TimeoutProposal(proposalID uint64, proposalType uint64) (*types.Transaction, error) {
	return _OpCommitter.Contract.TimeoutProposal(&_OpCommitter.TransactOpts, proposalID, proposalType)
}

// TimeoutProposal is a paid mutator transaction binding the contract method 0x04956a75.
//
// Solidity: function timeoutProposal(uint64 proposalID, uint64 proposalType) returns()
func (_OpCommitter *OpCommitterTransactorSession) TimeoutProposal(proposalID uint64, proposalType uint64) (*types.Transaction, error) {
	return _OpCommitter.Contract.TimeoutProposal(&_OpCommitter.TransactOpts, proposalID, proposalType)
}

// OpCommitterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OpCommitter contract.
type OpCommitterInitializedIterator struct {
	Event *OpCommitterInitialized // Event containing the contract specifics and raw log

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
func (it *OpCommitterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpCommitterInitialized)
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
		it.Event = new(OpCommitterInitialized)
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
func (it *OpCommitterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpCommitterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpCommitterInitialized represents a Initialized event raised by the OpCommitter contract.
type OpCommitterInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OpCommitter *OpCommitterFilterer) FilterInitialized(opts *bind.FilterOpts) (*OpCommitterInitializedIterator, error) {

	logs, sub, err := _OpCommitter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OpCommitterInitializedIterator{contract: _OpCommitter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OpCommitter *OpCommitterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OpCommitterInitialized) (event.Subscription, error) {

	logs, sub, err := _OpCommitter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpCommitterInitialized)
				if err := _OpCommitter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OpCommitter *OpCommitterFilterer) ParseInitialized(log types.Log) (*OpCommitterInitialized, error) {
	event := new(OpCommitterInitialized)
	if err := _OpCommitter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpCommitterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the OpCommitter contract.
type OpCommitterRoleAdminChangedIterator struct {
	Event *OpCommitterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OpCommitterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpCommitterRoleAdminChanged)
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
		it.Event = new(OpCommitterRoleAdminChanged)
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
func (it *OpCommitterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpCommitterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpCommitterRoleAdminChanged represents a RoleAdminChanged event raised by the OpCommitter contract.
type OpCommitterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OpCommitter *OpCommitterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OpCommitterRoleAdminChangedIterator, error) {

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

	logs, sub, err := _OpCommitter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OpCommitterRoleAdminChangedIterator{contract: _OpCommitter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_OpCommitter *OpCommitterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OpCommitterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _OpCommitter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpCommitterRoleAdminChanged)
				if err := _OpCommitter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_OpCommitter *OpCommitterFilterer) ParseRoleAdminChanged(log types.Log) (*OpCommitterRoleAdminChanged, error) {
	event := new(OpCommitterRoleAdminChanged)
	if err := _OpCommitter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpCommitterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the OpCommitter contract.
type OpCommitterRoleGrantedIterator struct {
	Event *OpCommitterRoleGranted // Event containing the contract specifics and raw log

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
func (it *OpCommitterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpCommitterRoleGranted)
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
		it.Event = new(OpCommitterRoleGranted)
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
func (it *OpCommitterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpCommitterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpCommitterRoleGranted represents a RoleGranted event raised by the OpCommitter contract.
type OpCommitterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpCommitter *OpCommitterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OpCommitterRoleGrantedIterator, error) {

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

	logs, sub, err := _OpCommitter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OpCommitterRoleGrantedIterator{contract: _OpCommitter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpCommitter *OpCommitterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OpCommitterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OpCommitter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpCommitterRoleGranted)
				if err := _OpCommitter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_OpCommitter *OpCommitterFilterer) ParseRoleGranted(log types.Log) (*OpCommitterRoleGranted, error) {
	event := new(OpCommitterRoleGranted)
	if err := _OpCommitter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpCommitterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the OpCommitter contract.
type OpCommitterRoleRevokedIterator struct {
	Event *OpCommitterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *OpCommitterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpCommitterRoleRevoked)
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
		it.Event = new(OpCommitterRoleRevoked)
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
func (it *OpCommitterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpCommitterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpCommitterRoleRevoked represents a RoleRevoked event raised by the OpCommitter contract.
type OpCommitterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpCommitter *OpCommitterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OpCommitterRoleRevokedIterator, error) {

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

	logs, sub, err := _OpCommitter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OpCommitterRoleRevokedIterator{contract: _OpCommitter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_OpCommitter *OpCommitterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OpCommitterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OpCommitter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpCommitterRoleRevoked)
				if err := _OpCommitter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_OpCommitter *OpCommitterFilterer) ParseRoleRevoked(log types.Log) (*OpCommitterRoleRevoked, error) {
	event := new(OpCommitterRoleRevoked)
	if err := _OpCommitter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpCommitterStateRootSubmittedIterator is returned from FilterStateRootSubmitted and is used to iterate over the raw logs and unpacked data for StateRootSubmitted events raised by the OpCommitter contract.
type OpCommitterStateRootSubmittedIterator struct {
	Event *OpCommitterStateRootSubmitted // Event containing the contract specifics and raw log

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
func (it *OpCommitterStateRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpCommitterStateRootSubmitted)
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
		it.Event = new(OpCommitterStateRootSubmitted)
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
func (it *OpCommitterStateRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpCommitterStateRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpCommitterStateRootSubmitted represents a StateRootSubmitted event raised by the OpCommitter contract.
type OpCommitterStateRootSubmitted struct {
	Proposer           common.Address
	ProposalID         uint64
	OutputRoot         string
	StartL1Timestamp   uint64
	EndL1Timestamp     uint64
	StartL2BlockNumber uint64
	EndL2BlockNumber   uint64
	OutputStartIndex   uint64
	OutputEndIndex     uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterStateRootSubmitted is a free log retrieval operation binding the contract event 0x1c0e748730db48af63aa72cf24a6070a26522decb8e17f58bd44ed079be45804.
//
// Solidity: event StateRootSubmitted(address indexed proposer, uint64 indexed proposalID, string outputRoot, uint64 startL1Timestamp, uint64 endL1Timestamp, uint64 startL2BlockNumber, uint64 endL2BlockNumber, uint64 outputStartIndex, uint64 outputEndIndex)
func (_OpCommitter *OpCommitterFilterer) FilterStateRootSubmitted(opts *bind.FilterOpts, proposer []common.Address, proposalID []uint64) (*OpCommitterStateRootSubmittedIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _OpCommitter.contract.FilterLogs(opts, "StateRootSubmitted", proposerRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &OpCommitterStateRootSubmittedIterator{contract: _OpCommitter.contract, event: "StateRootSubmitted", logs: logs, sub: sub}, nil
}

// WatchStateRootSubmitted is a free log subscription operation binding the contract event 0x1c0e748730db48af63aa72cf24a6070a26522decb8e17f58bd44ed079be45804.
//
// Solidity: event StateRootSubmitted(address indexed proposer, uint64 indexed proposalID, string outputRoot, uint64 startL1Timestamp, uint64 endL1Timestamp, uint64 startL2BlockNumber, uint64 endL2BlockNumber, uint64 outputStartIndex, uint64 outputEndIndex)
func (_OpCommitter *OpCommitterFilterer) WatchStateRootSubmitted(opts *bind.WatchOpts, sink chan<- *OpCommitterStateRootSubmitted, proposer []common.Address, proposalID []uint64) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _OpCommitter.contract.WatchLogs(opts, "StateRootSubmitted", proposerRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpCommitterStateRootSubmitted)
				if err := _OpCommitter.contract.UnpackLog(event, "StateRootSubmitted", log); err != nil {
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

// ParseStateRootSubmitted is a log parse operation binding the contract event 0x1c0e748730db48af63aa72cf24a6070a26522decb8e17f58bd44ed079be45804.
//
// Solidity: event StateRootSubmitted(address indexed proposer, uint64 indexed proposalID, string outputRoot, uint64 startL1Timestamp, uint64 endL1Timestamp, uint64 startL2BlockNumber, uint64 endL2BlockNumber, uint64 outputStartIndex, uint64 outputEndIndex)
func (_OpCommitter *OpCommitterFilterer) ParseStateRootSubmitted(log types.Log) (*OpCommitterStateRootSubmitted, error) {
	event := new(OpCommitterStateRootSubmitted)
	if err := _OpCommitter.contract.UnpackLog(event, "StateRootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpCommitterTxsRootSubmittedIterator is returned from FilterTxsRootSubmitted and is used to iterate over the raw logs and unpacked data for TxsRootSubmitted events raised by the OpCommitter contract.
type OpCommitterTxsRootSubmittedIterator struct {
	Event *OpCommitterTxsRootSubmitted // Event containing the contract specifics and raw log

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
func (it *OpCommitterTxsRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpCommitterTxsRootSubmitted)
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
		it.Event = new(OpCommitterTxsRootSubmitted)
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
func (it *OpCommitterTxsRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpCommitterTxsRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpCommitterTxsRootSubmitted represents a TxsRootSubmitted event raised by the OpCommitter contract.
type OpCommitterTxsRootSubmitted struct {
	Proposer         common.Address
	ProposalID       uint64
	StartTimestamp   uint64
	EndTimestamp     uint64
	StartBlockNumber uint64
	EndBlockNumber   uint64
	TxsRoot          string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTxsRootSubmitted is a free log retrieval operation binding the contract event 0x3ef8030bfab26ec3eaa1bb9e561677860e6794684c43e6b8a5bb50729281abad.
//
// Solidity: event TxsRootSubmitted(address indexed proposer, uint64 indexed proposalID, uint64 startTimestamp, uint64 endTimestamp, uint64 startBlockNumber, uint64 endBlockNumber, string txsRoot)
func (_OpCommitter *OpCommitterFilterer) FilterTxsRootSubmitted(opts *bind.FilterOpts, proposer []common.Address, proposalID []uint64) (*OpCommitterTxsRootSubmittedIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _OpCommitter.contract.FilterLogs(opts, "TxsRootSubmitted", proposerRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &OpCommitterTxsRootSubmittedIterator{contract: _OpCommitter.contract, event: "TxsRootSubmitted", logs: logs, sub: sub}, nil
}

// WatchTxsRootSubmitted is a free log subscription operation binding the contract event 0x3ef8030bfab26ec3eaa1bb9e561677860e6794684c43e6b8a5bb50729281abad.
//
// Solidity: event TxsRootSubmitted(address indexed proposer, uint64 indexed proposalID, uint64 startTimestamp, uint64 endTimestamp, uint64 startBlockNumber, uint64 endBlockNumber, string txsRoot)
func (_OpCommitter *OpCommitterFilterer) WatchTxsRootSubmitted(opts *bind.WatchOpts, sink chan<- *OpCommitterTxsRootSubmitted, proposer []common.Address, proposalID []uint64) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _OpCommitter.contract.WatchLogs(opts, "TxsRootSubmitted", proposerRule, proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpCommitterTxsRootSubmitted)
				if err := _OpCommitter.contract.UnpackLog(event, "TxsRootSubmitted", log); err != nil {
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

// ParseTxsRootSubmitted is a log parse operation binding the contract event 0x3ef8030bfab26ec3eaa1bb9e561677860e6794684c43e6b8a5bb50729281abad.
//
// Solidity: event TxsRootSubmitted(address indexed proposer, uint64 indexed proposalID, uint64 startTimestamp, uint64 endTimestamp, uint64 startBlockNumber, uint64 endBlockNumber, string txsRoot)
func (_OpCommitter *OpCommitterFilterer) ParseTxsRootSubmitted(log types.Log) (*OpCommitterTxsRootSubmitted, error) {
	event := new(OpCommitterTxsRootSubmitted)
	if err := _OpCommitter.contract.UnpackLog(event, "TxsRootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
