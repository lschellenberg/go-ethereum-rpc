package types

import (
	"testing"
)

func TestHexString_Concat(t *testing.T) {
	hs1, err := NewHexString("0x1234")
	if err != nil {
		t.Error(err)
	}
	hs2, err := NewHexString("0x56")
	if err != nil {
		t.Error(err)
	}

	result := hs1.Concat(hs2)
	expected := "0x123456"

	if result.String() != expected {
		t.Errorf("concat is wrong,[Expected: %v, Actual: %v", expected, result.String())
	}

}

func TestHexString_PadTo(t *testing.T) {
	hs, err := NewHexString("0x1234")

	if err != nil {
		t.Error(err)
	}

	// pad to same
	result := hs.PadTo(2)
	expected := "0x1234"

	if result.Hash() != expected {
		t.Errorf("pad is wrong for same size,[Expected: %v, Actual: %v", expected, result.Hash())
		return
	}

	// pad to bigger
	result = hs.PadTo(3)
	expected = "0x001234"

	if result.Hash() != expected {
		t.Errorf("pad is wrong for bigger pad,[Expected: %v, Actual: %v", expected, result.Hash())
		return
	}

	// pad to smaller
	result = hs.PadTo(1)
	expected = "0x34"

	if result.Hash() != expected {
		t.Errorf("pad is wrong for bigger pad,[Expected: %v, Actual: %v", expected, result.Hash())
		return
	}
}
