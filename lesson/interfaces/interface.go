package interfaces

func Qosh(x, y int) int {
	return x + y
}

// func Dars() {
// 	fmt.Println("Nichiksan")
// }
type UserInterface interface {
	Println(User) User
	Change(User) User
}

type User struct {
	Name string
	Age  int
}
