package protocol

import (
	"../"
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

func (s *testShare) Difficulty() *big.Int {
	return s.d
}
