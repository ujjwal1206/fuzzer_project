package main

import (
	"fmt"
)

func integerOverflow() {
	var num uint32 = 4294967295
	fmt.Println(num + 1)
}

func main() {
	integerOverflow()
}
