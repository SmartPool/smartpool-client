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
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"os"
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

func Initialize(c *cli.Context) *smartpool.Input {
	// Setting
	// rpcEndPoint := "http://localhost:8545"
	rpcEndPoint := c.String("rpc")
	// rpcEndPoint := "/Users/victor/Library/Application Support/io.parity.ethereum/jsonrpc.ipc"
	// keystorePath := "/Users/victor/Library/Application Support/io.parity.ethereum/keys/kovan"
	keystorePath := c.String("keystore")
	// shareThreshold := 5
	shareThreshold := int(c.Uint("threshold"))
	// shareDifficulty := big.NewInt(100000)
	shareDifficulty := big.NewInt(int64(c.Uint("diff")))
	submitInterval := 1 * time.Minute
	// contractAddr := "0x92a71342C2EaBc92d09b83a8C82D48F41C0ddbaf"
	contractAddr := c.String("spcontract")
	// minerAddr := "0x001aDBc838eDe392B5B054A47f8B8c28f2fA9F3F"
	minerAddr := c.String("miner")
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

func Run(c *cli.Context) error {
	input := Initialize(c)
	output := &smartpool.StdOut{}
	ethereumWorkPool := &ethereum.WorkPool{}
	address, ok := geth.GetAddress(
		input.KeystorePath(),
		common.HexToAddress(input.MinerAddress()),
	)
	input.SetMinerAddress(address)
	input.SetExtraData(buildExtraData(
		common.HexToAddress(input.MinerAddress()),
		input.ShareDifficulty()))
	kovanRPC, _ := geth.NewKovanRPC(
		input.RPCEndpoint(), input.ContractAddress(), input.ExtraData(),
	)
	client, err := kovanRPC.ClientVersion()
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
		return err
	}
	fmt.Printf("Connected to Ethereum node: %s\n", client)
	ethereumNetworkClient := ethereum.NewNetworkClient(
		kovanRPC,
		ethereumWorkPool,
	)
	ethereumClaimRepo := protocol.NewInMemClaimRepo()
	var gethContractClient *geth.GethContractClient
	for {
		if ok {
			passphrase, _ := promptUserPassPhrase(
				input.MinerAddress(),
			)
			gethContractClient, err = geth.NewGethContractClient(
				common.HexToAddress(input.ContractAddress()), kovanRPC,
				common.HexToAddress(input.MinerAddress()),
				input.RPCEndpoint(), input.KeystorePath(), passphrase,
			)
			if gethContractClient != nil {
				break
			} else {
				fmt.Printf("error: %s\n", err)
			}
		} else {
			if input.KeystorePath() == "" {
				fmt.Printf("You have to specify keystore path by --keystore. Abort!\n")
			} else {
				fmt.Printf("Your keystore: %s\n", input.KeystorePath())
				fmt.Printf("Your miner address: %s\n", input.MinerAddress())
				fmt.Printf("We couldn't find your miner address private key in the keystore path you specified.\nPlease make sure your keystore path exists.\nAbort!\n")
			}
			return nil
		}
	}
	ethereumContract := ethereum.NewContract(gethContractClient)
	ethminer.SmartPool = protocol.NewSmartPool(
		ethereumWorkPool, ethereumNetworkClient,
		ethereumClaimRepo, output, ethereumContract,
		common.HexToAddress(input.MinerAddress()),
		input.SubmitInterval(), input.ShareThreshold(),
	)
	server := ethminer.NewRPCServer(
		output,
		uint16(1633),
	)
	server.Start()
	return nil
}

func BuildAppCommandLine() *cli.App {
	app := cli.NewApp()
	app.Description = "Efficient Decentralized Mining Pools for Existing Cryptocurrencies Based on Ethereum Smart Contracts"
	app.Name = "SmartPool commandline tool"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "rpc",
			Value: "http://localhost:8545",
			Usage: "RPC endpoint of Ethereum node",
		},
		cli.StringFlag{
			Name:  "keystore",
			Usage: "Keystore path to your ethereum account private key. SmartPool will look for private key of the miner address you specified in that path.",
		},
		cli.UintFlag{
			Name:  "threshold",
			Value: 10,
			Usage: "Minimum number of shares in a claim. SmartPool will not submit the claim if it does not have more than or equal to this threshold numer of share.",
		},
		cli.UintFlag{
			Name:  "diff",
			Value: 100000,
			Usage: "Difficulty of a share.",
		},
		cli.StringFlag{
			Name:  "spcontract",
			Value: "0x3dC682397e93E46EBb5bE7463658fdD658365e9D",
			Usage: "SmartPool latest contract address.",
		},
		cli.StringFlag{
			Name:  "miner",
			Usage: "The address that would be paid by SmartPool. This is often your address. (Default: First account in your keystore.)",
		},
	}
	app.Action = Run
	return app
}

func main() {
	app := BuildAppCommandLine()
	app.Run(os.Args)
}
