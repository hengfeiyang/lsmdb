package db

import (
	"log"
	"os"
	"path"
)

const BlockKeyNum uint16 = 512  // each block contains N keys
const TableBlockNum uint16 = 64 // each table contains N blocks

var DB *MEMSSTable

func init() {
	var err error
	DB, err = NewMEMSSTable("./data", BlockKeyNum, TableBlockNum)
	if err != nil {
		log.Fatal(err)
	}
	if err := restoreWAL(); err != nil {
		log.Fatal(err)
	}
	if err := loadSparseIndex(); err != nil {
		log.Fatal(err)
	}
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
