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

// RewardsMetaData contains all meta data concerning the Rewards contract.
var RewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swingby\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_barn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ackFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceBefore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"barn\",\"outputs\":[{\"internalType\":\"contractIBarn\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPullTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pullFeature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"registerUserAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_barn\",\"type\":\"address\"}],\"name\":\"setBarn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"}],\"name\":\"setNewAPR\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTs\",\"type\":\"uint256\"}],\"name\":\"setupPullToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardsMetaData.ABI instead.
var RewardsABI = RewardsMetaData.ABI

// Rewards is an auto generated Go binding around an Ethereum contract.
type Rewards struct {
	RewardsCaller     // Read-only binding to the contract
	RewardsTransactor // Write-only binding to the contract
	RewardsFilterer   // Log filterer for contract events
}

// RewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardsSession struct {
	Contract     *Rewards          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardsCallerSession struct {
	Contract *RewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardsTransactorSession struct {
	Contract     *RewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardsRaw struct {
	Contract *Rewards // Generic contract binding to access the raw methods on
}

// RewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardsCallerRaw struct {
	Contract *RewardsCaller // Generic read-only contract binding to access the raw methods on
}

// RewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardsTransactorRaw struct {
	Contract *RewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewards creates a new instance of Rewards, bound to a specific deployed contract.
func NewRewards(address common.Address, backend bind.ContractBackend) (*Rewards, error) {
	contract, err := bindRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rewards{RewardsCaller: RewardsCaller{contract: contract}, RewardsTransactor: RewardsTransactor{contract: contract}, RewardsFilterer: RewardsFilterer{contract: contract}}, nil
}

// NewRewardsCaller creates a new read-only instance of Rewards, bound to a specific deployed contract.
func NewRewardsCaller(address common.Address, caller bind.ContractCaller) (*RewardsCaller, error) {
	contract, err := bindRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsCaller{contract: contract}, nil
}

// NewRewardsTransactor creates a new write-only instance of Rewards, bound to a specific deployed contract.
func NewRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardsTransactor, error) {
	contract, err := bindRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsTransactor{contract: contract}, nil
}

// NewRewardsFilterer creates a new log filterer instance of Rewards, bound to a specific deployed contract.
func NewRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardsFilterer, error) {
	contract, err := bindRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardsFilterer{contract: contract}, nil
}

// bindRewards binds a generic wrapper to an already deployed contract.
func bindRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewards *RewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewards.Contract.RewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewards *RewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.Contract.RewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewards *RewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewards.Contract.RewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewards *RewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewards *RewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewards *RewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewards.Contract.contract.Transact(opts, method, params...)
}

// Apr is a free data retrieval call binding the contract method 0x57ded9c9.
//
// Solidity: function apr() view returns(uint256)
func (_Rewards *RewardsCaller) Apr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "apr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Apr is a free data retrieval call binding the contract method 0x57ded9c9.
//
// Solidity: function apr() view returns(uint256)
func (_Rewards *RewardsSession) Apr() (*big.Int, error) {
	return _Rewards.Contract.Apr(&_Rewards.CallOpts)
}

// Apr is a free data retrieval call binding the contract method 0x57ded9c9.
//
// Solidity: function apr() view returns(uint256)
func (_Rewards *RewardsCallerSession) Apr() (*big.Int, error) {
	return _Rewards.Contract.Apr(&_Rewards.CallOpts)
}

// BalanceBefore is a free data retrieval call binding the contract method 0x94b5798a.
//
// Solidity: function balanceBefore() view returns(uint256)
func (_Rewards *RewardsCaller) BalanceBefore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "balanceBefore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceBefore is a free data retrieval call binding the contract method 0x94b5798a.
//
// Solidity: function balanceBefore() view returns(uint256)
func (_Rewards *RewardsSession) BalanceBefore() (*big.Int, error) {
	return _Rewards.Contract.BalanceBefore(&_Rewards.CallOpts)
}

// BalanceBefore is a free data retrieval call binding the contract method 0x94b5798a.
//
// Solidity: function balanceBefore() view returns(uint256)
func (_Rewards *RewardsCallerSession) BalanceBefore() (*big.Int, error) {
	return _Rewards.Contract.BalanceBefore(&_Rewards.CallOpts)
}

// Barn is a free data retrieval call binding the contract method 0x194f480e.
//
// Solidity: function barn() view returns(address)
func (_Rewards *RewardsCaller) Barn(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "barn")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Barn is a free data retrieval call binding the contract method 0x194f480e.
//
// Solidity: function barn() view returns(address)
func (_Rewards *RewardsSession) Barn() (common.Address, error) {
	return _Rewards.Contract.Barn(&_Rewards.CallOpts)
}

