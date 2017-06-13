package ethereum

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Contract struct {
	client ContractClient
	miner  common.Address
}

func (c *Contract) Version() string {
	return c.client.Version()
}

func (c *Contract) IsRegistered() bool {
	return c.client.IsRegistered()
}

func (c *Contract) CanRegister() bool {
	return c.client.CanRegister()
}

func (c *Contract) Register(paymentAddress common.Address) error {
	return c.client.Register(paymentAddress)
}

func (c *Contract) SubmitClaim(claim smartpool.Claim, lastClaim bool) error {
	smartpool.Output.Printf("Min: 0x%s - Max: 0x%s - Diff: 0x%s\n", claim.Min().Text(16), claim.Max().Text(16), claim.Difficulty().Text(16))
	return c.client.SubmitClaim(
		claim.NumShares(), claim.Difficulty(),
		claim.Min(), claim.Max(), claim.AugMerkle().Big(), lastClaim)
}

func (c *Contract) GetShareIndex(claim smartpool.Claim) (*big.Int, *big.Int, error) {
	seed := c.client.GetClaimSeed()
	data, err := c.client.CalculateSubmissionIndex(c.miner, seed)
	return data[0], data[1], err
}

func (c *Contract) NumOpenClaims() (*big.Int, error) {
	return c.client.NumOpenClaims(c.miner)
}

func (c *Contract) ResetOpenClaims() error {
	return c.client.ResetOpenClaims()
}

func (c *Contract) VerifyClaim(submissionIndex *big.Int, shareIndex *big.Int, claim smartpool.Claim) error {
	share := claim.GetShare(int(shareIndex.Int64())).(*Share)
	rlpHeader, _ := share.RlpHeaderWithoutNonce()
	nonce := share.NonceBig()
	claim.SetEvidence(shareIndex)
	augCountersBranch := claim.CounterBranch()
	augHashesBranch := claim.HashBranch()
	dataSetLookup := share.DAGElementArray()
	witnessForLookup := share.DAGProofArray()
	return c.client.VerifyClaim(
		rlpHeader,
		nonce,
		submissionIndex,
		shareIndex,
		dataSetLookup,
		witnessForLookup,
		augCountersBranch,
		augHashesBranch,
	)
}

func NewContract(client ContractClient, miner common.Address) *Contract {
	return &Contract{client, miner}
}
