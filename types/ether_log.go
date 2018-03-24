package types

import (
	"fmt"
	"encoding/json"
	"github.com/Leondroids/gox"
)

type EtherLog struct {
	Address          HexString   `json:"address"`
	Topics           []HexString `json:"topics"`
	Data             HexString   `json:"data"`
	BlockNumber      int64       `json:"blockNumber"`
	BlockHash        HexString   `json:"blockHash"`
	TransactionHash  HexString   `json:"transactionHash"`
	TransactionIndex int64       `json:"transactionIndex"`
	LogIndex         int64       `json:"logIndex"`
	Removed          bool        `json:"removed"`
}

func (el1 EtherLog) Compare(el2 EtherLog) error {
	if !el1.Address.IsEqual(&el2.Address) {
		return fmt.Errorf("error in address: [1: %v,2: %v]", el1.Address, el2.Address)
	}
	err := CompareHexStringList(el1.Topics, el2.Topics)
	if err != nil {
		return fmt.Errorf("error in topice: [1: %v,2: %v]", el1.Topics, el2.Topics)
	}
	if !el1.Data.IsEqual(&el2.Data) {
		return fmt.Errorf("error in data: [1: %v,2: %v]", el1.Data.String(), el2.Data.String())
	}
	if el1.BlockNumber != el2.BlockNumber {
		return fmt.Errorf("error in blockNumber: [1: %v,2: %v]", el1.BlockNumber, el2.BlockNumber)
	}
	if !el1.BlockHash.IsEqual(&el2.BlockHash) {
		return fmt.Errorf("error in blockHash: [1: %v,2: %v]", el1.BlockHash, el2.BlockHash)
	}
	if !el1.TransactionHash.IsEqual(&el2.TransactionHash) {
		return fmt.Errorf("error in transactionHash: [1: %v,2: %v]", el1.TransactionHash, el2.TransactionHash)
	}
	if el1.TransactionIndex != el2.TransactionIndex {
		return fmt.Errorf("error in transactionIndex: [1: %v,2: %v]", el1.TransactionIndex, el2.TransactionIndex)
	}
	if el1.LogIndex != el2.LogIndex {
		return fmt.Errorf("error in logIndex: [1: %v,2: %v]", el1.LogIndex, el2.LogIndex)
	}
	if el1.Removed != el2.Removed {
		return fmt.Errorf("error in removed: [1: %v,2: %v]", el1.Removed, el2.Removed)
	}
	return nil
}

func CompareEtherLogList(ell1 []EtherLog, ell2 []EtherLog) error {
	if len(ell1) != len(ell2) {
		return fmt.Errorf("wrong list sizes, left: %v, reight: %v", len(ell1), len(ell2))
	}

	for k, v := range ell1 {
		err := v.Compare(ell2[k])
		if err != nil {
			return fmt.Errorf("not equal ether log list at %v, left: %v - right: %v", k, v, ell2[k])
		}
	}
	return nil
}

type EtherLogJSONConverter interface {
	FromJSON([]byte) (*EtherLog, error)
	ToJSON(*EtherLog) ([]byte, error)
}

func (raw *EtherLogRaw) FromJSON(js []byte) (*EtherLog, error) {
	err := json.Unmarshal(js, raw)

	if err != nil {
		return nil, err
	}
	return raw.ToEtherLog()
}

func (raw *EtherLogRaw) ToEtherLog() (*EtherLog, error) {
	log := &EtherLog{}

	// address
	address, err := NewHexString(raw.Address)
	if err != nil {
		return nil, fmt.Errorf("error parsing address, %v", err)
	}
	log.Address = *address

	// topics
	topics, err := ToHexStringList(raw.Topics)
	if err != nil {
		return nil, fmt.Errorf("error parsing topics, %v", err)
	}
	log.Topics = topics

	// data
	data, err := NewHexString(raw.Data)
	if err != nil {
		return nil, fmt.Errorf("error parsing data, %v", err)
	}
	log.Data = *data

	// data
	blocknumber, err := NewHexString(raw.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("error parsing blocknumber, %v", err)
	}
	log.BlockNumber = blocknumber.Int64()

	// transactionHash
	transactionhash, err := NewHexString(raw.TransactionHash)
	if err != nil {
		return nil, fmt.Errorf("error parsing transactionHash, %v", err)
	}
	log.TransactionHash = *transactionhash

	// TransactionIndex
	transactionIndex, err := NewHexString(raw.TransactionIndex)
	if err != nil {
		return nil, fmt.Errorf("error parsing transactionIndex, %v", err)
	}
	log.TransactionIndex = transactionIndex.Int64()

	// blockHash
	blockhash, err := NewHexString(raw.BlockHash)
	if err != nil {
		return nil, fmt.Errorf("error parsing blockHash, %v", err)
	}
	log.BlockHash = *blockhash

	// LogIndex
	logIndex, err := NewHexString(raw.LogIndex)
	if err != nil {
		return nil, fmt.Errorf("error parsing logIndex, %v", err)
	}
	log.TransactionIndex = logIndex.Int64()

	// Removed
	log.Removed = raw.Removed

	return log, nil
}

func (raw *EtherLogRaw) ToJSON(log *EtherLog) ([]byte, error) {
	return nil, gox.NotImplementedYet("EtherLogRaw_ToJSON")
}

type EtherLogRaw struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
}
