package protocol

import (
	"math/big"
)

type testSolution struct {
	Counter *big.Int
}

func (s *testSolution) WorkID() string {
	return "work"
}
