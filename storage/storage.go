package storage

import (
	"blockchain-simulator/state"
	"blockchain-simulator/types"
)

// Storage defines the interface for blockchain persistence
type Storage interface {
	// Block operations
	PutBlock(block types.Block) error
	GetBlock(hash string) (types.Block, error)
	GetLatestBlock() (types.Block, error)

	// State operations
	PutState(stateRoot string, trie *state.MptTrie) error
	GetState(stateRoot string) (*state.MptTrie, error)

	// Transaction operations
	PutTransaction(tx types.Transaction) error
	GetTransaction(hash string) (types.Transaction, error)
	GetPendingTransactions() ([]types.Transaction, error)
	GetAllTransactions() ([]types.Transaction, error)
	RemoveTransaction(hash string) error
	RemoveBulkTransactions(hashes []string) error

	// Close
	Close() error
}
