package rpc

import (
	"testing"
	"math/big"
	"errors"
	"github.com/Leondroids/gox"
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
)

type TestConfig struct {
	address         string
	protocolVersion int64
	coinbase        string
	hashrate        int64
	mining          bool
}

func config() *TestConfig {
	return InfuraConfig()
}

func InfuraConfig() *TestConfig {
	return &TestConfig{
		address:         InfuraEndpoint,
		protocolVersion: 63,
	}
}

func LocalhostConfig() *TestConfig {
	return &TestConfig{
		address:         RPCEndpointLocalHost,
		protocolVersion: 99,
		coinbase:        "0x44a139cc0aed5eb5dbc6838b284fb051cad72dcb",
		hashrate:        0,
		mining:          false,
	}
}

func GCloudConfig() *TestConfig {
	return &TestConfig{
		address:         GCloudEndpoint,
		protocolVersion: 99,
		coinbase:        "0x0000000000000000000000000000000000000000",
		hashrate:        0,
		mining:          false,
	}
}

func TestRPCClient_EthAccounts(t *testing.T) {
	expectedStringList := []string{}
	s, err := NewRPCClient(config().address).Eth.Accounts()

	if err != nil {
		t.Error(err)
		return
	}

	if !gox.CompareStringList(expectedStringList, s) {
		t.Errorf("wrong eth accounts [Expected: %v, Actual: %v]", expectedStringList, s)
	}
}

func TestRPCClient_EthBlockNumber(t *testing.T) {
	var expectedBlockNumberMin int64 = 3000000
	block, err := NewRPCClient(config().address).Eth.BlockNumber()

	if err != nil {
		t.Error(err)
		return
	}

	if expectedBlockNumberMin > block {
		t.Errorf("wrong block number [At Least: %v, Actual: %v]", expectedBlockNumberMin, block)
	}
}

func TestEth_Call(t *testing.T) {
	params := new(EthCallParams).ToContract("0xd780ae2bf04cd96e577d3d014762f831d97129d0", "0x115976c4")
	result, err := NewRPCClient(config().address).Eth.Call(params, nil)

	if err != nil {
		t.Error(err)
		return
	}

	expectedResult := "0x7e49e51d4e30054915f11bd72c6c02f7423c1f40"

	if expectedResult != result.String() {
		t.Errorf("wrong eth_call response [Expected: %v, Actual: %v]", expectedResult, result.String())
	}
}

func TestEth_CallWithQuantity(t *testing.T) {
	params := new(EthCallParams).ToContract("0xd780ae2bf04cd96e577d3d014762f831d97129d0", "0xacc99bb7")
	result, err := NewRPCClient(config().address).Eth.Call(params, rpctypes.QuantityBlock(4900000))

	if err != nil {
		t.Error(err)
		return
	}

	expectedResult := "0x7e49e51d4e30054915f11bd72c6c02f7423c1f40"

	value, err := new(rpctypes.EtherValue).FromHexString(result.String())

	if err != nil {
		t.Error(err)
		return
	}

	t.Error("Value: ", value)
	if expectedResult != result.String() {
		t.Errorf("wrong eth_call response [Expected: %v, Actual: %v]", expectedResult, result.String())
	}
}

func TestEth_Coinbase(t *testing.T) {
	coinbase, err := NewRPCClient(config().address).Eth.Coinbase()

	if err != nil {
		t.Error(err)
		return
	}

	expectedCoinbase := config().coinbase

	if coinbase.String() != expectedCoinbase {
		t.Errorf("wrong coinbase [Expected: %v, Actual: %v]", expectedCoinbase, coinbase)
	}
}

func TestEth_GasPrice(t *testing.T) {
	gasprice, err := NewRPCClient(config().address).Eth.GasPrice()

	if err != nil {
		t.Error(err)
		return
	}

	if gasprice.IsEqual(rpctypes.EtherValueZero()) {
		t.Errorf("gas price cant be 0")
	}

	if gasprice.BigInt().Cmp(rpctypes.EtherValueOne().BigInt()) != -1 {
		t.Errorf("gas price cant be more than 1")
	}
}

func TestRPCClient_GetBalance(t *testing.T) {
	balance, err := NewRPCClient(config().address).Eth.GetBalance("0x6Da385A99A8799c986F215dDc14e8028eF0b8baF", rpctypes.QuantityLatest())

	if err != nil {
		t.Error(err)
		return
	}

	expected, ok := big.NewInt(0).SetString("24459365180000000000", 10)

	if !ok {
		t.Error(errors.New("couldn't convert big string int to big.Int"))
		return
	}

	bi := balance.BigInt()
	if bi.Cmp(expected) != 0 {
		t.Errorf("wrong balance [Expected: %v, Actual: %v]", expected.String(), bi.String())
	}
}

func TestRPCClient_GetBlockByHashWithTransactionHash(t *testing.T) {
	expectedBlock := getBlock_4ad331()
	result, err := NewRPCClient(config().address).Eth.GetBlockByHash(expectedBlock.Hash.String(), false)

	if err != nil {
		t.Error(err)
		return
	}

	if err = expectedBlock.Compare(result); err != nil {
		t.Error(err)
	}
}

func TestRPCClient_GetBlockByHashWithTransactionFull(t *testing.T) {
	expectedBlock := getBlock_4ad331()
	expectedBlock.TransactionsFull = HashToEtherTransactionFull(expectedBlock.Transactions)
	expectedBlock.Transactions = nil

	result, err := NewRPCClient(config().address).Eth.GetBlockByHash(expectedBlock.Hash.String(), true)

	if err != nil {
		t.Error(err)
		return
	}

	if err = expectedBlock.Compare(result); err != nil {
		t.Error(err)
	}

	// pick one transaction

	expectedTransaction := getTransaction_99192()
	actual := result.TransactionsFull[85]

	if err := expectedTransaction.Compare(&actual); err != nil {
		t.Error(err)
	}
}

