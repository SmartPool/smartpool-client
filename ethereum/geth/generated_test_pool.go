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

// TestPoolABI is the input ABI used to generate the binding from.
const TestPoolABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newVersionReleased\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPoolETHBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uncleRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creationBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"debugGetNumPendingSubmissions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"canRegister\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"poolFees\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"paymentAddress\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_uncleRate\",\"type\":\"uint256\"},{\"name\":\"_poolFees\",\"type\":\"uint256\"}],\"name\":\"setUnlceRateAndFees\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getAverageDifficulty\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"debugResetSubmissions\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"seed\",\"type\":\"uint256\"}],\"name\":\"calculateSubmissionIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getClaimSeed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rootHash\",\"type\":\"uint256\"},{\"name\":\"rootMin\",\"type\":\"uint256\"},{\"name\":\"rootMax\",\"type\":\"uint256\"},{\"name\":\"leafHash\",\"type\":\"uint256\"},{\"name\":\"leafCounter\",\"type\":\"uint256\"},{\"name\":\"branchIndex\",\"type\":\"uint256\"},{\"name\":\"countersBranch\",\"type\":\"uint256[]\"},{\"name\":\"hashesBranch\",\"type\":\"uint256[]\"}],\"name\":\"verifyAgtDebugForTesting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethashContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getShareIndexDebugForTestRPC\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"seed\",\"type\":\"uint256\"},{\"name\":\"submissionNumber\",\"type\":\"uint256\"},{\"name\":\"shareIndex\",\"type\":\"uint256\"}],\"name\":\"verifySubmissionIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"miner\",\"type\":\"address\"},{\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"updateWhiteList\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rlpHeader\",\"type\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"submissionIndex\",\"type\":\"uint256\"},{\"name\":\"shareIndex\",\"type\":\"uint256\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"},{\"name\":\"augCountersBranch\",\"type\":\"uint256[]\"},{\"name\":\"augHashesBranch\",\"type\":\"uint256[]\"}],\"name\":\"verifyClaim\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numShares\",\"type\":\"uint256\"},{\"name\":\"difficulty\",\"type\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\"},{\"name\":\"augRoot\",\"type\":\"uint256\"},{\"name\":\"lastClaimBeforeVerification\",\"type\":\"bool\"}],\"name\":\"submitClaim\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getMinerId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"whiteListEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"declareNewerVersion\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"existingIds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"numChars\",\"type\":\"uint256\"}],\"name\":\"to62Encoding\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[3]\"},{\"name\":\"_ethashContract\",\"type\":\"address\"},{\"name\":\"_whiteListEnabeled\",\"type\":\"bool\"}],\"payable\":true,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"numShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"ValidShares\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"UpdateWhiteList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"VerifyExtraData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"VerifyClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountInWei\",\"type\":\"uint256\"}],\"name\":\"IncomingFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"SetUnlceRateAndFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"GetShareIndexDebugForTestRPCSubmissionIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"GetShareIndexDebugForTestRPCShareIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"SubmitClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"DebugResetSubmissions\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"VerifyAgt\",\"type\":\"event\"}]"

// TestPool is an auto generated Go binding around an Ethereum contract.
type TestPool struct {
	TestPoolCaller     // Read-only binding to the contract
	TestPoolTransactor // Write-only binding to the contract
}

// TestPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestPoolSession struct {
	Contract     *TestPool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestPoolCallerSession struct {
	Contract *TestPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TestPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestPoolTransactorSession struct {
	Contract     *TestPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TestPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestPoolRaw struct {
	Contract *TestPool // Generic contract binding to access the raw methods on
}

// TestPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestPoolCallerRaw struct {
	Contract *TestPoolCaller // Generic read-only contract binding to access the raw methods on
}

// TestPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestPoolTransactorRaw struct {
	Contract *TestPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestPool creates a new instance of TestPool, bound to a specific deployed contract.
func NewTestPool(address common.Address, backend bind.ContractBackend) (*TestPool, error) {
	contract, err := bindTestPool(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestPool{TestPoolCaller: TestPoolCaller{contract: contract}, TestPoolTransactor: TestPoolTransactor{contract: contract}}, nil
}

// NewTestPoolCaller creates a new read-only instance of TestPool, bound to a specific deployed contract.
func NewTestPoolCaller(address common.Address, caller bind.ContractCaller) (*TestPoolCaller, error) {
	contract, err := bindTestPool(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TestPoolCaller{contract: contract}, nil
}

// NewTestPoolTransactor creates a new write-only instance of TestPool, bound to a specific deployed contract.
func NewTestPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*TestPoolTransactor, error) {
	contract, err := bindTestPool(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TestPoolTransactor{contract: contract}, nil
}

// bindTestPool binds a generic wrapper to an already deployed contract.
func bindTestPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPool *TestPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestPool.Contract.TestPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPool *TestPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPool.Contract.TestPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPool *TestPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPool.Contract.TestPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPool *TestPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPool *TestPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPool *TestPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPool.Contract.contract.Transact(opts, method, params...)
}

// CalculateSubmissionIndex is a free data retrieval call binding the contract method 0x5ed1057d.
//
// Solidity: function calculateSubmissionIndex(sender address, seed uint256) constant returns(uint256[2])
func (_TestPool *TestPoolCaller) CalculateSubmissionIndex(opts *bind.CallOpts, sender common.Address, seed *big.Int) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "calculateSubmissionIndex", sender, seed)
	return *ret0, err
}

// CalculateSubmissionIndex is a free data retrieval call binding the contract method 0x5ed1057d.
//
// Solidity: function calculateSubmissionIndex(sender address, seed uint256) constant returns(uint256[2])
func (_TestPool *TestPoolSession) CalculateSubmissionIndex(sender common.Address, seed *big.Int) ([2]*big.Int, error) {
	return _TestPool.Contract.CalculateSubmissionIndex(&_TestPool.CallOpts, sender, seed)
}

// CalculateSubmissionIndex is a free data retrieval call binding the contract method 0x5ed1057d.
//
// Solidity: function calculateSubmissionIndex(sender address, seed uint256) constant returns(uint256[2])
func (_TestPool *TestPoolCallerSession) CalculateSubmissionIndex(sender common.Address, seed *big.Int) ([2]*big.Int, error) {
	return _TestPool.Contract.CalculateSubmissionIndex(&_TestPool.CallOpts, sender, seed)
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(sender address) constant returns(bool)
func (_TestPool *TestPoolCaller) CanRegister(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "canRegister", sender)
	return *ret0, err
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(sender address) constant returns(bool)
func (_TestPool *TestPoolSession) CanRegister(sender common.Address) (bool, error) {
	return _TestPool.Contract.CanRegister(&_TestPool.CallOpts, sender)
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(sender address) constant returns(bool)
func (_TestPool *TestPoolCallerSession) CanRegister(sender common.Address) (bool, error) {
	return _TestPool.Contract.CanRegister(&_TestPool.CallOpts, sender)
}

// CreationBlockNumber is a free data retrieval call binding the contract method 0x1bf30929.
//
// Solidity: function creationBlockNumber() constant returns(uint256)
func (_TestPool *TestPoolCaller) CreationBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "creationBlockNumber")
	return *ret0, err
}

// CreationBlockNumber is a free data retrieval call binding the contract method 0x1bf30929.
//
// Solidity: function creationBlockNumber() constant returns(uint256)
func (_TestPool *TestPoolSession) CreationBlockNumber() (*big.Int, error) {
	return _TestPool.Contract.CreationBlockNumber(&_TestPool.CallOpts)
}

// CreationBlockNumber is a free data retrieval call binding the contract method 0x1bf30929.
//
// Solidity: function creationBlockNumber() constant returns(uint256)
func (_TestPool *TestPoolCallerSession) CreationBlockNumber() (*big.Int, error) {
	return _TestPool.Contract.CreationBlockNumber(&_TestPool.CallOpts)
}

// DebugGetNumPendingSubmissions is a free data retrieval call binding the contract method 0x21aa4548.
//
// Solidity: function debugGetNumPendingSubmissions(sender address) constant returns(uint256)
func (_TestPool *TestPoolCaller) DebugGetNumPendingSubmissions(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "debugGetNumPendingSubmissions", sender)
	return *ret0, err
}

