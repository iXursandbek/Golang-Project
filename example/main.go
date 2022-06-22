package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Golangda fayllarga xush kelibsiz!")
	content := "Bu yerga go file kerak"

	file, err := os.Create("./mygofile.txt")
	chechNillErr(err)

	length, err := io.WriteString(file, content)
	chechNillErr(err)

	fmt.Println("fayllar", length)
	defer file.Close()
	readFile("./mygofile.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	chechNillErr(err)
	fmt.Println("Text malumotlari fayl ichida \n", databyte)
}

func chechNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
