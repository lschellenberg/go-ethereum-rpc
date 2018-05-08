package rpcutils

import (
	"fmt"
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
	"log"
)

const (
	EthereumStandardByteLength = 32
)

var FPTAddress = FunctionParamType{"address", false, 0}
var FPTBool = FunctionParamType{"bool", false, 0}

var FPTUInt8 = FunctionParamType{"uint8", false, 0}
var FPTUInt16 = FunctionParamType{"uint16", false, 0}
var FPTUInt32 = FunctionParamType{"uint32", false, 0}
var FPTUInt64 = FunctionParamType{"uint64", false, 0}
var FPTUInt128 = FunctionParamType{"uint128", false, 0}
var FPTUInt256 = FunctionParamType{"uint256", false, 0}

var FPTBytes = FunctionParamType{"bytes", true, 0}
var FPTString = FunctionParamType{"string", true, 0}

type FunctionParamType struct {
	Type        string
	IsDynamic   bool
	ArrayLength int
}

type FunctionParam struct {
	Type  string
	Value interface{}
}

func NewFunctionParamType(t string, arraylength int) *FunctionParamType {
	fpt := &FunctionParamType{}
	switch t {
	case FPTString.Type:
		fpt = &FPTString
	case FPTBytes.Type:
		fpt = &FPTBytes
	case FPTAddress.Type:
		fpt = &FPTAddress
	case FPTBool.Type:
		fpt = &FPTBool
	case FPTUInt8.Type:
		fpt = &FPTUInt8
	case FPTUInt16.Type:
		fpt = &FPTUInt16
	case FPTUInt32.Type:
		fpt = &FPTUInt32
	case FPTUInt64.Type:
		fpt = &FPTUInt64
	case FPTUInt128.Type:
		fpt = &FPTUInt128
	case FPTUInt256.Type:
		fpt = &FPTUInt256
	}

	fpt.ArrayLength = arraylength

	return fpt
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

func (fs *FunctionSignature) DecodeEventData(input string) ([]FunctionParam, error) {

	log.Println("InputLength: ",len(input))
	log.Println("parameter length: ", 64*fs.Len())
	if len(input) < 2 + 64*fs.Len() {
		return nil, fmt.Errorf("input length should be at least methodsignature 10 + %v * 64 (32byte in hex) long", fs.Len())
	}

	hsinput, err := rpctypes.NewHexString(input)

	if err != nil {
		return nil, err
	}

	return fs.DecodeFunctionInputFromHex(hsinput)
}
func (fs *FunctionSignature) DecodeFunctionInput(input string) ([]FunctionParam, error) {

	if len(input) < 10+64*fs.Len() {
		return nil, fmt.Errorf("input length should be at least methodsignature 10 + %v * 64 (32byte in hex) long", fs.Len())
	}

	hsinput, err := rpctypes.NewHexString(input)

	if err != nil {
		return nil, err
	}
	b := hsinput.Bytes()


	return fs.DecodeFunctionInputFromHex(rpctypes.NewHexStringFromBytes(b[4:]))
}

func (fs *FunctionSignature) DecodeFunctionInputFromHex(input *rpctypes.HexString) ([]FunctionParam, error) {
	fp := make([]FunctionParam, len(fs.Params))

	//log.Println("input: ", input.Plain())

	head := fs.ReadHead(input)
	//log.Println("Head: ")
	//for k,v := range head {
	//	log.Printf("%v: %v",k, v.Hash())
	//}

	body, err := fs.ReadBody(input)
	if err != nil {
		return nil, err
	}
	headLength := len(fs.Params) * EthereumStandardByteLength

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
			case FPTString.Type:
				s, err := decodeBytes(body, location)
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

func (fs *FunctionSignature) ReadHead(input *rpctypes.HexString) []rpctypes.HexString {
	b := input.Bytes()

	result := make([]rpctypes.HexString, fs.Len())

	for k := range result {
		from := k * EthereumStandardByteLength
		to := EthereumStandardByteLength + from
		result[k] = *new(rpctypes.HexString).FromBytes(b[from:to])
	}

	return result
}

func (fs *FunctionSignature) ReadBody(input *rpctypes.HexString) (*rpctypes.HexString, error) {

	b := input.Bytes()[fs.Len()*EthereumStandardByteLength:]

	if len(b)%EthereumStandardByteLength != 0 {
		return nil, fmt.Errorf("function input body is not factor of EthereumStandardByteLength bytes")
	}

	return new(rpctypes.HexString).FromBytes(b), nil
}

func fromNonDynamicValue(val *rpctypes.HexString, paramType string) (interface{}, error) {
	switch paramType {
	case FPTAddress.Type:
		return new(rpctypes.EtherAddress).From32ByteHex(val)
	case FPTBytes.Type:
		return val.Bytes(), nil
	case FPTUInt8.Type:
		return int(val.BigInt().Int64()), nil
	case FPTUInt16.Type:
		return int(val.BigInt().Int64()), nil
	case FPTUInt32.Type:
		return int(val.BigInt().Int64()), nil
	case FPTUInt64.Type:
		return val.BigInt().Int64(), nil
	case FPTUInt128.Type:
		return val.BigInt(), nil
	case FPTUInt256.Type:
		return val.BigInt(), nil
	case FPTBool.Type:
		return val.Int64() > 0, nil
	}

	return "", nil
}

func decodeString(body *rpctypes.HexString, location int) (string, error) {
	if len(body.Bytes()) < location {
		return "", fmt.Errorf("function input body too short")
	}

	dataLengthPart := body.Bytes()[location:location+EthereumStandardByteLength]
	length := int(new(rpctypes.HexString).FromBytes(dataLengthPart).Int64())
	dataPart := body.Bytes()[location+EthereumStandardByteLength:location+EthereumStandardByteLength+length]
	result := new(rpctypes.HexString).FromBytes(dataPart).Text()
	return result, nil
}

func decodeBytes(body *rpctypes.HexString, location int) ([]byte, error) {
	if len(body.Bytes()) < location {
		return nil, fmt.Errorf("function input body too short")
	}

	dataLengthPart := body.Bytes()[location:location+EthereumStandardByteLength]
	length := int(new(rpctypes.HexString).FromBytes(dataLengthPart).Int64())
	dataPart := body.Bytes()[location+EthereumStandardByteLength:location+EthereumStandardByteLength+length]

	return dataPart, nil
}
