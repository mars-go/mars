package db

import (
	"encoding/json"
	bolt "go.etcd.io/bbolt"
)

func (db *DB) BucketSet(bucket string, key string, val interface{}) error {
	return db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		bk, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		// Marshal user data into bytes.
		buf, err := json.Marshal(val)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return bk.Put([]byte(key), buf)
		return nil
	})
}

func (db *DB) BucketGet(bucket string, key string, out interface{}) error {
	return db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		bk := tx.Bucket([]byte(bucket))

		data := bk.Get([]byte(key))
		if len(data) > 0 {
			err := json.Unmarshal(data, out)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (db *DB) BucketFind(bucket string, fn func(bk *bolt.Bucket) error) error {
	return db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		bk := tx.Bucket([]byte(bucket))

		return fn(bk)
	})
}
