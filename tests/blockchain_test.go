package tests

import (
	"testing"
	"vm/core"
)

func TestBlockchain(t *testing.T) {
	bc := core.NewBlockchain()

	tx1 := core.NewTransaction("Send 1 Coin to Alice")
	bc.AddBlock([]*core.Transaction{tx1})

	tx2 := core.NewTransaction("Send 2 Coins to Me")
	bc.AddBlock([]*core.Transaction{tx2})

	if len(bc.Blocks) != 3 {
		t.Fatalf("Expectred blockchain lenght of 3, got %d", len(bc.Blocks))
	}
	if string(bc.Blocks[1].Transactions[0].Data) != "Send 1 COin to Alice" {
		t.Fatalf("Unexpected transaction data: %s", bc.Blocks[1].Transactions[0].Data)
	}
	if string(bc.Blocks[2].Transactions[0].Data) != "Send 2 Coins to Me" {
		t.Fatalf("Unexpected transaction data: %s", bc.Blocks[2].Transactions[0].Data)
	}
}
