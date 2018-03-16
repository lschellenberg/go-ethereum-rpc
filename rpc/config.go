package rpc

import "errors"

const RPCEndpointLocalHost = "http://localhost:8545"

var NotImplemented = errors.New("rpc method not implemented yet")