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

// BarnFacetMetaData contains all meta data concerning the BarnFacet contract.
var BarnFacetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Delegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to_newDelegatedPower\",\"type\":\"uint256\"}],\"name\":\"DelegatedPowerDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to_newDelegatedPower\",\"type\":\"uint256\"}],\"name\":\"DelegatedPowerIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Lock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"name\":\"Timelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"dataType\",\"type\":\"uint8\"}],\"name\":\"UnlockTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWithdrew\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLeft\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_LOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"addOrAdjusttimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"balanceAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"bondStakedAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"checkTimeLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"delegatedPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"delegatedPowerAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"depositAndLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bond\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"}],\"name\":\"initBarn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"multiplierAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"multiplierOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"stakeAtTs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegatedTo\",\"type\":\"address\"}],\"internalType\":\"structLibBarnStorage.Stake\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userDelegatedTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userLockedUntil\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"votingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"votingPowerAtTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BarnFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use BarnFacetMetaData.ABI instead.
var BarnFacetABI = BarnFacetMetaData.ABI

// BarnFacet is an auto generated Go binding around an Ethereum contract.
type BarnFacet struct {
	BarnFacetCaller     // Read-only binding to the contract
	BarnFacetTransactor // Write-only binding to the contract
	BarnFacetFilterer   // Log filterer for contract events
}

// BarnFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BarnFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BarnFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BarnFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BarnFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BarnFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BarnFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BarnFacetSession struct {
	Contract     *BarnFacet        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BarnFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BarnFacetCallerSession struct {
	Contract *BarnFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BarnFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BarnFacetTransactorSession struct {
	Contract     *BarnFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BarnFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BarnFacetRaw struct {
	Contract *BarnFacet // Generic contract binding to access the raw methods on
}

// BarnFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BarnFacetCallerRaw struct {
	Contract *BarnFacetCaller // Generic read-only contract binding to access the raw methods on
}

// BarnFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BarnFacetTransactorRaw struct {
	Contract *BarnFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBarnFacet creates a new instance of BarnFacet, bound to a specific deployed contract.
func NewBarnFacet(address common.Address, backend bind.ContractBackend) (*BarnFacet, error) {
	contract, err := bindBarnFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BarnFacet{BarnFacetCaller: BarnFacetCaller{contract: contract}, BarnFacetTransactor: BarnFacetTransactor{contract: contract}, BarnFacetFilterer: BarnFacetFilterer{contract: contract}}, nil
}

// NewBarnFacetCaller creates a new read-only instance of BarnFacet, bound to a specific deployed contract.
func NewBarnFacetCaller(address common.Address, caller bind.ContractCaller) (*BarnFacetCaller, error) {
	contract, err := bindBarnFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BarnFacetCaller{contract: contract}, nil
}

// NewBarnFacetTransactor creates a new write-only instance of BarnFacet, bound to a specific deployed contract.
func NewBarnFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*BarnFacetTransactor, error) {
	contract, err := bindBarnFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BarnFacetTransactor{contract: contract}, nil
}

// NewBarnFacetFilterer creates a new log filterer instance of BarnFacet, bound to a specific deployed contract.
func NewBarnFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*BarnFacetFilterer, error) {
	contract, err := bindBarnFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BarnFacetFilterer{contract: contract}, nil
}

// bindBarnFacet binds a generic wrapper to an already deployed contract.
func bindBarnFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BarnFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BarnFacet *BarnFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BarnFacet.Contract.BarnFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BarnFacet *BarnFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BarnFacet.Contract.BarnFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BarnFacet *BarnFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BarnFacet.Contract.BarnFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BarnFacet *BarnFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BarnFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BarnFacet *BarnFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BarnFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BarnFacet *BarnFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BarnFacet.Contract.contract.Transact(opts, method, params...)
}

// MAXLOCK is a free data retrieval call binding the contract method 0x65a5d5f0.
//
// Solidity: function MAX_LOCK() view returns(uint256)
func (_BarnFacet *BarnFacetCaller) MAXLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "MAX_LOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLOCK is a free data retrieval call binding the contract method 0x65a5d5f0.
//
// Solidity: function MAX_LOCK() view returns(uint256)
func (_BarnFacet *BarnFacetSession) MAXLOCK() (*big.Int, error) {
	return _BarnFacet.Contract.MAXLOCK(&_BarnFacet.CallOpts)
}

// MAXLOCK is a free data retrieval call binding the contract method 0x65a5d5f0.
//
// Solidity: function MAX_LOCK() view returns(uint256)
func (_BarnFacet *BarnFacetCallerSession) MAXLOCK() (*big.Int, error) {
	return _BarnFacet.Contract.MAXLOCK(&_BarnFacet.CallOpts)
}

// BalanceAtTs is a free data retrieval call binding the contract method 0x417edd4d.
//
// Solidity: function balanceAtTs(address user, uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetCaller) BalanceAtTs(opts *bind.CallOpts, user common.Address, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "balanceAtTs", user, timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceAtTs is a free data retrieval call binding the contract method 0x417edd4d.
//
// Solidity: function balanceAtTs(address user, uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetSession) BalanceAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _BarnFacet.Contract.BalanceAtTs(&_BarnFacet.CallOpts, user, timestamp)
}

