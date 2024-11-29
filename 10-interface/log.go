package main

import (
	"log"
)

type user struct {
	name  string
	email string
}

func main() {
	u := user{"raj", "raj@email.com"}
	l := log.New(u, "log: ", log.LstdFlags)
	l.Println("Hello, log")
}
