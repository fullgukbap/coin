package main

import (
	"github.com/fullgukbap/coin/blockchain"
	"github.com/fullgukbap/coin/cli"
	"github.com/fullgukbap/coin/db"
)

func main() {
	blockchain.Blockchain()
	defer db.Close()

	cli.Start()
}
