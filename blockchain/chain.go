// blockchain 패키지는 blockchain을 구현하는 구조체 입니다.
package blockchain

import (
	"sync"
)

// blockchain 구조체는 blockchain을 정의하는 구조체 입니다.
// 데이터베이스 값을 저장함으로 NewestHash(가장 최신의 hash)와, 몇개의 블럭이 존재하는지만 저장하면 됩니다.
type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

// b 인스턴스는 singleton pattern으로, 해당 인스턴스를 application에 공유하는 것이 아니라 어떠한 함수를 통해 공유합니다.
// 이렇게 하면 instance를 b instance의 초기화를 제어할 수 있습니다.
var b *blockchain

// once는 Do method를 사용하기 위해 선언한 변수 입니다.
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height

}

// GetBlock() 함수는 b instance를 얻을 수 있는 함수 입니다.
// Singleton pattern을 사용함으로써 b instace의 초기화는 이 함수에서 다 이루어 집니다.
func Blockchain() *blockchain {
	if b == nil {
		// Only once
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Genesis block")
		})
	}
	return b
}
