package coding

import (
	b "huffman-coding/bits"
	"huffman-coding/hashmap"
	"huffman-coding/tree"
)

// Compress compresses the input string using Huffman coding
func Compress(input string) (b.BitVector, map[rune][]byte, map[rune]uint32) {
	frequencyMap := hashmap.GetFrequencyMap(input)
	sortedMap := hashmap.SortByValue(frequencyMap)
	huffmantree := tree.BuildTree(sortedMap)
	table, count := make(map[rune][]byte), make(map[rune]uint32)
	table, count = huffmantree.BuildTable(table, count)

	return b.Encode(table, count, input), table, count
}
