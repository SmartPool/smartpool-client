package geth

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/core/types"
	"math/rand"
	"time"
)

type TxProducer func() (*types.Transaction, error)

func EnsureTx(producer TxProducer, lower, upper int, action string) *types.Transaction {
	var (
		tx  *types.Transaction
		err error
	)
	for {
		tx, err = producer()
		if err != nil {
			waitTime := rand.Int()%(upper-lower) + lower
			smartpool.Output.Printf("%s failed. Error: %s. Retry in %d millisecond\n", action, err, waitTime)
			time.Sleep(time.Duration(waitTime) * time.Millisecond)
		} else {
			break
		}
	}
	return tx
}
