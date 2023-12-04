// blockchain 패키지는 blockchain을 구현하는 구조체 입니다.
package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

// Block 구조체는 blockchain의 block을 정의하는 구조체 입니다.
type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

// calucateHash 함수는 receiver pointer 값으로 넘어온 b(block)의 hash값을 대신 연산해주는 코드 입니다.
func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

// blockchain 구조체는 blockchain을 정의하는 구조체 입니다.
// 각각의 block은 block instance의 주소를 저장합니다.
type blockchain struct {
	blocks []*Block
}

// AddBlock함수는 pointer receiver인 b의 blocks 필드를 data 값만 인자로 받고 block을 추가해주는 함수 입니다.
func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

// AllBlocks 함수는 blockchain의 recevier포인터로, b의 blocks를 반환하는 함수 입니다.
func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

// b 인스턴스는 singleton pattern으로, 해당 인스턴스를 application에 공유하는 것이 아니라 어떠한 함수를 통해 공유합니다.
// 이렇게 하면 instance를 b instance의 초기화를 제어할 수 있습니다.
var b *blockchain

// once는 Do method를 사용하기 위해 선언한 변수 입니다.
var once sync.Once

// getLastHash 함수는 이전의 blockchain의 이전의 해쉬 값을 반환하는 함수 입니다.
// 만약 첫번쨰 블럭이 이 함수를 호출했을 경우 ""를 반환합니다.
func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

// createBlock 함수는 data를 인자로 받아 블럭을 생성하고,
// getLastHash()함수를 사용하여 전 블럭의 해쉬를 가져오고
// calculateHash() 함수를 사용해 hash값을 연산합니다.
func createBlock(data string) *Block {
	block := Block{data, "", getLastHash(), len(GetBlockchain().AllBlocks()) + 1}
	block.calculateHash()

	return &block
}

func (b *blockchain) GetBlock(height int) *Block {
	return b.blocks[height-1]
}

// GetBlock() 함수는 b instance를 얻을 수 있는 함수 입니다.
// Singleton pattern을 사용함으로써 b instace의 초기화는 이 함수에서 다 이루어 집니다.
func GetBlockchain() *blockchain {
	if b == nil {
		// Only once
		once.Do(func() {
			b = new(blockchain)
			b.AddBlock("Genesiss block")
		})
	}
	return b
}
