package rpcutils

import (
	"fmt"
	"github.com/Leondroids/go-ethereum-rpc/types"
)

var FPTAddress = FunctionParamType{"address", false}
var FPTUInt256 = FunctionParamType{"uint256", false}
var FPTString = FunctionParamType{"string", true}
var FPTBool = FunctionParamType{"bool", false}

type FunctionParamType struct {
	Type      string
	IsDynamic bool
}

type FunctionParam struct {
	Type  string
	Value interface{}
}

func NewFunctionParamType(t string) *FunctionParamType {
	switch t {
	case FPTString.Type:
		return &FPTString
	case FPTAddress.Type:
		return &FPTAddress
	case FPTBool.Type:
		return &FPTBool
	case FPTUInt256.Type:
		return &FPTUInt256
	}
	return nil
}

type FunctionSignature struct {
	Params []FunctionParamType
}

func NewFunctionSignature(values ...FunctionParamType) *FunctionSignature {
	params := make([]FunctionParamType, 0)

	for _, v := range values {
		params = append(params, v)
	}

	return &FunctionSignature{
		Params: params,
	}
}

func (fs *FunctionSignature) DecodeFunctionInput(input string) ([]FunctionParam, error) {

	if len(input) < 10+64*fs.Len() {
		return nil, fmt.Errorf("input length should be at least methodsignature 10 + %v * 64 (32byte in hex) long", fs.Len())
	}

	hsinput, err := types.NewHexString(input)

	if err != nil {
		return nil, err
	}

	return fs.DecodeFunctionInputFromHex(hsinput)
}

func (fs *FunctionSignature) DecodeFunctionInputFromHex(input *types.HexString) ([]FunctionParam, error) {
	fp := make([]FunctionParam, len(fs.Params))

	head := fs.ReadHead(input)
	body, err := fs.ReadBody(input)
	if err != nil {
		return nil, err
	}
	headLength := len(fs.Params) * 32

	for k, h := range head {
		p := fs.Params[k]
		pType := p.Type

		//log.Println("Type: ", pType)
		//log.Println("Value: ", h.Hash())
		//log.Println("..........................", h.Hash())
		if !p.IsDynamic {
			value, err := fromNonDynamicValue(&h, pType)

			if err != nil {
				return nil, err
			}

			fp[k] = FunctionParam{pType, value}
		} else {
			f := FunctionParam{pType, ""}
			location := int(h.Int64()) - headLength

			switch pType {
			case FPTString.Type:
				s, err := decodeString(body, location)
				if err != nil {
					return nil, err
				}
				f = FunctionParam{pType, s}
			}

			fp[k] = f
		}
	}

	return fp, nil
}

func (fs *FunctionSignature) Len() int {
	return len(fs.Params)
}

func (fs *FunctionSignature) ReadHead(input *types.HexString) []types.HexString {
	b := input.Bytes()[4:]

	result := make([]types.HexString, fs.Len())

	for k := range result {
		from := k * 32
		to := 32 + from
		result[k] = *new(types.HexString).FromBytes(b[from:to])
	}

	return result
}

func (fs *FunctionSignature) ReadBody(input *types.HexString) (*types.HexString, error) {

	b := input.Bytes()[4+fs.Len()*32:]

	if len(b)%32 != 0 {
		return nil, fmt.Errorf("function input body is not factor of 32 bytes")
	}

	return new(types.HexString).FromBytes(b), nil
}

func fromNonDynamicValue(val *types.HexString, paramType string) (interface{}, error) {
	switch paramType {
	case FPTAddress.Type:
		return new(types.EtherAddress).From32ByteHex(val)
	case FPTUInt256.Type:
		return val.BigInt(), nil
	case FPTBool.Type:
		return val.Int64() > 0, nil
	}

	return "", nil
}

func decodeString(body *types.HexString, location int) (string, error) {
	if len(body.Bytes()) < location {
		return "", fmt.Errorf("function input body too short")
	}

	dataLengthPart := body.Bytes()[location:location+32]
	length := int(new(types.HexString).FromBytes(dataLengthPart).Int64())
	dataPart := body.Bytes()[location+32:location+32+length]
	result := new(types.HexString).FromBytes(dataPart).Text()
	return result, nil
}
