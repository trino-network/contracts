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

// TroinPaymentMetaData contains all meta data concerning the TroinPayment contract.
var TroinPaymentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"invoice_id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"invoice_id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TroinPaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use TroinPaymentMetaData.ABI instead.
var TroinPaymentABI = TroinPaymentMetaData.ABI

// TroinPayment is an auto generated Go binding around an Ethereum contract.
type TroinPayment struct {
	TroinPaymentCaller     // Read-only binding to the contract
	TroinPaymentTransactor // Write-only binding to the contract
	TroinPaymentFilterer   // Log filterer for contract events
}

// TroinPaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type TroinPaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TroinPaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TroinPaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TroinPaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TroinPaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TroinPaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TroinPaymentSession struct {
	Contract     *TroinPayment     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TroinPaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TroinPaymentCallerSession struct {
	Contract *TroinPaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TroinPaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TroinPaymentTransactorSession struct {
	Contract     *TroinPaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TroinPaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type TroinPaymentRaw struct {
	Contract *TroinPayment // Generic contract binding to access the raw methods on
}

// TroinPaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TroinPaymentCallerRaw struct {
	Contract *TroinPaymentCaller // Generic read-only contract binding to access the raw methods on
}

// TroinPaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TroinPaymentTransactorRaw struct {
	Contract *TroinPaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTroinPayment creates a new instance of TroinPayment, bound to a specific deployed contract.
func NewTroinPayment(address common.Address, backend bind.ContractBackend) (*TroinPayment, error) {
	contract, err := bindTroinPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TroinPayment{TroinPaymentCaller: TroinPaymentCaller{contract: contract}, TroinPaymentTransactor: TroinPaymentTransactor{contract: contract}, TroinPaymentFilterer: TroinPaymentFilterer{contract: contract}}, nil
}

// NewTroinPaymentCaller creates a new read-only instance of TroinPayment, bound to a specific deployed contract.
func NewTroinPaymentCaller(address common.Address, caller bind.ContractCaller) (*TroinPaymentCaller, error) {
	contract, err := bindTroinPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TroinPaymentCaller{contract: contract}, nil
}

// NewTroinPaymentTransactor creates a new write-only instance of TroinPayment, bound to a specific deployed contract.
func NewTroinPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*TroinPaymentTransactor, error) {
	contract, err := bindTroinPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TroinPaymentTransactor{contract: contract}, nil
}

// NewTroinPaymentFilterer creates a new log filterer instance of TroinPayment, bound to a specific deployed contract.
func NewTroinPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*TroinPaymentFilterer, error) {
	contract, err := bindTroinPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TroinPaymentFilterer{contract: contract}, nil
}

