package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type ContractClient interface {
	Version() string
	IsRegistered() bool
	CanRegister() bool
	Register(paymentAddress common.Address) error
	GetClaimSeed() *big.Int
	SubmitClaim(
		numShares *big.Int,
		difficulty *big.Int,
		min *big.Int,
		max *big.Int,
		augMerkle *big.Int) error
	VerifyClaim(
		rlpHeader []byte,
		nonce *big.Int,
		shareIndex *big.Int,
		dataSetLookup []*big.Int,
		witnessForLookup []*big.Int,
		augCountersBranch []*big.Int,
		augHashesBranch []*big.Int) error
	SetEpochData(
		epoch *big.Int,
		fullSizeIn128Resolution *big.Int,
		branchDepth *big.Int,
		merkleNodes []*big.Int,
	) error
}
