package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func Print(u User) {
	fmt.Println(u)
}
