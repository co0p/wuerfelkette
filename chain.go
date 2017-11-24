package wuerfelkette

// Chain contains a list of Blocks
type Chain []Block

// CreateChain creates a new Chain with a valid genesis block
func CreateChain(data Data) Chain {
	var newChain Chain
	genesisBlock := GenesisBlock(data)
	newChain = append(newChain, genesisBlock)
	return newChain
}
