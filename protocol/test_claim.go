package protocol

import (
	"github.com/SmartPool/smartpool-client"
	"math/big"
)

type testClaim struct {
	c []smartpool.Share
}

func (c *testClaim) NumShares() *big.Int {
	return big.NewInt(int64(len(c.c)))
}

func (c *testClaim) Difficulty() *big.Int {
	return big.NewInt(100000)
}

func (c *testClaim) Min() *big.Int {
	return big.NewInt(0)
}

func (c *testClaim) Max() *big.Int {
	return big.NewInt(100)
}

func (c *testClaim) GetShare(index int) smartpool.Share {
	return c.c[index]
}

func (c *testClaim) SetEvidence(shareIndex *big.Int) {
}

func (c *testClaim) AugMerkle() smartpool.SPHash {
	return smartpool.SPHash{}
}

func (c *testClaim) CounterBranch() []*big.Int {
	return []*big.Int{}
}

func (c *testClaim) HashBranch() []*big.Int {
	return []*big.Int{}
}
