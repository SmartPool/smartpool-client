package geth

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type MinerAccount struct {
	keyFile    string
	passphrase string
	address    common.Address
}

func (ma MinerAccount) KeyFile() string         { return ma.keyFile }
func (ma MinerAccount) PassPhrase() string      { return ma.passphrase }
func (ma MinerAccount) Address() common.Address { return ma.address }

func GetAddress(keystorePath string, address common.Address) (common.Address, bool, []common.Address) {
	keys := keystore.NewKeyStore(
		keystorePath,
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)
	var acc *accounts.Account
	addresses := []common.Address{}
	for _, a := range keys.Accounts() {
		addresses = append(addresses, a.Address)
	}
	defaultAcc := big.NewInt(0)
	if address.Big().Cmp(defaultAcc) == 0 {
		if len(keys.Accounts()) == 0 {
			return address, false, addresses
		}
		acc = &keys.Accounts()[0]
	} else {
		for _, a := range keys.Accounts() {
			if a.Address == address {
				acc = &a
				break
			}
		}
	}
	if acc == nil {
		return address, false, addresses
	}
	return acc.Address, true, addresses
}

// Get the first account in key store
// Return nil if there's no account
func GetAccount(keystorePath string, address common.Address, passphrase string) *MinerAccount {
	keys := keystore.NewKeyStore(
		keystorePath,
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)
	var acc accounts.Account
	for _, a := range keys.Accounts() {
		if a.Address == address {
			acc = a
			break
		}
	}
	return &MinerAccount{
		acc.URL.Path,
		passphrase,
		acc.Address,
	}
}
