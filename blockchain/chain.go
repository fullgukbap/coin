/*
Blockchain 패키지는 blockchain의 내부 로직을 구현하는 패키지 입니다.
본 패키지를 사용하여 blockchain을 조작할 수 있으며, 싱글톤 패턴으로 구현되어 있습니다.
*/
package blockchain

import (
	"sync"

	"github.com/fullgukbap/coin/db"
	"github.com/fullgukbap/coin/utils"
)

// b 변수는 blockchain의 인스턴스 입니다.
var b *blockchain

const (
	// defaultDifficuly는 초기 값이 2이다.
	// 앞에 0이 두개가 되는 hash 값을 찾아야 한다는 뜻이다.
	defaultDifficulty int = 2

	// difficulty를 계산하는 주기
	// 단위는 블럭입니다. 즉, 5개의 블럭이 생성되었다면 difficulty를 계산하게 됩니다.
	difficultyInterval int = 5

	// 내가 희망하는 블럭 한 개 생성 주기, 매 2분 마다 블럭이 1개 생겼으면 좋겠다는 뜻이다.
	blockInterval int = 2

	// difficulty를  다시 계산할 때 사용되는 범위의 값
	allowedRange int = 2
)

// blockchain 구조체는 블럭체인의 구성요소 및 구조를 정의합니다.
// 데이터베이스를 이용해 영속성을 제공하기 때문에 구조가 아래와 같아도 됩니다.
type blockchain struct {
	// 가장 최신의 블럭의 Hash값
	NewestHash string `json:"newestHash"`
	// 현재 블럭의 개수
	Height int `json:"height"`

	// 현재 Difficulty의 수
	CurrentDifficulty int `json:"currentDifficulty"`
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
			b = &blockchain{
				Height: 0,
			}

			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock()
			} else {
				// restore b from bytes
				b.restore(checkpoint)
			}
		})
	}
	return b
}

// AddBlock 함수는 data 값만 전달하여 블럭체인의 블럭을 추가하는 함수 입니다.
func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

// txOuts 함수는 거래 이력 중 txOuts값만 추출하여 slice화 하여 반환합니다.
func (b *blockchain) txOuts() []*TxOut {
	var txOuts []*TxOut
	blocks := b.Blocks()
	for _, block := range blocks {
		for _, tx := range block.Transactions {
			txOuts = append(txOuts, tx.TxOuts...)
		}
	}
	return txOuts
}

// TxOutsByAddress 함수는 거래 이력 중 address인 것만 추출하여 반환합니다.
func (b *blockchain) TxOutsByAddress(address string) (ownedTxOuts []*TxOut) {
	txOuts := b.txOuts()
	for _, txOut := range txOuts {
		if txOut.Owner == address {
			ownedTxOuts = append(ownedTxOuts, txOut)
		}
	}

	return ownedTxOuts
}

// BalanceByAddress 함수는 address의 tx의 값을 총합하여 반환합니다.
func (b *blockchain) BalanceByAddress(address string) (amount int) {
	txOuts := b.TxOutsByAddress(address)
	for _, txOut := range txOuts {
		amount += txOut.Amount
	}
	return amount
}

// Blocks 함수는 blockchain의 모든 블럭을 반환합니다.
func (b *blockchain) Blocks() (blocks []*Block) {
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}

	return blocks
}

// recalculateDifficulty 함수는 difficulty를 다시 계산을 하여 반환해주는 함수 입니다.
// 전제: 이 함수는 difficulty Interval(difficulty를 생성하는 주기)에 따라 호출됩니다.
func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	// 가장 최신의 블럭
	newestBlock := allBlocks[0]

	// 지난 difficulty가 재설정된 block
	lastRecalculateBlock := allBlocks[difficultyInterval-1]

	// 실제로 두 블럭 사이의 걸린 시간
	// 단위가 초 이므로 /60으로 나누어서 분 단위로 변경
	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculateBlock.Timestamp / 60)

	// 이 두 블럭 사이의 이상적인 걸린 시간 == 예상 시간
	expectedTime := difficultyInterval * blockInterval

	// expectedTime 기준으로만 Difficulty를 계산하면 너무 빡빡해, 범위를 이용해 측정한다.
	if actualTime < (expectedTime - allowedRange) {
		// 빨리 생성되고 있기 때문에, Difficulty를 올림!
		return b.CurrentDifficulty + 1
	} else if actualTime > (expectedTime + allowedRange) {
		// 너무 느리게 생성되고 있기 때문에, Difficulty를 내림!
		return b.CurrentDifficulty - 1
	}

	return b.CurrentDifficulty
}

// difficulty 함수는 본 블럭체인의 difficulty의 값을 도출해냅니다.
func (b *blockchain) Difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	}
	if b.Height%difficultyInterval == 0 {
		// 다시 difficluty 계산
		return b.recalculateDifficulty()
	}

	return b.CurrentDifficulty
}

// restore 함수는 data의 값을 i로 복구시킵니다.
func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

// persist 함수는 데이터베이스에 값을 저장하여 영속성을 부여합니다.
func (b *blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(b))
}
