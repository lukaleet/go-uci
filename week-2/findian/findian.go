package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Please enter a word: ")

	fmt.Scan(&input)

	var loweredInput = strings.ToLower(input)

	if loweredInput[0] == 'i' && input[len(loweredInput) -1] == 'n' && strings.Contains(loweredInput, "a") {
		fmt.Println("Found!\n")
	} else {
		fmt.Println("Not Found!\n")
	}

}