package rpc

import (
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
	"fmt"
	"encoding/json"
)

func (client *Client) RequestEtherLogList(method string, params ...interface{}) ([]rpctypes.EtherLog, error) {
	response, err := checkRPCError(client.Call(method, params...))

	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("response returned without error but no logs found for %v", params)
	}

	js, err := json.Marshal(response.Result)

	if err != nil {
		return nil, err
	}

	return new(rpctypes.EtherLogRaw).FromJSONArray(js)
}
