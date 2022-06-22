package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	t := strconv.Itoa(n)

	fmt.Println(string(t[0]))
}
