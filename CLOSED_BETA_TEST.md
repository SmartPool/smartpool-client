## Architecture for a mining farm to use SmartPool
In traditional mining pools, every rig must be connected to the pool and submit shares to the pool directly, as shown in this figure.

[TODO: some pic that I will add later]

When using SmartPool, because the pool is maintained by a smart contract, only a few shares are needed to submit to the pool, so mining rigs won't submit the share to contract directly. Instead, they submit shares to a local server/ machine that acts as the gateway between the mining rigs and the SmartPool contract. The new structure for a farm is shown in this figure
[TODO: some other pic that I will add later]

In SmartPool, the local server will need to install an Ethereum client (either Geth or Parity) and a SmartPool client. The SmartPool client will interact with the mining rigs to collect all the shares. It also submit a few share to the SmartPool contract via the Ethereum client. Details on how to install these components are in the next section.

## Requirements and installations
Currently we only support linux for the local-server of the farm. Supporting other platforms will be in the future.
### Geth or Parity installed
You need to install a full-node client for Ethereum. We support both [Go-ethereum](https://github.com/ethereum/go-ethereum) and [Parity](https://github.com/paritytech/parity).

### Golang compiler
(You can skip this step if you already have Golang/Go installed when you installed Go-ethereum)
[Golang compiler](https://golang.org/) version 1.7 or higher. You can install go1.8 by following instructions as below:

```
sudo add-apt-repository ppa:gophers/archive
sudo apt-get update
sudo apt-get install golang-1.8
export PATH=$PATH:/usr/lib/go-1.8/bin (You can put this at the end of your ~/.profile and reload it by `source ~/.profile`)
```

### Install SmartPool client
Follow below instructions to install our client locally.
```
git clone https://github.com/SmartPool/smartpool-client
cd smartpool-client
git checkout develop
./compile.sh
```
The last command will create `smartpool` executable file in your current directory.

## Running SmartPool for your farm

Before running SmartPool-client on your local server, you need to start the Etherem client and get it to sync with the latest Ethereum blockchain. You also need to enable several apis in your client. Specifically, you need to run: 
`parity--jsonrpc-apis "web3,eth,net,parity,traces,rpc,parity_set"`
if you use Parity, or 
`geth --rpc --rpcapi "db,eth,net,web3,miner,personal"`
if you use Geth. 

At your smartpool-client folder, run
```
./smartpool --keystore ~/Library/Ethereum/testnet/keystore --diff 100000 --gateway <spc-gateway-address> --threshold 1000 --spcontract <spc-contract-address> --gasprice 10
```
All the values for the parameters will be provided by us before you launch. SmartPool0client binds its rpc server address at 0.0.0.0:1633 for the mining rigs to submit shares.

### Connect your rigs to SmartPool-client
SmartPool-cient will serve work and accept shares at `0.0.0.0:1633/<rig_or_worker_name>/`. The trailing `/` is a required. `rig_or_worker_name` must be different to `farm` and `rig`. They are reserved for collecting and showing statistics of the farm and the rigs.

For example:

`ethminer -F 192.168.1.100:1633/rig1/` if you mine with `ethminer`.
`EthDrcMiner64.exe -epool http://192.168.1.100/rig1/` if you mine with `claymore`.

Note: Remember to change `192.168.1.100` to the local IP address of the local server in your network.

### Collect stats to send us
SmartPool-client collects and views stats for the whole farm and each individual rig.
Farm stats is available at: `http://0.0.0.0:1633/farm`
Rig stats is available at: `http://0.0.0.0:1633/rig/<rig_or_worker_name>`

We need the stats to improve our protocol. If you are ok with sharing the farm stats with us, please go to the farm stats url, save the json file and send us at admin@smartpool.io.
