package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
	"strconv"
)

// Mutate input by randomly changing bytes
func mutateInput(input []byte, progNum string) []byte {
	rand.Seed(time.Now().UnixNano())
	mutated := make([]byte, len(input))
	copy(mutated, input)

	for i := range mutated {
		if rand.Intn(100) < 13 { // 13% chance of mutation
			newByte := byte(rand.Intn(256))
			
			// If testing divide-by-zero (case 5), avoid zero mutations
			if progNum == "5" && newByte == 0 {
				newByte = 1 + byte(rand.Intn(255)) // Ensure non-zero value
			}

			mutated[i] = newByte
		}
	}

	return mutated
}



func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run fuzzer.go <program_number> <seed_file> <num_iterations>")
		return
	}

	progNum := os.Args[1]
	seedFile := os.Args[2]
	numIterations := 1000
	if n, err := strconv.Atoi(os.Args[3]); err == nil {
		numIterations = n
	}

	// Read seed file
	seed, err := os.ReadFile(seedFile)
	if err != nil {
		fmt.Println("Error reading seed file:", err)
		return
	}

	input := seed
	for i := 0; i < numIterations; i++ {
		input = mutateInput(input, progNum)

		cmd := exec.Command("go", "run", "crashable.go", progNum, string(input))
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()

		if err != nil {
			fmt.Printf("Crash detected at iteration %d!\n", i)
			crashFile := fmt.Sprintf("crash_%d_%d.txt", progNum, i)
			os.WriteFile(crashFile, input, 0644)
			return
		}
	}
}
