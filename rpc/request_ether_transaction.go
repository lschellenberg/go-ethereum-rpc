package rpc

import (
	"encoding/json"
	"github.com/Leondroids/go-ethereum-rpc/types"
	"fmt"
)

func (client *Client) RequestEtherTransaction(method string, params ...interface{}) (*types.EtherTransaction, error) {
	response, err := checkRPCError(client.Call(method, params...))

	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("response returned without error but no transaction found for %v", params)
	}

	return getTransactionFromResponse(response.Result)
}

func getTransactionFromResponse(response interface{}) (*types.EtherTransaction, error) {
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

func (trans RPCEtherTransactionRaw) toEtherTransaction() (*types.EtherTransaction, error) {
	hash, err := types.NewHexString(trans.Hash)

	if err != nil {
		return nil, err
	}

	blockHash, err := types.NewHexString(trans.BlockHash)

	if err != nil {
		return nil, err
	}

	blockNumber, err := types.NewHexString(trans.BlockNumber)

	if err != nil {
		return nil, err
	}

	gas, err := types.NewHexString(trans.Gas)

	if err != nil {
		return nil, err
	}

	gasPrice, err := types.NewHexString(trans.GasPrice)

	if err != nil {
		return nil, err
	}

	to, err := new(types.EtherAddress).FromStringOrNull(trans.To)

	if err != nil {
		return nil, err
	}

	from, err := new(types.EtherAddress).FromStringOrNull(trans.From)

	if err != nil {
		return nil, err
	}

	nonce, err := types.NewHexString(trans.Nonce)

	if err != nil {
		return nil, err
	}

	input, err := types.NewHexString(trans.Input)

	if err != nil {
		return nil, err
	}

	transactionIndex, err := types.NewHexString(trans.TransactionIndex)

	if err != nil {
		return nil, err
	}

	value, err := new(types.EtherValue).FromHexString(trans.Value)

	if err != nil {
		return nil, err
	}

	v, err := types.NewHexString(trans.V)

	if err != nil {
		return nil, err
	}

	r, err := types.NewHexString(trans.R)

	if err != nil {
		return nil, err
	}

	s, err := types.NewHexString(trans.S)

	if err != nil {
		return nil, err
	}

	return &types.EtherTransaction{
		Hash:             *hash,
		BlockHash:        *blockHash,
		BlockNumber:      blockNumber.Int64(),
		Gas:              *gas.BigInt(),
		GasPrice:         *gasPrice.BigInt(),
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

func (client *Client) RequestEtherTransactionReceipt(method string, params ...interface{}) (*types.EtherTransactionReceipt, error) {
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

	return new(types.TransactionReceiptRaw).FromJSON(js)
}
