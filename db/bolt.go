package db

import (
	"log"
	"time"

	"github.com/boltdb/bolt"
)

type DB struct {
	DBPath string
}

func (d *DB) Put(key []byte, value []byte) error {
	db, err := bolt.Open(d.DBPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("lucid"))
		if err != nil {
			return err
		}
		b.Put(key, value)
		return nil
	})
	return nil
}

func (d *DB) Get(key []byte) string {
	var value string
	db, err := bolt.Open(d.DBPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Printf("%v", err)
		return ""
	}
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("lucid"))
		v := b.Get(key)
		value = string(v[:])
		return nil
	})
	return value
}
