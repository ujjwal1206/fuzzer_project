package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"io/ioutil"
)

// 1. Buffer Overflow (Panic on large input)
func bufferOverflow(input string) {
	var buf [10]byte
	copy(buf[:], input)
	fmt.Println(string(buf[:]))
	logCrash(1, input)
}

// 2. Infinite Recursion (Stack Overflow)
func stackOverflow() {
	var arr [1024]byte
	_ = arr // Prevent unused warning
	stackOverflow()
	logCrash(2, "Infinite Recursion")
}

// 3. Memory Leak Simulation
func memoryLeak() {
	leak := [][]byte{}
	for {
		leak = append(leak, make([]byte, 1024*1024))
		logCrash(3, "Memory Leak")
	}
}

// 4. Integer Overflow
func integerOverflow() {
	var num uint32 = 4294967295
	fmt.Println(num + 1)
	logCrash(4, "Integer Overflow")
}

// 6. Uninitialized Memory Read (Nil Pointer Dereference)
func nilPointer() {
	var ptr *int
	fmt.Println(*ptr)
	logCrash(6, "Nil Pointer Dereference")
}

// 7. Use-After-Free (Simulated by closing a file and using it)
func useAfterFree() {
	f, _ := os.Create("temp.txt")
	f.Close()
	fmt.Fprintln(f, "Writing after closing!")
	logCrash(7, "Use-After-Free")
}

// 8. Format String Exploit (Unsafe User Input)
func formatString(input string) {
	fmt.Printf(input)
	logCrash(8, input)
}

// 9. Out-of-Bounds Read
func outOfBounds() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[10])
	logCrash(9, "Out-of-Bounds Read")
}

// 10. String Explosion (Excessive Memory Allocation)
func stringExplosion() {
	_ = strings.Repeat("A", 1<<30)
	logCrash(10, "String Explosion")
}

// Log crashing input to .crash file
func logCrash(progNum int, input string) {
	fileName := fmt.Sprintf("crash_%d.crash", progNum)
	ioutil.WriteFile(fileName, []byte(input), 0644)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run crashable.go <program_number> <input>")
		return
	}

	progNum, _ := strconv.Atoi(os.Args[1])
	input := os.Args[2]

	// Save test input to .txt file
	testFileName := fmt.Sprintf("test_%d.txt", progNum)
	ioutil.WriteFile(testFileName, []byte(input), 0644)

	switch progNum {
	case 1:
		bufferOverflow(input)
	case 2:
		stackOverflow()
	case 3:
		memoryLeak()
	case 4:
		integerOverflow()
	case 5:
		divideByZero()
	case 6:
		nilPointer()
	case 7:
		useAfterFree()
	case 8:
		formatString(input)
	case 9:
		outOfBounds()
	case 10:
		stringExplosion()
	default:
		fmt.Println("Invalid program number")
	}
}
