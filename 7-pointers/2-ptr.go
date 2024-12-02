package main

import (
	"fmt"
)

func main() {
	x := 10 // x80
	update(&x)
	fmt.Println(x)
}

// let's assume p have address of x90
// x90 = x80
func update(p *int) {
	*p = 20 // changing the memory address directly of x variable from the main function

}
