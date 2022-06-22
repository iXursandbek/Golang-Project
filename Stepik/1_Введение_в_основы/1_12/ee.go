package main

import "fmt"

func main() {
	a := [5]int{5, 4, 3, 2, 1}

	for i := range a {
		fmt.Println(a[i])
	}

	for i, _ := range a {
		fmt.Print(a[i])
	}

	for _, v := range a {
		fmt.Println(v)
	}
}
