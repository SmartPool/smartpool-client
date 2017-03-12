// Package ethereum contains all necessary components to plug into smartpool
// to work with ethereum blockchain. Such as: Contract, Network Client,
// Share receiver...
// This package also provides interfaces for different ethereum clients to
// be able to work with smartpool.
package ethereum

import (
	"../"
	"./ethash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

const fixedDifficulty = 100000

// Work represents Ethereum pow work
type Work struct {
	blockHeader     *types.Header
	powHash         string
	seedHash        string
	shareDifficulty *big.Int
}

func (w *Work) ID() string {
	return w.powHash
}

func (w *Work) AcceptSolution(sol smartpool.Solution) smartpool.Share {
	solution := sol.(*Solution)
	eth := ethash.New()
	s := &Share{
		blockHeader: w.blockHeader,
		nonce:       solution.Nonce,
		mixDigest:   solution.MixDigest,
	}
	s.SolutionState = eth.SolutionState(s, w.ShareDifficulty())
	return s
}

func (w *Work) PoWHash() common.Hash {
	return common.HexToHash(w.powHash)
}

func (w Work) SeedHash() string {
	return w.seedHash
}

func (w Work) ShareDifficulty() *big.Int {
	return w.shareDifficulty
}

func (w Work) BlockHeader() *types.Header {
	return w.blockHeader
}

func NewWork(h *types.Header, ph string, sh string) *Work {
	return &Work{h, ph, sh, big.NewInt(fixedDifficulty)}
}

// func (w Work) PrintInfo() {
// 	fmt.Printf("Work Pow Hash:      %s\n", w.PoWHash().Hex())
// 	h := w.BlockHeader()
// 	fmt.Printf("Pow Hash of header: %s\n", h.HashNoNonce().Hex())
// 	fmt.Printf("Parent Hash: %s\n", h.ParentHash.Hex())
// 	fmt.Printf("Uncle Hash: %s\n", h.UncleHash.Hex())
// 	fmt.Printf("Coinbase: %s\n", h.Coinbase.Hex())
// 	fmt.Printf("Root: %s\n", h.Root.Hex())
// 	fmt.Printf("TxHash: %s\n", h.TxHash.Hex())
// 	fmt.Printf("ReceiptHash: %s\n", h.ReceiptHash.Hex())
// 	fmt.Printf("Bloom: %v\n", h.Bloom)
// 	fmt.Printf("Difficulty: 0x%s\n", h.Difficulty.Text(16))
// 	fmt.Printf("Number: 0x%s\n", h.Number.Text(16))
// 	fmt.Printf("GasLimit: 0x%s\n", h.GasLimit.Text(16))
// 	fmt.Printf("GasUsed: 0x%s\n", h.GasUsed.Text(16))
// 	fmt.Printf("Time: 0x%s\n", h.Time.Text(16))
// 	fmt.Printf("Extra: %v\n", h.Extra)
// 	fmt.Printf("Extra string --%s--\n", string(h.Extra))
// }
