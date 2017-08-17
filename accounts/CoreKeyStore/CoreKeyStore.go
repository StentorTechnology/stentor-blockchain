// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package CoreKeyStore

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// KeyStoreABI is the input ABI used to generate the binding from.
const KeyStoreABI = `[{"constant":false,"inputs":[{"name":"cluster","type":"address"},{"name":"account","type":"address"},{"name":"keyjson","type":"string"}],"name":"storeKey","outputs":[],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"cluster","type":"address"},{"name":"account","type":"address"}],"name":"getKey","outputs":[{"name":"","type":"string"}],"payable":false,"type":"function"}]`

// KeyStore is an auto generated Go binding around an Ethereum contract.
type KeyStore struct {
	KeyStoreCaller     // Read-only binding to the contract
	KeyStoreTransactor // Write-only binding to the contract
}

// KeyStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyStoreSession struct {
	Contract     *KeyStore         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyStoreCallerSession struct {
	Contract *KeyStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// KeyStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyStoreTransactorSession struct {
	Contract     *KeyStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KeyStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyStoreRaw struct {
	Contract *KeyStore // Generic contract binding to access the raw methods on
}

// KeyStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyStoreCallerRaw struct {
	Contract *KeyStoreCaller // Generic read-only contract binding to access the raw methods on
}

// KeyStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyStoreTransactorRaw struct {
	Contract *KeyStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyStore creates a new instance of KeyStore, bound to a specific deployed contract.
func NewKeyStore(address common.Address, backend bind.ContractBackend) (*KeyStore, error) {
	contract, err := bindKeyStore(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyStore{KeyStoreCaller: KeyStoreCaller{contract: contract}, KeyStoreTransactor: KeyStoreTransactor{contract: contract}}, nil
}

// NewKeyStoreCaller creates a new read-only instance of KeyStore, bound to a specific deployed contract.
func NewKeyStoreCaller(address common.Address, caller bind.ContractCaller) (*KeyStoreCaller, error) {
	contract, err := bindKeyStore(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &KeyStoreCaller{contract: contract}, nil
}

// NewKeyStoreTransactor creates a new write-only instance of KeyStore, bound to a specific deployed contract.
func NewKeyStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyStoreTransactor, error) {
	contract, err := bindKeyStore(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &KeyStoreTransactor{contract: contract}, nil
}

// bindKeyStore binds a generic wrapper to an already deployed contract.
func bindKeyStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyStore *KeyStoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KeyStore.Contract.KeyStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyStore *KeyStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyStore.Contract.KeyStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyStore *KeyStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyStore.Contract.KeyStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyStore *KeyStoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KeyStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyStore *KeyStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyStore *KeyStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyStore.Contract.contract.Transact(opts, method, params...)
}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(cluster address, account address) constant returns(string)
func (_KeyStore *KeyStoreCaller) GetKey(opts *bind.CallOpts, cluster common.Address, account common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _KeyStore.contract.Call(opts, out, "getKey", cluster, account)
	return *ret0, err
}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(cluster address, account address) constant returns(string)
func (_KeyStore *KeyStoreSession) GetKey(cluster common.Address, account common.Address) (string, error) {
	return _KeyStore.Contract.GetKey(&_KeyStore.CallOpts, cluster, account)
}

// GetKey is a free data retrieval call binding the contract method 0xbc298553.
//
// Solidity: function getKey(cluster address, account address) constant returns(string)
func (_KeyStore *KeyStoreCallerSession) GetKey(cluster common.Address, account common.Address) (string, error) {
	return _KeyStore.Contract.GetKey(&_KeyStore.CallOpts, cluster, account)
}

// StoreKey is a paid mutator transaction binding the contract method 0x1d6191d3.
//
// Solidity: function storeKey(cluster address, account address, keyjson string) returns()
func (_KeyStore *KeyStoreTransactor) StoreKey(opts *bind.TransactOpts, cluster common.Address, account common.Address, keyjson string) (*types.Transaction, error) {
	return _KeyStore.contract.Transact(opts, "storeKey", cluster, account, keyjson)
}

// StoreKey is a paid mutator transaction binding the contract method 0x1d6191d3.
//
// Solidity: function storeKey(cluster address, account address, keyjson string) returns()
func (_KeyStore *KeyStoreSession) StoreKey(cluster common.Address, account common.Address, keyjson string) (*types.Transaction, error) {
	return _KeyStore.Contract.StoreKey(&_KeyStore.TransactOpts, cluster, account, keyjson)
}

// StoreKey is a paid mutator transaction binding the contract method 0x1d6191d3.
//
// Solidity: function storeKey(cluster address, account address, keyjson string) returns()
func (_KeyStore *KeyStoreTransactorSession) StoreKey(cluster common.Address, account common.Address, keyjson string) (*types.Transaction, error) {
	return _KeyStore.Contract.StoreKey(&_KeyStore.TransactOpts, cluster, account, keyjson)
}
