/*
db패키지는 블럭체인 값의 영속성을 부여하기 위해 존재합니다.
db는 싱글톤패턴으로 구현되어 있습니다.
*/
package db

import (
	"github.com/boltdb/bolt"
	"github.com/fullgukbap/coin/utils"
)

const (
	dbName      = "blockchain.db"
	dataBucket  = "data"
	blockBucket = "blocks"

	checkpoint = "checkpoint"
)

// db 변수는 데이터베이스의 인스턴스 입니다.
var db *bolt.DB

// DB 함수는 데이터베이스를 호출하는 진입점 함수 입니다.
// 만약 처음 호출 시 데이터베이스 초기화 및 복구를 진행합니다.
// 만약 처음 호출 경험이 있는데, 호출했다면 인스턴스만 반환하게 됩니다.
func DB() *bolt.DB {
	if db == nil {
		// rwx , rwx, rwx
		// 110, 000, 000 -> only read and write
		dbPointer, err := bolt.Open("blockchain.db", 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blockBucket))
			utils.HandleErr(err)
			return err
		})
		utils.HandleErr(err)
	}
	return db
}

// SaveBlockchain 함수는 data 값을 데이터베이스 저장합니다.
func SaveBlockchain(data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)
}

// SaveBlock 함수는 말 그대로 Block을 저장하는 함수 입니다.
func SaveBlock(hash string, data []byte) {
	// fmt.Printf("Saving Block %s\nData: %b\n", hash, data)
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blockBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

// Checkpoint 함수는 blockchain의 구조체 값을 read 하여 []byte로 반환합니다.
func Checkpoint() []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	return data
}

// Block 함수는 hash값으로 block을 찾아 []byte로 반환합니다.
func Block(hash string) []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blockBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}
