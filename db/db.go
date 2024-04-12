// blockchain과 상호작용을 많이 할 패키지이다.
package db

import (
	"fmt"

	"github.com/JJerBum/nomadcoin/utils"
	"github.com/boltdb/bolt"
)

const (
	dbName      = "blockchain.db"
	dataBucket  = "data"
	blockBucket = "blocks"
)

var db *bolt.DB

// DB() 함수는 db 인스턴스를 load 및 초기화하는 함수 입니다.
// singleton pattern 기반입니다.
func DB() *bolt.DB {
	if db == nil {
		// initialize db
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

// SaveBlock 함수는 말 그대로 Block을 저장하는 함수 입니다.
func SaveBlock(hash string, data []byte) {
	fmt.Printf("Saving Block %s\nData: %b\n", hash, data)
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blockBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte("checkpoint"), data)
		return err
	})
	utils.HandleErr(err)
}
