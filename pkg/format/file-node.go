package format

import (
	"bytes"
	"encoding/binary"
	"time"
)

type FileNode struct {
	// fixed sized fields
	ID             [4]byte
	IsDir          bool
	ModTime        time.Time
	ParentID       [4]byte
	FileSize       uint64
	CheckSum       [4]byte
	DataNodesCount uint64
	DataPointer    uint64
	// variable sized fields
	Name string
}

func (node *FileNode) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	// Write fixed-size fields (ID, IsDir, Size, ModTime, Mode)
	if err := binary.Write(buf, binary.LittleEndian, node.ID); err != nil {
		return nil, err
	}

	nameBytes := []byte(node.Name)
	if err := binary.Write(buf, binary.LittleEndian, uint16(len(nameBytes))); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, node.IsDir); err != nil {
		return nil, err
	}

	modTimeUnix := node.ModTime.UnixNano()
	if err := binary.Write(buf, binary.LittleEndian, modTimeUnix); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, node.ParentID); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, node.FileSize); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, node.CheckSum); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, node.DataNodesCount); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, node.DataPointer); err != nil {
		return nil, err
	}
	// Write variable-length fields (Name, Data)

	if err := binary.Write(buf, binary.LittleEndian, uint16(len(nameBytes))); err != nil {
		return nil, err
	}
	if _, err := buf.Write(nameBytes); err != nil {
		return nil, err
	}
}
