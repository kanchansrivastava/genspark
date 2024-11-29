package main

import (
	"app/stores"
	"app/stores/models"
	"app/stores/mysql"
	"app/stores/postgress"
	"fmt"
)

// Refer question q3 from day 7 and q1 from day8

func main() {
	//Call postgres and mysql package methods using interface variable
	//var i stores.DataBase

	u1 := models.User{Id: 1, Name: "sandra"}
	u2 := models.User{Id: 2, Name: "komal"}
	u3 := models.User{Id: 3, Name: "diwakar"}

	m := mysql.NewConn()
	i := stores.NewStore(m)

	fmt.Printf("---------------- Printing type of interface %T ---------------------------------\n", i)
	//printMapReturn(i.FetchAll())

	printReturn(i.Create(u1))

	printReturn(i.FetchUser(1))

	printReturn(i.Update(1, "new me"))
	printReturn(i.Delete(2))

	printReturn(i.Create(u2))
	printMapReturn(i.FetchAll())

	printReturn(i.Delete(2))

	printMapReturn(i.FetchAll())
	p := postgress.NewConn()
	i = stores.NewStore(p)

	fmt.Printf("---------------- Printing type of interface %T ---------------------------------\n", i)
	printMapReturn(i.FetchAll())

	printReturn(i.Create(u1))

	printReturn(i.FetchUser(1))

	printReturn(i.Update(1, "new me"))
	printReturn(i.Delete(2))

	printReturn(i.Create(u2))
	printMapReturn(i.FetchAll())

	printReturn(i.Delete(2))

	printReturn(i.Create(u3))

	printMapReturn(i.FetchAll())

}

func printReturn(u *models.User, ok bool) {
	if !ok {
		fmt.Println("The command was not executed")
	}
	if u == nil {
		fmt.Println("u :", nil, " ok ", ok, "\n\n")
	} else {
		fmt.Println("u :", u, " ok ", ok, "\n\n")
	}

}

func printMapReturn(userDb map[int]*models.User, ok bool) {
	if !ok {
		fmt.Println("The command was not executed")
	}
	for key, value := range userDb {
		fmt.Println("Fetched values from db")
		fmt.Printf("Key: %v\tValue: %v\n", key, value)
	}

}
