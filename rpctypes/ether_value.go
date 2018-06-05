package rpctypes

import (
	"math/big"
	"bytes"
	"strings"
	"fmt"
	"strconv"
)

type EtherValue struct {
	value    big.Int
	decimals int
}

func NewEtherValueFromBigInt(bi *big.Int, decimals ...int) *EtherValue {
	ev := new(EtherValue)
	if len(decimals) == 1 {
		ev.decimals = decimals[0]
	} else {
		ev.decimals = 18
	}
	ev.value = *bi
	return ev
}

func NewEtherValue(decimals ...int) *EtherValue {
	ev := new(EtherValue)
	if len(decimals) == 1 {
		ev.decimals = decimals[0]
	} else {
		ev.decimals = 18
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

func (ev *EtherValue) Float64() (float64, error) {
	return strconv.ParseFloat(ev.String(), 64)
}

func (ev *EtherValue) FromBigIntString(intString string) (*EtherValue, error) {
	if ev.decimals == 0 {
		ev.decimals = 18
	}

	bi, ok := new(big.Int).SetString(intString, 10)
	if !ok {
		return nil, fmt.Errorf("couldn't convert %v to big.Int", intString)
	}
	ev.value = *bi
	return ev, nil
}
func (ev *EtherValue) FromFloat64String(floatString string) (*EtherValue, error) {
	if ev.decimals == 0 {
		ev.decimals = 18
	}

	bb := bytes.Buffer{}

	trimmed := strings.TrimLeft(floatString, "0")

	// no decimals
	if !strings.Contains(trimmed, ".") {

		bb.WriteString(trimmed)
		for i := 0; i < ev.decimals; i++ {
			bb.WriteString("0")
		}
	} else {
		a := strings.Split(trimmed, ".")
		if len(a) == 1 {
			bb.WriteString(a[0])
		} else {

			digit, decimal := a[0], a[1]
			pads := ev.decimals - len(decimal)

			bb.WriteString(digit)
			bb.WriteString(decimal)

			for i := 0; i < pads; i++ {
				bb.WriteString("0")
			}
		}
	}

	sv := bb.String()
	bi, ok := new(big.Int).SetString(sv, 10)

	if !ok {
		return nil, fmt.Errorf("couldnt convert %v to bi.Int", sv)
	}
	return ev.FromBigInt(bi), nil
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
	decimals := 18
	bi := ev.value
	s := bi.String()
	var buffer bytes.Buffer

	if len(s) > decimals {
		for i := 0; i < len(s); i++ {
			if i == len(s)-decimals {
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

	tr := strings.TrimRight(buffer.String(), "0")
	if tr[len(tr)-1] == '.' {
		return strings.TrimRight(tr, ".")
	}

	return tr
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

func (ev1 *EtherValue) Add(ev2 *EtherValue) *EtherValue {
	sum := big.NewInt(0).Add(ev1.BigInt(), ev2.BigInt())

	return new(EtherValue).FromBigInt(sum)
}

func (ev1 *EtherValue) Sub(ev2 *EtherValue) *EtherValue {
	difference := big.NewInt(0).Sub(ev1.BigInt(), ev2.BigInt())

	return new(EtherValue).FromBigInt(difference)
}
