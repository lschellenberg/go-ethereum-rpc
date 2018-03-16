package rpc

import "testing"

func TestNet_NetListening(t *testing.T) {
	listening, err := NewRPCClient(RPCEndpointLocalHost).Net.NetListening()

	if err != nil {
		t.Error(err)
		return
	}

	if listening != true {
		t.Errorf("net should listen")
	}
}

func TestNet_NetPeerCount(t *testing.T) {
	count, err := NewRPCClient(RPCEndpointLocalHost).Net.NetPeerCount()

	if err != nil {
		t.Error(err)
		return
	}

	if count < 5 {
		t.Errorf("net should listen at least 5 peers, but: %v", count)
	}
}

func TestNet_NetVersion(t *testing.T) {
	expectedVersion := "1"
	version, err := NewRPCClient(RPCEndpointLocalHost).Net.NetVersion()

	if err != nil {
		t.Error(err)
		return
	}

	if version != expectedVersion {
		t.Errorf("net version is wrong, [Expected: %v, Actual: %v]", expectedVersion, version)
	}
}
