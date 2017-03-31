package protocol

import (
	"math/big"
)

// PersistentStorage is the gateway for smartpool to interact with external
// persistent storage such as a file system, a database or even a cloud based
// service.
// Smartpool should only persist something via this interface.
type PersistentStorage interface {
	PersistLatestCounter(counter *big.Int) error
	LoadLatestCounter() (*big.Int, error)
}
