package main

import "fmt"

func main() {
	operate(add, 10, 20)
	//add(10, 20)
}

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operate(op func(int, int), x, y int) {
	op(x, y)
}

func print(i int, s string) {}
func add(a, b int) int {
	fmt.Println(a + b)
}
