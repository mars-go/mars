package db

import bolt "go.etcd.io/bbolt"

var db *bolt.DB

func Init(path string) error {
	var err error
	db, err = bolt.Open(path, 0666, nil)
	return err
}

func DB() *bolt.DB {
	return db
}

func Close() {
	if db != nil {
		db.Close()
	}
}
