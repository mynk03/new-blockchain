package types

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"blockchain-simulator/state"

	"github.com/ethereum/go-ethereum/common"
)

// Transaction represents a blockchain transaction
type Transaction struct {
	From            common.Address
	To              common.Address
	Amount          uint64
	Nonce           uint64
	BlockNumber     uint32
	Timestamp       uint64
	TransactionHash string
}

// GenerateHash generates a hash for the transaction
func (tx *Transaction) GenerateHash() string {
	data, _ := json.Marshal(tx)
	hashBytes := sha256.Sum256(data)
	return hex.EncodeToString(hashBytes[:])
}

// Validate checks if the transaction is valid
func (tx *Transaction) Validate() (bool, error) {
	if tx.Amount == 0 {
		return false, nil
	}
	if tx.From == (common.Address{}) {
		return false, nil
	}
	if tx.To == (common.Address{}) {
		return false, nil
	}
	if time.Now().Unix()-int64(tx.Timestamp) > 3600 {
		return false, nil
	}
	return true, nil
}

// ValidateWithState validates the transaction against the current state
func (tx *Transaction) ValidateWithState(stateTrie *state.MptTrie) (bool, error) {
	// Basic validation
	if valid, err := tx.Validate(); !valid {
		return false, err
	}

	// Check if sender has enough balance
	fromBalance := stateTrie.GetBalance(tx.From)
	if fromBalance < tx.Amount {
		return false, nil
	}

	return true, nil
}
