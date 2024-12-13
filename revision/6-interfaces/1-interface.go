package main

import "fmt"

// Polymorphism means that a piece of code changes its behavior depending on the
// concrete data it’s operating on // Tom Kurtz, Basic inventor

type Reader interface {
	// only method signatures could be added to an interface
	read(b []byte) (int, error)
	// interfaces are automatically implemented if method signature is same as of interfaces
	//write(b []byte) (int, error) // all methods must be implemented to implement the interface
}

type File struct {
	name string
}

func (f File) read(b []byte) (int, error) {
	fmt.Println("reading files and processing them", f.name)
	return 0, nil
}

type IO struct {
	name string
}

func (i IO) read(b []byte) (int, error) {
	fmt.Println("reading and processing ", i.name)
	return 0, nil
}

func DoWork(r Reader) {
	i, err := r.read(nil)
	_, _ = i, err
	fmt.Printf("DoWork %T\n", r)
	f, ok := r.(File)
	if ok {
		fmt.Println("File", f.name)
	}
}

func main() {
	f := File{name: "file1"}
	i := IO{name: "io1"}
	DoWork(f)
	DoWork(i)
}
