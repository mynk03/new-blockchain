package types

type Storage interface {
	PutBlock(block *Block) error
	GetBlock(hash []byte) (*Block, error)
	GetLatestBlock() (*Block, error)
	
	PutAccount(account *Account) error
	GetAccount(address []byte) (*Account, error)
	
	PutTransaction(tx *Transaction) error
	GetTransaction(txHash []byte) (*Transaction, error)
	
	Close() error
}
