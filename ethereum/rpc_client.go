package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type RPCClient interface {
	ClientVersion() (string, error)
	GetWork() *Work
	SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool
	SubmitWork(nonce types.BlockNonce, hash, mixDigest common.Hash) bool
	IsVerified(h common.Hash) bool
	Syncing() bool
	BlockNumber() (*big.Int, error)
	GetLog(from *big.Int, event *big.Int, sender *big.Int) (*big.Int, *big.Int)
	SetEtherbase(etherbase common.Address) error
	SetExtradata(extradata string) error
	Broadcast(raw []byte) (common.Hash, error)
}
