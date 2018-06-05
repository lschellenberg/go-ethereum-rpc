package rpc

import (
	"testing"
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
)

func TestNewFilterParams_ToMap(t *testing.T) {
	params := CreateNewFilterParams("address", rpctypes.QuantityLatest(), rpctypes.QuantityLatest(), CreateNewFilterTopics([]string{"t11", "t12"}, []string{"t21"}, []string{}))
	t.Errorf("%+v", params.ToMap())
}
