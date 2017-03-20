// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package geth

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PoolMonitorClientABI is the input ABI used to generate the binding from.
const PoolMonitorClientABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"version\",\"type\":\"bytes32\"}],\"name\":\"updateClientVersion\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"clientVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"poolContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"updatePoolContract\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[3]\"}],\"payable\":false,\"type\":\"constructor\"}]"

// PoolMonitorClient is an auto generated Go binding around an Ethereum contract.
type PoolMonitorClient struct {
	PoolMonitorClientCaller     // Read-only binding to the contract
	PoolMonitorClientTransactor // Write-only binding to the contract
}

// PoolMonitorClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolMonitorClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolMonitorClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolMonitorClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolMonitorClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolMonitorClientSession struct {
	Contract     *PoolMonitorClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PoolMonitorClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolMonitorClientCallerSession struct {
	Contract *PoolMonitorClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PoolMonitorClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolMonitorClientTransactorSession struct {
	Contract     *PoolMonitorClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PoolMonitorClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolMonitorClientRaw struct {
	Contract *PoolMonitorClient // Generic contract binding to access the raw methods on
}

// PoolMonitorClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolMonitorClientCallerRaw struct {
	Contract *PoolMonitorClientCaller // Generic read-only contract binding to access the raw methods on
}

// PoolMonitorClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolMonitorClientTransactorRaw struct {
	Contract *PoolMonitorClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolMonitorClient creates a new instance of PoolMonitorClient, bound to a specific deployed contract.
func NewPoolMonitorClient(address common.Address, backend bind.ContractBackend) (*PoolMonitorClient, error) {
	contract, err := bindPoolMonitorClient(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolMonitorClient{PoolMonitorClientCaller: PoolMonitorClientCaller{contract: contract}, PoolMonitorClientTransactor: PoolMonitorClientTransactor{contract: contract}}, nil
}

// NewPoolMonitorClientCaller creates a new read-only instance of PoolMonitorClient, bound to a specific deployed contract.
func NewPoolMonitorClientCaller(address common.Address, caller bind.ContractCaller) (*PoolMonitorClientCaller, error) {
	contract, err := bindPoolMonitorClient(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PoolMonitorClientCaller{contract: contract}, nil
}

// NewPoolMonitorClientTransactor creates a new write-only instance of PoolMonitorClient, bound to a specific deployed contract.
func NewPoolMonitorClientTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolMonitorClientTransactor, error) {
	contract, err := bindPoolMonitorClient(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PoolMonitorClientTransactor{contract: contract}, nil
}

// bindPoolMonitorClient binds a generic wrapper to an already deployed contract.
func bindPoolMonitorClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolMonitorClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolMonitorClient *PoolMonitorClientRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PoolMonitorClient.Contract.PoolMonitorClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolMonitorClient *PoolMonitorClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolMonitorClient.Contract.PoolMonitorClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolMonitorClient *PoolMonitorClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolMonitorClient.Contract.PoolMonitorClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolMonitorClient *PoolMonitorClientCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PoolMonitorClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolMonitorClient *PoolMonitorClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolMonitorClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolMonitorClient *PoolMonitorClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolMonitorClient.Contract.contract.Transact(opts, method, params...)
}

// ClientVersion is a free data retrieval call binding the contract method 0x044c1d4d.
//
// Solidity: function clientVersion() constant returns(bytes32)
func (_PoolMonitorClient *PoolMonitorClientCaller) ClientVersion(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PoolMonitorClient.contract.Call(opts, out, "clientVersion")
	return *ret0, err
}

// ClientVersion is a free data retrieval call binding the contract method 0x044c1d4d.
//
// Solidity: function clientVersion() constant returns(bytes32)
func (_PoolMonitorClient *PoolMonitorClientSession) ClientVersion() ([32]byte, error) {
	return _PoolMonitorClient.Contract.ClientVersion(&_PoolMonitorClient.CallOpts)
}

// ClientVersion is a free data retrieval call binding the contract method 0x044c1d4d.
//
// Solidity: function clientVersion() constant returns(bytes32)
func (_PoolMonitorClient *PoolMonitorClientCallerSession) ClientVersion() ([32]byte, error) {
	return _PoolMonitorClient.Contract.ClientVersion(&_PoolMonitorClient.CallOpts)
}

// PoolContract is a free data retrieval call binding the contract method 0x88d52ef7.
//
// Solidity: function poolContract() constant returns(address)
func (_PoolMonitorClient *PoolMonitorClientCaller) PoolContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PoolMonitorClient.contract.Call(opts, out, "poolContract")
	return *ret0, err
}

// PoolContract is a free data retrieval call binding the contract method 0x88d52ef7.
//
// Solidity: function poolContract() constant returns(address)
func (_PoolMonitorClient *PoolMonitorClientSession) PoolContract() (common.Address, error) {
	return _PoolMonitorClient.Contract.PoolContract(&_PoolMonitorClient.CallOpts)
}

// PoolContract is a free data retrieval call binding the contract method 0x88d52ef7.
//
// Solidity: function poolContract() constant returns(address)
func (_PoolMonitorClient *PoolMonitorClientCallerSession) PoolContract() (common.Address, error) {
	return _PoolMonitorClient.Contract.PoolContract(&_PoolMonitorClient.CallOpts)
}

// UpdateClientVersion is a paid mutator transaction binding the contract method 0x03e9598c.
//
// Solidity: function updateClientVersion(version bytes32) returns()
func (_PoolMonitorClient *PoolMonitorClientTransactor) UpdateClientVersion(opts *bind.TransactOpts, version [32]byte) (*types.Transaction, error) {
	return _PoolMonitorClient.contract.Transact(opts, "updateClientVersion", version)
}

// UpdateClientVersion is a paid mutator transaction binding the contract method 0x03e9598c.
//
// Solidity: function updateClientVersion(version bytes32) returns()
func (_PoolMonitorClient *PoolMonitorClientSession) UpdateClientVersion(version [32]byte) (*types.Transaction, error) {
	return _PoolMonitorClient.Contract.UpdateClientVersion(&_PoolMonitorClient.TransactOpts, version)
}

// UpdateClientVersion is a paid mutator transaction binding the contract method 0x03e9598c.
//
// Solidity: function updateClientVersion(version bytes32) returns()
func (_PoolMonitorClient *PoolMonitorClientTransactorSession) UpdateClientVersion(version [32]byte) (*types.Transaction, error) {
	return _PoolMonitorClient.Contract.UpdateClientVersion(&_PoolMonitorClient.TransactOpts, version)
}

// UpdatePoolContract is a paid mutator transaction binding the contract method 0xff1d9dd4.
//
// Solidity: function updatePoolContract(newAddress address) returns()
func (_PoolMonitorClient *PoolMonitorClientTransactor) UpdatePoolContract(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _PoolMonitorClient.contract.Transact(opts, "updatePoolContract", newAddress)
}

// UpdatePoolContract is a paid mutator transaction binding the contract method 0xff1d9dd4.
//
// Solidity: function updatePoolContract(newAddress address) returns()
func (_PoolMonitorClient *PoolMonitorClientSession) UpdatePoolContract(newAddress common.Address) (*types.Transaction, error) {
	return _PoolMonitorClient.Contract.UpdatePoolContract(&_PoolMonitorClient.TransactOpts, newAddress)
}

// UpdatePoolContract is a paid mutator transaction binding the contract method 0xff1d9dd4.
//
// Solidity: function updatePoolContract(newAddress address) returns()
func (_PoolMonitorClient *PoolMonitorClientTransactorSession) UpdatePoolContract(newAddress common.Address) (*types.Transaction, error) {
	return _PoolMonitorClient.Contract.UpdatePoolContract(&_PoolMonitorClient.TransactOpts, newAddress)
}
