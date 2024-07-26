package hashmap

import (
	"fmt"
	"sort"
)

type KV struct {
	Key   string
	Value int
}

func SortByValue(hashMap map[string]int) []KV {

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

func ReverseMap(data map[string][]byte) map[string]string {
	result := make(map[string]string)
	for key, value := range data {
		// Convert byte array to string using a custom bit-to-string conversion
		strVal := ""
		for _, bit := range value {
			strVal += fmt.Sprintf("%d", bit)
		}
		result[strVal] = key
	}
	return result
}

func GetFrequencyMap(data string) map[string]int {
	freqMap := make(map[string]int)
	// Count the frequency of each character (not byte) in the data
	for _, char := range data {
		freqMap[string(char)]++
	}
	return freqMap
}