// DebugGetNumPendingSubmissions is a free data retrieval call binding the contract method 0x21aa4548.
//
// Solidity: function debugGetNumPendingSubmissions(sender address) constant returns(uint256)
func (_TestPool *TestPoolSession) DebugGetNumPendingSubmissions(sender common.Address) (*big.Int, error) {
	return _TestPool.Contract.DebugGetNumPendingSubmissions(&_TestPool.CallOpts, sender)
}

// DebugGetNumPendingSubmissions is a free data retrieval call binding the contract method 0x21aa4548.
//
// Solidity: function debugGetNumPendingSubmissions(sender address) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) DebugGetNumPendingSubmissions(sender common.Address) (*big.Int, error) {
	return _TestPool.Contract.DebugGetNumPendingSubmissions(&_TestPool.CallOpts, sender)
}

// EthashContract is a free data retrieval call binding the contract method 0x7f949ac0.
//
// Solidity: function ethashContract() constant returns(address)
func (_TestPool *TestPoolCaller) EthashContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "ethashContract")
	return *ret0, err
}

// EthashContract is a free data retrieval call binding the contract method 0x7f949ac0.
//
// Solidity: function ethashContract() constant returns(address)
func (_TestPool *TestPoolSession) EthashContract() (common.Address, error) {
	return _TestPool.Contract.EthashContract(&_TestPool.CallOpts)
}

