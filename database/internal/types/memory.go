package types

import "errors"

// InMemoryStorage is an in-memory implementation of the Storage interface.
type InMemoryStorage struct {
	blocks     map[string]*Block
	accounts   map[string]*Account
	mem_pool   map[string]*Transaction
	last_block *Block
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		blocks:     make(map[string]*Block),
		accounts:   make(map[string]*Account),
		mem_pool:   make(map[string]*Transaction),
		last_block: nil,
	}
}

func (s *InMemoryStorage) PutBlock(block *Block) error {
	s.blocks[string(block.Hash)] = block
	s.last_block = block
	return nil
}

func (s *InMemoryStorage) GetBlock(hash []byte) (*Block, error) {
	block, exists := s.blocks[string(hash)]
	if !exists {
		return nil, errors.New("block not found")
	}
	return block, nil
}

func (s *InMemoryStorage) PutAccount(account *Account) error {
	s.accounts[string(account.Address)] = account
	return nil
}

func (s *InMemoryStorage) GetAccount(address []byte) (*Account, error) {
	account, exists := s.accounts[string(address)]
	if !exists {
		return nil, errors.New("account not found")
	}
	return account, nil
}

func (s *InMemoryStorage) PutTransaction(tx *Transaction) error {
	s.mem_pool[string(tx.Hash())] = tx
	return nil
}

func (s *InMemoryStorage) GetTransaction(hash []byte) (*Transaction, error) {
	tx, exists := s.mem_pool[string(hash)]
	if !exists {
		return nil, errors.New("transaction not found")
	}
	return tx, nil
}

func (s *InMemoryStorage) GetLatestBlock() (*Block, error) {
	if len(s.blocks) == 0 {
		return nil, errors.New("no blocks in storage")
	}

	return s.last_block, nil
}
