package main

import (
	// "fmt"
	// "io/ioutil"
	"log"
	"os"
)

func main() {
	// content, err := ioutil.ReadFile("./mygofile.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(content))

	file, err := os.Open("./mygofile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