// EthashContract is a free data retrieval call binding the contract method 0x7f949ac0.
//
// Solidity: function ethashContract() constant returns(address)
func (_TestPool *TestPoolCallerSession) EthashContract() (common.Address, error) {
	return _TestPool.Contract.EthashContract(&_TestPool.CallOpts)
}

// ExistingIds is a free data retrieval call binding the contract method 0xee385304.
//
// Solidity: function existingIds( bytes32) constant returns(bool)
func (_TestPool *TestPoolCaller) ExistingIds(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "existingIds", arg0)
	return *ret0, err
}

// ExistingIds is a free data retrieval call binding the contract method 0xee385304.
//
// Solidity: function existingIds( bytes32) constant returns(bool)
func (_TestPool *TestPoolSession) ExistingIds(arg0 [32]byte) (bool, error) {
	return _TestPool.Contract.ExistingIds(&_TestPool.CallOpts, arg0)
}

// ExistingIds is a free data retrieval call binding the contract method 0xee385304.
//
// Solidity: function existingIds( bytes32) constant returns(bool)
func (_TestPool *TestPoolCallerSession) ExistingIds(arg0 [32]byte) (bool, error) {
	return _TestPool.Contract.ExistingIds(&_TestPool.CallOpts, arg0)
}

// GetAverageDifficulty is a free data retrieval call binding the contract method 0x48d17a67.
//
// Solidity: function getAverageDifficulty(sender address) constant returns(uint256)
func (_TestPool *TestPoolCaller) GetAverageDifficulty(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "getAverageDifficulty", sender)
	return *ret0, err
}

// GetAverageDifficulty is a free data retrieval call binding the contract method 0x48d17a67.
//
// Solidity: function getAverageDifficulty(sender address) constant returns(uint256)
func (_TestPool *TestPoolSession) GetAverageDifficulty(sender common.Address) (*big.Int, error) {
	return _TestPool.Contract.GetAverageDifficulty(&_TestPool.CallOpts, sender)
}

// GetAverageDifficulty is a free data retrieval call binding the contract method 0x48d17a67.
//
// Solidity: function getAverageDifficulty(sender address) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) GetAverageDifficulty(sender common.Address) (*big.Int, error) {
	return _TestPool.Contract.GetAverageDifficulty(&_TestPool.CallOpts, sender)
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x7087ed2c.
//
// Solidity: function getClaimSeed(sender address) constant returns(uint256)
func (_TestPool *TestPoolCaller) GetClaimSeed(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "getClaimSeed", sender)
	return *ret0, err
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x7087ed2c.
//
// Solidity: function getClaimSeed(sender address) constant returns(uint256)
func (_TestPool *TestPoolSession) GetClaimSeed(sender common.Address) (*big.Int, error) {
	return _TestPool.Contract.GetClaimSeed(&_TestPool.CallOpts, sender)
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x7087ed2c.
//
// Solidity: function getClaimSeed(sender address) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) GetClaimSeed(sender common.Address) (*big.Int, error) {
	return _TestPool.Contract.GetClaimSeed(&_TestPool.CallOpts, sender)
}

