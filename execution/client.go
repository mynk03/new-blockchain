package execution

import (
	"context"
	"fmt"
	"sync"

	"blockchain-simulator/transactions"
	"blockchain-simulator/validator"
)

// ExecutionClient represents the execution client that broadcasts transactions to the network
type ExecutionClient struct {
	validators []*validator.Validator
	mu         sync.RWMutex
}

// NewExecutionClient creates a new execution client
func NewExecutionClient() *ExecutionClient {
	return &ExecutionClient{
		validators: make([]*validator.Validator, 0),
	}
}

// AddValidator adds a validator to the network
func (ec *ExecutionClient) AddValidator(v *validator.Validator) {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	ec.validators = append(ec.validators, v)
}

// BroadcastTransaction broadcasts a transaction to all validators in the network
func (ec *ExecutionClient) BroadcastTransaction(ctx context.Context, tx *transactions.Transaction) error {
	ec.mu.RLock()
	defer ec.mu.RUnlock()

	if len(ec.validators) == 0 {
		return fmt.Errorf("no validators in the network")
	}

	// Broadcast to all validators
	for _, v := range ec.validators {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := v.AddTransaction(*tx); err != nil {
				return fmt.Errorf("failed to add transaction to validator %s: %w", v.GetAddress(), err)
			}
		}
	}

	return nil
}

// GetValidatorCount returns the number of validators in the network
func (ec *ExecutionClient) GetValidatorCount() int {
	ec.mu.RLock()
	defer ec.mu.RUnlock()
	return len(ec.validators)
}

// GetValidators returns a copy of all validators in the network
func (ec *ExecutionClient) GetValidators() []*validator.Validator {
	ec.mu.RLock()
	defer ec.mu.RUnlock()

	validators := make([]*validator.Validator, len(ec.validators))
	copy(validators, ec.validators)
	return validators
}
