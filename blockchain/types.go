package blockchain

import (
	state "blockchain-simulator/state"
)

// Block represents a block in the blockchain.
type Block struct {
	Index        int           // Block height
	Timestamp    string        // Time of creation
	Transactions []Transaction // Transactions in the block
	PrevHash     string        // Hash of the previous block
	Hash         string        // Hash of the current block
	StateRoot    string        // Root hash of the state trie after applying transactions
	Validator    string        // Address of the validator who created the block (for PoS)
}

// Blockchain represents the entire chain.
type Blockchain struct {
	Chain      []Block       // Array of blocks
	StateTrie  *state.Trie   // Merkle Patricia Trie for account states
	PendingTxs []Transaction // Pending transactions (transaction pool)
	Validators []string      // List of validators (for PoS or round-robin)
}

// Transaction represents a transaction in the blockchain.
type Transaction struct {
	From   string // Sender's address
	To     string // Receiver's address
	Amount int    // Amount to transfer
	Nonce  int    // Sender's transaction count
}
