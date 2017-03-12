package geth

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type MinerAccount struct {
	keyFile    string
	passphrase string
}

func (ma MinerAccount) KeyFile() string    { return ma.keyFile }
func (ma MinerAccount) PassPhrase() string { return ma.passphrase }

// Get the first account in key store
// Return nil if there's no account
func GetAccount(keystorePath, passphrase string) *MinerAccount {
	keys := keystore.NewKeyStore(
		keystorePath,
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)
	if len(keys.Accounts()) == 0 {
		return nil
	}
	acc := keys.Accounts()[0]
	keyFile := acc.URL.Path
	return &MinerAccount{
		keyFile,
		passphrase,
	}
}
