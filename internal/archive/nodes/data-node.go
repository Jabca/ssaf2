package nodes

import (
	"bytes"
	"encoding/binary"
)

type DataNode struct {
	BatchNumber uint64
	StoredCount uint32
	Data        []byte
}

func (node *DataNode) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, node.BatchNumber); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, node.StoredCount); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, node.Data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}

func DecodeDataNode(data []byte, batchSize uint64) (*DataNode, error) {
	reader := bytes.NewReader(data)
	node := &DataNode{}

	if err := binary.Read(reader, binary.LittleEndian, &node.BatchNumber); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.StoredCount); err != nil {
		return nil, err
	}
	dataBytes := make([]byte, batchSize)
	if _, err := reader.Read(dataBytes); err != nil {
		return nil, err
	}

	node.Data = dataBytes

	return node, nil

}