// BalanceAtTs is a free data retrieval call binding the contract method 0x417edd4d.
//
// Solidity: function balanceAtTs(address user, uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetCallerSession) BalanceAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _BarnFacet.Contract.BalanceAtTs(&_BarnFacet.CallOpts, user, timestamp)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_BarnFacet *BarnFacetCaller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_BarnFacet *BarnFacetSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _BarnFacet.Contract.BalanceOf(&_BarnFacet.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_BarnFacet *BarnFacetCallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _BarnFacet.Contract.BalanceOf(&_BarnFacet.CallOpts, user)
}

// BondStaked is a free data retrieval call binding the contract method 0xc2077e81.
//
// Solidity: function bondStaked() view returns(uint256)
func (_BarnFacet *BarnFacetCaller) BondStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "bondStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondStaked is a free data retrieval call binding the contract method 0xc2077e81.
//
// Solidity: function bondStaked() view returns(uint256)
func (_BarnFacet *BarnFacetSession) BondStaked() (*big.Int, error) {
	return _BarnFacet.Contract.BondStaked(&_BarnFacet.CallOpts)
}

// BondStaked is a free data retrieval call binding the contract method 0xc2077e81.
//
// Solidity: function bondStaked() view returns(uint256)
func (_BarnFacet *BarnFacetCallerSession) BondStaked() (*big.Int, error) {
	return _BarnFacet.Contract.BondStaked(&_BarnFacet.CallOpts)
}

// BondStakedAtTs is a free data retrieval call binding the contract method 0xf77f962f.
//
// Solidity: function bondStakedAtTs(uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetCaller) BondStakedAtTs(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "bondStakedAtTs", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondStakedAtTs is a free data retrieval call binding the contract method 0xf77f962f.
//
// Solidity: function bondStakedAtTs(uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetSession) BondStakedAtTs(timestamp *big.Int) (*big.Int, error) {
	return _BarnFacet.Contract.BondStakedAtTs(&_BarnFacet.CallOpts, timestamp)
}

// BondStakedAtTs is a free data retrieval call binding the contract method 0xf77f962f.
//
// Solidity: function bondStakedAtTs(uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetCallerSession) BondStakedAtTs(timestamp *big.Int) (*big.Int, error) {
	return _BarnFacet.Contract.BondStakedAtTs(&_BarnFacet.CallOpts, timestamp)
}

// CheckTimeLock is a free data retrieval call binding the contract method 0x697eb324.
//
// Solidity: function checkTimeLock(address _user, uint8 _type) view returns(uint256 amount, uint256 expiry, bytes32 data)
func (_BarnFacet *BarnFacetCaller) CheckTimeLock(opts *bind.CallOpts, _user common.Address, _type uint8) (struct {
	Amount *big.Int
	Expiry *big.Int
	Data   [32]byte
}, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "checkTimeLock", _user, _type)

	outstruct := new(struct {
		Amount *big.Int
		Expiry *big.Int
		Data   [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Expiry = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// CheckTimeLock is a free data retrieval call binding the contract method 0x697eb324.
//
// Solidity: function checkTimeLock(address _user, uint8 _type) view returns(uint256 amount, uint256 expiry, bytes32 data)
func (_BarnFacet *BarnFacetSession) CheckTimeLock(_user common.Address, _type uint8) (struct {
	Amount *big.Int
	Expiry *big.Int
	Data   [32]byte
}, error) {
	return _BarnFacet.Contract.CheckTimeLock(&_BarnFacet.CallOpts, _user, _type)
}

// CheckTimeLock is a free data retrieval call binding the contract method 0x697eb324.
//
// Solidity: function checkTimeLock(address _user, uint8 _type) view returns(uint256 amount, uint256 expiry, bytes32 data)
func (_BarnFacet *BarnFacetCallerSession) CheckTimeLock(_user common.Address, _type uint8) (struct {
	Amount *big.Int
	Expiry *big.Int
	Data   [32]byte
}, error) {
	return _BarnFacet.Contract.CheckTimeLock(&_BarnFacet.CallOpts, _user, _type)
}

// DelegatedPower is a free data retrieval call binding the contract method 0x169df064.
//
// Solidity: function delegatedPower(address user) view returns(uint256)
func (_BarnFacet *BarnFacetCaller) DelegatedPower(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "delegatedPower", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedPower is a free data retrieval call binding the contract method 0x169df064.
//
// Solidity: function delegatedPower(address user) view returns(uint256)
func (_BarnFacet *BarnFacetSession) DelegatedPower(user common.Address) (*big.Int, error) {
	return _BarnFacet.Contract.DelegatedPower(&_BarnFacet.CallOpts, user)
}

// DelegatedPower is a free data retrieval call binding the contract method 0x169df064.
//
// Solidity: function delegatedPower(address user) view returns(uint256)
func (_BarnFacet *BarnFacetCallerSession) DelegatedPower(user common.Address) (*big.Int, error) {
	return _BarnFacet.Contract.DelegatedPower(&_BarnFacet.CallOpts, user)
}

// DelegatedPowerAtTs is a free data retrieval call binding the contract method 0xd265a115.
//
// Solidity: function delegatedPowerAtTs(address user, uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetCaller) DelegatedPowerAtTs(opts *bind.CallOpts, user common.Address, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "delegatedPowerAtTs", user, timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatedPowerAtTs is a free data retrieval call binding the contract method 0xd265a115.
//
// Solidity: function delegatedPowerAtTs(address user, uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetSession) DelegatedPowerAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _BarnFacet.Contract.DelegatedPowerAtTs(&_BarnFacet.CallOpts, user, timestamp)
}

// DelegatedPowerAtTs is a free data retrieval call binding the contract method 0xd265a115.
//
// Solidity: function delegatedPowerAtTs(address user, uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetCallerSession) DelegatedPowerAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _BarnFacet.Contract.DelegatedPowerAtTs(&_BarnFacet.CallOpts, user, timestamp)
}

// MultiplierAtTs is a free data retrieval call binding the contract method 0x7a141096.
//
// Solidity: function multiplierAtTs(address user, uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetCaller) MultiplierAtTs(opts *bind.CallOpts, user common.Address, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "multiplierAtTs", user, timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MultiplierAtTs is a free data retrieval call binding the contract method 0x7a141096.
//
// Solidity: function multiplierAtTs(address user, uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetSession) MultiplierAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _BarnFacet.Contract.MultiplierAtTs(&_BarnFacet.CallOpts, user, timestamp)
}

// MultiplierAtTs is a free data retrieval call binding the contract method 0x7a141096.
//
// Solidity: function multiplierAtTs(address user, uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetCallerSession) MultiplierAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _BarnFacet.Contract.MultiplierAtTs(&_BarnFacet.CallOpts, user, timestamp)
}

