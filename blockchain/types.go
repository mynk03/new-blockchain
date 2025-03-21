package blockchain

import (
	state "blockchain-simulator/state"

	"github.com/ethereum/go-ethereum/common"
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
	Chain      []Block          // Array of blocks
	StateTrie  *state.MptTrie      // Merkle Patricia Trie for account states
	Validators []common.Address // List of validators (for PoS or round-robin)
	Storage    Storage          // Add storage field
	TotalBlocks uint64           // Total number of blocks in the chain
}

// Transaction represents a transaction in the blockchain.
type Transaction struct {
	From   common.Address // Sender's address
	To     common.Address // Receiver's address
	Amount uint64         // Amount to transfer
	Nonce  uint64         // Sender's transaction count
}
