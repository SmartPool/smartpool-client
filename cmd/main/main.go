package main

import (
	"../../"
	"../../ethereum"
	"../../ethereum/ethminer"
	"../../ethereum/geth"
	"../../protocol"
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
	rpcEndPoint := "http://localhost:8545"
	keystorePath := "/Users/victor/Dropbox/Project/BlockChain/SmartPool/spclient_exp/.privatedata/keystore"
	shareThreshold := 1
	shareDifficulty := big.NewInt(100000)
	submitInterval := 3 * time.Minute
	contractAddr := "0xc071df9e80d2d13d3f6a7a062a764df4f34c65fd"
	minerAddr := "0x001aDBc838eDe392B5B054A47f8B8c28f2fA9F3F"
	extraData := buildExtraData(common.HexToAddress(minerAddr), shareDifficulty)
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
	output := &smartpool.StdOut{}
	ethereumWorkPool := &ethereum.WorkPool{}
	gethRPC, _ := geth.NewGethRPC(
		input.RPCEndpoint(), input.ContractAddress(),
		input.ExtraData(), input.ShareDifficulty(),
	)
	client, err := gethRPC.ClientVersion()
	if err != nil {
		fmt.Printf("Node RPC server is unavailable.\n")
		fmt.Printf("Make sure you have Geth or Parity installed. If you do, you can:\nRun Geth by following command (Note: --etherbase and --extradata params are required.):\n")
		fmt.Printf(
			"geth --rpc --etherbase \"%s\" --extradata \"%s\"\n",
			input.ContractAddress(), input.ExtraData())
		fmt.Printf("Or run Parity by following command (Note: --etherbase and --extradata params are required.)\n")
		fmt.Printf(
			"parity --author \"%s\" --extra-data \"%s\"\n",
			input.ContractAddress(), input.ExtraData())
		return
	}
	fmt.Printf("Connected to Ethereum node: %s\n", client)
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
			common.HexToAddress(input.MinerAddress()),
			input.RPCEndpoint(), input.KeystorePath(), passphrase,
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
		ethereumClaimRepo, ethereumContract,
		common.HexToAddress(input.MinerAddress()),
		input.SubmitInterval(), input.ShareThreshold(),
	)
	server := ethminer.NewRPCServer(
		output,
		uint16(1633),
	)
	server.Start()
}
