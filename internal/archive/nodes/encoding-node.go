package nodes

import (
	"bytes"
	"encoding/binary"
)

type EncodingNode struct {
	Ranks [256]byte
}

func (node *EncodingNode) Encode() ([]byte, error) {
	return node.Ranks[:], nil
}

func (node *EncodingNode) Decode(data []byte) (*EncodingNode, error) {
	buf := bytes.NewReader(data)
	node = &EncodingNode{}

	if err := binary.Read(buf, binary.LittleEndian, &node.Ranks); err != nil {
		return nil, err
	}

	return node, nil
}
