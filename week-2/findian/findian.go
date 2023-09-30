package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	// var input string
	fmt.Print("Please enter a word: ")

	// fmt.Scanln(&input)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	loweredInput := strings.ToLower(input)
	trimmedInput := strings.TrimSpace(loweredInput)

	if trimmedInput[0] == 'i' && trimmedInput[len(trimmedInput) -1] == 'n' && strings.Contains(trimmedInput, "a") {
		fmt.Println("Found!\n")
	} else {
		fmt.Println("Not Found!\n")
	}

}