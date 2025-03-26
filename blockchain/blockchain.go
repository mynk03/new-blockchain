package blockchain

import (
	"blockchain-simulator/state"
	"blockchain-simulator/storage"
	"blockchain-simulator/types"

	"time"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

// Blockchain represents the main blockchain structure
type Blockchain struct {
	Chain           []types.Block
	StateTrie       *state.MptTrie
	Validators      []common.Address
	Storage         storage.Storage
	LastBlockNumber uint64
}

func NewBlockchain(storage storage.Storage, accountsToFund []string, amountsToFund []uint64) *Blockchain {
	// Initialize the state trie
	stateTrie := state.NewMptTrie()

	// Create the genesis block
	genesisBlock := CreateGenesisBlock(accountsToFund, amountsToFund, stateTrie)

	// Store genesis block
	storage.PutBlock(genesisBlock)
	storage.PutState(genesisBlock.StateRoot, stateTrie)

	// Define validators (for PoS or round-robin)
	validators := make([]common.Address, len(accountsToFund))
	for i, addr := range accountsToFund {
		validators[i] = common.HexToAddress(addr)
	}

	return &Blockchain{
		Chain:           []types.Block{genesisBlock},
		StateTrie:       stateTrie,
		Validators:      validators,
		Storage:         storage,
		LastBlockNumber: genesisBlock.Index,
	}
}

// AddBlock adds a validated block to the chain and updates the state.
func (bc *Blockchain) AddBlock(newBlock types.Block) (bool, error) {
	// Store block and updated state
	if err := bc.Storage.PutBlock(newBlock); err != nil {
		return false, err
	}

	if err := bc.Storage.PutState(newBlock.StateRoot, bc.StateTrie); err != nil {
		return false, err
	}

	// Update the chain
	bc.Chain = append(bc.Chain, newBlock)
	bc.LastBlockNumber = newBlock.Index
	return true, nil
}

func (bc *Blockchain) GetLatestBlock() types.Block {
	return bc.Chain[bc.LastBlockNumber]
}

func (bc *Blockchain) GetLatestBlockHash() string {
	if len(bc.Chain) == 0 {
		return ""
	}
	return bc.Chain[bc.LastBlockNumber].Hash
}

func (bc *Blockchain) GetBlockByHash(hash string) types.Block {
	for _, block := range bc.Chain {
		if block.Hash == hash {
			return block
		}
	}
	log.WithFields(log.Fields{
		"type": "block_not_found",
		"hash": hash,
	}).Error("Block not found")
	return types.Block{}
}

// CreateGenesisBlock creates the first block in the chain
func CreateGenesisBlock(accountsToFund []string, amountsToFund []uint64, stateTrie *state.MptTrie) types.Block {
	// Create genesis block
	genesisBlock := types.Block{
		Index:        0,
		Timestamp:    time.Now().UTC().String(),
		Transactions: []types.Transaction{},
		PrevHash:     "",
		Hash:         "",
	}

	// Fund initial accounts
	for i, addr := range accountsToFund {
		account := &state.Account{
			Balance: amountsToFund[i],
			Nonce:   0,
		}
		stateTrie.PutAccount(common.HexToAddress(addr), account)
	}

	// Set state root
	genesisBlock.StateRoot = stateTrie.RootHash()
	genesisBlock.Hash = genesisBlock.GenerateHash()

	return genesisBlock
}

// ProcessBlock processes all transactions in a block and updates the state trie
func ProcessBlock(block types.Block, stateTrie *state.MptTrie) {
	for _, tx := range block.Transactions {
		ProcessTransaction(tx, stateTrie)
	}
}

// ProcessTransaction processes a single transaction and updates the state trie
func ProcessTransaction(tx types.Transaction, stateTrie *state.MptTrie) {
	// Get sender account
	senderAccount, _ := stateTrie.GetAccount(tx.From)
	if senderAccount == nil {
		senderAccount = &state.Account{Balance: 0, Nonce: 0}
	}

	// Get receiver account
	receiverAccount, _ := stateTrie.GetAccount(tx.To)
	if receiverAccount == nil {
		receiverAccount = &state.Account{Balance: 0, Nonce: 0}
	}

	// Update balances
	senderAccount.Balance -= tx.Amount
	receiverAccount.Balance += tx.Amount
	senderAccount.Nonce++

	// Store updated accounts
	stateTrie.PutAccount(tx.From, senderAccount)
	stateTrie.PutAccount(tx.To, receiverAccount)
}
