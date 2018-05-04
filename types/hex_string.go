package types

import (
	"strings"
	"encoding/hex"
	"fmt"
	"math/big"
	"encoding/binary"
	"bytes"
)

type HexString struct {
	value []byte
}

func NewHexString(s string) (*HexString, error) {
	temp := s
	if strings.Contains(s, "0x") {
		temp = s[2:]
	}

	// normalize
	if len(temp)%2 != 0 {
		temp = "0" + temp
	}

	b, err := hex.DecodeString(temp)

	if err != nil {
		return nil, err
	}

	return &HexString{
		value: b,
	}, nil
}

func NewHexStringFromBytes(b []byte) *HexString {
	return &HexString{
		value: b,
	}
}

func (hs HexString) Hash() string {
	// Ethereum Nodes are representing 0 as "0x0"
	if len(hs.value) == 0 {
		return "0x0"
	}

	s := hex.EncodeToString(hs.value)

	// Ethereum Nodes are representing 0 as "0x0"
	if s == "00" {
		return "0x0"
	}

	return "0x" + s
}

func (hs HexString) String() string {
	// Ethereum Nodes are representing 0 as "0x0"
	if len(hs.value) == 0 {
		return "0x0"
	}
	s := hex.EncodeToString(hs.value)

	// trim
	pos := 0
	for k, v := range s {
		if v != '0' {
			pos = k
			break
		}
	}

	// Ethereum Nodes are representing 0 as "0x0"
	if s == "00" {
		return "0x0"
	}

	return "0x" + s[pos:]
}

func (hs HexString) Bytes() []byte {
	return hs.value
}
func (hs HexString) Plain() string {
	return hex.EncodeToString(hs.Bytes())
}
// used to display text, from ascii 7 on there are meaningful chars,
// TODO sounds a bit arb.
func (hs HexString) Text() string {
	b := hs.Bytes()
	rb := make([]byte, 0)
	for _, v := range b {
		if v > 7 {
			rb = append(rb, v)
		}
	}

	return string(rb)
}

func (hs *HexString) FromBytes(b []byte) *HexString {
	hs.value = b
	return hs
}

func (hs HexString) BigInt() *big.Int {
	return new(big.Int).SetBytes(hs.value)
}

func (hs *HexString) FromInt64(i int64) *HexString {
	b := make([]byte, 8)

	binary.BigEndian.PutUint64(b, uint64(i))

	// trim byte array
	pos := 0
	for k, v := range b {
		if v > 0 {
			pos = k
			break
		}
	}

	return hs.FromBytes(b[pos:])
}

func (hs HexString) Int64() int64 {
	b := hs.Bytes()

	if len(b) == 0 {
		return 0
	}

	if len(b) > 8 {
		temp := b[len(b)-8:]
		b = temp
	}
	// pad if necessary
	if len(b) < 8 {
		temp := make([]byte, 8)

		for i := 8 - len(b); i < 8; i++ {
			temp[i] = b[i+len(b)-8]
		}

		b = temp
	}
	return int64(binary.BigEndian.Uint64(b))
}

func ToHexStringList(l []string) ([]HexString, error) {
	result := make([]HexString, len(l))

	for k, v := range l {
		hs, err := NewHexString(v)
		if err != nil {
			return nil, err
		}
		result[k] = *hs
	}

	return result, nil
}

func CompareHexStringList(s1 []HexString, s2 []HexString) error {
	if len(s1) != len(s2) {
		return fmt.Errorf("wrong sizes in HexString List %v - %v", len(s1), len(s2))
	}

	for k, v := range s1 {

		if bytes.Compare(s2[k].value, v.value) != 0 {
			return fmt.Errorf("not eqal at position %v", k)
		}
	}

	return nil
}

func HexStringListFromString(s []string) ([]HexString, error) {
	hsl := make([]HexString, len(s))

	for k, v := range s {
		hs, err := NewHexString(v)
		if err != nil {
			return nil, err
		}
		hsl[k] = *hs
	}

	return hsl, nil
}

func HexStringListToStringList(s []HexString) ([]string) {
	hsl := make([]string, len(s))

	for k, v := range s {
		hsl[k] = v.String()
	}

	return hsl
}

func (hs1 HexString) IsEqual(hs2 *HexString) bool {
	return bytes.Compare(hs1.value, hs2.value) == 0
}

func (hs1 *HexString) Concat(hs2 *HexString) *HexString {
	nb := make([]byte, len(hs1.value)+len(hs2.value))

	for k, v := range hs1.value {
		nb[k] = v
	}

	for k, v := range hs2.value {
		nb[k+len(hs1.value)] = v
	}

	return NewHexStringFromBytes(nb)
}

func (hs1 *HexString) PadTo(length int) (*HexString) {

	if len(hs1.value) == length {
		return NewHexStringFromBytes(hs1.value)
	}

	nb := make([]byte, length)
	paddingLength := length - len(hs1.value)

	if len(hs1.value) > length {
		copy(nb, hs1.value[-paddingLength:])
	} else {
		copy(nb[paddingLength:], hs1.value)
	}

	return NewHexStringFromBytes(nb)
}
