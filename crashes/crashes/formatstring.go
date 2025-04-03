package main

import (
	"fmt"
	"os"
)

func formatString(input string) {
	fmt.Printf(input)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run format_string.go <input>")
		return
	}
	formatString(os.Args[1])
}