func TestRPCClient_GetBlockByNumberWithTransactionHash(t *testing.T) {
	expectedBlock := getBlock_4ad331()
	result, err := NewRPCClient(config().address).Eth.GetBlockByNumber(expectedBlock.Number, false)

	if err != nil {
		t.Error(err)
		return
	}

	if err = expectedBlock.Compare(result); err != nil {
		t.Error(err)
	}
}

func TestRPCClient_GetBlockByNumberWithTransactionFull(t *testing.T) {
	expectedBlock := getBlock_4ad331()
	expectedBlock.TransactionsFull = HashToEtherTransactionFull(expectedBlock.Transactions)
	expectedBlock.Transactions = nil
	// 2 first full transaction

	result, err := NewRPCClient(config().address).Eth.GetBlockByNumber(expectedBlock.Number, true)

	if err != nil {
		t.Error(err)
		return
	}

	if err = expectedBlock.Compare(result); err != nil {
		t.Error(err)
	}

	// pick one transaction

	expectedTransaction := getTransaction_99192()
	actual := result.TransactionsFull[85]

	if err := expectedTransaction.Compare(&actual); err != nil {
		t.Error(err)
	}
}

// GetBlockTransactionCountByHash
func TestEth_GetBlockTransactionCountByHash(t *testing.T) {
	expectedBlock := getBlock_4ad331()

	result, err := NewRPCClient(config().address).Eth.GetBlockTransactionCountByHash(expectedBlock.Hash.String())

	if err != nil {
		t.Error(err)
		return
	}

	if result != int64(len(expectedBlock.Transactions)) {
		t.Errorf("wrong count [Expected: %v, Actual: %v]", len(expectedBlock.Transactions), result)
	}
}

// GetBlockTransactionCountByNumber
func TestEth_GetBlockTransactionCountByNumber(t *testing.T) {
	expectedBlock := getBlock_4ad331()

	result, err := NewRPCClient(config().address).Eth.GetBlockTransactionCountByNumber(expectedBlock.Number)

	if err != nil {
		t.Error(err)
		return
	}

	if result != int64(len(expectedBlock.Transactions)) {
		t.Errorf("wrong count [Expected: %v, Actual: %v]", len(expectedBlock.Transactions), result)
	}
}

// GetTransactionByBlockHashAndIndex
func TestEth_GetTransactionByBlockHashAndIndex(t *testing.T) {
	expected := getTransaction_99192()
	result, err := NewRPCClient(config().address).Eth.GetTransactionByBlockHashAndIndex(expected.BlockHash.String(), expected.TransactionIndex)

	if err != nil {
		t.Error(err)
		return
	}

	if err := expected.Compare(result); err != nil {
		t.Error(err)
	}
}

