package main

import (
	"fmt"
	"strings"
	"time"
)

// Stage 1: Convert string to uppercase
func toUpperCase(input <-chan string, output chan<- string) {
	for str := range input {
		output <- strings.ToUpper(str)
	}
	close(output)
}

// Stage 2: Add a suffix to the string
func addSuffix(input <-chan string, output chan<- string) {
	for str := range input {
		output <- str + "_SUFFIX"
	}
	close(output)
}

// Stage 3: Print the result
func printResult(input <-chan string) {
	for str := range input {
		fmt.Println("Processed:", str)
	}
}

func main() {
	// Initial input channel
	input := make(chan string)

	// Intermediate channels
	upperCaseChan := make(chan string)
	suffixChan := make(chan string)

	// Step 1: Start the input goroutine
	go func() {
		defer close(input)
		inputData := []string{"hello", "world", "golang"}
		for _, data := range inputData {
			input <- data
		}
	}()

	// Step 2: Start the upper case goroutine
	go toUpperCase(input, upperCaseChan)

	// Step 3: Start the suffix goroutine
	go addSuffix(upperCaseChan, suffixChan)

	// Step 4: Print the final results
	go printResult(suffixChan)

	// Wait for all goroutines to finish (simulate with sleep)
	time.Sleep(2 * time.Second)
}
