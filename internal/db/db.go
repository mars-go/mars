package db

import bolt "go.etcd.io/bbolt"

var db *DB

type DB struct {
	*bolt.DB
}

func Init(path string) error {
	bt, err := bolt.Open(path, 0666, nil)
	if err != nil {
		return err
	}
	db = &DB{
		DB: bt,
	}
	return nil
}

func Get() *DB {
	return db
}

func Close() {
	if db != nil {
		db.Close()
	}
}
