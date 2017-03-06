package protocol

import (
	"../"
)

type testShareReceiver struct {
}

func (spc *testShareReceiver) AcceptSolution(s smartpool.Solution) smartpool.Share {
	return &testShare{}
}
