package hashmap

import (
	"sort"
)

type KV struct {
	Key   rune
	Value int
}

func SortByValue(hashMap map[rune]int) []KV {

	var sortedMap []KV

	for k, v := range hashMap {
		sortedMap = append(sortedMap, KV{k, v})
	}

	// Sort the map by SortByValue
	sort.Slice(sortedMap, func(i, j int) bool {
		return sortedMap[i].Value < sortedMap[j].Value
	})

	return sortedMap
}

func GetFrequencyMap(data string) map[rune]int {
	freqMap := make(map[rune]int)
	// Count the frequency of each character (not byte) in the data
	for _, char := range data {
		freqMap[char]++
	}
	return freqMap
}
