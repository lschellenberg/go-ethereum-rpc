package rpcutils

import (
	"testing"
	"fmt"
	"reflect"
	"math/big"
	"log"
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
)

const inputDeliverToken = "0x3eeda7d3" +
	"000000000000000000000000499d4aab8acab015d319ebfe476e4c800079ce87" +
	"000000000000000000000000000000000000000000000000000000000000e9d6" +
	"0000000000000000000000000000000000000000000000000000000000000080" +
	"0000000000000000000000000000000000000000000000000000000000000001" +
	"0000000000000000000000000000000000000000000000000000000000000005" +
	"4a6e784e33000000000000000000000000000000000000000000000000000000"

var functionSignatureDeliverToken = NewFunctionSignature(FPTAddress, FPTUInt256, FPTString, FPTBool)

func TestFunctionSignature_ReadHead(t *testing.T) {
	hs, _ := rpctypes.NewHexString(inputDeliverToken)
	head := functionSignatureDeliverToken.ReadHead(hs)

	if len(head) != 4 {
		t.Errorf("wrong result length of input head, [Expected: %v, Actual: %v]", 4, len(head))
	}

	if err := CompareString("000000000000000000000000499d4aab8acab015d319ebfe476e4c800079ce87", head[0].Plain()); err != nil {
		t.Error(err)
		return
	}

	if err := CompareString("000000000000000000000000000000000000000000000000000000000000e9d6", head[1].Plain()); err != nil {
		t.Error(err)
		return
	}

	if err := CompareString("0000000000000000000000000000000000000000000000000000000000000080", head[2].Plain()); err != nil {
		t.Error(err)
		return
	}

	if err := CompareString("0000000000000000000000000000000000000000000000000000000000000001", head[3].Plain()); err != nil {
		t.Error(err)
		return
	}
}

func TestFunctionSignature_DecodeFunctionInputWithDeliverTokens(t *testing.T) {
	fs := NewFunctionSignature(FPTAddress, FPTUInt256, FPTString, FPTBool)

	result, err := fs.DecodeFunctionInput(inputDeliverToken)

	if err != nil {
		t.Error(err)
		return
	}

	if len(result) != 4 {
		t.Errorf("wrong result length of function params, [Expected: %v, Actual: %v]", 4, len(result))
		return
	}

	param0 := result[0]
	param1 := result[1]
	param2 := result[2]
	param3 := result[3]

	// check types
	if err := CompareString("address", param0.Type); err != nil {
		t.Error(err)
		return
	}
	if err := CompareString("uint256", param1.Type); err != nil {
		t.Error(err)
		return
	}
	if err := CompareString("string", param2.Type); err != nil {
		t.Error(err)
		return
	}
	if err := CompareString("bool", param3.Type); err != nil {
		t.Error(err)
		return
	}

	// check returned value types
	if err = CompareType("*types.EtherAddress", param0.Value); err != nil {
		t.Error(err)
		return
	}
	if err = CompareType("*big.Int", param1.Value); err != nil {
		t.Error(err)
		return
	}
	// JnxN3
	if err = CompareType("string", param2.Value); err != nil {
		t.Error(err)
		return
	}

	if err = CompareType("bool", param3.Value); err != nil {
		t.Error(err)
		return
	}

	// check content
	expectedAddress, _ := new(rpctypes.EtherAddress).FromString("0x499d4aab8acab015d319ebfe476e4c800079ce87")
	if err := CompareEtherAddress(expectedAddress, param0.Value.(*rpctypes.EtherAddress)); err != nil {
		t.Error(err)
		return
	}
	if err := CompareBigInt(new(big.Int).SetInt64(59862), param1.Value.(*big.Int)); err != nil {
		t.Error(err)
		return
	}
	if err := CompareString("JnxN3", param2.Value.(string)); err != nil {
		t.Error(err)
		return
	}
	if err := CompareBool(true, param3.Value.(bool)); err != nil {
		t.Error(err)
		return
	}
}

var inputBaz = "0xcdcd77c000000000000000000000000000000000000000000000000000000000000000450000000000000000000000000000000000000000000000000000000000000001"
var functionSignatureBaz = NewFunctionSignature(FPTUInt32, FPTBool)

