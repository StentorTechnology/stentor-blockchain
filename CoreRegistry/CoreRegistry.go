// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package CoreRegistry

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RegistrarABI is the input ABI used to generate the binding from.
const RegistrarABI = `[{"constant":false,"inputs":[{"name":"newCluster","type":"address"},{"name":"name","type":"string"}],"name":"register","outputs":[{"name":"","type":"bool"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"name","type":"string"},{"name":"newaddress","type":"address"}],"name":"updateAddress","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"name","type":"string"}],"name":"getCluster","outputs":[{"name":"","type":"address"}],"payable":false,"type":"function"}]`

// Registrar is an auto generated Go binding around an Ethereum contract.
type Registrar struct {
	RegistrarCaller     // Read-only binding to the contract
	RegistrarTransactor // Write-only binding to the contract
}

// RegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrarSession struct {
	Contract     *Registrar        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrarCallerSession struct {
	Contract *RegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrarTransactorSession struct {
	Contract     *RegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrarRaw struct {
	Contract *Registrar // Generic contract binding to access the raw methods on
}

// RegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrarCallerRaw struct {
	Contract *RegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrarTransactorRaw struct {
	Contract *RegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrar creates a new instance of Registrar, bound to a specific deployed contract.
func NewRegistrar(address common.Address, backend bind.ContractBackend) (*Registrar, error) {
	contract, err := bindRegistrar(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registrar{RegistrarCaller: RegistrarCaller{contract: contract}, RegistrarTransactor: RegistrarTransactor{contract: contract}}, nil
}

// NewRegistrarCaller creates a new read-only instance of Registrar, bound to a specific deployed contract.
func NewRegistrarCaller(address common.Address, caller bind.ContractCaller) (*RegistrarCaller, error) {
	contract, err := bindRegistrar(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrarCaller{contract: contract}, nil
}

// NewRegistrarTransactor creates a new write-only instance of Registrar, bound to a specific deployed contract.
func NewRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrarTransactor, error) {
	contract, err := bindRegistrar(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RegistrarTransactor{contract: contract}, nil
}

// bindRegistrar binds a generic wrapper to an already deployed contract.
func bindRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrar *RegistrarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registrar.Contract.RegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrar *RegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.Contract.RegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrar *RegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrar.Contract.RegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registrar *RegistrarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registrar *RegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registrar *RegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registrar.Contract.contract.Transact(opts, method, params...)
}

// GetCluster is a free data retrieval call binding the contract method 0x94845a35.
//
// Solidity: function getCluster(name string) constant returns(address)
func (_Registrar *RegistrarCaller) GetCluster(opts *bind.CallOpts, name string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registrar.contract.Call(opts, out, "getCluster", name)
	return *ret0, err
}

// GetCluster is a free data retrieval call binding the contract method 0x94845a35.
//
// Solidity: function getCluster(name string) constant returns(address)
func (_Registrar *RegistrarSession) GetCluster(name string) (common.Address, error) {
	return _Registrar.Contract.GetCluster(&_Registrar.CallOpts, name)
}

// GetCluster is a free data retrieval call binding the contract method 0x94845a35.
//
// Solidity: function getCluster(name string) constant returns(address)
func (_Registrar *RegistrarCallerSession) GetCluster(name string) (common.Address, error) {
	return _Registrar.Contract.GetCluster(&_Registrar.CallOpts, name)
}

// Register is a paid mutator transaction binding the contract method 0x32434a2e.
//
// Solidity: function register(newCluster address, name string) returns(bool)
func (_Registrar *RegistrarTransactor) Register(opts *bind.TransactOpts, newCluster common.Address, name string) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "register", newCluster, name)
}

// Register is a paid mutator transaction binding the contract method 0x32434a2e.
//
// Solidity: function register(newCluster address, name string) returns(bool)
func (_Registrar *RegistrarSession) Register(newCluster common.Address, name string) (*types.Transaction, error) {
	return _Registrar.Contract.Register(&_Registrar.TransactOpts, newCluster, name)
}

// Register is a paid mutator transaction binding the contract method 0x32434a2e.
//
// Solidity: function register(newCluster address, name string) returns(bool)
func (_Registrar *RegistrarTransactorSession) Register(newCluster common.Address, name string) (*types.Transaction, error) {
	return _Registrar.Contract.Register(&_Registrar.TransactOpts, newCluster, name)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x3e21750a.
//
// Solidity: function updateAddress(name string, newaddress address) returns(address)
func (_Registrar *RegistrarTransactor) UpdateAddress(opts *bind.TransactOpts, name string, newaddress common.Address) (*types.Transaction, error) {
	return _Registrar.contract.Transact(opts, "updateAddress", name, newaddress)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x3e21750a.
//
// Solidity: function updateAddress(name string, newaddress address) returns(address)
func (_Registrar *RegistrarSession) UpdateAddress(name string, newaddress common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.UpdateAddress(&_Registrar.TransactOpts, name, newaddress)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0x3e21750a.
//
// Solidity: function updateAddress(name string, newaddress address) returns(address)
func (_Registrar *RegistrarTransactorSession) UpdateAddress(name string, newaddress common.Address) (*types.Transaction, error) {
	return _Registrar.Contract.UpdateAddress(&_Registrar.TransactOpts, name, newaddress)
}
