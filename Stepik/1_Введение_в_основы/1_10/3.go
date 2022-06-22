package main

import "fmt"

func main() {
	var n, num int
	fmt.Scan(&n)

	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&num)

		if num >= 10 && num < 100 && num%8 == 0 {
			sum += num
		}
	}
	fmt.Println(sum)
}
