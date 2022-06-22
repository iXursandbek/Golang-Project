package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n >= 11 && n <= 14 {
		fmt.Println(n, "korov")
	} else {
		temp := n % 10
		if temp == 0 || (temp >= 5 && temp <= 9) {
			print(n, "korov")
		}

		if temp == 1 {
			print(n, "korova")
		}

		if temp >= 2 && temp <= 4 {
			print(n, "korovy")
		}
	}
}
