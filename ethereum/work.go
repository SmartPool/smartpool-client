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
	blockHeader     *types.Header
	powHash         string
	seedHash        string
	shareDifficulty *big.Int
	minerAddress    string
	createdAt       time.Time
}

func (w *Work) ID() string {
	return w.powHash
}

func (w *Work) CreatedAt() time.Time {
	return w.createdAt
}

func (w *Work) AcceptSolution(sol smartpool.Solution) smartpool.Share {
	solution := sol.(*Solution)
	s := &Share{
		blockHeader:     w.blockHeader,
		nonce:           solution.Nonce,
		mixDigest:       solution.MixDigest,
		shareDifficulty: w.ShareDifficulty(),
		minerAddress:    w.minerAddress,
	}
	s.SolutionState = ethash.Instance.SolutionState(s, w.ShareDifficulty())
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

func NewWork(h *types.Header, ph string, sh string, diff *big.Int, miner string) *Work {
	return &Work{h, ph, sh, diff, miner, time.Now()}
}
