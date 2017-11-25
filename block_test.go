package wuerfelkette_test

import (
	"testing"

	"github.com/co0p/wuerfelkette"
)

func TestGenesisBlock_should_create_always_same_block(t *testing.T) {

	genesis1 := wuerfelkette.GenesisBlock()
	genesis2 := wuerfelkette.GenesisBlock()

	if genesis1.GetHash() != genesis2.GetHash() {
		t.Errorf("hashes should not differ: %s != %s", genesis1.GetHash(), genesis2.GetHash())
	}
}
