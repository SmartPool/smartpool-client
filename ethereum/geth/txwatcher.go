package geth

import (
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

// TxWatcher keeps track of pending transactions
// and acknowledge corresponding channel when a transaction is
// confirmed.
// It also captures event information emitted during the transaction.
type TxWatcher struct {
	tx      *types.Transaction
	node    ethereum.RPCClient
	block   *big.Int
	event   *big.Int
	sender  *big.Int
	verChan chan bool
}

func (tw *TxWatcher) isVerified() bool {
	return tw.node.IsVerified(tw.tx.Hash())
}

// loop to check transactions verification
// if a transaction is verified, send it to verChan
func (tw *TxWatcher) loop() {
	for {
		if tw.isVerified() {
			tw.verChan <- true
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func (tw *TxWatcher) Wait() (*big.Int, *big.Int) {
	go tw.loop()
	<-tw.verChan
	return tw.node.GetLog(tw.block, tw.event, tw.sender)
}

func NewTxWatcher(
	tx *types.Transaction, node ethereum.RPCClient,
	blockNo *big.Int, event *big.Int, sender *big.Int) *TxWatcher {
	return &TxWatcher{tx, node, blockNo, event, sender, make(chan bool)}
}
