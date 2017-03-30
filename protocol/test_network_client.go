package protocol

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
)

type testNetworkClient struct {
	NotReadyToMine bool
}

func (n *testNetworkClient) GetWork() smartpool.Work {
	return &testWork{}
}

func (n *testNetworkClient) SubmitSolution(s smartpool.Solution) bool {
	return true
}

func (n *testNetworkClient) Configure(etherbase common.Address, extradata string) error {
	return nil
}

func (n *testNetworkClient) ReadyToMine() bool {
	return !n.NotReadyToMine
}
