package blockchain

import (
	"time"

	"github.com/fullgukbap/coin/utils"
)

const (
	minerReward int = 50
)

type Tx struct {
	Id        string   `json:"id"` // Id = hash(timestamp + txIns + TxOuts)
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

type TxIn struct {
	Owner  string
	Amount int
}

type TxOut struct {
	Owner  string
	Amount int
}

// 채굴자를 주소로 삼는 코인베이스 거래내역을 생성한다.
func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerReward},
	}

	txOuts := []*TxOut{
		{address, minerReward},
	}

	tx := &Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}

	tx.getId()

	return tx
}
