package main

import (
	"fmt"
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

	// Create accounts
	vm.AccountManager.CreateAccount("Alice", 1000)
	vm.AccountManager.CreateAccount("Bob", 500)

	// Print initial balances
	fmt.Println("Initial Balances:")
	printBalances(vm)

	// Create and add transactions
	tx1 := core.NewTransaction("Transfer 100 Coins from Alice to Bob", string([]byte{byte(core.OP_TRANSFER)}))
	copy(vm.Memory[0:], []byte("Alice"))
	copy(vm.Memory[32:], []byte("Bob"))
	vm.Registers["amount"] = 100
	bc.AddBlock([]*core.Transaction{tx1})

	// Print final balances
	fmt.Println("Final Balances:")
	printBalances(vm)
}

func printBalances(vm *core.VM) {
	for addr, acc := range vm.AccountManager.Accounts {
		fmt.Printf("Account: %s, Balance: %d\n", addr, acc.Balance)
	}
}
