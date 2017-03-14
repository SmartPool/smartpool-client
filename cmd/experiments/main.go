package main

import (
	"../../ethereum/geth"
	"fmt"
)

func main() {
	gethRPC, err := geth.NewGethRPC(
		"http://localhost:8545",
		"0xc071df9e80d2d13d3f6a7a062a764df4f34c65fd",
		"0x001aDBc838eDe392B5B054A47f8B8c28f2fA9F3F",
	)
	fmt.Printf("rpc: %v\nerr:%v\n", gethRPC, err)
	client, err := gethRPC.ClientVersion()
	fmt.Printf("client: %v\nerr: %v\n", client, err)
}
