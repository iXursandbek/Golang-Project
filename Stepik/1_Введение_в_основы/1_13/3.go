package main

import "fmt"

func main() {
	var sekund, soat, minut int
	fmt.Scan(&sekund)
	soat = sekund / 3600
	minut = sekund % 3600 / 60
	fmt.Print(`It is `, soat, ` hours `, minut, ` minutes.`)

}
