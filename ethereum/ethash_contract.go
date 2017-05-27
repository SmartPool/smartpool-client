package ethereum

import (
	"encoding/hex"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum/ethash"
	"github.com/SmartPool/smartpool-client/mtree"
	"math/big"
	"path/filepath"
)

type EthashContract struct {
	ethashClient EthashContractClient
}

func (c *EthashContract) SetEpochData(epoch int) error {
	smartpool.Output.Printf("Checking DAG file. Generate if needed...\n")
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
	// TODO: 10 is just an experimental level
	mt.RegisterStoredLevel(uint32(branchDepth), 10)
	processDuringRead(path, mt)
	mt.Finalize()
	err = c.ethashClient.SetEpochData(
		big.NewInt(int64(epoch)),
		big.NewInt(int64(fullSizeIn128Resolution)),
		big.NewInt(int64(branchDepth-10)),
		mt.MerkleNodes(),
	)
	if err != nil {
		fmt.Printf("Got error: %s\n", err)
		return err
	}
	fmt.Printf("Done.\n")
	return nil
}

func NewEthashContract(client EthashContractClient) *EthashContract {
	return &EthashContract{client}
}
