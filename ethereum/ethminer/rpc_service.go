package ethminer

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"net/http"
)

type RPCService struct{}

func (server *RPCService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		err        Error
		e          error
		res        interface{}
		hashrate   hexutil.Uint64
		hashrateID common.Hash
		nonce      types.BlockNonce
		hash       common.Hash
		mixDigest  common.Hash
	)
	rigName := r.URL.Query().Get(":rig")
	service := NewSmartPoolService(rigName)
	method, rawParams, id, err := extractRPCMsg(r)
	if err != nil {
		res = createErrorResponse(id, err)
		server.response(w, res)
		return
	}
	if method == "eth_getWork" {
		res, e = service.GetWork()
		if e != nil {
			err = &callbackError{e.Error()}
		}
	} else if method == "eth_submitHashrate" {
		hashrate, hashrateID, err = parseHashrateArguments(rawParams)
		if err == nil {
			res = service.SubmitHashrate(hashrate, hashrateID)
		}
	} else if method == "eth_submitWork" {
		nonce, hash, mixDigest, err = parseWorkArguments(rawParams)
		if err == nil {
			res = service.SubmitWork(nonce, hash, mixDigest)
		}
	}
	if err != nil {
		server.response(w, createErrorResponse(id, err))
	} else {
		server.response(w, createResponse(id, res))
	}
	return
}

func (server *RPCService) response(w http.ResponseWriter, resp interface{}) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(resp)
}

func NewRPCService() *RPCService {
	return &RPCService{}
}
