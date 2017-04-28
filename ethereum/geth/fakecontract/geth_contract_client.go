package fakecontract

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
)

type GethContractClient struct {
}

func (gcc *GethContractClient) Version() string {
	return "0.0.4"
}
func (gcc *GethContractClient) IsRegistered() bool {
	return true
}
func (gcc *GethContractClient) CanRegister() bool {
	return true
}
func (gcc *GethContractClient) Register(paymentAddress common.Address) error {
	return nil
}
func (gcc *GethContractClient) GetClaimSeed() *big.Int {
	return big.NewInt(1)
}
func (gcc *GethContractClient) SubmitClaim(
	numShares *big.Int,
	difficulty *big.Int,
	min *big.Int,
	max *big.Int,
	augMerkle *big.Int) error {
	filename := "test_input.js"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("//====================Got submission======================\n"))
	f.WriteString(fmt.Sprintf("var noShares = \"0x%s\"\n", numShares.Text(16)))
	f.WriteString(fmt.Sprintf("var diff = \"0x%s\"\n", difficulty.Text(16)))
	f.WriteString(fmt.Sprintf("var AugMin = \"0x%s\"\n", min.Text(16)))
	f.WriteString(fmt.Sprintf("var AugMax = \"0x%s\"\n", max.Text(16)))
	f.WriteString(fmt.Sprintf("var AugRoot = \"0x%s\"\n", augMerkle.Text(16)))
	f.WriteString(fmt.Sprintf("//==================Finish submission======================\n"))
	return nil
}
func (gcc *GethContractClient) VerifyClaim(
	rlpHeader []byte,
	nonce *big.Int,
	shareIndex *big.Int,
	dataSetLookup []*big.Int,
	witnessForLookup []*big.Int,
	augCountersBranch []*big.Int,
	augHashesBranch []*big.Int) error {
	filename := "test_input.js"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("//====================Got verification======================\n"))
	f.WriteString(fmt.Sprintf("var rlp%d = \"%s\";\n", shareIndex.Uint64(), hex.EncodeToString(rlpHeader)))
	f.WriteString(fmt.Sprintf("var nonce%d = \"0x%s\";\n", shareIndex.Uint64(), nonce.Text(16)))
	f.WriteString(fmt.Sprintf("var index%d = \"0x%s\";\n", shareIndex.Uint64(), shareIndex.Text(16)))
	f.WriteString(fmt.Sprintf("var dataSetLookup%d = [", shareIndex.Uint64()))
	for _, i := range dataSetLookup {
		f.WriteString(fmt.Sprintf("\"0x%s\", ", i.Text(16)))
	}
	f.WriteString(fmt.Sprintf("];\n"))
	f.WriteString(fmt.Sprintf("var witnessForLookup%d = [", shareIndex.Uint64()))
	for _, i := range witnessForLookup {
		f.WriteString(fmt.Sprintf("\"0x%s\", ", i.Text(16)))
	}
	f.WriteString(fmt.Sprintf("];\n"))
	f.WriteString(fmt.Sprintf("var augCountersBranch%d = [", shareIndex.Uint64()))
	for _, i := range augCountersBranch {
		f.WriteString(fmt.Sprintf("\"0x%s\", ", i.Text(16)))
	}
	f.WriteString(fmt.Sprintf("];\n"))
	f.WriteString(fmt.Sprintf("var augHashesBranch%d = [", shareIndex.Uint64()))
	for _, i := range augHashesBranch {
		f.WriteString(fmt.Sprintf("\"0x%s\", ", i.Text(16)))
	}
	f.WriteString(fmt.Sprintf("];\n"))
	f.WriteString(fmt.Sprintf("//====================End verification======================\n"))
	return nil
}

func (gcc *GethContractClient) SetEpochData(
	epoch *big.Int,
	fullSizeIn128Resolution *big.Int,
	branchDepth *big.Int,
	merkleNodes []*big.Int,
) error {
	filename := "test_epoch_data_input.js"
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	nodes := []*big.Int{}
	start := big.NewInt(0)
	part := 0
	fmt.Printf("No meaningful nodes: %d\n", len(merkleNodes))
	for k, n := range merkleNodes {
		nodes = append(nodes, n)
		if len(nodes) == 40 || k == len(merkleNodes)-1 {
			mnlen := big.NewInt(int64(len(nodes)))
			// tx, err := cc.pool.SetEpochData(
			// 	cc.transactor, epoch, fullSizeIn128Resolution,
			// 	branchDepth, nodes, start, mnlen)
			f.WriteString(fmt.Sprintf("var epoch%d = \"0x%s\"\n", part, epoch.Text(16)))
			f.WriteString(fmt.Sprintf("var fullSizeIn128Resultion%d = \"0x%s\"\n", part, fullSizeIn128Resolution.Text(16)))
			f.WriteString(fmt.Sprintf("var branchDepth%d = \"0x%s\"\n", part, branchDepth.Text(16)))
			f.WriteString(fmt.Sprintf("var start%d = \"0x%s\"\n", part, start.Text(16)))
			f.WriteString(fmt.Sprintf("var numElems%d = \"0x%s\"\n", part, mnlen.Text(16)))
			f.WriteString(fmt.Sprintf("var merkleNodes%d = [", part))
			for _, no := range nodes {
				f.WriteString(fmt.Sprintf("\"0x%s\", ", no.Text(16)))
			}
			f.WriteString(fmt.Sprintf("]\n"))
			start.Add(start, mnlen)
			nodes = []*big.Int{}
			part++
		}
	}
	return nil
}

func NewGethContractClient() *GethContractClient {
	return &GethContractClient{}
}
