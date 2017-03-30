package protocol

import (
	"math/big"
	"sync"
	"testing"
)

func TestInMemClaimRepoGetSpecificClaim(t *testing.T) {
	cr := &InMemClaimRepo{
		map[int]*Claim{0: NewClaim()},
		0,
		sync.Mutex{},
	}
	claim := cr.GetClaim(0)
	if claim.NumShares().Int64() != 0 {
		t.Fail()
	}
}

func TestInMemClaimAddShare(t *testing.T) {
	cr := NewInMemClaimRepo()
	cr.AddShare(&testShare{
		c: big.NewInt(10),
		d: big.NewInt(10),
		h: 1,
	})
	claim := cr.GetClaim(0)
	if claim.NumShares().Int64() != 1 {
		t.Fail()
	}
}

func TestInMemClaimGetCurrentClaim(t *testing.T) {
	cr := NewInMemClaimRepo()
	cr.AddShare(&testShare{
		c: big.NewInt(10),
		d: big.NewInt(10),
		h: 1,
	})
	claim := cr.GetCurrentClaim(1)
	if claim.NumShares().Int64() != 1 {
		t.Fail()
	}
}

func TestInMemClaimGetCurrentClaimThenNewClaimIsReady(t *testing.T) {
	cr := NewInMemClaimRepo()
	cr.AddShare(&testShare{
		c: big.NewInt(10),
		d: big.NewInt(10),
		h: 1,
	})
	cr.GetCurrentClaim(1)
	cr.AddShare(&testShare{
		c: big.NewInt(10),
		d: big.NewInt(10),
		h: 1,
	})
	claim := cr.GetClaim(1)
	if claim.NumShares().Int64() != 1 {
		t.Fail()
	}
}
