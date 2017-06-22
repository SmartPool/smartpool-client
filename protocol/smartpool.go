// Package protocol implements smartpool secured protocol between client side
// and contract side. It works on high abstraction level of interfaces and
// types of smartpool package.
package protocol

import (
	"errors"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
)

const COUNTER_FILE string = "counter"

// SmartPool represent smartpool protocol which interacts smartpool high level
// interfaces and types together to do following procedures:
// 1. Register the miner if needed
// 2. Give miner works
// 3. Accept solutions from miners and construct corresponding shares to
//    add into current active claim. It returns true when the share is
//    successfully added, false otherwise.
//    A share can only be added when it's counter is greater than the maximum
//    counter of the last verified claim
// 4. Package shares into a claim after interval of an amount of time.
// 5. Then Submit the claim to the contract
// 6. Then If successful, submit the claim proof to the contract.
type SmartPool struct {
	PoolMonitor       smartpool.PoolMonitor
	ShareReceiver     smartpool.ShareReceiver
	NetworkClient     smartpool.NetworkClient
	Contract          smartpool.Contract
	StatRecorder      smartpool.StatRecorder
	ClaimRepo         ClaimRepo
	Storage           smartpool.PersistentStorage
	LatestCounter     *big.Int
	ContractAddress   common.Address
	MinerAddress      common.Address
	ExtraData         string
	SubmitInterval    time.Duration
	ShareThreshold    int
	ClaimThreshold    int
	HotStop           bool
	loopStarted       bool
	ticker            <-chan time.Time
	counterMu         sync.RWMutex
	runMu             sync.Mutex
	SubmitterStopped  chan bool
	stopSubmitterChan chan bool
	signal            chan os.Signal
	Input             smartpool.UserInput
}

// Register registers miner address to the contract.
// It returns false if the miner address couldn't be able to register to the
// pool even though it didn't register before.
// It returns true otherwise, in this case, Register does nothing if the
// address registered before or registers the address if it didn't.
func (sp *SmartPool) Register(addr common.Address) bool {
	if sp.Contract.IsRegistered() {
		smartpool.Output.Printf("The address is already registered to the pool. Good to go.\n")
		return true
	}
	if !sp.Contract.CanRegister() {
		smartpool.Output.Printf("Your etherbase address couldn't register to the pool. You need to try another address.\n")
		return false
	}
	smartpool.Output.Printf("Registering to the pool. Please wait...")
	err := sp.Contract.Register(addr)
	if err != nil {
		smartpool.Output.Printf("Unable to register to the pool: %s\n", err)
		return false
	}
	if !sp.Contract.IsRegistered() {
		smartpool.Output.Printf("You are not accepted by the pool yet. Please wait about 30s and try again.\n")
		return false
	}
	smartpool.Output.Printf("Done.\n")
	return true
}

// GetWork returns miner work
func (sp *SmartPool) GetWork(rig smartpool.Rig) smartpool.Work {
	return sp.NetworkClient.GetWork()
}

func (sp *SmartPool) SubmitHashrate(rig smartpool.Rig, hashrate hexutil.Uint64, id common.Hash) bool {
	sp.StatRecorder.RecordHashrate(hashrate, id, rig)
	return sp.NetworkClient.SubmitHashrate(hashrate, id)
}

