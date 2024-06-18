package main

import (
	"fmt"
	"time"
	"vm/core"
)

func main() {
	vm := core.NewVM()
	validatorManager := core.NewValidatorManager()
	poh := core.NewProofOfHistory()

	// Add validators
	validatorManager.AddValidator("Validator1", 100)
	validatorManager.AddValidator("Validator2", 50)

	bc := core.NewBlockchain(vm, validatorManager, poh)

	node := core.NewNode("localhost:3000", bc)
	go node.Start()
	time.Sleep(1 * time.Second) // Kurze Verzögerung, um sicherzustellen, dass der Server startet

	// Simulate adding transactions
	tx1 := core.NewTransaction("Transfer 100 Coins from Alice to Bob", string([]byte{byte(core.OP_TRANSFER)}))
	node.Blockchain.AddTransaction(tx1)

	// Print the blockchain
	fmt.Println("Blockchain:")
	node.Blockchain.PrintChain()
}
