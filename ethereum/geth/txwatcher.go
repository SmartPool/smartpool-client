package geth

import (
	"bytes"
	"errors"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

// TxWatcher keeps track of pending transactions
// and acknowledge corresponding channel when a transaction is
// confirmed.
// It also captures event information emitted during the transaction
// and retry with higher gas price when the tx is not confirmed in time
type TxWatcher struct {
	tx         *types.Transaction
	transactor *bind.TransactOpts
	node       ethereum.RPCClient
	block      *big.Int
	event      *big.Int
	sender     *big.Int
	verChan    chan bool
}

func (tw *TxWatcher) isVerified() bool {
	return tw.node.IsVerified(tw.tx.Hash())
}

// loop to check transactions verification
// if a transaction is verified, send it to verChan
func (tw *TxWatcher) loop(timeout chan bool) {
	for {
		if tw.isVerified() {
			tw.verChan <- true
			return
		}
		select {
		case <-timeout:
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func (tw *TxWatcher) WaitAndRetry() (*big.Int, *big.Int, error) {
	errCode, errInfo, err := tw.Wait()
	if err != nil {
		newGasPrice := new(big.Int).Set(tw.tx.GasPrice())
		// setting new gas price as 25% higher than old gas price
		newGasPrice.Mul(newGasPrice, big.NewInt(125))
		newGasPrice.Div(newGasPrice, big.NewInt(100))
		oldTx := tw.tx
		tw.tx = types.NewTransaction(
			tw.tx.Nonce(), *tw.tx.To(), tw.tx.Value(),
			tw.tx.Gas(), newGasPrice, tw.tx.Data())
		signedTx, err := tw.transactor.Signer(types.HomesteadSigner{}, tw.transactor.From, tw.tx)
		if err != nil {
			return nil, nil, err
		}
		tw.tx = signedTx
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
		smartpool.Output.Printf(
			"Rebroadcast tx: %s by tx: %s with gas price %d...\n",
			oldTx.Hash().Hex(),
			tw.tx.Hash().Hex(),
			newGasPrice.Uint64())
		if hash.Big().Cmp(common.Big0) == 0 {
			return nil, nil, errors.New("Rebroadcast tx got 0 tx hash. This is not supposed to happend.")
		}
		return tw.Wait()
	} else {
		return errCode, errInfo, err
	}
}

func (tw *TxWatcher) Wait() (*big.Int, *big.Int, error) {
	smartpool.Output.Printf("Waiting for tx: %s to be mined...", tw.tx.Hash().Hex())
	timeout := make(chan bool, 1)
	go tw.loop(timeout)
	go func() {
		time.Sleep(10 * time.Minute)
		// push 2 timeouts for the above loop and below select
		timeout <- true
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

func GetTxResult(tx *types.Transaction, opts *bind.TransactOpts, node ethereum.RPCClient,
	blockNo *big.Int, event *big.Int, sender *big.Int) (*big.Int, *big.Int, error) {

	txWatcher := NewTxWatcher(tx, opts, node, blockNo, event, sender)
	errCode, errInfo, err := txWatcher.WaitAndRetry()
	if err != nil {
		smartpool.Output.Printf("Tx: %s was not approved by the network in time.\n", txWatcher.tx.Hash().Hex())
	}
	return errCode, errInfo, err
}

func NewTxWatcher(
	tx *types.Transaction, opts *bind.TransactOpts, node ethereum.RPCClient,
	blockNo *big.Int, event *big.Int, sender *big.Int) *TxWatcher {
	return &TxWatcher{tx, opts, node, blockNo, event, sender, make(chan bool)}
}
