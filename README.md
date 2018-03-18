# Golang Lib for Parity's Extended JSON RPC API

This library (aims to) offer a complete and continuously updated json rpc client written in Golang for a Parity node.

## Usage

Create a new client with

    node_address := "localhost:8545" // e.g.
    client := rpc.NewClient(node_address)

then access Ethereum's RPC method _module_function_ (e.g. __eth_getBalance__) by

    client.Module.Function(...params...) // e.g. client.Eth.GetBalance("0x00000000000000000001")


## Methods supported (so far...)

### Web3

- [x] [web3_clientVersion](https://wiki.parity.io/JSONRPC-web3-module#web3_clientversion)
- [x] [web3_sha3](https://wiki.parity.io/JSONRPC-web3-module#web3_sha3)

### Net

- [x] [net_listening](https://wiki.parity.io/JSONRPC-web3-module#net_listening)
- [x] [net_peerCount](https://wiki.parity.io/JSONRPC-web3-module#net_peercount)
- [x] [net_version](https://wiki.parity.io/JSONRPC-web3-module#net_version)

### Eth

- [x] [eth_accounts](https://wiki.parity.io/JSONRPC-eth-module#eth_accounts)
- [x] [eth_blockNumber](https://wiki.parity.io/JSONRPC-eth-module#eth_blocknumber)
- [ ] [eth_call](https://wiki.parity.io/JSONRPC-eth-module#eth_call)
- [x] [eth_coinbase](https://wiki.parity.io/JSONRPC-eth-module#eth_coinbase)
- [ ] [eth_estimateGas](https://wiki.parity.io/JSONRPC-eth-module#eth_estimategas)
- [x] [eth_gasPrice](https://wiki.parity.io/JSONRPC-eth-module#eth_gasprice)
- [x] [eth_getBalance](https://wiki.parity.io/JSONRPC-eth-module#eth_getbalance)
- [x] [eth_getBlockByHash](https://wiki.parity.io/JSONRPC-eth-module#eth_getblockbyhash)
- [x] [eth_getBlockByNumber](https://wiki.parity.io/JSONRPC-eth-module#eth_getblockByNumber)
- [x] [eth_getBlockTransactionCountByHash](https://wiki.parity.io/JSONRPC-eth-module#eth_getblocktransactioncountbyhash)
- [x] [eth_getBlockTransactionCountByNumber](https://wiki.parity.io/JSONRPC-eth-module#eth_getblocktransactioncountbynumber)
- [ ] [eth_getCode](https://wiki.parity.io/JSONRPC-eth-module#eth_getcode)
- [ ] [eth_getFilterChanges](https://wiki.parity.io/JSONRPC-eth-module#eth_getfilterchanges)
- [ ] [eth_getFilterLogs](https://wiki.parity.io/JSONRPC-eth-module#eth_getfilterlogs)
- [ ] [eth_getLogs](https://wiki.parity.io/JSONRPC-eth-module#eth_getlogs)
- [ ] [eth_getStorageAt](https://wiki.parity.io/JSONRPC-eth-module#eth_getstorageat)
- [x] [eth_getTransactionByBlockHashAndIndex](https://wiki.parity.io/JSONRPC-eth-module#eth_gettransactionbyblockhashandindex)
- [x] [eth_getTransactionByBlockNumberAndIndex](https://wiki.parity.io/JSONRPC-eth-module#eth_gettransactionbyblocknumberandindex)
- [x] [eth_getTransactionByHash](https://wiki.parity.io/JSONRPC-eth-module#eth_gettransactionbyhash)
- [x] [eth_getTransactionCount](https://wiki.parity.io/JSONRPC-eth-module#eth_gettransactioncount)
- [x] [eth_getTransactionReceipt](https://wiki.parity.io/JSONRPC-eth-module#eth_getTransactionreceipt)
- [ ] [eth_getUncleByBlockHashAndIndex](https://wiki.parity.io/JSONRPC-eth-module#eth_getunclebyblockhashandindex)
- [ ] [eth_getUncleByBlockNumberAndIndex](https://wiki.parity.io/JSONRPC-eth-module#eth_getunclebyblocknumberandindex)
- [ ] [eth_getUncleCountByBlockHash](https://wiki.parity.io/JSONRPC-eth-module#eth_getunclecountbyblockhash)
- [ ] [eth_getUncleCountByBlockNumber](https://wiki.parity.io/JSONRPC-eth-module#eth_getunclecountbyblocknumber)
- [ ] [eth_getWork](https://wiki.parity.io/JSONRPC-eth-module#eth_getwork)
- [x] [eth_hashrate](https://wiki.parity.io/JSONRPC-eth-module#eth_hashrate)
- [x] [eth_mining](https://wiki.parity.io/JSONRPC-eth-module#eth_mining)
- [ ] [eth_newBlockFilter](https://wiki.parity.io/JSONRPC-eth-module#eth_newblockfilter)
- [ ] [eth_newFilter](https://wiki.parity.io/JSONRPC-eth-module#eth_newfilter)
- [ ] [eth_newPendingTransactionFilter](https://wiki.parity.io/JSONRPC-eth-module#eth_newpendingtransactionfilter)
- [x] [eth_protocolVersion](https://wiki.parity.io/JSONRPC-eth-module#eth_protocolversion)
- [ ] [eth_sendRawTransaction](https://wiki.parity.io/JSONRPC-eth-module#eth_sendrawtransaction)
- [ ] [eth_sendTransaction](https://wiki.parity.io/JSONRPC-eth-module#eth_sendtransaction)
- [ ] [eth_sign](https://wiki.parity.io/JSONRPC-eth-module#eth_sign)
- [ ] [eth_signTransaction](https://wiki.parity.io/JSONRPC-eth-module#eth_signtransaction)
- [ ] [eth_submitHashrate](https://wiki.parity.io/JSONRPC-eth-module#eth_submithashrate)
- [ ] [eth_submitWork](https://wiki.parity.io/JSONRPC-eth-module#eth_submitwork)
- [x] [eth_syncing](https://wiki.parity.io/JSONRPC-eth-module#eth_syncing)
- [x] [eth_uninstallFilter](https://wiki.parity.io/JSONRPC-eth-module#eth_uninstallfilter)

### Personal

- [x] [personal_listAccounts](https://wiki.parity.io/JSONRPC-eth-module#personal_listaccounts)
- [x] [personal_newAccount](https://wiki.parity.io/JSONRPC-eth-module#personal_newaccount)
- [ ] [personal_sendTransaction](https://wiki.parity.io/JSONRPC-eth-module#personal_sendtransaction)
- [ ] [personal_signTransaction](https://wiki.parity.io/JSONRPC-eth-module#personal_signtransaction)
- [ ] [personal_unlockAccount](https://wiki.parity.io/JSONRPC-eth-module#personal_unlockaccount)
- [ ] [personal_sign](https://wiki.parity.io/JSONRPC-eth-module#personal_sign)
- [ ] [personal_ecRecover](https://wiki.parity.io/JSONRPC-eth-module#personal_ecrecover)