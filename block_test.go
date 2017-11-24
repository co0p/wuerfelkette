package wuerfelkette_test

import (
	"testing"

	"github.com/co0p/wuerfelkette"
)

func TestGenesisBlock_should_create_same_hash_for_same_data(t *testing.T) {

	testcases := []struct {
		data1  wuerfelkette.Data
		data2  wuerfelkette.Data
		equals bool
	}{
		{"some data", "some data", true},
	}

	for _, tc := range testcases {
		genesis1 := wuerfelkette.GenesisBlock(tc.data1)
		genesis2 := wuerfelkette.GenesisBlock(tc.data2)

		var isEqual bool = genesis1.GetHash() == genesis2.GetHash()
		if isEqual != tc.equals {
			t.Errorf("got %v; want %v", isEqual, tc.equals)
		}
	}
}
