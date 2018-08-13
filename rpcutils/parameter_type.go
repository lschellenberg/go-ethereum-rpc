package rpcutils

import (
	"fmt"
	"strings"
)

const (
	// uint, int: synonyms for uint256, int256 respectively (not to be used for computing the function selector).
	BaseTypeUInt = "uint" //unsigned integer type of M bits, 0 < M <= 256, M % 8 == 0. e.g. uint32, uint8, uint256.
	BaseTypeInt  = "int"  //two's complement signed integer type of M bits, 0 < M <= 256, M % 8 == 0.

	BaseTypeAddress = "address" //equivalent to uint160, except for the assumed interpretation and language typing.
	BaseTypeBool    = "bool"    // equivalent to uint8 restricted to the values 0 and 1

	// fixed, ufixed: synonyms for fixed128x19, ufixed128x19 respectively (not to be used for computing the function selector).
	BaseTypeFixed  = "fixed"  // fixed<M>x<N>: signed fixed-point decimal number of M bits, 0 < M <= 256, M % 8 ==0, and 0 < N <= 80, which denotes the value v as v / (10 ** N).
	BaseTypeUFixed = "ufixed" // ufixed<M>x<N>: unsigned variant of fixed<M>x<N>.

	BaseTypeBytes    = "bytes"    // bytes<M>: binary type of M bytes, 0 < M <= 32.
	BaseTypeFunction = "function" // equivalent to bytes24: an address, followed by a function selector
)

type ParameterType struct {
	Type        string
	ByteSize    int
	ArrayLength int
}

func CreateParameterType(t string, byteSize int, arrayLength int) (*ParameterType, error) {
	if byteSize%8 != 0 {
		return nil, fmt.Errorf("byteSize must be multiple of 8, actual %v", byteSize)
	}

	if t != BaseTypeUInt &&
		t != BaseTypeInt &&
		t != BaseTypeAddress &&
		t != BaseTypeBool &&
		t != BaseTypeFixed &&
		t != BaseTypeUFixed &&
		t != BaseTypeBytes &&
		t != BaseTypeFunction {
		return nil, fmt.Errorf("unrecognized parameter type %v", t)
	}

	return &ParameterType{t, byteSize, arrayLength}, nil
}

func ParameterTypeFrom(paramName string) {
	if strings.Contains(paramName,BaseTypeUInt) {

	}
}