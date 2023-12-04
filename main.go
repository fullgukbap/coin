package main

import (
	"github.com/JJerBum/nomadcoin/explorer"
	"github.com/JJerBum/nomadcoin/rest"
)

func main() {
	go rest.Start(2000)
	explorer.Start(4000)
}
