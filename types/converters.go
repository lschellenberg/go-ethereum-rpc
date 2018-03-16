package types

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"math/big"
)

/*
	returns a a byte array into a hex string in format '0x....'
 */
func ByteToHex(b []byte) string {
	return fmt.Sprintf("0x%v", hex.EncodeToString(b))
}

func HexToByte(s string) ([]byte, error) {
	if s == "0x0" {
		return []byte{0}, nil
	}
	if len(s) < 2 {
		return nil, fmt.Errorf("not a hex string %v", s)
	}
	return hex.DecodeString(s[2:])
}

func HexToBigInt(s string) (*big.Int, error) {
	if len(s) < 2 {
		return nil, fmt.Errorf("%v cannot be bigint", s)
	}
	ss := s[2:]
	// pad if not even
	if len(ss)%2 != 0 {
		ss = "0" + ss
	}

	b, err := hex.DecodeString(ss)

	if err != nil {
		return nil, err
	}

	return new(big.Int).SetBytes(b), nil
}

/*
	converts an interface list to a string list
 */
func InterfaceListToStringList(i []interface{}) ([]string, error) {
	s := make([]string, len(i))

	for k, v := range i {
		r, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("couldn't convert %v to string", reflect.TypeOf(v))
		}
		s[k] = r
	}

	return s, nil
}
