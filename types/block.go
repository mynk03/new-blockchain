package types

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index        uint64
	Timestamp    string
	Transactions []Transaction
	PrevHash     string
	Hash         string
	StateRoot    string
}

// CreateBlock creates a new block
func CreateBlock(transactions []Transaction, prevBlock Block) Block {
	newBlock := Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().UTC().String(),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
		Hash:         "", // Populated later
	}
	return newBlock
}

// GenerateHash generates a hash for the block
func (b *Block) GenerateHash() string {
	data, _ := json.Marshal(b)
	hashBytes := sha256.Sum256(data)
	return hex.EncodeToString(hashBytes[:])
}
