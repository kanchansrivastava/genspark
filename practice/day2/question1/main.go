// q1. Create a stringops package with three files. The package should export one function that utilizes internal (unexported) functions.
//
//	Files in stringops Package:
//	File 1: strings.go
//	Exported Function: ReverseAndUppercase(s1, s2 string) string
//	File 2: upper.go
//	Internal Function: toUpperCase(s string) string
//	File 3: reverse.go
//	Internal Function: reverseString(s string) string
package main

import (
	"fmt"
	"question1/stringops" // no numbers,hypen,upper case only alphabet for packages
)

func main() {
	result := stringops.ReverseAndUppercase("hello", "world     ")
	fmt.Println(result)
}
