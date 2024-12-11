package main

import "fmt"

func main() {

	var ptr *int     // ptr address = x120, value = nil
	updateValue(ptr) // updateAdd(nil)
	//fmt.Println(*ptr)
}

// p have address x150
// x150 = nil
func updateValue(p *int) {
	var x int = 10
	p = &x // // x150 = x80
	*p = 20
	fmt.Println(*p)
	//fmt.Println(*p)
	//fmt.Println(&p)
}
