package tests

import (
	"testing"
	"vm/core"
)

func TestBlockchain(t *testing.T) {
	vm := core.NewVM()
	validatorManager := core.NewValidatorManager()
	poh := core.NewProofOfHistory()

	validatorManager.AddValidator("Validator1", 100)
	validatorManager.AddValidator("Validator2", 50)

	bc := core.NewBlockchain(vm, validatorManager, poh)

	vm.AccountManager.CreateAccount("Alice", 1000)
	vm.AccountManager.CreateAccount("Bob", 500)

	tx1 := core.NewTransaction("Transfer 100 Coins from Alice to Bob", string([]byte{byte(core.OP_TRANSFER)}))
	copy(vm.Memory[0:], []byte("Alice"))
	copy(vm.Memory[32:], []byte("Bob"))
	vm.Registers["amount"] = 100
	bc.AddBlock([]*core.Transaction{tx1})

	balanceAlice, _ := vm.AccountManager.GetBalance("Alice")
	balanceBob, _ := vm.AccountManager.GetBalance("Bob")

	if balanceAlice != 900 {
		t.Fatalf("Expected balance of Alice: 900, got: %d", balanceAlice)
	}

	if balanceBob != 600 {
		t.Fatalf("Expected balance of Bob: 600, got: %d", balanceBob)
	}
}
