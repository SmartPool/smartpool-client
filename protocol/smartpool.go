// Package protocol implements smartpool secured protocol between client side
// and contract side. It works on high abstraction level of interfaces and
// types of smartpool package.
package protocol

import (
	"../"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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
	ShareReceiver smartpool.ShareReceiver
	NetworkClient smartpool.NetworkClient
	Contract      smartpool.Contract
	Output        smartpool.UserOutput
	ClaimRepo     ClaimRepo
	LatestCounter *big.Int
}

// Register registers miner address to the contract.
// It returns false if the miner address couldn't be able to register to the
// pool even though it didn't register before.
// It returns true otherwise, in this case, Register does nothing if the
// address registered before or registers the address if it didn't.
func (sp *SmartPool) Register(addr common.Address) bool {
	if sp.Contract.IsRegistered() {
		sp.Output.Printf("The address is already registered to the pool. Good to go.\n")
		return true
	}
	if !sp.Contract.CanRegister() {
		sp.Output.Printf("Your etherbase address couldn't register to the pool. You need to try another address.\n")
		return false
	}
	sp.Output.Printf("Registering to the pool. Please wait...")
	err := sp.Contract.Register(addr)
	if err != nil {
		sp.Output.Printf("Unable to register to the pool: %s\n", err)
		return false
	}
	if !sp.Contract.IsRegistered() {
		sp.Output.Printf("Unable to register to the pool. You might try again.")
		return false
	}
	sp.Output.Printf("Done.\n")
	return true
}

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
	if share.Counter().Cmp(sp.LatestCounter) <= 0 {
		return false
	}
	sp.ClaimRepo.AddShare(share)
	return true
}

// GetCurrentClaim returns new claim containing unsubmitted shares. If there
// is no new shares, it returns nil.
func (sp *SmartPool) GetCurrentClaim() smartpool.Claim {
	return sp.ClaimRepo.GetCurrentClaim()
}

func (sp *SmartPool) GetVerificationIndex(claim smartpool.Claim) *big.Int {
	return big.NewInt(0)
}

// Submit does all the protocol that communicates with the contract to submit
// the claim then verify it.
// It returns true when the claim is fully verified and accepted by the
// contract. It returns false otherwise.
func (sp *SmartPool) Submit() bool {
	claim := sp.GetCurrentClaim()
	subErr := sp.Contract.SubmitClaim(claim)
	if subErr != nil {
		sp.Output.Printf("Got error submitting claim to contract: %s\n", subErr)
		return false
	}
	index := sp.GetVerificationIndex(claim)
	verErr := sp.Contract.VerifyClaim(index, claim)
	if verErr != nil {
		sp.Output.Printf("Got error verifing claim: %s\n", verErr)
		return false
	}
	return true
}

func NewSmartPool(
	sr smartpool.ShareReceiver, nc smartpool.NetworkClient,
	cr ClaimRepo, uo smartpool.UserOutput, co smartpool.Contract) *SmartPool {
	return &SmartPool{
		ShareReceiver: sr,
		NetworkClient: nc,
		Output:        uo,
		ClaimRepo:     cr,
		Contract:      co,
		// TODO: should be persist between startups instead of having 0 hardcoded
		LatestCounter: big.NewInt(0),
	}
}
