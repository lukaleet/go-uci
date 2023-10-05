package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"sort"
)


func main() {
	sli := make([]int, 3)

	fmt.Println("Enter an integer. To quit write 'X': ")

	for i := 0; ; i++ {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		t := strings.TrimSpace(input)

		if t == "X" {
			break
		}

		num, _ := strconv.Atoi(t)

		if i < len(sli) {
			sli[i] = num
		} else {
			sli = append(sli, num)
		}
	}

	sort.Ints(sli)

	fmt.Println(sli)
}