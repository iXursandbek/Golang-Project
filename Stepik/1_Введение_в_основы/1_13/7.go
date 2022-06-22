package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n)
	var Slice = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&Slice[i])
		if Slice[i] == 0 {
			k++
		}
	}
	fmt.Println(k)
}
