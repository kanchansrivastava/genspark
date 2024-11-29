// q1. Create a stringops package with three files. The package should export one function that utilizes internal (unexported) functions.
//     Files in stringops Package:
//     File 1: strings.go
//     Exported Function: ReverseAndUppercase(s1, s2 string) string
//     File 2: upper.go
//     Internal Function: toUpperCase(s string) string
//     File 3: reverse.go
//     Internal Function: reverseString(s string) string
package main

import (
	"question1/stringops"
)

func main(){
	stringops.ReverseAndUppercase("hello", "world")
}