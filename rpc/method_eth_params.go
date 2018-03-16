package rpc

import "github.com/Leondroids/go-ethereum-rpc/types"

type EthCallParams struct {
	From     string         `json:"from"`      // (optional) 20 Bytes - The address the transaction is send from.
	To       string         `json:"to"`        // (optional when creating new contract) 20 Bytes - The address the transaction is directed to.
	Gas      int64          `json:"gas"`       // (optional) Integer of the gas provided for the transaction execution. eth_call consumes zero gas, but this parameter may be needed by some executions.
	GasPrice int64          `json:"gas_price"` // (optional) Integer of the gas price used for each paid gas.
	Value    int64          `json:"value"`     // (optional) Integer of the value sent with this transaction.
	Data     string         `json:"data"`      // (optional) 4 byte hash of the method signature followed by encoded parameters.
	Quantity types.Quantity `json:"quantity"`  // (optional) Integer block number, or the string 'latest', 'earliest' or 'pending', see the default block parameter.
}

type EthEstimateGasParams struct {
	From     string         `json:"from"`      // (optional) 20 Bytes - The address the transaction is send from.
	To       string         `json:"to"`        // (optional when creating new contract) 20 Bytes - The address the transaction is directed to.
	Gas      int64          `json:"gas"`       // (optional) Integer of the gas provided for the transaction execution. eth_call consumes zero gas, but this parameter may be needed by some executions.
	GasPrice int64          `json:"gas_price"` // (optional) Integer of the gas price used for each paid gas.
	Value    int64          `json:"value"`     // (optional) Integer of the value sent with this transaction.
	Data     string         `json:"data"`      // (optional) 4 byte hash of the method signature followed by encoded parameters.
	Quantity types.Quantity `json:"quantity"`  // (optional) Integer block number, or the string 'latest', 'earliest' or 'pending', see the default block parameter.
}


type NewFilterParams struct {

}