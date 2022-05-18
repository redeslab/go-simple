// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethapi

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

// AdvertiseAdItem is an auto generated low-level Go binding around an user-defined struct.
type AdvertiseAdItem struct {
	Name         string
	ConfigInJson string
}

// AdvertiseMetaData contains all meta data concerning the Advertise contract.
var AdvertiseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdList\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"configInJson\",\"type\":\"string\"}],\"internalType\":\"structAdvertise.AdItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Administrators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"QueryByOne\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"config\",\"type\":\"string\"}],\"name\":\"addItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"advertisements\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"configInJson\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"config\",\"type\":\"string\"}],\"name\":\"changeServer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"removeItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"sIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AdvertiseABI is the input ABI used to generate the binding from.
// Deprecated: Use AdvertiseMetaData.ABI instead.
var AdvertiseABI = AdvertiseMetaData.ABI

// Advertise is an auto generated Go binding around an Ethereum contract.
type Advertise struct {
	AdvertiseCaller     // Read-only binding to the contract
	AdvertiseTransactor // Write-only binding to the contract
	AdvertiseFilterer   // Log filterer for contract events
}

// AdvertiseCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdvertiseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvertiseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdvertiseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvertiseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdvertiseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvertiseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdvertiseSession struct {
	Contract     *Advertise        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdvertiseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdvertiseCallerSession struct {
	Contract *AdvertiseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AdvertiseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdvertiseTransactorSession struct {
	Contract     *AdvertiseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AdvertiseRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdvertiseRaw struct {
	Contract *Advertise // Generic contract binding to access the raw methods on
}

// AdvertiseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdvertiseCallerRaw struct {
	Contract *AdvertiseCaller // Generic read-only contract binding to access the raw methods on
}

// AdvertiseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdvertiseTransactorRaw struct {
	Contract *AdvertiseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdvertise creates a new instance of Advertise, bound to a specific deployed contract.
func NewAdvertise(address common.Address, backend bind.ContractBackend) (*Advertise, error) {
	contract, err := bindAdvertise(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Advertise{AdvertiseCaller: AdvertiseCaller{contract: contract}, AdvertiseTransactor: AdvertiseTransactor{contract: contract}, AdvertiseFilterer: AdvertiseFilterer{contract: contract}}, nil
}

// NewAdvertiseCaller creates a new read-only instance of Advertise, bound to a specific deployed contract.
func NewAdvertiseCaller(address common.Address, caller bind.ContractCaller) (*AdvertiseCaller, error) {
	contract, err := bindAdvertise(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdvertiseCaller{contract: contract}, nil
}

// NewAdvertiseTransactor creates a new write-only instance of Advertise, bound to a specific deployed contract.
func NewAdvertiseTransactor(address common.Address, transactor bind.ContractTransactor) (*AdvertiseTransactor, error) {
	contract, err := bindAdvertise(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdvertiseTransactor{contract: contract}, nil
}

// NewAdvertiseFilterer creates a new log filterer instance of Advertise, bound to a specific deployed contract.
func NewAdvertiseFilterer(address common.Address, filterer bind.ContractFilterer) (*AdvertiseFilterer, error) {
	contract, err := bindAdvertise(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdvertiseFilterer{contract: contract}, nil
}

// bindAdvertise binds a generic wrapper to an already deployed contract.
func bindAdvertise(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdvertiseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Advertise *AdvertiseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Advertise.Contract.AdvertiseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Advertise *AdvertiseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Advertise.Contract.AdvertiseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Advertise *AdvertiseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Advertise.Contract.AdvertiseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Advertise *AdvertiseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Advertise.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Advertise *AdvertiseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Advertise.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Advertise *AdvertiseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Advertise.Contract.contract.Transact(opts, method, params...)
}

// AdList is a free data retrieval call binding the contract method 0x9c2e1981.
//
// Solidity: function AdList() view returns((string,string)[])
func (_Advertise *AdvertiseCaller) AdList(opts *bind.CallOpts) ([]AdvertiseAdItem, error) {
	var out []interface{}
	err := _Advertise.contract.Call(opts, &out, "AdList")

	if err != nil {
		return *new([]AdvertiseAdItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]AdvertiseAdItem)).(*[]AdvertiseAdItem)

	return out0, err

}

// AdList is a free data retrieval call binding the contract method 0x9c2e1981.
//
// Solidity: function AdList() view returns((string,string)[])
func (_Advertise *AdvertiseSession) AdList() ([]AdvertiseAdItem, error) {
	return _Advertise.Contract.AdList(&_Advertise.CallOpts)
}

// AdList is a free data retrieval call binding the contract method 0x9c2e1981.
//
// Solidity: function AdList() view returns((string,string)[])
func (_Advertise *AdvertiseCallerSession) AdList() ([]AdvertiseAdItem, error) {
	return _Advertise.Contract.AdList(&_Advertise.CallOpts)
}

// Administrators is a free data retrieval call binding the contract method 0xf7735770.
//
// Solidity: function Administrators(address ) view returns(bool)
func (_Advertise *AdvertiseCaller) Administrators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Advertise.contract.Call(opts, &out, "Administrators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Administrators is a free data retrieval call binding the contract method 0xf7735770.
//
// Solidity: function Administrators(address ) view returns(bool)
func (_Advertise *AdvertiseSession) Administrators(arg0 common.Address) (bool, error) {
	return _Advertise.Contract.Administrators(&_Advertise.CallOpts, arg0)
}

// Administrators is a free data retrieval call binding the contract method 0xf7735770.
//
// Solidity: function Administrators(address ) view returns(bool)
func (_Advertise *AdvertiseCallerSession) Administrators(arg0 common.Address) (bool, error) {
	return _Advertise.Contract.Administrators(&_Advertise.CallOpts, arg0)
}

// QueryByOne is a free data retrieval call binding the contract method 0x52e1281f.
//
// Solidity: function QueryByOne(string name) view returns(string)
func (_Advertise *AdvertiseCaller) QueryByOne(opts *bind.CallOpts, name string) (string, error) {
	var out []interface{}
	err := _Advertise.contract.Call(opts, &out, "QueryByOne", name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// QueryByOne is a free data retrieval call binding the contract method 0x52e1281f.
//
// Solidity: function QueryByOne(string name) view returns(string)
func (_Advertise *AdvertiseSession) QueryByOne(name string) (string, error) {
	return _Advertise.Contract.QueryByOne(&_Advertise.CallOpts, name)
}

// QueryByOne is a free data retrieval call binding the contract method 0x52e1281f.
//
// Solidity: function QueryByOne(string name) view returns(string)
func (_Advertise *AdvertiseCallerSession) QueryByOne(name string) (string, error) {
	return _Advertise.Contract.QueryByOne(&_Advertise.CallOpts, name)
}

// Advertisements is a free data retrieval call binding the contract method 0xbcea610e.
//
// Solidity: function advertisements(uint256 ) view returns(string name, string configInJson)
func (_Advertise *AdvertiseCaller) Advertisements(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name         string
	ConfigInJson string
}, error) {
	var out []interface{}
	err := _Advertise.contract.Call(opts, &out, "advertisements", arg0)

	outstruct := new(struct {
		Name         string
		ConfigInJson string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ConfigInJson = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Advertisements is a free data retrieval call binding the contract method 0xbcea610e.
//
// Solidity: function advertisements(uint256 ) view returns(string name, string configInJson)
func (_Advertise *AdvertiseSession) Advertisements(arg0 *big.Int) (struct {
	Name         string
	ConfigInJson string
}, error) {
	return _Advertise.Contract.Advertisements(&_Advertise.CallOpts, arg0)
}

// Advertisements is a free data retrieval call binding the contract method 0xbcea610e.
//
// Solidity: function advertisements(uint256 ) view returns(string name, string configInJson)
func (_Advertise *AdvertiseCallerSession) Advertisements(arg0 *big.Int) (struct {
	Name         string
	ConfigInJson string
}, error) {
	return _Advertise.Contract.Advertisements(&_Advertise.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Advertise *AdvertiseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Advertise.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Advertise *AdvertiseSession) Owner() (common.Address, error) {
	return _Advertise.Contract.Owner(&_Advertise.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Advertise *AdvertiseCallerSession) Owner() (common.Address, error) {
	return _Advertise.Contract.Owner(&_Advertise.CallOpts)
}

// SIdx is a free data retrieval call binding the contract method 0xb14faf5a.
//
// Solidity: function sIdx(string ) view returns(uint256)
func (_Advertise *AdvertiseCaller) SIdx(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Advertise.contract.Call(opts, &out, "sIdx", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIdx is a free data retrieval call binding the contract method 0xb14faf5a.
//
// Solidity: function sIdx(string ) view returns(uint256)
func (_Advertise *AdvertiseSession) SIdx(arg0 string) (*big.Int, error) {
	return _Advertise.Contract.SIdx(&_Advertise.CallOpts, arg0)
}

// SIdx is a free data retrieval call binding the contract method 0xb14faf5a.
//
// Solidity: function sIdx(string ) view returns(uint256)
func (_Advertise *AdvertiseCallerSession) SIdx(arg0 string) (*big.Int, error) {
	return _Advertise.Contract.SIdx(&_Advertise.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Advertise *AdvertiseTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Advertise.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Advertise *AdvertiseSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Advertise.Contract.AddAdmin(&_Advertise.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_Advertise *AdvertiseTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Advertise.Contract.AddAdmin(&_Advertise.TransactOpts, admin)
}

// AddItem is a paid mutator transaction binding the contract method 0x7ed8a719.
//
// Solidity: function addItem(string name, string config) returns()
func (_Advertise *AdvertiseTransactor) AddItem(opts *bind.TransactOpts, name string, config string) (*types.Transaction, error) {
	return _Advertise.contract.Transact(opts, "addItem", name, config)
}

// AddItem is a paid mutator transaction binding the contract method 0x7ed8a719.
//
// Solidity: function addItem(string name, string config) returns()
func (_Advertise *AdvertiseSession) AddItem(name string, config string) (*types.Transaction, error) {
	return _Advertise.Contract.AddItem(&_Advertise.TransactOpts, name, config)
}

// AddItem is a paid mutator transaction binding the contract method 0x7ed8a719.
//
// Solidity: function addItem(string name, string config) returns()
func (_Advertise *AdvertiseTransactorSession) AddItem(name string, config string) (*types.Transaction, error) {
	return _Advertise.Contract.AddItem(&_Advertise.TransactOpts, name, config)
}

// ChangeServer is a paid mutator transaction binding the contract method 0xf5a7d9b9.
//
// Solidity: function changeServer(string name, string config) returns()
func (_Advertise *AdvertiseTransactor) ChangeServer(opts *bind.TransactOpts, name string, config string) (*types.Transaction, error) {
	return _Advertise.contract.Transact(opts, "changeServer", name, config)
}

// ChangeServer is a paid mutator transaction binding the contract method 0xf5a7d9b9.
//
// Solidity: function changeServer(string name, string config) returns()
func (_Advertise *AdvertiseSession) ChangeServer(name string, config string) (*types.Transaction, error) {
	return _Advertise.Contract.ChangeServer(&_Advertise.TransactOpts, name, config)
}

// ChangeServer is a paid mutator transaction binding the contract method 0xf5a7d9b9.
//
// Solidity: function changeServer(string name, string config) returns()
func (_Advertise *AdvertiseTransactorSession) ChangeServer(name string, config string) (*types.Transaction, error) {
	return _Advertise.Contract.ChangeServer(&_Advertise.TransactOpts, name, config)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Advertise *AdvertiseTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Advertise.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Advertise *AdvertiseSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Advertise.Contract.RemoveAdmin(&_Advertise.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_Advertise *AdvertiseTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _Advertise.Contract.RemoveAdmin(&_Advertise.TransactOpts, admin)
}

// RemoveItem is a paid mutator transaction binding the contract method 0x68dfa8ba.
//
// Solidity: function removeItem(string name) returns()
func (_Advertise *AdvertiseTransactor) RemoveItem(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Advertise.contract.Transact(opts, "removeItem", name)
}

// RemoveItem is a paid mutator transaction binding the contract method 0x68dfa8ba.
//
// Solidity: function removeItem(string name) returns()
func (_Advertise *AdvertiseSession) RemoveItem(name string) (*types.Transaction, error) {
	return _Advertise.Contract.RemoveItem(&_Advertise.TransactOpts, name)
}

// RemoveItem is a paid mutator transaction binding the contract method 0x68dfa8ba.
//
// Solidity: function removeItem(string name) returns()
func (_Advertise *AdvertiseTransactorSession) RemoveItem(name string) (*types.Transaction, error) {
	return _Advertise.Contract.RemoveItem(&_Advertise.TransactOpts, name)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Advertise *AdvertiseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Advertise.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Advertise *AdvertiseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Advertise.Contract.TransferOwnership(&_Advertise.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Advertise *AdvertiseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Advertise.Contract.TransferOwnership(&_Advertise.TransactOpts, newOwner)
}
