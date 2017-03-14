package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type RPCClient interface {
	ClientVersion() (string, error)
	GetWork() *Work
	SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool
	SubmitWork(nonce types.BlockNonce, hash, mixDigest common.Hash) bool
	IsVerified(h common.Hash) bool
}
