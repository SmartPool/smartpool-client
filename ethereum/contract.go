package ethereum

import (
	"encoding/hex"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum/ethash"
	"github.com/SmartPool/smartpool-client/mtree"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"path/filepath"
)

type Contract struct {
	client ContractClient
}

func (c *Contract) Version() string {
	return c.client.Version()
}

func (c *Contract) IsRegistered() bool {
	return c.client.IsRegistered()
}

func (c *Contract) CanRegister() bool {
	return c.client.CanRegister()
}

func (c *Contract) Register(paymentAddress common.Address) error {
	return c.client.Register(paymentAddress)
}

func (c *Contract) SubmitClaim(claim smartpool.Claim) error {
	smartpool.Output.Printf("Min: 0x%s - Max: 0x%s - Diff: 0x%s\n", claim.Min().Text(16), claim.Max().Text(16), claim.Difficulty().Text(16))
	return c.client.SubmitClaim(
		claim.NumShares(), claim.Difficulty(),
		claim.Min(), claim.Max(), claim.AugMerkle().Big())
}

func (c *Contract) GetShareIndex(claim smartpool.Claim) *big.Int {
	seed := c.client.GetClaimSeed()
	fmt.Printf("seed mod no.shares: 0x%s mod %s\n", seed.Text(16), claim.NumShares().Text(10))
	index := big.NewInt(0)
	index.Mod(seed, claim.NumShares())
	return index
}

func (c *Contract) VerifyClaim(shareIndex *big.Int, claim smartpool.Claim) error {
	share := claim.GetShare(int(shareIndex.Int64())).(*Share)
	rlpHeader, _ := share.RlpHeaderWithoutNonce()
	nonce := share.NonceBig()
	claim.SetEvidence(shareIndex)
	augCountersBranch := claim.CounterBranch()
	augHashesBranch := claim.HashBranch()
	dataSetLookup := share.DAGElementArray()
	witnessForLookup := share.DAGProofArray()
	return c.client.VerifyClaim(
		rlpHeader,
		nonce,
		shareIndex,
		dataSetLookup,
		witnessForLookup,
		augCountersBranch,
		augHashesBranch,
	)
}

func (c *Contract) SetEpochData(epoch int) error {
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
	err = c.client.SetEpochData(
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

func NewContract(client ContractClient) *Contract {
	return &Contract{client}
}
