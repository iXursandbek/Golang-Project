package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c := float64(a+b) / 2
	fmt.Println(c)
}
