package protocol

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
	"time"
)

func newTestSmartPool() *SmartPool {
	return NewSmartPool(
		&testPoolMonitor{},
		&testShareReceiver{}, &testNetworkClient{},
		&testClaimRepo{}, &testPersistentStorage{},
		&testContract{},
		common.HexToAddress("0x001aDBc838eDe392B5B054A47f8B8c28f2fA9F3F"),
		common.HexToAddress("0x001aDBc838eDe392B5B054A47f8B8c28f2fA9F3F"),
		"extradata", time.Minute,
		100, true,
	)
}

func TestSmartPoolRegisterMinerAfterRegister(t *testing.T) {
	sp := newTestSmartPool()
	testContract := sp.Contract.(*testContract)
	testContract.Registered = true
	if !sp.Register(common.Address{}) {
		t.Fail()
	}
}

func TestSmartPoolRegisterMinerWhenUnableToRegister(t *testing.T) {
	sp := newTestSmartPool()
	testContract := sp.Contract.(*testContract)
	testContract.Registerable = false
	if sp.Register(common.Address{}) {
		t.Fail()
	}
}

func TestSmartPoolRegisterMinerWhenAbleToRegister(t *testing.T) {
	sp := newTestSmartPool()
	testContract := sp.Contract.(*testContract)
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
	claim := sp.GetCurrentClaim(1)
	if claim != nil {
		t.Fail()
	}
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(9)})
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(8)})
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(10)})
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(5)})
	claim = sp.GetCurrentClaim(1)
	if claim.NumShares().Cmp(big.NewInt(3)) != 0 {
		t.Fail()
	}
}

func TestSmartPoolSubmitCorrectClaim(t *testing.T) {
	sp := newTestSmartPool()
	sp.ShareThreshold = 1
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

func TestSmartPoolReturnFalseIfNoClaim(t *testing.T) {
	sp := newTestSmartPool()
	if ok, _ := sp.Submit(); ok {
		t.Fail()
	}
}

func TestSmartPoolSuccessfullySubmitAndVerifyClaim(t *testing.T) {
	sp := newTestSmartPool()
	sp.ShareThreshold = 1
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(9)})
	if ok, _ := sp.Submit(); !ok {
		t.Fail()
	}
}

func TestSmartPoolGetCorrectShareIndex(t *testing.T) {
	sp := newTestSmartPool()
	sp.ShareThreshold = 1
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(9)})
	sp.Submit()
	c := sp.Contract.(*testContract)
	if c.IndexRequestedTime == nil {
		t.Fail()
	}
}

func TestSmartPoolGetCorrectShareIndexAfterSubmitClaim(t *testing.T) {
	sp := newTestSmartPool()
	sp.ShareThreshold = 1
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(9)})
	sp.Submit()
	c := sp.Contract.(*testContract)
	if (*c.SubmitTime).After(*c.IndexRequestedTime) {
		t.Fail()
	}
}

func TestSmartPoolSubmitReturnFalseWhenUnableToSubmit(t *testing.T) {
	sp := newTestSmartPool()
	c := sp.Contract.(*testContract)
	c.SubmitFailed = true
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(9)})
	if ok, _ := sp.Submit(); ok {
		t.Fail()
	}
}

func TestSmartPoolSubmitReturnFalseWhenUnableToVerify(t *testing.T) {
	sp := newTestSmartPool()
	c := sp.Contract.(*testContract)
	c.VerifyFailed = true
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(9)})
	if ok, _ := sp.Submit(); ok {
		t.Fail()
	}
}

func TestSmartPoolDoesntRunWhenMinerRegistered(t *testing.T) {
	sp := newTestSmartPool()
	if sp.Run() {
		t.Fail()
	}
}

func TestSmartPoolOnlySubmitPeriodly(t *testing.T) {
	sp := newTestSmartPool()
	ct := sp.Contract.(*testContract)
	ct.Registered = true
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(9)})
	c := sp.Contract.(*testContract)
	sp.SubmitInterval = 40 * time.Millisecond
	sp.ShareThreshold = 1
	sp.Run()
	if c.GetLastSubmittedClaim() != nil {
		t.Fail()
	}
	time.Sleep(60 * time.Millisecond)
	if c.GetLastSubmittedClaim() == nil {
		t.Fail()
	}
}

func TestSmartPoolOnlySubmitWhenMeetShareThreshold(t *testing.T) {
	sp := newTestSmartPool()
	sp.AcceptSolution(&testSolution{Counter: big.NewInt(9)})
	c := sp.Contract.(*testContract)
	sp.SubmitInterval = 40 * time.Millisecond
	sp.ShareThreshold = 3
	sp.Run()
	time.Sleep(60 * time.Millisecond)
	if c.GetLastSubmittedClaim() != nil {
		t.Fail()
	}
}

func TestSmartPoolOnlyRunAfterNetworkReady(t *testing.T) {
	sp := newTestSmartPool()
	testContract := sp.Contract.(*testContract)
	testContract.Registered = true
	nw := sp.NetworkClient.(*testNetworkClient)
	nw.NotReadyToMine = true
	ran := make(chan bool, 1)
	timeout := make(chan bool, 1)
	go func() {
		ran <- sp.Run()
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		timeout <- true
	}()
	select {
	case <-ran:
		t.Fail()
	case <-timeout:
		break
	}
}

func TestSmartPoolStopIfClientVersionChangedInHotStopMode(t *testing.T) {
	sp := newTestSmartPool()
	testContract := sp.Contract.(*testContract)
	testContract.Registered = true
	timeout := make(chan bool, 1)
	sp.PoolMonitor.(*testPoolMonitor).ClientUpdate = true
	sp.Run()
	go func() {
		time.Sleep(100 * time.Millisecond)
		timeout <- true
	}()
	select {
	case <-sp.SubmitterStopped:
		break
	case <-timeout:
		t.Fail()
	}
}

func TestSmartPoolDoesntStopIfHotStopModeIsDisabled(t *testing.T) {
	sp := newTestSmartPool()
	sp.HotStop = false
	testContract := sp.Contract.(*testContract)
	testContract.Registered = true
	timeout := make(chan bool, 1)
	sp.PoolMonitor.(*testPoolMonitor).ContractUpdate = true
	sp.Run()
	go func() {
		time.Sleep(100 * time.Millisecond)
		timeout <- true
	}()
	select {
	case <-sp.SubmitterStopped:
		t.Fail()
	case <-timeout:
		break
	}
}

func TestSmartPoolStopIfContractAddressChangedInHotStopMode(t *testing.T) {
	sp := newTestSmartPool()
	testContract := sp.Contract.(*testContract)
	testContract.Registered = true
	timeout := make(chan bool, 1)
	sp.PoolMonitor.(*testPoolMonitor).ContractUpdate = true
	sp.Run()
	go func() {
		time.Sleep(100 * time.Millisecond)
		timeout <- true
	}()
	select {
	case <-sp.SubmitterStopped:
		break
	case <-timeout:
		t.Fail()
	}
}
