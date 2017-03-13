package main

import (
	"../../"
	"../../ethereum"
	"../../ethereum/geth"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/ssh/terminal"
	"math/big"
	"syscall"
	"time"
)

func Initialize() *smartpool.Input {
	// Setting
	rpcEndPoint := "http://localhost:8545"
	// keystorePath := "/Users/victor/Dropbox/Project/BlockChain/SmartPool/spclient_exp/.privatedata/keystore"
	keystorePath := "/Users/victor/Library/Application Support/io.parity.ethereum/keys/kovan"
	shareThreshold := 1
	shareDifficulty := big.NewInt(100000)
	submitInterval := 3 * time.Minute
	contractAddr := "0xc071df9e80d2d13d3f6a7a062a764df4f34c65fd"
	minerAddr := "0x001aDBc838eDe392B5B054A47f8B8c28f2fA9F3F"
	extraData := ""
	return smartpool.NewInput(
		rpcEndPoint, keystorePath, shareThreshold, shareDifficulty,
		submitInterval, contractAddr, minerAddr, extraData,
	)
}

func promptUserPassPhrase(acc string) (string, error) {
	fmt.Printf("Using account address: %s\n", acc)
	fmt.Printf("Please enter passphrase: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Printf("\n")
	if err != nil {
		return "", err
	} else {
		return string(bytePassword), nil
	}
}

func main() {
	input := Initialize()
	gethRPC, _ := ethereum.NewGethRPC(
		input.RPCEndpoint(), input.ContractAddress(), input.ExtraData(),
	)
	var gethContractClient *geth.GethContractClient
	var err error
	for {
		passphrase, _ := promptUserPassPhrase(
			input.MinerAddress(),
		)
		gethContractClient, err = geth.NewGethContractClient(
			common.HexToAddress(input.ContractAddress()), gethRPC,
			input.RPCEndpoint(), input.KeystorePath(), passphrase,
		)
		if gethContractClient != nil {
			break
		} else {
			fmt.Printf("error: %s\n", err)
		}
	}
	ethereumContract := ethereum.NewContract(gethContractClient)
	epoch := 7
	err = ethereumContract.SetEpochData(epoch)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
	} else {
		fmt.Printf("Succeeded.\n", err)
	}
}
