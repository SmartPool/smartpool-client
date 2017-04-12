package geth

import (
	"encoding/hex"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
	"strings"
	"time"
)

type jsonHeader struct {
	ParentHash  *common.Hash      `json:"parentHash"`
	UncleHash   *common.Hash      `json:"sha3Uncles"`
	Coinbase    *common.Address   `json:"miner"`
	Root        *common.Hash      `json:"stateRoot"`
	TxHash      *common.Hash      `json:"transactionsRoot"`
	ReceiptHash *common.Hash      `json:"receiptsRoot"`
	Bloom       *types.Bloom      `json:"logsBloom"`
	Difficulty  *hexutil.Big      `json:"difficulty"`
	Number      *hexutil.Big      `json:"number"`
	GasLimit    *hexutil.Big      `json:"gasLimit"`
	GasUsed     *hexutil.Big      `json:"gasUsed"`
	Time        *hexutil.Big      `json:"timestamp"`
	Extra       *hexutil.Bytes    `json:"extraData"`
	MixDigest   *common.Hash      `json:"mixHash"`
	Nonce       *types.BlockNonce `json:"nonce"`
}

type GethRPC struct {
	client          *rpc.Client
	ContractAddr    common.Address
	ExtraData       []byte
	ShareDifficulty *big.Int
	MinerAddress    string
}

func (g *GethRPC) ClientVersion() (string, error) {
	result := ""
	err := g.client.Call(&result, "web3_clientVersion")
	return result, err
}

func (g *GethRPC) BlockNumber() (*big.Int, error) {
	str := ""
	err := g.client.Call(&str, "eth_blockNumber")
	result := common.HexToHash(str).Big()
	return result, err
}

func (g *GethRPC) GetPendingBlockHeader() *types.Header {
	header := jsonHeader{}
	err := g.client.Call(&header, "eth_getBlockByNumber", "pending", false)
	if err != nil {
		return nil
	}
	result := types.Header{}
	result.ParentHash = *header.ParentHash
	result.UncleHash = *header.UncleHash
	result.Root = *header.Root
	result.TxHash = *header.TxHash
	result.ReceiptHash = *header.ReceiptHash
	result.Difficulty = (*big.Int)(header.Difficulty)
	result.Number = (*big.Int)(header.Number)
	result.GasLimit = (*big.Int)(header.GasLimit)
	result.GasUsed = (*big.Int)(header.GasUsed)
	result.Time = (*big.Int)(header.Time)
	result.Coinbase = g.ContractAddr
	// result.Extra = []byte("0xd883010505846765746887676f312e372e348664617277696e")
	result.Extra = []byte(g.ExtraData)
	if header.Bloom == nil {
		result.Bloom = types.Bloom{}
	} else {
		result.Bloom = *header.Bloom
	}
	result.MixDigest = common.Hash{}
	result.Nonce = types.BlockNonce{}
	return &result
}

func (g *GethRPC) GetBlockHeader(number int) *types.Header {
	header := types.Header{}
	err := g.client.Call(&header, "eth_getBlockByNumber", number, false)
	if err != nil {
		log.Fatal("Couldn't get latest block:", err)
		return nil
	}
	return &header
}

type gethWork [3]string

func (w gethWork) PoWHash() string { return w[0] }

