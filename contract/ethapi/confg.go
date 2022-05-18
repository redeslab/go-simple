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

// ConfigServerItem is an auto generated low-level Go binding around an user-defined struct.
type ConfigServerItem struct {
	Addr string
	Host string
}

// ChainConfigMetaData contains all meta data concerning the ChainConfig contract.
var ChainConfigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"ServerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Administrators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AdvertiseAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"name\":\"QueryByOne\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ServerList\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"}],\"internalType\":\"structConfig.ServerItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"}],\"name\":\"addServer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"}],\"name\":\"changeServer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"name\":\"removeServer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"sIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"servers\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAdvertisAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ChainConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use ChainConfigMetaData.ABI instead.
var ChainConfigABI = ChainConfigMetaData.ABI

// ChainConfig is an auto generated Go binding around an Ethereum contract.
type ChainConfig struct {
	ChainConfigCaller     // Read-only binding to the contract
	ChainConfigTransactor // Write-only binding to the contract
	ChainConfigFilterer   // Log filterer for contract events
}

// ChainConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainConfigSession struct {
	Contract     *ChainConfig      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainConfigCallerSession struct {
	Contract *ChainConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ChainConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainConfigTransactorSession struct {
	Contract     *ChainConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ChainConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainConfigRaw struct {
	Contract *ChainConfig // Generic contract binding to access the raw methods on
}

// ChainConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainConfigCallerRaw struct {
	Contract *ChainConfigCaller // Generic read-only contract binding to access the raw methods on
}

// ChainConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainConfigTransactorRaw struct {
	Contract *ChainConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainConfig creates a new instance of ChainConfig, bound to a specific deployed contract.
func NewChainConfig(address common.Address, backend bind.ContractBackend) (*ChainConfig, error) {
	contract, err := bindChainConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainConfig{ChainConfigCaller: ChainConfigCaller{contract: contract}, ChainConfigTransactor: ChainConfigTransactor{contract: contract}, ChainConfigFilterer: ChainConfigFilterer{contract: contract}}, nil
}

// NewChainConfigCaller creates a new read-only instance of ChainConfig, bound to a specific deployed contract.
func NewChainConfigCaller(address common.Address, caller bind.ContractCaller) (*ChainConfigCaller, error) {
	contract, err := bindChainConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainConfigCaller{contract: contract}, nil
}

// NewChainConfigTransactor creates a new write-only instance of ChainConfig, bound to a specific deployed contract.
func NewChainConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainConfigTransactor, error) {
	contract, err := bindChainConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainConfigTransactor{contract: contract}, nil
}

// NewChainConfigFilterer creates a new log filterer instance of ChainConfig, bound to a specific deployed contract.
func NewChainConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainConfigFilterer, error) {
	contract, err := bindChainConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainConfigFilterer{contract: contract}, nil
}

// bindChainConfig binds a generic wrapper to an already deployed contract.
func bindChainConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainConfig *ChainConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainConfig.Contract.ChainConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainConfig *ChainConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainConfig.Contract.ChainConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainConfig *ChainConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainConfig.Contract.ChainConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainConfig *ChainConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainConfig *ChainConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainConfig *ChainConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainConfig.Contract.contract.Transact(opts, method, params...)
}

// Administrators is a free data retrieval call binding the contract method 0xf7735770.
//
// Solidity: function Administrators(address ) view returns(bool)
func (_ChainConfig *ChainConfigCaller) Administrators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "Administrators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Administrators is a free data retrieval call binding the contract method 0xf7735770.
//
// Solidity: function Administrators(address ) view returns(bool)
func (_ChainConfig *ChainConfigSession) Administrators(arg0 common.Address) (bool, error) {
	return _ChainConfig.Contract.Administrators(&_ChainConfig.CallOpts, arg0)
}

// Administrators is a free data retrieval call binding the contract method 0xf7735770.
//
// Solidity: function Administrators(address ) view returns(bool)
func (_ChainConfig *ChainConfigCallerSession) Administrators(arg0 common.Address) (bool, error) {
	return _ChainConfig.Contract.Administrators(&_ChainConfig.CallOpts, arg0)
}

