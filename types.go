package smartpool

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// Work represent SmartPool work that miner needs to solve to have valid
// shares. Work is easier (has smaller difficulty) than actual Ethereum
// work that the miner can get from Ethereum Client.
type Work interface {
	PoWHash() common.Hash
	SeedHash() string
	ShareDifficulty() *big.Int
	BlockHeader() *types.Header
	// AcceptSolution takes nonce and mixDigest to form a Share representing
	// the solution that came from miner.
	AcceptSolution(nonce types.BlockNonce, mixDigest common.Hash) Share
}

// Share represent a solution of a Work that comes from the miner.
type Share interface {
	Difficulty() *big.Int
	HashNoNonce() common.Hash
	Nonce() uint64
	MixDigest() common.Hash
	NumberU64() uint64
	NonceBig() *big.Int
	BlockHeader() *types.Header

	// RlpHeaderWithoutNonce is rlp encoded of all header of the share excluding
	// the nonce and mixDigest.
	RlpHeaderWithoutNonce() ([]byte, error)
	// Timestamp returns timestamp of the share itself.
	Timestamp() big.Int
	// Counter returns the counter to be used in augmented merkle tree of a claim
	// which contains many shares. This counter must be increasing as shares
	// share coming. In other words, later share must have bigger counter.
	Counter() *big.Int
	// Hash return the hash of the share to be used as leaf hash of the augmented
	// merkle tree.
	Hash() []byte
}
