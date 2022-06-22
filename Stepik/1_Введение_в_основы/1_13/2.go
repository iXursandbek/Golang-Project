package main

import (
	"fmt"
	"strconv"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)
	str := strconv.Itoa(n)
	fmt.Println(reverse(str))
}
