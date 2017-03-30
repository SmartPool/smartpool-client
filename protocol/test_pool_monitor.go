package protocol

import (
	"github.com/ethereum/go-ethereum/common"
)

type testPoolMonitor struct {
	ClientUpdate   bool
	ContractUpdate bool
	ContractAddr   common.Address
}

func (pm *testPoolMonitor) RequireClientUpdate() bool {
	return pm.ClientUpdate
}

func (pm *testPoolMonitor) ContractAddress() common.Address {
	return pm.ContractAddr
}

func (pm *testPoolMonitor) RequireContractUpdate() bool {
	return pm.ContractUpdate
}
