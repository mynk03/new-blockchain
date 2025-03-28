package execution

import (
	"context"
	"os"
	"testing"
	"time"

	"blockchain-simulator/blockchain"
	"blockchain-simulator/state"
	"blockchain-simulator/transactions"
	"blockchain-simulator/validator"
	"blockchain-simulator/wallet"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"
)

type ExecutionClientTestSuite struct {
	suite.Suite
	client         *ExecutionClient
	senderWallet   *wallet.MockWallet
	receiverWallet *wallet.MockWallet
	validator1     *validator.Validator
	validator2     *validator.Validator
	tempDirs       []string
}

func (suite *ExecutionClientTestSuite) SetupTest() {
	// Create execution client
	suite.client = NewExecutionClient()

	// Create wallets
	var err error
	suite.senderWallet, err = wallet.NewMockWallet()
	suite.NoError(err)

	suite.receiverWallet, err = wallet.NewMockWallet()
	suite.NoError(err)

	// Create transaction pools
	txPool1 := transactions.NewTransactionPool()
	txPool2 := transactions.NewTransactionPool()

	// Create temporary directories for storage
	dir1, err := os.MkdirTemp("", "validator1_*")
	suite.NoError(err)
	suite.tempDirs = append(suite.tempDirs, dir1)

	dir2, err := os.MkdirTemp("", "validator2_*")
	suite.NoError(err)
	suite.tempDirs = append(suite.tempDirs, dir2)

	// Create blockchains with unique storage paths
	storage1, err := blockchain.NewLevelDBStorage(dir1)
	suite.NoError(err)
	storage2, err := blockchain.NewLevelDBStorage(dir2)
	suite.NoError(err)

	chain1 := blockchain.NewBlockchain(storage1, []string{}, []uint64{})
	chain2 := blockchain.NewBlockchain(storage2, []string{}, []uint64{})

	// Set up state trie with sender's account
	account := state.Account{
		Balance: 1000,
		Nonce:   1,
	}
	err = chain1.StateTrie.PutAccount(suite.senderWallet.GetAddress(), &account)
	suite.NoError(err)
	err = chain2.StateTrie.PutAccount(suite.senderWallet.GetAddress(), &account)
	suite.NoError(err)

	// Create validators
	suite.validator1 = validator.NewValidator(suite.senderWallet.GetAddress(), txPool1, chain1)
	suite.validator2 = validator.NewValidator(suite.receiverWallet.GetAddress(), txPool2, chain2)

	// Add validators to the network
	suite.client.AddValidator(suite.validator1)
	suite.client.AddValidator(suite.validator2)
}

func (suite *ExecutionClientTestSuite) TearDownTest() {
	// Clean up storage
	if storage1, ok := suite.validator1.LocalChain.Storage.(*blockchain.LevelDBStorage); ok {
		storage1.Close()
	}
	if storage2, ok := suite.validator2.LocalChain.Storage.(*blockchain.LevelDBStorage); ok {
		storage2.Close()
	}

	// Remove temporary directories
	for _, dir := range suite.tempDirs {
		os.RemoveAll(dir)
	}
}

func TestExecutionClientTestSuite(t *testing.T) {
	suite.Run(t, new(ExecutionClientTestSuite))
}

func (suite *ExecutionClientTestSuite) TestBroadcastTransaction() {
	// Create a test transaction
	tx := transactions.Transaction{
		From:        suite.senderWallet.GetAddress(),
		To:          suite.receiverWallet.GetAddress(),
		Amount:      100,
		Nonce:       1,
		BlockNumber: 1,
		Timestamp:   uint64(time.Now().Unix()),
	}

	// Generate transaction hash and sign
	tx.TransactionHash = tx.GenerateHash()
	signature, err := suite.senderWallet.SignTransaction(common.HexToHash(tx.TransactionHash))
	suite.NoError(err)
	tx.Signature = signature

	// Broadcast transaction
	ctx := context.Background()
	err = suite.client.BroadcastTransaction(ctx, &tx)
	suite.NoError(err)

	// Verify transaction was added to both validators
	suite.True(suite.validator1.HasTransaction(tx.TransactionHash))
	suite.True(suite.validator2.HasTransaction(tx.TransactionHash))
}

func (suite *ExecutionClientTestSuite) TestBroadcastTransactionNoValidators() {
	// Create a new client without validators
	client := NewExecutionClient()

	// Create a test transaction
	tx := transactions.Transaction{
		From:        suite.senderWallet.GetAddress(),
		To:          suite.receiverWallet.GetAddress(),
		Amount:      100,
		Nonce:       1,
		BlockNumber: 1,
		Timestamp:   uint64(time.Now().Unix()),
	}

	// Try to broadcast transaction
	ctx := context.Background()
	err := client.BroadcastTransaction(ctx, &tx)
	suite.Error(err)
	suite.Equal("no validators in the network", err.Error())
}

func (suite *ExecutionClientTestSuite) TestAddValidator() {
	// Create a new validator
	newWallet, err := wallet.NewMockWallet()
	suite.NoError(err)

	txPool := transactions.NewTransactionPool()

	// Create temporary directory for storage
	dir, err := os.MkdirTemp("", "new_validator_*")
	suite.NoError(err)
	suite.tempDirs = append(suite.tempDirs, dir)

	storage, err := blockchain.NewLevelDBStorage(dir)
	suite.NoError(err)
	chain := blockchain.NewBlockchain(storage, []string{}, []uint64{})

	// Set up state trie with sender's account
	account := state.Account{
		Balance: 1000,
		Nonce:   1,
	}
	err = chain.StateTrie.PutAccount(suite.senderWallet.GetAddress(), &account)
	suite.NoError(err)

	newValidator := validator.NewValidator(newWallet.GetAddress(), txPool, chain)

	// Add validator to the network
	suite.client.AddValidator(newValidator)

	// Verify validator count increased
	suite.Equal(3, suite.client.GetValidatorCount())

	// Verify validator is in the list
	validators := suite.client.GetValidators()
	found := false
	for _, v := range validators {
		if v == newValidator {
			found = true
			break
		}
	}
	suite.True(found)
}
