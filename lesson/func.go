package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func Print(u interface{}) {
	fmt.Println(u)
}
