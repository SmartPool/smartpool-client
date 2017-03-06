package protocol

import (
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func newTestSmartPool() *SmartPool {
	return &SmartPool{
		ShareReceiver: &testShareReceiver{},
		NetworkClient: &testNetworkClient{},
		Output:        &testUserOutput{},
	}
}

func TestSmartPoolRegisterMinerAfterRegister(t *testing.T) {
	sp := newTestSmartPool()
	testContract := newTestContract()
	testContract.Registered = true
	sp.Contract = testContract
	if !sp.Register(common.Address{}) {
		t.Fail()
	}
}

func TestSmartPoolRegisterMinerWhenUnableToRegister(t *testing.T) {
	sp := newTestSmartPool()
	testContract := newTestContract()
	testContract.Registerable = false
	sp.Contract = testContract
	if sp.Register(common.Address{}) {
		t.Fail()
	}
}

func TestSmartPoolRegisterMinerWhenAbleToRegister(t *testing.T) {
	sp := newTestSmartPool()
	testContract := newTestContract()
	testContract.Registerable = true
	sp.Contract = testContract
	if !sp.Register(common.Address{}) {
		t.Fail()
	}
}

func TestSmartPoolReturnAWorkToMiner(t *testing.T) {
	sp := newTestSmartPool()
	sp.GetWork()
}

func TestSmartPoolAcceptSolution(t *testing.T) {
	sp := newTestSmartPool()
	sp.AcceptSolution(&testSolution{})
}

//
// func TestSmartPoolConstructsAShare(t *testing.T) {
// 	sp := newTestSmartPool()
// }
//
