package rpc

import (
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
	"math/big"
)

type SendTransaction struct {
	From     *rpctypes.EtherAddress `json:"from"`
	To       *rpctypes.EtherAddress `json:"to"`
	Gas      *big.Int               `json:"gas"`
	GasPrice *big.Int               `json:"gasPrice"`
	Value    *big.Int               `json:"value"`
	Data     *rpctypes.HexString    `json:"data"`
}

func (it *SendTransaction) ToMap() map[string]interface{} {
	m := make(map[string]interface{})

	m["from"] = it.From.Hash()
	m["to"] = it.To.Hash()
	m["gas"] = new(rpctypes.HexString).FromBytes(it.Gas.Bytes())
	m["gasPrice"] = new(rpctypes.HexString).FromBytes(it.GasPrice.Bytes())
	m["value"] = new(rpctypes.HexString).FromBytes(it.Value.Bytes())
	m["data"] = it.Data.Hash()

	return m
}

func (it *SendTransaction) WithFrom(from *rpctypes.EtherAddress) *SendTransaction {
	it.From = from
	return it
}
