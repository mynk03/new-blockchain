package main

import (
	"fmt"
	"log"
	"time"

	"blockchain_simulator/database/internal/core"
	"blockchain_simulator/database/internal/storage"
	"blockchain_simulator/database/internal/types"
)

func main() {
	// Initialize the LevelDB storage
	dbPath := "./temp/data/blocks"
	db, err := storage.NewLevelDBStorage(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize LevelDB: %v", err)
	}
	defer db.Close()

	// Initialize blockchain
	blockchain := core.NewBlockchain(db)

	// Create a test block
	block := &types.Block{
		Index:        1,
		Timestamp:    time.Now(),
		Transactions: []*types.Transaction{},
		Validator:    []byte("validator1"),
		PrevHash:     []byte("previous_hash"),
		Hash:         []byte("current_hash"),
		StateRoot:    []byte("state_root"),
	}

	// Store the block
	err = blockchain.AddBlock(block)
	if err != nil {
		log.Fatalf("Failed to add block: %v", err)
	}

	// Retrieve and verify the block
	retrievedBlock, err := blockchain.GetBlock(block.Hash)
	if err != nil {
		log.Fatalf("Failed to retrieve block: %v", err)
	}

	fmt.Printf("Block stored and retrieved successfully:\n")
	fmt.Printf("Index: %d\n", retrievedBlock.Index)
	fmt.Printf("Timestamp: %v\n", retrievedBlock.Timestamp)
	fmt.Printf("Validator: %x\n", retrievedBlock.Validator)
	fmt.Printf("Hash: %x\n", retrievedBlock.Hash)
}
