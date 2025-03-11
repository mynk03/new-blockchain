package core

import (
	"blockchain_simulator/database/internal/types"
	"crypto/sha256"
	"encoding/json"
	"errors"
)

type Blockchain struct {
	storage types.Storage
}

func NewBlockchain(s types.Storage) *Blockchain {
	return &Blockchain{storage: s}
}

func (bc *Blockchain) AddBlock(block *types.Block) error {
	// Validate block structure
	if block.Index == 0 && !isGenesisBlock(block) {
		return errors.New("invalid genesis block")
	}

	// Store block
	if err := bc.storage.PutBlock(block); err != nil {
		return err
	}

	// Update state
	return bc.updateState(block)
}

func (bc *Blockchain) GetBlock(hash []byte) (*types.Block, error) {
	return bc.storage.GetBlock(hash)
}

func (bc *Blockchain) updateState(block *types.Block) error {
	transactions := block.Transactions
	accounts := make(map[string]*types.Account)

	for _, tx := range transactions {
		fromAcc, err := bc.storage.GetAccount(tx.From)
		if err != nil {
			return err
		}
		toAcc, err := bc.storage.GetAccount(tx.To)
		if err != nil {
			return err
		}
		accounts[string(tx.From)] = fromAcc
		accounts[string(tx.To)] = toAcc

		// Update balances
		fromAcc.Balance -= tx.Amount + tx.Fee
		toAcc.Balance += tx.Amount

		// Update nonces
		fromAcc.Nonce++
	}

	// Store updated accounts
	for _, acc := range accounts {
		if err := bc.storage.PutAccount(acc); err != nil {
			return err
		}
	}

	// Calculate and update state root
	stateRoot, err := bc.CalculateStateRoot()
	if err != nil {
		return err
	}
	block.StateRoot = stateRoot

	return nil
}

func isGenesisBlock(block *types.Block) bool {
	// Add genesis block validation logic
	if block.Index != 0 || block.PrevHash != nil || block.Hash == nil {
		return false
	}
	return true
}

func (bc *Blockchain) CalculateStateRoot() ([]byte, error) {
	// TODO: Implement proper Merkle Patricia Trie
	// For now, return a simple hash of all account states
	accounts := make([]*types.Account, 0)
	// Get all accounts and create a hash
	hash := sha256.New()
	jsonData, err := json.Marshal(accounts)
	if err != nil {
		return nil, err
	}
	hash.Write(jsonData)
	return hash.Sum(nil), nil
}

// func (bc *Blockchain) GetLatestBlock() (*types.Block, error) {
// 	return bc.storage.GetLatestBlock()
// }
