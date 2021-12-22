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

// IDiamondCutFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondCutFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// DiamondCutFacetMetaData contains all meta data concerning the DiamondCutFacet contract.
var DiamondCutFacetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamondCut.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondCut.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DiamondCutFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondCutFacetMetaData.ABI instead.
var DiamondCutFacetABI = DiamondCutFacetMetaData.ABI

// DiamondCutFacet is an auto generated Go binding around an Ethereum contract.
type DiamondCutFacet struct {
	DiamondCutFacetCaller     // Read-only binding to the contract
	DiamondCutFacetTransactor // Write-only binding to the contract
	DiamondCutFacetFilterer   // Log filterer for contract events
}

// DiamondCutFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondCutFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondCutFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondCutFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondCutFacetSession struct {
	Contract     *DiamondCutFacet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiamondCutFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondCutFacetCallerSession struct {
	Contract *DiamondCutFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DiamondCutFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondCutFacetTransactorSession struct {
	Contract     *DiamondCutFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DiamondCutFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondCutFacetRaw struct {
	Contract *DiamondCutFacet // Generic contract binding to access the raw methods on
}

// DiamondCutFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondCutFacetCallerRaw struct {
	Contract *DiamondCutFacetCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondCutFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondCutFacetTransactorRaw struct {
	Contract *DiamondCutFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondCutFacet creates a new instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacet(address common.Address, backend bind.ContractBackend) (*DiamondCutFacet, error) {
	contract, err := bindDiamondCutFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacet{DiamondCutFacetCaller: DiamondCutFacetCaller{contract: contract}, DiamondCutFacetTransactor: DiamondCutFacetTransactor{contract: contract}, DiamondCutFacetFilterer: DiamondCutFacetFilterer{contract: contract}}, nil
}

// NewDiamondCutFacetCaller creates a new read-only instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacetCaller(address common.Address, caller bind.ContractCaller) (*DiamondCutFacetCaller, error) {
	contract, err := bindDiamondCutFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetCaller{contract: contract}, nil
}

// NewDiamondCutFacetTransactor creates a new write-only instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondCutFacetTransactor, error) {
	contract, err := bindDiamondCutFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetTransactor{contract: contract}, nil
}

// NewDiamondCutFacetFilterer creates a new log filterer instance of DiamondCutFacet, bound to a specific deployed contract.
func NewDiamondCutFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondCutFacetFilterer, error) {
	contract, err := bindDiamondCutFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetFilterer{contract: contract}, nil
}

// bindDiamondCutFacet binds a generic wrapper to an already deployed contract.
func bindDiamondCutFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiamondCutFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondCutFacet *DiamondCutFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondCutFacet.Contract.DiamondCutFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondCutFacet *DiamondCutFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCutFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondCutFacet *DiamondCutFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCutFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondCutFacet *DiamondCutFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondCutFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondCutFacet *DiamondCutFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondCutFacet *DiamondCutFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.contract.Transact(opts, method, params...)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCutFacet *DiamondCutFacetTransactor) DiamondCut(opts *bind.TransactOpts, _diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCutFacet.contract.Transact(opts, "diamondCut", _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCutFacet *DiamondCutFacetSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCut(&_DiamondCutFacet.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCutFacet *DiamondCutFacetTransactorSession) DiamondCut(_diamondCut []IDiamondCutFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCutFacet.Contract.DiamondCut(&_DiamondCutFacet.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCutFacetDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the DiamondCutFacet contract.
type DiamondCutFacetDiamondCutIterator struct {
	Event *DiamondCutFacetDiamondCut // Event containing the contract specifics and raw log

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
func (it *DiamondCutFacetDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondCutFacetDiamondCut)
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
		it.Event = new(DiamondCutFacetDiamondCut)
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
func (it *DiamondCutFacetDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondCutFacetDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondCutFacetDiamondCut represents a DiamondCut event raised by the DiamondCutFacet contract.
type DiamondCutFacetDiamondCut struct {
	DiamondCut []IDiamondCutFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*DiamondCutFacetDiamondCutIterator, error) {

	logs, sub, err := _DiamondCutFacet.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &DiamondCutFacetDiamondCutIterator{contract: _DiamondCutFacet.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *DiamondCutFacetDiamondCut) (event.Subscription, error) {

	logs, sub, err := _DiamondCutFacet.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondCutFacetDiamondCut)
				if err := _DiamondCutFacet.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCutFacet *DiamondCutFacetFilterer) ParseDiamondCut(log types.Log) (*DiamondCutFacetDiamondCut, error) {
	event := new(DiamondCutFacetDiamondCut)
	if err := _DiamondCutFacet.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
