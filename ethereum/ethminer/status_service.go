package ethminer

import (
	"encoding/json"
	"math/big"
	"net/http"
)

type StatusService struct{}

type StatusData struct {
	RPCEndpoint    string   `json:"rpc_endpoint"`
	ShareThreshold int      `json:"share_threshold"`
	ShareDiff      *big.Int `json:"share_difficulty"`
	Contract       string   `json:"contract_address"`
	Miner          string   `json:"miner_address"`
	Extra          string   `json:"extra_data"`
	HotStop        bool     `json:"hotstop_mode"`
}

func (server *StatusService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	result := StatusData{
		SmartPool.Input.RPCEndpoint(),
		SmartPool.Input.ShareThreshold(),
		SmartPool.Input.ShareDifficulty(),
		SmartPool.Input.ContractAddress(),
		SmartPool.Input.MinerAddress(),
		SmartPool.Input.ExtraData(),
		SmartPool.Input.HotStop(),
	}
	encoder.Encode(&result)
}

func NewStatusService() *StatusService {
	return &StatusService{}
}
