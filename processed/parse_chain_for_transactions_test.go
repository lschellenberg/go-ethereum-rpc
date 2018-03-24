package processed

import (
	"testing"
	"github.com/Leondroids/go-ethereum-rpc/rpc"
)

type TestConfig struct {
	address string
}

func config() *TestConfig {
	return LocalhostConfig()
}

func InfuraConfig() *TestConfig {
	return &TestConfig{
		address: rpc.InfuraEndpoint,
	}
}

func LocalhostConfig() *TestConfig {
	return &TestConfig{
		address: rpc.RPCEndpointLocalHost,
	}
}

func TestParseTransactionsFromChainByFrom(t *testing.T) {
	err := ParseTransactionsFromChainByTo(IconomiTokenAddress, 2400000, 25000001, &rpc.NewRPCClient(config().address).Eth)

	if err != nil {
		t.Error(err)
		return
	}
	t.Error("")

}