// Barn is a free data retrieval call binding the contract method 0x194f480e.
//
// Solidity: function barn() view returns(address)
func (_Rewards *RewardsCallerSession) Barn() (common.Address, error) {
	return _Rewards.Contract.Barn(&_Rewards.CallOpts)
}

// CurrentMultiplier is a free data retrieval call binding the contract method 0x6fbaaa1e.
//
// Solidity: function currentMultiplier() view returns(uint256)
func (_Rewards *RewardsCaller) CurrentMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "currentMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentMultiplier is a free data retrieval call binding the contract method 0x6fbaaa1e.
//
// Solidity: function currentMultiplier() view returns(uint256)
func (_Rewards *RewardsSession) CurrentMultiplier() (*big.Int, error) {
	return _Rewards.Contract.CurrentMultiplier(&_Rewards.CallOpts)
}

// CurrentMultiplier is a free data retrieval call binding the contract method 0x6fbaaa1e.
//
// Solidity: function currentMultiplier() view returns(uint256)
func (_Rewards *RewardsCallerSession) CurrentMultiplier() (*big.Int, error) {
	return _Rewards.Contract.CurrentMultiplier(&_Rewards.CallOpts)
}

// Disabled is a free data retrieval call binding the contract method 0xee070805.
//
// Solidity: function disabled() view returns(bool)
func (_Rewards *RewardsCaller) Disabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "disabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Disabled is a free data retrieval call binding the contract method 0xee070805.
//
// Solidity: function disabled() view returns(bool)
func (_Rewards *RewardsSession) Disabled() (bool, error) {
	return _Rewards.Contract.Disabled(&_Rewards.CallOpts)
}

// Disabled is a free data retrieval call binding the contract method 0xee070805.
//
// Solidity: function disabled() view returns(bool)
func (_Rewards *RewardsCallerSession) Disabled() (bool, error) {
	return _Rewards.Contract.Disabled(&_Rewards.CallOpts)
}

// LastPullTs is a free data retrieval call binding the contract method 0x100223bb.
//
// Solidity: function lastPullTs() view returns(uint256)
func (_Rewards *RewardsCaller) LastPullTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "lastPullTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPullTs is a free data retrieval call binding the contract method 0x100223bb.
//
// Solidity: function lastPullTs() view returns(uint256)
func (_Rewards *RewardsSession) LastPullTs() (*big.Int, error) {
	return _Rewards.Contract.LastPullTs(&_Rewards.CallOpts)
}

// LastPullTs is a free data retrieval call binding the contract method 0x100223bb.
//
// Solidity: function lastPullTs() view returns(uint256)
func (_Rewards *RewardsCallerSession) LastPullTs() (*big.Int, error) {
	return _Rewards.Contract.LastPullTs(&_Rewards.CallOpts)
}

// Owed is a free data retrieval call binding the contract method 0xdf18e047.
//
// Solidity: function owed(address ) view returns(uint256)
func (_Rewards *RewardsCaller) Owed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "owed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Owed is a free data retrieval call binding the contract method 0xdf18e047.
//
// Solidity: function owed(address ) view returns(uint256)
func (_Rewards *RewardsSession) Owed(arg0 common.Address) (*big.Int, error) {
	return _Rewards.Contract.Owed(&_Rewards.CallOpts, arg0)
}

