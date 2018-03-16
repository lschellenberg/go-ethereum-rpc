package types

import (
	"math/big"
	"fmt"
)

type EtherBlock struct {
	Number           int64              `json:"number"`            // The block number, -1 if it is a pending block
	Hash             HexString          `json:"hash"`              // 32 Bytes - hash of the block. nil when its pending block
	ParentHash       HexString          `json:"parent_hash"`       // 32 Bytes - hash of the parent block
	Nonce            HexString          `json:"nonce"`             // 8 Bytes - hash of the generated proof-of-work. nil when its pending block
	Sha3Uncles       HexString          `json:"sha_3_uncles"`      // 32 Bytes - SHA3 of the uncles data in the block
	LogsBloom        HexString          `json:"logs_bloom"`        // 256 Bytes - the bloom filter for the logs of the block.null when its pending block
	TransactionsRoot HexString          `json:"transactions_root"` // 32 Bytes - the root of the transaction trie of the block
	StateRoot        HexString          `json:"state_root"`        // 32 Bytes - the root of the final state trie of the block
	ReceiptsRoot     HexString          `json:"receipts_root"`     // 32 Bytes - the root of the receipts trie of the block
	Miner            EtherAddress       `json:"miner"`             // 20 Bytes - the address of the beneficiary to whom the mining rewards were given
	Difficulty       big.Int            `json:"difficulty"`        // integer of the difficulty for this block
	TotalDifficulty  big.Int            `json:"total_difficulty"`  // integer of the total difficulty of the chain until this block
	ExtraData        HexString          `json:"extra_data"`        // the ‘extra data’ field of this block
	Size             int64              `json:"size"`              // integer the size of this block in bytes
	GasLimit         big.Int            `json:"gas_limit"`         // the maximum gas allowed in this block
	GasUsed          big.Int            `json:"gas_used"`          // the total used gas by all transactions in this block
	Timestamp        int64              `json:"timestamp"`         // the unix timestamp for when the block was collated
	Transactions     []HexString        `json:"transactions"`      // Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter
	TransactionsFull []EtherTransaction `json:"transactions_full"` // Array of transaction objects, or 32 Bytes transaction hashes depending on the last given parameter
	Uncles           []HexString        `json:"uncles"`            // Array - Array of uncle hashes
}

func (b1 EtherBlock) Compare(b2 *EtherBlock) error {
	if b1.Number != b2.Number {
		return fmt.Errorf("error in number: [1: %v,2: %v]", b1.Number, b2.Number)
	}
	if !b1.Hash.IsEqual(&b2.Hash) {
		return fmt.Errorf("error in hash: [1: %v,2: %v]", b1.Hash, b2.Hash)
	}
	if !b1.ParentHash.IsEqual(&b2.ParentHash) {
		return fmt.Errorf("error in ParentHash: [1: %v,2: %v]", b1.ParentHash, b2.ParentHash)
	}
	if !b1.Nonce.IsEqual(&b2.Nonce) {
		return fmt.Errorf("error in nonce: [1: %v,2: %v]", b1.Nonce, b2.Nonce)
	}
	if !b1.Sha3Uncles.IsEqual(&b2.Sha3Uncles) {
		return fmt.Errorf("error in sha3Uncles: [1: %v,2: %v]", b1.Sha3Uncles, b2.Sha3Uncles)
	}
	if !b1.LogsBloom.IsEqual(&b2.LogsBloom) {
		return fmt.Errorf("error in logsBloom: [1: %v,2: %v]", b1.LogsBloom, b2.LogsBloom)
	}
	if !b1.TransactionsRoot.IsEqual(&b2.TransactionsRoot) {
		return fmt.Errorf("error in transactionsRoot: [1: %v,2: %v]", b1.TransactionsRoot, b2.TransactionsRoot)
	}
	if !b1.StateRoot.IsEqual(&b2.StateRoot) {
		return fmt.Errorf("error in stateRoot: [1: %v,2: %v]", b1.StateRoot, b2.StateRoot)
	}
	if !b1.ReceiptsRoot.IsEqual(&b2.ReceiptsRoot) {
		return fmt.Errorf("error in receiptsRoot: [1: %v,2: %v]", b1.ReceiptsRoot, b2.ReceiptsRoot)
	}
	if b1.Miner.value != b2.Miner.value {
		return fmt.Errorf("error in miner: [1: %v,2: %v]", b1.Miner, b2.Miner)
	}
	if b1.Difficulty.Cmp(&b2.Difficulty) != 0 {
		return fmt.Errorf("error in difficulty: [1: %v,2: %v]", b1.Difficulty, b2.Difficulty)
	}
	if b1.TotalDifficulty.Cmp(&b2.TotalDifficulty) != 0 {
		return fmt.Errorf("error in totalDifficulty: [1: %v,2: %v]", b1.TotalDifficulty, b2.TotalDifficulty)
	}
	if !b1.ExtraData.IsEqual(&b2.ExtraData) {
		return fmt.Errorf("error in extraData: [1: %v,2: %v]", b1.ExtraData, b2.ExtraData)
	}
	if b1.Size != b2.Size {
		return fmt.Errorf("error in size: [1: %v,2: %v]", b1.Size, b2.Size)
	}
	if b1.GasLimit.Cmp(&b2.GasLimit) != 0 {
		return fmt.Errorf("error in gasLimit: [1: %v,2: %v]", b1.GasLimit, b2.GasLimit)
	}
	if b1.GasUsed.Cmp(&b2.GasUsed) != 0 {
		return fmt.Errorf("error in gasUsed: [1: %v,2: %v]", b1.GasUsed, b2.GasUsed)
	}
	if b1.Timestamp != b2.Timestamp {
		return fmt.Errorf("error in timestamp: [1: %v,2: %v]", b1.Timestamp, b2.Timestamp)
	}
	if err := CompareHexStringList(b1.Transactions, b2.Transactions); err != nil {
		return fmt.Errorf("error in transaction hash list: message, %v", err.Error())
	}
	if err := CompareEtherTransactionList(b1.TransactionsFull, b2.TransactionsFull, true); err != nil {
		return fmt.Errorf("error in transactionsFull: [1: %v,2: %v, message: %v]", len(b1.TransactionsFull), len(b2.TransactionsFull), err.Error())
	}
	if err := CompareHexStringList(b1.Uncles, b2.Uncles); err != nil {
		return fmt.Errorf("error in uncles: [1: %v,2: %v], message, %v", b1.Uncles, b2.Uncles, err.Error())
	}
	return nil
}
