package geth

import (
	"../"
	"../ethash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

type KovanRPC struct {
	*GethRPC
}

func (k KovanRPC) GetWork() *ethereum.Work {
	var h *types.Header
	h = k.GethRPC.GetPendingBlockHeader()
	seedHash, err := ethash.GetSeedHash(uint64(h.Number.Int64()))
	if err != nil {
		panic(err)
	}
	seed := common.BytesToHash(seedHash).Hex()
	return ethereum.NewWork(h, h.HashNoNonce().Hex(), seed, k.ShareDifficulty)
}

// never submit solution to the node because in Kovan, miners can't propose blocks
func (k KovanRPC) SubmitWork(nonce types.BlockNonce, hash, mixDigest common.Hash) bool {
	return false
}

func NewKovanRPC(endpoint, contractAddr, extraData string, diff *big.Int) (*KovanRPC, error) {
	client, err := rpc.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	return &KovanRPC{&GethRPC{client, common.HexToAddress(contractAddr), []byte(extraData), diff}}, nil
}