// AcceptSolution accepts solutions from miners and construct corresponding
// shares to add into current active claim. It returns true when the share is
// successfully added, false otherwise.
// A share can only be added when it's counter is greater than the maximum
// counter of the last verified claim
func (sp *SmartPool) AcceptSolution(rig smartpool.Rig, s smartpool.Solution) bool {
	share := sp.ShareReceiver.AcceptSolution(s)
	sp.counterMu.RLock()
	defer sp.counterMu.RUnlock()
	if share != nil && share.FullSolution() {
		smartpool.Output.Printf("-->Yay! We found potential block!<--\n")
		sp.NetworkClient.SubmitSolution(s)
	}
	var success bool
	if share == nil || share.Counter().Cmp(sp.LatestCounter) <= 0 {
		smartpool.Output.Printf("Share is discarded.\n")
		if share != nil && share.Counter().Cmp(sp.LatestCounter) <= 0 {
			smartpool.Output.Printf("Share's counter (0x%s) is lower than last claim max counter (0x%s)\n", share.Counter().Text(16), sp.LatestCounter.Text(16))
		}
		success = false
	} else {
		err := sp.ClaimRepo.AddShare(share)
		if err != nil {
			smartpool.Output.Printf("Discarded because of %s.\n", err.Error())
			success = false
		} else {
			fmt.Print(".")
			success = true
		}
	}

	go func() {
		sp.StatRecorder.RecordShare("submitted", share, rig)
		if success {
			if share.FullSolution() {
				sp.StatRecorder.RecordShare("fullsolution", share, rig)
			} else {
				sp.StatRecorder.RecordShare("accepted", share, rig)
			}
		} else {
			sp.StatRecorder.RecordShare("rejected", share, rig)
		}
	}()

	return success
}

// GetCurrentClaim returns new claim containing unsubmitted shares. If there
// is no new shares, it returns nil.
func (sp *SmartPool) GetCurrentClaim(threshold int) smartpool.Claim {
	return sp.ClaimRepo.GetCurrentClaim(threshold)
}

func (sp *SmartPool) GetVerificationIndex(claim smartpool.Claim) (*big.Int, *big.Int) {
	var claimIndex, shareIndex *big.Int
	var err error
	for {
		claimIndex, shareIndex, err = sp.Contract.GetShareIndex(claim)
		if err != nil {
			smartpool.Output.Printf("Got error(%s) while trying to get verification index. Retry in 10s...\n", err.Error())
			time.Sleep(10 * time.Second)
		} else {
			return claimIndex, shareIndex
		}
	}
}

func (sp *SmartPool) SealClaim() smartpool.Claim {
	sp.counterMu.Lock()
	defer sp.counterMu.Unlock()
	claim := sp.GetCurrentClaim(sp.ShareThreshold)
	if claim != nil {
		sp.LatestCounter = claim.Max()
		smartpool.Output.Printf("Set Latest Counter to 0x%s.\n", sp.LatestCounter.Text(16))
		smartpool.Output.Printf("Persisting Latest Counter to storage...")
		persistLatestCounter(sp.Storage, sp.LatestCounter)
		smartpool.Output.Printf("Done.\n")
	}
	return claim
}

func persistLatestCounter(ps smartpool.PersistentStorage, counter *big.Int) error {
	return ps.Persist(counter, COUNTER_FILE)
}

func (sp *SmartPool) consistencyCheck() error {
	waited := 0
	for {
		numOpenClaimsContract, err := sp.Contract.NumOpenClaims()
		if err != nil {
			return err
		}
		if numOpenClaimsContract.Uint64() != sp.ClaimRepo.NumOpenClaims() {
			if waited >= 140 {
				smartpool.Output.Printf("Unrecoverable inconsistent state between client and contract. Resetting both sides...")
				sp.ClaimRepo.ResetOpenClaims()
				err := sp.Contract.ResetOpenClaims()
				if err != nil {
					return err
				} else {
					smartpool.Output.Printf("Done.\n")
				}
				break
			} else {
				smartpool.Output.Printf(
					"Inconsistent open claim list between client(%d claims) and contract(%d claims). Recheck in 14s...\n",
					sp.ClaimRepo.NumOpenClaims(),
					numOpenClaimsContract.Uint64(),
				)
				time.Sleep(14 * time.Second)
				waited += 14
			}
		} else {
			break
		}
	}
	return nil
}

