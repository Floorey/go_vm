package main

import (
	"fmt"
	"vm/core"
)

func main() {
	bc := core.NewBlockchain()

	tx1 := core.NewTransaction("Send 1 BTC to Alice")
	bc.AddBlock([]*core.Transaction{tx1})

	tx2 := core.NewTransaction("Send 2 BTC to Bob")
	bc.AddBlock([]*core.Transaction{tx2})

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		for _, tx := range block.Transactions {
			fmt.Printf("Transaction: %s\n", tx.Data)
		}
		fmt.Println()
	}
}
