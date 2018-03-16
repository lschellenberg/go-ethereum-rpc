package types

import "testing"


func TestEtherAddress_FromStringOrNull(t *testing.T) {
	a := "0x462cd55eae454ced675a267d22ebafe743b96528"

	ea, err := new(EtherAddress).FromStringOrNull(a)

	if err!= nil {
		t.Error(err)
	}

	if ea.String() != a {
		t.Errorf("wrong address result, [Expected %v, Actual: %v]", a, ea.String())
	}
}

func TestEtherAddress_FromBytes(t *testing.T) {
	bytes := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	ea, err := new(EtherAddress).FromBytes(bytes)

	if err != nil {
		t.Error(err)
		return
	}

	expected := "0x0102030405060708090a0b0c0d0e0f1011121314"
	if expected != ea.String() {
		t.Errorf("wrong address result, [Expected %v, Actual: %v]", expected, ea.String())
	}

	wrongbytes := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	_, err = new(EtherAddress).FromBytes(wrongbytes)
	if err == nil {
		t.Errorf("should return error if wrong address size (%v)", wrongbytes)
	}

}

func TestEtherAddress_FromHexString(t *testing.T) {
	address := "0x3333333333333333333333333333333333333333"

	ea := new(EtherAddress)

	var expected EtherAddress

	for k, _ := range expected.Bytes() {
		expected.value[k] = 51
	}

	hs, err := NewHexString(address)
	if err != nil {
		t.Error(err)
		return
	}

	result, err := ea.FromHexString(hs)

	if err != nil {
		t.Error(err)
		return
	}

	if *result != expected {
		t.Errorf("wrong address result, [Expected %v, Actual: %v]", expected, result)
	}

	if *ea != expected {
		t.Errorf("wrong address result, [Expected %v, Actual: %v]", expected, ea)
	}

	wrongaddress := "0x333333333333333333333333333333333333"

	hs,err = NewHexString(wrongaddress)

	if err != nil {
		t.Error(err)
		return
	}

	_, err = new(EtherAddress).FromHexString(hs)

	if err == nil {
		t.Errorf("should return error if wrong address size (%v)", wrongaddress)
	}
}
