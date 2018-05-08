package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/Leondroids/go-ethereum-rpc/rpctypes"
)

func (client *Client) RequestEtherBlock(method string, params ...interface{}) (*rpctypes.EtherBlock, error) {
	p := []interface{}(params)

	if len(p) != 2 {
		return nil, fmt.Errorf("wrong params in reqeusting ether block: %v", p)
	}

	full := p[1].(bool)

	response, err := checkRPCError(client.Call(method, params...))

	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("response returned without error but no block found for %v. is block not synced yet?", params)
	}

	return getBlockFromResponse(response.Result, full)
}

func getBlockFromResponse(response interface{}, full bool) (*rpctypes.EtherBlock, error) {
	js, err := json.Marshal(response)

	if err != nil {
		return nil, err
	}

	if !full {
		r := new(RPCEtherBlockRawWithTransactionHash)

		err = json.Unmarshal(js, r)

		if err != nil {
			return nil, err
		}

		return r.ToEtherBlock()
	}

	r := new(RPCEtherBlockWithFullTransactions)

	err = json.Unmarshal(js, r)

	if err != nil {
		return nil, err
	}

	return r.ToEtherBlock()
}

////////////////////////////////
//
// AUX Types
//
///////////////////////////////

func (rpcBlock RPCEtherBlockRawWithTransactionHash) ToEtherBlock() (*rpctypes.EtherBlock, error) {
	number, err := rpctypes.NewHexString(rpcBlock.Number)

	if err != nil {
		return nil, err
	}

	hash, err := rpctypes.NewHexString(rpcBlock.Hash)

	if err != nil {
		return nil, err
	}

	parentHash, err := rpctypes.NewHexString(rpcBlock.ParentHash)

	if err != nil {
		return nil, err
	}

	nonce, err := rpctypes.NewHexString(rpcBlock.Nonce)

	if err != nil {
		return nil, err
	}

	sha3Uncles, err := rpctypes.NewHexString(rpcBlock.Sha3Uncles)

	if err != nil {
		return nil, err
	}

	logsBloom, err := rpctypes.NewHexString(rpcBlock.LogsBloom)

	if err != nil {
		return nil, err
	}

	transactionsRoot, err := rpctypes.NewHexString(rpcBlock.TransactionsRoot)

	if err != nil {
		return nil, err
	}

	stateRoot, err := rpctypes.NewHexString(rpcBlock.StateRoot)

	if err != nil {
		return nil, err
	}

	receiptsRoot, err := rpctypes.NewHexString(rpcBlock.ReceiptsRoot)

	if err != nil {
		return nil, err
	}

	miner, err := new(rpctypes.EtherAddress).FromString(rpcBlock.Miner)

	if err != nil {
		return nil, err
	}

	difficulty, err := rpctypes.NewHexString(rpcBlock.Difficulty)

	if err != nil {
		return nil, err
	}

	totalDifficulty, err := rpctypes.NewHexString(rpcBlock.TotalDifficulty)

	if err != nil {
		return nil, err
	}

	extraData, err := rpctypes.NewHexString(rpcBlock.ExtraData)

	if err != nil {
		return nil, err
	}

	size, err := rpctypes.NewHexString(rpcBlock.Size)

	if err != nil {
		return nil, err
	}

	gasLimit, err := rpctypes.NewHexString(rpcBlock.GasLimit)

	if err != nil {
		return nil, err
	}

	gasUsed, err := rpctypes.NewHexString(rpcBlock.GasUsed)

	if err != nil {
		return nil, err
	}

	timestamp, err := rpctypes.NewHexString(rpcBlock.Timestamp)

	if err != nil {
		return nil, err
	}

	transactions, err := rpctypes.ToHexStringList(rpcBlock.Transactions)

	if err != nil {
		return nil, err
	}

	uncles, err := rpctypes.ToHexStringList(rpcBlock.Uncles)

	if err != nil {
		return nil, err
	}

	return &rpctypes.EtherBlock{
		Number:           number.Int64(),
		Hash:             *hash,
		ParentHash:       *parentHash,
		Nonce:            *nonce,
		Sha3Uncles:       *sha3Uncles,
		LogsBloom:        *logsBloom,
		TransactionsRoot: *transactionsRoot,
		StateRoot:        *stateRoot,
		ReceiptsRoot:     *receiptsRoot,
		Miner:            *miner,
		Difficulty:       *difficulty.BigInt(),
		TotalDifficulty:  *totalDifficulty.BigInt(),
		ExtraData:        *extraData,
		Size:             size.Int64(),
		GasLimit:         *gasLimit.BigInt(),
		GasUsed:          *gasUsed.BigInt(),
		Timestamp:        timestamp.Int64(),
		Transactions:     transactions,
		Uncles:           uncles,
	}, nil
}

