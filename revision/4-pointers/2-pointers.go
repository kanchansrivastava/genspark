package main

import (
	"fmt"
)

func main() {
	x := 10           // x80
	var ptr *int = &x // ptr address x120, value is x80
	update(ptr)       // update(x80)
	fmt.Println(x)
	fmt.Println(*ptr)
}

// let's assume p have address of x90
// x90 = x80
func update(p *int) {
	*p = 20 // changing the memory address directly of x variable from the main function
	// x80 = 20
}
