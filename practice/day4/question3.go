/*
q3. Create a function named as StringManipulation.

    	StringManipulation accepts a function and string type as an argument, and it returns string value
    	Possible Functions that it can accept:- trimSpace, toUpper, greet

    Create 3 functions trimSpace, toUpper, greet

    TrimSpace:- TrimSpace returns a string, with all leading and trailing white space removed, as defined by Unicode.
    ToUpper:- ToUpper returns string with all Unicode letters mapped to their upper case.
    Greet: - It takes a name as input, add hello as greeting and return the greeting
    
	Hint: use strings package for TrimSpace and ToUpper functionalities
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Trimspace:", StringManipulation(trimSpace, " \t\n Hello, Gophers \n\t\r\n"))
	fmt.Println("ToUpper:", StringManipulation(toUpper, "what is happening here"))
	fmt.Println("Greet:", StringManipulation(greet, "George"))
}

func StringManipulation(myop func(string) string, str string) string{
	result := myop(str)
	return result
}

func trimSpace(str string) string{
	return  strings.TrimSpace(str)
}

func toUpper(str string) string{
	return strings.ToUpper(str)
}

func greet(name string) string{
	return "Hello! " + name
}