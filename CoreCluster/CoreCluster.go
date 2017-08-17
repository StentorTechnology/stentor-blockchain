// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package CoreCluster

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ClusterABI is the input ABI used to generate the binding from.
const ClusterABI = `[{"constant":true,"inputs":[],"name":"getName","outputs":[{"name":"","type":"string"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"newAddress","type":"address"}],"name":"updateAddress","outputs":[],"payable":false,"type":"function"},{"inputs":[{"name":"_name","type":"string"}],"payable":false,"type":"constructor"},{"payable":true,"type":"fallback"}]`

// ClusterBin is the compiled bytecode used for deploying new contracts.
const ClusterBin = `6060604052341561000f57600080fd5b6040516106a83803806106a8833981016040528080518201919050505b600073f683b6c648d0999b7725330adb4d4e5ea3d4849973ffffffffffffffffffffffffffffffffffffffff166332434a2e30846000604051602001526040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101075780820151818401525b6020810190506100eb565b50505050905090810190601f1680156101345780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b151561015357600080fd5b6102c65a03f1151561016457600080fd5b5050506040518051905090506001151581151514151561018357600080fd5b81600190805190602001906101999291906101e2565b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5050610287565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061022357805160ff1916838001178555610251565b82800160010185558215610251579182015b82811115610250578251825591602001919060010190610235565b5b50905061025e9190610262565b5090565b61028491905b80821115610280576000816000905550600101610268565b5090565b90565b610412806102966000396000f3006060604052361561004a576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806317d7de7c1461004e578063ed6d9096146100dd575b5b5b005b341561005957600080fd5b610061610116565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100a25780820151818401525b602081019050610086565b50505050905090810190601f1680156100cf5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156100e857600080fd5b610114600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506101bf565b005b61011e6103d2565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101b45780601f10610189576101008083540402835291602001916101b4565b820191906000526020600020905b81548152906001019060200180831161019757829003601f168201915b505050505090505b90565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561021c57600080fd5b73f683b6c648d0999b7725330adb4d4e5ea3d4849973ffffffffffffffffffffffffffffffffffffffff16633e21750a6001846000604051602001526040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182810382528481815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561033e5780601f106103135761010080835404028352916020019161033e565b820191906000526020600020905b81548152906001019060200180831161032157829003601f168201915b50509350505050602060405180830381600087803b151561035e57600080fd5b6102c65a03f1151561036f57600080fd5b5050506040518051905090508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156103b557600080fd5b8173ffffffffffffffffffffffffffffffffffffffff16ff5b5050565b6020604051908101604052806000815250905600a165627a7a72305820a8f7f9598d3d7c149abf6a2154371845bd893015251655bd4a0dfdb7aa3fd6970029`

// DeployCluster deploys a new Ethereum contract, binding an instance of Cluster to it.
func DeployCluster(auth *bind.TransactOpts, backend bind.ContractBackend, _name string) (common.Address, *types.Transaction, *Cluster, error) {
	parsed, err := abi.JSON(strings.NewReader(ClusterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClusterBin), backend, _name)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cluster{ClusterCaller: ClusterCaller{contract: contract}, ClusterTransactor: ClusterTransactor{contract: contract}}, nil
}

// Cluster is an auto generated Go binding around an Ethereum contract.
type Cluster struct {
	ClusterCaller     // Read-only binding to the contract
	ClusterTransactor // Write-only binding to the contract
}

// ClusterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClusterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClusterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClusterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClusterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClusterSession struct {
	Contract     *Cluster          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClusterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClusterCallerSession struct {
	Contract *ClusterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ClusterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClusterTransactorSession struct {
	Contract     *ClusterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ClusterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClusterRaw struct {
	Contract *Cluster // Generic contract binding to access the raw methods on
}

// ClusterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClusterCallerRaw struct {
	Contract *ClusterCaller // Generic read-only contract binding to access the raw methods on
}

// ClusterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClusterTransactorRaw struct {
	Contract *ClusterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCluster creates a new instance of Cluster, bound to a specific deployed contract.
func NewCluster(address common.Address, backend bind.ContractBackend) (*Cluster, error) {
	contract, err := bindCluster(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cluster{ClusterCaller: ClusterCaller{contract: contract}, ClusterTransactor: ClusterTransactor{contract: contract}}, nil
}

// NewClusterCaller creates a new read-only instance of Cluster, bound to a specific deployed contract.
func NewClusterCaller(address common.Address, caller bind.ContractCaller) (*ClusterCaller, error) {
	contract, err := bindCluster(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ClusterCaller{contract: contract}, nil
}

// NewClusterTransactor creates a new write-only instance of Cluster, bound to a specific deployed contract.
func NewClusterTransactor(address common.Address, transactor bind.ContractTransactor) (*ClusterTransactor, error) {
	contract, err := bindCluster(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ClusterTransactor{contract: contract}, nil
}

// bindCluster binds a generic wrapper to an already deployed contract.
func bindCluster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClusterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cluster *ClusterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cluster.Contract.ClusterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cluster *ClusterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cluster.Contract.ClusterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cluster *ClusterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cluster.Contract.ClusterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cluster *ClusterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cluster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cluster *ClusterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cluster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cluster *ClusterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cluster.Contract.contract.Transact(opts, method, params...)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_Cluster *ClusterCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Cluster.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_Cluster *ClusterSession) GetName() (string, error) {
	return _Cluster.Contract.GetName(&_Cluster.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_Cluster *ClusterCallerSession) GetName() (string, error) {
	return _Cluster.Contract.GetName(&_Cluster.CallOpts)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0xed6d9096.
//
// Solidity: function updateAddress(newAddress address) returns()
func (_Cluster *ClusterTransactor) UpdateAddress(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Cluster.contract.Transact(opts, "updateAddress", newAddress)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0xed6d9096.
//
// Solidity: function updateAddress(newAddress address) returns()
func (_Cluster *ClusterSession) UpdateAddress(newAddress common.Address) (*types.Transaction, error) {
	return _Cluster.Contract.UpdateAddress(&_Cluster.TransactOpts, newAddress)
}

// UpdateAddress is a paid mutator transaction binding the contract method 0xed6d9096.
//
// Solidity: function updateAddress(newAddress address) returns()
func (_Cluster *ClusterTransactorSession) UpdateAddress(newAddress common.Address) (*types.Transaction, error) {
	return _Cluster.Contract.UpdateAddress(&_Cluster.TransactOpts, newAddress)
}
