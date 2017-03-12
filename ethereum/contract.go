package ethereum

import (
	"../"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

type Contract struct {
	client ContractClient
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

func (c *Contract) SubmitClaim(claim smartpool.Claim) error {
	fmt.Printf("contract client: %v\n", c.client)
	return c.client.SubmitClaim(
		claim.NumShares(), claim.Difficulty(),
		claim.Min(), claim.Max(), claim.AugMerkle().Big())
}

func (c *Contract) GetShareIndex(claim smartpool.Claim) *big.Int {
	zero := big.NewInt(0)
	var seed *big.Int
	for {
		seed = c.client.GetClaimSeed()
		if seed.Cmp(zero) != 0 {
			break
		}
		time.Sleep(14 * time.Second)
	}
	index := big.NewInt(0)
	index.Mod(seed, claim.NumShares())
	return index
}

func (c *Contract) VerifyClaim(shareIndex *big.Int, claim smartpool.Claim) error {
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
		shareIndex,
		dataSetLookup,
		witnessForLookup,
		augCountersBranch,
		augHashesBranch,
	)
}

func NewContract(client ContractClient) *Contract {
	return &Contract{client}
}
