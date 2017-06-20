package main

import (
	"fmt"
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/SmartPool/smartpool-client/ethereum/geth"
	"github.com/SmartPool/smartpool-client/storage"
	"math/big"
	// "net/http"
	"time"
)

func request(client *ethereum.NetworkClient, timeout chan bool, shutdown chan bool) {
	for {
		select {
		case <-timeout:
			shutdown <- true
			return
		default:
			client.GetWork()
			fmt.Print(".")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	// http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000
	fileStorage := storage.NewGobFileStorage()
	ethereumWorkPool := ethereum.NewWorkPool(fileStorage)
	gethRPC, _ := geth.NewGethRPC(
		"http://localhost:8545", "0x4e899e19e31cb6d86aefc0f3d2b2122e613a3f5b",
		"SmartPool-NsjdZFWvUonU0q00000000", big.NewInt(100000),
		"0xe034afdcc2ba0441ff215ee9ba0da3e86450108d",
	)
	networkClient := ethereum.NewNetworkClient(gethRPC, ethereumWorkPool)
	timeout := make(chan bool, 1)
	shutdown := make(chan bool, 1)
	for i := 0; i < 750; i++ {
		go request(networkClient, timeout, shutdown)
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
