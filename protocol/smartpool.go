// Package protocol implements smartpool secured protocol between client side
// and contract side. It works on high abstraction level of interfaces and
// types of smartpool package.
package protocol

import (
	"../"
	"github.com/ethereum/go-ethereum/common"
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

// AcceptSolution takes Solution to
func (sp *SmartPool) AcceptSolution(s smartpool.Solution) bool {
	sp.ShareReceiver.AcceptSolution(s)
	return true
}
