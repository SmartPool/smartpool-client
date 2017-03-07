package protocol

import (
	"../"
	"math/big"
)

type Shares []smartpool.Share

func (s Shares) Len() int      { return len(s) }
func (s Shares) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Shares) Less(i, j int) bool {
	return s[i].Counter().Cmp(s[j].Counter()) == -1
}

type Claim struct {
	shares Shares
}

func NewClaim() *Claim {
	return &Claim{Shares{}}
}

func (c *Claim) AddShare(s smartpool.Share) {
	c.shares = append(c.shares[:], s)
}

func (c *Claim) NumShares() *big.Int {
	return big.NewInt(int64(c.shares.Len()))
}

func (c *Claim) Difficulty() *big.Int {
	var m *big.Int
	if len(c.shares) > 0 {
		m = c.shares[0].Difficulty()
		for _, s := range c.shares {
			if m.Cmp(s.Difficulty()) >= 0 {
				m = s.Difficulty()
			}
		}
	}
	return m
}

func (c Claim) Min() *big.Int {
	var m *big.Int
	if len(c.shares) > 0 {
		m = c.shares[0].Counter()
		for _, s := range c.shares {
			if m.Cmp(s.Counter()) >= 0 {
				m = s.Counter()
			}
		}
	}
	return m
}

func (c Claim) Max() *big.Int {
	var m *big.Int
	if len(c.shares) > 0 {
		m = c.shares[0].Counter()
		for _, s := range c.shares {
			if m.Cmp(s.Counter()) <= 0 {
				m = s.Counter()
			}
		}
	}
	return m
}

func (c Claim) AugMerkle() []byte {
	return []byte{}
}
