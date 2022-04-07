package db

import (
	"bytes"
	"encoding/binary"
)

type SparseIndex struct {
	Key        string // key name
	DataStart  uint32 // data start position in table
	BlockIndex uint32 // which block contains the key
	TableName  string //  which table contains the block
}

func (t *SparseIndex) Bytes() (int, []byte) {
	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, t.BlockIndex)
	binary.Write(buf, binary.LittleEndian, t.DataStart)
	binary.Write(buf, binary.LittleEndian, uint32(len(t.Key)))
	buf.Write([]byte(t.Key))
	return buf.Len(), buf.Bytes()
}

func (t *SparseIndex) Restore(data []byte) {
	var n uint32
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &t.BlockIndex)
	binary.Read(buf, binary.LittleEndian, &t.DataStart)
	binary.Read(buf, binary.LittleEndian, &n)
	t.Key = string(buf.Next(int(n)))
	buf = nil
}
