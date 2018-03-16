package rpc

import (
	"testing"
	"strings"
)

var client = NewRPCClient(InfuraEndpoint)

func TestRPCClient_RequestWeb3ClientVersion(t *testing.T) {
	expectedVersionStart := "Geth/v"
	version, err := client.Web3.ClientVersion()

	if err != nil {
		t.Error(err)
		return
	}

	if !strings.Contains(version, expectedVersionStart) {
		t.Errorf("wrong Client version [Should Contain: %v, Actual: %v]", expectedVersionStart, version)
	}
}

func TestRPCClient_RequestWeb3Sha3(t *testing.T) {
	bytes := []byte("unlockKyc(address)")
	expectedKeccak := "0xf8db5d07f0ddf02621381eb3116de7cb49d036120a615265528d71f5e941ba3b"
	keccak, err := client.Web3.Sha3(bytes)

	if err != nil {
		t.Error(err)
		return
	}

	if expectedKeccak != keccak {
		t.Errorf("wrong keccak calculation [Expected: %v, Actual: %v]", expectedKeccak, keccak)
	}
}
