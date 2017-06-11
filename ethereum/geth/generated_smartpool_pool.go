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

// SmartPoolABI is the input ABI used to generate the binding from.
const SmartPoolABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newVersionReleased\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPoolETHBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uncleRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"debugGetNumPendingSubmissions\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"canRegister\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"poolFees\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"paymentAddress\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_uncleRate\",\"type\":\"uint256\"},{\"name\":\"_poolFees\",\"type\":\"uint256\"}],\"name\":\"setUnlceRateAndFees\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"debugResetSubmissions\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"seed\",\"type\":\"uint256\"}],\"name\":\"calculateSubmissionIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getClaimSeed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rootHash\",\"type\":\"uint256\"},{\"name\":\"rootMin\",\"type\":\"uint256\"},{\"name\":\"rootMax\",\"type\":\"uint256\"},{\"name\":\"leafHash\",\"type\":\"uint256\"},{\"name\":\"leafCounter\",\"type\":\"uint256\"},{\"name\":\"branchIndex\",\"type\":\"uint256\"},{\"name\":\"countersBranch\",\"type\":\"uint256[]\"},{\"name\":\"hashesBranch\",\"type\":\"uint256[]\"}],\"name\":\"verifyAgtDebugForTesting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethashContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getShareIndexDebugForTestRPC\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"seed\",\"type\":\"uint256\"},{\"name\":\"submissionNumber\",\"type\":\"uint256\"},{\"name\":\"shareIndex\",\"type\":\"uint256\"}],\"name\":\"verifySubmissionIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"miner\",\"type\":\"address\"},{\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"updateWhiteList\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rlpHeader\",\"type\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"submissionIndex\",\"type\":\"uint256\"},{\"name\":\"shareIndex\",\"type\":\"uint256\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"},{\"name\":\"augCountersBranch\",\"type\":\"uint256[]\"},{\"name\":\"augHashesBranch\",\"type\":\"uint256[]\"}],\"name\":\"verifyClaim\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"miner\",\"type\":\"address\"}],\"name\":\"storeClaimSeed\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numShares\",\"type\":\"uint256\"},{\"name\":\"difficulty\",\"type\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\"},{\"name\":\"augRoot\",\"type\":\"uint256\"},{\"name\":\"lastClaimBeforeVerification\",\"type\":\"bool\"}],\"name\":\"submitClaim\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getMinerId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"whiteListEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"declareNewerVersion\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"existingIds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawalAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"numChars\",\"type\":\"uint256\"}],\"name\":\"to62Encoding\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[3]\"},{\"name\":\"_ethashContract\",\"type\":\"address\"},{\"name\":\"_withdrawalAddress\",\"type\":\"address\"},{\"name\":\"_whiteListEnabeled\",\"type\":\"bool\"}],\"payable\":true,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"UpdateWhiteList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"VerifyExtraData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"VerifyClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountInWei\",\"type\":\"uint256\"}],\"name\":\"IncomingFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"SetUnlceRateAndFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"paymentAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"valueInWei\",\"type\":\"uint256\"}],\"name\":\"DoPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"GetShareIndexDebugForTestRPCSubmissionIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"GetShareIndexDebugForTestRPCShareIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"SubmitClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"StoreClaimSeed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"DebugResetSubmissions\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"VerifyAgt\",\"type\":\"event\"}]"

// SmartPool is an auto generated Go binding around an Ethereum contract.
type SmartPool struct {
	SmartPoolCaller     // Read-only binding to the contract
	SmartPoolTransactor // Write-only binding to the contract
}

// SmartPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartPoolSession struct {
	Contract     *SmartPool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartPoolCallerSession struct {
	Contract *SmartPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SmartPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartPoolTransactorSession struct {
	Contract     *SmartPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SmartPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartPoolRaw struct {
	Contract *SmartPool // Generic contract binding to access the raw methods on
}

// SmartPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartPoolCallerRaw struct {
	Contract *SmartPoolCaller // Generic read-only contract binding to access the raw methods on
}

// SmartPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartPoolTransactorRaw struct {
	Contract *SmartPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartPool creates a new instance of SmartPool, bound to a specific deployed contract.
func NewSmartPool(address common.Address, backend bind.ContractBackend) (*SmartPool, error) {
	contract, err := bindSmartPool(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartPool{SmartPoolCaller: SmartPoolCaller{contract: contract}, SmartPoolTransactor: SmartPoolTransactor{contract: contract}}, nil
}

// NewSmartPoolCaller creates a new read-only instance of SmartPool, bound to a specific deployed contract.
func NewSmartPoolCaller(address common.Address, caller bind.ContractCaller) (*SmartPoolCaller, error) {
	contract, err := bindSmartPool(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SmartPoolCaller{contract: contract}, nil
}

// NewSmartPoolTransactor creates a new write-only instance of SmartPool, bound to a specific deployed contract.
func NewSmartPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartPoolTransactor, error) {
	contract, err := bindSmartPool(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SmartPoolTransactor{contract: contract}, nil
}

// bindSmartPool binds a generic wrapper to an already deployed contract.
func bindSmartPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartPool *SmartPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartPool.Contract.SmartPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartPool *SmartPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartPool.Contract.SmartPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartPool *SmartPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartPool.Contract.SmartPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartPool *SmartPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartPool *SmartPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartPool *SmartPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartPool.Contract.contract.Transact(opts, method, params...)
}

// CalculateSubmissionIndex is a free data retrieval call binding the contract method 0x5ed1057d.
//
// Solidity: function calculateSubmissionIndex(sender address, seed uint256) constant returns(uint256[2])
func (_SmartPool *SmartPoolCaller) CalculateSubmissionIndex(opts *bind.CallOpts, sender common.Address, seed *big.Int) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "calculateSubmissionIndex", sender, seed)
	return *ret0, err
}

// CalculateSubmissionIndex is a free data retrieval call binding the contract method 0x5ed1057d.
//
// Solidity: function calculateSubmissionIndex(sender address, seed uint256) constant returns(uint256[2])
func (_SmartPool *SmartPoolSession) CalculateSubmissionIndex(sender common.Address, seed *big.Int) ([2]*big.Int, error) {
	return _SmartPool.Contract.CalculateSubmissionIndex(&_SmartPool.CallOpts, sender, seed)
}

// CalculateSubmissionIndex is a free data retrieval call binding the contract method 0x5ed1057d.
//
// Solidity: function calculateSubmissionIndex(sender address, seed uint256) constant returns(uint256[2])
func (_SmartPool *SmartPoolCallerSession) CalculateSubmissionIndex(sender common.Address, seed *big.Int) ([2]*big.Int, error) {
	return _SmartPool.Contract.CalculateSubmissionIndex(&_SmartPool.CallOpts, sender, seed)
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(sender address) constant returns(bool)
func (_SmartPool *SmartPoolCaller) CanRegister(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "canRegister", sender)
	return *ret0, err
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(sender address) constant returns(bool)
func (_SmartPool *SmartPoolSession) CanRegister(sender common.Address) (bool, error) {
	return _SmartPool.Contract.CanRegister(&_SmartPool.CallOpts, sender)
}

// CanRegister is a free data retrieval call binding the contract method 0x320d46d4.
//
// Solidity: function canRegister(sender address) constant returns(bool)
func (_SmartPool *SmartPoolCallerSession) CanRegister(sender common.Address) (bool, error) {
	return _SmartPool.Contract.CanRegister(&_SmartPool.CallOpts, sender)
}

// DebugGetNumPendingSubmissions is a free data retrieval call binding the contract method 0x21aa4548.
//
// Solidity: function debugGetNumPendingSubmissions(sender address) constant returns(uint256)
func (_SmartPool *SmartPoolCaller) DebugGetNumPendingSubmissions(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "debugGetNumPendingSubmissions", sender)
	return *ret0, err
}

// DebugGetNumPendingSubmissions is a free data retrieval call binding the contract method 0x21aa4548.
//
// Solidity: function debugGetNumPendingSubmissions(sender address) constant returns(uint256)
func (_SmartPool *SmartPoolSession) DebugGetNumPendingSubmissions(sender common.Address) (*big.Int, error) {
	return _SmartPool.Contract.DebugGetNumPendingSubmissions(&_SmartPool.CallOpts, sender)
}

// DebugGetNumPendingSubmissions is a free data retrieval call binding the contract method 0x21aa4548.
//
// Solidity: function debugGetNumPendingSubmissions(sender address) constant returns(uint256)
func (_SmartPool *SmartPoolCallerSession) DebugGetNumPendingSubmissions(sender common.Address) (*big.Int, error) {
	return _SmartPool.Contract.DebugGetNumPendingSubmissions(&_SmartPool.CallOpts, sender)
}

// EthashContract is a free data retrieval call binding the contract method 0x7f949ac0.
//
// Solidity: function ethashContract() constant returns(address)
func (_SmartPool *SmartPoolCaller) EthashContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "ethashContract")
	return *ret0, err
}

