package protocol

type testPoolMonitor struct {
	ClientUpdate   bool
	ContractUpdate bool
}

func (pm *testPoolMonitor) RequireClientUpdate() bool {
	return pm.ClientUpdate
}

func (pm *testPoolMonitor) RequireContractUpdate() bool {
	return pm.ContractUpdate
}
