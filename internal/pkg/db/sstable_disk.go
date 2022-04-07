package db

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/pierrec/lz4"
)

type DiskSSTable struct {
	filename string
	Blocks   map[uint32]*SSTable
}

func NewDiskSSTable(filename string) (*DiskSSTable, error) {
	return &DiskSSTable{filename, map[uint32]*SSTable{}}, nil
}

func (t *DiskSSTable) Query(blockIndex uint32, seek uint32, key string) (*Command, error) {
	if t.Blocks[blockIndex] == nil {
		if err := t.LoadBlock(blockIndex, seek); err != nil {
			return nil, err
		}
	}
	if v := t.Blocks[blockIndex].Query(key); v != nil {
		return v, nil
	}

	return nil, errors.New("key not exists in DiskSSTable")
}

func (t *DiskSSTable) LoadBlock(blockIndex uint32, seek uint32) error {
	f, err := os.Open(t.filename)
	if err != nil {
		return err
	}
	f.Seek(int64(seek), io.SeekStart)
	var n uint32
	if err = binary.Read(f, binary.LittleEndian, &n); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}
	data := make([]byte, int(n))
	nn, err := f.Read(data)
	if err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}
	if nn != int(n) {
		return fmt.Errorf("DiskSSTable.LoadBlock error: data length mismatch %d, %d", n, nn)
	}

	// decompress
	lz4buf := bytes.NewBuffer(data)
	lz4r := lz4.NewReader(lz4buf)
	unData := bytes.NewBuffer(nil)
	nnn, err := io.Copy(unData, lz4r)
	fmt.Println(blockIndex, seek, n, nn, nnn, err)
	block := NewSSTable()
	if err := block.Restore(unData.Bytes()); err != nil {
		return err
	}
	t.Blocks[blockIndex] = block

	f.Close()
	return nil
}
