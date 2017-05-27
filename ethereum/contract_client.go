package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type EthashContractClient interface {
	SetEpochData(
		epoch *big.Int,
		fullSizeIn128Resolution *big.Int,
		branchDepth *big.Int,
		merkleNodes []*big.Int,
	) error
}

type ContractClient interface {
	Version() string
	IsRegistered() bool
	CanRegister() bool
	Register(paymentAddress common.Address) error
	GetClaimSeed() *big.Int
	NumOpenClaims(sender common.Address) (*big.Int, error)
	ResetOpenClaims() error
	CalculateSubmissionIndex(sender common.Address, seed *big.Int) ([2]*big.Int, error)
	SubmitClaim(
		numShares *big.Int,
		difficulty *big.Int,
		min *big.Int,
		max *big.Int,
		augMerkle *big.Int,
		lastClaim bool) error
	VerifyClaim(
		rlpHeader []byte,
		nonce *big.Int,
		submission *big.Int,
		shareIndex *big.Int,
		dataSetLookup []*big.Int,
		witnessForLookup []*big.Int,
		augCountersBranch []*big.Int,
		augHashesBranch []*big.Int) error
}
