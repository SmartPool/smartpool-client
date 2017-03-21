// Package protocol implements smartpool secured protocol between client side
// and contract side. It works on high abstraction level of interfaces and
// types of smartpool package.
package protocol

import (
	"errors"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
	"sync"
	"time"
)

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
	PoolMonitor    smartpool.PoolMonitor
	ShareReceiver  smartpool.ShareReceiver
	NetworkClient  smartpool.NetworkClient
	Contract       smartpool.Contract
	ClaimRepo      ClaimRepo
	LatestCounter  *big.Int
	MinerAddress   common.Address
	SubmitInterval time.Duration
	ShareThreshold int
	HotStop        bool
	loopStarted    bool
	ticker         <-chan time.Time
	counterMu      sync.RWMutex
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
func (sp *SmartPool) GetWork() smartpool.Work {
	return sp.NetworkClient.GetWork()
}

// AcceptSolution accepts solutions from miners and construct corresponding
// shares to add into current active claim. It returns true when the share is
// successfully added, false otherwise.
// A share can only be added when it's counter is greater than the maximum
// counter of the last verified claim
func (sp *SmartPool) AcceptSolution(s smartpool.Solution) bool {
	share := sp.ShareReceiver.AcceptSolution(s)
	sp.counterMu.RLock()
	defer sp.counterMu.RUnlock()
	if share.Counter().Cmp(sp.LatestCounter) <= 0 {
		smartpool.Output.Printf("Share's counter (0x%s) is lower than last claim max counter (0x%s)\n", share.Counter().Text(16), sp.LatestCounter.Text(16))
	}
	if share == nil || share.Counter().Cmp(sp.LatestCounter) <= 0 {
		smartpool.Output.Printf("Share is discarded.\n")
		return false
	}
	sp.ClaimRepo.AddShare(share)
	smartpool.Output.Printf(".")
	return true
}

// GetCurrentClaim returns new claim containing unsubmitted shares. If there
// is no new shares, it returns nil.
func (sp *SmartPool) GetCurrentClaim(threshold int) smartpool.Claim {
	return sp.ClaimRepo.GetCurrentClaim(threshold)
}

func (sp *SmartPool) GetVerificationIndex(claim smartpool.Claim) *big.Int {
	return sp.Contract.GetShareIndex(claim)
}

func (sp *SmartPool) SealClaim() smartpool.Claim {
	sp.counterMu.Lock()
	defer sp.counterMu.Unlock()
	claim := sp.GetCurrentClaim(sp.ShareThreshold)
	if claim != nil {
		sp.LatestCounter = claim.Max()
		smartpool.Output.Printf("Set Latest Counter to 0x%s.\n", sp.LatestCounter.Text(16))
	}
	return claim
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
	if sp.PoolMonitor.RequireClientUpdate() {
		return false, errors.New("client update required")
	}
	if sp.PoolMonitor.RequireContractUpdate() {
		return false, errors.New("contract update required")
	}
	smartpool.Output.Printf("Submitting the claim with %d shares.\n", claim.NumShares().Int64())
	subErr := sp.Contract.SubmitClaim(claim)
	if subErr != nil {
		smartpool.Output.Printf("Got error submitting claim to contract: %s\n", subErr)
		return false, subErr
	}
	smartpool.Output.Printf("The claim is successfully submitted.\n")
	smartpool.Output.Printf("Waiting for verification index...")
	index := sp.GetVerificationIndex(claim)
	smartpool.Output.Printf("Verification index of %d has been requested. Submitting claim verification.\n", index.Int64())
	if sp.PoolMonitor.RequireClientUpdate() {
		return false, errors.New("client update required")
	}
	if sp.PoolMonitor.RequireContractUpdate() {
		return false, errors.New("contract update required")
	}
	verErr := sp.Contract.VerifyClaim(index, claim)
	if verErr != nil {
		smartpool.Output.Printf("%s\n", verErr)
		return false, verErr
	}
	smartpool.Output.Printf("Claim is successfully verified.\n")
	return true, nil
}

func (sp *SmartPool) actOnTick() {
	defer func() {
		if r := recover(); r != nil {
			smartpool.Output.Printf("Recovered in actOnTick: %v\n", r)
		}
	}()
	var err error
	for _ = range sp.ticker {
		_, err = sp.Submit()
		if err != nil && err.Error() == "client update required" {
			smartpool.Output.Printf("Your SmartPool client is too old. Please update to new version by going to https://github.com/SmartPool/smartpool-client.\n")
		}
		if err != nil && err.Error() == "contract update required" {
			smartpool.Output.Printf(
				"We deployed new contract at %s. Please restart SmartPool client with --spcontract %s.\n",
				sp.PoolMonitor.ContractAddress().Hex(),
				sp.PoolMonitor.ContractAddress().Hex())
		}
		if err != nil && sp.HotStop {
			smartpool.Output.Printf("SmartPool stopped. If you want SmartPool to keep running, please use \"--no-hot-stop\" to disable Hot Stop mode.\n")
			break
		}
	}
	os.Exit(1)
}

// Run can be called at most once to start a loop to submit and verify claims
// after an interval.
// If the loop has not been started, it starts the loop and return true, it
// return false otherwise.
// TODO: we need to have some lock here in case of concurrent invokes
func (sp *SmartPool) Run() bool {
	if sp.Register(sp.MinerAddress) {
		if sp.loopStarted {
			smartpool.Output.Printf("Warning: calling Run() multiple times\n")
			return false
		}
		sp.ticker = time.Tick(sp.SubmitInterval)
		go sp.actOnTick()
		sp.loopStarted = true
		smartpool.Output.Printf("Share collector is running...\n")
		return true
	} else {
		return false
	}
}

func NewSmartPool(
	pm smartpool.PoolMonitor,
	sr smartpool.ShareReceiver, nc smartpool.NetworkClient,
	cr ClaimRepo, co smartpool.Contract, ma common.Address,
	interval time.Duration, threshold int, hotStop bool) *SmartPool {
	return &SmartPool{
		PoolMonitor:    pm,
		ShareReceiver:  sr,
		NetworkClient:  nc,
		ClaimRepo:      cr,
		Contract:       co,
		MinerAddress:   ma,
		SubmitInterval: interval,
		ShareThreshold: threshold,
		HotStop:        hotStop,
		loopStarted:    false,
		// TODO: should be persist between startups instead of having 0 hardcoded
		LatestCounter: big.NewInt(0),
		counterMu:     sync.RWMutex{},
	}
}
