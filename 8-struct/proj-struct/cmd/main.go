package main

import (
	"fmt"
	"proj-struct/database"
)

type user struct {
	name string
}

func main() {

	// c.db = "mysql" // not allowed, db is not exported
	c := database.NewConf("mysql")
	fmt.Println(c)
	c.Ping()

	//u := new(user)
	//u1 := &user{}

}
