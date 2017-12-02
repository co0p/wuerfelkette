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
	noonce       string
}

// NewBlock constructs a new Block with a valid hash
func NewBlock(data Data, previousBlock Block, noonce string) Block {
	newBlock := Block{
		timestamp:    time.Now(),
		data:         data,
		previousHash: previousBlock.GetHash(),
		noonce:       noonce,
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

// GetPreviousHash returns the hash of the previous block
func (block *Block) GetPreviousHash() Hash {
	return block.previousHash
}

// hash() hashes the
func (block *Block) hash() {
	hasher := sha256.New()
	str := fmt.Sprintf("%s - %s - %s - %s - %s", block.data, block.timestamp, block.ownHash, block.previousHash, block.noonce)
	hasher.Write([]byte(str))
	demo := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	block.ownHash = Hash(demo)
}

func (block *Block) String() string {
	return "block: " + string(block.data)
}
