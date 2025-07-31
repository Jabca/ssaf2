package ssaf2_nodes

import (
	"bytes"
	"encoding/binary"
)

type EncodingNode struct {
	data [256]byte
}

func (node *EncodingNode) Encode() ([]byte, error) {
	return node.data[:], nil
}

func (node *EncodingNode) Decode(data []byte) (*EncodingNode, error) {
	buf := bytes.NewReader(data)
	node = &EncodingNode{}

	if err := binary.Read(buf, binary.LittleEndian, &node.data); err != nil {
		return nil, err
	}

	return node, nil
}
