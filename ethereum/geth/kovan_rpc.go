package geth

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/SmartPool/smartpool-client/ethereum/ethash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var TimestampFile = getTimestampFile()
var BlockTime = 6 * time.Second

func getTimestampFile() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return filepath.Join(dir, "smartpool_timestamp.json")
}

type KovanRPC struct {
	*GethRPC
	mu            sync.Mutex
	lastTimestamp *big.Int
	lastTime      time.Time
}

func (k *KovanRPC) GetWork() *ethereum.Work {
	var h *types.Header
	h = k.GethRPC.GetPendingBlockHeader()
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.lastTimestamp != nil {
		if time.Since(k.lastTime) > BlockTime {
			// fmt.Printf("Stuck timestamp detected.")
			// fmt.Printf("Force increase block timestamp from 0x%s ", h.Time.Text(16))
			h.Time.Add(k.lastTimestamp, big.NewInt(1))
			// fmt.Printf("to 0x%s\n", h.Time.Text(16))
			// fmt.Printf("k.lastTimestamp: 0x%s\n", k.lastTimestamp.Text(16))
		} else if h.Time.Cmp(k.lastTimestamp) < 0 {
			// fmt.Printf("--> assign timestamp to lastTimestamp\n")
			h.Time.Add(k.lastTimestamp, big.NewInt(0))
		}
	}
	if k.lastTimestamp == nil || k.lastTimestamp.Cmp(h.Time) < 0 {
		k.lastTimestamp = big.NewInt(0)
		k.lastTimestamp.Set(h.Time)
		timeBytes, _ := k.lastTimestamp.GobEncode()
		err := ioutil.WriteFile(TimestampFile, timeBytes, 0666)
		if err != nil {
			panic(err)
		}
		k.lastTime = time.Now()
		// fmt.Printf("assign k.lastTimestamp: 0x%s. LastTime: %v\n", k.lastTimestamp.Text(16), k.lastTime)
	}
	seedHash, err := ethash.GetSeedHash(uint64(h.Number.Int64()))
	if err != nil {
		panic(err)
	}
	seed := common.BytesToHash(seedHash).Hex()
	return ethereum.NewWork(h, h.HashNoNonce().Hex(), seed, k.ShareDifficulty, k.MinerAddress)
}

// never submit solution to the node because in Kovan, miners can't propose blocks
func (k *KovanRPC) SubmitWork(nonce types.BlockNonce, hash, mixDigest common.Hash) bool {
	return false
}

func NewKovanRPC(endpoint, contractAddr, extraData string, diff *big.Int, miner string) (*KovanRPC, error) {
	client, err := rpc.Dial(endpoint)
	if err != nil {
		return nil, err
	}
	t := big.NewInt(0)
	timeBytes, err := ioutil.ReadFile(TimestampFile)
	if err != nil {
		t = nil
	} else {
		err := t.GobDecode(timeBytes)
		if err != nil {
			smartpool.Output.Printf("Failed to parse timestamp from file. Error: %s. Setting timestamp to nil.\n", err)
			t = nil
		}
	}
	return &KovanRPC{
		&GethRPC{client, common.HexToAddress(contractAddr), []byte(extraData), diff, miner},
		sync.Mutex{},
		t,
		time.Now(),
	}, nil
}
