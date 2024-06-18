package core

import (
	"crypto/sha256"
)

type Transaction struct {
	ID   []byte
	Data []byte
	Code string // Smart Contract code
}

func (tx *Transaction) Hash() []byte {
	var hash [32]byte
	hash = sha256.Sum256(tx.Data)
	return hash[:]
}

func NewTransaction(data string, code string) *Transaction {
	tx := &Transaction{[]byte{}, []byte(data), code}
	tx.ID = tx.Hash()
	return tx
}
