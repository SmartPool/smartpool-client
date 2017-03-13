package smartpool

import (
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
}

func (i *Input) RPCEndpoint() string           { return i.rpcEndPoint }
func (i *Input) KeystorePath() string          { return i.keystorePath }
func (i *Input) ShareThreshold() int           { return i.shareThreshold }
func (i *Input) ShareDifficulty() *big.Int     { return i.shareDifficulty }
func (i *Input) SubmitInterval() time.Duration { return i.submitInterval }
func (i *Input) ContractAddress() string       { return i.contractAddr }
func (i *Input) MinerAddress() string          { return i.minerAddr }
func (i *Input) ExtraData() string             { return i.extraData }

func NewInput(
	rpcEndPoint string,
	keystorePath string,
	shareThreshold int,
	shareDifficulty *big.Int,
	submitInterval time.Duration,
	contractAddr string,
	minerAddr string,
	extraData string,
) *Input {
	return &Input{
		rpcEndPoint, keystorePath, shareThreshold, shareDifficulty,
		submitInterval, contractAddr, minerAddr, extraData,
	}
}
