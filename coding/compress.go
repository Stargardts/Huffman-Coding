package coding

import (
	b "huffman/bits"
	"huffman/hashmap"
	"huffman/tree"
)

// Compress compresses the input string using Huffman coding
func Compress(input string) (b.BitVector, map[any][]byte, map[any]uint32) {
	frequencyMap := hashmap.GetFrequencyMap(input)
	sortedMap := hashmap.SortByValue(frequencyMap)
	huffmantree := tree.BuildTree(sortedMap)
	table, count := make(map[any][]byte), make(map[any]uint32)
	table, count = huffmantree.BuildTable(table, count)

	return b.Encode(table, count, input), table, count
}
