package rpcutils

import (
	"testing"
	"log"
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
	"math/big"
)

var dataEventKittiesData = "0x" +
	"00000000000000000000000079bd592415ff6c91cfe69a7f9cd091354fc65a18" +
	"00000000000000000000000000000000000000000000000000000000000004a2" +
	"0000000000000000000000000000000000000000000000000000000000000487" +
	"0000000000000000000000000000000000000000000000000000000000465b93"

// input [{address false 0} {uint256 false 0} {uint256 false 0} {uint256 false 0}]
var fs = NewFunctionSignature(
	*NewFunctionParamType("address", 0),
	*NewFunctionParamType("uint256", 0),
	*NewFunctionParamType("uint256", 0),
	*NewFunctionParamType("uint256", 0),
)

func TestFunctionSignature_DecodeEventData(t *testing.T) {
	result, err := fs.DecodeEventData(dataEventKittiesData)
	if err != nil {
		t.Error(err)
		return
	}

	log.Println(result)

	if len(result) != 4 {
		t.Error("wrong length,[Expected:%v, Actual: %v]", 4, len(result))
		return
	}

	if err := CompareString("0x79bd592415ff6c91cfe69a7f9cd091354fc65a18", result[0].Value.(*rpctypes.EtherAddress).String()); err != nil {
		t.Error(err)
		return
	}

	if err := CompareBigInt(new(big.Int).SetInt64(1186), result[1].Value.(*big.Int)); err != nil {
		t.Error(err)
		return
	}
	if err := CompareBigInt(new(big.Int).SetInt64(1159), result[2].Value.(*big.Int)); err != nil {
		t.Error(err)
		return
	}
	if err := CompareBigInt(new(big.Int).SetInt64(4610963), result[3].Value.(*big.Int)); err != nil {
		t.Error(err)
		return
	}

}
