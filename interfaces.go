// Package smartpool defines interfaces for interaction between SmartPool
// and user, SmartPool and external resources such as Ethereum client
// (geth, partity), ethminer, persistent storage.
//
// It also defines some core interfaces for its sub packages to interact
// to each other.
package smartpool

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

// UserInput represents all necessary user specific inputs to run SmartPool
// client. Some of them can have default values depend on the actual structs
// implementing the interface.
type UserInput interface {
	IPCPath() string
	KeystorePath() string
	NoSharePerClaim() uint32
	ShareDifficulty() *big.Int
	SubmitInterval() time.Duration
	ContractAddress() string
	MinerAddress() string
	ExtraData() string
	Passphrase() string
}

// UserOutput accepts all the information that SmartPool wants to tell the user.
// It's only responsibility is to accept information. How the information is
// delivered to user is upto structs implementing the interface.
type UserOutput interface {
	// TODO: Add necessary methods
	// TODO: This might take information about internal detail such as:
	// number of claim submitted, number of claim verified,
	// number of claim accepted, number of share per claim on average,
	// average hash rate,...
	Printf(format string, a ...interface{}) (n int, err error)
}

// PersistentStorage is the gateway for smartpool to interact with external
// persistent storage such as a file system, a database or even a cloud based
// service.
// Smartpool should only persist something via this interface.
type PersistentStorage interface {
	// TODO: Add necessary methods
}

// DAGReader provides a way for smartpool to retrieve DAG dataset. How the DAG
// is retrieve is upto structs implementing the interface.
type DAGReader interface {
	// NextWord return next data chunk of the DAG dataset. First 8 bytes must
	// be ignored.
	NextWord() ([]byte, error)
}

// Contract is the interface for smartpool to interact with contract side of
// SmartPool protocol.
// Contract can be used for only one caller (Ethereum account) per
// instance.
type Contract interface {
	// Version return contract version which is useful for backward and forward
	// compatibility when the contract is redeployed in some occasions.
	Version() string
	// IsRegistered returns true when the miner's address is already recognized
	// as a user of the pool. It returns false otherwise.
	IsRegistered() bool
	// CanRegister returns true when the miner's address can actually register
	// to the pool. It returns false when the contract side decided to refuse
	// the address.
	CanRegister() bool
	// Register takes an address and register it to the pool.
	Register(paymentAddress common.Address) error
	// SubmitClaim takes some necessary parameters that represent a claim and
	// submit to the contract using miner's address. The address should be
	// unlocked first.
	SubmitClaim(claim Claim) error
	// GetShareIndex returns index of the share that is requested to submit
	// proof to the contract to represent correctness of the submitted claim.
	// GetShareIndex must be called after SubmitClaim to get shareIndex which
	// is used to pass to VerifyClaim. If GetShareIndex is called before
	// SubmitClaim, the index will have no meaning to contract.
	GetShareIndex() *big.Int
	// VerifyClaim takes some necessary parameters that provides complete proof
	// of a share with index shareIndex in the cliam and submit to contract side
	// in order to prove that the claim is valid so the miner can take credit
	// of it.
	VerifyClaim(shareIndex *big.Int, claim Claim) error

	// // SubmitClaim takes some necessary parameters that represent a claim and
	// // submit to the contract using miner's address. The address should be
	// // unlocked first.
	// // TODO: explain all parameters including examples.
	// SubmitClaim(
	// 	numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int,
	// 	augMerkle *big.Int) (*types.Transaction, error)
	// // VerifyClaim takes some necessary parameters that provides complete proof
	// // including ethash proof and augmented merkle tree proof and submit to
	// // contract side in order to prove that the claim is valid so the miner
	// // can take credit of it.
	// // TODO: explain all parameters including examples.
	// VerifyClaim(
	// 	rlpHeader []byte, nonce *big.Int, shareIndex *big.Int,
	// 	dataSetLookup []*big.Int, witnessForLookup []*big.Int,
	// 	augCountersBranch []*big.Int,
	// 	augHashesBranch []*big.Int) (*types.Transaction, error)
}

// NetworkClient represents client for blockchain network that miner is mining
// on. Network can be Ethereum, Ethereum Classic, ZCash, Bitcoin... For
// Ethereum, client can be Geth or Parity.
// Smartpool should only interact with network client via this interface and
// it doesn't care if the client is Geth or Partity or any other clients.
// Communication mechanism is upto structs implementing this interface.
type NetworkClient interface {
	// GetWork returns a Work for SmartPool to give to the miner. How the work
	// is formed is upto structs implementing this interface.
	GetWork() Work
	// SubmitSolution submits the solution that miner has submitted to SmartPool
	// so the full block solution can take credits. It also maintain workflow
	// between miner and the network client.
	SubmitSolution(s Solution) bool
}

// type EthereumClient interface {
// 	// GetPendingBlockHeader returns Ethereum block header for pending block
// 	// which potentially is the work that Ethereum client is wanting miner to
// 	// work on.
// 	GetPendingBlockHeader() *types.Header
// 	// GetBlockHeader returns Ethereum block for specific number.
// 	// TODO: this function should be removed because its redundant to actual
// 	// protocol.
// 	GetBlockHeader(number int) *types.Header
// 	// GetWork returns a Work for SmartPool to give to the miner. How the work
// 	// is formed is upto structs implementing this interface.
// 	GetWork() Work
// 	// SubmitHashrate submits the current hashrate to Ethereum Client in order
// 	// to maintain workflow between ethminer and Ethereum Client.
// 	SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool
// 	// SubmitWork submits the solution that miner has submitted to SmartPool
// 	// so the full block solution can take credits. It also maintain workflow
// 	// between ethminer and Ethereum Client.
// 	SubmitWork(nonce types.BlockNonce, hash, mixDigest common.Hash) bool
// 	// IsVerified returns true if the transaction with hash `h` in included in a
// 	// block. It returns false if the transaction is not yet included in any
// 	// blocks.
// 	IsVerified(h common.Hash) bool
// }

// ShareReceiver represents SmartPool itself which accepts solutions from
// miners.
type ShareReceiver interface {
	AcceptSolution(s Solution) Share
}
