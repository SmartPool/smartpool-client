package ethereum

import (
	"errors"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
)

type NetworkClient struct {
	rpc      RPCClient
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

func (nc *NetworkClient) ReadyToMine() bool {
	return !nc.rpc.Syncing()
}

func (nc *NetworkClient) Configure(etherbase common.Address, extradata string) error {
	client, err := nc.rpc.ClientVersion()
	if err != nil {
		return err
	}
	if strings.HasPrefix(client, "Geth") {
		smartpool.Output.Printf("Trying to set etherbase to SmartPool contract address: %s...", etherbase.Hex())
		err = nc.rpc.SetEtherbase(etherbase)
		if err != nil {
			smartpool.Output.Printf("Trying to set etherbase to SmartPool contract address failed: %s\n", err)
			smartpool.Output.Printf("Please make sure you used Geth option --rpcapi \"db,eth,net,web3,miner\"\n")
			return err
		}
		smartpool.Output.Printf("Done.\n")
		smartpool.Output.Printf("Trying to set extradata to SmartPool extradata convention: %s...", extradata)
		err = nc.rpc.SetExtradata(extradata)
		if err != nil {
			smartpool.Output.Printf("Trying to set extra data to SmartPool extradata convention failed: %s\n", err)
			smartpool.Output.Printf("Please make sure you used Geth option --rpcapi \"db,eth,net,web3,miner\"\n")
			return err
		}
		smartpool.Output.Printf("Done.\n")
	} else if strings.HasPrefix(client, "Parity") {
		smartpool.Output.Printf("Trying to set etherbase to SmartPool contract address: %s...", etherbase.Hex())
		err = nc.rpc.SetEtherbase(etherbase)
		if err != nil {
			smartpool.Output.Printf("Trying to set author to SmartPool contract address failed: %s\n", err)
			smartpool.Output.Printf("Please make sure you used Parity option --jsonrpc-apis \"web3,eth,net,parity,traces,rpc,parity_set\"\n")
			return err
		}
		smartpool.Output.Printf("Done.\n")
		smartpool.Output.Printf("Trying to set extradata to SmartPool extradata convention: %s...", extradata)
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
	return &NetworkClient{
		rpc, workpool,
	}
}
