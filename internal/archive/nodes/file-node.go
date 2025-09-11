package nodes

import (
	"bytes"
	"encoding/binary"
	"time"
)

type FileNode struct {
	// fixed sized fields
	ID             uint64
	ModTime        time.Time
	ParentID       [4]byte
	FileSize       uint64
	CheckSum       [32]byte
	DataNodesCount uint64
	DataOffset     int64
	IsDir          bool
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

	if err := binary.Write(buf, binary.LittleEndian, node.DataOffset); err != nil {
		return nil, err
	}
	// Write variable-length fields (Name)

	if _, err := buf.Write(nameBytes); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DecodeFileNode(data []byte) (*FileNode, error) {
	reader := bytes.NewReader(data)
	node := &FileNode{}

	// Read fixed-size fields
	if err := binary.Read(reader, binary.LittleEndian, &node.ID); err != nil {
		return nil, err
	}
	var nameLen uint16
	if err := binary.Read(reader, binary.LittleEndian, nameLen); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.IsDir); err != nil {
		return nil, err
	}
	var modTimeUnix int64
	if err := binary.Read(reader, binary.LittleEndian, &modTimeUnix); err != nil {
		return nil, err
	}
	node.ModTime = time.Unix(0, modTimeUnix)
	if err := binary.Read(reader, binary.LittleEndian, &node.ParentID); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.FileSize); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.CheckSum); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.DataNodesCount); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.DataOffset); err != nil {
		return nil, err
	}

	nameBytes := make([]byte, nameLen)
	if _, err := reader.Read(nameBytes); err != nil {
		return nil, err
	}
	node.Name = string(nameBytes)

	return node, nil
}
