// package main
//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"huffman/coding"
// 	"os"
// )
//
// func main() {
// 	data, err := os.ReadFile("../files/measurements.txt")
//     if err != nil {
//         fmt.Println("Error reading file: ", err)
//         return
//     }
//     // Convert byte array to string
//     dataStrig := string(data)
//
//     compressedData, table, count := coding.Compress(dataStrig)
//     compressed, err := json.Marshal(compressedData)
//     err = os.WriteFile("../files/compressed.txt", compressed, 0644)
//     if err != nil {
//         fmt.Println("Error writing file: ", err)
//         return
//     }
//     fmt.Println("File written successfully")
//     fmt.Println(table)
//     for key, value := range table {
//         key = string(int(key.(int32)))
//         fmt.Printf("Key: %v, Value: %v\n", key, value)
//     }
// 	decompressedData := coding.Decompress(compressedData, table, count, len(dataStrig))
//     fmt.Printf("Decompressed data: %v\n", decompressedData)
//     // Write decompressed data to file
//     decompressed, err := json.Marshal(decompressedData)
//     if err != nil {
//         fmt.Println("Error writing file: ", err)
//         return
//     }
//     err = os.WriteFile("../files/decompressed.txt", decompressed, 0644)
//     if err != nil {
//         fmt.Println("Error writing file: ", err)
//         return
//     }
//     fmt.Println("File written successfully")
// }

package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "huffman/coding"
    "os"
)

func main() {
    file, err := os.Open("../files/measurements.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var dataStr string
    for scanner.Scan() {
        dataStr += scanner.Text()
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    compressedData, table, count := coding.Compress(dataStr)
    compressed, err := json.Marshal(compressedData)
    if err != nil {
        fmt.Println("Error marshalling compressed data:", err)
        return
    }
    err = os.WriteFile("../files/compressed.txt", compressed, 0644)
    if err != nil {
        fmt.Println("Error writing compressed file:", err)
        return
    }
    fmt.Println("Compressed file written successfully")

    fmt.Println(table)
    for key, value := range table {
        key = string(int(key.(int32)))
        fmt.Printf("Key: %v, Value: %v\n", key, value)
    }

    decompressedData := coding.Decompress(compressedData, table, count, len(dataStr))
    fmt.Printf("Decompressed data: %v\n", decompressedData)

    decompressed, err := json.Marshal(decompressedData)
    if err != nil {
        fmt.Println("Error marshalling decompressed data:", err)
        return
    }
    err = os.WriteFile("../files/decompressed.txt", decompressed, 0644)
    if err != nil {
        fmt.Println("Error writing decompressed file:", err)
        return
    }
    fmt.Println("Decompressed file written successfully")
}
