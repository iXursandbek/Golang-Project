package main

import "fmt"

func main() {
	var A int
	fmt.Scan(&A)
	nFib, notFib := 1, -1
	a, b := 1, 1

	for a != A {
		a, b = b, a+b

		nFib++
		if a == A {
			fmt.Println(nFib)
		} else if a > A {
			fmt.Println(notFib)
			break
		}
	}
}
