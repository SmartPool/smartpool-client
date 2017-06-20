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
	"math/rand"
	"time"
)

// TxWatcher keeps track of pending transactions
// and acknowledge corresponding channel when a transaction is
// confirmed.
// It also captures event information emitted during the transaction
// and retry with higher gas price when the tx is not confirmed in time
type TxWatcher struct {
	txs        []*types.Transaction
	verifiedTx *types.Transaction
	transactor *bind.TransactOpts
	node       ethereum.RPCClient
	block      *big.Int
	event      *big.Int
	sender     *big.Int
	verChan    chan bool
}

func (tw *TxWatcher) isVerified() bool {
	for _, tx := range tw.txs {
		if tw.node.IsVerified(tx.Hash()) {
			tw.verifiedTx = tx
			smartpool.Output.Printf("tx %s is verified.\n", tx.Hash().Hex())
			return true
		}
	}
	return false
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

func (tw *TxWatcher) lastTx() *types.Transaction {
	return tw.txs[len(tw.txs)-1]
}

func (tw *TxWatcher) newTx(oldTx *types.Transaction) (*types.Transaction, error) {
	newGasPrice := new(big.Int).Set(oldTx.GasPrice())
	// setting new gas price as 25% higher than old gas price
	newGasPrice.Mul(newGasPrice, big.NewInt(125))
	newGasPrice.Div(newGasPrice, big.NewInt(100))
	newTx := types.NewTransaction(
		oldTx.Nonce(), *oldTx.To(), oldTx.Value(),
		oldTx.Gas(), newGasPrice, oldTx.Data())
	return tw.transactor.Signer(types.HomesteadSigner{}, tw.transactor.From, newTx)
}

func (tw *TxWatcher) rebroadcastByPublicNode(signedTx *types.Transaction) error {
	buff := bytes.NewBuffer([]byte{})
	if err := signedTx.EncodeRLP(buff); err != nil {
		return err
	}
	return SendRawTransaction(buff.Bytes())
}

func (tw *TxWatcher) rebroadcast(oldTx, signedTx *types.Transaction) error {
	buff := bytes.NewBuffer([]byte{})
	if err := signedTx.EncodeRLP(buff); err != nil {
		return err
	}
	var (
		hash common.Hash
		err  error
	)
	for {
		hash, err = tw.node.Broadcast(buff.Bytes())
		if err != nil {
			waitTime := rand.Int()%10000 + 1000
			smartpool.Output.Printf("Broadcast error: %s\n", err)
			time.Sleep(time.Duration(waitTime) * time.Millisecond)
		} else {
			break
		}
	}
	smartpool.Output.Printf(
		"Rebroadcast tx: %s by tx: %s with gas price %d...\n",
		oldTx.Hash().Hex(),
		signedTx.Hash().Hex(),
		signedTx.GasPrice().Uint64())
	if hash.Big().Cmp(common.Big0) == 0 {
		return errors.New("Rebroadcast tx got 0 tx hash. This is not supposed to happend.")
	}
	return nil
}

func (tw *TxWatcher) WaitAndRetry() (*big.Int, *big.Int, error) {
	errCode, errInfo, err := tw.Wait()
	if err != nil {
		oldTx := tw.lastTx()
		signedTx, err := tw.newTx(oldTx)
		if err != nil {
			return nil, nil, err
		}
		tw.txs = append(tw.txs, signedTx)
		err = tw.rebroadcast(oldTx, signedTx)
		if err != nil {
			return nil, nil, err
		}
		errCode, errInfo, err = tw.Wait()
		if err != nil {
			err = tw.rebroadcastByPublicNode(signedTx)
			if err != nil {
				smartpool.Output.Printf("Rebroadcast by public node of tx: %s failed. Error: %s\n", signedTx.Hash().Hex(), err)
			}
		}
		errCode, errInfo, err = tw.Wait()
	}
	return errCode, errInfo, err
}

func (tw *TxWatcher) Wait() (*big.Int, *big.Int, error) {
	smartpool.Output.Printf("Waiting for txs: [")
	for _, tx := range tw.txs {
		smartpool.Output.Printf("%s, ", tx.Hash().Hex())
	}
	smartpool.Output.Printf("] to be mined...")
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
	errCode, errInfo := tw.node.GetLog(tw.txs, tw.block, tw.event, tw.sender)
	return errCode, errInfo, nil
}

func GetTxResult(tx *types.Transaction, opts *bind.TransactOpts, node ethereum.RPCClient,
	blockNo *big.Int, event *big.Int, sender *big.Int) (*big.Int, *big.Int, error) {

	txWatcher := NewTxWatcher(tx, opts, node, blockNo, event, sender)
	errCode, errInfo, err := txWatcher.WaitAndRetry()
	if err != nil {
		smartpool.Output.Printf("No tx in: [")
		for _, tx := range txWatcher.txs {
			smartpool.Output.Printf("%s, ", tx.Hash().Hex())
		}
		smartpool.Output.Printf("] was approved by the network in time.\n")
	}
	return errCode, errInfo, err
}

func NewTxWatcher(
	tx *types.Transaction, opts *bind.TransactOpts, node ethereum.RPCClient,
	blockNo *big.Int, event *big.Int, sender *big.Int) *TxWatcher {
	return &TxWatcher{
		[]*types.Transaction{tx}, nil,
		opts, node, blockNo, event, sender, make(chan bool)}
}
