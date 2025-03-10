package storage

import (
	pb "blockchain_simulator/database/internal/core/pb"
	"blockchain_simulator/database/internal/types"
	"bytes"

	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/protobuf/proto"
)

type LevelDBStorage struct {
	db *leveldb.DB
}

var (
	blockPrefix    = []byte{'b'}
	accountPrefix  = []byte{'a'}
	txPrefix       = []byte{'t'}
	latestBlockKey = []byte("latest")
)

func NewLevelDBStorage(path string) (*LevelDBStorage, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDBStorage{db: db}, nil
}

func (l *LevelDBStorage) PutBlock(block *types.Block) error {
	pbBlock := block.ToProto()
	data, err := proto.Marshal(pbBlock)
	if err != nil {
		return err
	}

	batch := new(leveldb.Batch)
	blockKey := bytes.Join([][]byte{blockPrefix, block.Hash}, nil)
	batch.Put(blockKey, data)
	batch.Put(latestBlockKey, block.Hash)
	return l.db.Write(batch, nil)
}

func (l *LevelDBStorage) GetBlock(hash []byte) (*types.Block, error) {
	blockKey := bytes.Join([][]byte{blockPrefix, hash}, nil)
	data, err := l.db.Get(blockKey, nil)
	if err != nil {
		return nil, err
	}

	var pbBlock pb.Block
	if err := proto.Unmarshal(data, &pbBlock); err != nil {
		return nil, err
	}

	block := &types.Block{}
	block.FromProto(&pbBlock)
	return block, nil
}

func (l *LevelDBStorage) PutAccount(account *types.Account) error {
	pbAcc := account.ToProto()
	data, err := proto.Marshal(pbAcc)
	if err != nil {
		return err
	}

	key := bytes.Join([][]byte{accountPrefix, account.Address}, nil)
	return l.db.Put(key, data, nil)
}

func (l *LevelDBStorage) GetAccount(address []byte) (*types.Account, error) {
	key := bytes.Join([][]byte{accountPrefix, address}, nil)
	data, err := l.db.Get(key, nil)
	if err != nil {
		return nil, err
	}

	var pbAcc pb.Account
	if err := proto.Unmarshal(data, &pbAcc); err != nil {
		return nil, err
	}

	account := &types.Account{}
	account.FromProto(&pbAcc)
	return account, nil
}

func (l *LevelDBStorage) Close() error {
	return l.db.Close()
}
