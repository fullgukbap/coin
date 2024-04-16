package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	// blockchain.Blockchain()
	// defer db.Close()

	// cli.Start()

	difficulty := 8
	target := strings.Repeat("0", difficulty)
	nonce := 0
	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("1"+fmt.Sprint(nonce))))

		if strings.HasPrefix(hash, target) {
			fmt.Printf("Chache!\ntarget: %s\nnonce: %d\nhash: %s\n\n", target, nonce, hash)
			return
		} else {
			fmt.Printf("mining!\ntarget: %s\nnonce: %d\nhash: %s\n\n", target, nonce, hash)
			nonce++
		}

	}

}
