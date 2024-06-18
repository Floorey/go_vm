package core

import "crypto/sha256"

type Transaction struct {
	ID   []byte
	Data []byte
}

func (tx *Transaction) Hash() []byte {
	var hash [32]byte
	hash = sha256.Sum256(tx.Data)
	return hash[:]
}
func NewTransaction(data string) *Transaction {
	tx := &Transaction{[]byte{}, []byte(data)}
	tx.ID = tx.Hash()
	return tx
}
