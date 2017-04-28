package fakecontract

import (
	"github.com/ethereum/go-ethereum/common"
)

type PoolMonitor struct {
}

func (pm *PoolMonitor) RequireClientUpdate() bool {
	return false
}

func (pm *PoolMonitor) RequireContractUpdate() bool {
	return false
}

func (pm *PoolMonitor) ContractAddress() common.Address {
	return common.HexToAddress("0x07a457d878bf363e0bb5aa0b096092f941e19962")
}

func NewPoolMonitor() *PoolMonitor {
	return &PoolMonitor{}
}
