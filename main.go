package main

import (
	"github.com/fullgukbap/coin/blockchain"
	"github.com/fullgukbap/coin/cli"
)

func main() {
	blockchain.Blockchain().AddBlock("first")
	blockchain.Blockchain().AddBlock("third")
	blockchain.Blockchain().AddBlock("second")

	cli.Start()
}
