package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum/ethash"
	"github.com/SmartPool/smartpool-client/ethereum/geth"
	"github.com/SmartPool/smartpool-client/mtree"
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
	if err != nil {
		fmt.Printf("Failed to create authorized transactor: %s\n", err)
		return
	}
	epoch := 23
	fmt.Printf("Checking DAG file for epoch %d. Generate if needed...\n", epoch)
	seedHash, err := ethash.GetSeedHash(uint64(epoch * 30000))
	if err != nil {
		panic(err)
	}
	path := filepath.Join(
		ethash.DefaultDir,
		fmt.Sprintf("full-R%s-%s", "23", hex.EncodeToString(seedHash[:8])),
	)
	mt := mtree.NewDagTree()
	mt.RegisterIndex(4744867)
	fullSize, _ := ethash.MakeDAGWithSize(uint64(epoch*30000), "")
	fullSizeIn128Resolution := fullSize / 128
	branchDepth := len(fmt.Sprintf("%b", fullSizeIn128Resolution-1))
	mt.RegisterStoredLevel(uint32(branchDepth), uint32(10))
	processDuringRead(path, mt)
	mt.Finalize()
	elements := []*big.Int{}
	for _, w := range mt.AllDAGElements() {
		elements = append(elements, w.ToUint256Array()...)
	}
	proof := []*big.Int{}
	branch := mt.AllBranchesArray()
	fmt.Printf("len branch: %d\n", len(branch))
	for i := 0; i < len(branch); i++ {
		proof = append(proof, branch[i].Big())
	}
	result, err := testClient.TestOptimization(
		nil,
		[]*big.Int{big.NewInt(4744867)},
		elements,
		proof,
		big.NewInt(23),
	)
	fmt.Printf("Elements: [")
	for _, e := range elements {
		fmt.Printf("0x%s, ", e.Text(16))
	}
	fmt.Printf("]\n")
	fmt.Printf("Proof: [")
	for _, e := range proof {
		fmt.Printf("0x%s, ", e.Text(16))
	}
	fmt.Printf("]\n")
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("0x%s 0x%s 0x%s 0x%s", result[0].Text(16), result[1].Text(16), result[2].Text(16), result[3].Text(16))
}
