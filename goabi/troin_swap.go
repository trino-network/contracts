// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goabi

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

// TroinSwapMetaData contains all meta data concerning the TroinSwap contract.
var TroinSwapMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spend\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"SwapWithETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spend\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"SwapWithUSDT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"enumTrinoSwap.SwapMode\",\"name\":\"_mode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mode\",\"outputs\":[{\"internalType\":\"enumTrinoSwap.SwapMode\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTrinoSwap.SwapMode\",\"name\":\"_mode\",\"type\":\"uint8\"}],\"name\":\"setMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"swapWithETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"swapWithUSDT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TroinSwapABI is the input ABI used to generate the binding from.
// Deprecated: Use TroinSwapMetaData.ABI instead.
var TroinSwapABI = TroinSwapMetaData.ABI

// TroinSwap is an auto generated Go binding around an Ethereum contract.
type TroinSwap struct {
	TroinSwapCaller     // Read-only binding to the contract
	TroinSwapTransactor // Write-only binding to the contract
	TroinSwapFilterer   // Log filterer for contract events
}

// TroinSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type TroinSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TroinSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TroinSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TroinSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TroinSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TroinSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TroinSwapSession struct {
	Contract     *TroinSwap        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TroinSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TroinSwapCallerSession struct {
	Contract *TroinSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TroinSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TroinSwapTransactorSession struct {
	Contract     *TroinSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TroinSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type TroinSwapRaw struct {
	Contract *TroinSwap // Generic contract binding to access the raw methods on
}

// TroinSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TroinSwapCallerRaw struct {
	Contract *TroinSwapCaller // Generic read-only contract binding to access the raw methods on
}

// TroinSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TroinSwapTransactorRaw struct {
	Contract *TroinSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTroinSwap creates a new instance of TroinSwap, bound to a specific deployed contract.
func NewTroinSwap(address common.Address, backend bind.ContractBackend) (*TroinSwap, error) {
	contract, err := bindTroinSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TroinSwap{TroinSwapCaller: TroinSwapCaller{contract: contract}, TroinSwapTransactor: TroinSwapTransactor{contract: contract}, TroinSwapFilterer: TroinSwapFilterer{contract: contract}}, nil
}

// NewTroinSwapCaller creates a new read-only instance of TroinSwap, bound to a specific deployed contract.
func NewTroinSwapCaller(address common.Address, caller bind.ContractCaller) (*TroinSwapCaller, error) {
	contract, err := bindTroinSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TroinSwapCaller{contract: contract}, nil
}

// NewTroinSwapTransactor creates a new write-only instance of TroinSwap, bound to a specific deployed contract.
func NewTroinSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*TroinSwapTransactor, error) {
	contract, err := bindTroinSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TroinSwapTransactor{contract: contract}, nil
}

// NewTroinSwapFilterer creates a new log filterer instance of TroinSwap, bound to a specific deployed contract.
func NewTroinSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*TroinSwapFilterer, error) {
	contract, err := bindTroinSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TroinSwapFilterer{contract: contract}, nil
}