// Submit does all the protocol that communicates with the contract to submit
// the claim then verify it.
// It returns true when the claim is fully verified and accepted by the
// contract. It returns false otherwise.
func (sp *SmartPool) Submit() (bool, error) {
	claim := sp.SealClaim()
	if claim == nil {
		return false, nil
	}
	if err := sp.consistencyCheck(); err != nil {
		return false, err
	}
	smartpool.Output.Printf("Submitting the claim with %d shares.\n", claim.NumShares().Int64())
	var lastClaim bool
	if int(sp.ClaimRepo.NumOpenClaims()+1) >= sp.ClaimThreshold {
		lastClaim = true
	} else {
		lastClaim = false
	}
	sp.ClaimRepo.PutOpenClaim(claim)
	smartpool.Output.Printf("The claim is successfully put into open claims queue.\n")
	subErr := sp.Contract.SubmitClaim(claim, lastClaim)
	if subErr != nil {
		smartpool.Output.Printf("Got error submitting claim to contract: %s\n", subErr)
		sp.ClaimRepo.RemoveOpenClaim(claim)
		sp.StatRecorder.RecordClaim("error", claim)
		return false, subErr
	}
	sp.StatRecorder.RecordClaim("submitted", claim)
	smartpool.Output.Printf("The claim is successfully submitted.\n")
	if lastClaim {
		sp.ClaimRepo.SealClaimBatch()
		smartpool.Output.Printf("Waiting for verification index...")
		claimIndex, shareIndex := sp.GetVerificationIndex(claim)
		smartpool.Output.Printf("Verification for share index(%d) in claim index(%d) has been requested.\n", shareIndex.Int64(), claimIndex.Int64())
		claim = sp.ClaimRepo.GetOpenClaim(int(claimIndex.Int64()))
		if claim == nil {
			smartpool.Output.Printf("Got nil claim for share index(%d) in claim index(%d). This is a bug. Please report it to SmartPool Team.\n", shareIndex.Int64(), claimIndex.Int64())
			return false, errors.New("Nil claim. Incorrect verification indexes")
		}
		smartpool.Output.Printf("Submitting claim verification...\n")
		verErr := sp.Contract.VerifyClaim(claimIndex, shareIndex, claim)
		if verErr != nil {
			smartpool.Output.Printf("%s\n", verErr)
			sp.StatRecorder.RecordClaim("rejected", claim)
			return false, verErr
		}
		smartpool.Output.Printf("Claim is successfully verified.\n")
		sp.StatRecorder.RecordClaim("accepted", claim)
	}
	return true, nil
}

func (sp *SmartPool) stopSubmitter() {
	sp.stopSubmitterChan <- true
}

func (sp *SmartPool) monitor() {
	for {
		if sp.PoolMonitor.RequireContractUpdate() {
			smartpool.Output.Printf(
				"We deployed new contract at %s. Please restart SmartPool client with --spcontract %s.\n",
				sp.PoolMonitor.ContractAddress().Hex(),
				sp.PoolMonitor.ContractAddress().Hex())
			if sp.HotStop {
				sp.stopSubmitter()
				return
			}
		}
		if sp.PoolMonitor.RequireClientUpdate() {
			smartpool.Output.Printf("Your SmartPool client is too old. Please update to new version by going to https://github.com/SmartPool/smartpool-client.\n")
			if sp.HotStop {
				sp.stopSubmitter()
				return
			}
		}
		time.Sleep(1 * time.Minute)
	}
}

func (sp *SmartPool) SubmitterRunning() bool {
	return sp.loopStarted
}

func (sp *SmartPool) shouldStop(err error) bool {
	if err == nil {
		return false
	}
	if err.Error() == "timeout error" {
		smartpool.Output.Printf("The tx might not be verified. Current claim is dropped. Continue with next claim.\n")
		return false
	} else if sp.HotStop {
		return true
	} else {
		return false
	}
}

func (sp *SmartPool) runPersister() {
	for {
		sp.persist()
		time.Sleep(time.Minute)
	}
}

func (sp *SmartPool) persist() {
	sp.ClaimRepo.Persist(sp.Storage)
	sp.StatRecorder.Persist(sp.Storage)
	sp.ShareReceiver.Persist(sp.Storage)
}

