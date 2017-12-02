package wuerfelkette_test

import (
	"testing"

	"github.com/co0p/wuerfelkette"
)

func TestCreateChain_should_create_valid_chain(t *testing.T) {
	chain := wuerfelkette.CreateChain(wuerfelkette.Easy)
	if !chain.IsValid() {
		t.Errorf("Initial chain should be valid, got %v", chain.IsValid())
	}
}

func TestCreateChain_should_create_chain_with_genesis_block(t *testing.T) {
	chain := wuerfelkette.CreateChain(wuerfelkette.Easy)
	genesisBlock := wuerfelkette.GenesisBlock()
	block := chain.LatestBlock()

	if chain.Size() != 1 {
		t.Errorf("initial size of chain should be 1, got %v", chain.Size())
	}

	if genesisBlock.GetHash() != block.GetHash() {
		t.Errorf("first block should be equal to genesis block, got %v expected %v", block.GetHash(), genesisBlock.GetHash())
	}
}

func TestCreateChain_allows_to_add_new_blocks(t *testing.T) {
	chain := wuerfelkette.CreateChain(wuerfelkette.Easy)
	miner := wuerfelkette.SimpleMiner{}

	block1 := miner.Mine("some data1", chain.LatestBlock(), chain.GetDifficulty())
	err := chain.Append(block1)

	if err != nil {
		t.Errorf("chain.Append should not throw error, got: %v with block: %v", err, block1)
	}

	block2 := miner.Mine("some data2", chain.LatestBlock(), chain.GetDifficulty())
	err = chain.Append(block2)

	if err != nil {
		t.Errorf("chain.Append should not throw error, got: %v with block: %v", err, block2)
	}

	if chain.Size() != 3 {
		t.Errorf("chain.Size should be 3, got: %v ", chain.Size())
	}
}
