package protocol

import (
	"github.com/SmartPool/smartpool-client"
	"sync"
)

// InMemClaimRepo implements ClaimRepo interface. It stores claims in one start
// up. However this claim repo doesn't persist verified claims and even the
// active claim. So if the client is shutdown, all past claims and current
// shares information will be lost.
// This shouldn't be used in production.
type InMemClaimRepo struct {
	claims       map[int]*Claim
	cClaimNumber uint64
	mu           sync.Mutex
}

func NewInMemClaimRepo() *InMemClaimRepo {
	return &InMemClaimRepo{
		map[int]*Claim{0: NewClaim()},
		0,
		sync.Mutex{},
	}
}

func (cr *InMemClaimRepo) Persist(storage smartpool.PersistentStorage) error {
	return nil
}

func (cr *InMemClaimRepo) NoActiveShares() uint64 {
	c := cr.GetClaim(cr.cClaimNumber)
	return c.NumShares().Uint64()
}

func (cr *InMemClaimRepo) AddShare(s smartpool.Share) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	cr.claims[int(cr.cClaimNumber)].AddShare(s)
	return nil
}

func (cr *InMemClaimRepo) GetClaim(number uint64) smartpool.Claim {
	return cr.claims[int(number)]
}

func (cr *InMemClaimRepo) GetCurrentClaim(threshold int) smartpool.Claim {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	c := cr.GetClaim(cr.cClaimNumber)
	if c.NumShares().Int64() < int64(threshold) {
		return nil
	}
	cr.cClaimNumber++
	cr.claims[int(cr.cClaimNumber)] = NewClaim()
	return c
}
