package main

import (
	"fmt"
)

func stackOverflow() {
	var arr [1024]byte
	_ = arr // Prevent unused warning
	stackOverflow()
}

func main() {
	stackOverflow()
	fmt.Println("Crash Logged")
}
