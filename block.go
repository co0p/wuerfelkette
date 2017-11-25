package wuerfelkette

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

// Data represents the data the block chain contains
type Data string

// Hash is a hash for a block
type Hash string

// Block is part of a chain ... a block chain :-)
type Block struct {
	timestamp    time.Time
	data         Data
	previousHash Hash
	ownHash      Hash
}

// NewBlock constructs a new Block with a valid hash
func NewBlock(data Data, previousBlock Block) Block {
	newBlock := Block{
		timestamp:    time.Now(),
		data:         data,
		previousHash: previousBlock.GetHash(),
	}
	newBlock.hash()

	return newBlock
}

// GenesisBlock constructs a new Block with a valid hash and no previous block
func GenesisBlock() Block {
	newBlock := Block{
		timestamp: time.Date(1, 1, 1, 1, 1, 1, 1, time.Local),
		data:      Data("genesis"),
	}
	newBlock.hash()

	return newBlock
}

// GetHash returns the hash for the block
func (block *Block) GetHash() Hash {
	return block.ownHash
}

// hash() hashes the
func (block *Block) hash() {
	hasher := sha256.New()
	str := fmt.Sprintf("%v", block)
	hasher.Write([]byte(str))
	demo := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	block.ownHash = Hash(demo)
}
