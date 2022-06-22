package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)

	rs := []rune(text)

	// fmt.Println(rs)

	if unicode.IsUpper(rs[0]) && strings.HasSuffix(text, ".") {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