// EthashContract is a free data retrieval call binding the contract method 0x7f949ac0.
//
// Solidity: function ethashContract() constant returns(address)
func (_SmartPool *SmartPoolSession) EthashContract() (common.Address, error) {
	return _SmartPool.Contract.EthashContract(&_SmartPool.CallOpts)
}

// EthashContract is a free data retrieval call binding the contract method 0x7f949ac0.
//
// Solidity: function ethashContract() constant returns(address)
func (_SmartPool *SmartPoolCallerSession) EthashContract() (common.Address, error) {
	return _SmartPool.Contract.EthashContract(&_SmartPool.CallOpts)
}

// ExistingIds is a free data retrieval call binding the contract method 0xee385304.
//
// Solidity: function existingIds( bytes32) constant returns(bool)
func (_SmartPool *SmartPoolCaller) ExistingIds(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "existingIds", arg0)
	return *ret0, err
}

// ExistingIds is a free data retrieval call binding the contract method 0xee385304.
//
// Solidity: function existingIds( bytes32) constant returns(bool)
func (_SmartPool *SmartPoolSession) ExistingIds(arg0 [32]byte) (bool, error) {
	return _SmartPool.Contract.ExistingIds(&_SmartPool.CallOpts, arg0)
}

// ExistingIds is a free data retrieval call binding the contract method 0xee385304.
//
// Solidity: function existingIds( bytes32) constant returns(bool)
func (_SmartPool *SmartPoolCallerSession) ExistingIds(arg0 [32]byte) (bool, error) {
	return _SmartPool.Contract.ExistingIds(&_SmartPool.CallOpts, arg0)
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x7087ed2c.
//
// Solidity: function getClaimSeed(sender address) constant returns(uint256)
func (_SmartPool *SmartPoolCaller) GetClaimSeed(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "getClaimSeed", sender)
	return *ret0, err
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x7087ed2c.
//
// Solidity: function getClaimSeed(sender address) constant returns(uint256)
func (_SmartPool *SmartPoolSession) GetClaimSeed(sender common.Address) (*big.Int, error) {
	return _SmartPool.Contract.GetClaimSeed(&_SmartPool.CallOpts, sender)
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x7087ed2c.
//
// Solidity: function getClaimSeed(sender address) constant returns(uint256)
func (_SmartPool *SmartPoolCallerSession) GetClaimSeed(sender common.Address) (*big.Int, error) {
	return _SmartPool.Contract.GetClaimSeed(&_SmartPool.CallOpts, sender)
}

// GetMinerId is a free data retrieval call binding the contract method 0xe2dea715.
//
// Solidity: function getMinerId(sender address) constant returns(bytes32)
func (_SmartPool *SmartPoolCaller) GetMinerId(opts *bind.CallOpts, sender common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "getMinerId", sender)
	return *ret0, err
}

// GetMinerId is a free data retrieval call binding the contract method 0xe2dea715.
//
// Solidity: function getMinerId(sender address) constant returns(bytes32)
func (_SmartPool *SmartPoolSession) GetMinerId(sender common.Address) ([32]byte, error) {
	return _SmartPool.Contract.GetMinerId(&_SmartPool.CallOpts, sender)
}

// GetMinerId is a free data retrieval call binding the contract method 0xe2dea715.
//
// Solidity: function getMinerId(sender address) constant returns(bytes32)
func (_SmartPool *SmartPoolCallerSession) GetMinerId(sender common.Address) ([32]byte, error) {
	return _SmartPool.Contract.GetMinerId(&_SmartPool.CallOpts, sender)
}

// GetPoolETHBalance is a free data retrieval call binding the contract method 0x0abb8409.
//
// Solidity: function getPoolETHBalance() constant returns(uint256)
func (_SmartPool *SmartPoolCaller) GetPoolETHBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "getPoolETHBalance")
	return *ret0, err
}

