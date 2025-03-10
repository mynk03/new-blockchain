package core

import (
	"blockchain_simulator/database/internal/types"
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
	// Implement state transition logic
	transactions := block.Transactions
	accounts := make(map[string]*types.Account)

	for _, tx := range transactions {
		// Fetch accounts from storage
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
	}

	return nil
}

func isGenesisBlock(block *types.Block) bool {
	// Add genesis block validation logic
	if block.Index != 0 || block.PrevHash != nil || block.Hash == nil {
		return false
	}
	return true
}

// func (bc *Blockchain) GetLatestBlock() (*types.Block, error) {
// 	return bc.storage.GetLatestBlock()
// }
