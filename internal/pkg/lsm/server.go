package lsm

import (
	"errors"
	"log"
	"os"
	"path"
	"time"
)

const BlockKeyNum uint16 = 512  // each block contains N keys
const TableBlockNum uint16 = 64 // each table contains N blocks

var ErrorNotExist = errors.New("key not exists")

var DB *MEMTable

func init() {
	var err error
	DB, err = NewMEMTable("./data", BlockKeyNum, TableBlockNum)
	if err != nil {
		log.Fatal(err)
	}
	if err := restoreWAL(); err != nil {
		log.Fatal(err)
	}
	if err := loadSparseIndex(); err != nil {
		log.Fatal(err)
	}
	go compactionSSTable()
}

func loadSparseIndex() error {
	fs, err := os.ReadDir("./data")
	if err != nil {
		return err
	}
	for _, f := range fs {
		if path.Ext(f.Name()) == ".sdb" {
			sf, err := os.Open("./data/" + f.Name())
			if err != nil {
				return err
			}
			if err := DB.LoadFromDiskTable(sf); err != nil {
				return err
			}
			sf.Close()
		}
	}

	return nil
}

func restoreWAL() error {
	fs, err := os.ReadDir("./data")
	if err != nil {
		return err
	}
	for _, f := range fs {
		if path.Ext(f.Name()) == ".wal" {
			sf, err := os.Open("./data/" + f.Name())
			if err != nil {
				return err
			}
			if err := DB.LoadFromWAL(sf); err != nil {
				return err
			}
			sf.Close()
		}
	}

	return nil
}

func compactionSSTable() {
	for {
		time.Sleep(10 * time.Minute)
		if err := compaction(); err != nil {
			log.Printf("compactionSSTable error: %v", err)
		}
	}
}

func compaction() error {
	return nil
}
