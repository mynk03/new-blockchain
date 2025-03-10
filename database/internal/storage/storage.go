package storage

import "blockchain_simulator/database/internal/types"

type Storage interface {
	PutBlock(block *types.Block) error
	GetBlock(hash []byte) (*types.Block, error)
	GetLatestBlock() (*types.Block, error)

	PutAccount(account *types.Account) error
	GetAccount(address []byte) (*types.Account, error)

	PutTransaction(tx *types.Transaction) error
	GetTransaction(txHash []byte) (*types.Transaction, error)

	Close() error
}
