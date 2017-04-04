package main

import (
	"fmt"
	"github.com/SmartPool/smartpool-client/ethereum/geth"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func main() {
	gethRPC, err := geth.NewGethRPC(
		"http://localhost:8545",
		"0xc071df9e80d2d13d3f6a7a062a764df4f34c65fd",
		"0x001aDBc838eDe392B5B054A47f8B8c28f2fA9F3F",
		big.NewInt(100000),
		"0xc071df9e80d2d13d3f6a7a062a764df4f34c65fd",
	)
	fmt.Printf("rpc: %v\nerr:%v\n", gethRPC, err)
	client, err := gethRPC.ClientVersion()
	fmt.Printf("client: %v\nerr: %v\n", client, err)
	from := big.NewInt(314075)
	event := geth.VerifyClaimEventTopic
	sender := common.HexToHash("0x001adbc838ede392b5b054a47f8b8c28f2fa9f3f").Big()
	errCode, errInfo := gethRPC.GetLog(
		from, event, sender,
	)
	fmt.Printf("Error code: 0x%s - Error info: 0x%s\n", errCode.Text(16), errInfo.Text(16))
	fmt.Printf("Error message: %s\n", geth.ErrorMsg(errCode, errInfo))
	blockNo, err := gethRPC.BlockNumber()
	fmt.Printf("Last block number: 0x%s. Error %v\n", blockNo.Text(16), err)
}
