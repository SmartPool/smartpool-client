package protocol

import (
	"../"
)

type testNetworkClient struct {
}

func (n *testNetworkClient) GetWork() smartpool.Work {
	return &testWork{}
}

func (n *testNetworkClient) SubmitSolution(s smartpool.Solution) bool {
	return true
}
