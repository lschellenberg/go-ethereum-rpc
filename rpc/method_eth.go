package rpc

import (
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
)

const (
	MethodEthAccounts                         = "eth_accounts"
	MethodEthBlockNumber                      = "eth_blockNumber"
	MethodEthCall                             = "eth_call"
	MethodCoinbase                            = "eth_coinbase"
	MethodGas                                 = "eth_estimate_gas"
	MethodGasPrice                            = "eth_gasPrice"
	MethodGetBalance                          = "eth_getBalance"
	MethodGetBlockByHash                      = "eth_getBlockByHash"
	MethodGetBlockByNumber                    = "eth_getBlockByNumber"
	MethodGetBlockTransactionCountByHash      = "eth_getBlockTransactionCountByHash"
	MethodGetBlockTransactionCountByNumber    = "eth_getBlockTransactionCountByNumber"
	MethodGetCode                             = "eth_getCode"
	MethodGetFilterChanges                    = "eth_getFilterChanges"
	MethodGetFilterLogs                       = "eth_getFilterLogs"
	MethodGetLogs                             = "eth_getLogs"
	MethodGetStorageAt                        = "eth_getStorageAt"
	MethodGetTransactionByBlockNumberAndIndex = "eth_getTransactionByBlockNumberAndIndex"
	MethodGetTransactionByBlockHashAndIndex   = "eth_getTransactionByBlockHashAndIndex"
	MethodGetTransactionByHash                = "eth_getTransactionByHash"
	MethodGetTransactionCount                 = "eth_getTransactionCount"
	MethodGetTransactionReceipt               = "eth_getTransactionReceipt"
	MethodGetUncleByBlockHashAndIndex         = "eth_getUncleByBlockHashAndIndex"
	MethodGetUncleByBlockNumberAndIndex       = "eth_getUncleByBlockNumberAndIndex"
	MethodGetUncleCountByBlockHash            = "eth_getUncleCountByBlockHash"
	MethodGetUncleCountByBlockNumber          = "eth_getUncleCountByBlockNumber"
	MethodGetWork                             = "eth_getWork"
	MethodHashrate                            = "eth_hashrate"
	MethodMining                              = "eth_mining"
	MethodNewBlockFilter                      = "eth_newBlockFilter"
	MethodNewFilter                           = "eth_newFilter"
	MethodNewPendingTransactionFilter         = "eth_newPendingTransactionFilter"
	MethodProtocolVersion                     = "eth_protocolVersion"
	MethodSendRawTransaction                  = "eth_sendRawTransaction"
	MethodSendTransaction                     = "eth_sendTransaction"
	MethodSign                                = "eth_sign"
	MethodSignTransaction                     = "eth_signTransaction"
	MethodSubmitHashrate                      = "eth_submitHashrate"
	MethodSubmitWork                          = "eth_submitWork"
	MethodSyncing                             = "eth_syncing"
	MethodUninstallFilter                     = "eth_uninstallFilter"
)

type Eth struct {
	client *Client
}

/*
	rpc method: "eth_accounts"
	returns []string, a list of addresses owned by eth.client.
 */
func (eth Eth) Accounts() ([]string, error) {
	response, err := checkRPCError(eth.client.Call(MethodEthAccounts))

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.GetStringList()
}

/*
	rpc method: "eth_blockNumber"
	returns int64, the number of most recent block.
	curl --data '{"method":"eth_blockNumber","params":[],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
 */
func (eth Eth) BlockNumber() (int64, error) {
	return eth.client.RequestInt64(MethodEthBlockNumber)
}

/*
	rpc method: "eth_call"
	Executes a new message call immediately without creating a transaction on the block chain.
	returns hex, the return value of the function call
	curl --data '{"method":"eth_call","params":[{"to":"0xd780ae2bf04cd96e577d3d014762f831d97129d0","value":"0x115976c4"}],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545

*/
func (eth Eth) Call(callParams *EthCallParams, quantity *rpctypes.Quantity) (*rpctypes.HexString, error) {
	if quantity == nil {
		quantity = rpctypes.QuantityLatest()
	}

	return eth.client.RequestHexString(MethodEthCall, callParams.ToMap(), quantity.HexStringOrTag())
}

