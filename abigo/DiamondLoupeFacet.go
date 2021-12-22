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

// IDiamondLoupeFacet is an auto generated low-level Go binding around an user-defined struct.
type IDiamondLoupeFacet struct {
	FacetAddress      common.Address
	FunctionSelectors [][4]byte
}

// DiamondLoupeFacetMetaData contains all meta data concerning the DiamondLoupeFacet contract.
var DiamondLoupeFacetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"_facetFunctionSelectors\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DiamondLoupeFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondLoupeFacetMetaData.ABI instead.
var DiamondLoupeFacetABI = DiamondLoupeFacetMetaData.ABI

// DiamondLoupeFacet is an auto generated Go binding around an Ethereum contract.
type DiamondLoupeFacet struct {
	DiamondLoupeFacetCaller     // Read-only binding to the contract
	DiamondLoupeFacetTransactor // Write-only binding to the contract
	DiamondLoupeFacetFilterer   // Log filterer for contract events
}

// DiamondLoupeFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondLoupeFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondLoupeFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondLoupeFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondLoupeFacetSession struct {
	Contract     *DiamondLoupeFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DiamondLoupeFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondLoupeFacetCallerSession struct {
	Contract *DiamondLoupeFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DiamondLoupeFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondLoupeFacetTransactorSession struct {
	Contract     *DiamondLoupeFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DiamondLoupeFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondLoupeFacetRaw struct {
	Contract *DiamondLoupeFacet // Generic contract binding to access the raw methods on
}

// DiamondLoupeFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondLoupeFacetCallerRaw struct {
	Contract *DiamondLoupeFacetCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondLoupeFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondLoupeFacetTransactorRaw struct {
	Contract *DiamondLoupeFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondLoupeFacet creates a new instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacet(address common.Address, backend bind.ContractBackend) (*DiamondLoupeFacet, error) {
	contract, err := bindDiamondLoupeFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacet{DiamondLoupeFacetCaller: DiamondLoupeFacetCaller{contract: contract}, DiamondLoupeFacetTransactor: DiamondLoupeFacetTransactor{contract: contract}, DiamondLoupeFacetFilterer: DiamondLoupeFacetFilterer{contract: contract}}, nil
}

// NewDiamondLoupeFacetCaller creates a new read-only instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacetCaller(address common.Address, caller bind.ContractCaller) (*DiamondLoupeFacetCaller, error) {
	contract, err := bindDiamondLoupeFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacetCaller{contract: contract}, nil
}

// NewDiamondLoupeFacetTransactor creates a new write-only instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondLoupeFacetTransactor, error) {
	contract, err := bindDiamondLoupeFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacetTransactor{contract: contract}, nil
}

// NewDiamondLoupeFacetFilterer creates a new log filterer instance of DiamondLoupeFacet, bound to a specific deployed contract.
func NewDiamondLoupeFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondLoupeFacetFilterer, error) {
	contract, err := bindDiamondLoupeFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFacetFilterer{contract: contract}, nil
}

// bindDiamondLoupeFacet binds a generic wrapper to an already deployed contract.
func bindDiamondLoupeFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiamondLoupeFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondLoupeFacet *DiamondLoupeFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondLoupeFacet.Contract.DiamondLoupeFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondLoupeFacet *DiamondLoupeFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.DiamondLoupeFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondLoupeFacet *DiamondLoupeFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.DiamondLoupeFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondLoupeFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondLoupeFacet *DiamondLoupeFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondLoupeFacet *DiamondLoupeFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondLoupeFacet.Contract.contract.Transact(opts, method, params...)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) FacetAddress(opts *bind.CallOpts, _functionSelector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facetAddress", _functionSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddress(&_DiamondLoupeFacet.CallOpts, _functionSelector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddress(&_DiamondLoupeFacet.CallOpts, _functionSelector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) FacetAddresses() ([]common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddresses(&_DiamondLoupeFacet.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) FacetAddresses() ([]common.Address, error) {
	return _DiamondLoupeFacet.Contract.FacetAddresses(&_DiamondLoupeFacet.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] _facetFunctionSelectors)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) FacetFunctionSelectors(opts *bind.CallOpts, _facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facetFunctionSelectors", _facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] _facetFunctionSelectors)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _DiamondLoupeFacet.Contract.FacetFunctionSelectors(&_DiamondLoupeFacet.CallOpts, _facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] _facetFunctionSelectors)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _DiamondLoupeFacet.Contract.FacetFunctionSelectors(&_DiamondLoupeFacet.CallOpts, _facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) Facets(opts *bind.CallOpts) ([]IDiamondLoupeFacet, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _DiamondLoupeFacet.Contract.Facets(&_DiamondLoupeFacet.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _DiamondLoupeFacet.Contract.Facets(&_DiamondLoupeFacet.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupeFacet *DiamondLoupeFacetCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DiamondLoupeFacet.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupeFacet *DiamondLoupeFacetSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _DiamondLoupeFacet.Contract.SupportsInterface(&_DiamondLoupeFacet.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupeFacet *DiamondLoupeFacetCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _DiamondLoupeFacet.Contract.SupportsInterface(&_DiamondLoupeFacet.CallOpts, _interfaceId)
}
