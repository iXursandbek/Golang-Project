package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := (strconv.FormatBool(true)) == (10 == int(float64(100/10)))
	fmt.Println(res)
}
