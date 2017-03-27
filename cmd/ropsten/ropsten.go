package main

import (
	"errors"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/SmartPool/smartpool-client/ethereum/ethminer"
	"github.com/SmartPool/smartpool-client/ethereum/geth"
	"github.com/SmartPool/smartpool-client/protocol"
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
	rpcEndPoint := c.String("rpc")
	keystorePath := c.String("keystore")
	shareThreshold := int(c.Uint("threshold"))
	shareDifficulty := big.NewInt(int64(c.Uint("diff")))
	submitInterval := 1 * time.Minute
	contractAddr := c.String("spcontract")
	minerAddr := c.String("miner")
	hotStop := !c.Bool("no-hot-stop")
	if hotStop {
		fmt.Printf("SmartPool is in Hot-Stop mode: It will exit immediately if the contract returns errors.\n")
	}
	extraData := ""
	return smartpool.NewInput(
		rpcEndPoint, keystorePath, shareThreshold, shareDifficulty,
		submitInterval, contractAddr, minerAddr, extraData, hotStop,
	)
}

func promptUserPassPhrase(acc string) (string, error) {
	fmt.Printf("Using miner address: %s\n", acc)
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
	if input.KeystorePath() == "" {
		fmt.Printf("You have to specify keystore path by --keystore. Abort!\n")
		return nil
	}
	smartpool.Output = &smartpool.StdOut{}
	ethereumWorkPool := &ethereum.WorkPool{}
	go ethereumWorkPool.Cleanning()
	address, ok, addresses := geth.GetAddress(
		input.KeystorePath(),
		common.HexToAddress(input.MinerAddress()),
	)
	if len(addresses) == 0 {
		fmt.Printf("We couldn't find any private keys in your keystore path.\n")
		fmt.Printf("Please make sure your keystore path exists.\nAbort!\n")
		return nil
	}
	fmt.Printf("Using miner address: %s\n", address.Hex())
	input.SetMinerAddress(address)
	input.SetExtraData(buildExtraData(
		common.HexToAddress(input.MinerAddress()),
		input.ShareDifficulty()))
	gethRPC, _ := geth.NewGethRPC(
		input.RPCEndpoint(), input.ContractAddress(),
		input.ExtraData(), input.ShareDifficulty(),
	)
	client, err := gethRPC.ClientVersion()
	if err != nil {
		fmt.Printf("Node RPC server is unavailable.\n")
		fmt.Printf("Make sure you have Geth or Parity installed. If you do, you can:\nRun Geth by following command (Note: --etherbase and --extradata params are required.):\n")
		fmt.Printf(
			"geth --testnet --rpc --etherbase \"%s\" --extradata \"%s\"\n",
			input.ContractAddress(), input.ExtraData())
		fmt.Printf("Or run Parity by following command (Note: --author and --extra-data params are required.)\n")
		fmt.Printf(
			"parity --chain ropsten --author \"%s\" --extra-data \"%s\"\n",
			input.ContractAddress(), input.ExtraData())
		return err
	}
	fmt.Printf("Connected to Ethereum node: %s\n", client)
	ethereumNetworkClient := ethereum.NewNetworkClient(
		gethRPC,
		ethereumWorkPool,
	)
	ethereumClaimRepo := ethereum.NewTimestampClaimRepo()
	ethereumPoolMonitor, err := geth.NewPoolMonitor(
		common.HexToAddress(input.ContractAddress()),
		smartpool.VERSION,
		input.RPCEndpoint(),
	)
	contractAddr := ethereumPoolMonitor.ContractAddress()
	if err != nil {
		fmt.Printf("Couln't connect to gateway.\n")
		return err
	}
	if contractAddr.Big().Cmp(common.Big0) == 0 {
		fmt.Printf("Couldn't get SmartPool contract address from gateway.\n")
		return errors.New("Contract address is not set on the gateway")
	}
	var gethContractClient *geth.GethContractClient
	for {
		if ok {
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
		} else {
			if input.KeystorePath() == "" {
				fmt.Printf("You have to specify keystore path by --keystore. Abort!\n")
			} else {
				fmt.Printf("Your keystore: %s\n", input.KeystorePath())
				fmt.Printf("Your miner address: %s\n", input.MinerAddress())
				if len(addresses) > 0 {
					fmt.Printf("We couldn't find the private key of your miner address in the keystore path you specified. We found following addresses:\n")
					for i, addr := range addresses {
						fmt.Printf("%d. %s\n", i+1, addr.Hex())
					}
					fmt.Printf("Please make sure you entered correct miner address.\n")
				} else {
					fmt.Printf("We couldn't find any private keys in your keystore path.\n")
					fmt.Printf("Please make sure your keystore path exists.\nAbort!\n")
				}
			}
			return nil
		}
	}
	ethereumContract := ethereum.NewContract(gethContractClient)
	ethminer.SmartPool = protocol.NewSmartPool(
		ethereumPoolMonitor,
		ethereumWorkPool, ethereumNetworkClient,
		ethereumClaimRepo, ethereumContract,
		common.HexToAddress(input.MinerAddress()),
		input.SubmitInterval(), input.ShareThreshold(),
		input.HotStop(),
	)
	server := ethminer.NewRPCServer(
		smartpool.Output,
		uint16(1633),
	)
	server.Start()
	return nil
}

func BuildAppCommandLine() *cli.App {
	app := cli.NewApp()
	app.Description = "Efficient Decentralized Mining Pools for Existing Cryptocurrencies Based on Ethereum Smart Contracts"
	app.Name = "SmartPool commandline tool"
	app.Usage = "SmartPool client for ropsten ethereum chain"
	app.Version = smartpool.VERSION
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
			Value: 50,
			Usage: "Minimum number of shares in a claim. SmartPool will not submit the claim if it does not have more than or equal to this threshold numer of share.",
		},
		cli.UintFlag{
			Name:  "diff",
			Value: 1000000,
			Usage: "Difficulty of a share.",
		},
		cli.StringFlag{
			Name:  "spcontract",
			Value: "0xf7d93BCB8e4372F46383ecee82f9adF1aA397BA9",
			Usage: "SmartPool latest contract address.",
		},
		cli.StringFlag{
			Name:  "miner",
			Usage: "The address that would be paid by SmartPool. This is often your address. (Default: First account in your keystore.)",
		},
		cli.BoolFlag{
			Name:  "no-hot-stop",
			Usage: "If hot-stop is true, SmartPool will stop running once it got an error returned from the Contract",
		},
	}
	app.Action = Run
	return app
}

func main() {
	app := BuildAppCommandLine()
	app.Run(os.Args)
}
