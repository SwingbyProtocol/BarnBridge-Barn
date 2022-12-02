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

// INodeRewardsMetaData contains all meta data concerning the INodeRewards contract.
var INodeRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"totalNodeStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// INodeRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use INodeRewardsMetaData.ABI instead.
var INodeRewardsABI = INodeRewardsMetaData.ABI

// INodeRewards is an auto generated Go binding around an Ethereum contract.
type INodeRewards struct {
	INodeRewardsCaller     // Read-only binding to the contract
	INodeRewardsTransactor // Write-only binding to the contract
	INodeRewardsFilterer   // Log filterer for contract events
}

// INodeRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type INodeRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INodeRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INodeRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INodeRewardsSession struct {
	Contract     *INodeRewards     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INodeRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INodeRewardsCallerSession struct {
	Contract *INodeRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// INodeRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INodeRewardsTransactorSession struct {
	Contract     *INodeRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// INodeRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type INodeRewardsRaw struct {
	Contract *INodeRewards // Generic contract binding to access the raw methods on
}

// INodeRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INodeRewardsCallerRaw struct {
	Contract *INodeRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// INodeRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INodeRewardsTransactorRaw struct {
	Contract *INodeRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINodeRewards creates a new instance of INodeRewards, bound to a specific deployed contract.
func NewINodeRewards(address common.Address, backend bind.ContractBackend) (*INodeRewards, error) {
	contract, err := bindINodeRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INodeRewards{INodeRewardsCaller: INodeRewardsCaller{contract: contract}, INodeRewardsTransactor: INodeRewardsTransactor{contract: contract}, INodeRewardsFilterer: INodeRewardsFilterer{contract: contract}}, nil
}

// NewINodeRewardsCaller creates a new read-only instance of INodeRewards, bound to a specific deployed contract.
func NewINodeRewardsCaller(address common.Address, caller bind.ContractCaller) (*INodeRewardsCaller, error) {
	contract, err := bindINodeRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INodeRewardsCaller{contract: contract}, nil
}

// NewINodeRewardsTransactor creates a new write-only instance of INodeRewards, bound to a specific deployed contract.
func NewINodeRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*INodeRewardsTransactor, error) {
	contract, err := bindINodeRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INodeRewardsTransactor{contract: contract}, nil
}

// NewINodeRewardsFilterer creates a new log filterer instance of INodeRewards, bound to a specific deployed contract.
func NewINodeRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*INodeRewardsFilterer, error) {
	contract, err := bindINodeRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INodeRewardsFilterer{contract: contract}, nil
}

// bindINodeRewards binds a generic wrapper to an already deployed contract.
func bindINodeRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(INodeRewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INodeRewards *INodeRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INodeRewards.Contract.INodeRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INodeRewards *INodeRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INodeRewards.Contract.INodeRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INodeRewards *INodeRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INodeRewards.Contract.INodeRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INodeRewards *INodeRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INodeRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INodeRewards *INodeRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INodeRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INodeRewards *INodeRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INodeRewards.Contract.contract.Transact(opts, method, params...)
}

// TotalNodeStaked is a free data retrieval call binding the contract method 0x8ca0a7e5.
//
// Solidity: function totalNodeStaked() view returns(uint256)
func (_INodeRewards *INodeRewardsCaller) TotalNodeStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INodeRewards.contract.Call(opts, &out, "totalNodeStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNodeStaked is a free data retrieval call binding the contract method 0x8ca0a7e5.
//
// Solidity: function totalNodeStaked() view returns(uint256)
func (_INodeRewards *INodeRewardsSession) TotalNodeStaked() (*big.Int, error) {
	return _INodeRewards.Contract.TotalNodeStaked(&_INodeRewards.CallOpts)
}

// TotalNodeStaked is a free data retrieval call binding the contract method 0x8ca0a7e5.
//
// Solidity: function totalNodeStaked() view returns(uint256)
func (_INodeRewards *INodeRewardsCallerSession) TotalNodeStaked() (*big.Int, error) {
	return _INodeRewards.Contract.TotalNodeStaked(&_INodeRewards.CallOpts)
}
