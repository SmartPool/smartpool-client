package protocol

import (
	"github.com/SmartPool/smartpool-client"
)

type testShareReceiver struct {
}

func (spc *testShareReceiver) AcceptSolution(s smartpool.Solution) smartpool.Share {
	sol := s.(*testSolution)
	return &testShare{c: sol.Counter}
}
