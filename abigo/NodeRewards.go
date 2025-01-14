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

// NodeRewardsMetaData contains all meta data concerning the NodeRewards contract.
var NodeRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swingby\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_source\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ackFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceBefore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"barn\",\"outputs\":[{\"internalType\":\"contractIBarn\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastPullTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"owed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"resetUnstakedNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_barn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_swap\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startTs\",\"type\":\"uint256\"}],\"name\":\"setBarnAndSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"}],\"name\":\"setNewAPR\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"source\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapContract\",\"outputs\":[{\"internalType\":\"contractISwapContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalNodeStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateNodes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isStaker\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NodeRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeRewardsMetaData.ABI instead.
var NodeRewardsABI = NodeRewardsMetaData.ABI

// NodeRewards is an auto generated Go binding around an Ethereum contract.
type NodeRewards struct {
	NodeRewardsCaller     // Read-only binding to the contract
	NodeRewardsTransactor // Write-only binding to the contract
	NodeRewardsFilterer   // Log filterer for contract events
}

// NodeRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeRewardsSession struct {
	Contract     *NodeRewards      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeRewardsCallerSession struct {
	Contract *NodeRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodeRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeRewardsTransactorSession struct {
	Contract     *NodeRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodeRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRewardsRaw struct {
	Contract *NodeRewards // Generic contract binding to access the raw methods on
}

// NodeRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeRewardsCallerRaw struct {
	Contract *NodeRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// NodeRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeRewardsTransactorRaw struct {
	Contract *NodeRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeRewards creates a new instance of NodeRewards, bound to a specific deployed contract.
func NewNodeRewards(address common.Address, backend bind.ContractBackend) (*NodeRewards, error) {
	contract, err := bindNodeRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeRewards{NodeRewardsCaller: NodeRewardsCaller{contract: contract}, NodeRewardsTransactor: NodeRewardsTransactor{contract: contract}, NodeRewardsFilterer: NodeRewardsFilterer{contract: contract}}, nil
}

// NewNodeRewardsCaller creates a new read-only instance of NodeRewards, bound to a specific deployed contract.
func NewNodeRewardsCaller(address common.Address, caller bind.ContractCaller) (*NodeRewardsCaller, error) {
	contract, err := bindNodeRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRewardsCaller{contract: contract}, nil
}

// NewNodeRewardsTransactor creates a new write-only instance of NodeRewards, bound to a specific deployed contract.
func NewNodeRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeRewardsTransactor, error) {
	contract, err := bindNodeRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRewardsTransactor{contract: contract}, nil
}

// NewNodeRewardsFilterer creates a new log filterer instance of NodeRewards, bound to a specific deployed contract.
func NewNodeRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeRewardsFilterer, error) {
	contract, err := bindNodeRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeRewardsFilterer{contract: contract}, nil
}

// bindNodeRewards binds a generic wrapper to an already deployed contract.
func bindNodeRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeRewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRewards *NodeRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeRewards.Contract.NodeRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRewards *NodeRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRewards.Contract.NodeRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRewards *NodeRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRewards.Contract.NodeRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRewards *NodeRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRewards *NodeRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRewards *NodeRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRewards.Contract.contract.Transact(opts, method, params...)
}