// AdvertiseAddr is a free data retrieval call binding the contract method 0x06816e43.
//
// Solidity: function AdvertiseAddr() view returns(address)
func (_ChainConfig *ChainConfigCaller) AdvertiseAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "AdvertiseAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdvertiseAddr is a free data retrieval call binding the contract method 0x06816e43.
//
// Solidity: function AdvertiseAddr() view returns(address)
func (_ChainConfig *ChainConfigSession) AdvertiseAddr() (common.Address, error) {
	return _ChainConfig.Contract.AdvertiseAddr(&_ChainConfig.CallOpts)
}

// AdvertiseAddr is a free data retrieval call binding the contract method 0x06816e43.
//
// Solidity: function AdvertiseAddr() view returns(address)
func (_ChainConfig *ChainConfigCallerSession) AdvertiseAddr() (common.Address, error) {
	return _ChainConfig.Contract.AdvertiseAddr(&_ChainConfig.CallOpts)
}

// QueryByOne is a free data retrieval call binding the contract method 0x52e1281f.
//
// Solidity: function QueryByOne(string addr) view returns(string)
func (_ChainConfig *ChainConfigCaller) QueryByOne(opts *bind.CallOpts, addr string) (string, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "QueryByOne", addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// QueryByOne is a free data retrieval call binding the contract method 0x52e1281f.
//
// Solidity: function QueryByOne(string addr) view returns(string)
func (_ChainConfig *ChainConfigSession) QueryByOne(addr string) (string, error) {
	return _ChainConfig.Contract.QueryByOne(&_ChainConfig.CallOpts, addr)
}

// QueryByOne is a free data retrieval call binding the contract method 0x52e1281f.
//
// Solidity: function QueryByOne(string addr) view returns(string)
func (_ChainConfig *ChainConfigCallerSession) QueryByOne(addr string) (string, error) {
	return _ChainConfig.Contract.QueryByOne(&_ChainConfig.CallOpts, addr)
}

// ServerList is a free data retrieval call binding the contract method 0xad013d1d.
//
// Solidity: function ServerList() view returns((string,string)[])
func (_ChainConfig *ChainConfigCaller) ServerList(opts *bind.CallOpts) ([]ConfigServerItem, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "ServerList")

	if err != nil {
		return *new([]ConfigServerItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]ConfigServerItem)).(*[]ConfigServerItem)

	return out0, err

}

// ServerList is a free data retrieval call binding the contract method 0xad013d1d.
//
// Solidity: function ServerList() view returns((string,string)[])
func (_ChainConfig *ChainConfigSession) ServerList() ([]ConfigServerItem, error) {
	return _ChainConfig.Contract.ServerList(&_ChainConfig.CallOpts)
}

