package protocol

import (
	"../"
	"math/big"
)

type InMemClaimRepo struct {
	claims       map[int]*Claim
	cClaimNumber uint64
}

func NewInMemClaimRepo() *InMemClaimRepo {
	return &InMemClaimRepo{
		map[int]*Claim{0: NewClaim()},
		0,
	}
}

func (cr *InMemClaimRepo) AddShare(s smartpool.Share) {
	cr.claims[int(cr.cClaimNumber)].AddShare(s)
}

func (cr *InMemClaimRepo) GetClaim(number uint64) *Claim {
	return cr.claims[int(number)]
}

// TODO: This needs lock to prevent concurrent writes
func (cr *InMemClaimRepo) GetCurrentClaim() *Claim {
	c := cr.GetClaim(cr.cClaimNumber)
	if c.NumShares().Cmp(big.NewInt(0)) == 0 {
		return nil
	}
	cr.cClaimNumber++
	cr.claims[int(cr.cClaimNumber)] = NewClaim()
	return c
}
