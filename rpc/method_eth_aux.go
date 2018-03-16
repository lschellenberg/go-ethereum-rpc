package rpc

import (
	"github.com/Leondroids/go-ethereum-rpc/types"
)

func (eth Eth) GetLatestBalance(address string, quantity *types.Quantity) (*types.EtherValue, error) {
	return eth.GetBalance(address, types.QuantityLatest())
}
