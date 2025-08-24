package nodes

import (
	"bytes"
	"encoding/binary"
)

type HeaderNode struct {
	magicBytes      [4]byte
	headerLEngth    uint8
	version         [3]uint8
	batchSize       uint64
	nodesCount      uint64
	filesSectionPtr uint64
	dataSectionPtr  uint64
}

func (node *HeaderNode) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, node.magicBytes); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, node.headerLEngth); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, node.version); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, node.batchSize); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, node.nodesCount); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, node.filesSectionPtr); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, node.dataSectionPtr); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}

func DecodeHeader(data []byte) (*HeaderNode, error) {
	reader := new(bytes.Reader)
	node := &HeaderNode{}
	if err := binary.Read(reader, binary.LittleEndian, &node.magicBytes); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.headerLEngth); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.version); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.batchSize); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.nodesCount); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.filesSectionPtr); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.LittleEndian, &node.dataSectionPtr); err != nil {
		return nil, err
	}

	return node, nil

}
