package main

import "fmt"

var abc = 20

func main() {
	x := 10

	updateVal(&x)
	fmt.Println(x)
	fmt.Println(abc)
}
func updateVal(px *int) {

	px = &abc
	*px = 100
}