// ServerList is a free data retrieval call binding the contract method 0xad013d1d.
//
// Solidity: function ServerList() view returns((string,string)[])
func (_ChainConfig *ChainConfigCallerSession) ServerList() ([]ConfigServerItem, error) {
	return _ChainConfig.Contract.ServerList(&_ChainConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainConfig *ChainConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainConfig *ChainConfigSession) Owner() (common.Address, error) {
	return _ChainConfig.Contract.Owner(&_ChainConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainConfig *ChainConfigCallerSession) Owner() (common.Address, error) {
	return _ChainConfig.Contract.Owner(&_ChainConfig.CallOpts)
}

// SIdx is a free data retrieval call binding the contract method 0xb14faf5a.
//
// Solidity: function sIdx(string ) view returns(uint256)
func (_ChainConfig *ChainConfigCaller) SIdx(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "sIdx", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIdx is a free data retrieval call binding the contract method 0xb14faf5a.
//
// Solidity: function sIdx(string ) view returns(uint256)
func (_ChainConfig *ChainConfigSession) SIdx(arg0 string) (*big.Int, error) {
	return _ChainConfig.Contract.SIdx(&_ChainConfig.CallOpts, arg0)
}

// SIdx is a free data retrieval call binding the contract method 0xb14faf5a.
//
// Solidity: function sIdx(string ) view returns(uint256)
func (_ChainConfig *ChainConfigCallerSession) SIdx(arg0 string) (*big.Int, error) {
	return _ChainConfig.Contract.SIdx(&_ChainConfig.CallOpts, arg0)
}

// Servers is a free data retrieval call binding the contract method 0x5cf0f357.
//
// Solidity: function servers(uint256 ) view returns(string addr, string host)
func (_ChainConfig *ChainConfigCaller) Servers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr string
	Host string
}, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "servers", arg0)

	outstruct := new(struct {
		Addr string
		Host string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Host = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Servers is a free data retrieval call binding the contract method 0x5cf0f357.
//
// Solidity: function servers(uint256 ) view returns(string addr, string host)
func (_ChainConfig *ChainConfigSession) Servers(arg0 *big.Int) (struct {
	Addr string
	Host string
}, error) {
	return _ChainConfig.Contract.Servers(&_ChainConfig.CallOpts, arg0)
}

// Servers is a free data retrieval call binding the contract method 0x5cf0f357.
//
// Solidity: function servers(uint256 ) view returns(string addr, string host)
func (_ChainConfig *ChainConfigCallerSession) Servers(arg0 *big.Int) (struct {
	Addr string
	Host string
}, error) {
	return _ChainConfig.Contract.Servers(&_ChainConfig.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_ChainConfig *ChainConfigTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_ChainConfig *ChainConfigSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _ChainConfig.Contract.AddAdmin(&_ChainConfig.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns()
func (_ChainConfig *ChainConfigTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _ChainConfig.Contract.AddAdmin(&_ChainConfig.TransactOpts, admin)
}

// AddServer is a paid mutator transaction binding the contract method 0x36ba41bb.
//
// Solidity: function addServer(string addr, string host) returns()
func (_ChainConfig *ChainConfigTransactor) AddServer(opts *bind.TransactOpts, addr string, host string) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "addServer", addr, host)
}

// AddServer is a paid mutator transaction binding the contract method 0x36ba41bb.
//
// Solidity: function addServer(string addr, string host) returns()
func (_ChainConfig *ChainConfigSession) AddServer(addr string, host string) (*types.Transaction, error) {
	return _ChainConfig.Contract.AddServer(&_ChainConfig.TransactOpts, addr, host)
}

// AddServer is a paid mutator transaction binding the contract method 0x36ba41bb.
//
// Solidity: function addServer(string addr, string host) returns()
func (_ChainConfig *ChainConfigTransactorSession) AddServer(addr string, host string) (*types.Transaction, error) {
	return _ChainConfig.Contract.AddServer(&_ChainConfig.TransactOpts, addr, host)
}

// ChangeServer is a paid mutator transaction binding the contract method 0xf5a7d9b9.
//
// Solidity: function changeServer(string addr, string host) returns()
func (_ChainConfig *ChainConfigTransactor) ChangeServer(opts *bind.TransactOpts, addr string, host string) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "changeServer", addr, host)
}

// ChangeServer is a paid mutator transaction binding the contract method 0xf5a7d9b9.
//
// Solidity: function changeServer(string addr, string host) returns()
func (_ChainConfig *ChainConfigSession) ChangeServer(addr string, host string) (*types.Transaction, error) {
	return _ChainConfig.Contract.ChangeServer(&_ChainConfig.TransactOpts, addr, host)
}

// ChangeServer is a paid mutator transaction binding the contract method 0xf5a7d9b9.
//
// Solidity: function changeServer(string addr, string host) returns()
func (_ChainConfig *ChainConfigTransactorSession) ChangeServer(addr string, host string) (*types.Transaction, error) {
	return _ChainConfig.Contract.ChangeServer(&_ChainConfig.TransactOpts, addr, host)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_ChainConfig *ChainConfigTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_ChainConfig *ChainConfigSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _ChainConfig.Contract.RemoveAdmin(&_ChainConfig.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_ChainConfig *ChainConfigTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _ChainConfig.Contract.RemoveAdmin(&_ChainConfig.TransactOpts, admin)
}

// RemoveServer is a paid mutator transaction binding the contract method 0x46ace8fd.
//
// Solidity: function removeServer(string addr) returns()
func (_ChainConfig *ChainConfigTransactor) RemoveServer(opts *bind.TransactOpts, addr string) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "removeServer", addr)
}

// RemoveServer is a paid mutator transaction binding the contract method 0x46ace8fd.
//
// Solidity: function removeServer(string addr) returns()
func (_ChainConfig *ChainConfigSession) RemoveServer(addr string) (*types.Transaction, error) {
	return _ChainConfig.Contract.RemoveServer(&_ChainConfig.TransactOpts, addr)
}

// RemoveServer is a paid mutator transaction binding the contract method 0x46ace8fd.
//
// Solidity: function removeServer(string addr) returns()
func (_ChainConfig *ChainConfigTransactorSession) RemoveServer(addr string) (*types.Transaction, error) {
	return _ChainConfig.Contract.RemoveServer(&_ChainConfig.TransactOpts, addr)
}

// SetAdvertisAddr is a paid mutator transaction binding the contract method 0xa8ae3090.
//
// Solidity: function setAdvertisAddr(address addr) returns()
func (_ChainConfig *ChainConfigTransactor) SetAdvertisAddr(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "setAdvertisAddr", addr)
}

