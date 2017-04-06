package protocol

import (
	"github.com/SmartPool/smartpool-client"
)

type testClaimRepo struct {
	c []smartpool.Share
}

func newClaimRepo() *testClaimRepo {
	return &testClaimRepo{[]smartpool.Share{}}
}

func (cr *testClaimRepo) GetCurrentClaim(threshold int) smartpool.Claim {
	if len(cr.c) < threshold {
		return nil
	}
	claim := &testClaim{cr.c}
	cr.c = []smartpool.Share{}
	return claim
}

func (cr *testClaimRepo) AddShare(s smartpool.Share) error {
	cr.c = append(cr.c, s)
	return nil
}
