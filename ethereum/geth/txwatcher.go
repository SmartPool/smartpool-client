package geth

import (
	"bytes"
	"errors"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/ethereum/go-ethereum/common"
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

func (tw *TxWatcher) WaitAndRetry() (*big.Int, *big.Int, error) {
	errCode, errInfo, err := tw.Wait()
	if err != nil {
		smartpool.Output.Printf("Rebroadcast tx: %s...\n", tw.tx.Hash().Hex())
		buff := bytes.NewBuffer([]byte{})
		err = tw.tx.EncodeRLP(buff)
		if err != nil {
			return nil, nil, err
		}
		hash, err := tw.node.Broadcast(buff.Bytes())
		if err != nil {
			smartpool.Output.Printf("Broadcast error: %s\n", err)
			return nil, nil, err
		}
		if hash.Big().Cmp(common.Big0) == 0 {
			return nil, nil, errors.New("Rebroadcast tx got 0 tx hash. This is not supposed to happend.")
		}
		return tw.Wait()
	} else {
		return errCode, errInfo, err
	}
}

func (tw *TxWatcher) Wait() (*big.Int, *big.Int, error) {
	timeout := make(chan bool, 1)
	go tw.loop()
	go func() {
		time.Sleep(10 * time.Minute)
		timeout <- true
	}()
	select {
	case <-tw.verChan:
		break
	case <-timeout:
		return nil, nil, errors.New("timeout error")
	}
	errCode, errInfo := tw.node.GetLog(tw.block, tw.event, tw.sender)
	return errCode, errInfo, nil
}

func NewTxWatcher(
	tx *types.Transaction, node ethereum.RPCClient,
	blockNo *big.Int, event *big.Int, sender *big.Int) *TxWatcher {
	return &TxWatcher{tx, node, blockNo, event, sender, make(chan bool)}
}
