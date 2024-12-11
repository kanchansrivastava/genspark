package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) greet() {
	fmt.Println("Hello, my name is", p.name)
}
func (p *person) setAge(age int) {
	p.age = age // no dereference
}
func main() {
	p := person{"John", 30}
	p.greet()
	p.setAge(31) // you don't have to pass p with &
	fmt.Println(p.age)
}
