package rpcutils

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
)

func SignatureToTopic(signature string) string {
	method := []byte(signature)
	result := crypto.Keccak256(method)
	return common.Bytes2Hex(result)
}

func Signature2MethodId(signature string) (string, error) {
	method := []byte(signature)
	result := crypto.Keccak256(method)
	if len(result) < 4 {
		return "", fmt.Errorf("keccak doesnt have at least length 4 for %v -> %v", signature, result)
	}
	return common.Bytes2Hex(result[0:4]), nil
}
