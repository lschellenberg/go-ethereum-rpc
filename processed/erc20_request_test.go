package processed

import (
	"testing"
	"github.com/Leondroids/go-ethereum-rpc/rpc"
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
	"log"
)

func TestRequestERC20Transfer(t *testing.T) {
	client := rpc.NewRPCClient(rpc.GCloudEndpoint)
	params := &ERC20TransferParam{
		FromBlock: rpctypes.QuantityBlock(5705956),
		ToBlock:   rpctypes.QuantityBlock(5730114),
		Address:   "0xd780ae2bf04cd96e577d3d014762f831d97129d0",
	}
	result, err := RequestERC20Transfer(params, client)

	if err != nil {
		t.Error(err)
	}

	log.Println(result)
}
