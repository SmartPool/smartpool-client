package main

import (
	"../"
	"../ethereum"
	"../ethereum/ethminer"
	"../ethereum/geth"
	"../protocol"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/ssh/terminal"
	"math/big"
	"syscall"
	"time"
)

func buildExtraData(address common.Address, diff *big.Int) string {
	// id = address % (26+26+10)**11
	base := big.NewInt(0)
	base.Exp(big.NewInt(62), big.NewInt(11), nil)
	id := big.NewInt(0)
	id.Mod(address.Big(), base)
	return fmt.Sprintf("SmartPool-%s%s", smartpool.BigToBase62(id), smartpool.BigToBase62(diff))
}

func Initialize() *smartpool.Input {
	// Setting
	ipcPath := "/Users/victor/Dropbox/Project/BlockChain/SmartPool/spclient_exp/.privatedata/geth.ipc"
	rpcEndPoint := "http://localhost:8545"
	keystorePath := "/Users/victor/Dropbox/Project/BlockChain/SmartPool/spclient_exp/.privatedata/keystore"
	shareThreshold := 1
	shareDifficulty := big.NewInt(100000)
	submitInterval := 10 * time.Second
	contractAddr := "0xeb69b29551f5830581a29858d1aca0e39ec14d57"
	minerAddr := "0xad42beeb07db31149f5d2c4bd33d01c6d7c34116"
	extraData := buildExtraData(
		common.HexToAddress(minerAddr),
		shareDifficulty)
	inp := smartpool.NewInput(
		ipcPath, rpcEndPoint, keystorePath, shareThreshold, shareDifficulty,
		submitInterval, contractAddr, minerAddr, extraData,
	)
	return inp
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
	output := &smartpool.StdOut{}
	ethereumWorkPool := &ethereum.WorkPool{}
	gethRPC, err := ethereum.NewGethRPC(
		input.RPCEndpoint(), input.ContractAddress(), input.ExtraData(),
	)
	if err != nil {
		fmt.Printf("Geth RPC server is unavailable.\n")
		fmt.Printf("Make sure you have Geth installed. If you do, you can run geth by following command (Note: --etherbase and --extradata params are required.):\n")
		fmt.Printf(
			"geth --rpc --etherbase \"%s\" --extradata \"%s\"\n",
			input.ContractAddress(), input.ExtraData())
		return
	}
	ethereumNetworkClient := ethereum.NewNetworkClient(
		gethRPC,
		ethereumWorkPool,
	)
	ethereumClaimRepo := protocol.NewInMemClaimRepo()
	var gethContractClient *geth.GethContractClient
	for {
		passphrase, _ := promptUserPassPhrase(
			input.MinerAddress(),
		)
		gethContractClient, err = geth.NewGethContractClient(
			common.HexToAddress(input.ContractAddress()), gethRPC,
			input.IPCPath(), input.KeystorePath(), passphrase,
		)
		if gethContractClient != nil {
			break
		} else {
			fmt.Printf("error: %s\n", err)
		}
	}
	ethereumContract := ethereum.NewContract(gethContractClient)
	ethminer.SmartPool = protocol.NewSmartPool(
		ethereumWorkPool, ethereumNetworkClient,
		ethereumClaimRepo, output, ethereumContract,
		input.SubmitInterval(), input.ShareThreshold(),
	)
	server := ethminer.NewRPCServer(
		output,
		uint16(1633),
	)
	server.Start()
}