// GetMinerId is a free data retrieval call binding the contract method 0xe2dea715.
//
// Solidity: function getMinerId(sender address) constant returns(bytes32)
func (_TestPool *TestPoolCaller) GetMinerId(opts *bind.CallOpts, sender common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "getMinerId", sender)
	return *ret0, err
}

// GetMinerId is a free data retrieval call binding the contract method 0xe2dea715.
//
// Solidity: function getMinerId(sender address) constant returns(bytes32)
func (_TestPool *TestPoolSession) GetMinerId(sender common.Address) ([32]byte, error) {
	return _TestPool.Contract.GetMinerId(&_TestPool.CallOpts, sender)
}

// GetMinerId is a free data retrieval call binding the contract method 0xe2dea715.
//
// Solidity: function getMinerId(sender address) constant returns(bytes32)
func (_TestPool *TestPoolCallerSession) GetMinerId(sender common.Address) ([32]byte, error) {
	return _TestPool.Contract.GetMinerId(&_TestPool.CallOpts, sender)
}

// GetPoolETHBalance is a free data retrieval call binding the contract method 0x0abb8409.
//
// Solidity: function getPoolETHBalance() constant returns(uint256)
func (_TestPool *TestPoolCaller) GetPoolETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "getPoolETHBalance")
	return *ret0, err
}

// GetPoolETHBalance is a free data retrieval call binding the contract method 0x0abb8409.
//
// Solidity: function getPoolETHBalance() constant returns(uint256)
func (_TestPool *TestPoolSession) GetPoolETHBalance() (*big.Int, error) {
	return _TestPool.Contract.GetPoolETHBalance(&_TestPool.CallOpts)
}

