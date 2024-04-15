/*
Blockchain 패키지는 blockchain의 내부 로직을 구현하는 패키지 입니다.
본 패키지를 사용하여 blockchain을 조작할 수 있으며, 싱글톤 패턴으로 구현되어 있습니다.
*/
package blockchain

import (
	"fmt"
	"sync"

	"github.com/fullgukbap/coin/db"
	"github.com/fullgukbap/coin/utils"
)

// b 변수는 blockchain의 인스턴스 입니다.
var b *blockchain

// blockchain 구조체는 블럭체인의 구성요소 및 구조를 정의합니다.
// 데이터베이스를 이용해 영속성을 제공하기 때문에 구조가 아래와 같아도 됩니다.
type blockchain struct {
	// 가장 최신의 블럭의 Hash값
	NewestHash string `json:"newestHash"`
	// 현재 블럭의 개수
	Height int `json:"height"`
}

// once 변수는 Blockchain 함수의 once.Do 함수를 호출하기 위해 사용됩니다.
// 정교한 싱글톤 패턴 구현을 위해 사용됩니다.
var once sync.Once

// Blockchain 함수는 Blockchain 호출의 진입점 입니다.
// 만약 처음 호출 시 blockchain인스턴스가 초기화 및 복구되며
// 두 번 이상 호출 시 인스턴스를 반환합니다.
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}

			//
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis block")
			} else {
				// restore b from bytes
				b.restore(checkpoint)
			}

		})
	}
	return b
}

// AddBlock 함수는 data 값만 전달하여 블럭체인의 블럭을 추가하는 함수 입니다.
func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	fmt.Println(block.Hash)
	b.persist()
}

// restore 함수는 data의 값을 i로 복구시킵니다.
func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

// persist 함수는 데이터베이스에 값을 저장하여 영속성을 부여합니다.
func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}
