package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N, n int
	fmt.Scan(&N, &n)

	strN := strconv.Itoa(N)
	strn := strconv.Itoa(n)

	res := strings.ReplaceAll(strN, strn, "")
	fmt.Println(res)
}
