package core

import "log"

type Blockchain struct {
	Blocks []*Block
	VM     *VM // VM instance
}

func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

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

func NewBlockchain(vm *VM) *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}, vm}
}