// SetAdvertisAddr is a paid mutator transaction binding the contract method 0xa8ae3090.
//
// Solidity: function setAdvertisAddr(address addr) returns()
func (_ChainConfig *ChainConfigSession) SetAdvertisAddr(addr common.Address) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetAdvertisAddr(&_ChainConfig.TransactOpts, addr)
}

// SetAdvertisAddr is a paid mutator transaction binding the contract method 0xa8ae3090.
//
// Solidity: function setAdvertisAddr(address addr) returns()
func (_ChainConfig *ChainConfigTransactorSession) SetAdvertisAddr(addr common.Address) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetAdvertisAddr(&_ChainConfig.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainConfig *ChainConfigTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainConfig *ChainConfigSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChainConfig.Contract.TransferOwnership(&_ChainConfig.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ChainConfig *ChainConfigTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ChainConfig.Contract.TransferOwnership(&_ChainConfig.TransactOpts, newOwner)
}

// ChainConfigServerChangedIterator is returned from FilterServerChanged and is used to iterate over the raw logs and unpacked data for ServerChanged events raised by the ChainConfig contract.
type ChainConfigServerChangedIterator struct {
	Event *ChainConfigServerChanged // Event containing the contract specifics and raw log

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
func (it *ChainConfigServerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainConfigServerChanged)
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
		it.Event = new(ChainConfigServerChanged)
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
func (it *ChainConfigServerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainConfigServerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainConfigServerChanged represents a ServerChanged event raised by the ChainConfig contract.
type ChainConfigServerChanged struct {
	Arg0 string
	Arg1 string
	Arg2 uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterServerChanged is a free log retrieval operation binding the contract event 0xd377c3fbfc4223ee2ce4676c79f04087240dc96d7c0b31ea87f5bd5d3961b013.
//
// Solidity: event ServerChanged(string arg0, string arg1, uint8 arg2)
func (_ChainConfig *ChainConfigFilterer) FilterServerChanged(opts *bind.FilterOpts) (*ChainConfigServerChangedIterator, error) {

	logs, sub, err := _ChainConfig.contract.FilterLogs(opts, "ServerChanged")
	if err != nil {
		return nil, err
	}
	return &ChainConfigServerChangedIterator{contract: _ChainConfig.contract, event: "ServerChanged", logs: logs, sub: sub}, nil
}

// WatchServerChanged is a free log subscription operation binding the contract event 0xd377c3fbfc4223ee2ce4676c79f04087240dc96d7c0b31ea87f5bd5d3961b013.
//
// Solidity: event ServerChanged(string arg0, string arg1, uint8 arg2)
func (_ChainConfig *ChainConfigFilterer) WatchServerChanged(opts *bind.WatchOpts, sink chan<- *ChainConfigServerChanged) (event.Subscription, error) {

	logs, sub, err := _ChainConfig.contract.WatchLogs(opts, "ServerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainConfigServerChanged)
				if err := _ChainConfig.contract.UnpackLog(event, "ServerChanged", log); err != nil {
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

// ParseServerChanged is a log parse operation binding the contract event 0xd377c3fbfc4223ee2ce4676c79f04087240dc96d7c0b31ea87f5bd5d3961b013.
//
// Solidity: event ServerChanged(string arg0, string arg1, uint8 arg2)
func (_ChainConfig *ChainConfigFilterer) ParseServerChanged(log types.Log) (*ChainConfigServerChanged, error) {
	event := new(ChainConfigServerChanged)
	if err := _ChainConfig.contract.UnpackLog(event, "ServerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
