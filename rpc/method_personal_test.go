package rpc

import (
	"testing"
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
)

func TestPersonal_ListAccounts(t *testing.T) {
	accounts, err := NewRPCClient(config().address).Personal.ListAccounts()

	if err != nil {
		t.Error(err)
		return
	}

	expected1, _ := rpctypes.NewHexString("0x44a139cc0aed5eb5dbc6838b284fb051cad72dcb")
	expected2, _ := rpctypes.NewHexString("0xe96f31db85aa516b5a6ab2973d333f0406ddcb9b")

	err = rpctypes.CompareHexStringList([]rpctypes.HexString{*expected1, *expected2}, accounts)
	if err != nil {
		t.Error(err)
		return
	}

}

func TestPersonal_NewAccount(t *testing.T) {
	// stop testing unless on private test net
	//test := "password"
	//_, err := NewRPCClient(config().address).Personal.NewAccount(test)
	//
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
}
