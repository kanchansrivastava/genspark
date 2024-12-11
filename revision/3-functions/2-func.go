package main

import (
	"fmt"
)

type operation func(int, int) int
type money int

type person struct{}
type i = interface{} // alias

func main() {

	operate(add, 10, 20)
	operateV2(add, 10, 23)
	//add(10, 20)
}

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operate(op func(int, int) int, x, y int) {
	fmt.Println(op(x, y))
}

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operateV2(op operation, x, y int) {
	fmt.Println(op(x, y))
}

func print(i int, s string) {}
func add(a, b int) int {
	return a + b
}
