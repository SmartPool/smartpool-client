package geth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

func SendRawTransaction(data []byte) error {
	client, err := rpc.DialHTTP("https://mainnet.infura.io/0BRKxQ0SFvAxGL72cbXi")
	// client, err := rpc.DialHTTP("https://ropsten.infura.io/0BRKxQ0SFvAxGL72cbXi")
	if err != nil {
		return err
	}
	return client.Call(nil, "eth_sendRawTransaction",
		fmt.Sprintf("0x%s", common.Bytes2Hex(data)))
}
