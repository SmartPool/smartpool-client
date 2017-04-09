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

// TestClientABI is the input ABI used to generate the binding from.
const TestClientABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"dagIndices\",\"type\":\"uint256[]\"},{\"name\":\"dagElements\",\"type\":\"uint256[]\"},{\"name\":\"witnessArray\",\"type\":\"uint256[]\"},{\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"testOptimization\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"computeLeaf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"indexInElementsArray\",\"type\":\"uint256\"},{\"name\":\"elements\",\"type\":\"uint256[]\"},{\"name\":\"witness\",\"type\":\"uint256[]\"},{\"name\":\"branchSize\",\"type\":\"uint256\"}],\"name\":\"computeCacheRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"},{\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint256\"},{\"name\":\"branchDepth\",\"type\":\"uint256\"},{\"name\":\"merkleNodes\",\"type\":\"uint256[]\"},{\"name\":\"start\",\"type\":\"uint256\"},{\"name\":\"numElems\",\"type\":\"uint256\"}],\"name\":\"setOptEpochData\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// TestClient is an auto generated Go binding around an Ethereum contract.
type TestClient struct {
	TestClientCaller     // Read-only binding to the contract
	TestClientTransactor // Write-only binding to the contract
}

// TestClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestClientSession struct {
	Contract     *TestClient       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestClientCallerSession struct {
	Contract *TestClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestClientTransactorSession struct {
	Contract     *TestClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestClientRaw struct {
	Contract *TestClient // Generic contract binding to access the raw methods on
}

// TestClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestClientCallerRaw struct {
	Contract *TestClientCaller // Generic read-only contract binding to access the raw methods on
}

// TestClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestClientTransactorRaw struct {
	Contract *TestClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestClient creates a new instance of TestClient, bound to a specific deployed contract.
func NewTestClient(address common.Address, backend bind.ContractBackend) (*TestClient, error) {
	contract, err := bindTestClient(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestClient{TestClientCaller: TestClientCaller{contract: contract}, TestClientTransactor: TestClientTransactor{contract: contract}}, nil
}

// NewTestClientCaller creates a new read-only instance of TestClient, bound to a specific deployed contract.
func NewTestClientCaller(address common.Address, caller bind.ContractCaller) (*TestClientCaller, error) {
	contract, err := bindTestClient(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TestClientCaller{contract: contract}, nil
}

// NewTestClientTransactor creates a new write-only instance of TestClient, bound to a specific deployed contract.
func NewTestClientTransactor(address common.Address, transactor bind.ContractTransactor) (*TestClientTransactor, error) {
	contract, err := bindTestClient(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TestClientTransactor{contract: contract}, nil
}

// bindTestClient binds a generic wrapper to an already deployed contract.
func bindTestClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestClient *TestClientRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestClient.Contract.TestClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestClient *TestClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestClient.Contract.TestClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestClient *TestClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestClient.Contract.TestClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestClient *TestClientCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestClient *TestClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestClient *TestClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestClient.Contract.contract.Transact(opts, method, params...)
}

// ComputeCacheRoot is a free data retrieval call binding the contract method 0x3aa4868a.
//
// Solidity: function computeCacheRoot(index uint256, indexInElementsArray uint256, elements uint256[], witness uint256[], branchSize uint256) constant returns(uint256)
func (_TestClient *TestClientCaller) ComputeCacheRoot(opts *bind.CallOpts, index *big.Int, indexInElementsArray *big.Int, elements []*big.Int, witness []*big.Int, branchSize *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestClient.contract.Call(opts, out, "computeCacheRoot", index, indexInElementsArray, elements, witness, branchSize)
	return *ret0, err
}

// ComputeCacheRoot is a free data retrieval call binding the contract method 0x3aa4868a.
//
// Solidity: function computeCacheRoot(index uint256, indexInElementsArray uint256, elements uint256[], witness uint256[], branchSize uint256) constant returns(uint256)
func (_TestClient *TestClientSession) ComputeCacheRoot(index *big.Int, indexInElementsArray *big.Int, elements []*big.Int, witness []*big.Int, branchSize *big.Int) (*big.Int, error) {
	return _TestClient.Contract.ComputeCacheRoot(&_TestClient.CallOpts, index, indexInElementsArray, elements, witness, branchSize)
}

// ComputeCacheRoot is a free data retrieval call binding the contract method 0x3aa4868a.
//
// Solidity: function computeCacheRoot(index uint256, indexInElementsArray uint256, elements uint256[], witness uint256[], branchSize uint256) constant returns(uint256)
func (_TestClient *TestClientCallerSession) ComputeCacheRoot(index *big.Int, indexInElementsArray *big.Int, elements []*big.Int, witness []*big.Int, branchSize *big.Int) (*big.Int, error) {
	return _TestClient.Contract.ComputeCacheRoot(&_TestClient.CallOpts, index, indexInElementsArray, elements, witness, branchSize)
}

// ComputeLeaf is a free data retrieval call binding the contract method 0x275ccb13.
//
// Solidity: function computeLeaf(dataSetLookup uint256[], index uint256) constant returns(uint256)
func (_TestClient *TestClientCaller) ComputeLeaf(opts *bind.CallOpts, dataSetLookup []*big.Int, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestClient.contract.Call(opts, out, "computeLeaf", dataSetLookup, index)
	return *ret0, err
}

// ComputeLeaf is a free data retrieval call binding the contract method 0x275ccb13.
//
// Solidity: function computeLeaf(dataSetLookup uint256[], index uint256) constant returns(uint256)
func (_TestClient *TestClientSession) ComputeLeaf(dataSetLookup []*big.Int, index *big.Int) (*big.Int, error) {
	return _TestClient.Contract.ComputeLeaf(&_TestClient.CallOpts, dataSetLookup, index)
}

// ComputeLeaf is a free data retrieval call binding the contract method 0x275ccb13.
//
// Solidity: function computeLeaf(dataSetLookup uint256[], index uint256) constant returns(uint256)
func (_TestClient *TestClientCallerSession) ComputeLeaf(dataSetLookup []*big.Int, index *big.Int) (*big.Int, error) {
	return _TestClient.Contract.ComputeLeaf(&_TestClient.CallOpts, dataSetLookup, index)
}

// TestOptimization is a free data retrieval call binding the contract method 0x1b8a090a.
//
// Solidity: function testOptimization(dagIndices uint256[], dagElements uint256[], witnessArray uint256[], epoch uint256) constant returns(uint256[4])
func (_TestClient *TestClientCaller) TestOptimization(opts *bind.CallOpts, dagIndices []*big.Int, dagElements []*big.Int, witnessArray []*big.Int, epoch *big.Int) ([4]*big.Int, error) {
	var (
		ret0 = new([4]*big.Int)
	)
	out := ret0
	err := _TestClient.contract.Call(opts, out, "testOptimization", dagIndices, dagElements, witnessArray, epoch)
	return *ret0, err
}

// TestOptimization is a free data retrieval call binding the contract method 0x1b8a090a.
//
// Solidity: function testOptimization(dagIndices uint256[], dagElements uint256[], witnessArray uint256[], epoch uint256) constant returns(uint256[4])
func (_TestClient *TestClientSession) TestOptimization(dagIndices []*big.Int, dagElements []*big.Int, witnessArray []*big.Int, epoch *big.Int) ([4]*big.Int, error) {
	return _TestClient.Contract.TestOptimization(&_TestClient.CallOpts, dagIndices, dagElements, witnessArray, epoch)
}

// TestOptimization is a free data retrieval call binding the contract method 0x1b8a090a.
//
// Solidity: function testOptimization(dagIndices uint256[], dagElements uint256[], witnessArray uint256[], epoch uint256) constant returns(uint256[4])
func (_TestClient *TestClientCallerSession) TestOptimization(dagIndices []*big.Int, dagElements []*big.Int, witnessArray []*big.Int, epoch *big.Int) ([4]*big.Int, error) {
	return _TestClient.Contract.TestOptimization(&_TestClient.CallOpts, dagIndices, dagElements, witnessArray, epoch)
}

// SetOptEpochData is a paid mutator transaction binding the contract method 0x7e043973.
//
// Solidity: function setOptEpochData(epoch uint256, fullSizeIn128Resultion uint256, branchDepth uint256, merkleNodes uint256[], start uint256, numElems uint256) returns()
func (_TestClient *TestClientTransactor) SetOptEpochData(opts *bind.TransactOpts, epoch *big.Int, fullSizeIn128Resultion *big.Int, branchDepth *big.Int, merkleNodes []*big.Int, start *big.Int, numElems *big.Int) (*types.Transaction, error) {
	return _TestClient.contract.Transact(opts, "setOptEpochData", epoch, fullSizeIn128Resultion, branchDepth, merkleNodes, start, numElems)
}

// SetOptEpochData is a paid mutator transaction binding the contract method 0x7e043973.
//
// Solidity: function setOptEpochData(epoch uint256, fullSizeIn128Resultion uint256, branchDepth uint256, merkleNodes uint256[], start uint256, numElems uint256) returns()
func (_TestClient *TestClientSession) SetOptEpochData(epoch *big.Int, fullSizeIn128Resultion *big.Int, branchDepth *big.Int, merkleNodes []*big.Int, start *big.Int, numElems *big.Int) (*types.Transaction, error) {
	return _TestClient.Contract.SetOptEpochData(&_TestClient.TransactOpts, epoch, fullSizeIn128Resultion, branchDepth, merkleNodes, start, numElems)
}

// SetOptEpochData is a paid mutator transaction binding the contract method 0x7e043973.
//
// Solidity: function setOptEpochData(epoch uint256, fullSizeIn128Resultion uint256, branchDepth uint256, merkleNodes uint256[], start uint256, numElems uint256) returns()
func (_TestClient *TestClientTransactorSession) SetOptEpochData(epoch *big.Int, fullSizeIn128Resultion *big.Int, branchDepth *big.Int, merkleNodes []*big.Int, start *big.Int, numElems *big.Int) (*types.Transaction, error) {
	return _TestClient.Contract.SetOptEpochData(&_TestClient.TransactOpts, epoch, fullSizeIn128Resultion, branchDepth, merkleNodes, start, numElems)
}
