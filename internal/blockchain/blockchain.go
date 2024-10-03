package blockchain

import "blockchain-project/internal/transaction"

type Blockchain struct {
	Blocks []Block
}

func (bc *Blockchain) AddBlock(transactions []transaction.Transaction) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := CreateBlock(previousBlock, transactions)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	genesisBlock := Block{
		Index:        0,
		Timestamp:    "2023-01-01T00:00:00Z",
		Transactions: []transaction.Transaction{},
		PreviousHash: "0",
		Hash:         "",
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)
	return &Blockchain{
		Blocks: []Block{genesisBlock},
	}
}
