package smartpool

import (
	"math/big"
)

// Work represents SmartPool work that miner needs to solve to have valid
// shares. Work is easier (has smaller difficulty) than actual Ethereum
// work that the miner can get from the network.
type Work interface {
	ID() string
	// AcceptSolution takes solution to construct and return a Share representing
	// the solution that came from miner.
	AcceptSolution(sol Solution) Share
}

// Solution represents a solution for a work
type Solution interface {
	// WorkID returns the ID to identify the work it is trying to solve
	WorkID() string
}

// Share represent a solution of a Work that comes from the miner.
type Share interface {
	// Counter returns the counter to be used in augmented merkle tree of a claim
	// which contains many shares. This counter must be increasing as shares
	// share coming. In other words, later share must have bigger counter.
	Counter() *big.Int
	// Hash return the hash of the share to be used as leaf hash of the augmented
	// merkle tree.
	Hash() []byte
}

// Claim represent a batch of shares which needs to reorganize its shares in
// ascending order of share counter.
type Claim interface {
	// NumShares returns number of shares that the claim is holding
	NumShares() *big.Int
	// Difficulty returns the min difficulty across all of its shares
	Difficulty() *big.Int
	// Min returns the min counter of the augmented merkle root
	Min() *big.Int
	// Max returns the max counter of the augmented merkle root
	Max() *big.Int
	// AugMerkle returns the hash of the augmented merkle root
	AugMerkle() []byte
}
