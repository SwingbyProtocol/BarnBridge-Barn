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

// LibBarnStorageStake is an auto generated low-level Go binding around an user-defined struct.
type LibBarnStorageStake struct {
	Timestamp       *big.Int
	Amount          *big.Int
	ExpiryTimestamp *big.Int
	DelegatedTo     common.Address
}

// IBarnMetaData contains all meta data concerning the IBarn contract.
var IBarnMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"balanceAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondCirculatingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"bondStakedAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"delegatedPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"delegatedPowerAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"lockCreatorBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"multiplierAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"stakeAtTs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"}],\"internalType\":\"structLibBarnStorage.Stake\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userDelegatedTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userLockedUntil\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"votingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"votingPowerAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBarnABI is the input ABI used to generate the binding from.
// Deprecated: Use IBarnMetaData.ABI instead.
var IBarnABI = IBarnMetaData.ABI

// IBarn is an auto generated Go binding around an Ethereum contract.
type IBarn struct {
	IBarnCaller     // Read-only binding to the contract
	IBarnTransactor // Write-only binding to the contract
	IBarnFilterer   // Log filterer for contract events
}

// IBarnCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBarnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBarnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBarnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBarnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBarnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBarnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBarnSession struct {
	Contract     *IBarn            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBarnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBarnCallerSession struct {
	Contract *IBarnCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IBarnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBarnTransactorSession struct {
	Contract     *IBarnTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBarnRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBarnRaw struct {
	Contract *IBarn // Generic contract binding to access the raw methods on
}

// IBarnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBarnCallerRaw struct {
	Contract *IBarnCaller // Generic read-only contract binding to access the raw methods on
}

// IBarnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBarnTransactorRaw struct {
	Contract *IBarnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBarn creates a new instance of IBarn, bound to a specific deployed contract.
func NewIBarn(address common.Address, backend bind.ContractBackend) (*IBarn, error) {
	contract, err := bindIBarn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBarn{IBarnCaller: IBarnCaller{contract: contract}, IBarnTransactor: IBarnTransactor{contract: contract}, IBarnFilterer: IBarnFilterer{contract: contract}}, nil
}

// NewIBarnCaller creates a new read-only instance of IBarn, bound to a specific deployed contract.
func NewIBarnCaller(address common.Address, caller bind.ContractCaller) (*IBarnCaller, error) {
	contract, err := bindIBarn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBarnCaller{contract: contract}, nil
}

// NewIBarnTransactor creates a new write-only instance of IBarn, bound to a specific deployed contract.
func NewIBarnTransactor(address common.Address, transactor bind.ContractTransactor) (*IBarnTransactor, error) {
	contract, err := bindIBarn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBarnTransactor{contract: contract}, nil
}

// NewIBarnFilterer creates a new log filterer instance of IBarn, bound to a specific deployed contract.
func NewIBarnFilterer(address common.Address, filterer bind.ContractFilterer) (*IBarnFilterer, error) {
	contract, err := bindIBarn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBarnFilterer{contract: contract}, nil
}

// bindIBarn binds a generic wrapper to an already deployed contract.
func bindIBarn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBarnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBarn *IBarnRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBarn.Contract.IBarnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBarn *IBarnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBarn.Contract.IBarnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBarn *IBarnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBarn.Contract.IBarnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBarn *IBarnCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBarn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBarn *IBarnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBarn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBarn *IBarnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBarn.Contract.contract.Transact(opts, method, params...)
}

// BalanceAtTs is a free data retrieval call binding the contract method 0x417edd4d.
//
// Solidity: function balanceAtTs(address user, uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnCaller) BalanceAtTs(opts *bind.CallOpts, user common.Address, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "balanceAtTs", user, timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceAtTs is a free data retrieval call binding the contract method 0x417edd4d.
//
// Solidity: function balanceAtTs(address user, uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnSession) BalanceAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _IBarn.Contract.BalanceAtTs(&_IBarn.CallOpts, user, timestamp)
}

// BalanceAtTs is a free data retrieval call binding the contract method 0x417edd4d.
//
// Solidity: function balanceAtTs(address user, uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnCallerSession) BalanceAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _IBarn.Contract.BalanceAtTs(&_IBarn.CallOpts, user, timestamp)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_IBarn *IBarnCaller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_IBarn *IBarnSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _IBarn.Contract.BalanceOf(&_IBarn.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_IBarn *IBarnCallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _IBarn.Contract.BalanceOf(&_IBarn.CallOpts, user)
}

