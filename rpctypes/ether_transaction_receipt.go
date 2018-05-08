package rpctypes

import (
	"fmt"
	"encoding/json"
	"github.com/Leondroids/gox"
)

type EtherTransactionReceipt struct {
	TransactionHash   HexString    `json:"transactionHash"`
	TransactionIndex  int64        `json:"transactionIndex"`
	BlockNumber       int64        `json:"blockNumber"`
	BlockHash         HexString    `json:"blockHash"`
	From              EtherAddress `json:"from"`
	To                EtherAddress `json:"to"`
	CumulativeGasUsed EtherValue   `json:"cumulativeGasUsed"`
	GasUsed           EtherValue   `json:"gasUsed"`
	ContractAddress   EtherAddress `json:"contractAddress"`
	Status            int64        `json:"status"`
	LogsBloom         HexString    `json:"logsBloom"`
	Logs              []EtherLog   `json:"logs"`
}

func (tr1 EtherTransactionReceipt) Compare(tr2 *EtherTransactionReceipt) error {
	if !tr1.TransactionHash.IsEqual(&tr2.TransactionHash) {
		return fmt.Errorf("not equal transactionHash %v %v", tr1.TransactionHash.String(), tr2.TransactionHash.String())
	}
	if tr1.TransactionIndex != tr2.TransactionIndex {
		return fmt.Errorf("not equal transactionIndex %v %v", tr1.TransactionIndex, tr2.TransactionIndex)
	}
	if tr1.BlockNumber != tr2.BlockNumber {
		return fmt.Errorf("not equal blockNumber %v %v", tr1.BlockNumber, tr2.BlockNumber)
	}
	if !tr1.BlockHash.IsEqual(&tr2.BlockHash) {
		return fmt.Errorf("not equal blockHash %v %v", tr1.BlockHash.String(), tr2.BlockHash.String())
	}
	if !tr1.From.IsEqual(&tr2.From) {
		return fmt.Errorf("error in from: [1: %v,2: %v]", tr1.From.String(), tr2.From.String())
	}
	if !tr1.To.IsEqual(&tr2.To) {
		return fmt.Errorf("error in to: [1: %v,2: %v]", tr1.To.String(), tr2.To.String())
	}
	if !tr1.CumulativeGasUsed.IsEqual(&tr2.CumulativeGasUsed) {
		return fmt.Errorf("not equal cumulativeGasUsed %v %v", tr1.CumulativeGasUsed.String(), tr2.CumulativeGasUsed.String())
	}
	if !tr1.GasUsed.IsEqual(&tr2.GasUsed) {
		return fmt.Errorf("not equal gasUsed %v %v", tr1.GasUsed.String(), tr2.GasUsed.String())
	}
	if !tr1.ContractAddress.IsEqual(&tr2.ContractAddress) {
		return fmt.Errorf("not equal contractAddress %v %v", tr1.TransactionHash.String(), tr2.TransactionHash.String())
	}
	if !tr1.LogsBloom.IsEqual(&tr2.LogsBloom) {
		return fmt.Errorf("error in logsBloom: [1: %v,2: %v]", tr1.LogsBloom, tr2.LogsBloom)
	}
	if tr1.Status != tr2.Status {
		return fmt.Errorf("error in status: [1: %v,2: %v]", tr1.Status, tr2.Status)
	}
	if err := CompareEtherLogList(tr1.Logs, tr2.Logs); err != nil {
		return err
	}

	return nil
}

// JSON Converter
type TransactionReceiptJSONConverter interface {
	FromJSON([]byte) (*EtherTransactionReceipt, error)
	ToJSON(*EtherTransactionReceipt) ([]byte, error)
}

func (raw *TransactionReceiptRaw) FromJSON(js []byte) (*EtherTransactionReceipt, error) {
	raw = new(TransactionReceiptRaw)
	err := json.Unmarshal(js, raw)

	if err != nil {
		return nil, err
	}

	return raw.ToEtherTransactionReceipt()
}

