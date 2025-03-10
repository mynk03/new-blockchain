package types

type Storage interface {
	PutBlock(block *Block) error
	GetBlock(hash []byte) (*Block, error)
	PutAccount(account *Account) error
	GetAccount(address []byte) (*Account, error)
	Close() error
}
