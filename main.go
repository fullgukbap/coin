package main

import (
	"fmt"

	"github.com/JJerBum/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlock()
	chain.AddBlock("Two block")

	for _, block := range chain.AllBlocks() {
		fmt.Println(block)
	}
}
