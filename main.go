package main

import (
	"crypto/sha256"
	"encoding/hex"
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
