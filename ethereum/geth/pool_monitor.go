package geth

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"strings"
)

// var PoolMonitorAddress = common.HexToAddress("0x7727D4535f1A9c9ECC59FB17B6bF8145C7d5D58c")
// var PoolMonitorAddress = common.HexToAddress("0xddcdad6b099b1b237bdb1341cc6881eb63ee3b28")

type PoolMonitor struct {
	client   *PoolMonitorClient
	version  [3]byte
	contract common.Address
}

func (pm *PoolMonitor) ContractAddress() common.Address {
	address, err := pm.client.PoolContract(nil)
	if err != nil {
		smartpool.Output.Printf("Getting pool contract from gateway failed. Error: %s\n", err)
	}
	return address
}

func (pm *PoolMonitor) RequireClientUpdate() bool {
	version, err := pm.client.ClientVersion(nil)
	if err != nil {
		smartpool.Output.Printf("Getting client version from gateway failed. Error: %s\n", err)
		return false
	}
	if version[0] != pm.version[0] || version[1] != pm.version[1] {
		return true
	}
	return false
}

func (pm *PoolMonitor) RequireContractUpdate() bool {
	addr, err := pm.client.PoolContract(nil)
	if err != nil {
		smartpool.Output.Printf("Getting contract address from gateway failed. Error: %s\n", err)
		return false
	}
	return addr != pm.contract
}

func versionBytes(version string) [3]byte {
	vs := strings.Split(version, ".")
	major, _ := strconv.Atoi(vs[0])
	minor, _ := strconv.Atoi(vs[1])
	patch, _ := strconv.Atoi(vs[2])
	return [3]byte{byte(major), byte(minor), byte(patch)}
}

func NewPoolMonitor(
	gatewayAddr common.Address, contractAddr common.Address,
	version string, ipc string) (*PoolMonitor, error) {
	client, err := getClient(ipc)
	if err != nil {
		smartpool.Output.Printf("Couldn't connect to Geth via IPC file. Error: %s\n", err)
		return nil, err
	}
	poolMonitor, err := NewPoolMonitorClient(gatewayAddr, client)
	if err != nil {
		smartpool.Output.Printf("Couldn't get SmartPool information from Ethereum Blockchain. Error: %s\n", err)
		return nil, err
	}
	return &PoolMonitor{
		poolMonitor,
		versionBytes(version),
		contractAddr}, nil
}
