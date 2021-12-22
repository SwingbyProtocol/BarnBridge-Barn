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

// BarnMetaData contains all meta data concerning the Barn contract.
var BarnMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BarnABI is the input ABI used to generate the binding from.
// Deprecated: Use BarnMetaData.ABI instead.
var BarnABI = BarnMetaData.ABI

// Barn is an auto generated Go binding around an Ethereum contract.
type Barn struct {
	BarnCaller     // Read-only binding to the contract
	BarnTransactor // Write-only binding to the contract
	BarnFilterer   // Log filterer for contract events
}

// BarnCaller is an auto generated read-only Go binding around an Ethereum contract.
type BarnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BarnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BarnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BarnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BarnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BarnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BarnSession struct {
	Contract     *Barn             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BarnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BarnCallerSession struct {
	Contract *BarnCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BarnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BarnTransactorSession struct {
	Contract     *BarnTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BarnRaw is an auto generated low-level Go binding around an Ethereum contract.
type BarnRaw struct {
	Contract *Barn // Generic contract binding to access the raw methods on
}

// BarnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BarnCallerRaw struct {
	Contract *BarnCaller // Generic read-only contract binding to access the raw methods on
}

// BarnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BarnTransactorRaw struct {
	Contract *BarnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBarn creates a new instance of Barn, bound to a specific deployed contract.
func NewBarn(address common.Address, backend bind.ContractBackend) (*Barn, error) {
	contract, err := bindBarn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Barn{BarnCaller: BarnCaller{contract: contract}, BarnTransactor: BarnTransactor{contract: contract}, BarnFilterer: BarnFilterer{contract: contract}}, nil
}

// NewBarnCaller creates a new read-only instance of Barn, bound to a specific deployed contract.
func NewBarnCaller(address common.Address, caller bind.ContractCaller) (*BarnCaller, error) {
	contract, err := bindBarn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BarnCaller{contract: contract}, nil
}

// NewBarnTransactor creates a new write-only instance of Barn, bound to a specific deployed contract.
func NewBarnTransactor(address common.Address, transactor bind.ContractTransactor) (*BarnTransactor, error) {
	contract, err := bindBarn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BarnTransactor{contract: contract}, nil
}

// NewBarnFilterer creates a new log filterer instance of Barn, bound to a specific deployed contract.
func NewBarnFilterer(address common.Address, filterer bind.ContractFilterer) (*BarnFilterer, error) {
	contract, err := bindBarn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BarnFilterer{contract: contract}, nil
}

// bindBarn binds a generic wrapper to an already deployed contract.
func bindBarn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BarnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Barn *BarnRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Barn.Contract.BarnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Barn *BarnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Barn.Contract.BarnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Barn *BarnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Barn.Contract.BarnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Barn *BarnCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Barn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Barn *BarnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Barn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Barn *BarnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Barn.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Barn *BarnTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Barn.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Barn *BarnSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Barn.Contract.Fallback(&_Barn.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Barn *BarnTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Barn.Contract.Fallback(&_Barn.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Barn *BarnTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Barn.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Barn *BarnSession) Receive() (*types.Transaction, error) {
	return _Barn.Contract.Receive(&_Barn.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Barn *BarnTransactorSession) Receive() (*types.Transaction, error) {
	return _Barn.Contract.Receive(&_Barn.TransactOpts)
}
