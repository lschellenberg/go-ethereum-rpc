package types

import (
	"math/big"
	"bytes"
	"strings"
)

type EtherValue struct {
	value    big.Int
	decimals int
}

func NewEtherValue(decimals ...int) *EtherValue {
	ev := new(EtherValue)
	if len(decimals) == 1 {
		ev.decimals = decimals[0]
	}
	return ev
}

func (ev *EtherValue) FromHexString(hex string) (*EtherValue, error) {
	r, err := HexToBigInt(hex)
	if err != nil {
		return nil, err
	}
	ev.value = *r
	return ev, nil
}

func (ev *EtherValue) FromBigInt(bi *big.Int) *EtherValue {
	ev.value = *bi
	return ev
}

func (ev *EtherValue) BigInt() *big.Int {
	return &ev.value
}

func (ev1 *EtherValue) IsEqual(ev2 *EtherValue) bool {
	return ev1.value.Cmp(ev2.BigInt()) == 0
}

func (ev *EtherValue) String() string {
	bi := ev.value
	s := bi.String()
	var buffer bytes.Buffer

	if len(s) > 18 {
		for i := 0; i < len(s); i++ {
			if i == len(s)-18 {
				buffer.WriteString(".")
			}
			buffer.WriteByte(s[i])
		}
	}

	if len(s) <= 18 {
		buffer.WriteString("0.")
		for i := 0; i < 18-len(s); i++ {
			buffer.WriteString("0")
		}
		for i := 0; i < len(s); i++ {
			buffer.WriteByte(s[i])
		}
	}

	return strings.TrimRight(buffer.String(), "0")
}

func (ev *EtherValue) Hash() string {
	return ev.HexString().Hash()
}

func (ev *EtherValue) HexString() (*HexString) {
	return new(HexString).FromBytes(ev.value.Bytes())
}

func EtherValueZero() *EtherValue {
	return new(EtherValue).FromBigInt(new(big.Int))
}

func EtherValueOne() *EtherValue {
	return new(EtherValue).FromBigInt(big.NewInt(1000000000000000000))
}
