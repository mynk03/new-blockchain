package blockchain

import (
	"blockchain-simulator/state"
	"blockchain-simulator/transactions"
)

// Storage defines the interface for blockchain persistence
type Storage interface {
	// Block operations
	PutBlock(block Block) error
	GetBlock(hash string) (Block, error)
	GetLatestBlock() (Block, error)

	// State operations
	PutState(stateRoot string, trie *state.MptTrie) error
	GetState(stateRoot string) (*state.MptTrie, error)

	// Transaction operations
	PutTransaction(tx transactions.Transaction) error
	GetTransaction(hash string) (transactions.Transaction, error)
	GetPendingTransactions() ([]transactions.Transaction, error)
	GetAllTransactions() ([]transactions.Transaction, error)
	RemoveTransaction(hash string) error
	RemoveBulkTransactions(hashes []string) error

	// Close
	Close() error
}
