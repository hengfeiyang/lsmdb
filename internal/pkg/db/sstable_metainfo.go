package db

import (
	"bytes"
	"encoding/binary"
)

type SSTableMetaInfo struct {
	DataStart     uint64 // data segment position
	DataLength    uint64 // data segment length
	IndexStart    uint64 // sparse index position
	IndexLength   uint64 // sparse index length
	BlockKeyNum   uint16 // each block contains N keys
	TableBlockNum uint16 // each table contains N blocks
	Version       uint32 // data version
}

func (t *SSTableMetaInfo) Bytes() []byte {
	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, t.DataStart)
	binary.Write(buf, binary.LittleEndian, t.DataLength)
	binary.Write(buf, binary.LittleEndian, t.IndexStart)
	binary.Write(buf, binary.LittleEndian, t.IndexLength)
	binary.Write(buf, binary.LittleEndian, t.BlockKeyNum)
	binary.Write(buf, binary.LittleEndian, t.TableBlockNum)
	binary.Write(buf, binary.LittleEndian, t.Version)
	return buf.Bytes()
}

func (t *SSTableMetaInfo) Restore(data []byte) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &t.DataStart)
	binary.Read(buf, binary.LittleEndian, &t.DataLength)
	binary.Read(buf, binary.LittleEndian, &t.IndexStart)
	binary.Read(buf, binary.LittleEndian, &t.IndexLength)
	binary.Read(buf, binary.LittleEndian, &t.BlockKeyNum)
	binary.Read(buf, binary.LittleEndian, &t.TableBlockNum)
	binary.Read(buf, binary.LittleEndian, &t.Version)
	buf = nil
}
