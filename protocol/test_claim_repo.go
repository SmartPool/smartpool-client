package protocol

import (
	"../"
)

type testClaimRepo struct {
	c []smartpool.Share
}

func newClaimRepo() *testClaimRepo {
	return &testClaimRepo{[]smartpool.Share{}}
}

func (cr *testClaimRepo) GetCurrentClaim() smartpool.Claim {
	if len(cr.c) == 0 {
		return nil
	}
	claim := &testClaim{cr.c}
	cr.c = []smartpool.Share{}
	return claim
}

func (cr *testClaimRepo) AddShare(s smartpool.Share) {
	cr.c = append(cr.c, s)
}
