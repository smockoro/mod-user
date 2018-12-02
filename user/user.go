package user

import "fmt"

type User struct {
	Name   string
	Age    int
	Height int
	Weight int
}

func (u *User) UserHello() {
	fmt.Println("mod-user v2.1.0")
	fmt.Println(u.Name)
	fmt.Println(u.Age)
	fmt.Println(u.Height)
	fmt.Println(u.Weight)
}
