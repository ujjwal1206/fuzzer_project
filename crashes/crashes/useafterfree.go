package main

import (
	"fmt"
	"os"
)

func useAfterFree() {
	f, _ := os.Create("temp.txt")
	f.Close()
	fmt.Fprintln(f, "Writing after closing!")
}

func main() {
	useAfterFree()
}
