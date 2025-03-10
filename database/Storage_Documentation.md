# Storage System Documentation

## Overview
The storage system implements a persistent storage layer for the blockchain simulator using LevelDB. It handles three main types of data: Blocks, Accounts, and Transactions.

## Data Types

### Block
- Represents a block in the blockchain
- Contains:
  - Index (uint64)
  - Timestamp (time.Time)
  - List of Transactions
  - Validator address
  - Previous block hash
  - Current block hash
  - State root hash

### Account
- Represents a blockchain account
- Contains:
  - Address ([]byte)
  - Balance (uint64)
  - Stake (uint64)
  - Nonce (uint64)

### Transaction
- Represents a blockchain transaction
- Contains:
  - From address ([]byte)
  - To address ([]byte)
  - Amount (uint64)
  - Fee (uint64)
  - Nonce (uint64)
  - Gas limit (uint64)
  - Gas price (uint64)
  - Signature ([]byte)
  - Public key (ecdsa.PublicKey)

## Storage Flow

### Block Storage
1. When a new block is added:
   - Block is validated
   - Block is serialized to protobuf format
   - Block is stored with prefix 'b' + block hash
   - Latest block hash is updated

2. When retrieving a block:
   - Block is fetched using block hash
   - Block is deserialized from protobuf format
   - Block is returned with all its transactions

### Account Storage
1. When updating an account:
   - Account is serialized to protobuf format
   - Account is stored with prefix 'a' + account address

2. When retrieving an account:
   - Account is fetched using account address
   - Account is deserialized from protobuf format
   - Account state is returned

### Transaction Storage
1. When storing a transaction:
   - Transaction is serialized to protobuf format
   - Transaction is stored with prefix 't' + transaction hash

2. When retrieving a transaction:
   - Transaction is fetched using transaction hash
   - Transaction is deserialized from protobuf format
   - Transaction details are returned

## State Management
1. When a new block is added:
   - All transactions in the block are processed
   - For each transaction:
     - From account is fetched and updated
     - To account is fetched and updated
     - Transaction is stored
   - State root is calculated and updated in block

## Key Prefixes
- Block: 'b'
- Account: 'a'
- Transaction: 't'
- Latest block: 'latest'

## Error Handling
- Storage operations return errors for:
  - Invalid data format
  - Missing data
  - Storage system failures
  - Database errors

## Dependencies
- LevelDB for persistent storage
- Protocol Buffers for data serialization
- ECDSA for cryptographic operations


## Note - TODOS

### Core Implementation
1. Implement `SerializePublicKey` and `ParsePublicKey` functions in types package
   - Currently returns empty/placeholder values
   - Need to implement proper ECDSA key serialization/deserialization

2. Complete `updateState` function in blockchain.go
   - Implement state transition logic
   - Add balance updates for transactions
   - Add stake management
   - Implement proper state root calculation

### Storage Implementation
1. Add Transaction Storage Methods
   - Implement `PutTransaction` in LevelDB storage
   - Implement `GetTransaction` in LevelDB storage
   - Add transaction indexing

2. Add Latest Block Tracking
   - Implement `GetLatestBlock` method
   - Add proper latest block hash management
   - Add block height tracking

### State Management
1. Implement Merkle Patricia Trie
   - Add state trie implementation
   - Implement state root calculation
   - Add state snapshot functionality

2. Add Account State Management
   - Implement account balance updates
   - Add stake management
   - Add nonce management
   - Implement account state rollback

