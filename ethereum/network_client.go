package ethereum

import (
	"../"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type NetworkClient struct {
	rpc      *GethRPC
	workpool *WorkPool
}

func (nc *NetworkClient) GetWork() smartpool.Work {
	work := nc.rpc.GetWork()
	nc.workpool.AddWork(work)
	return work
}

func (nc *NetworkClient) SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool {
	return nc.rpc.SubmitHashrate(hashrate, id)
}

func (nc *NetworkClient) SubmitSolution(s smartpool.Solution) bool {
	sol := s.(*Solution)
	return nc.rpc.SubmitWork(sol.Nonce, sol.Hash, sol.MixDigest)
}

func NewNetworkClient(rpc *GethRPC, workpool *WorkPool) *NetworkClient {
	return &NetworkClient{
		rpc, workpool,
	}
}
