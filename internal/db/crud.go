package db

import (
	"encoding/json"
	"fmt"
	bolt "go.etcd.io/bbolt"
)

func Set(bucket string, key string, val interface{}) error {
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
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
}

func Get(bucket string, key string, out interface{}) error {
	tx, err := db.Begin(false)
	if err != nil {
		return err
	}
	// Assume bucket exists and has keys
	bk := tx.Bucket([]byte(bucket))

	data := bk.Get([]byte(key))
	if len(data) > 0 {
		err = json.Unmarshal(data, out)
		if err != nil {
			return err
		}
	}
	return nil
}

func Get(bucket string, key string, out interface{}) error {

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucket))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})

	tx, err := db.Begin(false)
	if err != nil {
		return err
	}
	// Assume bucket exists and has keys
	bk := tx.Bucket([]byte(bucket))

	data := bk.Get([]byte(key))
	if len(data) > 0 {
		err = json.Unmarshal(data, out)
		if err != nil {
			return err
		}
	}
	return nil
}
