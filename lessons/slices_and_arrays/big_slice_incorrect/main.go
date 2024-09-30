package main

import (
	"fmt"
	"runtime"
)

func printAllocs() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}

func FindData(filename string) []byte {
	data := make([]byte, 1<<30) // for example read from file

	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0x00 && data[i+1] == 0x00 {
			return data[i : i+20]
		}
	}

	return nil
}

func main() {
	data := FindData("data.bin")
	_ = data

	printAllocs()
	runtime.GC()
	printAllocs()

	runtime.KeepAlive(data)
}