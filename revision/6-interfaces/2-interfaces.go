package main

import "fmt"

type Speaker interface {
	Speak() string
	Modify() string
}

type Person struct {
	name string
}

func (p Person) Speak() string {
	return fmt.Sprintf("Hi, my name is %s", p.name)
}
func (p *Person) Modify() string {
	return fmt.Sprintf("Hi, modifying %s", p.name)
}

func main() {
	p := Person{name: "John"}
	p.Modify()
	var s Speaker = &p

	s.Speak()

}

/*
                          +-------------------------------------+
                          |              Method Set             |
                          +------------------+------------------+
                          |    Value Type    |   Pointer Type   |
+-------------------------+------------------+------------------+
| Function w/ Value Rec.  |        Yes       |       Yes        |
| ( func (t T) )          |                  |                  |
+-------------------------+------------------+------------------+
| Function w/ Ptr Rec.    |        No        |       Yes        |
| ( func (t *T) )         |                  |                  |
+-------------------------+------------------+------------------+

If a function is implemented with a value receiver (func (t T)), it can be called through a value or a pointer.
If a function is implemented with a pointer receiver (func (t *T)), it can only be called through a pointer.
*/
