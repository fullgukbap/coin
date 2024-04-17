package blockchain

// Tx 구조체는 거래를 하기 위해 사용되는 정보로 구성됩니다.
type Tx struct {
	// Id 변수는 각각의 거래를 식별하기 위해 사용됩니다.
	Id string

	// Timestamp는 거래가 발생한 시간을 기록하기 위해 사용됩니다.
	Timestamp int
	TxIns     []*TxIn
	TxOuts    []*TxOuts
}

// TxIn 구조체는 거래의 Input의 구성요소가 존재합니다.
type TxIn struct {
	// Owner는 소유주를 나타냅니다.
	Owner  string
	Amount int
}

// TxOuts 구조체는 거래의 Output의 구성요소가 존재합니다.
type TxOuts struct {
	// Owner는 소유주를 나타냅니다.
	Owner  string
	Amount int
}
