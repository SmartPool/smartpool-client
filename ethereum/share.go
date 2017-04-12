package ethereum

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum/ethash"
	"github.com/SmartPool/smartpool-client/mtree"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"io"
	"log"
	"math/big"
	"os"
	"path/filepath"
)

type Share struct {
	blockHeader     *types.Header
	nonce           types.BlockNonce
	mixDigest       common.Hash
	shareDifficulty *big.Int
	minerAddress    string
	SolutionState   int
	dt              *mtree.DagTree
}

func (s *Share) Difficulty() *big.Int      { return s.blockHeader.Difficulty }
func (s *Share) ShareDifficulty() *big.Int { return s.shareDifficulty }
func (s *Share) HashNoNonce() common.Hash  { return s.blockHeader.HashNoNonce() }
func (s *Share) Nonce() uint64             { return s.nonce.Uint64() }
func (s *Share) MixDigest() common.Hash    { return s.mixDigest }
func (s *Share) NumberU64() uint64         { return s.blockHeader.Number.Uint64() }
func (s *Share) MinerAddress() string      { return s.minerAddress }
func (s *Share) NonceBig() *big.Int {
	n := new(big.Int)
	n.SetBytes(s.nonce[:])
	return n
}

func (s *Share) FullSolution() bool {
	return s.SolutionState == 2
}

func (s *Share) BlockHeader() *types.Header {
	return s.blockHeader
}

func (s *Share) RlpHeaderWithoutNonce() ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := rlp.Encode(buffer, []interface{}{
		s.BlockHeader().ParentHash,
		s.BlockHeader().UncleHash,
		s.BlockHeader().Coinbase,
		s.BlockHeader().Root,
		s.BlockHeader().TxHash,
		s.BlockHeader().ReceiptHash,
		s.BlockHeader().Bloom,
		s.BlockHeader().Difficulty,
		s.BlockHeader().Number,
		s.BlockHeader().GasLimit,
		s.BlockHeader().GasUsed,
		s.BlockHeader().Time,
		s.BlockHeader().Extra,
	})
	fmt.Printf("RLP: 0x%s\n", hex.EncodeToString(buffer.Bytes()))
	return buffer.Bytes(), err
}

func (s *Share) Timestamp() *big.Int {
	return s.blockHeader.Time
}

// We use concatenation of timestamp and nonce
// as share counter
// Nonce in ethereum is 8 bytes so counter = timestamp << 64 + nonce
func (s *Share) Counter() *big.Int {
	t := big.NewInt(0)
	t.Set(s.Timestamp())
	t.Lsh(t, 64)
	n := big.NewInt(0).SetBytes(s.nonce[:])
	return t.Add(t, n)
}

func (s *Share) Hash() (result smartpool.SPHash) {
	h := s.blockHeader.HashNoNonce()
	copy(result[:smartpool.HashLength], h[smartpool.HashLength:])
	return
}

func processDuringRead(
	datasetPath string, mt *mtree.DagTree) {

	f, err := os.Open(datasetPath)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	buf := [128]byte{}
	// ignore first 8 bytes magic number at the beginning
	// of dataset. See more at https://gopkg.in/ethereum/wiki/wiki/Ethash-DAG-Disk-Storage-Format
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

func (s *Share) buildDagTree() {
	indices := ethash.Instance.GetVerificationIndices(s)
	fmt.Print("indices: %v\n", indices)
	s.dt = mtree.NewDagTree()
	s.dt.RegisterIndex(indices...)
	fullSize, _ := ethash.MakeDAGWithSize(s.NumberU64(), "")
	fullSizeIn128Resolution := fullSize / 128
	branchDepth := len(fmt.Sprintf("%b", fullSizeIn128Resolution-1))
	s.dt.RegisterStoredLevel(uint32(branchDepth), uint32(10))
	seedHash, err := ethash.GetSeedHash(s.NumberU64())
	if err != nil {
		panic(err)
	}
	path := filepath.Join(
		ethash.DefaultDir,
		fmt.Sprintf("full-R%s-%s", "23", hex.EncodeToString(seedHash[:8])),
	)
	processDuringRead(path, s.dt)
	s.dt.Finalize()
}

func (s *Share) DAGElementArray() []*big.Int {
	if s.dt == nil {
		s.buildDagTree()
	}
	result := []*big.Int{}
	for _, w := range s.dt.AllDAGElements() {
		result = append(result, w.ToUint256Array()...)
	}
	return result
}

func (s *Share) DAGProofArray() []*big.Int {
	if s.dt == nil {
		s.buildDagTree()
	}
	result := []*big.Int{}
	for _, be := range s.dt.AllBranchesArray() {
		result = append(result, be.Big())
	}
	return result
}

func NewShare(h *types.Header, dif *big.Int, miner string) *Share {
	return &Share{
		h,
		types.BlockNonce{},
		common.Hash{},
		dif,
		miner,
		0,
		nil,
	}
}
