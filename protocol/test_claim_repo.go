package protocol

import (
	"github.com/SmartPool/smartpool-client"
)

type testClaimRepo struct {
	c  []smartpool.Share
	oc []smartpool.Claim
}

func newClaimRepo() *testClaimRepo {
	return &testClaimRepo{[]smartpool.Share{}, []smartpool.Claim{}}
}

func (cr *testClaimRepo) GetCurrentClaim(threshold int) smartpool.Claim {
	if len(cr.c) < threshold {
		return nil
	}
	claim := &testClaim{cr.c}
	cr.c = []smartpool.Share{}
	cr.oc = []smartpool.Claim{claim}
	return claim
}

func (cr *testClaimRepo) AddShare(s smartpool.Share) error {
	cr.c = append(cr.c, s)
	return nil
}

func (cr *testClaimRepo) NoActiveShares() uint64 {
	return 0
}

func (cr *testClaimRepo) Persist(storage smartpool.PersistentStorage) error {
	return nil
}

func (cr *testClaimRepo) PutOpenClaim(claim smartpool.Claim) {
}

func (cr *testClaimRepo) GetOpenClaim(claimIndex int) smartpool.Claim {
	return cr.oc[0]
}

func (cr *testClaimRepo) ResetOpenClaims() {
}

func (cr *testClaimRepo) RemoveOpenClaim(claim smartpool.Claim) {
}

func (cr *testClaimRepo) NumOpenClaims() uint64 {
	return 0
}

func (cr *testClaimRepo) SealClaimBatch() {
}
