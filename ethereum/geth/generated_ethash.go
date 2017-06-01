// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package geth

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EthashABI is the input ABI used to generate the binding from.
const EthashABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes32\"},{\"name\":\"nonceLe\",\"type\":\"bytes8\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"},{\"name\":\"epochIndex\",\"type\":\"uint256\"}],\"name\":\"hashimoto\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"epochIndex\",\"type\":\"uint256\"},{\"name\":\"nodeIndex\",\"type\":\"uint256\"}],\"name\":\"getEpochData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[3]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"epochIndex\",\"type\":\"uint256\"}],\"name\":\"isEpochDataSet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"},{\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint256\"},{\"name\":\"branchDepth\",\"type\":\"uint256\"},{\"name\":\"merkleNodes\",\"type\":\"uint256[]\"},{\"name\":\"start\",\"type\":\"uint256\"},{\"name\":\"numElems\",\"type\":\"uint256\"}],\"name\":\"setEpochData\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[3]\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"SetEpochData\",\"type\":\"event\"}]"

// Ethash is an auto generated Go binding around an Ethereum contract.
type Ethash struct {
	EthashCaller     // Read-only binding to the contract
	EthashTransactor // Write-only binding to the contract
}

// EthashCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthashSession struct {
	Contract     *Ethash           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthashCallerSession struct {
	Contract *EthashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthashTransactorSession struct {
	Contract     *EthashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthashRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthashRaw struct {
	Contract *Ethash // Generic contract binding to access the raw methods on
}

// EthashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthashCallerRaw struct {
	Contract *EthashCaller // Generic read-only contract binding to access the raw methods on
}

// EthashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthashTransactorRaw struct {
	Contract *EthashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthash creates a new instance of Ethash, bound to a specific deployed contract.
func NewEthash(address common.Address, backend bind.ContractBackend) (*Ethash, error) {
	contract, err := bindEthash(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethash{EthashCaller: EthashCaller{contract: contract}, EthashTransactor: EthashTransactor{contract: contract}}, nil
}

// NewEthashCaller creates a new read-only instance of Ethash, bound to a specific deployed contract.
func NewEthashCaller(address common.Address, caller bind.ContractCaller) (*EthashCaller, error) {
	contract, err := bindEthash(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EthashCaller{contract: contract}, nil
}

// NewEthashTransactor creates a new write-only instance of Ethash, bound to a specific deployed contract.
func NewEthashTransactor(address common.Address, transactor bind.ContractTransactor) (*EthashTransactor, error) {
	contract, err := bindEthash(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EthashTransactor{contract: contract}, nil
}

// bindEthash binds a generic wrapper to an already deployed contract.
func bindEthash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethash *EthashRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethash.Contract.EthashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethash *EthashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethash.Contract.EthashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethash *EthashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethash.Contract.EthashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethash *EthashCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethash *EthashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethash *EthashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethash.Contract.contract.Transact(opts, method, params...)
}

// GetEpochData is a free data retrieval call binding the contract method 0x4c80c937.
//
// Solidity: function getEpochData(epochIndex uint256, nodeIndex uint256) constant returns(uint256[3])
func (_Ethash *EthashCaller) GetEpochData(opts *bind.CallOpts, epochIndex *big.Int, nodeIndex *big.Int) ([3]*big.Int, error) {
	var (
		ret0 = new([3]*big.Int)
	)
	out := ret0
	err := _Ethash.contract.Call(opts, out, "getEpochData", epochIndex, nodeIndex)
	return *ret0, err
}

// GetEpochData is a free data retrieval call binding the contract method 0x4c80c937.
//
// Solidity: function getEpochData(epochIndex uint256, nodeIndex uint256) constant returns(uint256[3])
func (_Ethash *EthashSession) GetEpochData(epochIndex *big.Int, nodeIndex *big.Int) ([3]*big.Int, error) {
	return _Ethash.Contract.GetEpochData(&_Ethash.CallOpts, epochIndex, nodeIndex)
}

// GetEpochData is a free data retrieval call binding the contract method 0x4c80c937.
//
// Solidity: function getEpochData(epochIndex uint256, nodeIndex uint256) constant returns(uint256[3])
func (_Ethash *EthashCallerSession) GetEpochData(epochIndex *big.Int, nodeIndex *big.Int) ([3]*big.Int, error) {
	return _Ethash.Contract.GetEpochData(&_Ethash.CallOpts, epochIndex, nodeIndex)
}

// Hashimoto is a free data retrieval call binding the contract method 0x40ed79f4.
//
// Solidity: function hashimoto(header bytes32, nonceLe bytes8, dataSetLookup uint256[], witnessForLookup uint256[], epochIndex uint256) constant returns(uint256)
func (_Ethash *EthashCaller) Hashimoto(opts *bind.CallOpts, header [32]byte, nonceLe [8]byte, dataSetLookup []*big.Int, witnessForLookup []*big.Int, epochIndex *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethash.contract.Call(opts, out, "hashimoto", header, nonceLe, dataSetLookup, witnessForLookup, epochIndex)
	return *ret0, err
}

// Hashimoto is a free data retrieval call binding the contract method 0x40ed79f4.
//
// Solidity: function hashimoto(header bytes32, nonceLe bytes8, dataSetLookup uint256[], witnessForLookup uint256[], epochIndex uint256) constant returns(uint256)
func (_Ethash *EthashSession) Hashimoto(header [32]byte, nonceLe [8]byte, dataSetLookup []*big.Int, witnessForLookup []*big.Int, epochIndex *big.Int) (*big.Int, error) {
	return _Ethash.Contract.Hashimoto(&_Ethash.CallOpts, header, nonceLe, dataSetLookup, witnessForLookup, epochIndex)
}

// Hashimoto is a free data retrieval call binding the contract method 0x40ed79f4.
//
// Solidity: function hashimoto(header bytes32, nonceLe bytes8, dataSetLookup uint256[], witnessForLookup uint256[], epochIndex uint256) constant returns(uint256)
func (_Ethash *EthashCallerSession) Hashimoto(header [32]byte, nonceLe [8]byte, dataSetLookup []*big.Int, witnessForLookup []*big.Int, epochIndex *big.Int) (*big.Int, error) {
	return _Ethash.Contract.Hashimoto(&_Ethash.CallOpts, header, nonceLe, dataSetLookup, witnessForLookup, epochIndex)
}

// IsEpochDataSet is a free data retrieval call binding the contract method 0xc7b81f4f.
//
// Solidity: function isEpochDataSet(epochIndex uint256) constant returns(bool)
func (_Ethash *EthashCaller) IsEpochDataSet(opts *bind.CallOpts, epochIndex *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ethash.contract.Call(opts, out, "isEpochDataSet", epochIndex)
	return *ret0, err
}

// IsEpochDataSet is a free data retrieval call binding the contract method 0xc7b81f4f.
//
// Solidity: function isEpochDataSet(epochIndex uint256) constant returns(bool)
func (_Ethash *EthashSession) IsEpochDataSet(epochIndex *big.Int) (bool, error) {
	return _Ethash.Contract.IsEpochDataSet(&_Ethash.CallOpts, epochIndex)
}

// IsEpochDataSet is a free data retrieval call binding the contract method 0xc7b81f4f.
//
// Solidity: function isEpochDataSet(epochIndex uint256) constant returns(bool)
func (_Ethash *EthashCallerSession) IsEpochDataSet(epochIndex *big.Int) (bool, error) {
	return _Ethash.Contract.IsEpochDataSet(&_Ethash.CallOpts, epochIndex)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(bool)
func (_Ethash *EthashCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ethash.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(bool)
func (_Ethash *EthashSession) Owners(arg0 common.Address) (bool, error) {
	return _Ethash.Contract.Owners(&_Ethash.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(bool)
func (_Ethash *EthashCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Ethash.Contract.Owners(&_Ethash.CallOpts, arg0)
}

// SetEpochData is a paid mutator transaction binding the contract method 0xc891a29d.
//
// Solidity: function setEpochData(epoch uint256, fullSizeIn128Resultion uint256, branchDepth uint256, merkleNodes uint256[], start uint256, numElems uint256) returns()
func (_Ethash *EthashTransactor) SetEpochData(opts *bind.TransactOpts, epoch *big.Int, fullSizeIn128Resultion *big.Int, branchDepth *big.Int, merkleNodes []*big.Int, start *big.Int, numElems *big.Int) (*types.Transaction, error) {
	return _Ethash.contract.Transact(opts, "setEpochData", epoch, fullSizeIn128Resultion, branchDepth, merkleNodes, start, numElems)
}

// SetEpochData is a paid mutator transaction binding the contract method 0xc891a29d.
//
// Solidity: function setEpochData(epoch uint256, fullSizeIn128Resultion uint256, branchDepth uint256, merkleNodes uint256[], start uint256, numElems uint256) returns()
func (_Ethash *EthashSession) SetEpochData(epoch *big.Int, fullSizeIn128Resultion *big.Int, branchDepth *big.Int, merkleNodes []*big.Int, start *big.Int, numElems *big.Int) (*types.Transaction, error) {
	return _Ethash.Contract.SetEpochData(&_Ethash.TransactOpts, epoch, fullSizeIn128Resultion, branchDepth, merkleNodes, start, numElems)
}

// SetEpochData is a paid mutator transaction binding the contract method 0xc891a29d.
//
// Solidity: function setEpochData(epoch uint256, fullSizeIn128Resultion uint256, branchDepth uint256, merkleNodes uint256[], start uint256, numElems uint256) returns()
func (_Ethash *EthashTransactorSession) SetEpochData(epoch *big.Int, fullSizeIn128Resultion *big.Int, branchDepth *big.Int, merkleNodes []*big.Int, start *big.Int, numElems *big.Int) (*types.Transaction, error) {
	return _Ethash.Contract.SetEpochData(&_Ethash.TransactOpts, epoch, fullSizeIn128Resultion, branchDepth, merkleNodes, start, numElems)
}
