package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(vote(a, b, c))
}

func vote(x, y, z int) int {
	summa := x + y + z
	if summa >= 2 {
		return 1
	} else {
		return 0
	}
}
