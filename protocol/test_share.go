package protocol

import (
	"math/big"
)

type testShare struct {
}

func (s *testShare) Counter() *big.Int {
	return big.NewInt(0)
}

func (s *testShare) Hash() []byte {
	return []byte{0}
}
