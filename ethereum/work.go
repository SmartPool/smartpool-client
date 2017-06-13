// Package ethereum contains all necessary components to plug into smartpool
// to work with ethereum blockchain. Such as: Contract, Network Client,
// Share receiver...
// This package also provides interfaces for different ethereum clients to
// be able to work with smartpool.
package ethereum

import (
	"encoding/binary"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"time"
)

var maxUint256 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))

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

func (w *Work) SolutionState(s *Share, shareDiff *big.Int) int {
	hash := s.BlockHeader().HashNoNonce().Bytes()
	seed := make([]byte, 40)
	copy(seed, hash)
	binary.LittleEndian.PutUint64(seed[32:], s.Nonce())

	seed = crypto.Keccak512(seed)
	hashimoto := new(big.Int).SetBytes(crypto.Keccak256(append(seed, s.MixDigest().Bytes()...)))

	blockTarget := new(big.Int).Div(maxUint256, s.BlockHeader().Difficulty)
	shareTarget := new(big.Int).Div(maxUint256, s.ShareDifficulty())

	if hashimoto.Cmp(blockTarget) <= 0 {
		return 2
	}
	if hashimoto.Cmp(shareTarget) <= 0 {
		return 1
	}
	return 0
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
	s.SolutionState = w.SolutionState(s, s.ShareDifficulty())
	return s
}

func (w *Work) PoWHash() common.Hash {
	return common.HexToHash(w.Hash)
}

func NewWork(h *types.Header, ph string, sh string, diff *big.Int, miner string) *Work {
	return &Work{h, ph, sh, diff, miner, time.Now()}
}
