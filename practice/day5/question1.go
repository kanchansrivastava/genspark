/*

q1. In main function create a slice of names & Add two elements in it.

    Create a function AddNames which appends new names to the slice
    Use double pointer concept to make AddNames function work

*/

package main

import "fmt"

func main(){
	names:= []string {"Kanchan", "Divya"};
	fmt.Println(names)
	AddNames(&names)
	fmt.Println("After addnames", names)
}

func AddNames(names *[]string){ // double pointer passed
	*names = append(*names, "Myname")
}