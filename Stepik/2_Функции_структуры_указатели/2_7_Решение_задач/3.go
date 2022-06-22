package main

import "fmt"

func main() {
	var text string
	fmt.Scan(&text)

	runeStr := []rune(text)

	max := runeStr[0]
	for i := 0; i < len(runeStr); i++ {
		if max < runeStr[i] {
			max = runeStr[i]
		}
	}
	fmt.Println(string(max))
}
