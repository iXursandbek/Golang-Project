package main

import "fmt"

func main() {
	var n, a, b, c int
	fmt.Scan(&n)

	a = n / 100
	b = n % 100 / 10
	c = n % 10

	if a != b && b != c && a != c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
