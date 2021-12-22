// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigo

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
)

// ChangeRewardsFacetMetaData contains all meta data concerning the ChangeRewardsFacet contract.
var ChangeRewardsFacetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"}],\"name\":\"changeRewardsAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ChangeRewardsFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use ChangeRewardsFacetMetaData.ABI instead.
var ChangeRewardsFacetABI = ChangeRewardsFacetMetaData.ABI

// ChangeRewardsFacet is an auto generated Go binding around an Ethereum contract.
type ChangeRewardsFacet struct {
	ChangeRewardsFacetCaller     // Read-only binding to the contract
	ChangeRewardsFacetTransactor // Write-only binding to the contract
	ChangeRewardsFacetFilterer   // Log filterer for contract events
}

// ChangeRewardsFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChangeRewardsFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChangeRewardsFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChangeRewardsFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChangeRewardsFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChangeRewardsFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChangeRewardsFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChangeRewardsFacetSession struct {
	Contract     *ChangeRewardsFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ChangeRewardsFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChangeRewardsFacetCallerSession struct {
	Contract *ChangeRewardsFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ChangeRewardsFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChangeRewardsFacetTransactorSession struct {
	Contract     *ChangeRewardsFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ChangeRewardsFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChangeRewardsFacetRaw struct {
	Contract *ChangeRewardsFacet // Generic contract binding to access the raw methods on
}

// ChangeRewardsFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChangeRewardsFacetCallerRaw struct {
	Contract *ChangeRewardsFacetCaller // Generic read-only contract binding to access the raw methods on
}

// ChangeRewardsFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChangeRewardsFacetTransactorRaw struct {
	Contract *ChangeRewardsFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChangeRewardsFacet creates a new instance of ChangeRewardsFacet, bound to a specific deployed contract.
func NewChangeRewardsFacet(address common.Address, backend bind.ContractBackend) (*ChangeRewardsFacet, error) {
	contract, err := bindChangeRewardsFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChangeRewardsFacet{ChangeRewardsFacetCaller: ChangeRewardsFacetCaller{contract: contract}, ChangeRewardsFacetTransactor: ChangeRewardsFacetTransactor{contract: contract}, ChangeRewardsFacetFilterer: ChangeRewardsFacetFilterer{contract: contract}}, nil
}

// NewChangeRewardsFacetCaller creates a new read-only instance of ChangeRewardsFacet, bound to a specific deployed contract.
func NewChangeRewardsFacetCaller(address common.Address, caller bind.ContractCaller) (*ChangeRewardsFacetCaller, error) {
	contract, err := bindChangeRewardsFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChangeRewardsFacetCaller{contract: contract}, nil
}

// NewChangeRewardsFacetTransactor creates a new write-only instance of ChangeRewardsFacet, bound to a specific deployed contract.
func NewChangeRewardsFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*ChangeRewardsFacetTransactor, error) {
	contract, err := bindChangeRewardsFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChangeRewardsFacetTransactor{contract: contract}, nil
}

// NewChangeRewardsFacetFilterer creates a new log filterer instance of ChangeRewardsFacet, bound to a specific deployed contract.
func NewChangeRewardsFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*ChangeRewardsFacetFilterer, error) {
	contract, err := bindChangeRewardsFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChangeRewardsFacetFilterer{contract: contract}, nil
}

// bindChangeRewardsFacet binds a generic wrapper to an already deployed contract.
func bindChangeRewardsFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChangeRewardsFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChangeRewardsFacet *ChangeRewardsFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChangeRewardsFacet.Contract.ChangeRewardsFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChangeRewardsFacet *ChangeRewardsFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChangeRewardsFacet.Contract.ChangeRewardsFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChangeRewardsFacet *ChangeRewardsFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChangeRewardsFacet.Contract.ChangeRewardsFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChangeRewardsFacet *ChangeRewardsFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChangeRewardsFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChangeRewardsFacet *ChangeRewardsFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChangeRewardsFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChangeRewardsFacet *ChangeRewardsFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChangeRewardsFacet.Contract.contract.Transact(opts, method, params...)
}

// ChangeRewardsAddress is a paid mutator transaction binding the contract method 0x8d240d8b.
//
// Solidity: function changeRewardsAddress(address _rewards) returns()
func (_ChangeRewardsFacet *ChangeRewardsFacetTransactor) ChangeRewardsAddress(opts *bind.TransactOpts, _rewards common.Address) (*types.Transaction, error) {
	return _ChangeRewardsFacet.contract.Transact(opts, "changeRewardsAddress", _rewards)
}

// ChangeRewardsAddress is a paid mutator transaction binding the contract method 0x8d240d8b.
//
// Solidity: function changeRewardsAddress(address _rewards) returns()
func (_ChangeRewardsFacet *ChangeRewardsFacetSession) ChangeRewardsAddress(_rewards common.Address) (*types.Transaction, error) {
	return _ChangeRewardsFacet.Contract.ChangeRewardsAddress(&_ChangeRewardsFacet.TransactOpts, _rewards)
}

// ChangeRewardsAddress is a paid mutator transaction binding the contract method 0x8d240d8b.
//
// Solidity: function changeRewardsAddress(address _rewards) returns()
func (_ChangeRewardsFacet *ChangeRewardsFacetTransactorSession) ChangeRewardsAddress(_rewards common.Address) (*types.Transaction, error) {
	return _ChangeRewardsFacet.Contract.ChangeRewardsAddress(&_ChangeRewardsFacet.TransactOpts, _rewards)
}
