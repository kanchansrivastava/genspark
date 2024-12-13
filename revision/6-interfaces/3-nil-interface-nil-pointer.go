package main

import (
	"database/sql"
	"fmt"
)

type DB interface {
	ReadAll() (string, bool)
}

type conn struct {
	db *sql.DB
}

func (c *conn) ReadAll() (string, bool) {
	if c.db == nil {
		return "", false
	}
	return "all values from the db", true
}

func main() {
	c := &conn{}

	//var db DB = c
	Read(c)

	fmt.Println("main func done")

}

func Read(db DB) (string, bool) {
	// interface nil checking
	if db == nil {
		return "", false
	}

	c, ok := db.(*conn) // type assertion
	if !ok {
		fmt.Println("conf is not present in the interface")
		return "", false
	}
	if c == nil {
		fmt.Println("conf is not set")
		return "", false
	}
	s, ok := db.ReadAll()
	fmt.Println(s, ok)
	return s, ok
}
