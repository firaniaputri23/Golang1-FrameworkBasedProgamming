package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// function that takes a string as an argument
// and returns the reverse of the string.
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	var str string
	reader := bufio.NewReader(os.Stdin)

	for {
		// Prompting the user for input
		fmt.Print("Enter a string with at least three words to reverse: ")
		str, _ = reader.ReadString('\n')
		str = strings.TrimSpace(str)

		// Checking if the input has at least three words
		words := strings.Fields(str)
		if len(words) >= 3 {
			break
		} else {
			fmt.Println("Please enter at least three words.")
		}
	}

	// Reversing the input string
	strRev := reverse(str)

	// Displaying the original and reversed strings
	fmt.Println("Original String:", str)
	fmt.Println("Reversed String:", strRev)
}
