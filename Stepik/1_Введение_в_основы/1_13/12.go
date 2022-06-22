package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		nn := math.Pow(2, float64(i))
		if nn <= float64(n) {
			fmt.Print(nn, " ")
		} else {
			break
		}
	}
}