// GetPoolETHBalance is a free data retrieval call binding the contract method 0x0abb8409.
//
// Solidity: function getPoolETHBalance() constant returns(uint256)
func (_SmartPool *SmartPoolSession) GetPoolETHBalance() (*big.Int, error) {
	return _SmartPool.Contract.GetPoolETHBalance(&_SmartPool.CallOpts)
}

// GetPoolETHBalance is a free data retrieval call binding the contract method 0x0abb8409.
//
// Solidity: function getPoolETHBalance() constant returns(uint256)
func (_SmartPool *SmartPoolCallerSession) GetPoolETHBalance() (*big.Int, error) {
	return _SmartPool.Contract.GetPoolETHBalance(&_SmartPool.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(sender address) constant returns(bool)
func (_SmartPool *SmartPoolCaller) IsRegistered(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "isRegistered", sender)
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(sender address) constant returns(bool)
func (_SmartPool *SmartPoolSession) IsRegistered(sender common.Address) (bool, error) {
	return _SmartPool.Contract.IsRegistered(&_SmartPool.CallOpts, sender)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(sender address) constant returns(bool)
func (_SmartPool *SmartPoolCallerSession) IsRegistered(sender common.Address) (bool, error) {
	return _SmartPool.Contract.IsRegistered(&_SmartPool.CallOpts, sender)
}

// NewVersionReleased is a free data retrieval call binding the contract method 0x0289e966.
//
// Solidity: function newVersionReleased() constant returns(bool)
func (_SmartPool *SmartPoolCaller) NewVersionReleased(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "newVersionReleased")
	return *ret0, err
}

// NewVersionReleased is a free data retrieval call binding the contract method 0x0289e966.
//
// Solidity: function newVersionReleased() constant returns(bool)
func (_SmartPool *SmartPoolSession) NewVersionReleased() (bool, error) {
	return _SmartPool.Contract.NewVersionReleased(&_SmartPool.CallOpts)
}

// NewVersionReleased is a free data retrieval call binding the contract method 0x0289e966.
//
// Solidity: function newVersionReleased() constant returns(bool)
func (_SmartPool *SmartPoolCallerSession) NewVersionReleased() (bool, error) {
	return _SmartPool.Contract.NewVersionReleased(&_SmartPool.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(bool)
func (_SmartPool *SmartPoolCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(bool)
func (_SmartPool *SmartPoolSession) Owners(arg0 common.Address) (bool, error) {
	return _SmartPool.Contract.Owners(&_SmartPool.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(bool)
func (_SmartPool *SmartPoolCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _SmartPool.Contract.Owners(&_SmartPool.CallOpts, arg0)
}

// PoolFees is a free data retrieval call binding the contract method 0x33580959.
//
// Solidity: function poolFees() constant returns(uint256)
func (_SmartPool *SmartPoolCaller) PoolFees(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "poolFees")
	return *ret0, err
}

// PoolFees is a free data retrieval call binding the contract method 0x33580959.
//
// Solidity: function poolFees() constant returns(uint256)
func (_SmartPool *SmartPoolSession) PoolFees() (*big.Int, error) {
	return _SmartPool.Contract.PoolFees(&_SmartPool.CallOpts)
}

// PoolFees is a free data retrieval call binding the contract method 0x33580959.
//
// Solidity: function poolFees() constant returns(uint256)
func (_SmartPool *SmartPoolCallerSession) PoolFees() (*big.Int, error) {
	return _SmartPool.Contract.PoolFees(&_SmartPool.CallOpts)
}

// To62Encoding is a free data retrieval call binding the contract method 0xff5d2c39.
//
// Solidity: function to62Encoding(id uint256, numChars uint256) constant returns(bytes32)
func (_SmartPool *SmartPoolCaller) To62Encoding(opts *bind.CallOpts, id *big.Int, numChars *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "to62Encoding", id, numChars)
	return *ret0, err
}

// To62Encoding is a free data retrieval call binding the contract method 0xff5d2c39.
//
// Solidity: function to62Encoding(id uint256, numChars uint256) constant returns(bytes32)
func (_SmartPool *SmartPoolSession) To62Encoding(id *big.Int, numChars *big.Int) ([32]byte, error) {
	return _SmartPool.Contract.To62Encoding(&_SmartPool.CallOpts, id, numChars)
}

// To62Encoding is a free data retrieval call binding the contract method 0xff5d2c39.
//
// Solidity: function to62Encoding(id uint256, numChars uint256) constant returns(bytes32)
func (_SmartPool *SmartPoolCallerSession) To62Encoding(id *big.Int, numChars *big.Int) ([32]byte, error) {
	return _SmartPool.Contract.To62Encoding(&_SmartPool.CallOpts, id, numChars)
}

// UncleRate is a free data retrieval call binding the contract method 0x1bc41284.
//
// Solidity: function uncleRate() constant returns(uint256)
func (_SmartPool *SmartPoolCaller) UncleRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "uncleRate")
	return *ret0, err
}

// UncleRate is a free data retrieval call binding the contract method 0x1bc41284.
//
// Solidity: function uncleRate() constant returns(uint256)
func (_SmartPool *SmartPoolSession) UncleRate() (*big.Int, error) {
	return _SmartPool.Contract.UncleRate(&_SmartPool.CallOpts)
}

// UncleRate is a free data retrieval call binding the contract method 0x1bc41284.
//
// Solidity: function uncleRate() constant returns(uint256)
func (_SmartPool *SmartPoolCallerSession) UncleRate() (*big.Int, error) {
	return _SmartPool.Contract.UncleRate(&_SmartPool.CallOpts)
}

// VerifySubmissionIndex is a free data retrieval call binding the contract method 0x9e133be4.
//
// Solidity: function verifySubmissionIndex(sender address, seed uint256, submissionNumber uint256, shareIndex uint256) constant returns(bool)
func (_SmartPool *SmartPoolCaller) VerifySubmissionIndex(opts *bind.CallOpts, sender common.Address, seed *big.Int, submissionNumber *big.Int, shareIndex *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "verifySubmissionIndex", sender, seed, submissionNumber, shareIndex)
	return *ret0, err
}

// VerifySubmissionIndex is a free data retrieval call binding the contract method 0x9e133be4.
//
// Solidity: function verifySubmissionIndex(sender address, seed uint256, submissionNumber uint256, shareIndex uint256) constant returns(bool)
func (_SmartPool *SmartPoolSession) VerifySubmissionIndex(sender common.Address, seed *big.Int, submissionNumber *big.Int, shareIndex *big.Int) (bool, error) {
	return _SmartPool.Contract.VerifySubmissionIndex(&_SmartPool.CallOpts, sender, seed, submissionNumber, shareIndex)
}

// VerifySubmissionIndex is a free data retrieval call binding the contract method 0x9e133be4.
//
// Solidity: function verifySubmissionIndex(sender address, seed uint256, submissionNumber uint256, shareIndex uint256) constant returns(bool)
func (_SmartPool *SmartPoolCallerSession) VerifySubmissionIndex(sender common.Address, seed *big.Int, submissionNumber *big.Int, shareIndex *big.Int) (bool, error) {
	return _SmartPool.Contract.VerifySubmissionIndex(&_SmartPool.CallOpts, sender, seed, submissionNumber, shareIndex)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_SmartPool *SmartPoolCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_SmartPool *SmartPoolSession) Version() (string, error) {
	return _SmartPool.Contract.Version(&_SmartPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_SmartPool *SmartPoolCallerSession) Version() (string, error) {
	return _SmartPool.Contract.Version(&_SmartPool.CallOpts)
}

// WhiteListEnabled is a free data retrieval call binding the contract method 0xe2e616bb.
//
// Solidity: function whiteListEnabled() constant returns(bool)
func (_SmartPool *SmartPoolCaller) WhiteListEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "whiteListEnabled")
	return *ret0, err
}

// WhiteListEnabled is a free data retrieval call binding the contract method 0xe2e616bb.
//
// Solidity: function whiteListEnabled() constant returns(bool)
func (_SmartPool *SmartPoolSession) WhiteListEnabled() (bool, error) {
	return _SmartPool.Contract.WhiteListEnabled(&_SmartPool.CallOpts)
}

// WhiteListEnabled is a free data retrieval call binding the contract method 0xe2e616bb.
//
// Solidity: function whiteListEnabled() constant returns(bool)
func (_SmartPool *SmartPoolCallerSession) WhiteListEnabled() (bool, error) {
	return _SmartPool.Contract.WhiteListEnabled(&_SmartPool.CallOpts)
}

// WithdrawalAddress is a free data retrieval call binding the contract method 0xf2bcd022.
//
// Solidity: function withdrawalAddress() constant returns(address)
func (_SmartPool *SmartPoolCaller) WithdrawalAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartPool.contract.Call(opts, out, "withdrawalAddress")
	return *ret0, err
}

