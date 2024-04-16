package main

import (
	"github.com/fullgukbap/coin/cli"
	"github.com/fullgukbap/coin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
