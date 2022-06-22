package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)

	c = int(math.Sqrt(float64(a*a) + float64(b*b)))
	fmt.Println(c)

}