// WithdrawalAddress is a free data retrieval call binding the contract method 0xf2bcd022.
//
// Solidity: function withdrawalAddress() constant returns(address)
func (_SmartPool *SmartPoolSession) WithdrawalAddress() (common.Address, error) {
	return _SmartPool.Contract.WithdrawalAddress(&_SmartPool.CallOpts)
}

// WithdrawalAddress is a free data retrieval call binding the contract method 0xf2bcd022.
//
// Solidity: function withdrawalAddress() constant returns(address)
func (_SmartPool *SmartPoolCallerSession) WithdrawalAddress() (common.Address, error) {
	return _SmartPool.Contract.WithdrawalAddress(&_SmartPool.CallOpts)
}

// DebugResetSubmissions is a paid mutator transaction binding the contract method 0x5eadd607.
//
// Solidity: function debugResetSubmissions() returns()
func (_SmartPool *SmartPoolTransactor) DebugResetSubmissions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartPool.contract.Transact(opts, "debugResetSubmissions")
}

// DebugResetSubmissions is a paid mutator transaction binding the contract method 0x5eadd607.
//
// Solidity: function debugResetSubmissions() returns()
func (_SmartPool *SmartPoolSession) DebugResetSubmissions() (*types.Transaction, error) {
	return _SmartPool.Contract.DebugResetSubmissions(&_SmartPool.TransactOpts)
}

