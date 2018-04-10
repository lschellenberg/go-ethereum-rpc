#!/usr/bin/env bash

export infura=https://mainnet.infura.io/3l5dxBOP3wPspnRDdG1u
export parity=http://ec2-54-93-108-139.eu-central-1.compute.amazonaws.com:8545
echo $1
echo $2
export trdata='{"method":"eth_call","params":["'$1','$2'"],"id":1,"jsonrpc":"2.0"}'
echo $trdata
curl --data $trdata -H "Content-Type: application/json" -X POST $infura | jq .
