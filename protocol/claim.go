package protocol

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/mtree"
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
	pamt       *mtree.AugTree
}

func NewClaim() *Claim {
	return &Claim{Shares{}, nil, nil, nil}
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
		m = c.shares[0].ShareDifficulty()
		for _, s := range c.shares {
			if m.Cmp(s.ShareDifficulty()) >= 0 {
				m = s.ShareDifficulty()
			}
		}
	}
	return m
}

func (c *Claim) Min() *big.Int {
	if c.amt == nil {
		c.buildAugTree()
	}
	return c.amt.RootMin()
}

func (c *Claim) Max() *big.Int {
	if c.amt == nil {
		c.buildAugTree()
	}
	return c.amt.RootMax()
}

func (c *Claim) AugMerkle() smartpool.SPHash {
	if c.amt == nil {
		c.buildAugTree()
	}
	return c.amt.RootHash()
}

func (c *Claim) SetEvidence(shareIndex *big.Int) {
	c.shareIndex = shareIndex
}

func (c *Claim) CounterBranch() []*big.Int {
	if c.pamt == nil {
		c.buildProofAugTree()
	}
	return c.pamt.CounterBranchArray()
}

func (c *Claim) HashBranch() []*big.Int {
	if c.pamt == nil {
		c.buildProofAugTree()
	}
	return c.pamt.HashBranchArray()
}

func (c *Claim) buildProofAugTree() {
	sort.Sort(c.shares)
	c.pamt = mtree.NewAugTree()
	c.pamt.RegisterIndex(uint32(c.shareIndex.Int64()))
	for i, s := range c.shares {
		c.pamt.Insert(s, uint32(i))
	}
	c.pamt.Finalize()
}

func (c *Claim) buildAugTree() {
	sort.Sort(c.shares)
	c.amt = mtree.NewAugTree()
	for i, s := range c.shares {
		c.amt.Insert(s, uint32(i))
	}
	c.amt.Finalize()
}
