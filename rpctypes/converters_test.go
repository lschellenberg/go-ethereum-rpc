package rpctypes

import (
	"testing"
	"math/big"
	"errors"
	"fmt"
)

func TestByteToHex(t *testing.T) {
	s := "test"
	expectedHex := "0x74657374"
	hex := ByteToHex([]byte(s))
	if hex != expectedHex {
		t.Errorf("wrong hex string [Expected: %v, Actual: %v]", expectedHex, hex)
	}
}

func TestInterfaceToStringArray(t *testing.T) {
	expected := []string{"test1", "test2"}
	i := make([]interface{}, len(expected))

	for k, v := range expected {
		i[k] = v
	}

	result, err := InterfaceListToStringList(i)

	if err != nil {
		t.Error(err)
		return
	}

	for k, v := range expected {
		if result[k] != v {
			t.Errorf("couldn't convert interface to string at %v [Expected: %v, Actual: %v]", k, v, result[k])
		}
	}
}

func TestHexToBigInt(t *testing.T) {
	h := "0x153711f0a39755800"

	expected, ok := big.NewInt(0).SetString("24459365180000000000", 10)
	if !ok {
		t.Error(errors.New("couldn't convert big string int to big.Int"))
		return
	}

	result, err := HexToBigInt(h)

	if err != nil {
		t.Error(err)
		return
	}

	if result.Cmp(expected) != 0 {
		t.Errorf("couldn't convert hex to big int [Expected: %v, Actual: %v]", expected, result)
	}
}

func TestHexToByte(t *testing.T) {
	s := "0x0"
	result, err := HexToByte(s)
	if err != nil {
		t.Error(err)
		return
	}
	if len(result) != 1 || result[0] != 0 {
		t.Error(fmt.Errorf("error in parsing string Result: %v", result))
	}
}
