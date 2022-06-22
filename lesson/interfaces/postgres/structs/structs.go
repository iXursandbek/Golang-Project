package structs

import (
	"lesson/interfaces"
)

type UserInter interface {
	Xullas() interfaces.UserInterface
}

type UserStruct struct {
	userI interfaces.UserInterface
}

func (u UserStruct) Xullas() interfaces.UserInterface {
	return u.userI
}

func (u UserStruct) Change(interfaces.User) interfaces.User {
	return interfaces.User{}
}

func (u UserStruct) Println(interfaces.User) interfaces.User {
	return interfaces.User{}
}
