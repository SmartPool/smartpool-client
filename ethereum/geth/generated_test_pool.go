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
const TestPoolABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"newVersionReleased\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getShareIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isRegistered\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMinerId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rlpHeader\",\"type\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"shareIndex\",\"type\":\"uint256\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"},{\"name\":\"augCountersBranch\",\"type\":\"uint256[]\"},{\"name\":\"augHashesBranch\",\"type\":\"uint256[]\"}],\"name\":\"verifyClaim\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"modul\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"paymentAddress\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getClaimSeed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canRegister\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes32\"},{\"name\":\"nonceLe\",\"type\":\"bytes8\"},{\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint256\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"},{\"name\":\"branchSize\",\"type\":\"uint256\"},{\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"hashimoto\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"rlpHeader\",\"type\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"shareIndex\",\"type\":\"uint256\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"},{\"name\":\"augCountersBranch\",\"type\":\"uint256[]\"},{\"name\":\"augHashesBranch\",\"type\":\"uint256[]\"}],\"name\":\"verifyClaim_debug\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochData\",\"outputs\":[{\"name\":\"merkleRoot\",\"type\":\"uint128\"},{\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint64\"},{\"name\":\"branchDepth\",\"type\":\"uint64\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"merkleRoot\",\"type\":\"uint128\"},{\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint64\"},{\"name\":\"branchDepth\",\"type\":\"uint64\"},{\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"setEpochData\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"extraData\",\"type\":\"bytes32\"},{\"name\":\"minerId\",\"type\":\"bytes32\"},{\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"verifyExtraData\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"extraData\",\"type\":\"bytes32\"},{\"name\":\"minerId\",\"type\":\"bytes32\"},{\"name\":\"difficulty\",\"type\":\"uint256\"}],\"name\":\"verifyExtraData_debug\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"declareNewerVersion\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numShares\",\"type\":\"uint256\"},{\"name\":\"difficulty\",\"type\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\"},{\"name\":\"augMerkle\",\"type\":\"uint256\"}],\"name\":\"submitClaim\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"existingIds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"numChars\",\"type\":\"uint256\"}],\"name\":\"to62Encoding\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"Debug\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"ErrorLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"VerifyAgt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"Log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"Result\",\"type\":\"event\"}]"

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

// CanRegister is a free data retrieval call binding the contract method 0x589384b9.
//
// Solidity: function canRegister() constant returns(bool)
func (_TestPool *TestPoolCaller) CanRegister(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "canRegister")
	return *ret0, err
}

// CanRegister is a free data retrieval call binding the contract method 0x589384b9.
//
// Solidity: function canRegister() constant returns(bool)
func (_TestPool *TestPoolSession) CanRegister() (bool, error) {
	return _TestPool.Contract.CanRegister(&_TestPool.CallOpts)
}

// CanRegister is a free data retrieval call binding the contract method 0x589384b9.
//
// Solidity: function canRegister() constant returns(bool)
func (_TestPool *TestPoolCallerSession) CanRegister() (bool, error) {
	return _TestPool.Contract.CanRegister(&_TestPool.CallOpts)
}

// EpochData is a free data retrieval call binding the contract method 0x6e821b2e.
//
// Solidity: function epochData( uint256) constant returns(merkleRoot uint128, fullSizeIn128Resultion uint64, branchDepth uint64)
func (_TestPool *TestPoolCaller) EpochData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MerkleRoot             *big.Int
	FullSizeIn128Resultion uint64
	BranchDepth            uint64
}, error) {
	ret := new(struct {
		MerkleRoot             *big.Int
		FullSizeIn128Resultion uint64
		BranchDepth            uint64
	})
	out := ret
	err := _TestPool.contract.Call(opts, out, "epochData", arg0)
	return *ret, err
}

// EpochData is a free data retrieval call binding the contract method 0x6e821b2e.
//
// Solidity: function epochData( uint256) constant returns(merkleRoot uint128, fullSizeIn128Resultion uint64, branchDepth uint64)
func (_TestPool *TestPoolSession) EpochData(arg0 *big.Int) (struct {
	MerkleRoot             *big.Int
	FullSizeIn128Resultion uint64
	BranchDepth            uint64
}, error) {
	return _TestPool.Contract.EpochData(&_TestPool.CallOpts, arg0)
}