func (g *GethRPC) GetWork() *ethereum.Work {
	w := gethWork{}
	var h *types.Header
	for {
		h = g.GetPendingBlockHeader()
		g.client.Call(&w, "eth_getWork")
		if w.PoWHash() != "" && w.PoWHash() == h.HashNoNonce().Hex() {
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return ethereum.NewWork(h, w[0], w[1], g.ShareDifficulty, g.MinerAddress)
}

func (g *GethRPC) SubmitHashrate(hashrate hexutil.Uint64, id common.Hash) bool {
	var result bool
	g.client.Call(&result, "eth_submitHashrate", hashrate, id)
	return result
}

func (g *GethRPC) SubmitWork(nonce types.BlockNonce, hash, mixDigest common.Hash) bool {
	var result bool
	g.client.Call(&result, "eth_submitWork", nonce, hash, mixDigest)
	return result
}

type filter struct {
	FromBlock string   `json:"fromBlock,omitempty"`
	ToBlock   string   `json:"toBlock,omitempty"`
	Topics    []string `json:"topics,omitempty"`
}

type logs []struct {
	LogIndex         string   `json:"logIndex,omitempty"`
	BlockNumber      string   `json:"blockNumber,omitempty"`
	BlockHash        string   `json:"blockHash,omitempty"`
	TransactionHash  string   `json:"transactionHash,omitempty"`
	TransactionIndex string   `json:"transactionIndex,omitempty"`
	Address          string   `json:"address,omitempty"`
	Data             string   `json:"data,omitempty"`
	Topics           []string `json:"topics,omitempty"`
}

func (g *GethRPC) GetLog(
	from *big.Int, event *big.Int,
	sender *big.Int) (*big.Int, *big.Int) {
	param := filter{
		fmt.Sprintf("0x%s", from.Text(16)),
		"latest",
		[]string{
			common.BigToHash(event).Hex(),
			common.BigToHash(sender).Hex(),
		},
	}
	result := logs{}
	g.client.Call(&result, "eth_getLogs", param)
	if len(result) != 1 {
		smartpool.Output.Printf(
			"Got %d logs. Contract unexpectedly threw. Topic: %s, sender: %s, from: %s\n",
			len(result),
			common.BigToHash(event).Hex(),
			common.BigToHash(sender).Hex(),
			fmt.Sprintf("0x%s", from.Text(16)),
		)
		return nil, nil
	} else {
		theLog := result[0]
		dataInByte, err := hex.DecodeString(theLog.Data[2:])
		if err != nil {
			smartpool.Output.Printf(
				"Error while converting log data to bytes. Log(%s), Error(%v)\n",
				theLog.Data, err,
			)
		}
		errCode := big.NewInt(0)
		errCode.SetBytes(dataInByte[:32])
		errInfo := big.NewInt(0)
		errInfo.SetBytes(dataInByte[32:])
		return errCode, errInfo
	}
}

type jsonTransaction struct {
	BlockHash string `json:"blockHash"`
}

func (g *GethRPC) IsVerified(h common.Hash) bool {
	result := jsonTransaction{}
	g.client.Call(&result, "eth_getTransactionByHash", h)
	return result.BlockHash != "" && result.BlockHash != "0x0000000000000000000000000000000000000000000000000000000000000000"
}

func (g *GethRPC) Syncing() bool {
	result := ""
	g.client.Call(&result, "net_peerCount")
	peerCount := common.HexToHash(result).Big().Uint64()
	fmt.Printf("peerCount: %d\n", peerCount)
	return peerCount == uint64(0)
}

func (g *GethRPC) SetEtherbase(etherbase common.Address) error {
	client, err := g.ClientVersion()
	if err != nil {
		return err
	}
	result := false
	if strings.HasPrefix(client, "Geth") {
		err = g.client.Call(&result, "miner_setEtherbase", etherbase)
	} else {
		// Client must be Parity
		err = g.client.Call(&result, "parity_setAuthor", etherbase)
	}
	return err
}

func (g *GethRPC) SetExtradata(extradata string) error {
	client, err := g.ClientVersion()
	if err != nil {
		return err
	}
	result := false
	if strings.HasPrefix(client, "Geth") {
		err = g.client.Call(&result, "miner_setExtra", extradata)
	} else {
		// Client must be Parity
		err = g.client.Call(&result, "parity_setExtraData",
			common.StringToHash(extradata))
	}
	return err
}

func (g *GethRPC) Broadcast(data []byte) (common.Hash, error) {
	hash := common.Hash{}
	err := g.client.Call(&hash, "eth_sendRawTransaction",
		fmt.Sprintf("0x%s", common.Bytes2Hex(data)))
	return hash, err
}

func NewGethRPC(endpoint, contractAddr, extraData string, diff *big.Int, miner string) (*GethRPC, error) {
	client, err := rpc.DialHTTP(endpoint)
	if err != nil {
		return nil, err
	}
	return &GethRPC{client, common.HexToAddress(contractAddr), []byte(extraData), diff, miner}, nil
}
