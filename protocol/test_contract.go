package protocol

import (
	"../"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type testContract struct {
	Registered   bool
	Registerable bool
	SubmitFailed bool
	VerifyFailed bool
	claim        *testClaim
}

func newTestContract() *testContract {
	return &testContract{false, false, false, false, nil}
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
	return nil
}
func (c *testContract) VerifyClaim(shareIndex *big.Int, claim smartpool.Claim) error {
	if c.VerifyFailed {
		return errors.New("fail")
	}
	return nil
}
func (c *testContract) GetLastSubmittedClaim() *testClaim {
	return c.claim
}
