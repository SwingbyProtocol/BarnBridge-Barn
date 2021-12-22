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

// IRewardsMetaData contains all meta data concerning the IRewards contract.
var IRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"registerUserAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"}],\"name\":\"setNewAPR\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use IRewardsMetaData.ABI instead.
var IRewardsABI = IRewardsMetaData.ABI

// IRewards is an auto generated Go binding around an Ethereum contract.
type IRewards struct {
	IRewardsCaller     // Read-only binding to the contract
	IRewardsTransactor // Write-only binding to the contract
	IRewardsFilterer   // Log filterer for contract events
}

// IRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRewardsSession struct {
	Contract     *IRewards         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRewardsCallerSession struct {
	Contract *IRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRewardsTransactorSession struct {
	Contract     *IRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRewardsRaw struct {
	Contract *IRewards // Generic contract binding to access the raw methods on
}

// IRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRewardsCallerRaw struct {
	Contract *IRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// IRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRewardsTransactorRaw struct {
	Contract *IRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRewards creates a new instance of IRewards, bound to a specific deployed contract.
func NewIRewards(address common.Address, backend bind.ContractBackend) (*IRewards, error) {
	contract, err := bindIRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRewards{IRewardsCaller: IRewardsCaller{contract: contract}, IRewardsTransactor: IRewardsTransactor{contract: contract}, IRewardsFilterer: IRewardsFilterer{contract: contract}}, nil
}

// NewIRewardsCaller creates a new read-only instance of IRewards, bound to a specific deployed contract.
func NewIRewardsCaller(address common.Address, caller bind.ContractCaller) (*IRewardsCaller, error) {
	contract, err := bindIRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardsCaller{contract: contract}, nil
}

// NewIRewardsTransactor creates a new write-only instance of IRewards, bound to a specific deployed contract.
func NewIRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*IRewardsTransactor, error) {
	contract, err := bindIRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardsTransactor{contract: contract}, nil
}

// NewIRewardsFilterer creates a new log filterer instance of IRewards, bound to a specific deployed contract.
func NewIRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*IRewardsFilterer, error) {
	contract, err := bindIRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRewardsFilterer{contract: contract}, nil
}

// bindIRewards binds a generic wrapper to an already deployed contract.
func bindIRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewards *IRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewards.Contract.IRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewards *IRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewards.Contract.IRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewards *IRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewards.Contract.IRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewards *IRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewards *IRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewards *IRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewards.Contract.contract.Transact(opts, method, params...)
}

// RegisterUserAction is a paid mutator transaction binding the contract method 0x66a7d821.
//
// Solidity: function registerUserAction(address user) returns()
func (_IRewards *IRewardsTransactor) RegisterUserAction(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _IRewards.contract.Transact(opts, "registerUserAction", user)
}

// RegisterUserAction is a paid mutator transaction binding the contract method 0x66a7d821.
//
// Solidity: function registerUserAction(address user) returns()
func (_IRewards *IRewardsSession) RegisterUserAction(user common.Address) (*types.Transaction, error) {
	return _IRewards.Contract.RegisterUserAction(&_IRewards.TransactOpts, user)
}

// RegisterUserAction is a paid mutator transaction binding the contract method 0x66a7d821.
//
// Solidity: function registerUserAction(address user) returns()
func (_IRewards *IRewardsTransactorSession) RegisterUserAction(user common.Address) (*types.Transaction, error) {
	return _IRewards.Contract.RegisterUserAction(&_IRewards.TransactOpts, user)
}

// SetNewAPR is a paid mutator transaction binding the contract method 0x5168e544.
//
// Solidity: function setNewAPR(uint256 _apr) returns()
func (_IRewards *IRewardsTransactor) SetNewAPR(opts *bind.TransactOpts, _apr *big.Int) (*types.Transaction, error) {
	return _IRewards.contract.Transact(opts, "setNewAPR", _apr)
}

// SetNewAPR is a paid mutator transaction binding the contract method 0x5168e544.
//
// Solidity: function setNewAPR(uint256 _apr) returns()
func (_IRewards *IRewardsSession) SetNewAPR(_apr *big.Int) (*types.Transaction, error) {
	return _IRewards.Contract.SetNewAPR(&_IRewards.TransactOpts, _apr)
}

// SetNewAPR is a paid mutator transaction binding the contract method 0x5168e544.
//
// Solidity: function setNewAPR(uint256 _apr) returns()
func (_IRewards *IRewardsTransactorSession) SetNewAPR(_apr *big.Int) (*types.Transaction, error) {
	return _IRewards.Contract.SetNewAPR(&_IRewards.TransactOpts, _apr)
}
