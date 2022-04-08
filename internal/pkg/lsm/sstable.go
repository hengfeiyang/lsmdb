package lsm

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/huandu/skiplist"
)

type SSTable struct {
	data *skiplist.SkipList
}

func NewSSTable() *SSTable {
	return &SSTable{
		data: skiplist.New(skiplist.String),
	}
}

func (t *SSTable) Set(c *Command) {
	t.data.Set(c.Key, c)
}

func (t *SSTable) Query(key string) *Command {
	if e := t.data.Get(key); e != nil {
		return e.Value.(*Command)
	}
	return nil
}

func (t *SSTable) Len() int {
	return t.data.Len()
}

func (t *SSTable) Bytes() []byte {
	buf := bytes.NewBuffer(nil)
	e := t.data.Front()
	for {
		if e == nil {
			break
		}
		v := e.Value.(*Command)
		n, body := v.Bytes()
		binary.Write(buf, binary.LittleEndian, uint32(n))
		buf.Write(body)
		e = e.Next()
	}
	return buf.Bytes()
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
		t.Set(cmd)
	}
	buf = nil

	return nil
}
