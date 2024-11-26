package main

import "fmt"

func main() {
	var p *int
	updateNilPointer(&p)
	fmt.Println(p)
	fmt.Println(*p)
}

func updateNilPointer(p1 **int) {

	x := 10
	fmt.Println(&x)
	*p1 = &x
}
