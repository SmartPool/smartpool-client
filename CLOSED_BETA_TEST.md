# Farm setup instructions

## Requirements

In this installation guide, we suppose you already have Ubuntu and a Ethereum full node client
(we support only Ethereum Foundation's Geth and Parity) running.

### Golang compiler
[Golang compiler](https://golang.org/) version 1.7 or higher.
If you don't mind using third party PPA, you can install go1.8 following instructions below:
1. sudo add-apt-repository ppa:gophers/archive
2. sudo apt-get update
3. sudo apt-get install golang-1.8
4. export PATH=$PATH:/usr/lib/go-1.8/bin (You can put this at the end of your ~/.profile and reload it by `source ~/.profile`)

### Install SmartPool
1. git clone https://github.com/SmartPool/smartpool-client
2. cd smartpool-client
3. git checkout develop
4. ./compile.sh (this will create `smartpool` executable commandline tool in your current directory)

### Running SmartPool for your farm
```
./smartpool --keystore ~/Library/Ethereum/testnet/keystore --diff 100000 --gateway 0x79A09eab4Cb39A43115cF34D9DDCD26AD73e03ea --threshold 20 --spcontract 0x9af93376af1ddd22fa2e94fd0a030b3dea96bb96 --gasprice 10
```
Note: SmartPool binds its rpc server at 0.0.0.0:1633.
### Connect your rigs to SmartPool
SmartPool will serve works and accept shares at `0.0.0.0:1633/<rig_or_worker_name>/`. The trailing `/` is a required. `rig_or_worker_name` must be different to `farm` and `rig`. They are reserved for stats entrypoint.
For example:

`ethminer -F 192.168.1.100:1633/rig1/` if you mine with `ethminer`.
`EthDrcMiner64.exe -epool http://192.168.1.100/rig1/` if you mine with `claymore`.

Note: Remember to change `192.168.1.100` to IP of the machine you run `SmartPool` in your network.

### Collect stats to send us
SmartPool is giving stats for the whole farm and its rigs individually.
Farm stat is at: `http://0.0.0.0:1633/farm`
Rig stat is at: `http://0.0.0.0:1633/rig/<rig_or_worker_name>`

Please go to those urls, save them as a file and send us.
