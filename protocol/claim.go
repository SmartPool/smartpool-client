package protocol

import (
	"../"
	"../mtree"
	"math/big"
	"sort"
)

type Shares []smartpool.Share

func (s Shares) Len() int      { return len(s) }
func (s Shares) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Shares) Less(i, j int) bool {
	return s[i].Counter().Cmp(s[j].Counter()) == -1
}

type Claim struct {
	shares     Shares
	shareIndex *big.Int
	amt        *mtree.AugTree
}

func NewClaim() *Claim {
	return &Claim{Shares{}, nil, nil}
}

func (c *Claim) AddShare(s smartpool.Share) {
	c.shares = append(c.shares[:], s)
}

func (c *Claim) GetShare(index int) smartpool.Share {
	return c.shares[index]
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

func (c *Claim) Min() *big.Int {
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

func (c *Claim) Max() *big.Int {
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

func (c *Claim) AugMerkle() smartpool.SPHash {
	sort.Sort(c.shares)
	amt := mtree.NewAugTree()
	for i, s := range c.shares {
		amt.Insert(s, uint32(i))
	}
	amt.Finalize()
	return amt.RootHash()
}

func (c *Claim) SetEvidence(shareIndex *big.Int) {
	c.shareIndex = shareIndex
}

func (c *Claim) CounterBranch() []*big.Int {
	if c.amt == nil {
		c.buildAugTree()
	}
	return c.amt.CounterBranchArray()
}

func (c *Claim) HashBranch() []*big.Int {
	if c.amt == nil {
		c.buildAugTree()
	}
	return c.amt.HashBranchArray()
}

func (c *Claim) buildAugTree() {
	sort.Sort(c.shares)
	c.amt = mtree.NewAugTree()
	c.amt.RegisterIndex(uint32(c.shareIndex.Int64()))
	for i, s := range c.shares {
		c.amt.Insert(s, uint32(i))
	}
	c.amt.Finalize()
}
