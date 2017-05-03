package ethminer

import (
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/SmartPool/smartpool-client/protocol"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

var SmartPool *protocol.SmartPool

type SmartPoolService struct {
	rig *ethereum.Rig
}

func (sps *SmartPoolService) GetWork() ([3]string, error) {
	var res [3]string
	w := SmartPool.GetWork(sps.rig).(*ethereum.Work)
	res[0] = w.PoWHash().Hex()
	res[1] = w.SeedHash()
	n := big.NewInt(1)
	n.Lsh(n, 255)
	n.Div(n, w.ShareDifficulty())
	n.Lsh(n, 1)
	res[2] = common.BytesToHash(n.Bytes()).Hex()
	return res, nil
}

func (sps *SmartPoolService) SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool {
	// nc := SmartPool.NetworkClient.(*ethereum.NetworkClient)
	// return nc.SubmitHashrate(sps.rig, hashrate, id)
	return SmartPool.SubmitHashrate(sps.rig, hashrate, id)
}

func (sps *SmartPoolService) SubmitWork(nonce types.BlockNonce, hash, mixDigest common.Hash) bool {
	sol := &ethereum.Solution{
		Nonce:     nonce,
		Hash:      hash,
		MixDigest: mixDigest,
	}
	return SmartPool.AcceptSolution(sps.rig, sol)
}

func NewSmartPoolService(rigName string) *SmartPoolService {
	return &SmartPoolService{ethereum.NewRig(rigName)}
}
