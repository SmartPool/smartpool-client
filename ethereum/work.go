// Package ethereum contains all necessary components to plug into smartpool
// to work with ethereum blockchain. Such as: Contract, Network Client,
// Share receiver...
// This package also provides interfaces for different ethereum clients to
// be able to work with smartpool.
package ethereum

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum/ethash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

// Work represents Ethereum pow work
type Work struct {
	BlockHeader     *types.Header
	Hash            string
	SeedHash        string
	ShareDifficulty *big.Int
	MinerAddress    string
	CreatedAt       time.Time
}

func (w *Work) ID() string {
	return w.Hash
}

func (w *Work) AcceptSolution(sol smartpool.Solution) smartpool.Share {
	solution := sol.(*Solution)
	s := &Share{
		blockHeader:     w.BlockHeader,
		nonce:           solution.Nonce,
		mixDigest:       solution.MixDigest,
		shareDifficulty: w.ShareDifficulty,
		minerAddress:    w.MinerAddress,
	}
	s.SolutionState = ethash.Instance.SolutionState(s, w.ShareDifficulty)
	return s
}

func (w *Work) PoWHash() common.Hash {
	return common.HexToHash(w.Hash)
}

func NewWork(h *types.Header, ph string, sh string, diff *big.Int, miner string) *Work {
	return &Work{h, ph, sh, diff, miner, time.Now()}
}