// GetPoolETHBalance is a free data retrieval call binding the contract method 0x0abb8409.
//
// Solidity: function getPoolETHBalance() constant returns(uint256)
func (_TestPool *TestPoolCallerSession) GetPoolETHBalance() (*big.Int, error) {
	return _TestPool.Contract.GetPoolETHBalance(&_TestPool.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(sender address) constant returns(bool)
func (_TestPool *TestPoolCaller) IsRegistered(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "isRegistered", sender)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(sender address) constant returns(bool)
func (_TestPool *TestPoolSession) IsRegistered(sender common.Address) (bool, error) {
	return _TestPool.Contract.IsRegistered(&_TestPool.CallOpts, sender)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(sender address) constant returns(bool)
func (_TestPool *TestPoolCallerSession) IsRegistered(sender common.Address) (bool, error) {
	return _TestPool.Contract.IsRegistered(&_TestPool.CallOpts, sender)
}

// NewVersionReleased is a free data retrieval call binding the contract method 0x0289e966.
//
// Solidity: function newVersionReleased() constant returns(bool)
func (_TestPool *TestPoolCaller) NewVersionReleased(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "newVersionReleased")
	return *ret0, err
}

// NewVersionReleased is a free data retrieval call binding the contract method 0x0289e966.
//
// Solidity: function newVersionReleased() constant returns(bool)
func (_TestPool *TestPoolSession) NewVersionReleased() (bool, error) {
	return _TestPool.Contract.NewVersionReleased(&_TestPool.CallOpts)
}

// NewVersionReleased is a free data retrieval call binding the contract method 0x0289e966.
//
// Solidity: function newVersionReleased() constant returns(bool)
func (_TestPool *TestPoolCallerSession) NewVersionReleased() (bool, error) {
	return _TestPool.Contract.NewVersionReleased(&_TestPool.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(bool)
func (_TestPool *TestPoolCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(bool)
func (_TestPool *TestPoolSession) Owners(arg0 common.Address) (bool, error) {
	return _TestPool.Contract.Owners(&_TestPool.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(bool)
func (_TestPool *TestPoolCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _TestPool.Contract.Owners(&_TestPool.CallOpts, arg0)
}

// PoolFees is a free data retrieval call binding the contract method 0x33580959.
//
// Solidity: function poolFees() constant returns(uint256)
func (_TestPool *TestPoolCaller) PoolFees(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "poolFees")
	return *ret0, err
}

// PoolFees is a free data retrieval call binding the contract method 0x33580959.
//
// Solidity: function poolFees() constant returns(uint256)
func (_TestPool *TestPoolSession) PoolFees() (*big.Int, error) {
	return _TestPool.Contract.PoolFees(&_TestPool.CallOpts)
}

// PoolFees is a free data retrieval call binding the contract method 0x33580959.
//
// Solidity: function poolFees() constant returns(uint256)
func (_TestPool *TestPoolCallerSession) PoolFees() (*big.Int, error) {
	return _TestPool.Contract.PoolFees(&_TestPool.CallOpts)
}

// To62Encoding is a free data retrieval call binding the contract method 0xff5d2c39.
//
// Solidity: function to62Encoding(id uint256, numChars uint256) constant returns(bytes32)
func (_TestPool *TestPoolCaller) To62Encoding(opts *bind.CallOpts, id *big.Int, numChars *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "to62Encoding", id, numChars)
	return *ret0, err
}

// To62Encoding is a free data retrieval call binding the contract method 0xff5d2c39.
//
// Solidity: function to62Encoding(id uint256, numChars uint256) constant returns(bytes32)
func (_TestPool *TestPoolSession) To62Encoding(id *big.Int, numChars *big.Int) ([32]byte, error) {
	return _TestPool.Contract.To62Encoding(&_TestPool.CallOpts, id, numChars)
}

// To62Encoding is a free data retrieval call binding the contract method 0xff5d2c39.
//
// Solidity: function to62Encoding(id uint256, numChars uint256) constant returns(bytes32)
func (_TestPool *TestPoolCallerSession) To62Encoding(id *big.Int, numChars *big.Int) ([32]byte, error) {
	return _TestPool.Contract.To62Encoding(&_TestPool.CallOpts, id, numChars)
}

// UncleRate is a free data retrieval call binding the contract method 0x1bc41284.
//
// Solidity: function uncleRate() constant returns(uint256)
func (_TestPool *TestPoolCaller) UncleRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "uncleRate")
	return *ret0, err
}

// UncleRate is a free data retrieval call binding the contract method 0x1bc41284.
//
// Solidity: function uncleRate() constant returns(uint256)
func (_TestPool *TestPoolSession) UncleRate() (*big.Int, error) {
	return _TestPool.Contract.UncleRate(&_TestPool.CallOpts)
}

// UncleRate is a free data retrieval call binding the contract method 0x1bc41284.
//
// Solidity: function uncleRate() constant returns(uint256)
func (_TestPool *TestPoolCallerSession) UncleRate() (*big.Int, error) {
	return _TestPool.Contract.UncleRate(&_TestPool.CallOpts)
}

// VerifySubmissionIndex is a free data retrieval call binding the contract method 0x9e133be4.
//
// Solidity: function verifySubmissionIndex(sender address, seed uint256, submissionNumber uint256, shareIndex uint256) constant returns(bool)
func (_TestPool *TestPoolCaller) VerifySubmissionIndex(opts *bind.CallOpts, sender common.Address, seed *big.Int, submissionNumber *big.Int, shareIndex *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "verifySubmissionIndex", sender, seed, submissionNumber, shareIndex)
	return *ret0, err
}

// VerifySubmissionIndex is a free data retrieval call binding the contract method 0x9e133be4.
//
// Solidity: function verifySubmissionIndex(sender address, seed uint256, submissionNumber uint256, shareIndex uint256) constant returns(bool)
func (_TestPool *TestPoolSession) VerifySubmissionIndex(sender common.Address, seed *big.Int, submissionNumber *big.Int, shareIndex *big.Int) (bool, error) {
	return _TestPool.Contract.VerifySubmissionIndex(&_TestPool.CallOpts, sender, seed, submissionNumber, shareIndex)
}

// VerifySubmissionIndex is a free data retrieval call binding the contract method 0x9e133be4.
//
// Solidity: function verifySubmissionIndex(sender address, seed uint256, submissionNumber uint256, shareIndex uint256) constant returns(bool)
func (_TestPool *TestPoolCallerSession) VerifySubmissionIndex(sender common.Address, seed *big.Int, submissionNumber *big.Int, shareIndex *big.Int) (bool, error) {
	return _TestPool.Contract.VerifySubmissionIndex(&_TestPool.CallOpts, sender, seed, submissionNumber, shareIndex)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TestPool *TestPoolCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TestPool *TestPoolSession) Version() (string, error) {
	return _TestPool.Contract.Version(&_TestPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TestPool *TestPoolCallerSession) Version() (string, error) {
	return _TestPool.Contract.Version(&_TestPool.CallOpts)
}

// WhiteListEnabled is a free data retrieval call binding the contract method 0xe2e616bb.
//
// Solidity: function whiteListEnabled() constant returns(bool)
func (_TestPool *TestPoolCaller) WhiteListEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "whiteListEnabled")
	return *ret0, err
}

// WhiteListEnabled is a free data retrieval call binding the contract method 0xe2e616bb.
//
// Solidity: function whiteListEnabled() constant returns(bool)
func (_TestPool *TestPoolSession) WhiteListEnabled() (bool, error) {
	return _TestPool.Contract.WhiteListEnabled(&_TestPool.CallOpts)
}

// WhiteListEnabled is a free data retrieval call binding the contract method 0xe2e616bb.
//
// Solidity: function whiteListEnabled() constant returns(bool)
func (_TestPool *TestPoolCallerSession) WhiteListEnabled() (bool, error) {
	return _TestPool.Contract.WhiteListEnabled(&_TestPool.CallOpts)
}

// DebugResetSubmissions is a paid mutator transaction binding the contract method 0x5eadd607.
//
// Solidity: function debugResetSubmissions() returns()
func (_TestPool *TestPoolTransactor) DebugResetSubmissions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "debugResetSubmissions")
}

