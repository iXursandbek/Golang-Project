package main

import "fmt"

func main() {
	var n, a, b, c int
	fmt.Scan(&n)
	a = n / 100
	b = n % 100 / 10
	c = n % 10
	d := a + b + c
	fmt.Println(d)
}
