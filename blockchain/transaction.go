package blockchain

import (
	"errors"
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

func makeTx(from, to string, amount int) (*Tx, error) {
	// 유효한 Tx를 생성할 수 있나?
	if Blockchain().BalanceByAddress(from) < amount {
		return nil, errors.New("not enough money")
	}

	// 유효한 Tx을 생성하자.
	var txIns []*TxIn
	var txOuts []*TxOut
	var total int

	oldTxOuts := Blockchain().TxOutsByAddress(from)

	for _, txOut := range oldTxOuts {
		if total > amount {
			break
		}

		txIn := &TxIn{
			Owner:  txOut.Owner,
			Amount: txOut.Amount,
		}
		txIns = append(txIns, txIn)
		total += txOut.Amount

	}

	// 잔돈 구현
	change := total - amount
	if change != 0 {
		changeTxOut := &TxOut{
			Owner:  from,
			Amount: change,
		}

		txOuts = append(txOuts, changeTxOut)
	}

	txOut := &TxOut{
		Owner:  to,
		Amount: amount,
	}

	txOuts = append(txOuts, txOut)

	tx := &Tx{
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}

	tx.getId()

	return tx, nil
}

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
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
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

func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("fullgukbap")
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil

	return txs
}
