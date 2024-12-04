/*
q3. Create an Interface with one method square(int) int
    Create a type that implements this interface
    Create a function Operation that can call the square method using the interface

    In the main function, create a nil pointer to the concrete type
    Pass the value to the operation function

    Operation function calls the method that implements the interface

    Try to do recovery from panic at different levels


*/
package main

import (
	"fmt"
)



type SquareInterface interface {
	square(int)
}

type SquareInterfaceImpl struct{
	squareVal int
}

func (s *SquareInterfaceImpl) square(x int) {
	s.squareVal =  x * x
}

func Operation(s SquareInterface, value int) {
	// if s == nil {
	// 	panic("Nil pointer passed to Operation!")
	// }
	defer recoverPanic() // comment and uncomment to understand
	s.square(value)
	fmt.Println("operation executed successfully:")
}

func recoverPanic () {
	if r := recover(); r != nil {
		fmt.Println("Recovered in main:", r)
	}
}

func main() {
	var squareImpl *SquareInterfaceImpl
	defer recoverPanic()
	Operation(squareImpl, 5)
	fmt.Println("Main executed!!")
}