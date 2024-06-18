package main

import (
	"fmt"
	"vm/core"
)

func main() {
	vm := core.NewVM()
	bc := core.NewBlockchain(vm)

	tx1 := core.NewTransaction("Send 1 Coin to Alice", string([]byte{byte(core.OP_PRINT), byte(core.OP_ADD)}))
	bc.AddBlock([]*core.Transaction{tx1})

	tx2 := core.NewTransaction("Send 2 Coins to Bob", string([]byte{byte(core.OP_PRINT), byte(core.OP_SUB)}))
	bc.AddBlock([]*core.Transaction{tx2})

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		for _, tx := range block.Transactions {
			fmt.Printf("Transaction: %s\n", tx.Data)
			fmt.Printf("Smart Contract Code: %s\n", tx.Code)
		}
		fmt.Println()
	}
}
