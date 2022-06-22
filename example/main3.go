package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	var f *os.File

	fil1, _ := os.Create("text.txt")
	fil1.WriteString("1")
	fil1.WriteString("3")

	switch fil1.Name() {
	case "text.txt":
		os.Rename("text.txt", "txttxt.txt")
		fallthrough
	case "texttt.txt":
		os.Remove("text.txt")
		fallthrough
	default:
		var err error
		f, err = os.Open("textxt.txt")
		fil1.WriteString("2")
		if err != nil {
			panic(err)
		}
	}

	res, _ := ioutil.ReadAll(f)
	resInt, err := strconv.Atoi(string(res[2]))
	if err != nil {
		panic(err)
	}
	
	fmt.Println(resInt + 1)
	fil1.Close()
}