// DebugResetSubmissions is a paid mutator transaction binding the contract method 0x5eadd607.
//
// Solidity: function debugResetSubmissions() returns()
func (_TestPool *TestPoolSession) DebugResetSubmissions() (*types.Transaction, error) {
	return _TestPool.Contract.DebugResetSubmissions(&_TestPool.TransactOpts)
}

// DebugResetSubmissions is a paid mutator transaction binding the contract method 0x5eadd607.
//
// Solidity: function debugResetSubmissions() returns()
func (_TestPool *TestPoolTransactorSession) DebugResetSubmissions() (*types.Transaction, error) {
	return _TestPool.Contract.DebugResetSubmissions(&_TestPool.TransactOpts)
}

// DeclareNewerVersion is a paid mutator transaction binding the contract method 0xe3d86998.
//
// Solidity: function declareNewerVersion() returns()
func (_TestPool *TestPoolTransactor) DeclareNewerVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "declareNewerVersion")
}

// DeclareNewerVersion is a paid mutator transaction binding the contract method 0xe3d86998.
//
// Solidity: function declareNewerVersion() returns()
func (_TestPool *TestPoolSession) DeclareNewerVersion() (*types.Transaction, error) {
	return _TestPool.Contract.DeclareNewerVersion(&_TestPool.TransactOpts)
}

// DeclareNewerVersion is a paid mutator transaction binding the contract method 0xe3d86998.
//
// Solidity: function declareNewerVersion() returns()
func (_TestPool *TestPoolTransactorSession) DeclareNewerVersion() (*types.Transaction, error) {
	return _TestPool.Contract.DeclareNewerVersion(&_TestPool.TransactOpts)
}

// GetShareIndexDebugForTestRPC is a paid mutator transaction binding the contract method 0x84c65296.
//
// Solidity: function getShareIndexDebugForTestRPC(sender address) returns()
func (_TestPool *TestPoolTransactor) GetShareIndexDebugForTestRPC(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "getShareIndexDebugForTestRPC", sender)
}

