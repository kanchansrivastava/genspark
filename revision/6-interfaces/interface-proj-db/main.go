package main

import (
	"interface-proj-db/stores"
	"interface-proj-db/stores/models"
	"interface-proj-db/stores/mysql"
	postgress "interface-proj-db/stores/postgres"
)

func main() {
	u := models.User{
		Id:   1,
		Name: "test",
	}
	pg := postgress.NewConn()
	mq := mysql.NewConn()
	stores.Create(pg, u)
	stores.Create(mq, u)
	//var s stores.DataBase = pg
	//fmt.Println(s.Create(u))
	//fmt.Println(s.FetchAll())
	//
	//s = mq
	//fmt.Println(s.Create(u))
}