// BondCirculatingSupply is a free data retrieval call binding the contract method 0x2082b4d1.
//
// Solidity: function bondCirculatingSupply() view returns(uint256)
func (_IBarn *IBarnCaller) BondCirculatingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "bondCirculatingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondCirculatingSupply is a free data retrieval call binding the contract method 0x2082b4d1.
//
// Solidity: function bondCirculatingSupply() view returns(uint256)
func (_IBarn *IBarnSession) BondCirculatingSupply() (*big.Int, error) {
	return _IBarn.Contract.BondCirculatingSupply(&_IBarn.CallOpts)
}

// BondCirculatingSupply is a free data retrieval call binding the contract method 0x2082b4d1.
//
// Solidity: function bondCirculatingSupply() view returns(uint256)
func (_IBarn *IBarnCallerSession) BondCirculatingSupply() (*big.Int, error) {
	return _IBarn.Contract.BondCirculatingSupply(&_IBarn.CallOpts)
}

// BondStaked is a free data retrieval call binding the contract method 0xc2077e81.
//
// Solidity: function bondStaked() view returns(uint256)
func (_IBarn *IBarnCaller) BondStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "bondStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondStaked is a free data retrieval call binding the contract method 0xc2077e81.
//
// Solidity: function bondStaked() view returns(uint256)
func (_IBarn *IBarnSession) BondStaked() (*big.Int, error) {
	return _IBarn.Contract.BondStaked(&_IBarn.CallOpts)
}

// BondStaked is a free data retrieval call binding the contract method 0xc2077e81.
//
// Solidity: function bondStaked() view returns(uint256)
func (_IBarn *IBarnCallerSession) BondStaked() (*big.Int, error) {
	return _IBarn.Contract.BondStaked(&_IBarn.CallOpts)
}

// BondStakedAtTs is a free data retrieval call binding the contract method 0xf77f962f.
//
// Solidity: function bondStakedAtTs(uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnCaller) BondStakedAtTs(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "bondStakedAtTs", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondStakedAtTs is a free data retrieval call binding the contract method 0xf77f962f.
//
// Solidity: function bondStakedAtTs(uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnSession) BondStakedAtTs(timestamp *big.Int) (*big.Int, error) {
	return _IBarn.Contract.BondStakedAtTs(&_IBarn.CallOpts, timestamp)
}

// BondStakedAtTs is a free data retrieval call binding the contract method 0xf77f962f.
//
// Solidity: function bondStakedAtTs(uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnCallerSession) BondStakedAtTs(timestamp *big.Int) (*big.Int, error) {
	return _IBarn.Contract.BondStakedAtTs(&_IBarn.CallOpts, timestamp)
}

// DelegatedPower is a free data retrieval call binding the contract method 0x169df064.
//
// Solidity: function delegatedPower(address user) view returns(uint256)
func (_IBarn *IBarnCaller) DelegatedPower(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "delegatedPower", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedPower is a free data retrieval call binding the contract method 0x169df064.
//
// Solidity: function delegatedPower(address user) view returns(uint256)
func (_IBarn *IBarnSession) DelegatedPower(user common.Address) (*big.Int, error) {
	return _IBarn.Contract.DelegatedPower(&_IBarn.CallOpts, user)
}

// DelegatedPower is a free data retrieval call binding the contract method 0x169df064.
//
// Solidity: function delegatedPower(address user) view returns(uint256)
func (_IBarn *IBarnCallerSession) DelegatedPower(user common.Address) (*big.Int, error) {
	return _IBarn.Contract.DelegatedPower(&_IBarn.CallOpts, user)
}

// DelegatedPowerAtTs is a free data retrieval call binding the contract method 0xd265a115.
//
// Solidity: function delegatedPowerAtTs(address user, uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnCaller) DelegatedPowerAtTs(opts *bind.CallOpts, user common.Address, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "delegatedPowerAtTs", user, timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedPowerAtTs is a free data retrieval call binding the contract method 0xd265a115.
//
// Solidity: function delegatedPowerAtTs(address user, uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnSession) DelegatedPowerAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _IBarn.Contract.DelegatedPowerAtTs(&_IBarn.CallOpts, user, timestamp)
}

// DelegatedPowerAtTs is a free data retrieval call binding the contract method 0xd265a115.
//
// Solidity: function delegatedPowerAtTs(address user, uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnCallerSession) DelegatedPowerAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _IBarn.Contract.DelegatedPowerAtTs(&_IBarn.CallOpts, user, timestamp)
}

