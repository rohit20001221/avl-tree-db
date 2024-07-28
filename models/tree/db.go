package tree

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

type DB struct {
	Store *Tree
	file  *os.File
	size  int64
}

func OpenDB(path string) *DB {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		panic(err)
	}

	fileInfo, _ := file.Stat()

	return &DB{
		file: file,
		size: fileInfo.Size(),
	}
}

func (db *DB) Insert(key uint64, value []byte) {
	wal := WalEntry{
		Value: value,
		WalHeader: WalHeader{
			Key:       key,
			Size:      uint64(len(value)),
			Timestamp: time.Now().Unix(),
		},
	}

	binary.Write(db.file, binary.BigEndian, wal.WalHeader)
	binary.Write(db.file, binary.BigEndian, wal.Value)

	n, err := db.file.Write(value)
	if err != nil {
		fmt.Println(err)
	}

	db.Store = Insert(db.Store, Node{
		Key:    key,
		Offset: db.size + int64(n),
	})

	db.size += int64(n)
}
