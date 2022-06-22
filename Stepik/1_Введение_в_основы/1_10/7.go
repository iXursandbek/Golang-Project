package main

import "fmt"

func main() {
	var x, p, y, yil int
	fmt.Scan(&x, &p, &y)

	for ; x < y; yil++ {
		x += x * p / 100
	}
	fmt.Println(yil)
}