// EpochData is a free data retrieval call binding the contract method 0x6e821b2e.
//
// Solidity: function epochData( uint256) constant returns(merkleRoot uint128, fullSizeIn128Resultion uint64, branchDepth uint64)
func (_TestPool *TestPoolCallerSession) EpochData(arg0 *big.Int) (struct {
	MerkleRoot             *big.Int
	FullSizeIn128Resultion uint64
	BranchDepth            uint64
}, error) {
	return _TestPool.Contract.EpochData(&_TestPool.CallOpts, arg0)
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

// GetClaimSeed is a free data retrieval call binding the contract method 0x47740bdc.
//
// Solidity: function getClaimSeed() constant returns(uint256)
func (_TestPool *TestPoolCaller) GetClaimSeed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "getClaimSeed")
	return *ret0, err
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x47740bdc.
//
// Solidity: function getClaimSeed() constant returns(uint256)
func (_TestPool *TestPoolSession) GetClaimSeed() (*big.Int, error) {
	return _TestPool.Contract.GetClaimSeed(&_TestPool.CallOpts)
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x47740bdc.
//
// Solidity: function getClaimSeed() constant returns(uint256)
func (_TestPool *TestPoolCallerSession) GetClaimSeed() (*big.Int, error) {
	return _TestPool.Contract.GetClaimSeed(&_TestPool.CallOpts)
}

// GetMinerId is a free data retrieval call binding the contract method 0x24e74bdf.
//
// Solidity: function getMinerId() constant returns(bytes32)
func (_TestPool *TestPoolCaller) GetMinerId(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "getMinerId")
	return *ret0, err
}

// GetMinerId is a free data retrieval call binding the contract method 0x24e74bdf.
//
// Solidity: function getMinerId() constant returns(bytes32)
func (_TestPool *TestPoolSession) GetMinerId() ([32]byte, error) {
	return _TestPool.Contract.GetMinerId(&_TestPool.CallOpts)
}

// GetMinerId is a free data retrieval call binding the contract method 0x24e74bdf.
//
// Solidity: function getMinerId() constant returns(bytes32)
func (_TestPool *TestPoolCallerSession) GetMinerId() ([32]byte, error) {
	return _TestPool.Contract.GetMinerId(&_TestPool.CallOpts)
}

// GetShareIndex is a free data retrieval call binding the contract method 0x1f7bc1e3.
//
// Solidity: function getShareIndex() constant returns(uint256)
func (_TestPool *TestPoolCaller) GetShareIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "getShareIndex")
	return *ret0, err
}

// GetShareIndex is a free data retrieval call binding the contract method 0x1f7bc1e3.
//
// Solidity: function getShareIndex() constant returns(uint256)
func (_TestPool *TestPoolSession) GetShareIndex() (*big.Int, error) {
	return _TestPool.Contract.GetShareIndex(&_TestPool.CallOpts)
}

// GetShareIndex is a free data retrieval call binding the contract method 0x1f7bc1e3.
//
// Solidity: function getShareIndex() constant returns(uint256)
func (_TestPool *TestPoolCallerSession) GetShareIndex() (*big.Int, error) {
	return _TestPool.Contract.GetShareIndex(&_TestPool.CallOpts)
}

// Hashimoto is a free data retrieval call binding the contract method 0x58e69c5a.
//
// Solidity: function hashimoto(header bytes32, nonceLe bytes8, fullSizeIn128Resultion uint256, dataSetLookup uint256[], witnessForLookup uint256[], branchSize uint256, root uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) Hashimoto(opts *bind.CallOpts, header [32]byte, nonceLe [8]byte, fullSizeIn128Resultion *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, branchSize *big.Int, root *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "hashimoto", header, nonceLe, fullSizeIn128Resultion, dataSetLookup, witnessForLookup, branchSize, root)
	return *ret0, err
}

// Hashimoto is a free data retrieval call binding the contract method 0x58e69c5a.
//
// Solidity: function hashimoto(header bytes32, nonceLe bytes8, fullSizeIn128Resultion uint256, dataSetLookup uint256[], witnessForLookup uint256[], branchSize uint256, root uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) Hashimoto(header [32]byte, nonceLe [8]byte, fullSizeIn128Resultion *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, branchSize *big.Int, root *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Hashimoto(&_TestPool.CallOpts, header, nonceLe, fullSizeIn128Resultion, dataSetLookup, witnessForLookup, branchSize, root)
}

