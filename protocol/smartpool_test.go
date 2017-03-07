package protocol

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func newTestSmartPool() *SmartPool {
	return NewSmartPool(
		&testShareReceiver{},
		&testNetworkClient{},
		&testClaimRepo{},
		&testUserOutput{},
		&testContract{},
	)
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
	if !sp.AcceptSolution(&testSolution{Counter: big.NewInt(10)}) {
		t.Fail()
	}
}

func TestSmartPoolNotAcceptSolution(t *testing.T) {
	sp := newTestSmartPool()
	sp.LatestCounter = big.NewInt(10)
	if sp.AcceptSolution(&testSolution{Counter: big.NewInt(9)}) {
		t.Fail()
	}
}

func TestSmartPoolPackageAllCurrentShares(t *testing.T) {
	sp := newTestSmartPool()
	sp.LatestCounter = big.NewInt(5)
	claim := sp.GetCurrentClaim()
	if claim != nil {
		t.Fail()
	}
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(9)})
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(8)})
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(10)})
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(5)})
	claim = sp.GetCurrentClaim()
	if claim.NumShares().Cmp(big.NewInt(3)) != 0 {
		t.Fail()
	}
}

func TestSmartPoolSubmitCorrectClaim(t *testing.T) {
	sp := newTestSmartPool()
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(9)})
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(8)})
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(10)})
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(5)})
	sp.Submit()

	testContract := sp.Contract.(*testContract)
	claim := testContract.GetLastSubmittedClaim()
	if claim.NumShares().Cmp(big.NewInt(4)) != 0 {
		t.Fail()
	}
}

//
// func TestSmartPoolConstructsAShare(t *testing.T) {
// 	sp := newTestSmartPool()
// }
//
