package protocol

import (
	"../"
)

type testWork struct {
}

func (w *testWork) ID() string {
	return "work"
}

func (w *testWork) AcceptSolution(sol smartpool.Solution) smartpool.Share {
	return &testShare{}
}
