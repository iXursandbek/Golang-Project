package main

import "fmt"

func minimumFour() int {
	var m int
	arr := make([]int, 4)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	for i, e := range arr {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

func main() {
	fmt.Println(minimumFour())
}
