package main

import (
	"fmt"
)

func nilPointer() {
	var ptr *int
	fmt.Println(*ptr)
}

func main() {
	nilPointer()
}
