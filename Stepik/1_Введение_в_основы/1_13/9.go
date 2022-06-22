package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	var summa int

	for n > 0 {
		digit := n % 10
		summa += int(digit)
		n = n / 10
		if summa > 10 {
			summa = summa%10 + summa/10
		}
	}
	fmt.Println(summa)
}
