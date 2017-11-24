package wuerfelkette

import (
	"crypto/sha256"
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
func GenesisBlock(data Data) Block {
	newBlock := Block{
		timestamp: time.Now(),
		data:      data,
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
	h := sha256.New()
	str := fmt.Sprintf("%v", block)
	h.Write([]byte(str))
}