// DebugResetSubmissions is a paid mutator transaction binding the contract method 0x5eadd607.
//
// Solidity: function debugResetSubmissions() returns()
func (_SmartPool *SmartPoolTransactorSession) DebugResetSubmissions() (*types.Transaction, error) {
	return _SmartPool.Contract.DebugResetSubmissions(&_SmartPool.TransactOpts)
}

// DeclareNewerVersion is a paid mutator transaction binding the contract method 0xe3d86998.
//
// Solidity: function declareNewerVersion() returns()
func (_SmartPool *SmartPoolTransactor) DeclareNewerVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartPool.contract.Transact(opts, "declareNewerVersion")
}

// DeclareNewerVersion is a paid mutator transaction binding the contract method 0xe3d86998.
//
// Solidity: function declareNewerVersion() returns()
func (_SmartPool *SmartPoolSession) DeclareNewerVersion() (*types.Transaction, error) {
	return _SmartPool.Contract.DeclareNewerVersion(&_SmartPool.TransactOpts)
}

// DeclareNewerVersion is a paid mutator transaction binding the contract method 0xe3d86998.
//
// Solidity: function declareNewerVersion() returns()
func (_SmartPool *SmartPoolTransactorSession) DeclareNewerVersion() (*types.Transaction, error) {
	return _SmartPool.Contract.DeclareNewerVersion(&_SmartPool.TransactOpts)
}

