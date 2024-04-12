package utils

import (
	"bytes"
	"encoding/gob"
)

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ToBytes(i any) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}
