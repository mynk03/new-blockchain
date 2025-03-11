package types

import (
	pb "blockchain_simulator/database/internal/core/pb"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Block struct {
	Index        uint64
	Timestamp    time.Time
	Transactions []*Transaction
	Validator    []byte
	PrevHash     []byte
	Hash         []byte
	StateRoot    []byte
}

func (b *Block) ToProto() *pb.Block {
	txs := make([]*pb.Transaction, len(b.Transactions))
	for i, tx := range b.Transactions {
		txs[i] = tx.ToProto()
	}
	return &pb.Block{
		Index:        b.Index,
		Timestamp:    timestamppb.New(b.Timestamp),
		Transactions: txs,
		Validator:    b.Validator,
		PrevHash:     b.PrevHash,
		Hash:         b.Hash,
		StateRoot:    b.StateRoot,
	}
}

func (b *Block) FromProto(pbBlock *pb.Block) {
	b.Index = pbBlock.Index
	b.Timestamp = pbBlock.Timestamp.AsTime()
	b.Transactions = make([]*Transaction, len(pbBlock.Transactions))
	for i, tx := range pbBlock.Transactions {
		b.Transactions[i] = &Transaction{}
		b.Transactions[i].FromProto(tx)
	}
	b.Validator = pbBlock.Validator
	b.PrevHash = pbBlock.PrevHash
	b.Hash = pbBlock.Hash
	b.StateRoot = pbBlock.StateRoot
}

type Account struct {
	Address []byte
	Balance uint64
	Stake   uint64
	Nonce   uint64
}

func (a *Account) ToProto() *pb.Account {
	return &pb.Account{
		Address: a.Address,
		Balance: a.Balance,
		Stake:   a.Stake,
		Nonce:   a.Nonce,
	}
}

func (a *Account) FromProto(pbAcc *pb.Account) {
	a.Address = pbAcc.Address
	a.Balance = pbAcc.Balance
	a.Stake = pbAcc.Stake
	a.Nonce = pbAcc.Nonce
}

type Transaction struct {
	From      []byte
	To        []byte
	Amount    uint64
	Fee       uint64
	Nonce     uint64
	GasLimit  uint64
	GasPrice  uint64
	Signature []byte
	PublicKey *ecdsa.PublicKey
}

func (t *Transaction) ToProto() *pb.Transaction {
	pubKey, _ := SerializePublicKey(t.PublicKey)
	return &pb.Transaction{
		From:      t.From,
		To:        t.To,
		Amount:    t.Amount,
		Fee:       t.Fee,
		Nonce:     t.Nonce,
		GasLimit:  t.GasLimit,
		GasPrice:  t.GasPrice,
		Signature: t.Signature,
		PublicKey: pubKey,
	}
}

func (t *Transaction) FromProto(pbTx *pb.Transaction) error {
	pubKey, err := ParsePublicKey(pbTx.PublicKey)
	if err != nil {
		return err
	}

	t.From = pbTx.From
	t.To = pbTx.To
	t.Amount = pbTx.Amount
	t.Fee = pbTx.Fee
	t.Nonce = pbTx.Nonce
	t.GasLimit = pbTx.GasLimit
	t.GasPrice = pbTx.GasPrice
	t.Signature = pbTx.Signature
	t.PublicKey = pubKey
	return nil
}

func (t *Transaction) Hash() []byte {
	// Create a hash of the transaction data
	hash := sha256.New()
	hash.Write(t.From)
	hash.Write(t.To)
	hash.Write([]byte{byte(t.Amount)})
	hash.Write([]byte{byte(t.Fee)})
	hash.Write([]byte{byte(t.Nonce)})
	hash.Write([]byte{byte(t.GasLimit)})
	hash.Write([]byte{byte(t.GasPrice)})
	return hash.Sum(nil)
}

func SerializePublicKey(pub *ecdsa.PublicKey) ([]byte, error) {
	if pub == nil {
		return nil, errors.New("public key is nil")
	}
	der, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: der,
	}
	return pem.EncodeToMemory(block), nil
}

func ParsePublicKey(data []byte) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch pub := pub.(type) {
	case *ecdsa.PublicKey:
		return pub, nil
	default:
		return nil, errors.New("not ECDSA public key")
	}
}