func (rpcBlock RPCEtherBlockWithFullTransactions) ToEtherBlock() (*rpctypes.EtherBlock, error) {
	number, err := rpctypes.NewHexString(rpcBlock.Number)

	if err != nil {
		return nil, err
	}

	hash, err := rpctypes.NewHexString(rpcBlock.Hash)

	if err != nil {
		return nil, err
	}

	parentHash, err := rpctypes.NewHexString(rpcBlock.ParentHash)

	if err != nil {
		return nil, err
	}

	nonce, err := rpctypes.NewHexString(rpcBlock.Nonce)

	if err != nil {
		return nil, err
	}

	sha3Uncles, err := rpctypes.NewHexString(rpcBlock.Sha3Uncles)

	if err != nil {
		return nil, err
	}

	logsBloom, err := rpctypes.NewHexString(rpcBlock.LogsBloom)

	if err != nil {
		return nil, err
	}

	transactionsRoot, err := rpctypes.NewHexString(rpcBlock.TransactionsRoot)

	if err != nil {
		return nil, err
	}

	stateRoot, err := rpctypes.NewHexString(rpcBlock.StateRoot)

	if err != nil {
		return nil, err
	}

	receiptsRoot, err := rpctypes.NewHexString(rpcBlock.ReceiptsRoot)

	if err != nil {
		return nil, err
	}

	miner, err := new(rpctypes.EtherAddress).FromString(rpcBlock.Miner)

	if err != nil {
		return nil, err
	}

	difficulty, err := rpctypes.NewHexString(rpcBlock.Difficulty)

	if err != nil {
		return nil, err
	}

	totalDifficulty, err := rpctypes.NewHexString(rpcBlock.TotalDifficulty)

	if err != nil {
		return nil, err
	}

	extraData, err := rpctypes.NewHexString(rpcBlock.ExtraData)

	if err != nil {
		return nil, err
	}

	size, err := rpctypes.NewHexString(rpcBlock.Size)

	if err != nil {
		return nil, err
	}

	gasLimit, err := rpctypes.NewHexString(rpcBlock.GasLimit)

	if err != nil {
		return nil, err
	}

	gasUsed, err := rpctypes.NewHexString(rpcBlock.GasUsed)

	if err != nil {
		return nil, err
	}

	timestamp, err := rpctypes.NewHexString(rpcBlock.Timestamp)

	if err != nil {
		return nil, err
	}

	etl := make([]rpctypes.EtherTransaction, len(rpcBlock.Transactions))
	for k, v := range rpcBlock.Transactions {
		et, err := v.toEtherTransaction()

		if err != nil {
			return nil, err
		}
		etl[k] = *et
	}

	uncles, err := rpctypes.ToHexStringList(rpcBlock.Uncles)

	if err != nil {
		return nil, err
	}

	return &rpctypes.EtherBlock{
		Number:           number.Int64(),
		Hash:             *hash,
		ParentHash:       *parentHash,
		Nonce:            *nonce,
		Sha3Uncles:       *sha3Uncles,
		LogsBloom:        *logsBloom,
		TransactionsRoot: *transactionsRoot,
		StateRoot:        *stateRoot,
		ReceiptsRoot:     *receiptsRoot,
		Miner:            *miner,
		Difficulty:       *difficulty.BigInt(),
		TotalDifficulty:  *totalDifficulty.BigInt(),
		ExtraData:        *extraData,
		Size:             size.Int64(),
		GasLimit:         *gasLimit.BigInt(),
		GasUsed:          *gasUsed.BigInt(),
		Timestamp:        timestamp.Int64(),
		TransactionsFull: etl,
		Uncles:           uncles,
	}, nil
}

