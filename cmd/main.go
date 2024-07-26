package main

import (
	"fmt"
	"huffman-coding/coding"
	"os"
)

func main() {
	data, err := os.ReadFile("../sample/sample-input.txt")
    if err != nil {
        fmt.Println("Error reading file: ", err)
        return
    }
    // Convert byte array to string
    dataStrig := string(data)

    compressedData, table, count := coding.Compress(dataStrig)
    for key, value := range table {
        fmt.Printf("Key: %s, Value: %08b\n", key, value)
    }
	decompressedData := coding.Decompress(compressedData, table, count, len(dataStrig))
    fmt.Printf("Decompressed data: %s\n", decompressedData)
    // Write decompressed data to file
    err = os.WriteFile("../sample/sample-output.txt", []byte(decompressedData), 0644)
    if err != nil {
        fmt.Println("Error writing file: ", err)
        return
    }
    fmt.Println("File written successfully")
}
