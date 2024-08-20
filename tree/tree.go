package tree

import (
	"huffman/hashmap"
    b "huffman/bits"
)

type HuffmanTree struct {
	currentNode Node
	leftNode    Node
	rightNode   Node
	size        int
}

func (h *HuffmanTree) Insert(kv hashmap.KV) {
	if h.leftNode == nil {
		node := &primaryNode{
			name:      kv.Key,
			frequency: kv.Value,
		}
		h.leftNode = node
		h.rightNode = node
		h.size++
		h.currentNode = node
	} else if h.rightNode == h.leftNode {
		node := &primaryNode{
			name:      kv.Key,
			frequency: kv.Value,
		}
		nodeFreq := node.Frequency()
		leftFreq := h.leftNode.Frequency()
		// rightFreq := h.rightNode.Frequency()
		if nodeFreq <= leftFreq {
			h.leftNode = node
		} else {
			h.rightNode = node
		}
		h.size++
		h.currentNode = node
	} else {
		var leftFreq, rightFreq int
		leftFreq = h.leftNode.Frequency()
		rightFreq = h.rightNode.Frequency()

		node := &linkNode{
			value:      leftFreq + rightFreq,
			leftChild:  h.leftNode,
			rightChild: h.rightNode,
		}
		h.leftNode = node
		h.rightNode = node
		h.size++
		h.currentNode = node
		h.Insert(kv)
	}
}

func (h *HuffmanTree) inserthead() {
	node := &linkNode{
		value:      h.leftNode.Frequency() + h.rightNode.Frequency(),
		leftChild:  h.leftNode,
		rightChild: h.rightNode,
	}
	h.leftNode = nil
	h.rightNode = nil
	h.size++
	h.currentNode = node
}

type primaryNode struct {
	name              any
	frequency         int
	bitRepresentation b.BitVector
}

type linkNode struct {
	value             int
	leftChild         Node
	rightChild        Node
	bitRepresentation b.BitVector
}

func (n *primaryNode) nodeType() string {
	return "primary"
}

func (n *linkNode) setVector(b b.BitVector) {
	n.bitRepresentation = b
}

func (n *primaryNode) setVector(b b.BitVector) {
	n.bitRepresentation = b
}

func (n *primaryNode) Frequency() int {
	return n.frequency
}

func (n *linkNode) nodeType() string {
	return "link"
}

func (n *linkNode) Frequency() int {
	return n.value
}

func (n *primaryNode) Name() any {
	return n.name
}

func (n *linkNode) Name() any {
	return 0
}

type Node interface {
	nodeType() string
	Frequency() int
	Name() any
	setVector(b.BitVector)
}

func NewTree() *HuffmanTree {
	return &HuffmanTree{}
}

func BuildTree(slice []hashmap.KV) *HuffmanTree {
	tree := NewTree()
	for _, kv := range slice {
		tree.Insert(kv)
	}
	tree.inserthead()
	return tree
}

func traverse(node Node, table map[any][]byte, count map[any]uint32) {
	if node.nodeType() == "primary" {
		table[node.(*primaryNode).name] = node.(*primaryNode).bitRepresentation.Vector
        count[node.(*primaryNode).name] = node.(*primaryNode).bitRepresentation.SignificantBits
	} else {
        significantBits := node.(*linkNode).bitRepresentation.SignificantBits
		currentByte := node.(*linkNode).bitRepresentation.CurrentByte
		currentBit := node.(*linkNode).bitRepresentation.CurrentBit
		vector := node.(*linkNode).bitRepresentation.Vector
		left := make([]byte, len(vector))
		copy(left, vector)
		right := make([]byte, len(vector))
		copy(right, vector)
        leftVector := b.BitVector{Vector: left, CurrentByte: currentByte, CurrentBit: currentBit, SignificantBits: significantBits}
        rightVector := b.BitVector{Vector: right, CurrentByte: currentByte, CurrentBit: currentBit, SignificantBits: significantBits}
		leftVector.AddBit(0)
		rightVector.AddBit(1)
		node.(*linkNode).leftChild.setVector(leftVector)
		node.(*linkNode).rightChild.setVector(rightVector)
		traverse(node.(*linkNode).leftChild, table, count)
		traverse(node.(*linkNode).rightChild, table, count)
    }
}

func (h *HuffmanTree) BuildTable(table map[any][]byte, count map[any]uint32) (map[any][]byte, map[any]uint32) {
	if table == nil {
		table = make(map[any][]byte)
	}
    if count == nil {
        count = make(map[any]uint32)
    }
	traverse(h.currentNode, table, count)
	return table, count
}