// bindTroinSwap binds a generic wrapper to an already deployed contract.
func bindTroinSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TroinSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TroinSwap *TroinSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TroinSwap.Contract.TroinSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TroinSwap *TroinSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TroinSwap.Contract.TroinSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TroinSwap *TroinSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TroinSwap.Contract.TroinSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TroinSwap *TroinSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TroinSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TroinSwap *TroinSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TroinSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TroinSwap *TroinSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TroinSwap.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TroinSwap *TroinSwapCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TroinSwap.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TroinSwap *TroinSwapSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TroinSwap.Contract.DEFAULTADMINROLE(&_TroinSwap.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TroinSwap *TroinSwapCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TroinSwap.Contract.DEFAULTADMINROLE(&_TroinSwap.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_TroinSwap *TroinSwapCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TroinSwap.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_TroinSwap *TroinSwapSession) PAUSERROLE() ([32]byte, error) {
	return _TroinSwap.Contract.PAUSERROLE(&_TroinSwap.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_TroinSwap *TroinSwapCallerSession) PAUSERROLE() ([32]byte, error) {
	return _TroinSwap.Contract.PAUSERROLE(&_TroinSwap.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_TroinSwap *TroinSwapCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TroinSwap.contract.Call(opts, &out, "beneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_TroinSwap *TroinSwapSession) Beneficiary() (common.Address, error) {
	return _TroinSwap.Contract.Beneficiary(&_TroinSwap.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() view returns(address)
func (_TroinSwap *TroinSwapCallerSession) Beneficiary() (common.Address, error) {
	return _TroinSwap.Contract.Beneficiary(&_TroinSwap.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(int256)
func (_TroinSwap *TroinSwapCaller) GetLatestPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TroinSwap.contract.Call(opts, &out, "getLatestPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(int256)
func (_TroinSwap *TroinSwapSession) GetLatestPrice() (*big.Int, error) {
	return _TroinSwap.Contract.GetLatestPrice(&_TroinSwap.CallOpts)
}

// GetLatestPrice is a free data retrieval call binding the contract method 0x8e15f473.
//
// Solidity: function getLatestPrice() view returns(int256)
func (_TroinSwap *TroinSwapCallerSession) GetLatestPrice() (*big.Int, error) {
	return _TroinSwap.Contract.GetLatestPrice(&_TroinSwap.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TroinSwap *TroinSwapCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TroinSwap.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TroinSwap *TroinSwapSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TroinSwap.Contract.GetRoleAdmin(&_TroinSwap.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TroinSwap *TroinSwapCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TroinSwap.Contract.GetRoleAdmin(&_TroinSwap.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TroinSwap *TroinSwapCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TroinSwap.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TroinSwap *TroinSwapSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TroinSwap.Contract.HasRole(&_TroinSwap.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TroinSwap *TroinSwapCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TroinSwap.Contract.HasRole(&_TroinSwap.CallOpts, role, account)
}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(uint8)
func (_TroinSwap *TroinSwapCaller) Mode(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TroinSwap.contract.Call(opts, &out, "mode")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(uint8)
func (_TroinSwap *TroinSwapSession) Mode() (uint8, error) {
	return _TroinSwap.Contract.Mode(&_TroinSwap.CallOpts)
}

// Mode is a free data retrieval call binding the contract method 0x295a5212.
//
// Solidity: function mode() view returns(uint8)
func (_TroinSwap *TroinSwapCallerSession) Mode() (uint8, error) {
	return _TroinSwap.Contract.Mode(&_TroinSwap.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TroinSwap *TroinSwapCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TroinSwap.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TroinSwap *TroinSwapSession) Paused() (bool, error) {
	return _TroinSwap.Contract.Paused(&_TroinSwap.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TroinSwap *TroinSwapCallerSession) Paused() (bool, error) {
	return _TroinSwap.Contract.Paused(&_TroinSwap.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_TroinSwap *TroinSwapCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TroinSwap.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_TroinSwap *TroinSwapSession) Rate() (*big.Int, error) {
	return _TroinSwap.Contract.Rate(&_TroinSwap.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_TroinSwap *TroinSwapCallerSession) Rate() (*big.Int, error) {
	return _TroinSwap.Contract.Rate(&_TroinSwap.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TroinSwap *TroinSwapCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TroinSwap.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TroinSwap *TroinSwapSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TroinSwap.Contract.SupportsInterface(&_TroinSwap.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TroinSwap *TroinSwapCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TroinSwap.Contract.SupportsInterface(&_TroinSwap.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TroinSwap *TroinSwapTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinSwap.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TroinSwap *TroinSwapSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinSwap.Contract.GrantRole(&_TroinSwap.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TroinSwap *TroinSwapTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinSwap.Contract.GrantRole(&_TroinSwap.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xd616d923.
//
// Solidity: function initialize(address _token, address _usdt, address _oracle, uint8 _mode, uint256 _rate) returns()
func (_TroinSwap *TroinSwapTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _usdt common.Address, _oracle common.Address, _mode uint8, _rate *big.Int) (*types.Transaction, error) {
	return _TroinSwap.contract.Transact(opts, "initialize", _token, _usdt, _oracle, _mode, _rate)
}

// Initialize is a paid mutator transaction binding the contract method 0xd616d923.
//
// Solidity: function initialize(address _token, address _usdt, address _oracle, uint8 _mode, uint256 _rate) returns()
func (_TroinSwap *TroinSwapSession) Initialize(_token common.Address, _usdt common.Address, _oracle common.Address, _mode uint8, _rate *big.Int) (*types.Transaction, error) {
	return _TroinSwap.Contract.Initialize(&_TroinSwap.TransactOpts, _token, _usdt, _oracle, _mode, _rate)
}

// Initialize is a paid mutator transaction binding the contract method 0xd616d923.
//
// Solidity: function initialize(address _token, address _usdt, address _oracle, uint8 _mode, uint256 _rate) returns()
func (_TroinSwap *TroinSwapTransactorSession) Initialize(_token common.Address, _usdt common.Address, _oracle common.Address, _mode uint8, _rate *big.Int) (*types.Transaction, error) {
	return _TroinSwap.Contract.Initialize(&_TroinSwap.TransactOpts, _token, _usdt, _oracle, _mode, _rate)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TroinSwap *TroinSwapTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TroinSwap.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TroinSwap *TroinSwapSession) Pause() (*types.Transaction, error) {
	return _TroinSwap.Contract.Pause(&_TroinSwap.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TroinSwap *TroinSwapTransactorSession) Pause() (*types.Transaction, error) {
	return _TroinSwap.Contract.Pause(&_TroinSwap.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TroinSwap *TroinSwapTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinSwap.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TroinSwap *TroinSwapSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinSwap.Contract.RenounceRole(&_TroinSwap.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TroinSwap *TroinSwapTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinSwap.Contract.RenounceRole(&_TroinSwap.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TroinSwap *TroinSwapTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinSwap.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TroinSwap *TroinSwapSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinSwap.Contract.RevokeRole(&_TroinSwap.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TroinSwap *TroinSwapTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinSwap.Contract.RevokeRole(&_TroinSwap.TransactOpts, role, account)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address _beneficiary) returns()
func (_TroinSwap *TroinSwapTransactor) SetBeneficiary(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _TroinSwap.contract.Transact(opts, "setBeneficiary", _beneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address _beneficiary) returns()
func (_TroinSwap *TroinSwapSession) SetBeneficiary(_beneficiary common.Address) (*types.Transaction, error) {
	return _TroinSwap.Contract.SetBeneficiary(&_TroinSwap.TransactOpts, _beneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address _beneficiary) returns()
func (_TroinSwap *TroinSwapTransactorSession) SetBeneficiary(_beneficiary common.Address) (*types.Transaction, error) {
	return _TroinSwap.Contract.SetBeneficiary(&_TroinSwap.TransactOpts, _beneficiary)
}

// SetMode is a paid mutator transaction binding the contract method 0x21175b4a.
//
// Solidity: function setMode(uint8 _mode) returns()
func (_TroinSwap *TroinSwapTransactor) SetMode(opts *bind.TransactOpts, _mode uint8) (*types.Transaction, error) {
	return _TroinSwap.contract.Transact(opts, "setMode", _mode)
}

// SetMode is a paid mutator transaction binding the contract method 0x21175b4a.
//
// Solidity: function setMode(uint8 _mode) returns()
func (_TroinSwap *TroinSwapSession) SetMode(_mode uint8) (*types.Transaction, error) {
	return _TroinSwap.Contract.SetMode(&_TroinSwap.TransactOpts, _mode)
}

// SetMode is a paid mutator transaction binding the contract method 0x21175b4a.
//
// Solidity: function setMode(uint8 _mode) returns()
func (_TroinSwap *TroinSwapTransactorSession) SetMode(_mode uint8) (*types.Transaction, error) {
	return _TroinSwap.Contract.SetMode(&_TroinSwap.TransactOpts, _mode)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _rate) returns()
func (_TroinSwap *TroinSwapTransactor) SetRate(opts *bind.TransactOpts, _rate *big.Int) (*types.Transaction, error) {
	return _TroinSwap.contract.Transact(opts, "setRate", _rate)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _rate) returns()
func (_TroinSwap *TroinSwapSession) SetRate(_rate *big.Int) (*types.Transaction, error) {
	return _TroinSwap.Contract.SetRate(&_TroinSwap.TransactOpts, _rate)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _rate) returns()
func (_TroinSwap *TroinSwapTransactorSession) SetRate(_rate *big.Int) (*types.Transaction, error) {
	return _TroinSwap.Contract.SetRate(&_TroinSwap.TransactOpts, _rate)
}

// SwapWithETH is a paid mutator transaction binding the contract method 0x7f9be7f8.
//
// Solidity: function swapWithETH(uint256 amount, address recipient) payable returns()
func (_TroinSwap *TroinSwapTransactor) SwapWithETH(opts *bind.TransactOpts, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _TroinSwap.contract.Transact(opts, "swapWithETH", amount, recipient)
}

// SwapWithETH is a paid mutator transaction binding the contract method 0x7f9be7f8.
//
// Solidity: function swapWithETH(uint256 amount, address recipient) payable returns()
func (_TroinSwap *TroinSwapSession) SwapWithETH(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _TroinSwap.Contract.SwapWithETH(&_TroinSwap.TransactOpts, amount, recipient)
}

// SwapWithETH is a paid mutator transaction binding the contract method 0x7f9be7f8.
//
// Solidity: function swapWithETH(uint256 amount, address recipient) payable returns()
func (_TroinSwap *TroinSwapTransactorSession) SwapWithETH(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _TroinSwap.Contract.SwapWithETH(&_TroinSwap.TransactOpts, amount, recipient)
}

// SwapWithUSDT is a paid mutator transaction binding the contract method 0xa40a6dca.
//
// Solidity: function swapWithUSDT(uint256 amount, address recipient) returns()
func (_TroinSwap *TroinSwapTransactor) SwapWithUSDT(opts *bind.TransactOpts, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _TroinSwap.contract.Transact(opts, "swapWithUSDT", amount, recipient)
}

// SwapWithUSDT is a paid mutator transaction binding the contract method 0xa40a6dca.
//
// Solidity: function swapWithUSDT(uint256 amount, address recipient) returns()
func (_TroinSwap *TroinSwapSession) SwapWithUSDT(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _TroinSwap.Contract.SwapWithUSDT(&_TroinSwap.TransactOpts, amount, recipient)
}

// SwapWithUSDT is a paid mutator transaction binding the contract method 0xa40a6dca.
//
// Solidity: function swapWithUSDT(uint256 amount, address recipient) returns()
func (_TroinSwap *TroinSwapTransactorSession) SwapWithUSDT(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _TroinSwap.Contract.SwapWithUSDT(&_TroinSwap.TransactOpts, amount, recipient)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TroinSwap *TroinSwapTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TroinSwap.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TroinSwap *TroinSwapSession) Unpause() (*types.Transaction, error) {
	return _TroinSwap.Contract.Unpause(&_TroinSwap.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TroinSwap *TroinSwapTransactorSession) Unpause() (*types.Transaction, error) {
	return _TroinSwap.Contract.Unpause(&_TroinSwap.TransactOpts)
}

// TroinSwapInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TroinSwap contract.
type TroinSwapInitializedIterator struct {
	Event *TroinSwapInitialized // Event containing the contract specifics and raw log

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
func (it *TroinSwapInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinSwapInitialized)
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
		it.Event = new(TroinSwapInitialized)
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
func (it *TroinSwapInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinSwapInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinSwapInitialized represents a Initialized event raised by the TroinSwap contract.
type TroinSwapInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TroinSwap *TroinSwapFilterer) FilterInitialized(opts *bind.FilterOpts) (*TroinSwapInitializedIterator, error) {

	logs, sub, err := _TroinSwap.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TroinSwapInitializedIterator{contract: _TroinSwap.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TroinSwap *TroinSwapFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TroinSwapInitialized) (event.Subscription, error) {

	logs, sub, err := _TroinSwap.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinSwapInitialized)
				if err := _TroinSwap.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TroinSwap *TroinSwapFilterer) ParseInitialized(log types.Log) (*TroinSwapInitialized, error) {
	event := new(TroinSwapInitialized)
	if err := _TroinSwap.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinSwapPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TroinSwap contract.
type TroinSwapPausedIterator struct {
	Event *TroinSwapPaused // Event containing the contract specifics and raw log

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
func (it *TroinSwapPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinSwapPaused)
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
		it.Event = new(TroinSwapPaused)
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
func (it *TroinSwapPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinSwapPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinSwapPaused represents a Paused event raised by the TroinSwap contract.
type TroinSwapPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TroinSwap *TroinSwapFilterer) FilterPaused(opts *bind.FilterOpts) (*TroinSwapPausedIterator, error) {

	logs, sub, err := _TroinSwap.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TroinSwapPausedIterator{contract: _TroinSwap.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TroinSwap *TroinSwapFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TroinSwapPaused) (event.Subscription, error) {

	logs, sub, err := _TroinSwap.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinSwapPaused)
				if err := _TroinSwap.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TroinSwap *TroinSwapFilterer) ParsePaused(log types.Log) (*TroinSwapPaused, error) {
	event := new(TroinSwapPaused)
	if err := _TroinSwap.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinSwapRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TroinSwap contract.
type TroinSwapRoleAdminChangedIterator struct {
	Event *TroinSwapRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TroinSwapRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinSwapRoleAdminChanged)
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
		it.Event = new(TroinSwapRoleAdminChanged)
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
func (it *TroinSwapRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinSwapRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinSwapRoleAdminChanged represents a RoleAdminChanged event raised by the TroinSwap contract.
type TroinSwapRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TroinSwap *TroinSwapFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TroinSwapRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TroinSwap.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TroinSwapRoleAdminChangedIterator{contract: _TroinSwap.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TroinSwap *TroinSwapFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TroinSwapRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TroinSwap.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinSwapRoleAdminChanged)
				if err := _TroinSwap.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TroinSwap *TroinSwapFilterer) ParseRoleAdminChanged(log types.Log) (*TroinSwapRoleAdminChanged, error) {
	event := new(TroinSwapRoleAdminChanged)
	if err := _TroinSwap.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinSwapRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TroinSwap contract.
type TroinSwapRoleGrantedIterator struct {
	Event *TroinSwapRoleGranted // Event containing the contract specifics and raw log

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
func (it *TroinSwapRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinSwapRoleGranted)
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
		it.Event = new(TroinSwapRoleGranted)
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
func (it *TroinSwapRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinSwapRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinSwapRoleGranted represents a RoleGranted event raised by the TroinSwap contract.
type TroinSwapRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TroinSwap *TroinSwapFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TroinSwapRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TroinSwap.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TroinSwapRoleGrantedIterator{contract: _TroinSwap.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TroinSwap *TroinSwapFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TroinSwapRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TroinSwap.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinSwapRoleGranted)
				if err := _TroinSwap.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TroinSwap *TroinSwapFilterer) ParseRoleGranted(log types.Log) (*TroinSwapRoleGranted, error) {
	event := new(TroinSwapRoleGranted)
	if err := _TroinSwap.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinSwapRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TroinSwap contract.
type TroinSwapRoleRevokedIterator struct {
	Event *TroinSwapRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TroinSwapRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinSwapRoleRevoked)
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
		it.Event = new(TroinSwapRoleRevoked)
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
func (it *TroinSwapRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinSwapRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinSwapRoleRevoked represents a RoleRevoked event raised by the TroinSwap contract.
type TroinSwapRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TroinSwap *TroinSwapFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TroinSwapRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TroinSwap.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TroinSwapRoleRevokedIterator{contract: _TroinSwap.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TroinSwap *TroinSwapFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TroinSwapRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TroinSwap.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinSwapRoleRevoked)
				if err := _TroinSwap.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TroinSwap *TroinSwapFilterer) ParseRoleRevoked(log types.Log) (*TroinSwapRoleRevoked, error) {
	event := new(TroinSwapRoleRevoked)
	if err := _TroinSwap.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinSwapSwapWithETHIterator is returned from FilterSwapWithETH and is used to iterate over the raw logs and unpacked data for SwapWithETH events raised by the TroinSwap contract.
type TroinSwapSwapWithETHIterator struct {
	Event *TroinSwapSwapWithETH // Event containing the contract specifics and raw log

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
func (it *TroinSwapSwapWithETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinSwapSwapWithETH)
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
		it.Event = new(TroinSwapSwapWithETH)
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
func (it *TroinSwapSwapWithETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinSwapSwapWithETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinSwapSwapWithETH represents a SwapWithETH event raised by the TroinSwap contract.
type TroinSwapSwapWithETH struct {
	Amount    *big.Int
	Spend     *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapWithETH is a free log retrieval operation binding the contract event 0x5c9d6f1ac5c04962bf07e6e38c4dd148e39e9d97b14b5ac8600eebed83013809.
//
// Solidity: event SwapWithETH(uint256 amount, uint256 spend, address recipient)
func (_TroinSwap *TroinSwapFilterer) FilterSwapWithETH(opts *bind.FilterOpts) (*TroinSwapSwapWithETHIterator, error) {

	logs, sub, err := _TroinSwap.contract.FilterLogs(opts, "SwapWithETH")
	if err != nil {
		return nil, err
	}
	return &TroinSwapSwapWithETHIterator{contract: _TroinSwap.contract, event: "SwapWithETH", logs: logs, sub: sub}, nil
}

// WatchSwapWithETH is a free log subscription operation binding the contract event 0x5c9d6f1ac5c04962bf07e6e38c4dd148e39e9d97b14b5ac8600eebed83013809.
//
// Solidity: event SwapWithETH(uint256 amount, uint256 spend, address recipient)
func (_TroinSwap *TroinSwapFilterer) WatchSwapWithETH(opts *bind.WatchOpts, sink chan<- *TroinSwapSwapWithETH) (event.Subscription, error) {

	logs, sub, err := _TroinSwap.contract.WatchLogs(opts, "SwapWithETH")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinSwapSwapWithETH)
				if err := _TroinSwap.contract.UnpackLog(event, "SwapWithETH", log); err != nil {
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

// ParseSwapWithETH is a log parse operation binding the contract event 0x5c9d6f1ac5c04962bf07e6e38c4dd148e39e9d97b14b5ac8600eebed83013809.
//
// Solidity: event SwapWithETH(uint256 amount, uint256 spend, address recipient)
func (_TroinSwap *TroinSwapFilterer) ParseSwapWithETH(log types.Log) (*TroinSwapSwapWithETH, error) {
	event := new(TroinSwapSwapWithETH)
	if err := _TroinSwap.contract.UnpackLog(event, "SwapWithETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinSwapSwapWithUSDTIterator is returned from FilterSwapWithUSDT and is used to iterate over the raw logs and unpacked data for SwapWithUSDT events raised by the TroinSwap contract.
type TroinSwapSwapWithUSDTIterator struct {
	Event *TroinSwapSwapWithUSDT // Event containing the contract specifics and raw log

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
func (it *TroinSwapSwapWithUSDTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinSwapSwapWithUSDT)
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
		it.Event = new(TroinSwapSwapWithUSDT)
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
func (it *TroinSwapSwapWithUSDTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinSwapSwapWithUSDTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinSwapSwapWithUSDT represents a SwapWithUSDT event raised by the TroinSwap contract.
type TroinSwapSwapWithUSDT struct {
	Amount    *big.Int
	Spend     *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapWithUSDT is a free log retrieval operation binding the contract event 0x5684976c8baaeb722611941a7d1c52c858db1740f79ab7437c46545648c8bbb4.
//
// Solidity: event SwapWithUSDT(uint256 amount, uint256 spend, address recipient)
func (_TroinSwap *TroinSwapFilterer) FilterSwapWithUSDT(opts *bind.FilterOpts) (*TroinSwapSwapWithUSDTIterator, error) {

	logs, sub, err := _TroinSwap.contract.FilterLogs(opts, "SwapWithUSDT")
	if err != nil {
		return nil, err
	}
	return &TroinSwapSwapWithUSDTIterator{contract: _TroinSwap.contract, event: "SwapWithUSDT", logs: logs, sub: sub}, nil
}

// WatchSwapWithUSDT is a free log subscription operation binding the contract event 0x5684976c8baaeb722611941a7d1c52c858db1740f79ab7437c46545648c8bbb4.
//
// Solidity: event SwapWithUSDT(uint256 amount, uint256 spend, address recipient)
func (_TroinSwap *TroinSwapFilterer) WatchSwapWithUSDT(opts *bind.WatchOpts, sink chan<- *TroinSwapSwapWithUSDT) (event.Subscription, error) {

	logs, sub, err := _TroinSwap.contract.WatchLogs(opts, "SwapWithUSDT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinSwapSwapWithUSDT)
				if err := _TroinSwap.contract.UnpackLog(event, "SwapWithUSDT", log); err != nil {
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

// ParseSwapWithUSDT is a log parse operation binding the contract event 0x5684976c8baaeb722611941a7d1c52c858db1740f79ab7437c46545648c8bbb4.
//
// Solidity: event SwapWithUSDT(uint256 amount, uint256 spend, address recipient)
func (_TroinSwap *TroinSwapFilterer) ParseSwapWithUSDT(log types.Log) (*TroinSwapSwapWithUSDT, error) {
	event := new(TroinSwapSwapWithUSDT)
	if err := _TroinSwap.contract.UnpackLog(event, "SwapWithUSDT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinSwapUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TroinSwap contract.
type TroinSwapUnpausedIterator struct {
	Event *TroinSwapUnpaused // Event containing the contract specifics and raw log

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
func (it *TroinSwapUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinSwapUnpaused)
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
		it.Event = new(TroinSwapUnpaused)
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
func (it *TroinSwapUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinSwapUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinSwapUnpaused represents a Unpaused event raised by the TroinSwap contract.
type TroinSwapUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TroinSwap *TroinSwapFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TroinSwapUnpausedIterator, error) {

	logs, sub, err := _TroinSwap.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TroinSwapUnpausedIterator{contract: _TroinSwap.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TroinSwap *TroinSwapFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TroinSwapUnpaused) (event.Subscription, error) {

	logs, sub, err := _TroinSwap.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinSwapUnpaused)
				if err := _TroinSwap.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TroinSwap *TroinSwapFilterer) ParseUnpaused(log types.Log) (*TroinSwapUnpaused, error) {
	event := new(TroinSwapUnpaused)
	if err := _TroinSwap.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
