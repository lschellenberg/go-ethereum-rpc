package processed

import (
	"github.com/Leondroids/go-ethereum-rpc/types"
	"fmt"
	"github.com/Leondroids/go-ethereum-rpc/rpc"
)

type EtherTransactionWithReceipt struct {
	Hash              types.HexString    `json:"hash"`
	From              types.EtherAddress `json:"from"`
	To                types.EtherAddress `json:"to"`
	Input             types.HexString    `json:"input"`
	Value             types.EtherValue   `json:"value"`
	BlockHash         types.HexString    `json:"blockHash"`
	BlockNumber       int64              `json:"blockNumber"`
	Gas               types.EtherValue   `json:"gas"`
	GasPrice          types.EtherValue   `json:"gasPrice"`
	CumulativeGasUsed types.EtherValue   `json:"cumulativeGasUsed"`
	GasUsed           types.EtherValue   `json:"gasUsed"`
	ContractAddress   types.EtherAddress `json:"contractAddress"`
	Status            int64              `json:"status"`
	LogsBloom         types.HexString    `json:"logsBloom"`
	Logs              []types.EtherLog   `json:"logs"`
	Nonce             types.HexString    `json:"nonce"`
	TransactionIndex  int64              `json:"transactionIndex"`
	V                 types.HexString    `json:"v"`
	R                 types.HexString    `json:"r"`
	S                 types.HexString    `json:"s"`
}

type EtherTransactionWithReceiptDB interface {
	write(db *EtherTransactionWithReceiptDB) error
}

func LoadTransactionReceiptAndMerge(et *types.EtherTransaction, eth *rpc.Eth) (*EtherTransactionWithReceipt, error) {
	receipt, err := GetTransactionReceipt(et, eth)
	if err != nil {
		return nil, err
	}

	return MergeTransactionWithReceipt(et, receipt)
}

func GetTransactionReceipt(et *types.EtherTransaction, eth *rpc.Eth) (*types.EtherTransactionReceipt, error) {
	return eth.GetTransactionReceipt(et.Hash.Hash())
}

func MergeTransactionWithReceipt(transaction *types.EtherTransaction, receipt *types.EtherTransactionReceipt) (*EtherTransactionWithReceipt, error) {
	if !transaction.Hash.IsEqual(&receipt.TransactionHash) {
		return nil, fmt.Errorf("transaction and receipt are not equal, [transaction: %v, receipt: %v]", transaction.Hash.String(), receipt.TransactionHash.String())
	}

	return &EtherTransactionWithReceipt{
		Hash:              transaction.Hash,
		GasPrice:          transaction.GasPrice,
		Logs:              receipt.Logs,
		LogsBloom:         receipt.LogsBloom,
		Status:            receipt.Status,
		ContractAddress:   receipt.ContractAddress,
		GasUsed:           receipt.GasUsed,
		CumulativeGasUsed: receipt.CumulativeGasUsed,
		TransactionIndex:  receipt.TransactionIndex,
		BlockNumber:       receipt.BlockNumber,
		BlockHash:         receipt.BlockHash,
		Value:             transaction.Value,
		V:                 transaction.V,
		Input:             transaction.Input,
		Nonce:             transaction.Nonce,
		From:              transaction.From,
		To:                transaction.To,
		Gas:               transaction.Gas,
		R:                 transaction.R,
		S:                 transaction.S,
	}, nil
}
