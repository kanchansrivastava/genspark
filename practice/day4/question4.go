/*


q4. Modify the above function to perform the following action
    stringManipulation(trimSpace(), "\ngfdngbk \n"))
    Instead of passing trimSpace you need to call the trimSpace function and make the program work

    Hint: you need to return a function with the signature of what stringManipulation accepts as first parameter


*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Trimspace:", StringManipulation1(trimSpaceV1(), " \t\n Hello, Gophers \n\t\r\n"))
	fmt.Println("ToUpper:", StringManipulation1(toUpperV1(), "what is happening here"))
	fmt.Println("Greet:", StringManipulation1(greetV1(), "George"))
}

func StringManipulation1(myop func(string) string, str string) string{
	result := myop(str)
	return result
}


func trimSpaceV1() func(string) string{
	return func(str string) string {
		return  strings.TrimSpace(str)
	}
}

func toUpperV1() func(string) string{
	return func(str string) string {
		return strings.ToUpper(str)
	}
}

func greetV1() func(string) string{
	return func(name string) string {
		return "Hello! " + name
	}
}