// MultiplierAtTs is a free data retrieval call binding the contract method 0x7a141096.
//
// Solidity: function multiplierAtTs(address user, uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnCaller) MultiplierAtTs(opts *bind.CallOpts, user common.Address, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "multiplierAtTs", user, timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MultiplierAtTs is a free data retrieval call binding the contract method 0x7a141096.
//
// Solidity: function multiplierAtTs(address user, uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnSession) MultiplierAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _IBarn.Contract.MultiplierAtTs(&_IBarn.CallOpts, user, timestamp)
}

// MultiplierAtTs is a free data retrieval call binding the contract method 0x7a141096.
//
// Solidity: function multiplierAtTs(address user, uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnCallerSession) MultiplierAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _IBarn.Contract.MultiplierAtTs(&_IBarn.CallOpts, user, timestamp)
}

// StakeAtTs is a free data retrieval call binding the contract method 0x18ab6a3c.
//
// Solidity: function stakeAtTs(address user, uint256 timestamp) view returns((uint256,uint256,uint256,address))
func (_IBarn *IBarnCaller) StakeAtTs(opts *bind.CallOpts, user common.Address, timestamp *big.Int) (LibBarnStorageStake, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "stakeAtTs", user, timestamp)

	if err != nil {
		return *new(LibBarnStorageStake), err
	}

	out0 := *abi.ConvertType(out[0], new(LibBarnStorageStake)).(*LibBarnStorageStake)

	return out0, err

}

// StakeAtTs is a free data retrieval call binding the contract method 0x18ab6a3c.
//
// Solidity: function stakeAtTs(address user, uint256 timestamp) view returns((uint256,uint256,uint256,address))
func (_IBarn *IBarnSession) StakeAtTs(user common.Address, timestamp *big.Int) (LibBarnStorageStake, error) {
	return _IBarn.Contract.StakeAtTs(&_IBarn.CallOpts, user, timestamp)
}

// StakeAtTs is a free data retrieval call binding the contract method 0x18ab6a3c.
//
// Solidity: function stakeAtTs(address user, uint256 timestamp) view returns((uint256,uint256,uint256,address))
func (_IBarn *IBarnCallerSession) StakeAtTs(user common.Address, timestamp *big.Int) (LibBarnStorageStake, error) {
	return _IBarn.Contract.StakeAtTs(&_IBarn.CallOpts, user, timestamp)
}

