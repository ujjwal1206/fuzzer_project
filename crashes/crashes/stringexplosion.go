package main

import (
	"fmt"
	"strings"
)

func stringExplosion() {
	_ = strings.Repeat("A", 1<<30)
}

func main() {
	stringExplosion()
	fmt.Println("String Explosion Simulated")
}
