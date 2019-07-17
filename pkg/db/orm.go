package db

import (
	"encoding/json"
	bolt "go.etcd.io/bbolt"
	"strconv"
	"time"
)

func (db *DB) BucketCreate(bucket string, v Modeler) error {
	return db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		bk, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		//检查是否是Model，自动添加时间
		model, ok := checkModel(v)
		if ok {
			now := time.Now()
			model.CreatedAt = now
			model.UpdatedAt = now
		}

		id := v.PK()
		if id <= 0 {
			id, err = bk.NextSequence()
			if err != nil {
				return err
			}
			v.SetPK(id)
		}

		// Marshal user data into bytes.
		buf, err := json.Marshal(v)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return bk.Put([]byte(strconv.FormatUint(id, 10)), buf)
		return nil
	})
}

func (db *DB) BucketUpdate(bucket string, v Modeler) error {
	return db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		bk, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		//检查是否是Model，自动添加时间
		model, ok := checkModel(v)
		if ok {
			model.UpdatedAt = time.Now()
		}

		id := v.PK()

		// Marshal user data into bytes.
		buf, err := json.Marshal(v)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return bk.Put([]byte(strconv.FormatUint(id, 10)), buf)
		return nil
	})
}
