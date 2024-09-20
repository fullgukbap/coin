package main

import (
	"coin/cli"
	"coin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
