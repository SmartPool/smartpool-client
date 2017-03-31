package ethereum

import (
	"encoding/gob"
)

type FileStorage struct{}

func (*FileStorage) PersistLatestCounter(counter *big.Int) error {

}

func (*FileStorage) LoadLatestCounter() (*big.Int, error) {
}

func (*FileStorage) PersistActiveShares(shares []*Share) error {
}

func (*FileStorage) LoadActiveShares() ([]*Share, error) {
}
