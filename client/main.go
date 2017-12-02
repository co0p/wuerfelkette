package main

import (
	"fmt"

	"github.com/co0p/wuerfelkette"
)

func main() {

	chain := wuerfelkette.CreateChain(2)
	miner := wuerfelkette.SimpleMiner{}

	data := wuerfelkette.Data("hi there")
	previousBlock := chain.LatestBlock()
	difficulty := chain.GetDifficulty()

	newBlock := miner.Mine(data, previousBlock, difficulty)
	if err := chain.Append(newBlock); err != nil {
		fmt.Println(err)
	}
}
