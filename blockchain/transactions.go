package blockchain

import (
	"errors"
	"time"

	"github.com/fullgukbap/coin/utils"
)

const (
	// minerReward 변수는 광부에게 주는 포상금 입니다.
	minerReward int = 50
)

// mempool 구조체는 trasaction(거래)들을 저장하기 위해 사용됩니다.
type mempool struct {
	Txs []*Tx
}

// Mempool 변수는 mempool type의 인스턴스 이며, 효용은 구조체 주석과 같습니다.
// export인 이유는 초기화 단계를 제어하는 것이 아님(-> 싱글톤 패턴을 사용할 필요가 없음)
// 메모리에 있는 값을 보여주기만 해줘도 되기 때문이다.
var Mempool *mempool = &mempool{}

// makeTx 함수는 거래를 생성합니다.
func makeTx(from /*보내는 이*/, to /*받는 이*/ string, amount /*보낼 돈*/ int) (*Tx, error) {
	// 돈이 충분한가?
	if Blockchain().BalanceByAddress(from) < amount {
		return nil, errors.New("not enough money")
	}

	var txIns []*TxIn
	var txOuts []*TxOut
	total := 0

	oldTxOuts := Blockchain().TxOutsByAddress(from)
	for _, txOut := range oldTxOuts {
		if total > amount {
			break
		}
		total += txOut.Amount
		txIns = append(txIns, &TxIn{txOut.Owner, txOut.Amount})
	}

	change := total - amount
	// 값이 0이면 txOuts(from)에 표현하지 않음
	if change != 0 {
		txOuts = append(txOuts, &TxOut{from, change})
	}

	txOuts = append(txOuts, &TxOut{to, amount})
	tx := &Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return tx, nil
}

// addTx 함수는 tx를 하나 만든 뒤, m.Txs에 tx을 추가하는 함수 입니다.
func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("nico", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

// Tx 구조체는 거래를 하기 위해 사용되는 정보로 구성됩니다.
type Tx struct {
	// Id 변수는 각각의 거래를 식별하기 위해 사용됩니다.
	Id string `json:"id"`

	// Timestamp는 거래가 발생한 시간을 기록하기 위해 사용됩니다.
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

// getId 함수는 Id 값을 자동적으로 생성하여 대입해줍니다.
func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

// TxIn 구조체는 거래의 Input의 구성요소가 존재합니다.
type TxIn struct {
	// Owner는 소유주를 나타냅니다.
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

// TxOuts 구조체는 거래의 Output의 구성요소가 존재합니다.
type TxOut struct {
	// Owner는 소유주를 나타냅니다.
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

// makeCoinbaseTx 함수는 채굴자를 주소로 삼는 코인베이스 거래내역을 생성해 *Tx 포인터를 반환합니다.
func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerReward},
	}

	txOuts := []*TxOut{
		{address, minerReward},
	}

	tx := Tx{
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}
