package protocol

import (
	"math/big"
)

type testShare struct {
	c *big.Int
}

func (s *testShare) Counter() *big.Int {
	return s.c
}

func (s *testShare) Hash() []byte {
	return []byte{0}
}

func (s *testShare) Difficulty() *big.Int {
	return big.NewInt(100000)
}
