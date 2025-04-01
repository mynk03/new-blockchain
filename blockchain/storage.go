package blockchain

import (
	"blockchain-simulator/state"
	"blockchain-simulator/transaction"
)

// Storage defines the interface for persistent storage operations in the blockchain.
// It provides methods for managing blocks, state, and transactions.
type Storage interface {
	// Block operations
	PutBlock(block Block) error
	GetBlock(hash string) (Block, error)
	GetLatestBlock() (Block, error)

	// State operations
	PutState(stateRoot string, trie *state.MptTrie) error
	GetState(stateRoot string) (*state.MptTrie, error)

	// Transaction operations
	PutTransaction(tx transaction.Transaction) error

	// Transaction Getters
	GetTransaction(hash string) (transaction.Transaction, error)
	GetPendingTransactions() ([]transaction.Transaction, error)

	// Remove Transaction Operations
	RemoveTransaction(hash string) error
	RemoveBulkTransactions(hashes []string) error

	// Close
	Close() error
}