func (sp *SmartPool) actOnTick() {
	defer func() {
		if r := recover(); r != nil {
			smartpool.Output.Printf("Recovered in actOnTick: %v\n", r)
			debug.PrintStack()
		}
	}()
	var err error
Loop:
	for {
		select {
		case <-sp.ticker:
			_, err = sp.Submit()
			if sp.shouldStop(err) {
				smartpool.Output.Printf("SmartPool stopped. If you want SmartPool to keep running, please use \"--no-hot-stop\" to disable Hot Stop mode.\n")
				break Loop
			}
		case <-sp.stopSubmitterChan:
			break Loop
		}
	}
	sp.Exit()
}

func (sp *SmartPool) Exit() {
	smartpool.Output.Printf("Persisting current state to disk...\n")
	sp.persist()
	smartpool.Output.Printf("Gracefully stopped SmartPool.\n")
	sp.SubmitterStopped <- true
	smartpool.Output.Printf("Close log file.\n")
	smartpool.Output.Close()
}

// Run can be called at most once to start a loop to submit and verify claims
// after an interval.
// If the loop has not been started, it starts the loop and return true, it
// return false otherwise.
func (sp *SmartPool) Run() bool {
	sp.runMu.Lock()
	defer sp.runMu.Unlock()
	if sp.Register(sp.MinerAddress) {
		if sp.loopStarted {
			smartpool.Output.Printf("Warning: calling Run() multiple times\n")
			return false
		}
		err := sp.NetworkClient.Configure(
			sp.ContractAddress,
			sp.ExtraData,
		)
		if err != nil {
			return false
		}
		for {
			if sp.NetworkClient.ReadyToMine() {
				smartpool.Output.Printf("The network is ready for mining.\n")
				sp.ticker = time.Tick(sp.SubmitInterval)
				go sp.monitor()
				go sp.actOnTick()
				smartpool.Output.Printf("Share collector is running...\n")
				go sp.runPersister()
				smartpool.Output.Printf("Share persister and stat persister are running...\n")
				go sp.handleSignal()
				sp.loopStarted = true
				break
			}
			smartpool.Output.Printf("The network is not ready for mining yet. Retry in 10s...\n")
			time.Sleep(10 * time.Second)
		}
		return true
	} else {
		return false
	}
}

func (sp *SmartPool) handleSignal() {
	<-sp.signal
	smartpool.Output.Printf("Got shutdown signal:\n")
	sp.Exit()
}

func loadLatestCounter(ps smartpool.PersistentStorage) (*big.Int, error) {
	counter := big.NewInt(0)
	loadedCounter, err := ps.Load(counter, COUNTER_FILE)
	counter = loadedCounter.(*big.Int)
	return counter, err
}

func NewSmartPool(
	pm smartpool.PoolMonitor, sr smartpool.ShareReceiver,
	nc smartpool.NetworkClient, cr ClaimRepo, ps smartpool.PersistentStorage,
	co smartpool.Contract, stat smartpool.StatRecorder, ca common.Address,
	ma common.Address, ed string, interval time.Duration, shareThreshold int,
	claimThreshold int, hotStop bool, input smartpool.UserInput) *SmartPool {
	counter, err := loadLatestCounter(ps)
	if err != nil {
		smartpool.Output.Printf("Couldn't load counter from storage. Initialize it to 0.\n")
		counter = big.NewInt(0)
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return &SmartPool{
		PoolMonitor:       pm,
		ShareReceiver:     sr,
		NetworkClient:     nc,
		ClaimRepo:         cr,
		Storage:           ps,
		Contract:          co,
		StatRecorder:      stat,
		ContractAddress:   ca,
		MinerAddress:      ma,
		ExtraData:         ed,
		SubmitInterval:    interval,
		ShareThreshold:    shareThreshold,
		ClaimThreshold:    claimThreshold,
		HotStop:           hotStop,
		loopStarted:       false,
		LatestCounter:     counter,
		counterMu:         sync.RWMutex{},
		runMu:             sync.Mutex{},
		SubmitterStopped:  make(chan bool, 1),
		stopSubmitterChan: make(chan bool, 1),
		Input:             input,
		signal:            sig,
	}
}
