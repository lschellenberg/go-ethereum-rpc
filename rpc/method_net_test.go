package rpc

import "testing"

func TestNet_NetListening(t *testing.T) {
	listening, err := NewRPCClient(config().address).Net.NetListening()

	if err != nil {
		t.Error(err)
		return
	}

	if listening != true {
		t.Errorf("net should listen")
	}
}

func TestNet_NetPeerCount(t *testing.T) {
	count, err := NewRPCClient(config().address).Net.NetPeerCount()

	if err != nil {
		t.Error(err)
		return
	}

	var expectedPeers int64 = 2

	if count < expectedPeers {
		t.Errorf("net should listen at least %v peers, but: %v", expectedPeers, count)
	}
}

func TestNet_NetVersion(t *testing.T) {
	expectedVersion := "1"
	version, err := NewRPCClient(config().address).Net.NetVersion()

	if err != nil {
		t.Error(err)
		return
	}

	if version != expectedVersion {
		t.Errorf("net version is wrong, [Expected: %v, Actual: %v]", expectedVersion, version)
	}
}
