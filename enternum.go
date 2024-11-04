package main

import "fmt"

func main() {
	// User's input Variable
	var input int
	// Prompt to ask input from user
	fmt.Print("Input: ")
	// Read input
	fmt.Scan(&input)

	// Check whether input = 42
	if input == 42 {
		// If condition fulfilled, output: "Hello Universe"
		fmt.Println("Hello Universe")
	} else {
		// If not fulfilled, print input
		fmt.Println(input)
	}
}
