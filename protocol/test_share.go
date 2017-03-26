package protocol

import (
	"github.com/SmartPool/smartpool-client"
	"math/big"
)

type testShare struct {
	c *big.Int
	d *big.Int
	h byte
}

func (s *testShare) Counter() *big.Int {
	return s.c
}

func (s *testShare) Hash() smartpool.SPHash {
	h := smartpool.SPHash{}
	copy(h[:], []byte{s.h})
	return h
}

func (s *testShare) ShareDifficulty() *big.Int {
	return s.d
}

func (s *testShare) FullSolution() bool {
	return true
}
