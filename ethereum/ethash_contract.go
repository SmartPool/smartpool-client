package ethereum

import (
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum/ethash"
	"github.com/SmartPool/smartpool-client/mtree"
	"math/big"
)

type EthashContract struct {
	ethashClient EthashContractClient
}

func (c *EthashContract) SetEpochData(epoch int) error {
	var err error
	smartpool.Output.Printf("Checking DAG file. Generate if needed...\n")
	ethash.MakeDataset(uint64(epoch*30000), ethash.DefaultDir)
	fullSize := ethash.DAGSize(uint64(epoch * 30000))
	fullSizeIn128Resolution := fullSize / 128
	path := ethash.PathToDAG(uint64(epoch), ethash.DefaultDir)
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
