package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func bufferOverflow(input string) {
	var buf [10]byte
	copy(buf[:], input)
	fmt.Println(string(buf[:]))
	logCrash(1, input)
}

func logCrash(progNum int, input string) {
	fileName := fmt.Sprintf("crash_%d.crash", progNum)
	ioutil.WriteFile(fileName, []byte(input), 0644)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run buffer_overflow.go <input>")
		return
	}
	bufferOverflow(os.Args[1])
}
