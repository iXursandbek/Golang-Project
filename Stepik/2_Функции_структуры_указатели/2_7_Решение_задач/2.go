package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	str = strings.TrimSpace(str)
	
	bs := []rune(str)

	fmt.Print(string(bs[0]))
	fmt.Print(strings.Replace(str[1:len(str)-1], "", "*", -1))
	fmt.Print(string(bs[len(str)-1]))

}