// Apr is a free data retrieval call binding the contract method 0x57ded9c9.
//
// Solidity: function apr() view returns(uint256)
func (_NodeRewards *NodeRewardsCaller) Apr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeRewards.contract.Call(opts, &out, "apr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Apr is a free data retrieval call binding the contract method 0x57ded9c9.
//
// Solidity: function apr() view returns(uint256)
func (_NodeRewards *NodeRewardsSession) Apr() (*big.Int, error) {
	return _NodeRewards.Contract.Apr(&_NodeRewards.CallOpts)
}

// Apr is a free data retrieval call binding the contract method 0x57ded9c9.
//
// Solidity: function apr() view returns(uint256)
func (_NodeRewards *NodeRewardsCallerSession) Apr() (*big.Int, error) {
	return _NodeRewards.Contract.Apr(&_NodeRewards.CallOpts)
}

// BalanceBefore is a free data retrieval call binding the contract method 0x94b5798a.
//
// Solidity: function balanceBefore() view returns(uint256)
func (_NodeRewards *NodeRewardsCaller) BalanceBefore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeRewards.contract.Call(opts, &out, "balanceBefore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceBefore is a free data retrieval call binding the contract method 0x94b5798a.
//
// Solidity: function balanceBefore() view returns(uint256)
func (_NodeRewards *NodeRewardsSession) BalanceBefore() (*big.Int, error) {
	return _NodeRewards.Contract.BalanceBefore(&_NodeRewards.CallOpts)
}

// BalanceBefore is a free data retrieval call binding the contract method 0x94b5798a.
//
// Solidity: function balanceBefore() view returns(uint256)
func (_NodeRewards *NodeRewardsCallerSession) BalanceBefore() (*big.Int, error) {
	return _NodeRewards.Contract.BalanceBefore(&_NodeRewards.CallOpts)
}

// Barn is a free data retrieval call binding the contract method 0x194f480e.
//
// Solidity: function barn() view returns(address)
func (_NodeRewards *NodeRewardsCaller) Barn(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeRewards.contract.Call(opts, &out, "barn")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Barn is a free data retrieval call binding the contract method 0x194f480e.
//
// Solidity: function barn() view returns(address)
func (_NodeRewards *NodeRewardsSession) Barn() (common.Address, error) {
	return _NodeRewards.Contract.Barn(&_NodeRewards.CallOpts)
}

// Barn is a free data retrieval call binding the contract method 0x194f480e.
//
// Solidity: function barn() view returns(address)
func (_NodeRewards *NodeRewardsCallerSession) Barn() (common.Address, error) {
	return _NodeRewards.Contract.Barn(&_NodeRewards.CallOpts)
}

// CurrentMultiplier is a free data retrieval call binding the contract method 0x6fbaaa1e.
//
// Solidity: function currentMultiplier() view returns(uint256)
func (_NodeRewards *NodeRewardsCaller) CurrentMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeRewards.contract.Call(opts, &out, "currentMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentMultiplier is a free data retrieval call binding the contract method 0x6fbaaa1e.
//
// Solidity: function currentMultiplier() view returns(uint256)
func (_NodeRewards *NodeRewardsSession) CurrentMultiplier() (*big.Int, error) {
	return _NodeRewards.Contract.CurrentMultiplier(&_NodeRewards.CallOpts)
}

// CurrentMultiplier is a free data retrieval call binding the contract method 0x6fbaaa1e.
//
// Solidity: function currentMultiplier() view returns(uint256)
func (_NodeRewards *NodeRewardsCallerSession) CurrentMultiplier() (*big.Int, error) {
	return _NodeRewards.Contract.CurrentMultiplier(&_NodeRewards.CallOpts)
}

// LastPullTs is a free data retrieval call binding the contract method 0x100223bb.
//
// Solidity: function lastPullTs() view returns(uint256)
func (_NodeRewards *NodeRewardsCaller) LastPullTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeRewards.contract.Call(opts, &out, "lastPullTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPullTs is a free data retrieval call binding the contract method 0x100223bb.
//
// Solidity: function lastPullTs() view returns(uint256)
func (_NodeRewards *NodeRewardsSession) LastPullTs() (*big.Int, error) {
	return _NodeRewards.Contract.LastPullTs(&_NodeRewards.CallOpts)
}

// LastPullTs is a free data retrieval call binding the contract method 0x100223bb.
//
// Solidity: function lastPullTs() view returns(uint256)
func (_NodeRewards *NodeRewardsCallerSession) LastPullTs() (*big.Int, error) {
	return _NodeRewards.Contract.LastPullTs(&_NodeRewards.CallOpts)
}

// Owed is a free data retrieval call binding the contract method 0xdf18e047.
//
// Solidity: function owed(address ) view returns(uint256)
func (_NodeRewards *NodeRewardsCaller) Owed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NodeRewards.contract.Call(opts, &out, "owed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Owed is a free data retrieval call binding the contract method 0xdf18e047.
//
// Solidity: function owed(address ) view returns(uint256)
func (_NodeRewards *NodeRewardsSession) Owed(arg0 common.Address) (*big.Int, error) {
	return _NodeRewards.Contract.Owed(&_NodeRewards.CallOpts, arg0)
}

// Owed is a free data retrieval call binding the contract method 0xdf18e047.
//
// Solidity: function owed(address ) view returns(uint256)
func (_NodeRewards *NodeRewardsCallerSession) Owed(arg0 common.Address) (*big.Int, error) {
	return _NodeRewards.Contract.Owed(&_NodeRewards.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeRewards *NodeRewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeRewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeRewards *NodeRewardsSession) Owner() (common.Address, error) {
	return _NodeRewards.Contract.Owner(&_NodeRewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeRewards *NodeRewardsCallerSession) Owner() (common.Address, error) {
	return _NodeRewards.Contract.Owner(&_NodeRewards.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_NodeRewards *NodeRewardsCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeRewards.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_NodeRewards *NodeRewardsSession) RewardToken() (common.Address, error) {
	return _NodeRewards.Contract.RewardToken(&_NodeRewards.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_NodeRewards *NodeRewardsCallerSession) RewardToken() (common.Address, error) {
	return _NodeRewards.Contract.RewardToken(&_NodeRewards.CallOpts)
}

// Source is a free data retrieval call binding the contract method 0x67e828bf.
//
// Solidity: function source() view returns(address)
func (_NodeRewards *NodeRewardsCaller) Source(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeRewards.contract.Call(opts, &out, "source")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Source is a free data retrieval call binding the contract method 0x67e828bf.
//
// Solidity: function source() view returns(address)
func (_NodeRewards *NodeRewardsSession) Source() (common.Address, error) {
	return _NodeRewards.Contract.Source(&_NodeRewards.CallOpts)
}

// Source is a free data retrieval call binding the contract method 0x67e828bf.
//
// Solidity: function source() view returns(address)
func (_NodeRewards *NodeRewardsCallerSession) Source() (common.Address, error) {
	return _NodeRewards.Contract.Source(&_NodeRewards.CallOpts)
}

// SwapContract is a free data retrieval call binding the contract method 0x8ea83031.
//
// Solidity: function swapContract() view returns(address)
func (_NodeRewards *NodeRewardsCaller) SwapContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeRewards.contract.Call(opts, &out, "swapContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapContract is a free data retrieval call binding the contract method 0x8ea83031.
//
// Solidity: function swapContract() view returns(address)
func (_NodeRewards *NodeRewardsSession) SwapContract() (common.Address, error) {
	return _NodeRewards.Contract.SwapContract(&_NodeRewards.CallOpts)
}

// SwapContract is a free data retrieval call binding the contract method 0x8ea83031.
//
// Solidity: function swapContract() view returns(address)
func (_NodeRewards *NodeRewardsCallerSession) SwapContract() (common.Address, error) {
	return _NodeRewards.Contract.SwapContract(&_NodeRewards.CallOpts)
}

// TotalNodeStaked is a free data retrieval call binding the contract method 0x8ca0a7e5.
//
// Solidity: function totalNodeStaked() view returns(uint256)
func (_NodeRewards *NodeRewardsCaller) TotalNodeStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NodeRewards.contract.Call(opts, &out, "totalNodeStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNodeStaked is a free data retrieval call binding the contract method 0x8ca0a7e5.
//
// Solidity: function totalNodeStaked() view returns(uint256)
func (_NodeRewards *NodeRewardsSession) TotalNodeStaked() (*big.Int, error) {
	return _NodeRewards.Contract.TotalNodeStaked(&_NodeRewards.CallOpts)
}

// TotalNodeStaked is a free data retrieval call binding the contract method 0x8ca0a7e5.
//
// Solidity: function totalNodeStaked() view returns(uint256)
func (_NodeRewards *NodeRewardsCallerSession) TotalNodeStaked() (*big.Int, error) {
	return _NodeRewards.Contract.TotalNodeStaked(&_NodeRewards.CallOpts)
}

// UserMultiplier is a free data retrieval call binding the contract method 0xb1a03b6b.
//
// Solidity: function userMultiplier(address ) view returns(uint256)
func (_NodeRewards *NodeRewardsCaller) UserMultiplier(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NodeRewards.contract.Call(opts, &out, "userMultiplier", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMultiplier is a free data retrieval call binding the contract method 0xb1a03b6b.
//
// Solidity: function userMultiplier(address ) view returns(uint256)
func (_NodeRewards *NodeRewardsSession) UserMultiplier(arg0 common.Address) (*big.Int, error) {
	return _NodeRewards.Contract.UserMultiplier(&_NodeRewards.CallOpts, arg0)
}

// UserMultiplier is a free data retrieval call binding the contract method 0xb1a03b6b.
//
// Solidity: function userMultiplier(address ) view returns(uint256)
func (_NodeRewards *NodeRewardsCallerSession) UserMultiplier(arg0 common.Address) (*big.Int, error) {
	return _NodeRewards.Contract.UserMultiplier(&_NodeRewards.CallOpts, arg0)
}

// AckFunds is a paid mutator transaction binding the contract method 0xacfd9325.
//
// Solidity: function ackFunds() returns()
func (_NodeRewards *NodeRewardsTransactor) AckFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRewards.contract.Transact(opts, "ackFunds")
}

// AckFunds is a paid mutator transaction binding the contract method 0xacfd9325.
//
// Solidity: function ackFunds() returns()
func (_NodeRewards *NodeRewardsSession) AckFunds() (*types.Transaction, error) {
	return _NodeRewards.Contract.AckFunds(&_NodeRewards.TransactOpts)
}

// AckFunds is a paid mutator transaction binding the contract method 0xacfd9325.
//
// Solidity: function ackFunds() returns()
func (_NodeRewards *NodeRewardsTransactorSession) AckFunds() (*types.Transaction, error) {
	return _NodeRewards.Contract.AckFunds(&_NodeRewards.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_NodeRewards *NodeRewardsTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRewards.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_NodeRewards *NodeRewardsSession) Claim() (*types.Transaction, error) {
	return _NodeRewards.Contract.Claim(&_NodeRewards.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_NodeRewards *NodeRewardsTransactorSession) Claim() (*types.Transaction, error) {
	return _NodeRewards.Contract.Claim(&_NodeRewards.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_NodeRewards *NodeRewardsTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRewards.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_NodeRewards *NodeRewardsSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _NodeRewards.Contract.EmergencyWithdraw(&_NodeRewards.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_NodeRewards *NodeRewardsTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _NodeRewards.Contract.EmergencyWithdraw(&_NodeRewards.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeRewards *NodeRewardsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRewards.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeRewards *NodeRewardsSession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeRewards.Contract.RenounceOwnership(&_NodeRewards.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeRewards *NodeRewardsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeRewards.Contract.RenounceOwnership(&_NodeRewards.TransactOpts)
}

// ResetUnstakedNode is a paid mutator transaction binding the contract method 0x58411c09.
//
// Solidity: function resetUnstakedNode(address _node) returns()
func (_NodeRewards *NodeRewardsTransactor) ResetUnstakedNode(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _NodeRewards.contract.Transact(opts, "resetUnstakedNode", _node)
}

// ResetUnstakedNode is a paid mutator transaction binding the contract method 0x58411c09.
//
// Solidity: function resetUnstakedNode(address _node) returns()
func (_NodeRewards *NodeRewardsSession) ResetUnstakedNode(_node common.Address) (*types.Transaction, error) {
	return _NodeRewards.Contract.ResetUnstakedNode(&_NodeRewards.TransactOpts, _node)
}

// ResetUnstakedNode is a paid mutator transaction binding the contract method 0x58411c09.
//
// Solidity: function resetUnstakedNode(address _node) returns()
func (_NodeRewards *NodeRewardsTransactorSession) ResetUnstakedNode(_node common.Address) (*types.Transaction, error) {
	return _NodeRewards.Contract.ResetUnstakedNode(&_NodeRewards.TransactOpts, _node)
}

// SetBarnAndSwap is a paid mutator transaction binding the contract method 0x4209fdf3.
//
// Solidity: function setBarnAndSwap(address _barn, address _swap, uint256 _startTs) returns()
func (_NodeRewards *NodeRewardsTransactor) SetBarnAndSwap(opts *bind.TransactOpts, _barn common.Address, _swap common.Address, _startTs *big.Int) (*types.Transaction, error) {
	return _NodeRewards.contract.Transact(opts, "setBarnAndSwap", _barn, _swap, _startTs)
}

// SetBarnAndSwap is a paid mutator transaction binding the contract method 0x4209fdf3.
//
// Solidity: function setBarnAndSwap(address _barn, address _swap, uint256 _startTs) returns()
func (_NodeRewards *NodeRewardsSession) SetBarnAndSwap(_barn common.Address, _swap common.Address, _startTs *big.Int) (*types.Transaction, error) {
	return _NodeRewards.Contract.SetBarnAndSwap(&_NodeRewards.TransactOpts, _barn, _swap, _startTs)
}

// SetBarnAndSwap is a paid mutator transaction binding the contract method 0x4209fdf3.
//
// Solidity: function setBarnAndSwap(address _barn, address _swap, uint256 _startTs) returns()
func (_NodeRewards *NodeRewardsTransactorSession) SetBarnAndSwap(_barn common.Address, _swap common.Address, _startTs *big.Int) (*types.Transaction, error) {
	return _NodeRewards.Contract.SetBarnAndSwap(&_NodeRewards.TransactOpts, _barn, _swap, _startTs)
}

// SetNewAPR is a paid mutator transaction binding the contract method 0x5168e544.
//
// Solidity: function setNewAPR(uint256 _apr) returns()
func (_NodeRewards *NodeRewardsTransactor) SetNewAPR(opts *bind.TransactOpts, _apr *big.Int) (*types.Transaction, error) {
	return _NodeRewards.contract.Transact(opts, "setNewAPR", _apr)
}

// SetNewAPR is a paid mutator transaction binding the contract method 0x5168e544.
//
// Solidity: function setNewAPR(uint256 _apr) returns()
func (_NodeRewards *NodeRewardsSession) SetNewAPR(_apr *big.Int) (*types.Transaction, error) {
	return _NodeRewards.Contract.SetNewAPR(&_NodeRewards.TransactOpts, _apr)
}

// SetNewAPR is a paid mutator transaction binding the contract method 0x5168e544.
//
// Solidity: function setNewAPR(uint256 _apr) returns()
func (_NodeRewards *NodeRewardsTransactorSession) SetNewAPR(_apr *big.Int) (*types.Transaction, error) {
	return _NodeRewards.Contract.SetNewAPR(&_NodeRewards.TransactOpts, _apr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeRewards *NodeRewardsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NodeRewards.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeRewards *NodeRewardsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeRewards.Contract.TransferOwnership(&_NodeRewards.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeRewards *NodeRewardsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeRewards.Contract.TransferOwnership(&_NodeRewards.TransactOpts, newOwner)
}

// UpdateNodes is a paid mutator transaction binding the contract method 0xa96adec2.
//
// Solidity: function updateNodes() returns(bool isStaker)
func (_NodeRewards *NodeRewardsTransactor) UpdateNodes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRewards.contract.Transact(opts, "updateNodes")
}

// UpdateNodes is a paid mutator transaction binding the contract method 0xa96adec2.
//
// Solidity: function updateNodes() returns(bool isStaker)
func (_NodeRewards *NodeRewardsSession) UpdateNodes() (*types.Transaction, error) {
	return _NodeRewards.Contract.UpdateNodes(&_NodeRewards.TransactOpts)
}

// UpdateNodes is a paid mutator transaction binding the contract method 0xa96adec2.
//
// Solidity: function updateNodes() returns(bool isStaker)
func (_NodeRewards *NodeRewardsTransactorSession) UpdateNodes() (*types.Transaction, error) {
	return _NodeRewards.Contract.UpdateNodes(&_NodeRewards.TransactOpts)
}

// NodeRewardsClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the NodeRewards contract.
type NodeRewardsClaimIterator struct {
	Event *NodeRewardsClaim // Event containing the contract specifics and raw log

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
func (it *NodeRewardsClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRewardsClaim)
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
		it.Event = new(NodeRewardsClaim)
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
func (it *NodeRewardsClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRewardsClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRewardsClaim represents a Claim event raised by the NodeRewards contract.
type NodeRewardsClaim struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed user, uint256 amount)
func (_NodeRewards *NodeRewardsFilterer) FilterClaim(opts *bind.FilterOpts, user []common.Address) (*NodeRewardsClaimIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _NodeRewards.contract.FilterLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return &NodeRewardsClaimIterator{contract: _NodeRewards.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed user, uint256 amount)
func (_NodeRewards *NodeRewardsFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *NodeRewardsClaim, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _NodeRewards.contract.WatchLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRewardsClaim)
				if err := _NodeRewards.contract.UnpackLog(event, "Claim", log); err != nil {
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
func (_NodeRewards *NodeRewardsFilterer) ParseClaim(log types.Log) (*NodeRewardsClaim, error) {
	event := new(NodeRewardsClaim)
	if err := _NodeRewards.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRewardsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NodeRewards contract.
type NodeRewardsOwnershipTransferredIterator struct {
	Event *NodeRewardsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NodeRewardsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRewardsOwnershipTransferred)
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
		it.Event = new(NodeRewardsOwnershipTransferred)
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
func (it *NodeRewardsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRewardsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRewardsOwnershipTransferred represents a OwnershipTransferred event raised by the NodeRewards contract.
type NodeRewardsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeRewards *NodeRewardsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodeRewardsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeRewards.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeRewardsOwnershipTransferredIterator{contract: _NodeRewards.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeRewards *NodeRewardsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodeRewardsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeRewards.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRewardsOwnershipTransferred)
				if err := _NodeRewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NodeRewards *NodeRewardsFilterer) ParseOwnershipTransferred(log types.Log) (*NodeRewardsOwnershipTransferred, error) {
	event := new(NodeRewardsOwnershipTransferred)
	if err := _NodeRewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
