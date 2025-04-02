// Copyright (c) 2025 ANCILAR
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package transactions

import (
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	ErrDuplicateTransaction = errors.New("transaction already exists in pool")
)

// TransactionPool represents a pool of pending transactions
type TransactionPool struct {
	transactions map[string]Transaction
	mu           sync.RWMutex
}

// NewTransactionPool creates a new transaction pool
func NewTransactionPool() *TransactionPool {
	return &TransactionPool{
		transactions: make(map[string]Transaction),
	}
}

// AddTransaction adds a transaction to the pool
func (pool *TransactionPool) AddTransaction(tx Transaction) error {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	// Validate transaction
	if status, err := tx.Validate(); !status {
		return err
	}

	// Check if transaction already exists
	if _, exists := pool.transactions[tx.TransactionHash]; exists {
		return ErrDuplicateTransaction
	}

	pool.transactions[tx.TransactionHash] = tx
	return nil
}

// GetTransaction retrieves a transaction by its hash
func (pool *TransactionPool) GetTransaction(hash string) (Transaction, bool) {
	pool.mu.RLock()
	defer pool.mu.RUnlock()
	tx, exists := pool.transactions[hash]
	return tx, exists
}

// RemoveTransaction removes a transaction from the pool
func (pool *TransactionPool) RemoveTransaction(hash string) error {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	if _, exists := pool.transactions[hash]; !exists {
		return errors.New("transaction not found")
	}

	delete(pool.transactions, hash)
	return nil
}

// GetAllTransactions returns all transactions in the pool
func (pool *TransactionPool) GetAllTransactions() []Transaction {
	pool.mu.RLock()
	defer pool.mu.RUnlock()

	txs := make([]Transaction, 0, len(pool.transactions))
	for _, tx := range pool.transactions {
		txs = append(txs, tx)
	}
	return txs
}

// HasTransaction checks if a transaction exists in the pool
func (pool *TransactionPool) HasTransaction(hash string) bool {
	pool.mu.RLock()
	defer pool.mu.RUnlock()
	_, exists := pool.transactions[hash]
	return exists
}

// RemoveBulkTransactions removes multiple transactions from the PendingTransactions
func (tp *TransactionPool) RemoveBulkTransactions(hashes []string) {
	for _, hash := range hashes {
		err := tp.RemoveTransaction(hash)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"type":     "transaction_pool",
				"error":    err,
				"txn_hash": hash,
			}).Error("failed to remove transaction")
		}
	}
}

// GetPendingTransactions returns the PendingTransactions
func (tp *TransactionPool) GetPendingTransactions() []Transaction {
	return tp.GetAllTransactions()
}

// GetTransactionByHash returns a transaction by hash
func (tp *TransactionPool) GetTransactionByHash(hash string) *Transaction {
	tx, exists := tp.GetTransaction(hash)
	if exists {
		return &tx
	}
	return nil
}
