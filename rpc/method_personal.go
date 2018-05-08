package rpc

import (
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
)

const (
	MethodPersonalListAccounts    = "personal_listAccounts"
	MethodPersonalNewAccount      = "personal_newAccount"
	MethodPersonalSendTransaction = "personal_sendTransaction"
	MethodPersonalSignTransaction = "personal_signTransaction"
	MethodUnlockAccount           = "personal_unlockAccount"
	MethodPersonalSign            = "personal_sign"
	MethodPersnalECRecover        = "personal_ecRecover"
)

type Personal struct {
	client *Client
}

/*
	rpc method: personal_listAccounts
	Lists all stored accounts.

	curl --data '{"method":"personal_listAccounts","params":[],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
*/
func (personal *Personal) ListAccounts() ([]rpctypes.HexString, error) {
	return personal.client.RequestHexStringList(MethodPersonalListAccounts)
}

/*
	rpc method: personal_newAccount
	Creates new account.
	Note: it becomes the new current unlocked account. There can only be one unlocked account at a time.

	curl --data '{"method":"personal_newAccount","params":["hunter2"],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
 */
func (personal *Personal) NewAccount(password string) (*rpctypes.HexString, error) {
	return personal.client.RequestHexString(MethodPersonalNewAccount, password)
}
