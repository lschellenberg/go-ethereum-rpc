package main

import (
	"github.com/Leondroids/go-ethereum-rpc/rpc"
	"github.com/Leondroids/go-ethereum-rpc/types"
	"log"
)

func main() {
	client := rpc.NewRPCClient(rpc.InfuraEndpoint)

	result, err := client.Eth.GetBalance("0xd780ae2bf04cd96e577d3d014762f831d97129d0", types.QuantityLatest())

	if err != nil {
		panic(err)
	}

	log.Println(result.String())
}
