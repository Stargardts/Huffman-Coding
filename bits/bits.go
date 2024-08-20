package bits

import (
	"fmt"
)

type BitVector struct {
	Vector          []byte
	CurrentByte     uint32
	CurrentBit      uint32
	SignificantBits uint32
}

func (b *BitVector) AddBit(bit byte) {
	if b.CurrentBit == 8 {
		b.CurrentBit = 0
		b.CurrentByte += 1
	}
	if b.CurrentByte == uint32(len(b.Vector)) {
		b.Vector = append(b.Vector, 0)
	}
	b.Vector[b.CurrentByte] |= bit << b.CurrentBit
	b.CurrentBit += 1
	b.SignificantBits += 1
}

func ExtractNextBit(byteArray []byte, bitIndex uint32) (byte, error) {
	if bitIndex/8 >= uint32(len(byteArray)) {
		return 0, fmt.Errorf("Index out of range")
	}
	bit := (byteArray[bitIndex/8] << (7 - (bitIndex % 8)) >> 7)

	return bit, nil
}

func Encode(bitMap map[any][]byte, bitCount map[any]uint32, s string) BitVector {
	var bitVector BitVector
	bitVector.Vector = make([]byte, 0)

	for _, char := range s {
		for i := uint32(0); i < bitCount[char]; i++ {
			bit, err := ExtractNextBit(bitMap[char], i)
			if err != nil {
				continue
			}
			bitVector.AddBit(bit)
		}
	}

	return bitVector
}

func ExtreactBits(byteArray []byte, bitCount uint32, index int) []byte {
	var result BitVector
	for i := 0; i < int(bitCount); i++ {
		bit, err := ExtractNextBit(byteArray, uint32(index+i))
		if err != nil {
			continue
		}
		result.AddBit(bit)
	}
	return result.Vector
}

func Decode(bitMap map[any][]byte, bitCount map[any]uint32, bitVector []byte, count uint32) any {
	var result []any
	index := 0
	for i := uint32(0); i < count; i++ {
		for key, value := range bitMap {
			x := true
			buffer := ExtreactBits(bitVector, bitCount[key], index)
			if CompareBits(buffer, value) {
				result = append(result, key)
				index += int(bitCount[key])
				x = false
			}
			if !x {
				break
			}
		}
	}
	return result
}

func CompareBits(byteArray1 []byte, byteArray2 []byte) bool {
	for i := range byteArray1 {
		if byteArray1[i] != byteArray2[i] {
			return false
		}
	}

	return true
}
