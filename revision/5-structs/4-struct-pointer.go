package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{
		Name: "John",
		Age:  30,
	}
	updateUser(&u)
	fmt.Println(u)
}

func updateUser(u *User) {
	u.Name = "Jane"
}
