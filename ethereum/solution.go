package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Solution struct {
	Nonce     types.BlockNonce
	Hash      common.Hash
	MixDigest common.Hash
}

func (s *Solution) WorkID() string {
	return s.Hash.Hex()
}
