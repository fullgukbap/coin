package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/fullgukbap/coin/db"
	"github.com/fullgukbap/coin/utils"
)

// ErrNotFound 변수는 블럭을 찾이 못했을 때 사용됩니다.
var ErrNotFound = errors.New("block not found")

// Block 구조체는 블럭체인의 한 개의 노드에 해당되는 구조를 정의했습니다.
type Block struct {
	// 값을 저장하는 변수 입니다.
	Data string `json:"data"`

	// 현재 블럭의 Hash값을 저장하는 변수 입니다.
	// 현재 블럭의 Hash값은 일방향 함수(Data + PrevHash + Height)으로 구성됩니다.
	Hash string `json:"hash"`

	// 전 블럭의 Hash를 저장하는 변수 입니다.
	PrevHash string `json:"prevHash,omitempty"`

	// 현재 블럭의 index를 저장하는 변수 입니다.
	Height int `json:"height"`
}

// createBlock 함수는 data, prevHash, height 값을 받은 후 hash 값을 계산한 후 db에 저장합니다.
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

// FindBlock 함수는 hash 값으로 데이터베이스에 있는 블럭을 탐색한 뒤 최종적으로 *Block, error를 반환합니다.
func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes != nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

// persist 함수는 현재 b의 hash 값을 기준으로 []byte로 변환하여 데이터베이스에 저장합니다.
func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

// restore 함수는 현재 data의 값을 b 구조체를 복구합니다.
func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}