// bindTroinPayment binds a generic wrapper to an already deployed contract.
func bindTroinPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TroinPaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TroinPayment *TroinPaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TroinPayment.Contract.TroinPaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TroinPayment *TroinPaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TroinPayment.Contract.TroinPaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TroinPayment *TroinPaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TroinPayment.Contract.TroinPaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TroinPayment *TroinPaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TroinPayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TroinPayment *TroinPaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TroinPayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TroinPayment *TroinPaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TroinPayment.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TroinPayment *TroinPaymentCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TroinPayment.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TroinPayment *TroinPaymentSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TroinPayment.Contract.DEFAULTADMINROLE(&_TroinPayment.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TroinPayment *TroinPaymentCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TroinPayment.Contract.DEFAULTADMINROLE(&_TroinPayment.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_TroinPayment *TroinPaymentCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TroinPayment.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_TroinPayment *TroinPaymentSession) PAUSERROLE() ([32]byte, error) {
	return _TroinPayment.Contract.PAUSERROLE(&_TroinPayment.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_TroinPayment *TroinPaymentCallerSession) PAUSERROLE() ([32]byte, error) {
	return _TroinPayment.Contract.PAUSERROLE(&_TroinPayment.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TroinPayment *TroinPaymentCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TroinPayment.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TroinPayment *TroinPaymentSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TroinPayment.Contract.GetRoleAdmin(&_TroinPayment.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TroinPayment *TroinPaymentCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TroinPayment.Contract.GetRoleAdmin(&_TroinPayment.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TroinPayment *TroinPaymentCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TroinPayment.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TroinPayment *TroinPaymentSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TroinPayment.Contract.HasRole(&_TroinPayment.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TroinPayment *TroinPaymentCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TroinPayment.Contract.HasRole(&_TroinPayment.CallOpts, role, account)
}

// Orders is a free data retrieval call binding the contract method 0x1a948947.
//
// Solidity: function orders(string ) view returns(uint256)
func (_TroinPayment *TroinPaymentCaller) Orders(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _TroinPayment.contract.Call(opts, &out, "orders", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Orders is a free data retrieval call binding the contract method 0x1a948947.
//
// Solidity: function orders(string ) view returns(uint256)
func (_TroinPayment *TroinPaymentSession) Orders(arg0 string) (*big.Int, error) {
	return _TroinPayment.Contract.Orders(&_TroinPayment.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0x1a948947.
//
// Solidity: function orders(string ) view returns(uint256)
func (_TroinPayment *TroinPaymentCallerSession) Orders(arg0 string) (*big.Int, error) {
	return _TroinPayment.Contract.Orders(&_TroinPayment.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TroinPayment *TroinPaymentCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TroinPayment.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TroinPayment *TroinPaymentSession) Paused() (bool, error) {
	return _TroinPayment.Contract.Paused(&_TroinPayment.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TroinPayment *TroinPaymentCallerSession) Paused() (bool, error) {
	return _TroinPayment.Contract.Paused(&_TroinPayment.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TroinPayment *TroinPaymentCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TroinPayment.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TroinPayment *TroinPaymentSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TroinPayment.Contract.SupportsInterface(&_TroinPayment.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TroinPayment *TroinPaymentCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TroinPayment.Contract.SupportsInterface(&_TroinPayment.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TroinPayment *TroinPaymentTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinPayment.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TroinPayment *TroinPaymentSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinPayment.Contract.GrantRole(&_TroinPayment.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TroinPayment *TroinPaymentTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinPayment.Contract.GrantRole(&_TroinPayment.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _token) returns()
func (_TroinPayment *TroinPaymentTransactor) Initialize(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _TroinPayment.contract.Transact(opts, "initialize", _token)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _token) returns()
func (_TroinPayment *TroinPaymentSession) Initialize(_token common.Address) (*types.Transaction, error) {
	return _TroinPayment.Contract.Initialize(&_TroinPayment.TransactOpts, _token)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _token) returns()
func (_TroinPayment *TroinPaymentTransactorSession) Initialize(_token common.Address) (*types.Transaction, error) {
	return _TroinPayment.Contract.Initialize(&_TroinPayment.TransactOpts, _token)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TroinPayment *TroinPaymentTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TroinPayment.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TroinPayment *TroinPaymentSession) Pause() (*types.Transaction, error) {
	return _TroinPayment.Contract.Pause(&_TroinPayment.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TroinPayment *TroinPaymentTransactorSession) Pause() (*types.Transaction, error) {
	return _TroinPayment.Contract.Pause(&_TroinPayment.TransactOpts)
}

// Pay is a paid mutator transaction binding the contract method 0x0a90e179.
//
// Solidity: function pay(string invoice_id, address recipient, uint256 amount) returns()
func (_TroinPayment *TroinPaymentTransactor) Pay(opts *bind.TransactOpts, invoice_id string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TroinPayment.contract.Transact(opts, "pay", invoice_id, recipient, amount)
}

// Pay is a paid mutator transaction binding the contract method 0x0a90e179.
//
// Solidity: function pay(string invoice_id, address recipient, uint256 amount) returns()
func (_TroinPayment *TroinPaymentSession) Pay(invoice_id string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TroinPayment.Contract.Pay(&_TroinPayment.TransactOpts, invoice_id, recipient, amount)
}

// Pay is a paid mutator transaction binding the contract method 0x0a90e179.
//
// Solidity: function pay(string invoice_id, address recipient, uint256 amount) returns()
func (_TroinPayment *TroinPaymentTransactorSession) Pay(invoice_id string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TroinPayment.Contract.Pay(&_TroinPayment.TransactOpts, invoice_id, recipient, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TroinPayment *TroinPaymentTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinPayment.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TroinPayment *TroinPaymentSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinPayment.Contract.RenounceRole(&_TroinPayment.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TroinPayment *TroinPaymentTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinPayment.Contract.RenounceRole(&_TroinPayment.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TroinPayment *TroinPaymentTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinPayment.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TroinPayment *TroinPaymentSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinPayment.Contract.RevokeRole(&_TroinPayment.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TroinPayment *TroinPaymentTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TroinPayment.Contract.RevokeRole(&_TroinPayment.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TroinPayment *TroinPaymentTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TroinPayment.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TroinPayment *TroinPaymentSession) Unpause() (*types.Transaction, error) {
	return _TroinPayment.Contract.Unpause(&_TroinPayment.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TroinPayment *TroinPaymentTransactorSession) Unpause() (*types.Transaction, error) {
	return _TroinPayment.Contract.Unpause(&_TroinPayment.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_TroinPayment *TroinPaymentTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _TroinPayment.contract.Transact(opts, "withdraw", recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_TroinPayment *TroinPaymentSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _TroinPayment.Contract.Withdraw(&_TroinPayment.TransactOpts, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_TroinPayment *TroinPaymentTransactorSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _TroinPayment.Contract.Withdraw(&_TroinPayment.TransactOpts, recipient)
}

// TroinPaymentInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TroinPayment contract.
type TroinPaymentInitializedIterator struct {
	Event *TroinPaymentInitialized // Event containing the contract specifics and raw log

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
func (it *TroinPaymentInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinPaymentInitialized)
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
		it.Event = new(TroinPaymentInitialized)
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
func (it *TroinPaymentInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinPaymentInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinPaymentInitialized represents a Initialized event raised by the TroinPayment contract.
type TroinPaymentInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TroinPayment *TroinPaymentFilterer) FilterInitialized(opts *bind.FilterOpts) (*TroinPaymentInitializedIterator, error) {

	logs, sub, err := _TroinPayment.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TroinPaymentInitializedIterator{contract: _TroinPayment.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TroinPayment *TroinPaymentFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TroinPaymentInitialized) (event.Subscription, error) {

	logs, sub, err := _TroinPayment.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinPaymentInitialized)
				if err := _TroinPayment.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TroinPayment *TroinPaymentFilterer) ParseInitialized(log types.Log) (*TroinPaymentInitialized, error) {
	event := new(TroinPaymentInitialized)
	if err := _TroinPayment.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinPaymentPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TroinPayment contract.
type TroinPaymentPausedIterator struct {
	Event *TroinPaymentPaused // Event containing the contract specifics and raw log

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
func (it *TroinPaymentPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinPaymentPaused)
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
		it.Event = new(TroinPaymentPaused)
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
func (it *TroinPaymentPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinPaymentPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinPaymentPaused represents a Paused event raised by the TroinPayment contract.
type TroinPaymentPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TroinPayment *TroinPaymentFilterer) FilterPaused(opts *bind.FilterOpts) (*TroinPaymentPausedIterator, error) {

	logs, sub, err := _TroinPayment.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TroinPaymentPausedIterator{contract: _TroinPayment.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TroinPayment *TroinPaymentFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TroinPaymentPaused) (event.Subscription, error) {

	logs, sub, err := _TroinPayment.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinPaymentPaused)
				if err := _TroinPayment.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_TroinPayment *TroinPaymentFilterer) ParsePaused(log types.Log) (*TroinPaymentPaused, error) {
	event := new(TroinPaymentPaused)
	if err := _TroinPayment.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinPaymentPayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the TroinPayment contract.
type TroinPaymentPayIterator struct {
	Event *TroinPaymentPay // Event containing the contract specifics and raw log

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
func (it *TroinPaymentPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinPaymentPay)
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
		it.Event = new(TroinPaymentPay)
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
func (it *TroinPaymentPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinPaymentPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinPaymentPay represents a Pay event raised by the TroinPayment contract.
type TroinPaymentPay struct {
	InvoiceId string
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0xe804362d0950a73915a0e63be10f642d8dbcb58d8beabcf788ab015c17642e60.
//
// Solidity: event Pay(string invoice_id, address recipient, uint256 amount)
func (_TroinPayment *TroinPaymentFilterer) FilterPay(opts *bind.FilterOpts) (*TroinPaymentPayIterator, error) {

	logs, sub, err := _TroinPayment.contract.FilterLogs(opts, "Pay")
	if err != nil {
		return nil, err
	}
	return &TroinPaymentPayIterator{contract: _TroinPayment.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0xe804362d0950a73915a0e63be10f642d8dbcb58d8beabcf788ab015c17642e60.
//
// Solidity: event Pay(string invoice_id, address recipient, uint256 amount)
func (_TroinPayment *TroinPaymentFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *TroinPaymentPay) (event.Subscription, error) {

	logs, sub, err := _TroinPayment.contract.WatchLogs(opts, "Pay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinPaymentPay)
				if err := _TroinPayment.contract.UnpackLog(event, "Pay", log); err != nil {
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

// ParsePay is a log parse operation binding the contract event 0xe804362d0950a73915a0e63be10f642d8dbcb58d8beabcf788ab015c17642e60.
//
// Solidity: event Pay(string invoice_id, address recipient, uint256 amount)
func (_TroinPayment *TroinPaymentFilterer) ParsePay(log types.Log) (*TroinPaymentPay, error) {
	event := new(TroinPaymentPay)
	if err := _TroinPayment.contract.UnpackLog(event, "Pay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinPaymentRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TroinPayment contract.
type TroinPaymentRoleAdminChangedIterator struct {
	Event *TroinPaymentRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TroinPaymentRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinPaymentRoleAdminChanged)
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
		it.Event = new(TroinPaymentRoleAdminChanged)
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
func (it *TroinPaymentRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinPaymentRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinPaymentRoleAdminChanged represents a RoleAdminChanged event raised by the TroinPayment contract.
type TroinPaymentRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TroinPayment *TroinPaymentFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TroinPaymentRoleAdminChangedIterator, error) {

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

	logs, sub, err := _TroinPayment.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TroinPaymentRoleAdminChangedIterator{contract: _TroinPayment.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TroinPayment *TroinPaymentFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TroinPaymentRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TroinPayment.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinPaymentRoleAdminChanged)
				if err := _TroinPayment.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TroinPayment *TroinPaymentFilterer) ParseRoleAdminChanged(log types.Log) (*TroinPaymentRoleAdminChanged, error) {
	event := new(TroinPaymentRoleAdminChanged)
	if err := _TroinPayment.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinPaymentRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TroinPayment contract.
type TroinPaymentRoleGrantedIterator struct {
	Event *TroinPaymentRoleGranted // Event containing the contract specifics and raw log

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
func (it *TroinPaymentRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinPaymentRoleGranted)
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
		it.Event = new(TroinPaymentRoleGranted)
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
func (it *TroinPaymentRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinPaymentRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinPaymentRoleGranted represents a RoleGranted event raised by the TroinPayment contract.
type TroinPaymentRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TroinPayment *TroinPaymentFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TroinPaymentRoleGrantedIterator, error) {

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

	logs, sub, err := _TroinPayment.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TroinPaymentRoleGrantedIterator{contract: _TroinPayment.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TroinPayment *TroinPaymentFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TroinPaymentRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TroinPayment.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinPaymentRoleGranted)
				if err := _TroinPayment.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TroinPayment *TroinPaymentFilterer) ParseRoleGranted(log types.Log) (*TroinPaymentRoleGranted, error) {
	event := new(TroinPaymentRoleGranted)
	if err := _TroinPayment.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinPaymentRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TroinPayment contract.
type TroinPaymentRoleRevokedIterator struct {
	Event *TroinPaymentRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TroinPaymentRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinPaymentRoleRevoked)
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
		it.Event = new(TroinPaymentRoleRevoked)
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
func (it *TroinPaymentRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinPaymentRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinPaymentRoleRevoked represents a RoleRevoked event raised by the TroinPayment contract.
type TroinPaymentRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TroinPayment *TroinPaymentFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TroinPaymentRoleRevokedIterator, error) {

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

	logs, sub, err := _TroinPayment.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TroinPaymentRoleRevokedIterator{contract: _TroinPayment.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TroinPayment *TroinPaymentFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TroinPaymentRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TroinPayment.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinPaymentRoleRevoked)
				if err := _TroinPayment.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TroinPayment *TroinPaymentFilterer) ParseRoleRevoked(log types.Log) (*TroinPaymentRoleRevoked, error) {
	event := new(TroinPaymentRoleRevoked)
	if err := _TroinPayment.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinPaymentUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TroinPayment contract.
type TroinPaymentUnpausedIterator struct {
	Event *TroinPaymentUnpaused // Event containing the contract specifics and raw log

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
func (it *TroinPaymentUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinPaymentUnpaused)
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
		it.Event = new(TroinPaymentUnpaused)
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
func (it *TroinPaymentUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinPaymentUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinPaymentUnpaused represents a Unpaused event raised by the TroinPayment contract.
type TroinPaymentUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TroinPayment *TroinPaymentFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TroinPaymentUnpausedIterator, error) {

	logs, sub, err := _TroinPayment.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TroinPaymentUnpausedIterator{contract: _TroinPayment.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TroinPayment *TroinPaymentFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TroinPaymentUnpaused) (event.Subscription, error) {

	logs, sub, err := _TroinPayment.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinPaymentUnpaused)
				if err := _TroinPayment.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_TroinPayment *TroinPaymentFilterer) ParseUnpaused(log types.Log) (*TroinPaymentUnpaused, error) {
	event := new(TroinPaymentUnpaused)
	if err := _TroinPayment.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TroinPaymentWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the TroinPayment contract.
type TroinPaymentWithdrawIterator struct {
	Event *TroinPaymentWithdraw // Event containing the contract specifics and raw log

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
func (it *TroinPaymentWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TroinPaymentWithdraw)
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
		it.Event = new(TroinPaymentWithdraw)
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
func (it *TroinPaymentWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TroinPaymentWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TroinPaymentWithdraw represents a Withdraw event raised by the TroinPayment contract.
type TroinPaymentWithdraw struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address recipient, uint256 amount)
func (_TroinPayment *TroinPaymentFilterer) FilterWithdraw(opts *bind.FilterOpts) (*TroinPaymentWithdrawIterator, error) {

	logs, sub, err := _TroinPayment.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &TroinPaymentWithdrawIterator{contract: _TroinPayment.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address recipient, uint256 amount)
func (_TroinPayment *TroinPaymentFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TroinPaymentWithdraw) (event.Subscription, error) {

	logs, sub, err := _TroinPayment.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TroinPaymentWithdraw)
				if err := _TroinPayment.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address recipient, uint256 amount)
func (_TroinPayment *TroinPaymentFilterer) ParseWithdraw(log types.Log) (*TroinPaymentWithdraw, error) {
	event := new(TroinPaymentWithdraw)
	if err := _TroinPayment.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
