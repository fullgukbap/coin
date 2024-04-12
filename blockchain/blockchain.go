// blockchain 패키지는 blockchain을 구현하는 구조체 입니다.
package blockchain

import (
	"sync"
)

// Block 구조체는 blockchain의 block을 정의하는 구조체 입니다.
type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

// blockchain 구조체는 blockchain을 정의하는 구조체 입니다.
// 각각의 block은 block instance의 주소를 저장합니다.
type blockchain struct {
	blocks []*Block
}

// b 인스턴스는 singleton pattern으로, 해당 인스턴스를 application에 공유하는 것이 아니라 어떠한 함수를 통해 공유합니다.
// 이렇게 하면 instance를 b instance의 초기화를 제어할 수 있습니다.
var b *blockchain

// once는 Do method를 사용하기 위해 선언한 변수 입니다.
var once sync.Once

// GetBlock() 함수는 b instance를 얻을 수 있는 함수 입니다.
// Singleton pattern을 사용함으로써 b instace의 초기화는 이 함수에서 다 이루어 집니다.
func Blockchain() *blockchain {
	if b == nil {
		// Only once
		once.Do(func() {
			b = new(blockchain)
			b.AddBlock("Genesiss block")
		})
	}
	return b
}
