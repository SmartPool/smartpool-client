package ethminer

import (
	"encoding/json"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/bmizerany/pat"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"net/http"
	"os"
)

const jsonrpcVersion = "2.0"

type Server struct {
	Port      uint16
	rpcServer *RPCServer
	server    *http.Server
	output    smartpool.UserOutput
}

type RPCServer struct {
}

func (server *RPCServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (server *RPCServer) response(w http.ResponseWriter, resp interface{}) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(resp)
}

func NewRPCServer(output smartpool.UserOutput, port uint16) *Server {
	mux := pat.New()
	rpcServer := &RPCServer{}
	mux.Post("/:rig/", rpcServer)
	return &Server{port, rpcServer, &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: mux,
	}, output}
}

func (s *Server) Start() {
	if SmartPool == nil {
		panic("SmartPool instance must be initialized first.")
	}
	if SmartPool.Run() {
		go func() {
			<-SmartPool.SubmitterStopped
			os.Exit(1)
		}()
		s.output.Printf("RPC Server is running...\n")
		s.output.Printf("You can start mining now by running ethminer using following command:\n")
		s.output.Printf("--------------------------\n")
		s.output.Printf("ethminer -F localhost:1633\n")
		s.output.Printf("--------------------------\n")
		s.server.ListenAndServe()
	} else {
		s.output.Printf("SmartPool couldn't run. Exit.\n")
	}
}
