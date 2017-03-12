package geth

import (
	"../"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
)

type GethContractClient struct {
	// the contract implementation that holds all underlying
	// communication with Ethereum Contract
	pool       *TestPool
	transactor *bind.TransactOpts
	geth       *ethereum.GethRPC
}

func (cc *GethContractClient) Version() string {
	v, err := cc.pool.Version(nil)
	if err != nil {
		return ""
	}
	return v
}

func (cc *GethContractClient) IsRegistered() bool {
	ok, err := cc.pool.IsRegistered(nil)
	if err != nil {
		return false
	}
	return ok
}

func (cc *GethContractClient) CanRegister() bool {
	ok, err := cc.pool.CanRegister(nil)
	if err != nil {
		return false
	}
	return ok
}

func (cc *GethContractClient) Register(paymentAddress common.Address) error {
	tx, err := cc.pool.Register(cc.transactor, paymentAddress)
	if err != nil {
		return err
	}
	NewTxWatcher(tx, cc.geth).Wait()
	return nil
}

func (cc *GethContractClient) GetClaimSeed() *big.Int {
	seed, err := cc.pool.GetClaimSeed(nil)
	if err != nil {
		return big.NewInt(0)
	}
	return seed
}

func (cc *GethContractClient) SubmitClaim(
	numShares *big.Int,
	difficulty *big.Int,
	min *big.Int,
	max *big.Int,
	augMerkle *big.Int) error {
	tx, err := cc.pool.SubmitClaim(cc.transactor,
		numShares, difficulty, min, max, augMerkle)
	if err != nil {
		return err
	}
	NewTxWatcher(tx, cc.geth).Wait()
	return nil
}

func (cc *GethContractClient) VerifyClaim(
	rlpHeader []byte,
	nonce *big.Int,
	shareIndex *big.Int,
	dataSetLookup []*big.Int,
	witnessForLookup []*big.Int,
	augCountersBranch []*big.Int,
	augHashesBranch []*big.Int) error {
	tx, err := cc.pool.VerifyClaim(cc.transactor,
		rlpHeader, nonce, shareIndex, dataSetLookup,
		witnessForLookup, augCountersBranch, augHashesBranch)
	if err != nil {
		return err
	}
	NewTxWatcher(tx, cc.geth).Wait()
	return nil
}

func getClient(ipc string) (*ethclient.Client, error) {
	return ethclient.Dial(ipc)
}

func NewGethContractClient(
	contractAddr common.Address, rpc *ethereum.GethRPC,
	ipc, keystorePath, passphrase string) (*GethContractClient, error) {
	client, err := getClient(ipc)
	if err != nil {
		// fmt.Printf("Couldn't connect to Geth via IPC file. Error: %s\n", err)
		return nil, err
	}
	pool, err := NewTestPool(contractAddr, client)
	if err != nil {
		// fmt.Printf("Couldn't get SmartPool information from Ethereum Blockchain. Error: %s\n", err)
		return nil, err
	}
	account := GetAccount(keystorePath, passphrase)
	if account == nil {
		// fmt.Printf("Couldn't get any account from key store.\n")
		return nil, err
	}
	keyio, err := os.Open(account.KeyFile())
	if err != nil {
		// fmt.Printf("Failed to open key file: %s\n", err)
		return nil, err
	}
	// fmt.Printf("Unlocking account...")
	auth, err := bind.NewTransactor(keyio, account.PassPhrase())
	if err != nil {
		// fmt.Printf("Failed to create authorized transactor: %s\n", err)
		return nil, err
	}
	// fmt.Printf("Done.\n")
	return &GethContractClient{pool, auth, rpc}, nil
}
