package ethminer

import (
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/bmizerany/pat"
	"net/http"
	"os"
)

const jsonrpcVersion = "2.0"

type Server struct {
	Port      uint16
	rpcServer *RPCService
	server    *http.Server
	output    smartpool.UserOutput
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
		s.output.Printf("ethminer -F localhost:1633/:worker_name/\n")
		s.output.Printf("Change :worker_name to whichever name you want.\n")
		s.output.Printf("--------------------------\n")
		s.server.ListenAndServe()
	} else {
		s.output.Printf("SmartPool couldn't run. Exit.\n")
	}
}

func NewServer(output smartpool.UserOutput, port uint16) *Server {
	mux := pat.New()
	rpcService := NewRPCService()
	statService := NewStatService()
	statusService := NewStatusService()
	mux.Post("/:rig/", rpcService)
	mux.Get("/status", statusService)
	mux.Get("/:method/:rig", statService)
	mux.Get("/:method", statService)
	return &Server{port, rpcService, &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: mux,
	}, output}
}
