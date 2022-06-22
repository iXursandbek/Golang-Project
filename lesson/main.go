package main

import (
	"lesson/interfaces"
	// "lesson/interfaces/postgres"
	"lesson/interfaces/postgres/anonim"
)

func main() {
	foydalanuvchi := User{"palonkas", 34}
	Print(foydalanuvchi)

	Print(interfaces.Qosh(23, 4))
	// interfaces.Dars()

	// db := postgres.Load()

	funksiya := anonim.GetAn()
	Print(funksiya(23))
}
