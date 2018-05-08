package rpc

import (
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
)

const (
	MethodWeb3ClientVersion = "web3_clientVersion"
	MethodWeb3Sha3          = "web3_sha3"
)

type Web3 struct {
	client *Client
}

/*
	rpc method: "web3_clientVersion"
	returns string, the current Client version.
 */
func (web3 Web3) ClientVersion() (string, error) {
	return web3.client.RequestString(MethodWeb3ClientVersion)
}

/*
	rpc method: "web3_sha3"
	returns string, Keccak-256 (not the standardized SHA3-256) of the given data.
 */
func (web3 Web3) Sha3(b []byte) (string, error) {
	return web3.client.RequestString(MethodWeb3Sha3, rpctypes.ByteToHex(b))
}
