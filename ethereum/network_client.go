package ethereum

import (
	"errors"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type NetworkClient struct {
	rpc        RPCClient
	workpool   *WorkPool
	cachedWork *Work
	mu         sync.RWMutex
	ticker     <-chan time.Time
}

func (nc *NetworkClient) fetchNewWork() {
	nc.mu.Lock()
	defer nc.mu.Unlock()
	work := nc.rpc.GetWork()
	nc.cachedWork = work
	nc.workpool.AddWork(work)
}

func (nc *NetworkClient) fetchOnTick() {
	for _ = range nc.ticker {
		nc.fetchNewWork()
	}
}

func (nc *NetworkClient) fetchFromCache() smartpool.Work {
	nc.mu.RLock()
	defer nc.mu.RUnlock()
	return nc.cachedWork
}

func (nc *NetworkClient) GetWork() smartpool.Work {
	for {
		work := nc.fetchFromCache()
		if work != nil {
			return work
		} else {
			waitTime := rand.Int()%100 + 100
			time.Sleep(time.Duration(waitTime) * time.Millisecond)
		}
	}
}

func (nc *NetworkClient) SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool {
	return nc.rpc.SubmitHashrate(hashrate, id)
}

func (nc *NetworkClient) SubmitSolution(s smartpool.Solution) bool {
	sol := s.(*Solution)
	return nc.rpc.SubmitWork(sol.Nonce, sol.Hash, sol.MixDigest)
}

func (nc *NetworkClient) ReadyToMine() bool {
	return !nc.rpc.Syncing()
}

func (nc *NetworkClient) Configure(etherbase common.Address, extradata string) error {
	client, err := nc.rpc.ClientVersion()
	if err != nil {
		return err
	}
	if strings.HasPrefix(client, "Geth") {
		smartpool.Output.Printf("Trying to set etherbase to SmartPool contract address: %s...\n", etherbase.Hex())
		err = nc.rpc.SetEtherbase(etherbase)
		if err != nil {
			smartpool.Output.Printf("Trying to set etherbase to SmartPool contract address failed: %s\n", err)
			smartpool.Output.Printf("Please make sure you used Geth option --rpcapi \"db,eth,net,web3,miner\"\n")
			return err
		}
		smartpool.Output.Printf("Done.\n")
		smartpool.Output.Printf("Trying to set extradata to SmartPool extradata convention: %s...\n", extradata)
		err = nc.rpc.SetExtradata(extradata)
		if err != nil {
			smartpool.Output.Printf("Trying to set extra data to SmartPool extradata convention failed: %s\n", err)
			smartpool.Output.Printf("Please make sure you used Geth option --rpcapi \"db,eth,net,web3,miner\"\n")
			return err
		}
		smartpool.Output.Printf("Done.\n")
	} else if strings.HasPrefix(client, "Parity") {
		smartpool.Output.Printf("Trying to set etherbase to SmartPool contract address: %s...\n", etherbase.Hex())
		err = nc.rpc.SetEtherbase(etherbase)
		if err != nil {
			smartpool.Output.Printf("Trying to set author to SmartPool contract address failed: %s\n", err)
			smartpool.Output.Printf("Please make sure you used Parity option --jsonrpc-apis \"web3,eth,net,parity,traces,rpc,parity_set\"\n")
			return err
		}
		smartpool.Output.Printf("Done.\n")
		smartpool.Output.Printf("Trying to set extradata to SmartPool extradata convention: %s...\n", extradata)
		err = nc.rpc.SetExtradata(extradata)
		if err != nil {
			smartpool.Output.Printf("Trying to set extra data to SmartPool extradata convention failed: %s\n", err)
			smartpool.Output.Printf("Please make sure you used Parity option --jsonrpc-apis \"web3,eth,net,parity,traces,rpc,parity_set\"\n")
			return err
		}
		smartpool.Output.Printf("Done.\n")
	} else {
		return errors.New(
			fmt.Sprintf("Unsupported client: %s", client))
	}
	return nil
}

func NewNetworkClient(rpc RPCClient, workpool *WorkPool) *NetworkClient {
	networkClient := &NetworkClient{
		rpc, workpool, nil, sync.RWMutex{}, time.Tick(50 * time.Millisecond),
	}
	go networkClient.fetchOnTick()
	return networkClient
}