// Hashimoto is a free data retrieval call binding the contract method 0x58e69c5a.
//
// Solidity: function hashimoto(header bytes32, nonceLe bytes8, fullSizeIn128Resultion uint256, dataSetLookup uint256[], witnessForLookup uint256[], branchSize uint256, root uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) Hashimoto(header [32]byte, nonceLe [8]byte, fullSizeIn128Resultion *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, branchSize *big.Int, root *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Hashimoto(&_TestPool.CallOpts, header, nonceLe, fullSizeIn128Resultion, dataSetLookup, witnessForLookup, branchSize, root)
}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() constant returns(bool)
func (_TestPool *TestPoolCaller) IsRegistered(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "isRegistered")
	return *ret0, err
}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() constant returns(bool)
func (_TestPool *TestPoolSession) IsRegistered() (bool, error) {
	return _TestPool.Contract.IsRegistered(&_TestPool.CallOpts)
}

// IsRegistered is a free data retrieval call binding the contract method 0x22366844.
//
// Solidity: function isRegistered() constant returns(bool)
func (_TestPool *TestPoolCallerSession) IsRegistered() (bool, error) {
	return _TestPool.Contract.IsRegistered(&_TestPool.CallOpts)
}

// Modul is a free data retrieval call binding the contract method 0x426dd4e3.
//
// Solidity: function modul(x uint256, y uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) Modul(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "modul", x, y)
	return *ret0, err
}

// Modul is a free data retrieval call binding the contract method 0x426dd4e3.
//
// Solidity: function modul(x uint256, y uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) Modul(x *big.Int, y *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Modul(&_TestPool.CallOpts, x, y)
}

// Modul is a free data retrieval call binding the contract method 0x426dd4e3.
//
// Solidity: function modul(x uint256, y uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) Modul(x *big.Int, y *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Modul(&_TestPool.CallOpts, x, y)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TestPool *TestPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TestPool *TestPoolSession) Owner() (common.Address, error) {
	return _TestPool.Contract.Owner(&_TestPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TestPool *TestPoolCallerSession) Owner() (common.Address, error) {
	return _TestPool.Contract.Owner(&_TestPool.CallOpts)
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

// VerifyClaim_debug is a free data retrieval call binding the contract method 0x60e5b418.
//
// Solidity: function verifyClaim_debug(rlpHeader bytes, nonce uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) constant returns(uint256)
func (_TestPool *TestPoolCaller) VerifyClaim_debug(opts *bind.CallOpts, rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "verifyClaim_debug", rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
	return *ret0, err
}

// VerifyClaim_debug is a free data retrieval call binding the contract method 0x60e5b418.
//
// Solidity: function verifyClaim_debug(rlpHeader bytes, nonce uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) constant returns(uint256)
func (_TestPool *TestPoolSession) VerifyClaim_debug(rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*big.Int, error) {
	return _TestPool.Contract.VerifyClaim_debug(&_TestPool.CallOpts, rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// VerifyClaim_debug is a free data retrieval call binding the contract method 0x60e5b418.
//
// Solidity: function verifyClaim_debug(rlpHeader bytes, nonce uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) VerifyClaim_debug(rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*big.Int, error) {
	return _TestPool.Contract.VerifyClaim_debug(&_TestPool.CallOpts, rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// VerifyExtraData is a free data retrieval call binding the contract method 0xa442d820.
//
// Solidity: function verifyExtraData(extraData bytes32, minerId bytes32, difficulty uint256) constant returns(bool)
func (_TestPool *TestPoolCaller) VerifyExtraData(opts *bind.CallOpts, extraData [32]byte, minerId [32]byte, difficulty *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "verifyExtraData", extraData, minerId, difficulty)
	return *ret0, err
}

// VerifyExtraData is a free data retrieval call binding the contract method 0xa442d820.
//
// Solidity: function verifyExtraData(extraData bytes32, minerId bytes32, difficulty uint256) constant returns(bool)
func (_TestPool *TestPoolSession) VerifyExtraData(extraData [32]byte, minerId [32]byte, difficulty *big.Int) (bool, error) {
	return _TestPool.Contract.VerifyExtraData(&_TestPool.CallOpts, extraData, minerId, difficulty)
}

// VerifyExtraData is a free data retrieval call binding the contract method 0xa442d820.
//
// Solidity: function verifyExtraData(extraData bytes32, minerId bytes32, difficulty uint256) constant returns(bool)
func (_TestPool *TestPoolCallerSession) VerifyExtraData(extraData [32]byte, minerId [32]byte, difficulty *big.Int) (bool, error) {
	return _TestPool.Contract.VerifyExtraData(&_TestPool.CallOpts, extraData, minerId, difficulty)
}

// VerifyExtraData_debug is a free data retrieval call binding the contract method 0xe33a97d8.
//
// Solidity: function verifyExtraData_debug(extraData bytes32, minerId bytes32, difficulty uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) VerifyExtraData_debug(opts *bind.CallOpts, extraData [32]byte, minerId [32]byte, difficulty *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "verifyExtraData_debug", extraData, minerId, difficulty)
	return *ret0, err
}

// VerifyExtraData_debug is a free data retrieval call binding the contract method 0xe33a97d8.
//
// Solidity: function verifyExtraData_debug(extraData bytes32, minerId bytes32, difficulty uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) VerifyExtraData_debug(extraData [32]byte, minerId [32]byte, difficulty *big.Int) (*big.Int, error) {
	return _TestPool.Contract.VerifyExtraData_debug(&_TestPool.CallOpts, extraData, minerId, difficulty)
}

// VerifyExtraData_debug is a free data retrieval call binding the contract method 0xe33a97d8.
//
// Solidity: function verifyExtraData_debug(extraData bytes32, minerId bytes32, difficulty uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) VerifyExtraData_debug(extraData [32]byte, minerId [32]byte, difficulty *big.Int) (*big.Int, error) {
	return _TestPool.Contract.VerifyExtraData_debug(&_TestPool.CallOpts, extraData, minerId, difficulty)
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

// SetEpochData is a paid mutator transaction binding the contract method 0x820a22c0.
//
// Solidity: function setEpochData(merkleRoot uint128, fullSizeIn128Resultion uint64, branchDepth uint64, epoch uint256) returns()
func (_TestPool *TestPoolTransactor) SetEpochData(opts *bind.TransactOpts, merkleRoot *big.Int, fullSizeIn128Resultion uint64, branchDepth uint64, epoch *big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "setEpochData", merkleRoot, fullSizeIn128Resultion, branchDepth, epoch)
}

