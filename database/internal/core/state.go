package core

import "blockchain_simulator/database/internal/types"

type StateManager struct {
	storage types.Storage
}

func NewStateManager(s types.Storage) *StateManager {
	return &StateManager{storage: s}
}

func (sm *StateManager) GetAccount(address []byte) (*types.Account, error) {
	return sm.storage.GetAccount(address)
}

func (sm *StateManager) UpdateAccount(account *types.Account) error {
	return sm.storage.PutAccount(account)
}

func (sm *StateManager) CalculateStateRoot() ([]byte, error) {
	// Implement Merkle Patricia Trie logic
	return []byte{}, nil
}
