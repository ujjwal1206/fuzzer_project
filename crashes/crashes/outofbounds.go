package main

import (
	"fmt"
)

func outOfBounds() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[10])
}

func main() {
	outOfBounds()
}
