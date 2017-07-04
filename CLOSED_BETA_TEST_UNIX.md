The instructions below apply for running SmartPool on both Ethereum and Ethereum Classic. 

## Architecture for a mining farm to use SmartPool
In traditional mining pools, every rig must be connected to the pool and submit shares to the pool directly, as shown in this figure.

<img src="https://github.com/SmartPool/smartpool-client/blob/develop/miscs/normalpool-farm.png" width="500" alt="Architecture of a mining farm when working with a centralized pool">

When using SmartPool, because the pool is maintained by a smart contract, only a few shares are needed to submit to the pool, so mining rigs won't submit the share to contract directly. Instead, they submit shares to a local server/ machine that acts as the gateway between the mining rigs and the SmartPool contract. The new structure for a farm is shown in this figure

<img src="https://github.com/SmartPool/smartpool-client/blob/develop/miscs/smartpool-farm.png" width="500" alt="Architecture of a mining farm when working with SmartPool">

In SmartPool, the local server will need to install an Ethereum client (either Geth or Parity) and a SmartPool client. The SmartPool client will interact with the mining rigs to collect all the shares. It also submit a few share to the SmartPool contract via the Ethereum client. Details on how to install these components are in the next section.

## Requirements and installations
The below instructions are for  the the local-server of the farm which runs on linux. For the Windows platform, please check this [link](https://github.com/SmartPool/smartpool-client/blob/develop/CLOSED_BETA_TEST_WINDOWS.md)!
### Geth or Parity installed
You need to install a full-node client for Ethereum. We support both [Go-ethereum](https://github.com/ethereum/go-ethereum) and [Parity](https://github.com/paritytech/parity).

### Golang compiler
(You can skip this step if you already have Golang/Go installed when you installed Go-ethereum)
[Golang compiler](https://golang.org/) version 1.7 or higher. 

### Install SmartPool client

Follow below instructions to install our client locally.
```
git clone https://github.com/SmartPool/smartpool-client
cd smartpool-client
git checkout develop
./compile.sh
```
The last command will create a `smartpool` executable file in your current directory.



## Running SmartPool for your farm

Before running SmartPool-client on your local server, you need to start the Etherem client and get it to sync with the latest Ethereum blockchain. You also need to enable several apis in your client. Specifically, you need to run: 
`parity --jsonrpc-apis "web3,eth,net,parity,traces,rpc,parity_set"`
if you use Parity, or 
`geth --rpc --rpcapi "db,eth,net,web3,miner,personal"`
if you use Geth. 

At your smartpool-client folder, run
```
./smartpool --keystore <keystore_dir>  --gateway 0x09077D92F7a71Ae3C4EAc8dC9f35cE9aa5A06F7B --share-threshold 36000 --claim-threshold 2 --spcontract 0xfc668AE14b0F7702c04b105448fE733D96C558DF --gasprice 4
```
The `<keystore_dir>` is normally `~/Library/Ethereum/testnet/keystore`, but it may be different depending on machines. You can figure out the meaning of these parameters in our client by running `./smartpool --help`. Two important parameters are `share-threshold` and `claim-threshold`. You can reduce these parameters to get your payment faster, but the average costs for transaction fees will be higher. The above setup `-share-threshold 36000 --claim-threshold 2` is for a miner of 20 GH/s to get paid every 4 hours (the average fees is ~0.6-1%). You can reduce either of the parameters to get paid faster (linearly) according to your hashrate. SmartPool client binds its rpc server address at 0.0.0.0:1633 for the mining rigs to submit shares. Now you can just let the SmartPool client runs, do not close it.


### Connect your rigs to SmartPool-client
SmartPool-cient will serve work and accept shares at `0.0.0.0:1633/<rig_or_worker_name>/`. The trailing `/` is a required. `rig_or_worker_name` must be different to `farm`, `rig` and `status`. They are reserved for collecting and showing statistics of the farm and the rigs.

For example:

`ethminer -F 192.168.1.100:1633/rig1/` if you mine with `ethminer`.
`EthDrcMiner64.exe -epool http://192.168.1.100/rig1/` if you mine with `claymore`.

Note: Remember to change `192.168.1.100` to the local IP address of the local server in your network.

### Collect stats to send us
SmartPool-client collects and views stats for the whole farm and each individual rig.
Farm stats is available at: `http://0.0.0.0:1633/farm`
Rig stats is available at: `http://0.0.0.0:1633/rig/<rig_or_worker_name>`

We need the stats to improve our protocol. If you are ok with sharing the farm stats with us, please go to the farm stats url, save the json file and send us at admin@smartpool.io.
