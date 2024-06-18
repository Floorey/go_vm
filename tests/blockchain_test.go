package tests

import (
	"testing"
	"vm/core"
)

func TestBlockchain(t *testing.T) {
	vm := core.NewVM()
	bc := core.NewBlockchain(vm)

	tx1 := core.NewTransaction("Send 1 Coin to Alice", string([]byte{byte(core.OP_PRINT), byte(core.OP_ADD)}))
	bc.AddBlock([]*core.Transaction{tx1})

	tx2 := core.NewTransaction("Send 2 Coins to Bob", string([]byte{byte(core.OP_PRINT), byte(core.OP_SUB)}))
	bc.AddBlock([]*core.Transaction{tx2})

	if len(bc.Blocks) != 3 {
		t.Fatalf("Expected blockchain length of 3, got %d", len(bc.Blocks))
	}

	if string(bc.Blocks[1].Transactions[0].Data) != "Send 1 Coin to Alice" {
		t.Fatalf("Unexpected transaction data: %s", bc.Blocks[1].Transactions[0].Data)
	}

	if string(bc.Blocks[2].Transactions[0].Data) != "Send 2 Coins to Bob" {
		t.Fatalf("Unexpected transaction data: %s", bc.Blocks[2].Transactions[0].Data)
	}
}
