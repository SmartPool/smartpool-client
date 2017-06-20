package geth

import (
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"math/rand"
	"time"
)

func SendRawTransaction(data []byte) error {
	client, err := rpc.DialHTTP("https://mainnet.infura.io/0BRKxQ0SFvAxGL72cbXi")
	// client, err := rpc.DialHTTP("https://ropsten.infura.io/0BRKxQ0SFvAxGL72cbXi")
	if err != nil {
		return err
	}
	for {
		err = client.Call(nil, "eth_sendRawTransaction",
			fmt.Sprintf("0x%s", common.Bytes2Hex(data)))
		if err != nil {
			waitTime := rand.Int()%10000 + 1000
			smartpool.Output.Printf("Failed rebroadcasting via public node. Error: %s\n", err)
			time.Sleep(time.Duration(waitTime) * time.Millisecond)
		} else {
			break
		}
	}
	return nil
}
