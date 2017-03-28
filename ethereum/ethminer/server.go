package ethminer

import (
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/rpc"
	"net/http"
)

type Server struct {
	Port      uint16
	rpcServer *rpc.Server
	server    *http.Server
	output    smartpool.UserOutput
}

func NewRPCServer(output smartpool.UserOutput, port uint16) *Server {
	rpcServer := rpc.NewServer()
	service := SmartPoolService{}
	rpcServer.RegisterName("eth", service)
	return &Server{port, rpcServer, &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: rpcServer,
	}, output}
}

func (s *Server) Start() {
	if SmartPool == nil {
		panic("SmartPool instance must be initialized first.")
	}
	if SmartPool.Run() {
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