/*
	rpc method: "eth_coinbase"
	returns string, the client coinbase address.

	curl --data '{"method":"eth_coinbase","params":[],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
*/
func (eth Eth) Coinbase() (*rpctypes.HexString, error) {
	return eth.client.RequestHexString(MethodCoinbase)
}

/*
	rpc method: "eth_estimateGas"
	Makes a call or transaction, which won’t be added to the blockchain and returns the used gas, which can be used for estimating the used gas.
    TODO
 */
func (eth Eth) EstimateGas(params *EthEstimateGasParams) (int64, error) {
	return -1, NotImplemented
}

/*
	rpc method: "eth_gasPrice"
	Returns the current price per gas in wei.
 */
func (eth Eth) GasPrice() (*rpctypes.EtherValue, error) {
	return eth.client.RequestEtherValue(MethodGasPrice)
}

/*
	rpc method: "eth_getBalance"
	Returns the balance of the account of given address.
	curl --data '{"method":"eth_getBalance","params":["0x407d73d8a49eeb85d32cf465507dd71d507100c1"],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
*/
func (eth Eth) GetBalance(address string, quantity *rpctypes.Quantity) (*rpctypes.EtherValue, error) {
	return eth.client.RequestEtherValue(MethodGetBalance, address, quantity.HexStringOrTag())
}

/*
	rpc method: "eth_getBlockByHash"
	Returns information about a block by hash.
 */
func (eth Eth) GetBlockByHash(address string, full bool) (*rpctypes.EtherBlock, error) {
	return eth.client.RequestEtherBlock(MethodGetBlockByHash, address, full)
}

/*
	rpc method: "eth_getBlockByHash"
	Returns information about a block by block number.
 */
func (eth Eth) GetBlockByNumber(blockNumber int64, full bool) (*rpctypes.EtherBlock, error) {
	return eth.client.RequestEtherBlock(MethodGetBlockByNumber, new(rpctypes.HexString).FromInt64(blockNumber).String(), full)
}

/*
	rpc method: "eth_getBlockTransactionCountByHash"
	Returns the number of transactions in a block from a block matching the given block hash.
 */
func (eth Eth) GetBlockTransactionCountByHash(address string) (int64, error) {
	return eth.client.RequestInt64(MethodGetBlockTransactionCountByHash, address)
}

/*
	rpc method: "eth_getBlockTransactionCountByNumber"
	Returns the number of transactions in a block from a block matching the given block number.
 */
func (eth Eth) GetBlockTransactionCountByNumber(blockNumber int64) (int64, error) {
	return eth.client.RequestInt64(MethodGetBlockTransactionCountByNumber, new(rpctypes.HexString).FromInt64(blockNumber).String())
}

/*
	TODO
	eth_getCode
	eth_getFilterChanges
*/


/*
	rpc method: "eth_getFilterLogs"
	Returns array of log objects, or an empty array if nothing has changed since last poll.
	curl --data '{"method":"eth_getFilterLogs","params":["0x16"],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
*/
func (eth Eth) GetFilterLogs(filterID string) ([]rpctypes.EtherLog, error) {
	return eth.client.RequestEtherLogList(MethodGetFilterLogs, filterID)
}

/*
	rpc method: "eth_getLogs"
	Returns array of log objects, or an empty array if nothing has changed since last poll.
	curl --data '{"method":"eth_getLogs","params":[{"topics":["0x000000000000000000000000a94f5374fce5edbc8e2a8697c15331677e6ebf0b"]}],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
*/
func (eth Eth) GetLogs(params *NewFilterParams) ([]rpctypes.EtherLog, error) {
	return eth.client.RequestEtherLogList(MethodGetLogs, params.ToMap())
}

/*
	rpc method: "eth_getStorageAt"
	Returns the value from a storage position at a given address.
	curl --data '{"method":"eth_getStorageAt","params":["0x407d73d8a49eeb85d32cf465507dd71d507100c1","0x0","0x2"],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
*/
func (eth Eth) GetStorageAt(contractAddress string, index int64, quantity *rpctypes.Quantity) (*rpctypes.HexString, error) {
	return eth.client.RequestHexString(MethodGetStorageAt, contractAddress, new(rpctypes.HexString).FromInt64(index).Hash(), quantity.String())
}

