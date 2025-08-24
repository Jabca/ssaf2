package nodes

import (
	"bytes"
	"encoding/binary"
)

type DataNode struct {
	fileID      uint64
	batchNumber uint64
	storedCount uint64
	data        []byte
}

func (node *DataNode) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, node.fileID); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, node.batchNumber); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, node.storedCount); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, node.data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}

func DecodeDataNode(data []byte, batchSize uint64) (*DataNode, error) {
	reader := bytes.NewReader(data)
	node := &DataNode{}

	if err := binary.Read(reader, binary.LittleEndian, &node.fileID); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.batchNumber); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.storedCount); err != nil {
		return nil, err
	}
	dataBytes := make([]byte, batchSize)
	if _, err := reader.Read(dataBytes); err != nil {
		return nil, err
	}

	node.data = dataBytes

	return node, nil

}
