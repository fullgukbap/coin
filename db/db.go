// blockchain과 상호작용을 많이 할 패키지이다.
package db

import (
	"github.com/JJerBum/nomadcoin/utils"
	"github.com/boltdb/bolt"
)

const (
	dbName      = "blockchain.db"
	dataBucket  = "data"
	blockBucket = "blocks"
)

var db *bolt.DB

// singleton pattern
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
