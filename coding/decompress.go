package coding

import (
	b "huffman/bits"
)

// Decompress decompresses the input string using Huffman coding
func Decompress(input b.BitVector, table map[any][]byte, count map[any]uint32,size int) any {
    return b.Decode(table, count, input.Vector, uint32(size))
}
