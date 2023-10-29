// blockchain 패키지는 blockchain을 구현하는 구조체 입니다.
package blockchain

// block 구조체는 blockchain의 block을 정의하는 구조체 입니다.
type block struct {
	data     string
	hash     string
	prevHash string
}

// blockchain 구조체는 blockchain을 정의하는 구조체 입니다.
type blockchain struct {
	blocks []block
}

// b 인스턴스는 singleton pattern으로, 해당 인스턴스를 application에 공유하는 것이 아니라 어떠한 함수를 통해 공유합니다.
// 이렇게 하면 instance를 b instance의 초기화를 제어할 수 있습니다.
var b *blockchain

// GetBlock() 함수는 b instance를 얻을 수 있는 함수 입니다.
// Singleton pattern을 사용함으로써 b instace의 초기화는 이 함수에서 다 이루어 집니다.
func GetBlock() *blockchain {
	if b == nil {
		b = new(blockchain)
	}
	return b
}
