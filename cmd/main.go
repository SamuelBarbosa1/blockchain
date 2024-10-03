package main

import (
	"blockchain-project/internal/blockchain"
	"blockchain-project/internal/transaction"
	"fmt"
)

func main() {

	bc := blockchain.NewBlockchain()

	t1 := transaction.NewTransaction("Alice", "Bob", 10.0)
	t2 := transaction.NewTransaction("Bob", "Charlie", 5.0)

	bc.AddBlock([]transaction.Transaction{t1, t2})

	for _, block := range bc.Blocks {
		fmt.Printf("Bloco %d, Hash: %s, Previous Hash: %s\n", block.Index, block.Hash, block.PreviousHash)
	}
}
