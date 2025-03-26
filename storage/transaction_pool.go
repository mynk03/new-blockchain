package storage

import (
	"blockchain-simulator/types"
	"errors"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// TransactionPool manages pending transactions
type TransactionPool struct {
	storage Storage
	mu      sync.RWMutex
}

// NewTransactionPool creates a new transaction pool
func NewTransactionPool(storage Storage) (*TransactionPool, error) {
	return &TransactionPool{
		storage: storage,
	}, nil
}

// AddTransaction adds a transaction to the pool
func (tp *TransactionPool) AddTransaction(tx types.Transaction) error {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	// Validate transaction
	if err := tp.validateTransaction(tx); err != nil {
		return err
	}

	// Add to storage
	if err := tp.storage.PutTransaction(tx); err != nil {
		logrus.WithFields(logrus.Fields{
			"type":  "transaction_pool",
			"error": err,
		}).Error("Failed to add transaction to storage")
		return err
	}

	return nil
}

// GetPendingTransactions returns all pending transactions
func (tp *TransactionPool) GetPendingTransactions() []types.Transaction {
	tp.mu.RLock()
	defer tp.mu.RUnlock()

	txs, err := tp.storage.GetPendingTransactions()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"type":  "transaction_pool",
			"error": err,
		}).Error("Failed to get pending transactions")
		return []types.Transaction{}
	}

	return txs
}

// RemoveTransaction removes a transaction from the pool
func (tp *TransactionPool) RemoveTransaction(hash string) error {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	if err := tp.storage.RemoveTransaction(hash); err != nil {
		logrus.WithFields(logrus.Fields{
			"type":  "transaction_pool",
			"error": err,
		}).Error("Failed to remove transaction")
		return err
	}

	return nil
}

// RemoveBulkTransactions removes multiple transactions from the pool
func (tp *TransactionPool) RemoveBulkTransactions(hashes []string) error {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	if err := tp.storage.RemoveBulkTransactions(hashes); err != nil {
		logrus.WithFields(logrus.Fields{
			"type":  "transaction_pool",
			"error": err,
		}).Error("Failed to remove bulk transactions")
		return err
	}

	return nil
}

// validateTransaction validates a transaction
func (tp *TransactionPool) validateTransaction(tx types.Transaction) error {
	// Check if transaction already exists
	existingTx, err := tp.storage.GetTransaction(tx.TransactionHash)
	if err == nil && existingTx.TransactionHash != "" {
		return errors.New("transaction already exists")
	}

	// Check if transaction is too old (e.g., > 1 hour)
	if time.Now().Unix()-int64(tx.Timestamp) > 3600 {
		return errors.New("transaction too old")
	}

	// Check if sender has enough balance (this should be done with state trie)
	// This is a placeholder for now
	if tx.Amount == 0 {
		return errors.New("invalid amount")
	}

	return nil
}
