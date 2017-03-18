package protocol

import (
	"errors"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

type testContract struct {
	Registered          bool
	Registerable        bool
	SubmitFailed        bool
	VerifyFailed        bool
	SubmitTime          *time.Time
	IndexRequestedTime  *time.Time
	claim               *testClaim
	DelayedVerification bool
}

func newTestContract() *testContract {
	return &testContract{false, false, false, false, nil, nil, nil, false}
}

func (c *testContract) Version() string {
	return "1.0.0"
}
func (c *testContract) IsRegistered() bool {
	return c.Registered
}
func (c *testContract) CanRegister() bool {
	return c.Registerable
}
func (c *testContract) Register(paymentAddress common.Address) error {
	c.Registered = true
	return nil
}
func (c *testContract) SubmitClaim(claim smartpool.Claim) error {
	c.claim = claim.(*testClaim)
	if c.SubmitFailed {
		return errors.New("fail")
	}
	t := time.Now()
	c.SubmitTime = &t
	return nil
}
func (c *testContract) GetShareIndex(claim smartpool.Claim) *big.Int {
	t := time.Now()
	c.IndexRequestedTime = &t
	return big.NewInt(100)
}
func (c *testContract) VerifyClaim(shareIndex *big.Int, claim smartpool.Claim) error {
	if c.VerifyFailed {
		return errors.New("fail")
	}
	if c.DelayedVerification {
		time.Sleep(50 * time.Millisecond)
	}
	return nil
}
func (c *testContract) GetLastSubmittedClaim() *testClaim {
	return c.claim
}
