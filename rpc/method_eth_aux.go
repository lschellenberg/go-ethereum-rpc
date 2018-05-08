package rpc

import (
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
)

func (eth Eth) GetLatestBalance(address string, quantity *rpctypes.Quantity) (*rpctypes.EtherValue, error) {
	return eth.GetBalance(address, rpctypes.QuantityLatest())
}
