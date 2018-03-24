package rpc

import "errors"

const RPCEndpointLocalHost = "http://localhost:8545"
const EndpointRemote = "http://ec2-54-93-108-139.eu-central-1.compute.amazonaws.com:8545"

var NotImplemented = errors.New("rpc method not implemented yet")