// GetShareIndexDebugForTestRPC is a paid mutator transaction binding the contract method 0x84c65296.
//
// Solidity: function getShareIndexDebugForTestRPC(sender address) returns()
func (_TestPool *TestPoolSession) GetShareIndexDebugForTestRPC(sender common.Address) (*types.Transaction, error) {
	return _TestPool.Contract.GetShareIndexDebugForTestRPC(&_TestPool.TransactOpts, sender)
}

// GetShareIndexDebugForTestRPC is a paid mutator transaction binding the contract method 0x84c65296.
//
// Solidity: function getShareIndexDebugForTestRPC(sender address) returns()
func (_TestPool *TestPoolTransactorSession) GetShareIndexDebugForTestRPC(sender common.Address) (*types.Transaction, error) {
	return _TestPool.Contract.GetShareIndexDebugForTestRPC(&_TestPool.TransactOpts, sender)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(paymentAddress address) returns()
func (_TestPool *TestPoolTransactor) Register(opts *bind.TransactOpts, paymentAddress common.Address) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "register", paymentAddress)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(paymentAddress address) returns()
func (_TestPool *TestPoolSession) Register(paymentAddress common.Address) (*types.Transaction, error) {
	return _TestPool.Contract.Register(&_TestPool.TransactOpts, paymentAddress)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(paymentAddress address) returns()
func (_TestPool *TestPoolTransactorSession) Register(paymentAddress common.Address) (*types.Transaction, error) {
	return _TestPool.Contract.Register(&_TestPool.TransactOpts, paymentAddress)
}

// SetUnlceRateAndFees is a paid mutator transaction binding the contract method 0x46b3f696.
//
// Solidity: function setUnlceRateAndFees(_uncleRate uint256, _poolFees uint256) returns()
func (_TestPool *TestPoolTransactor) SetUnlceRateAndFees(opts *bind.TransactOpts, _uncleRate *big.Int, _poolFees *big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "setUnlceRateAndFees", _uncleRate, _poolFees)
}

// SetUnlceRateAndFees is a paid mutator transaction binding the contract method 0x46b3f696.
//
// Solidity: function setUnlceRateAndFees(_uncleRate uint256, _poolFees uint256) returns()
func (_TestPool *TestPoolSession) SetUnlceRateAndFees(_uncleRate *big.Int, _poolFees *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SetUnlceRateAndFees(&_TestPool.TransactOpts, _uncleRate, _poolFees)
}

// SetUnlceRateAndFees is a paid mutator transaction binding the contract method 0x46b3f696.
//
// Solidity: function setUnlceRateAndFees(_uncleRate uint256, _poolFees uint256) returns()
func (_TestPool *TestPoolTransactorSession) SetUnlceRateAndFees(_uncleRate *big.Int, _poolFees *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SetUnlceRateAndFees(&_TestPool.TransactOpts, _uncleRate, _poolFees)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xcedb217a.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augRoot uint256, lastClaimBeforeVerification bool) returns()
func (_TestPool *TestPoolTransactor) SubmitClaim(opts *bind.TransactOpts, numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augRoot *big.Int, lastClaimBeforeVerification bool) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "submitClaim", numShares, difficulty, min, max, augRoot, lastClaimBeforeVerification)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xcedb217a.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augRoot uint256, lastClaimBeforeVerification bool) returns()
func (_TestPool *TestPoolSession) SubmitClaim(numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augRoot *big.Int, lastClaimBeforeVerification bool) (*types.Transaction, error) {
	return _TestPool.Contract.SubmitClaim(&_TestPool.TransactOpts, numShares, difficulty, min, max, augRoot, lastClaimBeforeVerification)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xcedb217a.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augRoot uint256, lastClaimBeforeVerification bool) returns()
