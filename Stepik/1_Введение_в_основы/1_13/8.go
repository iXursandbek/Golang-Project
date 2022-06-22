package main

import "fmt"

func main() {
	var n, min, k int
	fmt.Scan(&n)

	slice := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	min = slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	for i := 0; i < n; i++ {
		if slice[i] == min {
			k++
		}
	}
	fmt.Println(k)

}