type RPCEtherBlockRawWithTransactionHash struct {
	Number           string   `json:"number"`           // The block number, -1 if it is a pending block
	Hash             string   `json:"hash"`             // 32 Bytes - hash of the block. nil when its pending block
	ParentHash       string   `json:"parentHash"`       // 32 Bytes - hash of the parent block
	Nonce            string   `json:"nonce"`            // 8 Bytes - hash of the generated proof-of-work. nil when its pending block
	Sha3Uncles       string   `json:"sha3Uncles"`       // 32 Bytes - SHA3 of the uncles data in the block
	LogsBloom        string   `json:"logsBloom"`        // 256 Bytes - the bloom filter for the logs of the block.null when its pending block
	TransactionsRoot string   `json:"transactionsRoot"` // 32 Bytes - the root of the transaction trie of the block
	StateRoot        string   `json:"stateRoot"`        // 32 Bytes - the root of the final state trie of the block
	ReceiptsRoot     string   `json:"receiptsRoot"`     // 32 Bytes - the root of the receipts trie of the block
	Miner            string   `json:"miner"`            // 20 Bytes - the address of the beneficiary to whom the mining rewards were given
	Difficulty       string   `json:"difficulty"`       // integer of the difficulty for this block
	TotalDifficulty  string   `json:"totalDifficulty"`  // integer of the total difficulty of the chain until this block
	ExtraData        string   `json:"extraData"`        // the ‘extra data’ field of this block
	Size             string   `json:"size"`             // integer the size of this block in bytes
	GasLimit         string   `json:"gasLimit"`         // the maximum gas allowed in this block
	GasUsed          string   `json:"gasUsed"`          // the total used gas by all transactions in this block
	Timestamp        string   `json:"timestamp"`        // the unix timestamp for when the block was collated
	Transactions     []string `json:"transactions"`     // Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter
	Uncles           []string `json:"uncles"`           // Array - Array of uncle hashes
}

type RPCEtherBlockWithFullTransactions struct {
	Number           string                   `json:"number"`           // The block number, -1 if it is a pending block
	Hash             string                   `json:"hash"`             // 32 Bytes - hash of the block. nil when its pending block
	ParentHash       string                   `json:"parentHash"`       // 32 Bytes - hash of the parent block
	Nonce            string                   `json:"nonce"`            // 8 Bytes - hash of the generated proof-of-work. nil when its pending block
	Sha3Uncles       string                   `json:"sha3Uncles"`       // 32 Bytes - SHA3 of the uncles data in the block
	LogsBloom        string                   `json:"logsBloom"`        // 256 Bytes - the bloom filter for the logs of the block.null when its pending block
	TransactionsRoot string                   `json:"transactionsRoot"` // 32 Bytes - the root of the transaction trie of the block
	StateRoot        string                   `json:"stateRoot"`        // 32 Bytes - the root of the final state trie of the block
	ReceiptsRoot     string                   `json:"receiptsRoot"`     // 32 Bytes - the root of the receipts trie of the block
	Miner            string                   `json:"miner"`            // 20 Bytes - the address of the beneficiary to whom the mining rewards were given
	Difficulty       string                   `json:"difficulty"`       // integer of the difficulty for this block
	TotalDifficulty  string                   `json:"totalDifficulty"`  // integer of the total difficulty of the chain until this block
	ExtraData        string                   `json:"extraData"`        // the ‘extra data’ field of this block
	Size             string                   `json:"size"`             // integer the size of this block in bytes
	GasLimit         string                   `json:"gasLimit"`         // the maximum gas allowed in this block
	GasUsed          string                   `json:"gasUsed"`          // the total used gas by all transactions in this block
	Timestamp        string                   `json:"timestamp"`        // the unix timestamp for when the block was collated
	Transactions     []RPCEtherTransactionRaw `json:"transactions"`     // Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter
	Uncles           []string                 `json:"uncles"`           // Array - Array of uncle hashes
}