// MultiplierOf is a free data retrieval call binding the contract method 0x8e4a5248.
//
// Solidity: function multiplierOf(address user) view returns(uint256)
func (_BarnFacet *BarnFacetCaller) MultiplierOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "multiplierOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MultiplierOf is a free data retrieval call binding the contract method 0x8e4a5248.
//
// Solidity: function multiplierOf(address user) view returns(uint256)
func (_BarnFacet *BarnFacetSession) MultiplierOf(user common.Address) (*big.Int, error) {
	return _BarnFacet.Contract.MultiplierOf(&_BarnFacet.CallOpts, user)
}

// MultiplierOf is a free data retrieval call binding the contract method 0x8e4a5248.
//
// Solidity: function multiplierOf(address user) view returns(uint256)
func (_BarnFacet *BarnFacetCallerSession) MultiplierOf(user common.Address) (*big.Int, error) {
	return _BarnFacet.Contract.MultiplierOf(&_BarnFacet.CallOpts, user)
}

// StakeAtTs is a free data retrieval call binding the contract method 0x18ab6a3c.
//
// Solidity: function stakeAtTs(address user, uint256 timestamp) view returns((uint256,uint256,uint256,address))
func (_BarnFacet *BarnFacetCaller) StakeAtTs(opts *bind.CallOpts, user common.Address, timestamp *big.Int) (LibBarnStorageStake, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "stakeAtTs", user, timestamp)

	if err != nil {
		return *new(LibBarnStorageStake), err
	}

	out0 := *abi.ConvertType(out[0], new(LibBarnStorageStake)).(*LibBarnStorageStake)

	return out0, err

}

// StakeAtTs is a free data retrieval call binding the contract method 0x18ab6a3c.
//
// Solidity: function stakeAtTs(address user, uint256 timestamp) view returns((uint256,uint256,uint256,address))
func (_BarnFacet *BarnFacetSession) StakeAtTs(user common.Address, timestamp *big.Int) (LibBarnStorageStake, error) {
	return _BarnFacet.Contract.StakeAtTs(&_BarnFacet.CallOpts, user, timestamp)
}

// StakeAtTs is a free data retrieval call binding the contract method 0x18ab6a3c.
//
// Solidity: function stakeAtTs(address user, uint256 timestamp) view returns((uint256,uint256,uint256,address))
func (_BarnFacet *BarnFacetCallerSession) StakeAtTs(user common.Address, timestamp *big.Int) (LibBarnStorageStake, error) {
	return _BarnFacet.Contract.StakeAtTs(&_BarnFacet.CallOpts, user, timestamp)
}

