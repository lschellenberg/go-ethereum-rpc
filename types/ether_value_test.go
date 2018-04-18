package types

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

	result, err := new(EtherValue).FromDotString(s)

	if err != nil {
		t.Error(err)
	}
	expected := "0.1"
	if result.String() != expected {
		t.Errorf("dot string not equal,[Expected: %v, Actual: %v]", expected, result.String())
	}

	// Test 2
	s1 := "266482000000000000000"
	result1, err := new(EtherValue).FromDotString(s1)

	if err != nil {
		t.Error(err)
	}

	if len(result1.value.Bytes()) != 9 {
		t.Errorf("wrong size of ether value, [Expected: %v, Acutal: %v]", 9, len(result1.value.Bytes()))
	}

	expected1 := "266.482"
	if result1.String() != expected1 {
		t.Errorf("dot string not equal,[Expected: %v, Actual: %v]", expected1, result1.String())
	}

	// Test 3
	s2 := "266000000000000000000"
	result2, err := new(EtherValue).FromDotString(s2)

	if err != nil {
		t.Error(err)
	}

	if len(result2.value.Bytes()) != 9 {
		t.Errorf("wrong size of ether value, [Expected: %v, Acutal: %v]", 9, len(result.value.Bytes()))
	}

	expected2 := "266"
	if result1.String() != expected2 {
		t.Errorf("dot string not equal,[Expected: %v, Actual: %v]", expected2, result2.String())
	}
}
