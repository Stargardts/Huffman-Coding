package coding

import (
	b "huffman-coding/bits"
)

// Decompress decompresses the input string using Huffman coding
func Decompress(input b.BitVector, table map[rune][]byte, count map[rune]uint32,size int) string {
    return b.Decode(table, count, input.Vector, uint32(size))
}
