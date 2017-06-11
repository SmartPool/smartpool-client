package geth

import (
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
)

type EthashContractClient struct {
	contract   *Ethash
	transactor *bind.TransactOpts
	node       ethereum.RPCClient
	sender     common.Address
}

func (cc *EthashContractClient) SetEpochData(
	epoch *big.Int,
	fullSizeIn128Resolution *big.Int,
	branchDepth *big.Int,
	merkleNodes []*big.Int) error {

	nodes := []*big.Int{}
	start := big.NewInt(0)
	fmt.Printf("No meaningful nodes: %d\n", len(merkleNodes))
	for k, n := range merkleNodes {
		nodes = append(nodes, n)
		if len(nodes) == 40 || k == len(merkleNodes)-1 {
			mnlen := big.NewInt(int64(len(nodes)))
			blockNo, err := cc.node.BlockNumber()
			blockNo.Add(blockNo, big.NewInt(1))
			if err != nil {
				smartpool.Output.Printf("Setting epoch data. Error: %s\n", err)
				return err
			}
			fmt.Printf("Going to do tx\n")
			fmt.Printf("Block Number: %d\n", blockNo.Int64())
			tx, err := cc.contract.SetEpochData(
				cc.transactor, epoch, fullSizeIn128Resolution,
				branchDepth, nodes, start, mnlen)
			if err != nil {
				smartpool.Output.Printf("Setting optimized epoch data. Error: %s\n", err)
				return err
			}
			errCode, errInfo, err := GetTxResult(
				tx, cc.transactor, cc.node, blockNo, SetEpochDataEventTopic,
				cc.sender.Big())
			if err != nil {
				smartpool.Output.Printf("Tx: %s was not approved by the network in time.\n", tx.Hash().Hex())
				return err
			}
			if errCode.Cmp(common.Big0) != 0 {
				smartpool.Output.Printf("Error code: 0x%s - Error info: 0x%s\n", errCode.Text(16), errInfo.Text(16))
				// return errors.New(ErrorMsg(errCode, errInfo))
			}
			start.Add(start, mnlen)
			nodes = []*big.Int{}
		}
	}
	return nil
}

func NewEthashContractClient(
	contractAddr common.Address, node ethereum.RPCClient, miner common.Address,
	ipc, keystorePath, passphrase string, gasprice uint64) (*EthashContractClient, error) {
	client, err := getClient(ipc)
	if err != nil {
		smartpool.Output.Printf("Couldn't connect to Geth/Parity. Error: %s\n", err)
		return nil, err
	}
	ethash, err := NewEthash(contractAddr, client)
	if err != nil {
		smartpool.Output.Printf("Couldn't get SmartPool information from Ethereum Blockchain. Error: %s\n", err)
		return nil, err
	}
	account := GetAccount(keystorePath, miner, passphrase)
	if account == nil {
		smartpool.Output.Printf("Couldn't get any account from key store.\n")
		return nil, err
	}
	keyio, err := os.Open(account.KeyFile())
	if err != nil {
		smartpool.Output.Printf("Failed to open key file: %s\n", err)
		return nil, err
	}
	smartpool.Output.Printf("Unlocking account...")
	auth, err := bind.NewTransactor(keyio, account.PassPhrase())
	if err != nil {
		smartpool.Output.Printf("Failed to create authorized transactor: %s\n", err)
		return nil, err
	}
	if gasprice != 0 {
		auth.GasPrice = big.NewInt(int64(gasprice * 1000000000))
		smartpool.Output.Printf("Gas price is set to: %s wei.\n", auth.GasPrice.Text(10))
	}
	smartpool.Output.Printf("Done.\n")
	return &EthashContractClient{ethash, auth, node, miner}, nil
}
