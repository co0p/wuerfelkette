package wuerfelkette_test

import (
	"testing"

	"github.com/co0p/wuerfelkette"
)

func TestCreateChain_should_create_valid_chain(t *testing.T) {
	chain := wuerfelkette.CreateChain()
	if !chain.IsValid() {
		t.Errorf("Initial chain should be valid, got %v", chain.IsValid())
	}
}

func TestCreateChain_should_create_chain_with_genesis_block(t *testing.T) {
	chain := wuerfelkette.CreateChain()
	genesisBlock := wuerfelkette.GenesisBlock()
	block := chain.Latest()

	if chain.Size() != 1 {
		t.Errorf("initial size of chain should be 1, got %v", chain.Size())
	}

	if genesisBlock.GetHash() != block.GetHash() {
		t.Errorf("first block should be equal to genesis block, got %v expected %v", block.GetHash(), genesisBlock.GetHash())
	}
}
