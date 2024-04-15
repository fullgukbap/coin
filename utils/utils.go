/*
Package utils는 본 프로젝트를 구현하는데 있어서 두루두루 사용되는 함수들을 구현한 패키지 입니다.
*/
package utils

import (
	"bytes"
	"encoding/gob"
)

// HandleErr 함수는 err의 값이 nil이 아니면 panic을 발생하는 함수 입니다.
func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// ToBytes 함수는 i의 값을 []byte로 변환해주는 함수 입니다.
// 중간에 에러가 발생하면 panic을 터트립니다.
func ToBytes(i any) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

// FromBytes 함수는 i, data의 값을 이용해 data의 값을 i 값으로 Decode 합니다.
// 중간에 에러가 발생하면 panic을 터트립니다.
func FromBytes(i any, data []byte) {
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(encoder.Decode(i))
}
