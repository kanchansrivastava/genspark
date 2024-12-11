package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) UpdateName(name string) {
	p.Name = name
}

func main() {
	p := Person{"John", 25}
	p1 := []Person{
		{"John", 25},
		{"Jane", 26},
	}
	fmt.Println(p1[0].Age)
	println(p.Name)

}
