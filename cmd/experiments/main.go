package main

import (
	"fmt"
	"github.com/SmartPool/smartpool-client/ethereum/geth"
	"math/big"
	// "net/http"
	"time"
)

func request(rpc *geth.GethRPC, timeout chan bool, shutdown chan bool) {
	for {
		select {
		case <-timeout:
			shutdown <- true
			return
		default:
			rpc.GetWork()
			fmt.Print(".")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	// http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000
	gethRPC, _ := geth.NewGethRPC(
		"http://localhost:8545", "0x4e899e19e31cb6d86aefc0f3d2b2122e613a3f5b",
		"SmartPool-NsjdZFWvUonU0q00000000", big.NewInt(4000000000),
		"0xe034afdcc2ba0441ff215ee9ba0da3e86450108d",
	)
	timeout := make(chan bool, 1)
	shutdown := make(chan bool, 1)
	for i := 0; i < 750; i++ {
		go request(gethRPC, timeout, shutdown)
	}
	time.Sleep(1800 * time.Second)
	fmt.Printf("Shutting down...\n")
	for i := 0; i < 750; i++ {
		timeout <- true
	}
	for i := 0; i < 750; i++ {
		<-shutdown
	}
}
