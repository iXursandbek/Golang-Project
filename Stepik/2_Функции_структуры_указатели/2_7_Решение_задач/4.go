package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string
	fmt.Scan(&text)

	for _, r := range text {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic(fmt.Sprintf("invalid int value: %c", r))
		}

		fmt.Print(n * n)
	}
}
