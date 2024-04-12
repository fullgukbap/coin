package main

import "github.com/JJerBum/nomadcoin/blockchain"

func main() {
	// cli.Start()
	blockchain.Blockchain().AddBlock("Good")
	blockchain.Blockchain().AddBlock("xx")
	blockchain.Blockchain().AddBlock("zz")
}
