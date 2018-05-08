package processed

import (
	"fmt"
	"github.com/Leondroids/go-ethereum-rpc/rpc"
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
)

type EtherTransactionWithReceipt struct {
	Hash              rpctypes.HexString    `json:"hash"`
	From              rpctypes.EtherAddress `json:"from"`
	To                rpctypes.EtherAddress `json:"to"`
	Input             rpctypes.HexString    `json:"input"`
	Value             rpctypes.EtherValue   `json:"value"`
	BlockHash         rpctypes.HexString    `json:"blockHash"`
	BlockNumber       int64                 `json:"blockNumber"`
	Gas               rpctypes.EtherValue   `json:"gas"`
	GasPrice          rpctypes.EtherValue   `json:"gasPrice"`
	CumulativeGasUsed rpctypes.EtherValue   `json:"cumulativeGasUsed"`
	GasUsed           rpctypes.EtherValue   `json:"gasUsed"`
	ContractAddress   rpctypes.EtherAddress `json:"contractAddress"`
	Status            int64                 `json:"status"`
	LogsBloom         rpctypes.HexString    `json:"logsBloom"`
	Logs              []rpctypes.EtherLog   `json:"logs"`
	Nonce             rpctypes.HexString    `json:"nonce"`
	TransactionIndex  int64                 `json:"transactionIndex"`
	V                 rpctypes.HexString    `json:"v"`
	R                 rpctypes.HexString    `json:"r"`
	S                 rpctypes.HexString    `json:"s"`
}

type EtherTransactionWithReceiptDB interface {
	write(db *EtherTransactionWithReceiptDB) error
}

func LoadTransactionReceiptAndMerge(et *rpctypes.EtherTransaction, eth *rpc.Eth) (*EtherTransactionWithReceipt, error) {
	receipt, err := GetTransactionReceipt(et, eth)
	if err != nil {
		return nil, err
	}

	return MergeTransactionWithReceipt(et, receipt)
}

func GetTransactionReceipt(et *rpctypes.EtherTransaction, eth *rpc.Eth) (*rpctypes.EtherTransactionReceipt, error) {
	return eth.GetTransactionReceipt(et.Hash.Hash())
}

func MergeTransactionWithReceipt(transaction *rpctypes.EtherTransaction, receipt *rpctypes.EtherTransactionReceipt) (*EtherTransactionWithReceipt, error) {
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
