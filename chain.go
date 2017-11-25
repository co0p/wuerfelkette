package wuerfelkette

// Chain contains a list of Blocks
type Chain struct {
	blocks []Block
}

// CreateChain creates a new Chain with a valid genesis block
func CreateChain() Chain {
	var newChain Chain
	genesisBlock := GenesisBlock()
	newChain.blocks = append(newChain.blocks, genesisBlock)
	return newChain
}

func (chain *Chain) Latest() Block {
	return chain.blocks[len(chain.blocks)-1]
}

func (chain *Chain) Size() int {
	return len(chain.blocks)
}

func (chain *Chain) Append(data Data) {
	previousBlock := chain.Latest()
	newBlock := NewBlock(data, previousBlock)
	chain.blocks = append(chain.blocks, newBlock)
}

func (chain *Chain) IsValid() bool {

	// TODO
	return true
}
