package main

import (
	"fmt"
	"github.com/SmartPool/smartpool-client/ethereum/ethash"
)

func main() {
	ethash.MakeDAG(uint64(3815769), ethash.DefaultDir)
	fmt.Printf("file: %s\n", ethash.PathToDAG(uint64(3815769/30000), ethash.DefaultDir))
}
