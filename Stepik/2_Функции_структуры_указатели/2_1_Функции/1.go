package main

import (
	"fmt"
)

func main() {
	text := "Hello World"
	f(text)
}

func f(text string) {
	fmt.Println(text)
}
