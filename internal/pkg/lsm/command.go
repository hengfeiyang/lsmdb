package lsm

import (
	"bytes"
	"encoding/binary"
)

type CommandType uint8

const (
	CommandTypeSet CommandType = iota
	CommandTypeDelete
)

type Command struct {
	Key     string
	Value   string
	Command CommandType
}

func (t *Command) Bytes() (int, []byte) {
	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, t.Command)
	binary.Write(buf, binary.LittleEndian, uint32(len(t.Key)))
	buf.Write([]byte(t.Key))
	binary.Write(buf, binary.LittleEndian, uint32(len(t.Value)))
	buf.Write([]byte(t.Value))
	return buf.Len(), buf.Bytes()
}

func (t *Command) Restore(data []byte) {
	var n uint32
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &t.Command)
	binary.Read(buf, binary.LittleEndian, &n)
	t.Key = string(buf.Next(int(n)))
	binary.Read(buf, binary.LittleEndian, &n)
	t.Value = string(buf.Next(int(n)))
	buf = nil
}

type CommandData []*Command

func (t CommandData) Len() int { return len(t) }

func (t CommandData) Less(i, j int) bool { return t[i].Key < t[j].Key }

func (t CommandData) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