func TestFunctionSignature_DecodeFunctionInputForBaz(t *testing.T) {
	fs := functionSignatureBaz

	result, err := fs.DecodeFunctionInput(inputBaz)

	if err != nil {
		t.Error(err)
		return
	}

	if len(result) != 2 {
		t.Errorf("wrong result length of function params, [Expected: %v, Actual: %v]", 4, len(result))
		return
	}

	param0 := result[0]
	param1 := result[1]

	// check types
	if err := CompareString("uint32", param0.Type); err != nil {
		t.Error(err)
		return
	}
	if err := CompareString("bool", param1.Type); err != nil {
		t.Error(err)
		return
	}

	// check returned value types
	if err = CompareType("int", param0.Value); err != nil {
		t.Error(err)
		return
	}
	if err = CompareType("bool", param1.Value); err != nil {
		t.Error(err)
		return
	}

	// check content
	if err := CompareInt(69, param0.Value.(int)); err != nil {
		t.Error(err)
		return
	}
	if err := CompareBool(true, param1.Value.(bool)); err != nil {
		t.Error(err)
		return
	}
}

var inputSam = "0xa5643bf20000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000464617665000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003"
var functionSignatureSam = NewFunctionSignature(FPTBytes, FPTBool, *NewFunctionParamType("uint256", -1))

func TestFunctionSignature_DecodeFunctionInputForSam(t *testing.T) {
	fs := functionSignatureSam

	result, err := fs.DecodeFunctionInput(inputSam)

	if err != nil {
		t.Error(err)
		return
	}

	if len(result) != 2 {
		t.Errorf("wrong result length of function params, [Expected: %v, Actual: %v]", 4, len(result))
		return
	}

	param0 := result[0]
	param1 := result[1]

	log.Println(param0)
	//	param2 := result[2]

	// check types
	if err := CompareString("uint32", param0.Type); err != nil {
		t.Error(err)
		return
	}
	if err := CompareString("bool", param1.Type); err != nil {
		t.Error(err)
		return
	}
	//
	//// check returned value types
	//if err = CompareType("int", param0.Value); err != nil {
	//	t.Error(err)
	//	return
	//}
	//if err = CompareType("bool", param1.Value); err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//// check content
	//if err := CompareInt(69, param0.Value.(int)); err != nil {
	//	t.Error(err)
	//	return
	//}
	//if err := CompareBool(true, param1.Value.(bool)); err != nil {
	//	t.Error(err)
	//	return

}

// utils

func CompareBool(expected bool, actual bool) error {
	if expected != actual {
		return fmt.Errorf("not equal,[\nExpected: %v \nActual  : %v", expected, actual)
	}

	return nil
}

func CompareString(expected string, actual string) error {
	if expected != actual {
		return fmt.Errorf("not equal,[\nExpected: %v \nActual  : %v", expected, actual)
	}

	return nil
}

func CompareEtherAddress(expected *rpctypes.EtherAddress, actual *rpctypes.EtherAddress) error {
	if !expected.IsEqual(actual) {
		return fmt.Errorf("not equal,[\nExpected: %v \nActual  : %v", expected, actual)
	}

	return nil
}

func CompareInt(expected int, actual int) error {
	if expected != actual {
		return fmt.Errorf("not equal,[\nExpected: %v \nActual  : %v", expected, actual)
	}

	return nil
}
func CompareInt64(expected int64, actual int64) error {
	if expected != actual {
		return fmt.Errorf("not equal,[\nExpected: %v \nActual  : %v", expected, actual)
	}

	return nil
}
func CompareBigInt(expected *big.Int, actual *big.Int) error {
	if expected.Cmp(actual) != 0 {
		return fmt.Errorf("not equal,[\nExpected: %v \nActual  : %v", expected, actual)
	}

	return nil
}

func CompareType(expected string, actual interface{}) error {
	r := reflect.TypeOf(actual)
	actualKind := r.String()
	if expected != actualKind {
		return fmt.Errorf("not equal,[\nExpected: %v \nActual  : %v", expected, actualKind)
	}

	return nil
}
