package blockchain

import (
	"time"

	"github.com/fullgukbap/coin/utils"
)

const (
	minerReward int = 50
)

type mempool struct {
	Txs []*Tx
}

var Mempool *mempool = &mempool{}

// func makeTx(from, to string, amount int) (*Tx, error) {
// 	if Blockchain().BalanceByAddress(from) < amount {
// 		return nil, errors.New("not enough money")
// 	}

// }

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("fullgukbap", to, amount)
	if err != nil {
		return err
	}

	m.Txs = append(m.Txs, tx)
	return nil
}

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
	// 참고한 utxout의 transcation id
	TxID string
	// 참고한 utxout의 id
	Index int
	Owner string `json:"owner"`
}

type UTxOut struct {
	TxID string

	// input을 생성한 output의 index
	Index  int
	Amount int
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

// 채굴자를 주소로 삼는 코인베이스 거래내역을 생성한다.
func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"", -1, "COINBASE"},
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

func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("fullgukbap")
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil

	return txs
}
