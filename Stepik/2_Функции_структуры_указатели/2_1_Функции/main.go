package main

import "fmt"

func main() {
	// var age, name = add(3, 4, "Jumaniyoz", "Sobirov")
	// _, ism := add(5, 6, "Ubaydullo", "Omon")
	// fmt.Println(ism)
	// fmt.Println(age, name)
	// // fmt.Println(name)

	a := 200
	// var b *int = &a
	b := &a
	*b++
	fmt.Println(b)
}

// func add(x, y int, firstname, lastname string) (int, string) {
// 	var z int = x + y
// 	var fullname string = lastname + " " + firstname
// 	return z, fullname
// }
