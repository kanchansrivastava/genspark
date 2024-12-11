package main

import (
	"errors"
	"fmt"
)

type operation func(int, int) int
type money int

// alias

func main() {

	operateV3(addV2(), 10, 20)
	//f := func(x int, y int) int {
	//	return x + y
	//}

	//add(10, 20)
}

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operateV3(op func(int, int) int, x, y int) {
	fmt.Println("inside the operate func")
	fmt.Println(op(x, y))
}

func addV2() func(int, int) int {
	return func(x int, y int) int {
		fmt.Println("doing add inside the add func")
		return x + y
	}
}

func DoSomething() error {
	return DoError()
}

func DoError() error {
	return errors.New("error")
}
