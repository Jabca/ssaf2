package encodedecode

import (
	"ssaf2/internal/archive/nodes"
	"ssaf2/internal/data_structs"
)

type tNode struct {
	data  byte
	left  *tNode
	right *tNode
}

// whether this node holds data or is it intermediate
func (node tNode) IsDataNode() bool {
	return node.left == nil && node.right == nil
}

// object to decode and encode data. supports encoding of individual bytes
type HuffmanTree struct {
	root *tNode
}

// create huffman tree from EncodingNode
func createHF(node *nodes.EncodingNode) HuffmanTree {
	pq := &data_structs.PriorityQueue[tNode]{}
	var value byte
	var rating uint8
	// initially populate with data nodes
	for i := range len(node.Data) {
		value = byte(i)
		rating = node.Data[i]
		item := &tNode{data: value}
		pq.Enqueue(item, int(rating))
	}

	// unite 2 least rated nodes into one
	for pq.Len() > 1 {
		i1, r1, _ := pq.Dequeue()
		i2, r2, _ := pq.Dequeue()

		pq.Enqueue(
			&tNode{
				left:  i1,
				right: i2,
			},
			r1+r2,
		)
	}

	last_el, _, _ := pq.Dequeue()
	return HuffmanTree{
		root: last_el,
	}
}

func dfs(node *tNode, ba data_structs.BitArray, em *map[byte]data_structs.BitArray) {
	if node.IsDataNode() {
		(*em)[node.data] = ba
	} else {
		if node.left != nil {
			dfs(node.left, ba.PushRight(0), em)
		} else if node.right != nil {
			dfs(node.right, ba.PushRight(1), em)
		}
	}

}

func CreateEncodeMap(node *nodes.EncodingNode) map[byte]data_structs.BitArray {
	tree := createHF(node)
	encMap := map[byte]data_structs.BitArray{}

	dfs(tree.root, data_structs.BitArray{}, &encMap)

	return encMap
}

func CreateDecodeMap(node *nodes.EncodingNode) map[data_structs.BitArray]byte {
	encMap := CreateEncodeMap(node)
	decMap := map[data_structs.BitArray]byte{}
	for key, val := range encMap {
		decMap[val] = key
	}

	return decMap
}
