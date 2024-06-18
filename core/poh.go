package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type ProofOfHistory struct {
	LastHash string
}

func NewProofOfHistory() *ProofOfHistory {
	return &ProofOfHistory{LastHash: ""}
}
func (poh *ProofOfHistory) RecordEvent(event string) string {
	timestamp := time.Now().String()
	data := timestamp + event + poh.LastHash
	hash := sha256.Sum256([]byte(data))
	poh.LastHash = hex.EncodeToString(hash[:])
	return poh.LastHash

}
