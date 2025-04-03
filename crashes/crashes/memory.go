package main

import (
	"fmt"
)

func memoryLeak() {
	leak := [][]byte{}
	for {
		leak = append(leak, make([]byte, 1024*1024))
	}
}

func main() {
	memoryLeak()
	fmt.Println("Memory Leak Simulated")
}
