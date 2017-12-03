package wuerfelkette_test

import (
	"strings"
	"testing"

	"github.com/co0p/wuerfelkette"
)

func TestMine_should_mine_a_new_block(t *testing.T) {
	chain := wuerfelkette.CreateChain(wuerfelkette.Easy)
	latestBlock := chain.LatestBlock()
	diffculty := chain.GetDifficulty()

	data := wuerfelkette.Data("some data")
	miner := wuerfelkette.SimpleMiner{}
	newBlock := miner.Mine(data, latestBlock, diffculty)

	if newBlock.GetPreviousHash() != latestBlock.GetHash() {
		t.Errorf("new block should have previousHash of previous block, got %v", newBlock.GetPreviousHash())
	}
	prefix := strings.Repeat("0", int(diffculty))
	if !strings.HasPrefix(string(newBlock.GetHash()), prefix) {
		t.Errorf("hash of new block should start with prefix %v, got %v", prefix, newBlock.GetHash())
	}
}