// UserDelegatedTo is a free data retrieval call binding the contract method 0xbef624d8.
//
// Solidity: function userDelegatedTo(address user) view returns(address)
func (_BarnFacet *BarnFacetCaller) UserDelegatedTo(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "userDelegatedTo", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserDelegatedTo is a free data retrieval call binding the contract method 0xbef624d8.
//
// Solidity: function userDelegatedTo(address user) view returns(address)
func (_BarnFacet *BarnFacetSession) UserDelegatedTo(user common.Address) (common.Address, error) {
	return _BarnFacet.Contract.UserDelegatedTo(&_BarnFacet.CallOpts, user)
}

// UserDelegatedTo is a free data retrieval call binding the contract method 0xbef624d8.
//
// Solidity: function userDelegatedTo(address user) view returns(address)
func (_BarnFacet *BarnFacetCallerSession) UserDelegatedTo(user common.Address) (common.Address, error) {
	return _BarnFacet.Contract.UserDelegatedTo(&_BarnFacet.CallOpts, user)
}

// UserLockedUntil is a free data retrieval call binding the contract method 0xbf0ae48c.
//
// Solidity: function userLockedUntil(address user) view returns(uint256)
func (_BarnFacet *BarnFacetCaller) UserLockedUntil(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "userLockedUntil", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLockedUntil is a free data retrieval call binding the contract method 0xbf0ae48c.
//
// Solidity: function userLockedUntil(address user) view returns(uint256)
func (_BarnFacet *BarnFacetSession) UserLockedUntil(user common.Address) (*big.Int, error) {
	return _BarnFacet.Contract.UserLockedUntil(&_BarnFacet.CallOpts, user)
}

// UserLockedUntil is a free data retrieval call binding the contract method 0xbf0ae48c.
//
// Solidity: function userLockedUntil(address user) view returns(uint256)
func (_BarnFacet *BarnFacetCallerSession) UserLockedUntil(user common.Address) (*big.Int, error) {
	return _BarnFacet.Contract.UserLockedUntil(&_BarnFacet.CallOpts, user)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address user) view returns(uint256)
func (_BarnFacet *BarnFacetCaller) VotingPower(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "votingPower", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address user) view returns(uint256)
func (_BarnFacet *BarnFacetSession) VotingPower(user common.Address) (*big.Int, error) {
	return _BarnFacet.Contract.VotingPower(&_BarnFacet.CallOpts, user)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address user) view returns(uint256)
func (_BarnFacet *BarnFacetCallerSession) VotingPower(user common.Address) (*big.Int, error) {
	return _BarnFacet.Contract.VotingPower(&_BarnFacet.CallOpts, user)
}

// VotingPowerAtTs is a free data retrieval call binding the contract method 0xcbf8eda9.
//
// Solidity: function votingPowerAtTs(address user, uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetCaller) VotingPowerAtTs(opts *bind.CallOpts, user common.Address, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BarnFacet.contract.Call(opts, &out, "votingPowerAtTs", user, timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPowerAtTs is a free data retrieval call binding the contract method 0xcbf8eda9.
//
// Solidity: function votingPowerAtTs(address user, uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetSession) VotingPowerAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _BarnFacet.Contract.VotingPowerAtTs(&_BarnFacet.CallOpts, user, timestamp)
}

// VotingPowerAtTs is a free data retrieval call binding the contract method 0xcbf8eda9.
//
// Solidity: function votingPowerAtTs(address user, uint256 timestamp) view returns(uint256)
func (_BarnFacet *BarnFacetCallerSession) VotingPowerAtTs(user common.Address, timestamp *big.Int) (*big.Int, error) {
	return _BarnFacet.Contract.VotingPowerAtTs(&_BarnFacet.CallOpts, user, timestamp)
}

// AddOrAdjusttimelock is a paid mutator transaction binding the contract method 0x20e77a84.
//
// Solidity: function addOrAdjusttimelock(uint256 _amount, uint256 _timestamp, uint8 _type, bytes32 _data) returns()
func (_BarnFacet *BarnFacetTransactor) AddOrAdjusttimelock(opts *bind.TransactOpts, _amount *big.Int, _timestamp *big.Int, _type uint8, _data [32]byte) (*types.Transaction, error) {
	return _BarnFacet.contract.Transact(opts, "addOrAdjusttimelock", _amount, _timestamp, _type, _data)
}

// AddOrAdjusttimelock is a paid mutator transaction binding the contract method 0x20e77a84.
//
// Solidity: function addOrAdjusttimelock(uint256 _amount, uint256 _timestamp, uint8 _type, bytes32 _data) returns()
func (_BarnFacet *BarnFacetSession) AddOrAdjusttimelock(_amount *big.Int, _timestamp *big.Int, _type uint8, _data [32]byte) (*types.Transaction, error) {
	return _BarnFacet.Contract.AddOrAdjusttimelock(&_BarnFacet.TransactOpts, _amount, _timestamp, _type, _data)
}

// AddOrAdjusttimelock is a paid mutator transaction binding the contract method 0x20e77a84.
//
// Solidity: function addOrAdjusttimelock(uint256 _amount, uint256 _timestamp, uint8 _type, bytes32 _data) returns()
func (_BarnFacet *BarnFacetTransactorSession) AddOrAdjusttimelock(_amount *big.Int, _timestamp *big.Int, _type uint8, _data [32]byte) (*types.Transaction, error) {
	return _BarnFacet.Contract.AddOrAdjusttimelock(&_BarnFacet.TransactOpts, _amount, _timestamp, _type, _data)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_BarnFacet *BarnFacetTransactor) Delegate(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BarnFacet.contract.Transact(opts, "delegate", to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_BarnFacet *BarnFacetSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _BarnFacet.Contract.Delegate(&_BarnFacet.TransactOpts, to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_BarnFacet *BarnFacetTransactorSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _BarnFacet.Contract.Delegate(&_BarnFacet.TransactOpts, to)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_BarnFacet *BarnFacetTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BarnFacet.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_BarnFacet *BarnFacetSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _BarnFacet.Contract.Deposit(&_BarnFacet.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_BarnFacet *BarnFacetTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _BarnFacet.Contract.Deposit(&_BarnFacet.TransactOpts, amount)
}

// DepositAndLock is a paid mutator transaction binding the contract method 0xbfc10279.
//
// Solidity: function depositAndLock(uint256 amount, uint256 timestamp) returns()
func (_BarnFacet *BarnFacetTransactor) DepositAndLock(opts *bind.TransactOpts, amount *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _BarnFacet.contract.Transact(opts, "depositAndLock", amount, timestamp)
}

// DepositAndLock is a paid mutator transaction binding the contract method 0xbfc10279.
//
// Solidity: function depositAndLock(uint256 amount, uint256 timestamp) returns()
func (_BarnFacet *BarnFacetSession) DepositAndLock(amount *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _BarnFacet.Contract.DepositAndLock(&_BarnFacet.TransactOpts, amount, timestamp)
}

// DepositAndLock is a paid mutator transaction binding the contract method 0xbfc10279.
//
// Solidity: function depositAndLock(uint256 amount, uint256 timestamp) returns()
func (_BarnFacet *BarnFacetTransactorSession) DepositAndLock(amount *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _BarnFacet.Contract.DepositAndLock(&_BarnFacet.TransactOpts, amount, timestamp)
}

// InitBarn is a paid mutator transaction binding the contract method 0x5107a3ae.
//
// Solidity: function initBarn(address _bond, address _rewards) returns()
func (_BarnFacet *BarnFacetTransactor) InitBarn(opts *bind.TransactOpts, _bond common.Address, _rewards common.Address) (*types.Transaction, error) {
	return _BarnFacet.contract.Transact(opts, "initBarn", _bond, _rewards)
}

// InitBarn is a paid mutator transaction binding the contract method 0x5107a3ae.
//
// Solidity: function initBarn(address _bond, address _rewards) returns()
func (_BarnFacet *BarnFacetSession) InitBarn(_bond common.Address, _rewards common.Address) (*types.Transaction, error) {
	return _BarnFacet.Contract.InitBarn(&_BarnFacet.TransactOpts, _bond, _rewards)
}

// InitBarn is a paid mutator transaction binding the contract method 0x5107a3ae.
//
// Solidity: function initBarn(address _bond, address _rewards) returns()
func (_BarnFacet *BarnFacetTransactorSession) InitBarn(_bond common.Address, _rewards common.Address) (*types.Transaction, error) {
	return _BarnFacet.Contract.InitBarn(&_BarnFacet.TransactOpts, _bond, _rewards)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 timestamp) returns()
func (_BarnFacet *BarnFacetTransactor) Lock(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _BarnFacet.contract.Transact(opts, "lock", timestamp)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 timestamp) returns()
func (_BarnFacet *BarnFacetSession) Lock(timestamp *big.Int) (*types.Transaction, error) {
	return _BarnFacet.Contract.Lock(&_BarnFacet.TransactOpts, timestamp)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 timestamp) returns()
func (_BarnFacet *BarnFacetTransactorSession) Lock(timestamp *big.Int) (*types.Transaction, error) {
	return _BarnFacet.Contract.Lock(&_BarnFacet.TransactOpts, timestamp)
}

// StopDelegate is a paid mutator transaction binding the contract method 0x6f121578.
//
// Solidity: function stopDelegate() returns()
func (_BarnFacet *BarnFacetTransactor) StopDelegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BarnFacet.contract.Transact(opts, "stopDelegate")
}

// StopDelegate is a paid mutator transaction binding the contract method 0x6f121578.
//
// Solidity: function stopDelegate() returns()
func (_BarnFacet *BarnFacetSession) StopDelegate() (*types.Transaction, error) {
	return _BarnFacet.Contract.StopDelegate(&_BarnFacet.TransactOpts)
}

// StopDelegate is a paid mutator transaction binding the contract method 0x6f121578.
//
// Solidity: function stopDelegate() returns()
func (_BarnFacet *BarnFacetTransactorSession) StopDelegate() (*types.Transaction, error) {
	return _BarnFacet.Contract.StopDelegate(&_BarnFacet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_BarnFacet *BarnFacetTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BarnFacet.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_BarnFacet *BarnFacetSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _BarnFacet.Contract.Withdraw(&_BarnFacet.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_BarnFacet *BarnFacetTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _BarnFacet.Contract.Withdraw(&_BarnFacet.TransactOpts, amount)
}

// BarnFacetDelegateIterator is returned from FilterDelegate and is used to iterate over the raw logs and unpacked data for Delegate events raised by the BarnFacet contract.
type BarnFacetDelegateIterator struct {
	Event *BarnFacetDelegate // Event containing the contract specifics and raw log

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
func (it *BarnFacetDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BarnFacetDelegate)
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
		it.Event = new(BarnFacetDelegate)
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
func (it *BarnFacetDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BarnFacetDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BarnFacetDelegate represents a Delegate event raised by the BarnFacet contract.
type BarnFacetDelegate struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0xab7d75eccd27c9989942a3a6e4137e415df0ad90ec428751b16361f16fe8780f.
//
// Solidity: event Delegate(address indexed from, address indexed to)
func (_BarnFacet *BarnFacetFilterer) FilterDelegate(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BarnFacetDelegateIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BarnFacet.contract.FilterLogs(opts, "Delegate", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BarnFacetDelegateIterator{contract: _BarnFacet.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0xab7d75eccd27c9989942a3a6e4137e415df0ad90ec428751b16361f16fe8780f.
//
// Solidity: event Delegate(address indexed from, address indexed to)
func (_BarnFacet *BarnFacetFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *BarnFacetDelegate, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BarnFacet.contract.WatchLogs(opts, "Delegate", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BarnFacetDelegate)
				if err := _BarnFacet.contract.UnpackLog(event, "Delegate", log); err != nil {
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

// ParseDelegate is a log parse operation binding the contract event 0xab7d75eccd27c9989942a3a6e4137e415df0ad90ec428751b16361f16fe8780f.
//
// Solidity: event Delegate(address indexed from, address indexed to)
func (_BarnFacet *BarnFacetFilterer) ParseDelegate(log types.Log) (*BarnFacetDelegate, error) {
	event := new(BarnFacetDelegate)
	if err := _BarnFacet.contract.UnpackLog(event, "Delegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BarnFacetDelegatedPowerDecreasedIterator is returned from FilterDelegatedPowerDecreased and is used to iterate over the raw logs and unpacked data for DelegatedPowerDecreased events raised by the BarnFacet contract.
type BarnFacetDelegatedPowerDecreasedIterator struct {
	Event *BarnFacetDelegatedPowerDecreased // Event containing the contract specifics and raw log

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
func (it *BarnFacetDelegatedPowerDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BarnFacetDelegatedPowerDecreased)
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
		it.Event = new(BarnFacetDelegatedPowerDecreased)
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
func (it *BarnFacetDelegatedPowerDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BarnFacetDelegatedPowerDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BarnFacetDelegatedPowerDecreased represents a DelegatedPowerDecreased event raised by the BarnFacet contract.
type BarnFacetDelegatedPowerDecreased struct {
	From                common.Address
	To                  common.Address
	Amount              *big.Int
	ToNewDelegatedPower *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDelegatedPowerDecreased is a free log retrieval operation binding the contract event 0xfb73cd22fb01f433ef312f758a708c1c7d1442ec871b9dd2546b3ec85a8b4e76.
//
// Solidity: event DelegatedPowerDecreased(address indexed from, address indexed to, uint256 amount, uint256 to_newDelegatedPower)
func (_BarnFacet *BarnFacetFilterer) FilterDelegatedPowerDecreased(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BarnFacetDelegatedPowerDecreasedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BarnFacet.contract.FilterLogs(opts, "DelegatedPowerDecreased", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BarnFacetDelegatedPowerDecreasedIterator{contract: _BarnFacet.contract, event: "DelegatedPowerDecreased", logs: logs, sub: sub}, nil
}

// WatchDelegatedPowerDecreased is a free log subscription operation binding the contract event 0xfb73cd22fb01f433ef312f758a708c1c7d1442ec871b9dd2546b3ec85a8b4e76.
//
// Solidity: event DelegatedPowerDecreased(address indexed from, address indexed to, uint256 amount, uint256 to_newDelegatedPower)
func (_BarnFacet *BarnFacetFilterer) WatchDelegatedPowerDecreased(opts *bind.WatchOpts, sink chan<- *BarnFacetDelegatedPowerDecreased, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BarnFacet.contract.WatchLogs(opts, "DelegatedPowerDecreased", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BarnFacetDelegatedPowerDecreased)
				if err := _BarnFacet.contract.UnpackLog(event, "DelegatedPowerDecreased", log); err != nil {
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

// ParseDelegatedPowerDecreased is a log parse operation binding the contract event 0xfb73cd22fb01f433ef312f758a708c1c7d1442ec871b9dd2546b3ec85a8b4e76.
//
// Solidity: event DelegatedPowerDecreased(address indexed from, address indexed to, uint256 amount, uint256 to_newDelegatedPower)
func (_BarnFacet *BarnFacetFilterer) ParseDelegatedPowerDecreased(log types.Log) (*BarnFacetDelegatedPowerDecreased, error) {
	event := new(BarnFacetDelegatedPowerDecreased)
	if err := _BarnFacet.contract.UnpackLog(event, "DelegatedPowerDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BarnFacetDelegatedPowerIncreasedIterator is returned from FilterDelegatedPowerIncreased and is used to iterate over the raw logs and unpacked data for DelegatedPowerIncreased events raised by the BarnFacet contract.
type BarnFacetDelegatedPowerIncreasedIterator struct {
	Event *BarnFacetDelegatedPowerIncreased // Event containing the contract specifics and raw log

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
func (it *BarnFacetDelegatedPowerIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BarnFacetDelegatedPowerIncreased)
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
		it.Event = new(BarnFacetDelegatedPowerIncreased)
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
func (it *BarnFacetDelegatedPowerIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BarnFacetDelegatedPowerIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BarnFacetDelegatedPowerIncreased represents a DelegatedPowerIncreased event raised by the BarnFacet contract.
type BarnFacetDelegatedPowerIncreased struct {
	From                common.Address
	To                  common.Address
	Amount              *big.Int
	ToNewDelegatedPower *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDelegatedPowerIncreased is a free log retrieval operation binding the contract event 0x9306546ca617a204223f7da51d942104c887cf8e53f8fd454af55a529aaa689a.
//
// Solidity: event DelegatedPowerIncreased(address indexed from, address indexed to, uint256 amount, uint256 to_newDelegatedPower)
func (_BarnFacet *BarnFacetFilterer) FilterDelegatedPowerIncreased(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BarnFacetDelegatedPowerIncreasedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BarnFacet.contract.FilterLogs(opts, "DelegatedPowerIncreased", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BarnFacetDelegatedPowerIncreasedIterator{contract: _BarnFacet.contract, event: "DelegatedPowerIncreased", logs: logs, sub: sub}, nil
}

// WatchDelegatedPowerIncreased is a free log subscription operation binding the contract event 0x9306546ca617a204223f7da51d942104c887cf8e53f8fd454af55a529aaa689a.
//
// Solidity: event DelegatedPowerIncreased(address indexed from, address indexed to, uint256 amount, uint256 to_newDelegatedPower)
func (_BarnFacet *BarnFacetFilterer) WatchDelegatedPowerIncreased(opts *bind.WatchOpts, sink chan<- *BarnFacetDelegatedPowerIncreased, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BarnFacet.contract.WatchLogs(opts, "DelegatedPowerIncreased", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BarnFacetDelegatedPowerIncreased)
				if err := _BarnFacet.contract.UnpackLog(event, "DelegatedPowerIncreased", log); err != nil {
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

// ParseDelegatedPowerIncreased is a log parse operation binding the contract event 0x9306546ca617a204223f7da51d942104c887cf8e53f8fd454af55a529aaa689a.
//
// Solidity: event DelegatedPowerIncreased(address indexed from, address indexed to, uint256 amount, uint256 to_newDelegatedPower)
func (_BarnFacet *BarnFacetFilterer) ParseDelegatedPowerIncreased(log types.Log) (*BarnFacetDelegatedPowerIncreased, error) {
	event := new(BarnFacetDelegatedPowerIncreased)
	if err := _BarnFacet.contract.UnpackLog(event, "DelegatedPowerIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BarnFacetDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BarnFacet contract.
type BarnFacetDepositIterator struct {
	Event *BarnFacetDeposit // Event containing the contract specifics and raw log

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
func (it *BarnFacetDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BarnFacetDeposit)
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
		it.Event = new(BarnFacetDeposit)
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
func (it *BarnFacetDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BarnFacetDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BarnFacetDeposit represents a Deposit event raised by the BarnFacet contract.
type BarnFacetDeposit struct {
	User       common.Address
	Amount     *big.Int
	NewBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 newBalance)
func (_BarnFacet *BarnFacetFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*BarnFacetDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BarnFacet.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &BarnFacetDepositIterator{contract: _BarnFacet.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 newBalance)
func (_BarnFacet *BarnFacetFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BarnFacetDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BarnFacet.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BarnFacetDeposit)
				if err := _BarnFacet.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 amount, uint256 newBalance)
func (_BarnFacet *BarnFacetFilterer) ParseDeposit(log types.Log) (*BarnFacetDeposit, error) {
	event := new(BarnFacetDeposit)
	if err := _BarnFacet.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BarnFacetLockIterator is returned from FilterLock and is used to iterate over the raw logs and unpacked data for Lock events raised by the BarnFacet contract.
type BarnFacetLockIterator struct {
	Event *BarnFacetLock // Event containing the contract specifics and raw log

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
func (it *BarnFacetLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BarnFacetLock)
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
		it.Event = new(BarnFacetLock)
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
func (it *BarnFacetLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BarnFacetLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BarnFacetLock represents a Lock event raised by the BarnFacet contract.
type BarnFacetLock struct {
	User      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLock is a free log retrieval operation binding the contract event 0x625fed9875dada8643f2418b838ae0bc78d9a148a18eee4ee1979ff0f3f5d427.
//
// Solidity: event Lock(address indexed user, uint256 timestamp)
func (_BarnFacet *BarnFacetFilterer) FilterLock(opts *bind.FilterOpts, user []common.Address) (*BarnFacetLockIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BarnFacet.contract.FilterLogs(opts, "Lock", userRule)
	if err != nil {
		return nil, err
	}
	return &BarnFacetLockIterator{contract: _BarnFacet.contract, event: "Lock", logs: logs, sub: sub}, nil
}

// WatchLock is a free log subscription operation binding the contract event 0x625fed9875dada8643f2418b838ae0bc78d9a148a18eee4ee1979ff0f3f5d427.
//
// Solidity: event Lock(address indexed user, uint256 timestamp)
func (_BarnFacet *BarnFacetFilterer) WatchLock(opts *bind.WatchOpts, sink chan<- *BarnFacetLock, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BarnFacet.contract.WatchLogs(opts, "Lock", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BarnFacetLock)
				if err := _BarnFacet.contract.UnpackLog(event, "Lock", log); err != nil {
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

// ParseLock is a log parse operation binding the contract event 0x625fed9875dada8643f2418b838ae0bc78d9a148a18eee4ee1979ff0f3f5d427.
//
// Solidity: event Lock(address indexed user, uint256 timestamp)
func (_BarnFacet *BarnFacetFilterer) ParseLock(log types.Log) (*BarnFacetLock, error) {
	event := new(BarnFacetLock)
	if err := _BarnFacet.contract.UnpackLog(event, "Lock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BarnFacetTimelockIterator is returned from FilterTimelock and is used to iterate over the raw logs and unpacked data for Timelock events raised by the BarnFacet contract.
type BarnFacetTimelockIterator struct {
	Event *BarnFacetTimelock // Event containing the contract specifics and raw log

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
func (it *BarnFacetTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BarnFacetTimelock)
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
		it.Event = new(BarnFacetTimelock)
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
func (it *BarnFacetTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BarnFacetTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BarnFacetTimelock represents a Timelock event raised by the BarnFacet contract.
type BarnFacetTimelock struct {
	User      common.Address
	Timestamp *big.Int
	DataType  uint8
	Data      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelock is a free log retrieval operation binding the contract event 0xc262f2ecc2db2f7833ae92a6147054879e39334843267f53fea2a25bf04407ba.
//
// Solidity: event Timelock(address indexed user, uint256 timestamp, uint8 dataType, bytes32 data)
func (_BarnFacet *BarnFacetFilterer) FilterTimelock(opts *bind.FilterOpts, user []common.Address) (*BarnFacetTimelockIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BarnFacet.contract.FilterLogs(opts, "Timelock", userRule)
	if err != nil {
		return nil, err
	}
	return &BarnFacetTimelockIterator{contract: _BarnFacet.contract, event: "Timelock", logs: logs, sub: sub}, nil
}

// WatchTimelock is a free log subscription operation binding the contract event 0xc262f2ecc2db2f7833ae92a6147054879e39334843267f53fea2a25bf04407ba.
//
// Solidity: event Timelock(address indexed user, uint256 timestamp, uint8 dataType, bytes32 data)
func (_BarnFacet *BarnFacetFilterer) WatchTimelock(opts *bind.WatchOpts, sink chan<- *BarnFacetTimelock, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BarnFacet.contract.WatchLogs(opts, "Timelock", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BarnFacetTimelock)
				if err := _BarnFacet.contract.UnpackLog(event, "Timelock", log); err != nil {
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

// ParseTimelock is a log parse operation binding the contract event 0xc262f2ecc2db2f7833ae92a6147054879e39334843267f53fea2a25bf04407ba.
//
// Solidity: event Timelock(address indexed user, uint256 timestamp, uint8 dataType, bytes32 data)
func (_BarnFacet *BarnFacetFilterer) ParseTimelock(log types.Log) (*BarnFacetTimelock, error) {
	event := new(BarnFacetTimelock)
	if err := _BarnFacet.contract.UnpackLog(event, "Timelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BarnFacetUnlockTimelockIterator is returned from FilterUnlockTimelock and is used to iterate over the raw logs and unpacked data for UnlockTimelock events raised by the BarnFacet contract.
type BarnFacetUnlockTimelockIterator struct {
	Event *BarnFacetUnlockTimelock // Event containing the contract specifics and raw log

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
func (it *BarnFacetUnlockTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BarnFacetUnlockTimelock)
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
		it.Event = new(BarnFacetUnlockTimelock)
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
func (it *BarnFacetUnlockTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BarnFacetUnlockTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BarnFacetUnlockTimelock represents a UnlockTimelock event raised by the BarnFacet contract.
type BarnFacetUnlockTimelock struct {
	User     common.Address
	DataType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnlockTimelock is a free log retrieval operation binding the contract event 0x6112f4eb95779372a95ff6c4cadf034aa699e547223fd4e615dd4746ce36fc19.
//
// Solidity: event UnlockTimelock(address indexed user, uint8 dataType)
func (_BarnFacet *BarnFacetFilterer) FilterUnlockTimelock(opts *bind.FilterOpts, user []common.Address) (*BarnFacetUnlockTimelockIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BarnFacet.contract.FilterLogs(opts, "UnlockTimelock", userRule)
	if err != nil {
		return nil, err
	}
	return &BarnFacetUnlockTimelockIterator{contract: _BarnFacet.contract, event: "UnlockTimelock", logs: logs, sub: sub}, nil
}

// WatchUnlockTimelock is a free log subscription operation binding the contract event 0x6112f4eb95779372a95ff6c4cadf034aa699e547223fd4e615dd4746ce36fc19.
//
// Solidity: event UnlockTimelock(address indexed user, uint8 dataType)
func (_BarnFacet *BarnFacetFilterer) WatchUnlockTimelock(opts *bind.WatchOpts, sink chan<- *BarnFacetUnlockTimelock, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BarnFacet.contract.WatchLogs(opts, "UnlockTimelock", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BarnFacetUnlockTimelock)
				if err := _BarnFacet.contract.UnpackLog(event, "UnlockTimelock", log); err != nil {
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

// ParseUnlockTimelock is a log parse operation binding the contract event 0x6112f4eb95779372a95ff6c4cadf034aa699e547223fd4e615dd4746ce36fc19.
//
// Solidity: event UnlockTimelock(address indexed user, uint8 dataType)
func (_BarnFacet *BarnFacetFilterer) ParseUnlockTimelock(log types.Log) (*BarnFacetUnlockTimelock, error) {
	event := new(BarnFacetUnlockTimelock)
	if err := _BarnFacet.contract.UnpackLog(event, "UnlockTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BarnFacetWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BarnFacet contract.
type BarnFacetWithdrawIterator struct {
	Event *BarnFacetWithdraw // Event containing the contract specifics and raw log

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
func (it *BarnFacetWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BarnFacetWithdraw)
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
		it.Event = new(BarnFacetWithdraw)
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
func (it *BarnFacetWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BarnFacetWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BarnFacetWithdraw represents a Withdraw event raised by the BarnFacet contract.
type BarnFacetWithdraw struct {
	User           common.Address
	AmountWithdrew *big.Int
	AmountLeft     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amountWithdrew, uint256 amountLeft)
func (_BarnFacet *BarnFacetFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*BarnFacetWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BarnFacet.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &BarnFacetWithdrawIterator{contract: _BarnFacet.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amountWithdrew, uint256 amountLeft)
func (_BarnFacet *BarnFacetFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BarnFacetWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BarnFacet.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BarnFacetWithdraw)
				if err := _BarnFacet.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 amountWithdrew, uint256 amountLeft)
func (_BarnFacet *BarnFacetFilterer) ParseWithdraw(log types.Log) (*BarnFacetWithdraw, error) {
	event := new(BarnFacetWithdraw)
	if err := _BarnFacet.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
