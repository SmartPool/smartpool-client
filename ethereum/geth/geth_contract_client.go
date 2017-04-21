package geth

import (
	"errors"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"time"
)

type GethContractClient struct {
	// the contract implementation that holds all underlying
	// communication with Ethereum Contract
	pool       *TestPool
	transactor *bind.TransactOpts
	node       ethereum.RPCClient
	sender     common.Address
}

func (cc *GethContractClient) Version() string {
	v, err := cc.pool.Version(nil)
	if err != nil {
		smartpool.Output.Printf("Couldn't get contract version: %s\n", err)
		return ""
	}
	return v
}

func (cc *GethContractClient) IsRegistered() bool {
	ok, err := cc.pool.IsRegistered(nil, cc.sender)
	if err != nil {
		smartpool.Output.Printf("Couldn't check the address's registration: %s\n", err)
		return false
	}
	return ok
}

func (cc *GethContractClient) CanRegister() bool {
	ok, err := cc.pool.CanRegister(nil, cc.sender)
	if err != nil {
		smartpool.Output.Printf("Couldn't check slot availability for the address: %s\n", err)
		return false
	}
	return ok
}

func (cc *GethContractClient) Register(paymentAddress common.Address) error {
	blockNo, err := cc.node.BlockNumber()
	if err != nil {
		return err
	}
	tx, err := cc.pool.Register(cc.transactor, paymentAddress)
	if err != nil {
		return err
	}
	smartpool.Output.Printf("Registering address %s to SmartPool contract by tx: %s\n", paymentAddress.Hex(), tx.Hash().Hex())
	errCode, errInfo, err := NewTxWatcher(
		tx, cc.node, blockNo, RegisterEventTopic,
		cc.sender.Big()).Wait()
	if err != nil {
		smartpool.Output.Printf("Tx: %s was not approved by the network in time.\n", tx.Hash().Hex())
		return err
	}
	if errCode.Cmp(common.Big0) != 0 {
		smartpool.Output.Printf("Error code: 0x%s - Error info: 0x%s\n", errCode.Text(16), errInfo.Text(16))
		return errors.New(ErrorMsg(errCode, errInfo))
	}
	smartpool.Output.Printf("Registered address %s to SmartPool contract. Tx %s is confirmed\n", paymentAddress.Hex(), tx.Hash().Hex())
	return nil
}

func (cc *GethContractClient) GetClaimSeed() *big.Int {
	var seed *big.Int
	var err error
	// Wait for 30s because the seed is only available after several blocks
	time.Sleep(30 * time.Second)
	for {
		seed, err = cc.pool.GetClaimSeed(nil, cc.sender)
		if err != nil {
			smartpool.Output.Printf("Getting claim seed failed. Error: %s\n", err)
			return big.NewInt(0)
		}
		if seed.Cmp(common.Big0) != 0 {
			break
		}
		time.Sleep(time.Second)
	}
	return seed
}

func (cc *GethContractClient) SubmitClaim(
	numShares *big.Int,
	difficulty *big.Int,
	min *big.Int,
	max *big.Int,
	augMerkle *big.Int) error {
	blockNo, err := cc.node.BlockNumber()
	if err != nil {
		smartpool.Output.Printf("Submitting claim failed. Error: %s\n", err)
		return err
	}
	tx, err := cc.pool.SubmitClaim(cc.transactor,
		numShares, difficulty, min, max, augMerkle)
	if err != nil {
		smartpool.Output.Printf("Submitting claim failed. Error: %s\n", err)
		return err
	}
	errCode, errInfo, err := NewTxWatcher(
		tx, cc.node, blockNo, SubmitClaimEventTopic,
		cc.sender.Big()).Wait()
	if err != nil {
		smartpool.Output.Printf("Tx: %s was not approved by the network in time.\n", tx.Hash().Hex())
		return err
	}
	if errCode.Cmp(common.Big0) != 0 {
		smartpool.Output.Printf("Error code: 0x%s - Error info: 0x%s\n", errCode.Text(16), errInfo.Text(16))
		return errors.New(ErrorMsg(errCode, errInfo))
	}
	return nil
}

func (cc *GethContractClient) VerifyClaim(
	rlpHeader []byte,
	nonce *big.Int,
	shareIndex *big.Int,
	dataSetLookup []*big.Int,
	witnessForLookup []*big.Int,
	augCountersBranch []*big.Int,
	augHashesBranch []*big.Int) error {
	blockNo, err := cc.node.BlockNumber()
	if err != nil {
		smartpool.Output.Printf("Verifying claim failed. Error: %s\n", err)
		return err
	}
	tx, err := cc.pool.VerifyClaim(cc.transactor,
		rlpHeader, nonce, shareIndex, dataSetLookup,
		witnessForLookup, augCountersBranch, augHashesBranch)
	if err != nil {
		smartpool.Output.Printf("Verifying claim failed. Error: %s\n", err)
		return err
	}
	errCode, errInfo, err := NewTxWatcher(
		tx, cc.node, blockNo, VerifyClaimEventTopic,
		cc.sender.Big()).Wait()
	if err != nil {
		smartpool.Output.Printf("Tx: %s was not approved by the network in time.\n", tx.Hash().Hex())
		return err
	}
	if errCode.Cmp(common.Big0) != 0 {
		smartpool.Output.Printf("Error code: 0x%s - Error info: 0x%s\n", errCode.Text(16), errInfo.Text(16))
		return errors.New(ErrorMsg(errCode, errInfo))
	}
	return nil
}

func (cc *GethContractClient) SetEpochData(
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
			tx, err := cc.pool.SetEpochData(
				cc.transactor, epoch, fullSizeIn128Resolution,
				branchDepth, nodes, start, mnlen)
			if err != nil {
				smartpool.Output.Printf("Setting optimized epoch data. Error: %s\n", err)
				return err
			}
			errCode, errInfo, err := NewTxWatcher(
				tx, cc.node, blockNo, SetEpochDataEventTopic,
				cc.sender.Big()).Wait()
			if err != nil {
				smartpool.Output.Printf("Tx: %s was not approved by the network in time.\n", tx.Hash().Hex())
				return err
			}
			if errCode.Cmp(common.Big0) != 0 {
				smartpool.Output.Printf("Error code: 0x%s - Error info: 0x%s\n", errCode.Text(16), errInfo.Text(16))
				return errors.New(ErrorMsg(errCode, errInfo))
			}
			start.Add(start, mnlen)
			nodes = []*big.Int{}
		}
	}
	return nil
}

func getClient(rpc string) (*ethclient.Client, error) {
	return ethclient.Dial(rpc)
}

func NewGethContractClient(
	contractAddr common.Address, node ethereum.RPCClient, miner common.Address,
	ipc, keystorePath, passphrase string) (*GethContractClient, error) {
	client, err := getClient(ipc)
	if err != nil {
		smartpool.Output.Printf("Couldn't connect to Geth via IPC file. Error: %s\n", err)
		return nil, err
	}
	pool, err := NewTestPool(contractAddr, client)
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
	// TODO: make gas price one command line flag
	auth.GasPrice = big.NewInt(20000000000)
	smartpool.Output.Printf("Done.\n")
	return &GethContractClient{pool, auth, node, miner}, nil
}
