package rpc

import (
	"encoding/json"
	"github.com/Leondroids/go-ethereum-rpc/types"
	"fmt"
)

func (client *Client) RequestEtherBlock(method string, params ...interface{}) (*types.EtherBlock, error) {
	p := []interface{}(params)

	if len(p) != 2 {
		return nil, fmt.Errorf("wrong params in reqeusting ether block: %v", p)
	}

	full := p[1].(bool)

	response, err := checkRPCError(client.Call(method, params...))

	if err != nil {
		return nil, err
	}

	return getBlockFromResponse(response.Result, full)
}

func getBlockFromResponse(response interface{}, full bool) (*types.EtherBlock, error) {
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

func (rpcBlock RPCEtherBlockRawWithTransactionHash) ToEtherBlock() (*types.EtherBlock, error) {
	number, err := types.NewHexString(rpcBlock.Number)

	if err != nil {
		return nil, err
	}

	hash, err := types.NewHexString(rpcBlock.Hash)

	if err != nil {
		return nil, err
	}

	parentHash, err := types.NewHexString(rpcBlock.ParentHash)

	if err != nil {
		return nil, err
	}

	nonce, err := types.NewHexString(rpcBlock.Nonce)

	if err != nil {
		return nil, err
	}

	sha3Uncles, err := types.NewHexString(rpcBlock.Sha3Uncles)

	if err != nil {
		return nil, err
	}

	logsBloom, err := types.NewHexString(rpcBlock.LogsBloom)

	if err != nil {
		return nil, err
	}

	transactionsRoot, err := types.NewHexString(rpcBlock.TransactionsRoot)

	if err != nil {
		return nil, err
	}

	stateRoot, err := types.NewHexString(rpcBlock.StateRoot)

	if err != nil {
		return nil, err
	}

	receiptsRoot, err := types.NewHexString(rpcBlock.ReceiptsRoot)

	if err != nil {
		return nil, err
	}

	miner, err := new(types.EtherAddress).FromString(rpcBlock.Miner)

	if err != nil {
		return nil, err
	}

	difficulty, err := types.NewHexString(rpcBlock.Difficulty)

	if err != nil {
		return nil, err
	}

	totalDifficulty, err := types.NewHexString(rpcBlock.TotalDifficulty)

	if err != nil {
		return nil, err
	}

	extraData, err := types.NewHexString(rpcBlock.ExtraData)

	if err != nil {
		return nil, err
	}

	size, err := types.NewHexString(rpcBlock.Size)

	if err != nil {
		return nil, err
	}

	gasLimit, err := types.NewHexString(rpcBlock.GasLimit)

	if err != nil {
		return nil, err
	}

	gasUsed, err := types.NewHexString(rpcBlock.GasUsed)

	if err != nil {
		return nil, err
	}

	timestamp, err := types.NewHexString(rpcBlock.Timestamp)

	if err != nil {
		return nil, err
	}

	transactions, err := types.ToHexStringList(rpcBlock.Transactions)

	if err != nil {
		return nil, err
	}

	uncles, err := types.ToHexStringList(rpcBlock.Uncles)

	if err != nil {
		return nil, err
	}

	return &types.EtherBlock{
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

func (rpcBlock RPCEtherBlockWithFullTransactions) ToEtherBlock() (*types.EtherBlock, error) {
	number, err := types.NewHexString(rpcBlock.Number)

	if err != nil {
		return nil, err
	}

	hash, err := types.NewHexString(rpcBlock.Hash)

	if err != nil {
		return nil, err
	}

	parentHash, err := types.NewHexString(rpcBlock.ParentHash)

	if err != nil {
		return nil, err
	}

	nonce, err := types.NewHexString(rpcBlock.Nonce)

	if err != nil {
		return nil, err
	}

	sha3Uncles, err := types.NewHexString(rpcBlock.Sha3Uncles)

	if err != nil {
		return nil, err
	}

	logsBloom, err := types.NewHexString(rpcBlock.LogsBloom)

	if err != nil {
		return nil, err
	}

	transactionsRoot, err := types.NewHexString(rpcBlock.TransactionsRoot)

	if err != nil {
		return nil, err
	}

	stateRoot, err := types.NewHexString(rpcBlock.StateRoot)

	if err != nil {
		return nil, err
	}

	receiptsRoot, err := types.NewHexString(rpcBlock.ReceiptsRoot)

	if err != nil {
		return nil, err
	}

	miner, err := new(types.EtherAddress).FromString(rpcBlock.Miner)

	if err != nil {
		return nil, err
	}

	difficulty, err := types.NewHexString(rpcBlock.Difficulty)

	if err != nil {
		return nil, err
	}

	totalDifficulty, err := types.NewHexString(rpcBlock.TotalDifficulty)

	if err != nil {
		return nil, err
	}

	extraData, err := types.NewHexString(rpcBlock.ExtraData)

	if err != nil {
		return nil, err
	}

	size, err := types.NewHexString(rpcBlock.Size)

	if err != nil {
		return nil, err
	}

	gasLimit, err := types.NewHexString(rpcBlock.GasLimit)

	if err != nil {
		return nil, err
	}

	gasUsed, err := types.NewHexString(rpcBlock.GasUsed)

	if err != nil {
		return nil, err
	}

	timestamp, err := types.NewHexString(rpcBlock.Timestamp)

	if err != nil {
		return nil, err
	}


	etl := make([]types.EtherTransaction, len(rpcBlock.Transactions))
	for k, v := range rpcBlock.Transactions {
		et, err := v.toEtherTransaction()

		if err != nil {
			return nil, err
		}
		etl[k] = *et
	}

	uncles, err := types.ToHexStringList(rpcBlock.Uncles)

	if err != nil {
		return nil, err
	}

	return &types.EtherBlock{
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
