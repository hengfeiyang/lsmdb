package db

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"sort"
)

type SSTable struct {
	data CommandData
}

func NewSSTable() *SSTable {
	return &SSTable{}
}

func (t *SSTable) Append(c *Command) {
	t.data = append(t.data, c)
}

func (t *SSTable) Query(key string) *Command {
	for i := range t.data {
		if t.data[i].Key == key {
			return t.data[i]
		}
	}
	return nil
}

func (t *SSTable) Sort() {
	sort.Sort(t.data)
}

func (t *SSTable) Len() int {
	return len(t.data)
}

func (t *SSTable) Bytes() (int, []byte) {
	buf := bytes.NewBuffer(nil)
	for i := 0; i < len(t.data); i++ {
		n, body := t.data[i].Bytes()
		binary.Write(buf, binary.LittleEndian, uint32(n))
		buf.Write(body)
	}
	return buf.Len(), buf.Bytes()
}

func (t *SSTable) Restore(data []byte) error {
	buf := bytes.NewBuffer(data)
	var n uint32
	for {
		if err := binary.Read(buf, binary.LittleEndian, &n); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		cmd := new(Command)
		cmd.Restore(buf.Next(int(n)))
		t.data = append(t.data, cmd)
		fmt.Printf("SSTable.Restore: %+v\n", cmd)
	}
	buf = nil

	return nil
}
