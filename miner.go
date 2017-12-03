package wuerfelkette

import (
	"strconv"
	"strings"
)

// Miner can Mine a new Block
type Miner interface {
	Mine(Data, Block, Difficulty) Block
}

// SimpleMiner mines a new block with a simple strategy
type SimpleMiner struct{}

func (m *SimpleMiner) Mine(data Data, pBlock Block, diff Difficulty) Block {
	noonce := ""
	var block = NewBlock(data, pBlock, noonce)

	counter := 0
	for !m.validate(block, diff) {
		counter = counter + 1
		noonce = strconv.Itoa(counter)
		block = NewBlock(data, pBlock, noonce)
	}
	return block
}

func (m *SimpleMiner) validate(block Block, diff Difficulty) bool {
	prefix := strings.Repeat("0", int(diff))
	return strings.HasPrefix(string(block.GetHash()), prefix)
}
