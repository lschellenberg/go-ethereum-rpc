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

/*
	rpc method: personal_sendTransaction
	Sends transaction and signs it in a single call. The account does not need to be unlocked to make this call, and will not be left unlocked after.

	curl --data '{"method":"personal_sendTransaction","params":[{"from":"0x00a329c0648769a73afac7f9381e08fb43dbea72","to":"0x7c9ef46acd73a8225c417ea1ba69a65a74be0d80","data":"0x0","value":"0x186a0"},"hunter2"],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
 */
func (personal *Personal) SendTransaction(sendTransactionParams *SendTransaction, password string) (*rpctypes.HexString, error) {
	return personal.client.RequestHexString(MethodPersonalSendTransaction, sendTransactionParams.ToMap(), password)
}