// Owed is a free data retrieval call binding the contract method 0xdf18e047.
//
// Solidity: function owed(address ) view returns(uint256)
func (_Rewards *RewardsCallerSession) Owed(arg0 common.Address) (*big.Int, error) {
	return _Rewards.Contract.Owed(&_Rewards.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewards *RewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewards *RewardsSession) Owner() (common.Address, error) {
	return _Rewards.Contract.Owner(&_Rewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewards *RewardsCallerSession) Owner() (common.Address, error) {
	return _Rewards.Contract.Owner(&_Rewards.CallOpts)
}

// PullFeature is a free data retrieval call binding the contract method 0x4831976c.
//
// Solidity: function pullFeature() view returns(address source, uint256 startTs, uint256 endTs, uint256 totalDuration)
func (_Rewards *RewardsCaller) PullFeature(opts *bind.CallOpts) (struct {
	Source        common.Address
	StartTs       *big.Int
	EndTs         *big.Int
	TotalDuration *big.Int
}, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "pullFeature")

	outstruct := new(struct {
		Source        common.Address
		StartTs       *big.Int
		EndTs         *big.Int
		TotalDuration *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Source = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StartTs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndTs = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalDuration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PullFeature is a free data retrieval call binding the contract method 0x4831976c.
//
// Solidity: function pullFeature() view returns(address source, uint256 startTs, uint256 endTs, uint256 totalDuration)
func (_Rewards *RewardsSession) PullFeature() (struct {
	Source        common.Address
	StartTs       *big.Int
	EndTs         *big.Int
	TotalDuration *big.Int
}, error) {
	return _Rewards.Contract.PullFeature(&_Rewards.CallOpts)
}

// PullFeature is a free data retrieval call binding the contract method 0x4831976c.
//
// Solidity: function pullFeature() view returns(address source, uint256 startTs, uint256 endTs, uint256 totalDuration)
func (_Rewards *RewardsCallerSession) PullFeature() (struct {
	Source        common.Address
	StartTs       *big.Int
	EndTs         *big.Int
	TotalDuration *big.Int
}, error) {
	return _Rewards.Contract.PullFeature(&_Rewards.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Rewards *RewardsCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Rewards *RewardsSession) RewardToken() (common.Address, error) {
	return _Rewards.Contract.RewardToken(&_Rewards.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Rewards *RewardsCallerSession) RewardToken() (common.Address, error) {
	return _Rewards.Contract.RewardToken(&_Rewards.CallOpts)
}

// UserMultiplier is a free data retrieval call binding the contract method 0xb1a03b6b.
//
// Solidity: function userMultiplier(address ) view returns(uint256)
func (_Rewards *RewardsCaller) UserMultiplier(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "userMultiplier", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMultiplier is a free data retrieval call binding the contract method 0xb1a03b6b.
//
// Solidity: function userMultiplier(address ) view returns(uint256)
func (_Rewards *RewardsSession) UserMultiplier(arg0 common.Address) (*big.Int, error) {
	return _Rewards.Contract.UserMultiplier(&_Rewards.CallOpts, arg0)
}

// UserMultiplier is a free data retrieval call binding the contract method 0xb1a03b6b.
//
// Solidity: function userMultiplier(address ) view returns(uint256)
func (_Rewards *RewardsCallerSession) UserMultiplier(arg0 common.Address) (*big.Int, error) {
	return _Rewards.Contract.UserMultiplier(&_Rewards.CallOpts, arg0)
}

// AckFunds is a paid mutator transaction binding the contract method 0xacfd9325.
//
// Solidity: function ackFunds() returns()
func (_Rewards *RewardsTransactor) AckFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "ackFunds")
}

// AckFunds is a paid mutator transaction binding the contract method 0xacfd9325.
//
// Solidity: function ackFunds() returns()
func (_Rewards *RewardsSession) AckFunds() (*types.Transaction, error) {
	return _Rewards.Contract.AckFunds(&_Rewards.TransactOpts)
}

// AckFunds is a paid mutator transaction binding the contract method 0xacfd9325.
//
// Solidity: function ackFunds() returns()
func (_Rewards *RewardsTransactorSession) AckFunds() (*types.Transaction, error) {
	return _Rewards.Contract.AckFunds(&_Rewards.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_Rewards *RewardsTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_Rewards *RewardsSession) Claim() (*types.Transaction, error) {
	return _Rewards.Contract.Claim(&_Rewards.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_Rewards *RewardsTransactorSession) Claim() (*types.Transaction, error) {
	return _Rewards.Contract.Claim(&_Rewards.TransactOpts)
}

// RegisterUserAction is a paid mutator transaction binding the contract method 0x66a7d821.
//
// Solidity: function registerUserAction(address user) returns()
func (_Rewards *RewardsTransactor) RegisterUserAction(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "registerUserAction", user)
}

// RegisterUserAction is a paid mutator transaction binding the contract method 0x66a7d821.
//
// Solidity: function registerUserAction(address user) returns()
func (_Rewards *RewardsSession) RegisterUserAction(user common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.RegisterUserAction(&_Rewards.TransactOpts, user)
}

// RegisterUserAction is a paid mutator transaction binding the contract method 0x66a7d821.
//
// Solidity: function registerUserAction(address user) returns()
func (_Rewards *RewardsTransactorSession) RegisterUserAction(user common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.RegisterUserAction(&_Rewards.TransactOpts, user)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewards *RewardsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewards *RewardsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewards.Contract.RenounceOwnership(&_Rewards.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewards *RewardsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewards.Contract.RenounceOwnership(&_Rewards.TransactOpts)
}

// SetBarn is a paid mutator transaction binding the contract method 0x3b342a85.
//
// Solidity: function setBarn(address _barn) returns()
func (_Rewards *RewardsTransactor) SetBarn(opts *bind.TransactOpts, _barn common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setBarn", _barn)
}

// SetBarn is a paid mutator transaction binding the contract method 0x3b342a85.
//
// Solidity: function setBarn(address _barn) returns()
func (_Rewards *RewardsSession) SetBarn(_barn common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetBarn(&_Rewards.TransactOpts, _barn)
}

// SetBarn is a paid mutator transaction binding the contract method 0x3b342a85.
//
// Solidity: function setBarn(address _barn) returns()
func (_Rewards *RewardsTransactorSession) SetBarn(_barn common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetBarn(&_Rewards.TransactOpts, _barn)
}

// SetNewAPR is a paid mutator transaction binding the contract method 0x5168e544.
//
// Solidity: function setNewAPR(uint256 _apr) returns()
func (_Rewards *RewardsTransactor) SetNewAPR(opts *bind.TransactOpts, _apr *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setNewAPR", _apr)
}

// SetNewAPR is a paid mutator transaction binding the contract method 0x5168e544.
//
// Solidity: function setNewAPR(uint256 _apr) returns()
func (_Rewards *RewardsSession) SetNewAPR(_apr *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.SetNewAPR(&_Rewards.TransactOpts, _apr)
}

// SetNewAPR is a paid mutator transaction binding the contract method 0x5168e544.
//
// Solidity: function setNewAPR(uint256 _apr) returns()
func (_Rewards *RewardsTransactorSession) SetNewAPR(_apr *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.SetNewAPR(&_Rewards.TransactOpts, _apr)
}

// SetupPullToken is a paid mutator transaction binding the contract method 0x8a8a6f59.
//
// Solidity: function setupPullToken(address source, uint256 startTs, uint256 endTs) returns()
func (_Rewards *RewardsTransactor) SetupPullToken(opts *bind.TransactOpts, source common.Address, startTs *big.Int, endTs *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setupPullToken", source, startTs, endTs)
}

// SetupPullToken is a paid mutator transaction binding the contract method 0x8a8a6f59.
//
// Solidity: function setupPullToken(address source, uint256 startTs, uint256 endTs) returns()
func (_Rewards *RewardsSession) SetupPullToken(source common.Address, startTs *big.Int, endTs *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.SetupPullToken(&_Rewards.TransactOpts, source, startTs, endTs)
}

// SetupPullToken is a paid mutator transaction binding the contract method 0x8a8a6f59.
//
// Solidity: function setupPullToken(address source, uint256 startTs, uint256 endTs) returns()
func (_Rewards *RewardsTransactorSession) SetupPullToken(source common.Address, startTs *big.Int, endTs *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.SetupPullToken(&_Rewards.TransactOpts, source, startTs, endTs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewards *RewardsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewards *RewardsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.TransferOwnership(&_Rewards.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewards *RewardsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.TransferOwnership(&_Rewards.TransactOpts, newOwner)
}

// RewardsClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Rewards contract.
type RewardsClaimIterator struct {
	Event *RewardsClaim // Event containing the contract specifics and raw log

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
func (it *RewardsClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsClaim)
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
		it.Event = new(RewardsClaim)
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
func (it *RewardsClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsClaim represents a Claim event raised by the Rewards contract.
type RewardsClaim struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed user, uint256 amount)
func (_Rewards *RewardsFilterer) FilterClaim(opts *bind.FilterOpts, user []common.Address) (*RewardsClaimIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return &RewardsClaimIterator{contract: _Rewards.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed user, uint256 amount)
func (_Rewards *RewardsFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *RewardsClaim, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsClaim)
				if err := _Rewards.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed user, uint256 amount)
func (_Rewards *RewardsFilterer) ParseClaim(log types.Log) (*RewardsClaim, error) {
	event := new(RewardsClaim)
	if err := _Rewards.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rewards contract.
type RewardsOwnershipTransferredIterator struct {
	Event *RewardsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RewardsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsOwnershipTransferred)
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
		it.Event = new(RewardsOwnershipTransferred)
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
func (it *RewardsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsOwnershipTransferred represents a OwnershipTransferred event raised by the Rewards contract.
type RewardsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewards *RewardsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewardsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewardsOwnershipTransferredIterator{contract: _Rewards.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewards *RewardsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewardsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsOwnershipTransferred)
				if err := _Rewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewards *RewardsFilterer) ParseOwnershipTransferred(log types.Log) (*RewardsOwnershipTransferred, error) {
	event := new(RewardsOwnershipTransferred)
	if err := _Rewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
