package main

import (
	"fmt"
)

func main() {
	var input float64
	fmt.Print("Please provide your floating point value: ")
	fmt.Scan(&input)
	fmt.Printf("Truncated number is: %d\n", int64(input))
}