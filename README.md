# Smartpool - The first decentralized mining pool based on smart contract (alpha)
[![Gitter](https://img.shields.io/gitter/room/nwjs/nw.js.svg)](https://gitter.im/SmartPool/Lobby)

## Ropsten testnet
[Smartpool](http://smartpool.io) is [live on Ropsten testnet](https://ropsten.etherscan.io/address/0xf7d93bcb8e4372f46383ecee82f9adf1aa397ba9).
This repository consists of the client software.
The smart contract repository is [here](https://github.com/SmartPool/contracts).

### Requirements
#### OS
The client is currently tested only on Mac OS and Ubuntu.

#### Golang compiler
[Golang compiler](https://golang.org/) version 1.7 or higher.

#### Parity client
[Ethereum Parity client](https://github.com/paritytech/parity/releases) version 1.5.9 or higher.

#### Geth client
[Ethereum Geth client](https://github.com/ethereum/go-ethereum) needs to be compiled from source.

#### Ethminer
We support CPU and GPU mining with [ethminer](https://github.com/ethereum/cpp-ethereum) version 1.2.9 or higher.

#### ETH balance
To run smartpool you must have a Ropsten testnet account with least 0.5 Ether. You can get testnet Ethers from [metamask faucets](https://faucet.metamask.io/) or ping us on our [gitter channel](https://gitter.im/SmartPool/Lobby).

*Note:* To get Ether from metamask faucet, you need to install metamask browser add-on.


#### Installation
1. `git clone https://github.com/SmartPool/smartpool-client.git`
2. `cd smartpool-client`
3. `./compile.sh`

*Note:* If you are on MacOS, there is a issue with Go and XCode 8.3 that might make you see `killed     ./smartpool` error. To fix this issue, please run `build/env.sh go build -o smartpool -ldflags -s cmd/ropsten/ropsten.go` instead of ./compile.sh.
### Running
1. Run Geth on Ropsten testnet: `geth --testnet --fast --rpc --rpcapi "db,eth,net,web3,miner"` or Parity: `parity --chain ropsten --jsonrpc-apis "web3,eth,net,parity,traces,rpc,parity_set"`
2. Run smartpool client `./smartpool --keystore keystore_path --miner account`.
Where
- `keystore_path` is a path to a directory that contains your account key. E.g., `$HOME/.local/share/io.parity.ethereum/keys/kovan/`.
- `account` is the address of your account. E.g., `0x2ba80fe2811f8e0ea5eabf8e07697f7e9f5ae56c`.
- E.g., `./smartpool --keystore ~/Library/Ethereum/testnet/keystore --miner 0xe034afdcc2ba0441ff215ee9ba0da3e86450108d`.
3. Enter your key passphrase.
4. Run `ethminer -F localhost:1633` or `ethminer -G -F localhost:1633` if you mine with your GPU.

## Kovan testnet

[Smartpool](http://smartpool.io) is also [live on Kovan testnet](https://kovan.etherscan.io/address/0x0398ae5a974fe8179b6b0ab9baf4d5f366e932bf).

### Parity client
[Ethereum Parity client](https://github.com/paritytech/parity/releases) version 1.5.9 or higher.

### ETH balance
To run smartpool you must have a testnet Kovan account with least 0.5 Ether. You can get testnet Ethers from [faucets](https://github.com/kovan-testnet/faucet).

### Installation
1. `git clone https://github.com/SmartPool/smartpool-client.git`
2. `cd smartpool-client`
3. `./kovan_compile.sh`

### Running
1. Run Parity on Kovan testnet. `parity --chain kovan --jsonrpc-apis "web3,eth,net,parity,traces,rpc,parity_set"`
2. Run smartpool client `./kovan --keystore keystore_path --miner account --spcontract 0x0398ae5a974FE8179B6B0ab9baF4d5f366E932Bf`.
Where
- `keystore_path` is a path to a directory that contains your account key. E.g., `$HOME/.local/share/io.parity.ethereum/keys/kovan/`.
- `account` is the address of your account. E.g., `0x2ba80fe2811f8e0ea5eabf8e07697f7e9f5ae56c`
3. Enter your key passphrase.
4. Run `ethminer -F localhost:1633` or `ethminer -G -F localhost:1633` if you mine with your GPU.

## Support
Contact us at [gitter](https://gitter.im/SmartPool/Lobby) for support.