func TestEth_GetFilterLogs(t *testing.T) {

	p := CreateNewFilterParamsWithOneTopic("0xd780ae2bf04cd96e577d3d014762f831d97129d0", rpctypes.QuantityBlock(5730114), rpctypes.QuantityBlock(5730114), "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	filter, err := NewRPCClient(config().address).Eth.NewFilter(p)

	if err != nil {
		t.Error(err)
	}

	result, err := NewRPCClient(config().address).Eth.GetFilterLogs(filter.Hash())


	if err != nil {
		t.Error(err)
	}

	if len(result) != 1 {
		t.Error("expected length of return is 1")
	}

	ea, _ := new(rpctypes.EtherAddress).FromString("0xd780ae2bf04cd96e577d3d014762f831d97129d0")
	t1, _ := rpctypes.NewHexString("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	t2, _ := rpctypes.NewHexString("0x1ce7ae555139c5ef5a57cc8d814a867ee6ee33d8")
	t3, _ := rpctypes.NewHexString("0xf5a276aa1062f19453ef96e156ea868ee3c1fccf")
	data, _ := rpctypes.NewHexString("0x00000000000000000000000000000000000000000000000ad78ebc5ac6200000")
	bh, _ := rpctypes.NewHexString("0xaeeddb13df5d5a9b95a1748106286543e8c5115debf280bfdfb81c9efab5287a")
	th, _ := rpctypes.NewHexString("0x95edd0f567e4617284c5cd64c659546d33828fac59a5594ca7d3f3d56bb3c186")

	expectedLogs := &rpctypes.EtherLog{
		Address:          *ea.HexString(),
		Topics:           []rpctypes.HexString{*t1, *t2, *t3},
		Data:             *data,
		BlockNumber:      5730114,
		BlockHash:        *bh,
		TransactionHash:  *th,
		TransactionIndex: 50,
	}

	if err := expectedLogs.Compare(result[0]); err != nil {
		t.Error(err)
	}
}

func TestEth_GetLogs(t *testing.T) {
	p := CreateNewFilterParamsWithOneTopic("0xd780ae2bf04cd96e577d3d014762f831d97129d0", rpctypes.QuantityBlock(5730114), rpctypes.QuantityBlock(5730114), "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	result, err := NewRPCClient(config().address).Eth.GetLogs(p)

	if err != nil {
		t.Error(err)
	}

	if len(result) != 1 {
		t.Error("expected length of return is 1")
	}

	ea, _ := new(rpctypes.EtherAddress).FromString("0xd780ae2bf04cd96e577d3d014762f831d97129d0")
	t1, _ := rpctypes.NewHexString("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	t2, _ := rpctypes.NewHexString("0x1ce7ae555139c5ef5a57cc8d814a867ee6ee33d8")
	t3, _ := rpctypes.NewHexString("0xf5a276aa1062f19453ef96e156ea868ee3c1fccf")
	data, _ := rpctypes.NewHexString("0x00000000000000000000000000000000000000000000000ad78ebc5ac6200000")
	bh, _ := rpctypes.NewHexString("0xaeeddb13df5d5a9b95a1748106286543e8c5115debf280bfdfb81c9efab5287a")
	th, _ := rpctypes.NewHexString("0x95edd0f567e4617284c5cd64c659546d33828fac59a5594ca7d3f3d56bb3c186")

	expectedLogs := &rpctypes.EtherLog{
		Address:          *ea.HexString(),
		Topics:           []rpctypes.HexString{*t1, *t2, *t3},
		Data:             *data,
		BlockNumber:      5730114,
		BlockHash:        *bh,
		TransactionHash:  *th,
		TransactionIndex: 50,
	}

	if err := expectedLogs.Compare(result[0]); err != nil {
		t.Error(err)
	}
}

func TestEth_GetStorageAt(t *testing.T) {
	result, err := NewRPCClient(config().address).Eth.GetStorageAt("0xd780ae2bf04cd96e577d3d014762f831d97129d0", 1, rpctypes.QuantityLatest())

	if err != nil {
		t.Error(err)
	}

	expected, _ := rpctypes.NewHexString("0x0000000000000000000000000000000000000000000000000000000000000000")

	if !result.IsEqual(expected) {
		t.Errorf("getStorageAt result not correct,[Expected: %v, Actual: %v]", expected, result.String())
	}
}

// GetTransactionByBlockNumberAndIndex
func TestEth_GetTransactionByBlockNumberAndIndex(t *testing.T) {
	expected := getTransaction_99192()
	result, err := NewRPCClient(config().address).Eth.GetTransactionByBlockNumberAndIndex(expected.BlockNumber, expected.TransactionIndex)

	if err != nil {
		t.Error(err)
		return
	}

	if err := expected.Compare(result); err != nil {
		t.Error(err)
	}
}

func TestEth_GetTransactionByHash(t *testing.T) {
	expected := getTransaction_99192()
	result, err := NewRPCClient(config().address).Eth.GetTransactionByHash(expected.Hash.String())

	if err != nil {
		t.Error(err)
		return
	}

	if err := expected.Compare(result); err != nil {
		t.Error(err)
	}
}

func TestEth_GetTransactionCount(t *testing.T) {
	result, err := NewRPCClient(config().address).Eth.GetTransactionCount("0xfdc795aa0c3b4b30bca8275d61f8dfbd49d9e912", rpctypes.QuantityBlock(5212393))

	if err != nil {
		t.Error(err)
		return
	}

	var expectedCount int64 = 9

	if result != expectedCount {
		t.Errorf("wrong count [Expected: %v, Actual: %v]", expectedCount, result)
	}
}

func TestEth_GetTransactionReceipt(t *testing.T) {
	trans := getTransactionReceipt_99192()

	result, err := NewRPCClient(config().address).Eth.GetTransactionReceipt(trans.TransactionHash.String())

	if err != nil {
		t.Error(err)
		return
	}

	err = result.Compare(trans)

	if err != nil {
		t.Error(err)
		return
	}
}

func TestEth_Hashrate(t *testing.T) {
	hashrate, err := NewRPCClient(config().address).Eth.Hashrate()

	if err != nil {
		t.Error(err)
		return
	}

	expectedHashrate := config().hashrate

	if hashrate != expectedHashrate {
		t.Errorf("wrong hashrate [Expected: %v, Actual: %v]", expectedHashrate, hashrate)
	}
}

func TestEth_Mining(t *testing.T) {
	mining, err := NewRPCClient(config().address).Eth.Mining()

	if err != nil {
		t.Error(err)
		return
	}

	expectedMining := config().mining

	if mining != expectedMining {
		t.Errorf("wrong mining [Expected: %v, Actual: %v]", expectedMining, mining)
	}
}

func TestEth_NewFilter(t *testing.T) {
	p := CreateNewFilterParamsWithOneTopic("0xd780ae2bf04cd96e577d3d014762f831d97129d0", rpctypes.QuantityLatest(), rpctypes.QuantityLatest(), "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	result, err := NewRPCClient(config().address).Eth.NewFilter(p)

	if err != nil {
		t.Error(err)
		return
	}

	t.Error(result)
}

func TestEth_ProtocolVersion(t *testing.T) {
	result, err := NewRPCClient(config().address).Eth.ProtocolVersion()

	if err != nil {
		t.Error(err)
		return
	}

	if result != config().protocolVersion {
		t.Errorf("protocol version wrong, [Expected: %v, Actual: %v]", config().protocolVersion, result)
		return
	}
}

func TestEth_Syncing(t *testing.T) {
	result, err := NewRPCClient(config().address).Eth.Syncing()

	if err != nil {
		t.Error(err)
		return
	}
	if !result.IsSyncing {
		return
	}

	if result.CurrentBlock < 10000 {
		t.Errorf("current block cant be smaller than 10000")
		return
	}
	if result.StartingBlock < 10000 {
		t.Errorf("starting block cant be smaller than 10000")
		return
	}
	if result.HighestBlock < 5000000 {
		t.Errorf("highest block cant be smaller than 5m")
		return
	}
}

func TestEth_UninstallFilter(t *testing.T) {
	result, err := NewRPCClient(config().address).Eth.UninstallFilter("0x12")

	if err != nil {
		t.Error(err)
		return
	}

	if !result {
		t.Error("should uninstall filter")
	}
}

func getBlock_4ad331() *rpctypes.EtherBlock {
	block := "0x4ad331e4f8076ca863fcfbc407d17a0f4c8086ca3288a0f7e00965f3b634d04b"
	hash, _ := rpctypes.NewHexString(block)
	parentHash, _ := rpctypes.NewHexString("0x7af09e57a36cc4362f944f328792c062e342f9ebf1e0a5aff9dfb1617f011c30")
	nonce, _ := rpctypes.NewHexString("0xe8931c7c0b7068cc")
	sha3Uncles, _ := rpctypes.NewHexString("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347")
	logsBloom, _ := rpctypes.NewHexString("0x8e00c8791560084986228500aa20011266fc840b0a10143741032d40692c8d08a0c25cb008480c00440066c0d0d14a8240c0470cb413207055d08b2500244110030022c1802006d16950580c1c202908523801213228c020e40dab21c600402502e0042a13c4e7520110430e68901902b0a1211b891cc9c4508c2038b002398160730810f0650802ac400686104c943a808135097944e9946580aed020574ce04a3980128cc371400a34184a05358901b400049303d181090611d101a358e2100a5b09138b05001002502c4440786c02385b28151e0a4ac024588280e601b1926c702126a2815c01448080c21638be5674608c91a140196ea300109c41002380")
	transactionsRoot, _ := rpctypes.NewHexString("0x2dfd960db39f5e8f0029e9d074ecb4585b21f44de7ee8cba625831870f0c90a7")
	stateRoot, _ := rpctypes.NewHexString("0x959dd2b6888cab28bf25b5324270d0f34f998821443309e8f5e7f31c90b45a7c")
	receiptsRoot, _ := rpctypes.NewHexString("0x805858cd10a988c1300811dcb48fbb78d65966f75eacea524496037c06e09022")

	ea, _ := new(rpctypes.EtherAddress).FromString("0xb2930b35844a230f00e51431acae96fe543a0347")
	difficulty, _ := big.NewInt(0).SetString("3090008557990385", 10)
	totalDifficulty, _ := big.NewInt(0).SetString("2982634332710548077770", 10)

	extraData, _ := rpctypes.NewHexString("0x7337")
	gasUsed, _ := big.NewInt(0).SetString("7979953", 10)
	gasLimit, _ := big.NewInt(0).SetString("8000029", 10)

	return &rpctypes.EtherBlock{
		Number:           5235555,
		Hash:             *hash,
		ParentHash:       *parentHash,
		Nonce:            *nonce, // deviates from Etherscan  https://etherscan.io/block/5235555
		Sha3Uncles:       *sha3Uncles,
		LogsBloom:        *logsBloom,
		TransactionsRoot: *transactionsRoot,
		StateRoot:        *stateRoot,
		ReceiptsRoot:     *receiptsRoot,
		Miner:            *ea,
		Difficulty:       *difficulty,
		TotalDifficulty:  *totalDifficulty,
		ExtraData:        *extraData,
		Size:             31169,
		GasLimit:         *gasLimit,
		GasUsed:          *gasUsed,
		Timestamp:        1520763355,
		Uncles:           *new([]rpctypes.HexString),
		Transactions:     transactions(),
	}
}

func transactions() []rpctypes.HexString {
	trans, _ :=
		rpctypes.HexStringListFromString([]string{"0x5ad2329221f1df749119e5103a9be79cb250c6bffd57c983e80cff80ebf74e24", "0xef02ed9fac01cbe4ef78f2070e2d55cffde3e7ece8e149a58fb49b5e71c971a3", "0x6cb044f4d9afd81d818f1328474aa509d513f11f3a7320d5faa1b0164a60726b", "0x605f1e8885e544b0b4f913053e67b51228b12fbc4fb25f66f0105531048aa01a", "0x81288daffcafa22a8444da78f3de5f799b8f7857efd72c53868d9a5b176caf30", "0xb44f7c82824caf6b29fdb940526fa02e4d91535c8f9e8092b9f04b0779f13a9e", "0xa2b321cab23942a68311fbe8279f4f5e2c8da5abf76e8c6acf9c0e48868cc430", "0x5f6233560cb5cf0793c8e57c7a541279be24894b9092aa4e49e8b972c7fe5bd3", "0x2d726e89d2d64527b4614a41520db61c542fc78ade15420570ee9450dcdd0e0d", "0xe5693364a4765c5a0b691773065fba48fa52521688d9d3fc1b2bd1ecc568c5a0", "0x76ff9dc57122a8337464a4b8899fecc0d1625d2d68a5ca8cf7866f9974914383", "0x6c11a014d56227aecf0c670f9d12ce9e3ec3f03c2b8eb678e1742e38aebcc474", "0xa0bd5e22a3b531eb3ab34b389f6524b59283d7759e75dba662dbe1b8bdc9caac", "0x7672b09c04012832867ff50bd4e81efa9cb7538e0453be9ac8829ae07247fdd1", "0x8fc5a4db5bac2d80ee3800bda25860e9616e3f2e61fbb1a4d93304cddb90da70", "0xab59b3e30be0702e8e958e7e9fecfb87f30e6f6d992bfcde16661aaead0b83c8", "0x6e072c52462be37a156f4e91f90ee96f3976cb63426304f9cbad3666038e7bb1", "0xc781614c207448000c349190348540d720b6ac94404dd1006206f8d6c55f127d", "0xc1e21db9b61ca90156bd73fe25816430e6b08a0c317b18a72e709a358a803b15", "0xc819e656ba0fd358212d4df98471ca9238cef7b79f1206c7bc1d2a411768dc86", "0x94f3edb1e94461f141545f0211dd9b2271be49fe58e11ff46d83153c5951abb4", "0x0361fb3bbd3e0f6e17e3f64c7d55c29f7579f76f7cab217deef94e6b6785d099", "0x208a8cb9d1891617d04995d2f640d1eab29b2de13bcde726fcfe010a538d9d21", "0x14640b1f03f63f06579ffe4a91122f226b87931d1704817388bce0abba8a24c6", "0x884a6108e710c26c978baed538c93eb291adbb7ffcd1375ae3871a073813ed68", "0x48a75683208c6ccac55b66be22c3050439664f0629bf3e8be19e90a6d3046141", "0x7abb7c9645cb143a75c1df59febb6769eb0a90fe6c922f9914e043040f41e01d", "0x479e3a73dbbb70a86d10889fdaddab9fe2d8ccc1261b731c91d2e26869a6f1fe", "0xa6963867c2e0a88513428da36dce46e8a90c879b5a5059bd0974d00c84415c1b", "0x83d92bbddad4cc7d0a0443643b386b9688568317275b96ea3fd01d8e602846e7", "0xd7efc97e54aa72587036ad66187b210968833d1ec8d69d05422651ce1d9ed08e", "0xa19f98570d43000bb8b9f6dd606572c6279af0032076d4de4e28e7e6833ae45d", "0x0a68e31f3595030254cd18c8111136853ed9fb07555e585c67222fd167f65be6", "0x1cb6f1be98e4f21d0ad2ed137230f99dce1d15adc1e440b5e0717887328cbf03", "0xe7687ba229b2a73ea8495186b4aca0fdd4418469bbd1b42f06469dffcd2acc82", "0x8e5d9b2e40690f57fcfafbdaeab54ba560dac77e56db1c75c41153027b54fcef", "0xf9d330fa2f7d421a33da180462ae5141aede25590974bb51e39cfd5cdfdcc1f5", "0x2b3bf673454c8a1098967b3961701e3c7f165775c96fcb13d6dce17f1c577187", "0xcaf37dc8aefc2eba99a355f4bb6f0635ccd4a768c7f6496bc9011d8c3392cb6e", "0xe0e7e4b0e971b4697f06e32a07d549db49eb9e17c0c997123d4efe8c99d8e7f0", "0x90eeab2a374771d441aa0bef8c73303dbddb3b21455d0cd0a56d92c85747d879", "0x633e7e8efbb735d41ba2ae5687b04588935847419fe893332f8ac650964feb55", "0x7d40dba467d06f7cd4336dc40d4cedc9dc915e6a3475e4451f50d30c8eefeb0c", "0x7d9df97a21a87b5a8d0b07adc4701708752a25aee1984b87c94d6552204dc681", "0x6f853f5cedfabebdf69ba63e6bfe990532ccc862def714a90f5eaa9e1b3f9255", "0xeeee7c800aa4426df0a89332f9832cd39b11c8068f23b8ff753916fc4d53fb8e", "0x116f764eac36dce859eddf1e9de4296255bed84e631bcd9e1bc4bef6e1b3cf07", "0x607c3aae51993c9bb1951bd969e618056092c47806fe628ebced699de0c11561", "0xf5068bc1a5efc33a289da712ff42c258d5a78d06082bea1a54bdfb4cced17d58", "0x3bb2ef42b71d098e0c57c667cb85704479486b8146772d6a1bbae52ad9b6401a", "0xfb1c526eb9b2d179832ab32c334983547945b630343b9368cbe436c7991aac74", "0x7af1d8848d596422de412698fab429bfca1191f00577cd109ff9f72a42c6f59d", "0x3dd11e98e4c818e0580db1c1b556591f310a3311b8464b0391ee9e982f68f31e", "0x0877a98f729d79c229da6e4634305eb8e39f4f582daa661dda52771ada0ce77f", "0xb7abaaf1cd08a051e20e5fbb56edf630c2ca8f065823b795a461c5228d3c80f9", "0x1ae87f8ebd352723ffaa37a6a734586d15af58d1662fd8c82cd7c4b02a629513", "0x6a2f62bf9b9cf131e8fcb03d7fbfb917e3f0ddf54418618fc1b394547cf72fff", "0xae37d691b225d73471708a31e95f2d66441d46179df2e934a287d1914b70beea", "0xe5b80a80a77450d9e6e37afe017d171c6934a6947976abce638083b9928e0881", "0xce40aec79390fbc280869061e29fc15864bc0473a89f853645c6424072a3cf48", "0xf045714e3f62459b56331a930e79b527beb3014ef3ba20c9361bec1c1bf6d0d7", "0xef5230688aff4bc7215ef0e2ff9f1df361c3a4530cb00559279ad7aad9b41ed3", "0x455b43728221bfa1b93b5537c41d0fb318ddd6f34913561db1ec2f3fab059641", "0x136d2de2921c36afb34d6dfa3b0a1e2ff1d659190fca8eb4d745b823a650bfa4", "0xfcf53c64e46ab747b81e289bdfc5b47a2d562c1127feac0f97313205a611c37a", "0x583a53ad5495d7cbf0048a18e0dacfd154b8d377df4e2238fbf883c09b6102b1", "0xc557ed57a54207f04a103dfd6d8bd1daca23d9673dd53718c5967e59904a61b0", "0xfc36b4348a53e3351821bc610a40160d999436f0077275aa97d54343c9587c83", "0x2fe0a6a5f350896c62ee02b1aaf1e471c430653595f6d2a6409ab779f67528ec", "0x0a9fe3f62c3ea20a8d7e80f029085f30bf1668b33373cfa37cbe5bc03d208cc2", "0xfae136bd77dcefb6c067378fff1cb691076b76594b197bba902b4fbfb9023abd", "0x2e747f0de84144b2b782b59935a508b437bbd17d75723375e176febb7c2eae91", "0x5031587ca39a98840066a093e6033596c74f00e2212fffef4c15ad25d5cd5956", "0xd63010fccbe60ad6bce05a535a9ebda6c96204f6e478090c7c60d8223dd6c288", "0x97d7ee282ad961241eed30dd4ec1cc04307570a9728da77b8669c1cf3291c237", "0x52e7691e9a237db43d71601502ce6c0ad40fcc50d2f1d6dea94da0438a47938d", "0x3474ad53dceda6e359ce43fe9a23d1a95b106db3778319f702978721bc2ed14b", "0xd60a5b98111bb5c1f3119397b3c6b624f0ac60caea10523a66c112cb6cd5aaeb", "0x3db2a63b0c4490e0543aa21fd4f0c7e125201536bd401fa4c32956966373fb80", "0x40e96b523e52228a87cb756c111ee1822d049ded743de0d3a8ebcf284f6c6c44", "0x4ab8adc74d20aeb8651a2f17226049b5130bc0efd0401a2a714d7c7da3753b02", "0x41b6307f95d798e32c367b820b537e530108c06b099d716cbc494f141d1a7688", "0x2e90535a246166581c9b0520054f4f11813be0000c7448fcc5e93bfbd0248ce1", "0x5974f8c31620d1218b770dbe090786bccc7b48d5dfe5a33b7276a24aee856a0b", "0x8735b89fea7603fc76239101d7649447087fda8aa290d94329999b56e6de6bd3", "0x99192da8574412d70d210b73d859904a0f3d1493051d751e4f32fb2905d1844a", "0xec3e1c2cdbbf2ffa8b4585fc6c418419bb5d36495fcb0aa206d89eda6b1e8033", "0xed646305054dc09dc7f192ccfc9a721fde52d34c9cf65867ce34560809c9bb22", "0x831dac73d565e455663432b3dac12ad55080e82085cbd0bc21eadba275b8dacd", "0x1555e172f273e278ec1e8c35a1c2bb2dce010cc5a5f158604670f008fc7264ed", "0x9dea3caa2ff11c451d23074ae4dddcfb21c1f078e0277f3efe5f63387a6a74cd", "0xec54180f11dff22522657a5f0f7b08e30f80c04b97d78cf2135c93c7fdf4744a", "0xc769ed131a0fffc458939e85c50c4108f70f46ddd85c84a622860e2d6c0e1d7e", "0xde2cf0f57b21491ce6eeb4cf060b0856fb4392a037809c1557db72c1b848612a", "0x1734d879369eec9c23b915755f38819942da299d77a38a2f8272748a658e50ff", "0xb1c33589428f8893303feaf054bd755473ef7c77b358972598d4b89b0564a98c", "0x5c0f77dadb3aea71def08bf6d02adbdd9e86b5492d0995bd6d47655fd5cd951d", "0xa80f753f2a57b924a216670a74b88c4c0059392fd54bfb13d408604247ebb9c8", "0x85b79e58c0d6487478aaa5d2c8dbd5a155dab41368396f065272c5944eed2f2f", "0xf8d046fe6bea83d84492f3042f6f7649c0e8a69fb252e6720ffa4b721fa16d78", "0xa03b7661377878837ba9322dd418d190c320b350580ec0ff3308e8dd86b3e790", "0x1a43eb51aba6cfc86bc9603551839ea929c7a99fbcc3d9f12c20f14883993c3b", "0xad92a829a3da7da71370ab4995d4e4bb91aba47d580b5e54ceb386cf270dab78", "0xb87deac73ac5262bf096b0fd398cc84a2d6b49c2fb4ed63b789520237d859ed5", "0xac72c4635e95d627801e23faa8b63058318cb3a4e3b61d7516412d142099deb4", "0x6abdacc01ee51317b0e9237d857ef65c94ea70da5357f80300e84f5b9d428593", "0x6f9be88801fd49b3cb28c9806698303af125d7b5a6e57d2e05931029acc0978d", "0x7a71fd312d55cca1d2ae0e3c5c04fa0e6862e744fdeb087d7a588f0d159a226c", "0xecbf0577d8c5fa85004cc99e70781644b3a63e1ee8d5e601b2609059f414d3c0", "0xce5601cde90db2b627cda9cab1fc083cb60543d6d651a3bc89310f69894da6a7", "0xcaa459189a90e2d76174103cfdb1a5772e15887b03c4b3550a5e835c5e280436", "0x030099deb34052e9308f3e88b10a2ba44a86c06f840d777000c217c06b8f0c5d", "0x2c67afdd8879c31fc736d810fb9ad81a7e2aec03235a0bd51bcb6bf95f8e6b25", "0x8405709fb8389c92ff0f334558bd4bbd19b00cb97dc9fe01b4aabd4efc669a1d", "0x32c39a3a6ca8de72b8b57ae46ef0878544a7698bda0467edcff09a517eb43bd9", "0x9e2a1d1a5f8dc16559d097801af4d74864e31717e1fe4e6b7e3243fc05c83001", "0xcc136d6faa5aad262ddd96bfbc54a8981039f7990ab97027f4a049ee2347f5a9", "0xf7519e65bf6b04b6b8c0ec6f0a5b74e54c4c10021be3d7f68ed8b3ce92603e85", "0x67b82f043764ce2b8604f40e1db78b898e1fef6bc445421827ce25d41fe53d8a", "0x588faaaaf2eb5f4eb0bf54cf99830b35a9ff559f16c2fef14a2c276934fcacaa", "0xe7b124fd25bb2c068731fa234c148c656f73b6077b07d2c11cb432f81d839b18", "0x8f371aaf2ccd43108aed73a309d143d350941353582f23b3ca942c1ce4d4d569", "0x3e4482ee2ae2733aff29b7f92910814fd45c63b4ec938fe4b5714b3436d2d136", "0x062ffcde2cc206a20eb91a3c41df29200e8b6371b6b0b823bd2b4f0f86822959", "0x7b6ec721b1b4412feeb29e096323a68ff3d427841e34555bc0ff9d60898952d5", "0x79fe9eb3edda0098372c7b57d9279c7c3fb4ebc94b579471ec0217c3c07aa334", "0xedcc789d57e48b1c32eb3ef07210004b1c8ae05db8e0659e479bd98b12099900", "0x58129d5c7181635430f241de5eaae748c21befd108052e9363d1cc99447698e0", "0x47dbd252cf83cfb34514c5fd5692f9f8bfc6201e8e83ee1c5dca312fec15214e", "0x929557653e869872f60ecc3bdd01f2bfa57c88da53219db6df7c0b592d540743", "0xbf9cf830a61b96926d776f122f0fff016c04f07a67a6578789212efca238bc90", "0x5fa55d920d2ca0404c6165318697eeab0e16bb88c2bbe7d2f97ff6c701bc548c", "0x1b765e227bbf8ffa8659fc9bbdb84eaa7e71a09e1c618fd7f216c240b45d6310", "0x0c32b90ae865d097107abf10e4eef156d3f6c89116d6c05fb5c37751face6643", "0xe74d1d1a0cf9e2255cc21fc3d6e6fd4de665e509fb248ba9468fba2fdea59121", "0xf963707caaf2d39893773fba4ad66f3b3a8d7a2c5697a22c9f27cc07e5a10a00", "0xe995cb68abe2bd492f8a3b43022f725e78cb6231b0d2ffdb66de06c707f053cd", "0x8c81541774ed61c584a26512c5b0a6555196951f923639716fc13aa286167303", "0x582334d3c26c7dd522de567fa0fae6fd1647289f5a190553ba72dd5591052ace", "0xcbd39ea26a41e2b51bf7c3e61784226228d6ea91517e565a7d626118d6174cba", "0x2ad039362885281fa60441b61475cf2420d3a6364f9fe8a95698e8941c8309f1", "0xef481b7fe3b2e57209c0eb1ca37c4cb0a1ee17a816366fd2a223afe2062fac81", "0xcc270c38b634904969a3cbed82903306e55a55d659be6463e76195fcdf1afa61", "0x2d13f985d2e38be293a90387a78f897ada33b5274d685508b24ad1b435dbb654", "0x6498075c52b9316862c27dec2cc744d54dcbf580ec5b6e092a69675e9afb8527", "0x55bd57465e076fdc9be8333157d42124e98a5d30cf8442e3131f3e1f7167806c", "0xa881649a27fb54d491ed719996831ec803dc3d73a5af43e80c4b192e5ed5e360", "0x138ed5ca81887eb81df24f24ee24806a0f223fa4a0efbe0ff0d8c6ac742c35b4", "0x00a7578bce9bdc890c2c2f8b9352d3c692f34dd0df0f2d092aff4c2f3e2e7f22", "0xac44dc89f6633b2497dacf3780c21e6e9a27e77675d736dd41128d20aecbc4ea", "0x457c6845ae37b029d4b5613e59f1b5ef1d7ab7a402e0d375f370c326e07372f8", "0x4af972b3ef98284ea8d011ad5d722e63cf4901f9ca622c24a7afff0169c8d2f7", "0xc0ca64edbb05f0af6d856f1f2a52bb5bd5f144333fafbe317a9a29f85edf24e5", "0x49483178c77701718b75d066b10a1b516b4bdcc08c06b04a4148d7e875269d5f", "0xb339a0bb8fca175bc859ddf3e9eae46b8bc77cefbe2cec9ee6c0a23270101596", "0xc43f8f1f98ae19f7ceba8fdd22f3fdf1bf71cd578b9f08ef4cd122865292e3f0", "0x660ef529265185031dc057a3a39232e3555e0e828dbe1dab0facbe6b9f435aa3", "0x0ea6d01397b08ede3ea3a3e5371827e5c4c9a8e9a9b66bbcbfdac285121b69b7", "0xbda08fc8a90a00309977c85ae01681c57532cea672ef42568d75bc828dba4d9f", "0x09d936744bb0e838fa3605b416f9b69b0a45ffd619a36e976dbfd193c2598c59", "0x7fb6901da74cc9d1836aff27abcec4489d494efd23c29aa28b648cb2784ad6f2", "0xb2221a4291247141d1cc2f1d0c0a4fd0df4cb900873459b5d3994548b08dd0b0", "0x9143b733bb8236f033beaa0aef50c5fd327d4ffbe2ea2895661af3cdf5ac905d", "0xaa91f550900e518ba35fa5c041ab1f053810844e8aee00dc5c6bf4d6301c228d", "0xdc939e30b98e70ee89b20493eaaeef79a913850ffaf68129efb52b4829b8779b", "0xd59e6c8e131dc23566ff357ef1be5a5597449e3cf32a8c5fb7b64ecab9cb54d3", "0xcad691b91b3139aa39e5fbdef267441d2071b889de078e20c3704d9e13766bbb", "0xcb773063d77738f50459d9fbe807b686ad37e318e108c5fbdf8b5215cfe11e38", "0x1abce240792e6ae6fcd970eb7c763762b66fae9049767fb686233b8450dcf6ed", "0x321f9e68655477e575b50f8a466887e06f86be8f0a6cf58438e5865e08090a76", "0x4d488a66d637157428b50279aa00a128ed921b53bdd56ec508c7b4863a1a0b79", "0x404fa0fd8c442f0d53444ebbec341677c21ed1931337bfbf131f00343fa92215", "0x7e7417f4e7a992dc8db0276bd83597077ac9d7c8b4fdc02fb27093c7bd372cc3", "0x8ba8bf871e8f7811c5667fe41e35ab0364c5a6f9afb1f2116fbf32a2f7fa7f97", "0xf5fab7832c3a43e4a3f9e3da9525838665557cf97f3044d71659a635e4053086", "0x5bb98e42415f18021c4edcd394c85f026721ec30c3086ad6190aadc2fe72ca51", "0xf5290c26d037cf827c5a277345af391c59e0501a0c5a66080abadc9d57e6ddd8", "0x7a9efdafcca92b72357caf1dc51b56a9980afbe08b2f9669aa1e9edee373c1b3", "0x0b86fd0428ef0646ccdf5325b8db7d92c6d046bdcd11de91ce17ab4d2f2c3ef0", "0xc3ad61044120a22ddc17e4c69d075c04167233782a431ea559359fd3884b75d3"})
	return trans
}
func getTransaction_99192() *rpctypes.EtherTransaction {
	temp, _ := big.NewInt(0).SetString("250000", 10) // gas limit
	gas := rpctypes.NewEtherValue().FromBigInt(temp)
	temp, _ = big.NewInt(0).SetString("8000000000", 10)
	gasPrice := rpctypes.NewEtherValue().FromBigInt(temp)
	from, _ := new(rpctypes.EtherAddress).FromString("0xfdc795aa0c3b4b30bca8275d61f8dfbd49d9e912")
	to, _ := new(rpctypes.EtherAddress).FromString("0x8d12a197cb00d4747a1fe03395095ce2a5cc6819")
	nonce, _ := rpctypes.NewHexString("0x0e")
	input, _ := rpctypes.NewHexString("0xd0e30db0")
	hash, _ := rpctypes.NewHexString("0x99192da8574412d70d210b73d859904a0f3d1493051d751e4f32fb2905d1844a")
	blockHash, _ := rpctypes.NewHexString("0x4ad331e4f8076ca863fcfbc407d17a0f4c8086ca3288a0f7e00965f3b634d04b")
	v, _ := rpctypes.NewHexString("0x1b")
	r, _ := rpctypes.NewHexString("0x948153c3a2f21176d6606a39f1442f3711352c745bf47de69a663d60c4334bc9")
	s, _ := rpctypes.NewHexString("0x88a88fbeda64ae8b5195efe296712ee2bd0a324f4ac45b50c863f41a2957531")

	return &rpctypes.EtherTransaction{
		Hash:             *hash,
		BlockHash:        *blockHash,
		BlockNumber:      5235555,
		Gas:              *gas,
		GasPrice:         *gasPrice,
		From:             *from,
		To:               *to,
		Nonce:            *nonce,
		Input:            *input,
		TransactionIndex: 85,
		Value:            *new(rpctypes.EtherValue).FromBigInt(big.NewInt(100000000000000000)),
		V:                *v,
		R:                *r,
		S:                *s,
	}
	/*
		"value": "0x16345785d8a0000",
},
	*/
}

func HashToEtherTransactionFull(l []rpctypes.HexString) []rpctypes.EtherTransaction {
	newList := make([]rpctypes.EtherTransaction, 0)
	for _, v := range l {
		newList = append(newList, rpctypes.EtherTransaction{Hash: v})
	}
	return newList
}

func getTransactionReceipt_99192() *rpctypes.EtherTransactionReceipt {
	logs := rpctypes.EtherLogRaw{
		Address: "0x8d12a197cb00d4747a1fe03395095ce2a5cc6819",
		Topics: []string{
			"0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7",
		},
		Data:             "0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000fdc795aa0c3b4b30bca8275d61f8dfbd49d9e912000000000000000000000000000000000000000000000000016345785d8a00000000000000000000000000000000000000000000000000000189a1b765ca13ca",
		BlockNumber:      "0x4fe363",
		TransactionHash:  "0x99192da8574412d70d210b73d859904a0f3d1493051d751e4f32fb2905d1844a",
		TransactionIndex: "0x55",
		BlockHash:        "0x4ad331e4f8076ca863fcfbc407d17a0f4c8086ca3288a0f7e00965f3b634d04b",
		LogIndex:         "0x40",
		Removed:          false,
	}

	raw := rpctypes.TransactionReceiptRaw{
		TransactionHash:   "0x99192da8574412d70d210b73d859904a0f3d1493051d751e4f32fb2905d1844a",
		TransactionIndex:  "0x55",
		BlockNumber:       "0x4fe363",
		BlockHash:         "0x4ad331e4f8076ca863fcfbc407d17a0f4c8086ca3288a0f7e00965f3b634d04b",
		From:              "0xfdc795aa0c3b4b30bca8275d61f8dfbd49d9e912",
		To:                "0x8d12a197cb00d4747a1fe03395095ce2a5cc6819",
		CumulativeGasUsed: "0x39a599",
		GasUsed:           "0x7335",
		ContractAddress:   "",
		Status:            "0x1",
		LogsBloom:         "0x00000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000040008000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		Logs:              []rpctypes.EtherLogRaw{logs},
	}

	log, _ := raw.ToEtherTransactionReceipt()
	return log
}
