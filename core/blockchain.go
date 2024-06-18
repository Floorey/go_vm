package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks           []*Block
	VM               *VM // VM instance
	ValidatorManager *ValidatorManager
	PoH              *ProofOfHistory
}

func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	validator := bc.ValidatorManager.SelectValidator()

	newBlock := NewBlock(transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

	// Record PoH event
	event := "Block created by " + validator.Address
	bc.PoH.RecordEvent(event)

	// Execute all transactions in the block
	for _, tx := range transactions {
		err := bc.VM.LoadContract(string(tx.ID), tx.Code)
		if err != nil {
			log.Printf("Failed to load contract: %v", err)
			continue
		}

		result, err := bc.VM.Execute(string(tx.ID))
		if err != nil {
			log.Printf("Failed to execute contract: %v", err)
			continue
		}

		log.Printf("Transaction %x executed with result: %s", tx.ID, result)
	}
}

func (bc *Blockchain) AddTransaction(tx *Transaction) {
	// In a real implementation, transactions would be pooled and added to the next block
	bc.Blocks[len(bc.Blocks)-1].Transactions = append(bc.Blocks[len(bc.Blocks)-1].Transactions, tx)
}

func (bc *Blockchain) PrintChain() {
	for i, block := range bc.Blocks {
		fmt.Printf("Block %d\n", i)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("Previous Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Println("Transactions:")
		for _, tx := range block.Transactions {
			fmt.Printf("  Transaction ID: %x\n", tx.ID)
			fmt.Printf("  Data: %s\n", tx.Data)
			fmt.Printf("  Code: %s\n", tx.Code)
		}
		fmt.Println()
	}
}

func NewBlockchain(vm *VM, validatorManager *ValidatorManager, poh *ProofOfHistory) *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}, vm, validatorManager, poh}
}
