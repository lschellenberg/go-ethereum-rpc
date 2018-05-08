package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
)

func (client *Client) RequestEtherTransaction(method string, params ...interface{}) (*rpctypes.EtherTransaction, error) {
	response, err := checkRPCError(client.Call(method, params...))

	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("response returned without error but no transaction found for %v", params)
	}

	return getTransactionFromResponse(response.Result)
}

func getTransactionFromResponse(response interface{}) (*rpctypes.EtherTransaction, error) {
	js, err := json.Marshal(response)

	if err != nil {
		return nil, err
	}

	r := new(RPCEtherTransactionRaw)
	err = json.Unmarshal(js, r)

	if err != nil {
		return nil, err
	}

	return r.toEtherTransaction()
}

type RPCEtherTransactionRaw struct {
	Hash             string `json:"hash"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	From             string `json:"from"`
	To               string `json:"to"`
	Nonce            string `json:"nonce"`
	Input            string `json:"input"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

func (trans RPCEtherTransactionRaw) toEtherTransaction() (*rpctypes.EtherTransaction, error) {
	hash, err := rpctypes.NewHexString(trans.Hash)

	if err != nil {
		return nil, err
	}

	blockHash, err := rpctypes.NewHexString(trans.BlockHash)

	if err != nil {
		return nil, err
	}

	blockNumber, err := rpctypes.NewHexString(trans.BlockNumber)

	if err != nil {
		return nil, err
	}

	gas, err := rpctypes.NewEtherValue().FromHexString(trans.Gas)

	if err != nil {
		return nil, err
	}

	gasPrice, err := rpctypes.NewEtherValue().FromHexString(trans.GasPrice)

	if err != nil {
		return nil, err
	}

	to, err := new(rpctypes.EtherAddress).FromStringOrNull(trans.To)

	if err != nil {
		return nil, err
	}

	from, err := new(rpctypes.EtherAddress).FromStringOrNull(trans.From)

	if err != nil {
		return nil, err
	}

	nonce, err := rpctypes.NewHexString(trans.Nonce)

	if err != nil {
		return nil, err
	}

	input, err := rpctypes.NewHexString(trans.Input)

	if err != nil {
		return nil, err
	}

	transactionIndex, err := rpctypes.NewHexString(trans.TransactionIndex)

	if err != nil {
		return nil, err
	}

	value, err := new(rpctypes.EtherValue).FromHexString(trans.Value)

	if err != nil {
		return nil, err
	}

	v, err := rpctypes.NewHexString(trans.V)

	if err != nil {
		return nil, err
	}

	r, err := rpctypes.NewHexString(trans.R)

	if err != nil {
		return nil, err
	}

	s, err := rpctypes.NewHexString(trans.S)

	if err != nil {
		return nil, err
	}

	return &rpctypes.EtherTransaction{
		Hash:             *hash,
		BlockHash:        *blockHash,
		BlockNumber:      blockNumber.Int64(),
		Gas:              *gas,
		GasPrice:         *gasPrice,
		From:             *from,
		To:               *to,
		Nonce:            *nonce,
		Input:            *input,
		TransactionIndex: transactionIndex.Int64(),
		Value:            *value,
		V:                *v,
		R:                *r,
		S:                *s,
	}, nil
}

// transaction receipt

func (client *Client) RequestEtherTransactionReceipt(method string, params ...interface{}) (*rpctypes.EtherTransactionReceipt, error) {
	response, err := checkRPCError(client.Call(method, params...))
	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("response returned without error but no transaction found for %v", params)
	}

	js, err := json.Marshal(response.Result)

	if err != nil {
		return nil, err
	}

	return new(rpctypes.TransactionReceiptRaw).FromJSON(js)
}
