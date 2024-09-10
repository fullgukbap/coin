package blockchain

import (
	"errors"
	"strings"
	"time"

	"github.com/fullgukbap/coin/db"
	"github.com/fullgukbap/coin/utils"
)

// ErrNotFound 변수는 블럭을 찾이 못했을 때 사용됩니다.
var ErrNotFound = errors.New("block not found")

// Block 구조체는 블럭체인의 한 개의 노드에 해당되는 구조를 정의했습니다.
type Block struct {
	Transactions []*Tx `json:"transactions"`

	// 현재 블럭의 Hash값을 저장하는 변수 입니다.
	// 현재 블럭의 Hash값은 일방향 함수(Data + PrevHash + Height)으로 구성됩니다.
	Hash string `json:"hash"`

	// 전 블럭의 Hash를 저장하는 변수 입니다.
	PrevHash string `json:"prevHash,omitempty"`

	// 현재 블럭의 index를 저장하는 변수 입니다.
	Height int `json:"height"`

	// 시작 부분의 0의 개수를 지정하는 변수 입니다.
	// 네트워가 보유한 힘의 양에 따라서 유동적으로 값이 변하게 된다.
	Difficulty int `json:"difficulty"`

	// Number used once의 약자로 채굴자들이 사용하는 변수 값
	Nonce int `json:"nonce"`

	// Timestamp 변수는 블럭의 생성일을 저장합니다.
	Timestamp int `json:"timestamp"`
}

// createBlock 함수는 data, prevHash, height 값을 받은 후 hash 값을 계산한 후 db에 저장합니다.
func createBlock(prevHash string, height int) *Block {
	block := &Block{
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: Blockchain().Difficulty(),
		Nonce:      0,
		// 이렇게 하게 되면 체굴 과정의 시간은 생략되기 때문에 정확하지 않은 생성일이 들어감, 고로 초기화에 시간을 넣지 않고, mine에 넣을 거임
		// Timestamp: int(time.Now().Unix()),
		// Transactions: []*Tx{makeCoinbaseTx("fullgukbap")},
	}
	block.mine()
	block.Transactions = Mempool.TxToConfirm()
	block.persist()
	return block
}

// mine 함수는 말 그래도 채굴을 수행합니다.
// 채굴을 수행하여 결과적으로 생긴 hash값을 대입해줍니다.
func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		// fmt.Printf("[MINING..]\nhash: %s\ntarget:%s\nNonce:%d\n\n", hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			// fmt.Printf("[CACHED!]\nhash: %s\ntarget:%s\nNonce:%d\n\n", hash, target, b.Nonce)
			break
		} else {
			b.Nonce++
		}
	}
}

// FindBlock 함수는 hash 값으로 데이터베이스에 있는 블럭을 탐색한 뒤 최종적으로 *Block, error를 반환합니다.
func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
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