func (raw *TransactionReceiptRaw) ToEtherTransactionReceipt() (*EtherTransactionReceipt, error) {
	receipt := &EtherTransactionReceipt{}

	// TransactionHash
	th, err := NewHexString(raw.TransactionHash)
	if err != nil {
		return nil, fmt.Errorf("error parsing transactinHash, %v", err)
	}
	receipt.TransactionHash = *th
	// TransactionIndex
	ti, err := NewHexString(raw.TransactionIndex)
	if err != nil {
		return nil, fmt.Errorf("error parsing transactinIndex, %v", err)
	}
	receipt.BlockNumber = ti.Int64()
	// BlockNumber
	bn, err := NewHexString(raw.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("error parsing blockNumber, %v", err)
	}
	receipt.BlockNumber = bn.Int64()
	// BlockHash
	bh, err := NewHexString(raw.BlockHash)
	if err != nil {
		return nil, fmt.Errorf("error parsing blockHash, %v", err)
	}
	receipt.BlockHash = *bh
	// From parity doesnt give from
	if raw.From != "" {
		from, err := new(EtherAddress).FromString(raw.From)
		if err != nil {
			return nil, fmt.Errorf("error parsing from, %v", err)
		}
		receipt.From = *from
	}
	// To (parity doesnt give to)
	if raw.To != "" {
		to, err := new(EtherAddress).FromString(raw.To)
		if err != nil {
			return nil, fmt.Errorf("error parsing to, %v", err)
		}
		receipt.To = *to
	}
	// CumulativeGasUsed
	cgu, err := new(EtherValue).FromHexString(raw.CumulativeGasUsed)
	if err != nil {
		return nil, fmt.Errorf("error parsing cumulativeGasUsed, %v", err)
	}
	receipt.CumulativeGasUsed = *cgu
	// GasUsed
	gu, err := new(EtherValue).FromHexString(raw.GasUsed)
	if err != nil {
		return nil, fmt.Errorf("error parsing gasUsed, %v", err)
	}
	receipt.GasUsed = *gu
	// ContractAddress
	ca, err := new(EtherAddress).FromStringOrNull(raw.ContractAddress)
	if err != nil {
		return nil, fmt.Errorf("error parsing contract address, %v", err)
	}
	receipt.ContractAddress = *ca
	// Status
	status, err := NewHexString(raw.Status)
	if err != nil {
		return nil, fmt.Errorf("error parsing status, %v", err)
	}
	receipt.Status = status.Int64()
	// LogsBloom
	lb, err := NewHexString(raw.LogsBloom)
	if err != nil {
		return nil, fmt.Errorf("error parsing logsBloom, %v", err)
	}
	receipt.LogsBloom = *lb
	// Logs              []EtherLogRaw `json:"logs"`
	logs := make([]EtherLog, len(raw.Logs))
	for k, v := range raw.Logs {
		elog, err := v.ToEtherLog()
		if err != nil {
			return nil, err
		}
		logs[k] = *elog
	}
	receipt.Logs = logs

	return receipt, nil
}

func (raw *TransactionReceiptRaw) ToJSON(receipt *EtherTransactionReceipt) ([]byte, error) {
	return nil, gox.NotImplementedYet("TransactionReceiptRaw_ToJSON")
}

type TransactionReceiptRaw struct {
	TransactionHash   string        `json:"transactionHash"`
	TransactionIndex  string        `json:"transactionIndex"`
	BlockNumber       string        `json:"blockNumber"`
	BlockHash         string        `json:"blockHash"`
	From              string        `json:"from"`
	To                string        `json:"to"`
	CumulativeGasUsed string        `json:"cumulativeGasUsed"`
	GasUsed           string        `json:"gasUsed"`
	ContractAddress   string        `json:"contractAddress"`
	Status            string        `json:"status"`
	LogsBloom         string        `json:"logsBloom"`
	Logs              []EtherLogRaw `json:"logs"`
}
