package main

import "fmt"

func main() {
	var n int
	max := 0
	number := 0

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		if n > max {
			max = n
			number = 0
		}
		if n == max {
			number++
		}
	}
	fmt.Println(number)
}
