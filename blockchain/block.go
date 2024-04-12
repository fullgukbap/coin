package blockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/JJerBum/nomadcoin/db"
	"github.com/JJerBum/nomadcoin/utils"
)

// Block 구조체는 blockchain의 block을 정의하는 구조체 입니다.
type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

// createBlock 함수는 block을 생성해주는 함수 입니다.
func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return block
}
