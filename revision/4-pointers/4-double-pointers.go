package main

import "fmt"

func main() {
	var p *int           // nil
	updateNilPointer(&p) // passing pointer reference
	fmt.Println(*p)
}

// p[p1[nil]]
// *p -> p1[]
// *p = &x // p1[&x]

func updateNilPointer(p1 **int) {
	x := 10
	*p1 = &x
}
