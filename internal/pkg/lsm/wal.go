package lsm

import (
	"encoding/binary"
	"os"
)

type wal struct {
	filename string
	f        *os.File
}

func NewWAL(filename string) (*wal, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &wal{filename, f}, nil
}

func (t *wal) Append(c *Command) error {
	n, body := c.Bytes()
	binary.Write(t.f, binary.LittleEndian, uint32(n))
	_, err := t.f.Write(body)
	return err
}

func (t *wal) Remove() error {
	if err := t.f.Close(); err != nil {
		return err
	}
	return os.Remove(t.filename)
}

func (t *wal) Close() error {
	return t.f.Close()
}
