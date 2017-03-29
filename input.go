package smartpool

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

type Input struct {
	rpcEndPoint     string
	keystorePath    string
	shareThreshold  int
	shareDifficulty *big.Int
	submitInterval  time.Duration
	contractAddr    string
	minerAddr       string
	extraData       string
	hotStop         bool
}

func (i *Input) RPCEndpoint() string           { return i.rpcEndPoint }
func (i *Input) KeystorePath() string          { return i.keystorePath }
func (i *Input) ShareThreshold() int           { return i.shareThreshold }
func (i *Input) ShareDifficulty() *big.Int     { return i.shareDifficulty }
func (i *Input) SubmitInterval() time.Duration { return i.submitInterval }
func (i *Input) ContractAddress() string       { return i.contractAddr }
func (i *Input) MinerAddress() string          { return i.minerAddr }
func (i *Input) ExtraData() string             { return i.extraData }
func (i *Input) HotStop() bool                 { return i.hotStop }
func (i *Input) SetMinerAddress(addr common.Address) {
	i.minerAddr = addr.Hex()
}
func (i *Input) SetExtraData(extra string) {
	i.extraData = extra
}
func (i *Input) SetContractAddress(addr common.Address) {
	i.contractAddr = addr.Hex()
}

func NewInput(
	rpcEndPoint string,
	keystorePath string,
	shareThreshold int,
	shareDifficulty *big.Int,
	submitInterval time.Duration,
	contractAddr string,
	minerAddr string,
	extraData string,
	hotStop bool,
) *Input {
	return &Input{
		rpcEndPoint, keystorePath, shareThreshold, shareDifficulty,
		submitInterval, contractAddr, minerAddr, extraData, hotStop,
	}
}
