package types

import (
	"errors"
	"fmt"
)

type EtherTransaction struct {
	Hash             HexString    `json:"hash"`
	BlockHash        HexString    `json:"blockHash"`
	BlockNumber      int64        `json:"blockNumber"`
	Gas              EtherValue   `json:"gas"`
	GasPrice         EtherValue   `json:"gasPrice"`
	From             EtherAddress `json:"from"`
	To               EtherAddress `json:"to"`
	Nonce            HexString    `json:"nonce"`
	Input            HexString    `json:"input"`
	TransactionIndex int64        `json:"transactionIndex"`
	Value            EtherValue   `json:"value"`
	V                HexString    `json:"v"`
	R                HexString    `json:"r"`
	S                HexString    `json:"s"`
}

func CompareEtherTransactionList(et1 []EtherTransaction, et2 []EtherTransaction, onlyHash bool) error {
	if len(et1) != len(et2) {
		return errors.New("wrong sizes in HexString List")
	}

	for k, v := range et1 {
		if onlyHash {
			if !v.Hash.IsEqual(&et2[k].Hash) {
				return fmt.Errorf("not equal at position %v, hash: %v, %v", k, v.Hash.value, et2[k].Hash.value)
			}
		} else {
			if err := v.Compare(&et2[k]); err != nil {
				return fmt.Errorf("not equal at position %v, msg: %v", k, err.Error())
			}
		}
	}

	return nil
}

func (et1 EtherTransaction) Compare(et2 *EtherTransaction) error {
	if !et1.Hash.IsEqual(&et2.Hash) {
		return fmt.Errorf("error in hash: [1: %v,2: %v]", et1.Hash, et2.Hash)
	}
	if !et1.BlockHash.IsEqual(&et2.BlockHash) {
		return fmt.Errorf("error in blockHash: [1: %v,2: %v]", et1.BlockHash, et2.BlockHash)
	}
	if et1.BlockNumber != et2.BlockNumber {
		return fmt.Errorf("error in blockNumber: [1: %v,2: %v]", et1.BlockNumber, et2.BlockNumber)
	}
	if !et1.Gas.IsEqual(&et2.Gas) {
		return fmt.Errorf("error in gas: [1: %v,2: %v]", et1.Gas.String(), et2.Gas.String())
	}
	if !et1.GasPrice.IsEqual(&et2.GasPrice) {
		return fmt.Errorf("error in gasPrice: [1: %v,2: %v]", et1.GasPrice.String(), et2.GasPrice.String())
	}
	if et1.From.value != et2.From.value {
		return fmt.Errorf("error in from: [1: %v,2: %v]", et1.From.String(), et2.From.String())
	}
	if et1.To.value != et2.To.value {
		return fmt.Errorf("error in to: [1: %v,2: %v]", et1.To.String(), et2.To.String())
	}
	if !et1.Nonce.IsEqual(&et2.Nonce) {
		return fmt.Errorf("error in nonce: [1: %v,2: %v]", et1.Nonce.value, et2.Nonce.value)
	}
	if !et1.Input.IsEqual(&et2.Input) {
		return fmt.Errorf("error in input: [1: %v,2: %v]", et1.Input.value, et2.Input.value)
	}
	if et1.TransactionIndex != et2.TransactionIndex {
		return fmt.Errorf("error in transactionIndex: [1: %v,2: %v]", et1.TransactionIndex, et2.TransactionIndex)
	}
	if et1.Value.value.Cmp(&et2.Value.value) != 0 {
		return fmt.Errorf("error in value: [1: %v,2: %v]", et1.Value.String(), et2.Value.String())
	}
	if !et1.V.IsEqual(&et2.V) {
		return fmt.Errorf("error in v: [1: %v,2: %v]", et1.V.value, et2.V.value)
	}
	if !et1.R.IsEqual(&et2.R) {
		return fmt.Errorf("error in r: [1: %v,2: %v]", et1.R.value, et2.R.value)
	}
	if !et1.S.IsEqual(&et2.S) {
		return fmt.Errorf("error in s: [1: %v,2: %v]", et1.S.value, et2.S.value)
	}
	return nil
}
