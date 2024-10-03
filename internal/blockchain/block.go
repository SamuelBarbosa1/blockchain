package blockchain

import (
	"blockchain-project/internal/transaction"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Transactions []transaction.Transaction
	PreviousHash string
	Hash         string
}

func CalculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func CreateBlock(previousBlock Block, transactions []transaction.Transaction) Block {
	newBlock := Block{
		Index:        previousBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PreviousHash: previousBlock.Hash,
		Hash:         "",
	}
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock
}
