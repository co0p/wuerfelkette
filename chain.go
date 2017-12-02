package wuerfelkette

import (
	"errors"
)

type Difficulty uint16

const (
	// Easy Difficulty means fasts mining
	Easy Difficulty = 1
	// Medium Difficulty means a little bit work for miners
	Medium = 2
	// Hard Difficulty means work for miners
	Hard = 4
	// Insane Difficulty means a lot of work for miners
	Insane = 8
)

// Chain contains a list of Blocks
type Chain struct {
	difficulty Difficulty
	blocks     []Block
}

// CreateChain creates a new Chain with a valid genesis block
func CreateChain(diff Difficulty) Chain {
	genesisBlock := GenesisBlock()
	blocks := []Block{}
	return Chain{difficulty: diff, blocks: append(blocks, genesisBlock)}
}

func (chain *Chain) LatestBlock() Block {
	return chain.blocks[len(chain.blocks)-1]
}

func (chain *Chain) Size() int {
	return len(chain.blocks)
}

func (chain *Chain) GetDifficulty() Difficulty {
	return chain.difficulty
}

func (chain *Chain) Append(b Block) error {

	p := chain.LatestBlock()
	if p.GetHash() != b.GetPreviousHash() {
		return errors.New("previous hash of new block is not equal latest hash")
	}

	chain.blocks = append(chain.blocks, b)

	if !chain.IsValid() {
		return errors.New("chain is not valid")
	}
	return nil
}

func (chain *Chain) IsValid() bool {

	for i := 1; i < len(chain.blocks); i++ {
		p := chain.blocks[i-1]
		b := chain.blocks[i]

		if p.GetHash() != b.GetPreviousHash() {
			return false
		}
	}
	return true
}
