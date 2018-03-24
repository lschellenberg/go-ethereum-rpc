package types

import (
	"fmt"
	"bytes"
	"encoding/hex"
)

const EtherAddressLength = 20

type EtherAddress struct {
	value [EtherAddressLength]byte
}

func (ea *EtherAddress) ShortFormat() string {
	return ea.String()[:8] + "..."
}

func (ea *EtherAddress) String() string {
	s := hex.EncodeToString(ea.Bytes())
	return "0x" + s
}

func (ea *EtherAddress) Hash() string {
	return ea.String()
}

func (ea *EtherAddress) HexString() *HexString {
	return NewHexStringFromBytes(ea.Bytes())
}

func (ea *EtherAddress) Bytes() []byte {
	newb := make([]byte, EtherAddressLength)
	copy(newb[:], ea.value[:])
	return newb
}

func (ea *EtherAddress) FromBytes(b []byte) (*EtherAddress, error) {
	if len(b) != EtherAddressLength {
		return nil, fmt.Errorf("%v not a valid ether address", b)
	}
	copy(ea.value[:], b[:])
	return ea, nil
}

func (ea *EtherAddress) FromHexString(hex *HexString) (*EtherAddress, error) {
	return ea.FromBytes(hex.Bytes())
}

func (ea *EtherAddress) FromString(s string) (*EtherAddress, error) {
	b, err := NewHexString(s)

	if err != nil {
		return nil, err
	}
	return ea.FromBytes(b.Bytes())
}

func (ea *EtherAddress) FromStringOrNull(s string) (*EtherAddress, error) {
	if len(s) != EtherAddressLength*2+2 {
		s = "0x0000000000000000000000000000000000000000"
	}
	return ea.FromString(s)
}

func (ea1 *EtherAddress) IsEqual(ea2 *EtherAddress) bool {
	return bytes.Compare(ea1.Bytes(), ea2.Bytes()) == 0
}
func (address *EtherAddress) From32ByteString(value string) (*EtherAddress, error) {
	if len(value) != 64+2 {
		return nil, fmt.Errorf("%v is not a 64 byte hex string", value)
	}

	return address.FromString(value[26:])
}
