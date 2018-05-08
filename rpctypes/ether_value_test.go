package rpctypes

import (
	"testing"
	"math/big"
)

func TestEtherValue_Add(t *testing.T) {
	ev1 := NewEtherValueFromBigInt(big.NewInt(1299))
	ev2 := NewEtherValueFromBigInt(big.NewInt(31701))

	result := ev1.Add(ev2)
	var expected int64 = 33000

	if expected != result.BigInt().Int64() {
		t.Errorf("sum is wrong,[Expected: %v, Actual: %v", expected, result.BigInt().Int64())
	}
}

func TestEtherValue_FromDotString(t *testing.T) {
	s := "0.1"

	result, err := new(EtherValue).FromFloat64String(s)

	if err != nil {
		t.Error(err)
	}
	actual := result.BigInt().String()
	expected := "100000000000000000"
	if actual != expected {
		t.Errorf("dot string not equal,[Expected: %v, Actual: %v]", expected, actual)
	}

	// Test 2
	s = "12345678.123981345"

	result, err = new(EtherValue).FromFloat64String(s)

	if err != nil {
		t.Error(err)
	}
	actual = result.BigInt().String()
	expected = "12345678123981345000000000"
	if actual != expected {
		t.Errorf("dot string not equal,[Expected: %v, Actual: %v]", expected, actual)
	}

	// Test 3
	s = "74"

	result, err = new(EtherValue).FromFloat64String(s)

	if err != nil {
		t.Error(err)
	}
	actual = result.BigInt().String()
	expected = "74000000000000000000"
	if actual != expected {
		t.Errorf("dot string not equal,[Expected: %v, Actual: %v]", expected, actual)
	}

	// Test 4
	s = "74.0"

	result, err = new(EtherValue).FromFloat64String(s)

	if err != nil {
		t.Error(err)
	}
	actual = result.BigInt().String()
	expected = "74000000000000000000"
	if actual != expected {
		t.Errorf("dot string not equal,[Expected: %v, Actual: %v]", expected, actual)
	}
}
