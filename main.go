package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
}

func calculateHash(block Block) string {
	concBlock := string(block.Index) + block.Timestamp + block.Data + block.PrevHash
	hash := sha256.New()
	hash.Write([]byte(concBlock))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

func createNewBlock(oldBlock Block, data string) (Block, error) {
	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	newBlock.PrevHash = oldBlock.Hash

	return newBlock, nil
}

func validateBlock(block Block) bool {

	if block.Index == 0 { // Genesis Block has no prevBlock
		return calculateHash(block) == block.Hash
	} else {
		// TODO: Assign block to prevBlock
		var prevBlock Block

		return calculateHash(block) == block.Hash && prevBlock.Hash == block.PrevHash
	}
}
