package blockchain

import (
	"blockchain-simulator/state"
	"blockchain-simulator/transactions"
)

// UnifiedStorage combines blockchain and transaction storage into a single interface
type UnifiedStorage interface {
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

// UnifiedLevelDBStorage implements UnifiedStorage using LevelDB
type UnifiedLevelDBStorage struct {
	blockStorage       Storage
	transactionStorage transactions.TransactionStorage
}

// NewUnifiedLevelDBStorage creates a new unified storage instance
func NewUnifiedLevelDBStorage(dbPath string) (UnifiedStorage, error) {
	blockStorage, err := NewLevelDBStorage(dbPath)
	if err != nil {
		return nil, err
	}

	txStorage := transactions.InitializeStorage(dbPath)

	return &UnifiedLevelDBStorage{
		blockStorage:       blockStorage,
		transactionStorage: txStorage,
	}, nil
}

// Block operations
func (s *UnifiedLevelDBStorage) PutBlock(block Block) error {
	return s.blockStorage.PutBlock(block)
}

func (s *UnifiedLevelDBStorage) GetBlock(hash string) (Block, error) {
	return s.blockStorage.GetBlock(hash)
}

func (s *UnifiedLevelDBStorage) GetLatestBlock() (Block, error) {
	return s.blockStorage.GetLatestBlock()
}

// State operations
func (s *UnifiedLevelDBStorage) PutState(stateRoot string, trie *state.MptTrie) error {
	return s.blockStorage.PutState(stateRoot, trie)
}

func (s *UnifiedLevelDBStorage) GetState(stateRoot string) (*state.MptTrie, error) {
	return s.blockStorage.GetState(stateRoot)
}

// Transaction operations
func (s *UnifiedLevelDBStorage) PutTransaction(tx transactions.Transaction) error {
	return s.transactionStorage.PutTransaction(tx)
}

func (s *UnifiedLevelDBStorage) GetTransaction(hash string) (transactions.Transaction, error) {
	return s.transactionStorage.GetTransaction(hash)
}

func (s *UnifiedLevelDBStorage) GetPendingTransactions() ([]transactions.Transaction, error) {
	return s.transactionStorage.GetPendingTransactions()
}

func (s *UnifiedLevelDBStorage) GetAllTransactions() ([]transactions.Transaction, error) {
	return s.transactionStorage.GetAllTransactions()
}

func (s *UnifiedLevelDBStorage) RemoveTransaction(hash string) error {
	return s.transactionStorage.RemoveTransaction(hash)
}

func (s *UnifiedLevelDBStorage) RemoveBulkTransactions(hashes []string) error {
	return s.transactionStorage.RemoveBulkTransactions(hashes)
}

// Close closes both storages
func (s *UnifiedLevelDBStorage) Close() error {
	if err := s.blockStorage.Close(); err != nil {
		return err
	}
	return s.transactionStorage.Close()
}