func (_TestPool *TestPoolTransactorSession) SubmitClaim(numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augRoot *big.Int, lastClaimBeforeVerification bool) (*types.Transaction, error) {
	return _TestPool.Contract.SubmitClaim(&_TestPool.TransactOpts, numShares, difficulty, min, max, augRoot, lastClaimBeforeVerification)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0xac1d0609.
//
// Solidity: function updateWhiteList(miner address, add bool) returns()
func (_TestPool *TestPoolTransactor) UpdateWhiteList(opts *bind.TransactOpts, miner common.Address, add bool) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "updateWhiteList", miner, add)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0xac1d0609.
//
// Solidity: function updateWhiteList(miner address, add bool) returns()
func (_TestPool *TestPoolSession) UpdateWhiteList(miner common.Address, add bool) (*types.Transaction, error) {
	return _TestPool.Contract.UpdateWhiteList(&_TestPool.TransactOpts, miner, add)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0xac1d0609.
//
// Solidity: function updateWhiteList(miner address, add bool) returns()
func (_TestPool *TestPoolTransactorSession) UpdateWhiteList(miner common.Address, add bool) (*types.Transaction, error) {
	return _TestPool.Contract.UpdateWhiteList(&_TestPool.TransactOpts, miner, add)
}

// VerifyAgtDebugForTesting is a paid mutator transaction binding the contract method 0x7ed996a3.
//
// Solidity: function verifyAgtDebugForTesting(rootHash uint256, rootMin uint256, rootMax uint256, leafHash uint256, leafCounter uint256, branchIndex uint256, countersBranch uint256[], hashesBranch uint256[]) returns(bool)
func (_TestPool *TestPoolTransactor) VerifyAgtDebugForTesting(opts *bind.TransactOpts, rootHash *big.Int, rootMin *big.Int, rootMax *big.Int, leafHash *big.Int, leafCounter *big.Int, branchIndex *big.Int, countersBranch []*big.Int, hashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "verifyAgtDebugForTesting", rootHash, rootMin, rootMax, leafHash, leafCounter, branchIndex, countersBranch, hashesBranch)
}

// VerifyAgtDebugForTesting is a paid mutator transaction binding the contract method 0x7ed996a3.
//
// Solidity: function verifyAgtDebugForTesting(rootHash uint256, rootMin uint256, rootMax uint256, leafHash uint256, leafCounter uint256, branchIndex uint256, countersBranch uint256[], hashesBranch uint256[]) returns(bool)
func (_TestPool *TestPoolSession) VerifyAgtDebugForTesting(rootHash *big.Int, rootMin *big.Int, rootMax *big.Int, leafHash *big.Int, leafCounter *big.Int, branchIndex *big.Int, countersBranch []*big.Int, hashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.VerifyAgtDebugForTesting(&_TestPool.TransactOpts, rootHash, rootMin, rootMax, leafHash, leafCounter, branchIndex, countersBranch, hashesBranch)
}

// VerifyAgtDebugForTesting is a paid mutator transaction binding the contract method 0x7ed996a3.
//
// Solidity: function verifyAgtDebugForTesting(rootHash uint256, rootMin uint256, rootMax uint256, leafHash uint256, leafCounter uint256, branchIndex uint256, countersBranch uint256[], hashesBranch uint256[]) returns(bool)
func (_TestPool *TestPoolTransactorSession) VerifyAgtDebugForTesting(rootHash *big.Int, rootMin *big.Int, rootMax *big.Int, leafHash *big.Int, leafCounter *big.Int, branchIndex *big.Int, countersBranch []*big.Int, hashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.VerifyAgtDebugForTesting(&_TestPool.TransactOpts, rootHash, rootMin, rootMax, leafHash, leafCounter, branchIndex, countersBranch, hashesBranch)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0xc2209acb.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, submissionIndex uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns()
func (_TestPool *TestPoolTransactor) VerifyClaim(opts *bind.TransactOpts, rlpHeader []byte, nonce *big.Int, submissionIndex *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "verifyClaim", rlpHeader, nonce, submissionIndex, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0xc2209acb.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, submissionIndex uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns()
func (_TestPool *TestPoolSession) VerifyClaim(rlpHeader []byte, nonce *big.Int, submissionIndex *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.VerifyClaim(&_TestPool.TransactOpts, rlpHeader, nonce, submissionIndex, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0xc2209acb.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, submissionIndex uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns()
func (_TestPool *TestPoolTransactorSession) VerifyClaim(rlpHeader []byte, nonce *big.Int, submissionIndex *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.VerifyClaim(&_TestPool.TransactOpts, rlpHeader, nonce, submissionIndex, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}