/*
	rpc method: "eth_getTransactionByBlockHashAndIndex"
	Returns information about a transaction by block hash and transaction index position.
*/
func (eth Eth) GetTransactionByBlockHashAndIndex(hash string, index int64) (*rpctypes.EtherTransaction, error) {
	return eth.client.RequestEtherTransaction(MethodGetTransactionByBlockHashAndIndex, hash, new(rpctypes.HexString).FromInt64(index).String())
}

/*
	rpc method: "eth_getTransactionByBlockHashAndIndex"
	Returns information about a transaction by block number and transaction index position.
*/
func (eth Eth) GetTransactionByBlockNumberAndIndex(blockNumber int64, index int64) (*rpctypes.EtherTransaction, error) {
	return eth.client.RequestEtherTransaction(MethodGetTransactionByBlockNumberAndIndex, new(rpctypes.HexString).FromInt64(blockNumber).String(), new(rpctypes.HexString).FromInt64(index).String())
}

/*
	rpc method: "eth_getTransactionByHash"
	Returns the information about a transaction requested by transaction hash.
*/
func (eth Eth) GetTransactionByHash(hash string) (*rpctypes.EtherTransaction, error) {
	return eth.client.RequestEtherTransaction(MethodGetTransactionByHash, hash)
}

/*
	rpc method: "eth_getTransactionCount"
	Returns the number of transactions sent from an address.
*/
func (eth Eth) GetTransactionCount(hash string, quantity *rpctypes.Quantity) (int64, error) {
	return eth.client.RequestInt64(MethodGetTransactionCount, hash, quantity.HexStringOrTag())
}

/*
	rpc method: "eth_getTransactionReceipt"
	Returns the number of transactions sent from an address.

	curl --data '{"method":"eth_getTransactionReceipt","params":["0x9676244c3a233b19a025184ea406fc5765f53edee7afabd901b470adcdeb5720"],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
*/
func (eth Eth) GetTransactionReceipt(hash string) (*rpctypes.EtherTransactionReceipt, error) {
	return eth.client.RequestEtherTransactionReceipt(MethodGetTransactionReceipt, hash)
}

/*
	eth_getUncleByBlockHashAndIndex
	eth_getUncleByBlockNumberAndIndex
	eth_getUncleCountByBlockHash
	eth_getUncleCountByBlockNumber
	eth_getWork
*/
/*
	rpc method: "eth_hashrate"
	Returns the number of hashes per second that the node is mining with.

	curl --data '{"method":"eth_hashrate","params":[],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
*/
func (eth Eth) Hashrate() (int64, error) {
	return eth.client.RequestInt64(MethodHashrate)
}

/*
	rpc method: "eth_mining"
	Returns true if client is actively mining new blocks.

	curl --data '{"method":"eth_mining","params":[],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
*/
func (eth Eth) Mining() (bool, error) {
	return eth.client.RequestBool(MethodMining)
}

/*
TODO
eth_newBlockFilter
*/

/*
	rpc method: "eth_newFilter"
	Creates a filter object, based on filter options, to notify when the state changes (logs). To check if the state has changed, call eth_getFilterChanges.
*/
func (eth Eth) NewFilter(params *NewFilterParams) (*rpctypes.HexString, error) {
	return eth.client.RequestHexString(MethodNewFilter, params.ToMap())
}

/*
TODO
eth_newPendingTransactionFilter
*/

/*
	rpc method: "eth_protocolVersion"
	Returns the current ethereum protocol version.
*/
func (eth Eth) ProtocolVersion() (int64, error) {
	return eth.client.RequestInt64(MethodProtocolVersion)
}

/*
	eth_sendRawTransaction
	eth_sendTransaction
	eth_sign
	eth_signTransaction
	eth_submitHashrate
	eth_submitWork
*/

/*
	rpc method: "eth_syncing"
	Returns an object with data about the sync status or false.
	curl --data '{"method":"eth_syncing","params":[],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
*/
func (eth Eth) Syncing() (*SyncStatus, error) {
	return eth.client.RequestSyncStatus(MethodSyncing)
}

/*
	rpc method: "eth_uninstallFilter"
	Uninstalls a filter with given id. Should always be called when watch is no longer needed. Additonally Filters timeout when they aren’t requested with eth_getFilterChanges for a period of time.
*/
func (eth Eth) UninstallFilter(filterID string) (bool, error) {
	return eth.client.RequestBool(MethodUninstallFilter, filterID)
}
