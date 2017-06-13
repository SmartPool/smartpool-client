package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum/ethash"
	"github.com/SmartPool/smartpool-client/ethereum/geth"
	"github.com/SmartPool/smartpool-client/mtree"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"log"
	"math/big"
	"os"
	"path/filepath"
)

func processDuringRead(
	datasetPath string, mt *mtree.DagTree) {

	f, err := os.Open(datasetPath)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	buf := [128]byte{}
	// ignore first 8 bytes magic number at the beginning
	// of dataset. See more at https://github.com/ethereum/wiki/wiki/Ethash-DAG-Disk-Storage-Format
	_, err = io.ReadFull(r, buf[:8])
	if err != nil {
		log.Fatal(err)
	}
	var i uint32 = 0
	for {
		n, err := io.ReadFull(r, buf[:128])
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if n != 128 {
			log.Fatal("Malformed dataset")
		}
		mt.Insert(smartpool.Word(buf), i)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		i++
	}
}

func getClient(rpc string) (*ethclient.Client, error) {
	return ethclient.Dial(rpc)
}

func main() {
	client, err := getClient("http://localhost:8545")
	if err != nil {
		fmt.Printf("Couldn't connect to Geth via IPC file. Error: %s\n", err)
		return
	}
	contractAddr := common.HexToAddress("0xda87714c91d62070ebc29675ec79a190e6ccdfba")
	testClient, err := geth.NewTestClient(contractAddr, client)
	if err != nil {
		fmt.Printf("Couldn't bind. Error: %s\n", err)
		return
	}
	account := geth.GetAccount(
		"/Users/victor/Library/Ethereum/testnet/keystore",
		common.HexToAddress("0xe034afdcc2ba0441ff215ee9ba0da3e86450108d"),
		"hoctinhocsong")
	if account == nil {
		fmt.Printf("Couldn't get any account from key store.\n")
		return
	}
	keyio, err := os.Open(account.KeyFile())
	if err != nil {
		fmt.Printf("Failed to open key file: %s\n", err)
		return
	}
	fmt.Printf("Unlocking account...")
	auth, err := bind.NewTransactor(keyio, account.PassPhrase())
	if err != nil {
		fmt.Printf("Failed to create authorized transactor: %s\n", err)
		return
	}
	epoch := 24
	fmt.Printf("Checking DAG file for epoch %d. Generate if needed...\n", epoch)
	fullSize, _ := ethash.MakeDAGWithSize(uint64(epoch*30000), "")
	fullSizeIn128Resolution := fullSize / 128
	seedHash, err := ethash.GetSeedHash(uint64(epoch * 30000))
	if err != nil {
		panic(err)
	}
	path := filepath.Join(
		ethash.DefaultDir,
		fmt.Sprintf("full-R%s-%s", "23", hex.EncodeToString(seedHash[:8])),
	)
	branchDepth := len(fmt.Sprintf("%b", fullSizeIn128Resolution-1))
	mt := mtree.NewDagTree()
	mt.RegisterStoredLevel(uint32(branchDepth), uint32(10))
	processDuringRead(path, mt)
	mt.Finalize()
	merkleNodes := []*big.Int{}
	fmt.Printf("len nodes: %d\n", len(mt.MerkleNodes()))
	start := big.NewInt(0)
	var temp int
	for k, n := range mt.MerkleNodes() {
		merkleNodes = append(merkleNodes, n)
		if len(merkleNodes) == 50 || k == len(mt.MerkleNodes())-1 {
			mnlen := big.NewInt(int64(len(merkleNodes)))
			tx, err := testClient.SetOptEpochData(
				auth,
				big.NewInt(int64(epoch)),
				big.NewInt(int64(fullSizeIn128Resolution)),
				big.NewInt(int64(branchDepth-10)),
				merkleNodes,
				start,
				mnlen,
			)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			fmt.Printf("Transaction: %s\n", tx.Hash().Hex())
			start.Add(start, mnlen)
			merkleNodes = []*big.Int{}
			fmt.Scanf("%d", &temp)
		}
	}
}