// GetShareIndexDebugForTestRPC is a paid mutator transaction binding the contract method 0x84c65296.
//
// Solidity: function getShareIndexDebugForTestRPC(sender address) returns()
func (_SmartPool *SmartPoolTransactor) GetShareIndexDebugForTestRPC(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _SmartPool.contract.Transact(opts, "getShareIndexDebugForTestRPC", sender)
}

// GetShareIndexDebugForTestRPC is a paid mutator transaction binding the contract method 0x84c65296.
//
// Solidity: function getShareIndexDebugForTestRPC(sender address) returns()
func (_SmartPool *SmartPoolSession) GetShareIndexDebugForTestRPC(sender common.Address) (*types.Transaction, error) {
	return _SmartPool.Contract.GetShareIndexDebugForTestRPC(&_SmartPool.TransactOpts, sender)
}

// GetShareIndexDebugForTestRPC is a paid mutator transaction binding the contract method 0x84c65296.
//
// Solidity: function getShareIndexDebugForTestRPC(sender address) returns()
func (_SmartPool *SmartPoolTransactorSession) GetShareIndexDebugForTestRPC(sender common.Address) (*types.Transaction, error) {
	return _SmartPool.Contract.GetShareIndexDebugForTestRPC(&_SmartPool.TransactOpts, sender)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(paymentAddress address) returns()
func (_SmartPool *SmartPoolTransactor) Register(opts *bind.TransactOpts, paymentAddress common.Address) (*types.Transaction, error) {
	return _SmartPool.contract.Transact(opts, "register", paymentAddress)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(paymentAddress address) returns()
func (_SmartPool *SmartPoolSession) Register(paymentAddress common.Address) (*types.Transaction, error) {
	return _SmartPool.Contract.Register(&_SmartPool.TransactOpts, paymentAddress)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(paymentAddress address) returns()
func (_SmartPool *SmartPoolTransactorSession) Register(paymentAddress common.Address) (*types.Transaction, error) {
	return _SmartPool.Contract.Register(&_SmartPool.TransactOpts, paymentAddress)
}

// SetUnlceRateAndFees is a paid mutator transaction binding the contract method 0x46b3f696.
//
// Solidity: function setUnlceRateAndFees(_uncleRate uint256, _poolFees uint256) returns()
func (_SmartPool *SmartPoolTransactor) SetUnlceRateAndFees(opts *bind.TransactOpts, _uncleRate *big.Int, _poolFees *big.Int) (*types.Transaction, error) {
	return _SmartPool.contract.Transact(opts, "setUnlceRateAndFees", _uncleRate, _poolFees)
}

// SetUnlceRateAndFees is a paid mutator transaction binding the contract method 0x46b3f696.
//
// Solidity: function setUnlceRateAndFees(_uncleRate uint256, _poolFees uint256) returns()
func (_SmartPool *SmartPoolSession) SetUnlceRateAndFees(_uncleRate *big.Int, _poolFees *big.Int) (*types.Transaction, error) {
	return _SmartPool.Contract.SetUnlceRateAndFees(&_SmartPool.TransactOpts, _uncleRate, _poolFees)
}

// SetUnlceRateAndFees is a paid mutator transaction binding the contract method 0x46b3f696.
//
// Solidity: function setUnlceRateAndFees(_uncleRate uint256, _poolFees uint256) returns()
func (_SmartPool *SmartPoolTransactorSession) SetUnlceRateAndFees(_uncleRate *big.Int, _poolFees *big.Int) (*types.Transaction, error) {
	return _SmartPool.Contract.SetUnlceRateAndFees(&_SmartPool.TransactOpts, _uncleRate, _poolFees)
}

// StoreClaimSeed is a paid mutator transaction binding the contract method 0xc7876940.
//
// Solidity: function storeClaimSeed(miner address) returns()
func (_SmartPool *SmartPoolTransactor) StoreClaimSeed(opts *bind.TransactOpts, miner common.Address) (*types.Transaction, error) {
	return _SmartPool.contract.Transact(opts, "storeClaimSeed", miner)
}

// StoreClaimSeed is a paid mutator transaction binding the contract method 0xc7876940.
//
// Solidity: function storeClaimSeed(miner address) returns()
func (_SmartPool *SmartPoolSession) StoreClaimSeed(miner common.Address) (*types.Transaction, error) {
	return _SmartPool.Contract.StoreClaimSeed(&_SmartPool.TransactOpts, miner)
}

// StoreClaimSeed is a paid mutator transaction binding the contract method 0xc7876940.
//
// Solidity: function storeClaimSeed(miner address) returns()
func (_SmartPool *SmartPoolTransactorSession) StoreClaimSeed(miner common.Address) (*types.Transaction, error) {
	return _SmartPool.Contract.StoreClaimSeed(&_SmartPool.TransactOpts, miner)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xcedb217a.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augRoot uint256, lastClaimBeforeVerification bool) returns()
func (_SmartPool *SmartPoolTransactor) SubmitClaim(opts *bind.TransactOpts, numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augRoot *big.Int, lastClaimBeforeVerification bool) (*types.Transaction, error) {
	return _SmartPool.contract.Transact(opts, "submitClaim", numShares, difficulty, min, max, augRoot, lastClaimBeforeVerification)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xcedb217a.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augRoot uint256, lastClaimBeforeVerification bool) returns()
func (_SmartPool *SmartPoolSession) SubmitClaim(numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augRoot *big.Int, lastClaimBeforeVerification bool) (*types.Transaction, error) {
	return _SmartPool.Contract.SubmitClaim(&_SmartPool.TransactOpts, numShares, difficulty, min, max, augRoot, lastClaimBeforeVerification)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xcedb217a.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augRoot uint256, lastClaimBeforeVerification bool) returns()
func (_SmartPool *SmartPoolTransactorSession) SubmitClaim(numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augRoot *big.Int, lastClaimBeforeVerification bool) (*types.Transaction, error) {
	return _SmartPool.Contract.SubmitClaim(&_SmartPool.TransactOpts, numShares, difficulty, min, max, augRoot, lastClaimBeforeVerification)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0xac1d0609.
//
// Solidity: function updateWhiteList(miner address, add bool) returns()
func (_SmartPool *SmartPoolTransactor) UpdateWhiteList(opts *bind.TransactOpts, miner common.Address, add bool) (*types.Transaction, error) {
	return _SmartPool.contract.Transact(opts, "updateWhiteList", miner, add)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0xac1d0609.
//
// Solidity: function updateWhiteList(miner address, add bool) returns()
func (_SmartPool *SmartPoolSession) UpdateWhiteList(miner common.Address, add bool) (*types.Transaction, error) {
	return _SmartPool.Contract.UpdateWhiteList(&_SmartPool.TransactOpts, miner, add)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0xac1d0609.
//
// Solidity: function updateWhiteList(miner address, add bool) returns()
func (_SmartPool *SmartPoolTransactorSession) UpdateWhiteList(miner common.Address, add bool) (*types.Transaction, error) {
	return _SmartPool.Contract.UpdateWhiteList(&_SmartPool.TransactOpts, miner, add)
}

// VerifyAgtDebugForTesting is a paid mutator transaction binding the contract method 0x7ed996a3.
//
// Solidity: function verifyAgtDebugForTesting(rootHash uint256, rootMin uint256, rootMax uint256, leafHash uint256, leafCounter uint256, branchIndex uint256, countersBranch uint256[], hashesBranch uint256[]) returns(bool)
func (_SmartPool *SmartPoolTransactor) VerifyAgtDebugForTesting(opts *bind.TransactOpts, rootHash *big.Int, rootMin *big.Int, rootMax *big.Int, leafHash *big.Int, leafCounter *big.Int, branchIndex *big.Int, countersBranch []*big.Int, hashesBranch []*big.Int) (*types.Transaction, error) {
	return _SmartPool.contract.Transact(opts, "verifyAgtDebugForTesting", rootHash, rootMin, rootMax, leafHash, leafCounter, branchIndex, countersBranch, hashesBranch)
}

// VerifyAgtDebugForTesting is a paid mutator transaction binding the contract method 0x7ed996a3.
//
// Solidity: function verifyAgtDebugForTesting(rootHash uint256, rootMin uint256, rootMax uint256, leafHash uint256, leafCounter uint256, branchIndex uint256, countersBranch uint256[], hashesBranch uint256[]) returns(bool)
func (_SmartPool *SmartPoolSession) VerifyAgtDebugForTesting(rootHash *big.Int, rootMin *big.Int, rootMax *big.Int, leafHash *big.Int, leafCounter *big.Int, branchIndex *big.Int, countersBranch []*big.Int, hashesBranch []*big.Int) (*types.Transaction, error) {
	return _SmartPool.Contract.VerifyAgtDebugForTesting(&_SmartPool.TransactOpts, rootHash, rootMin, rootMax, leafHash, leafCounter, branchIndex, countersBranch, hashesBranch)
}

// VerifyAgtDebugForTesting is a paid mutator transaction binding the contract method 0x7ed996a3.
//
// Solidity: function verifyAgtDebugForTesting(rootHash uint256, rootMin uint256, rootMax uint256, leafHash uint256, leafCounter uint256, branchIndex uint256, countersBranch uint256[], hashesBranch uint256[]) returns(bool)
func (_SmartPool *SmartPoolTransactorSession) VerifyAgtDebugForTesting(rootHash *big.Int, rootMin *big.Int, rootMax *big.Int, leafHash *big.Int, leafCounter *big.Int, branchIndex *big.Int, countersBranch []*big.Int, hashesBranch []*big.Int) (*types.Transaction, error) {
	return _SmartPool.Contract.VerifyAgtDebugForTesting(&_SmartPool.TransactOpts, rootHash, rootMin, rootMax, leafHash, leafCounter, branchIndex, countersBranch, hashesBranch)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0xc2209acb.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, submissionIndex uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns()
func (_SmartPool *SmartPoolTransactor) VerifyClaim(opts *bind.TransactOpts, rlpHeader []byte, nonce *big.Int, submissionIndex *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _SmartPool.contract.Transact(opts, "verifyClaim", rlpHeader, nonce, submissionIndex, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0xc2209acb.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, submissionIndex uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns()
func (_SmartPool *SmartPoolSession) VerifyClaim(rlpHeader []byte, nonce *big.Int, submissionIndex *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _SmartPool.Contract.VerifyClaim(&_SmartPool.TransactOpts, rlpHeader, nonce, submissionIndex, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0xc2209acb.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, submissionIndex uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns()
func (_SmartPool *SmartPoolTransactorSession) VerifyClaim(rlpHeader []byte, nonce *big.Int, submissionIndex *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _SmartPool.Contract.VerifyClaim(&_SmartPool.TransactOpts, rlpHeader, nonce, submissionIndex, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_SmartPool *SmartPoolTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SmartPool.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_SmartPool *SmartPoolSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _SmartPool.Contract.Withdraw(&_SmartPool.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_SmartPool *SmartPoolTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _SmartPool.Contract.Withdraw(&_SmartPool.TransactOpts, amount)
}
