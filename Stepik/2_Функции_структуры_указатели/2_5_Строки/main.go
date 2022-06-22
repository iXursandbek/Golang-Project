package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		strings.Contains("test", "est"),
		strings.Count("test", "t"),
		strings.HasPrefix("test", "te"),
		strings.HasSuffix("test", "st"),
		strings.Index("test", "s"),
		strings.Join([]string{"Assalom", "alekum"}, "-"),
		strings.Repeat("naGap ", 5),
		strings.Replace("bolaflar", "la", "***", -3),
		strings.Split("Xursandbek", ""),
		strings.ToLower("NEWS"),
		strings.ToUpper("yangiliklar"),
		strings.Trim("tetstet", "te"),
	)
}
