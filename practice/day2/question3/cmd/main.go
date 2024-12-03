package main

import (
	"fmt"
	"question3/auth"
)

func main() {
	Setup()       
	auth.Authenticate("David")
	fmt.Println("End of the main")
}
