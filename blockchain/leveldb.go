package blockchain

import (
	"blockchain-simulator/state"
	"blockchain-simulator/transactions"
	"encoding/json"

	"github.com/syndtr/goleveldb/leveldb"
)

const (
	blockPrefix       = "b:"      // Prefix for block data
	statePrefix       = "s:"      // Prefix for state data
	transactionPrefix = "t:"      // Prefix for transaction data
	pendingKey        = "pending" // Key for pending transactions
	allKey            = "all"     // Key for all transactions
)

// LevelDBStorage implements Storage interface using LevelDB
type LevelDBStorage struct {
	db *leveldb.DB
}

// NewLevelDBStorage creates a new LevelDB storage instance
func NewLevelDBStorage(path string) (Storage, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDBStorage{db: db}, nil
}

// Block operations
func (s *LevelDBStorage) PutBlock(block Block) error {
	data, err := json.Marshal(block)
	if err != nil {
		return err
	}
	return s.db.Put([]byte(blockPrefix+block.Hash), data, nil)
}

func (s *LevelDBStorage) GetBlock(hash string) (Block, error) {
	data, err := s.db.Get([]byte(blockPrefix+hash), nil)
	if err != nil {
		return Block{}, err
	}
	var block Block
	err = json.Unmarshal(data, &block)
	return block, err
}

func (s *LevelDBStorage) GetLatestBlock() (Block, error) {
	iter := s.db.NewIterator(nil, nil)
	defer iter.Release()

	var latestBlock Block
	for iter.Next() {
		key := string(iter.Key())
		if len(key) > len(blockPrefix) && key[:len(blockPrefix)] == blockPrefix {
			var block Block
			if err := json.Unmarshal(iter.Value(), &block); err != nil {
				continue
			}
			if block.Index > latestBlock.Index {
				latestBlock = block
			}
		}
	}
	return latestBlock, nil
}

// State operations
func (s *LevelDBStorage) PutState(stateRoot string, trie *state.MptTrie) error {
	data, err := json.Marshal(trie)
	if err != nil {
		return err
	}
	return s.db.Put([]byte(statePrefix+stateRoot), data, nil)
}

func (s *LevelDBStorage) GetState(stateRoot string) (*state.MptTrie, error) {
	data, err := s.db.Get([]byte(statePrefix+stateRoot), nil)
	if err != nil {
		return nil, err
	}
	var trie state.MptTrie
	err = json.Unmarshal(data, &trie)
	return &trie, err
}

// Transaction operations
func (s *LevelDBStorage) PutTransaction(tx transactions.Transaction) error {
	data, err := json.Marshal(tx)
	if err != nil {
		return err
	}

	// Store transaction
	if err := s.db.Put([]byte(transactionPrefix+tx.TransactionHash), data, nil); err != nil {
		return err
	}

	// Update pending transactions list
	pendingTxs, err := s.GetPendingTransactions()
	if err != nil {
		pendingTxs = []transactions.Transaction{}
	}
	pendingTxs = append(pendingTxs, tx)
	pendingData, err := json.Marshal(pendingTxs)
	if err != nil {
		return err
	}
	if err := s.db.Put([]byte(pendingKey), pendingData, nil); err != nil {
		return err
	}

	// Update all transactions list
	allTxs, err := s.GetAllTransactions()
	if err != nil {
		allTxs = []transactions.Transaction{}
	}
	allTxs = append(allTxs, tx)
	allData, err := json.Marshal(allTxs)
	if err != nil {
		return err
	}
	return s.db.Put([]byte(allKey), allData, nil)
}

func (s *LevelDBStorage) GetTransaction(hash string) (transactions.Transaction, error) {
	data, err := s.db.Get([]byte(transactionPrefix+hash), nil)
	if err != nil {
		return transactions.Transaction{}, err
	}
	var tx transactions.Transaction
	err = json.Unmarshal(data, &tx)
	return tx, err
}

func (s *LevelDBStorage) GetPendingTransactions() ([]transactions.Transaction, error) {
	data, err := s.db.Get([]byte(pendingKey), nil)
	if err != nil {
		return nil, err
	}
	var txs []transactions.Transaction
	err = json.Unmarshal(data, &txs)
	return txs, err
}

func (s *LevelDBStorage) GetAllTransactions() ([]transactions.Transaction, error) {
	data, err := s.db.Get([]byte(allKey), nil)
	if err != nil {
		return nil, err
	}
	var txs []transactions.Transaction
	err = json.Unmarshal(data, &txs)
	return txs, err
}

func (s *LevelDBStorage) RemoveTransaction(hash string) error {
	// Get pending transactions
	pendingTxs, err := s.GetPendingTransactions()
	if err != nil {
		return err
	}

	// Remove from pending transactions
	newPending := make([]transactions.Transaction, 0)
	for _, tx := range pendingTxs {
		if tx.TransactionHash != hash {
			newPending = append(newPending, tx)
		}
	}

	// Update pending transactions
	pendingData, err := json.Marshal(newPending)
	if err != nil {
		return err
	}
	if err := s.db.Put([]byte(pendingKey), pendingData, nil); err != nil {
		return err
	}

	// Remove transaction from storage
	return s.db.Delete([]byte(transactionPrefix+hash), nil)
}

func (s *LevelDBStorage) RemoveBulkTransactions(hashes []string) error {
	for _, hash := range hashes {
		if err := s.RemoveTransaction(hash); err != nil {
			return err
		}
	}
	return nil
}

// Close closes the database connection
func (s *LevelDBStorage) Close() error {
	return s.db.Close()
}