// UserDelegatedTo is a free data retrieval call binding the contract method 0xbef624d8.
//
// Solidity: function userDelegatedTo(address user) view returns(address)
func (_IBarn *IBarnCaller) UserDelegatedTo(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "userDelegatedTo", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserDelegatedTo is a free data retrieval call binding the contract method 0xbef624d8.
//
// Solidity: function userDelegatedTo(address user) view returns(address)
func (_IBarn *IBarnSession) UserDelegatedTo(user common.Address) (common.Address, error) {
	return _IBarn.Contract.UserDelegatedTo(&_IBarn.CallOpts, user)
}

// UserDelegatedTo is a free data retrieval call binding the contract method 0xbef624d8.
//
// Solidity: function userDelegatedTo(address user) view returns(address)
func (_IBarn *IBarnCallerSession) UserDelegatedTo(user common.Address) (common.Address, error) {
	return _IBarn.Contract.UserDelegatedTo(&_IBarn.CallOpts, user)
}

// UserLockedUntil is a free data retrieval call binding the contract method 0xbf0ae48c.
//
// Solidity: function userLockedUntil(address user) view returns(uint256)
func (_IBarn *IBarnCaller) UserLockedUntil(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "userLockedUntil", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLockedUntil is a free data retrieval call binding the contract method 0xbf0ae48c.
//
// Solidity: function userLockedUntil(address user) view returns(uint256)
func (_IBarn *IBarnSession) UserLockedUntil(user common.Address) (*big.Int, error) {
	return _IBarn.Contract.UserLockedUntil(&_IBarn.CallOpts, user)
}

// UserLockedUntil is a free data retrieval call binding the contract method 0xbf0ae48c.
//
// Solidity: function userLockedUntil(address user) view returns(uint256)
func (_IBarn *IBarnCallerSession) UserLockedUntil(user common.Address) (*big.Int, error) {
	return _IBarn.Contract.UserLockedUntil(&_IBarn.CallOpts, user)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address user) view returns(uint256)
func (_IBarn *IBarnCaller) VotingPower(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "votingPower", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address user) view returns(uint256)
func (_IBarn *IBarnSession) VotingPower(user common.Address) (*big.Int, error) {
	return _IBarn.Contract.VotingPower(&_IBarn.CallOpts, user)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address user) view returns(uint256)
func (_IBarn *IBarnCallerSession) VotingPower(user common.Address) (*big.Int, error) {
	return _IBarn.Contract.VotingPower(&_IBarn.CallOpts, user)
}

// VotingPowerAtTs is a free data retrieval call binding the contract method 0xcbf8eda9.
//
// Solidity: function votingPowerAtTs(address user, uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnCaller) VotingPowerAtTs(opts *bind.CallOpts, user common.Address, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IBarn.contract.Call(opts, &out, "votingPowerAtTs", user, timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPowerAtTs is a free data retrieval call binding the contract method 0xcbf8eda9.
//
// Solidity: function votingPowerAtTs(address user, uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnSession) VotingPowerAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _IBarn.Contract.VotingPowerAtTs(&_IBarn.CallOpts, user, timestamp)
}

// VotingPowerAtTs is a free data retrieval call binding the contract method 0xcbf8eda9.
//
// Solidity: function votingPowerAtTs(address user, uint256 timestamp) view returns(uint256)
func (_IBarn *IBarnCallerSession) VotingPowerAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _IBarn.Contract.VotingPowerAtTs(&_IBarn.CallOpts, user, timestamp)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_IBarn *IBarnTransactor) Delegate(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IBarn.contract.Transact(opts, "delegate", to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_IBarn *IBarnSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _IBarn.Contract.Delegate(&_IBarn.TransactOpts, to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_IBarn *IBarnTransactorSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _IBarn.Contract.Delegate(&_IBarn.TransactOpts, to)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IBarn *IBarnTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IBarn.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IBarn *IBarnSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _IBarn.Contract.Deposit(&_IBarn.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_IBarn *IBarnTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _IBarn.Contract.Deposit(&_IBarn.TransactOpts, amount)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 timestamp) returns()
func (_IBarn *IBarnTransactor) Lock(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _IBarn.contract.Transact(opts, "lock", timestamp)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 timestamp) returns()
func (_IBarn *IBarnSession) Lock(timestamp *big.Int) (*types.Transaction, error) {
	return _IBarn.Contract.Lock(&_IBarn.TransactOpts, timestamp)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 timestamp) returns()
func (_IBarn *IBarnTransactorSession) Lock(timestamp *big.Int) (*types.Transaction, error) {
	return _IBarn.Contract.Lock(&_IBarn.TransactOpts, timestamp)
}

// LockCreatorBalance is a paid mutator transaction binding the contract method 0x71ef7663.
//
// Solidity: function lockCreatorBalance(address user, uint256 timestamp) returns()
func (_IBarn *IBarnTransactor) LockCreatorBalance(opts *bind.TransactOpts, user common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _IBarn.contract.Transact(opts, "lockCreatorBalance", user, timestamp)
}

// LockCreatorBalance is a paid mutator transaction binding the contract method 0x71ef7663.
//
// Solidity: function lockCreatorBalance(address user, uint256 timestamp) returns()
func (_IBarn *IBarnSession) LockCreatorBalance(user common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _IBarn.Contract.LockCreatorBalance(&_IBarn.TransactOpts, user, timestamp)
}

// LockCreatorBalance is a paid mutator transaction binding the contract method 0x71ef7663.
//
// Solidity: function lockCreatorBalance(address user, uint256 timestamp) returns()
func (_IBarn *IBarnTransactorSession) LockCreatorBalance(user common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _IBarn.Contract.LockCreatorBalance(&_IBarn.TransactOpts, user, timestamp)
}

// StopDelegate is a paid mutator transaction binding the contract method 0x6f121578.
//
// Solidity: function stopDelegate() returns()
func (_IBarn *IBarnTransactor) StopDelegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBarn.contract.Transact(opts, "stopDelegate")
}

// StopDelegate is a paid mutator transaction binding the contract method 0x6f121578.
//
// Solidity: function stopDelegate() returns()
func (_IBarn *IBarnSession) StopDelegate() (*types.Transaction, error) {
	return _IBarn.Contract.StopDelegate(&_IBarn.TransactOpts)
}

// StopDelegate is a paid mutator transaction binding the contract method 0x6f121578.
//
// Solidity: function stopDelegate() returns()
func (_IBarn *IBarnTransactorSession) StopDelegate() (*types.Transaction, error) {
	return _IBarn.Contract.StopDelegate(&_IBarn.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_IBarn *IBarnTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IBarn.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_IBarn *IBarnSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _IBarn.Contract.Withdraw(&_IBarn.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_IBarn *IBarnTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _IBarn.Contract.Withdraw(&_IBarn.TransactOpts, amount)
}
