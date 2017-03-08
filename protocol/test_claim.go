package protocol

import (
	"../"
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

func (c *testClaim) AugMerkle() smartpool.SPHash {
	return smartpool.SPHash{}
}