// SetEpochData is a paid mutator transaction binding the contract method 0x820a22c0.
//
// Solidity: function setEpochData(merkleRoot uint128, fullSizeIn128Resultion uint64, branchDepth uint64, epoch uint256) returns()
func (_TestPool *TestPoolSession) SetEpochData(merkleRoot *big.Int, fullSizeIn128Resultion uint64, branchDepth uint64, epoch *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SetEpochData(&_TestPool.TransactOpts, merkleRoot, fullSizeIn128Resultion, branchDepth, epoch)
}

// SetEpochData is a paid mutator transaction binding the contract method 0x820a22c0.
//
// Solidity: function setEpochData(merkleRoot uint128, fullSizeIn128Resultion uint64, branchDepth uint64, epoch uint256) returns()
func (_TestPool *TestPoolTransactorSession) SetEpochData(merkleRoot *big.Int, fullSizeIn128Resultion uint64, branchDepth uint64, epoch *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SetEpochData(&_TestPool.TransactOpts, merkleRoot, fullSizeIn128Resultion, branchDepth, epoch)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xe7dac983.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augMerkle uint256) returns()
func (_TestPool *TestPoolTransactor) SubmitClaim(opts *bind.TransactOpts, numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augMerkle *big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "submitClaim", numShares, difficulty, min, max, augMerkle)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xe7dac983.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augMerkle uint256) returns()
func (_TestPool *TestPoolSession) SubmitClaim(numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augMerkle *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SubmitClaim(&_TestPool.TransactOpts, numShares, difficulty, min, max, augMerkle)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xe7dac983.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augMerkle uint256) returns()
func (_TestPool *TestPoolTransactorSession) SubmitClaim(numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augMerkle *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SubmitClaim(&_TestPool.TransactOpts, numShares, difficulty, min, max, augMerkle)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0x35ffbe74.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns(uint256)
func (_TestPool *TestPoolTransactor) VerifyClaim(opts *bind.TransactOpts, rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "verifyClaim", rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0x35ffbe74.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns(uint256)
func (_TestPool *TestPoolSession) VerifyClaim(rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.VerifyClaim(&_TestPool.TransactOpts, rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0x35ffbe74.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns(uint256)
func (_TestPool *TestPoolTransactorSession) VerifyClaim(rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.VerifyClaim(&_TestPool.TransactOpts, rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}
