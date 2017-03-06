package protocol

import (
	"../"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type testContract struct {
	Registered   bool
	Registerable bool
}

func newTestContract() *testContract {
	return &testContract{false, false}
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
	return nil
}
func (c *testContract) VerifyClaim(shareIndex *big.Int, share smartpool.Share) error {
	return nil
}
