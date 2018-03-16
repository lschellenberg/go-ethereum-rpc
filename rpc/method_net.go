package rpc

const (
	MethodNetListening = "net_listening"
	MethodNetPeerCount = "net_peerCount"
	MethodNetVersion   = "net_version"
)

type Net struct {
	client *Client
}

/*
	rpc method: "net_listening"
	returns true if client is actively listening for network connections.
 */
func (net *Net) NetListening() (bool, error) {
	return net.client.RequestBool(MethodNetListening)
}

/*
	rpc method: "net_peerCount"
	returns true if client is actively listening for network connections.
 */
func (net *Net) NetPeerCount() (int64, error) {
	return net.client.RequestInt64(MethodNetPeerCount)
}

/*
	rpc method: "net_version"
	returns the current network protocol version.
	"1": Ethereum Mainnet
	"2": Morden Testnet (deprecated)
	"3": Ropsten Testnet
	"4": Rinkeby Testnet
	"42": Kovan Testnet
 */
func (net *Net) NetVersion() (string, error) {
	return net.client.RequestString(MethodNetVersion)
}
