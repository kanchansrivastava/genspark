/* q2. Create 3 functions f1, f2, f3
    f1() call f2(), f2() call f3()
    each layer would return the error, wrap the error from each layer
    print stack trace using debug.Stack to get a complete stack trace

*/

package main

import (
	"fmt"
	"runtime/debug"
)

func f1() error {
	err := f2()
	if err != nil {
		return fmt.Errorf("f1 encountered an error: %w", err)
	}
	return nil
}

func f2() error {
	err := f3()
	if err != nil {
		return fmt.Errorf("f2 encountered an error: %w", err)
	}
	return nil
}

func f3() error {
	fmt.Println(string(debug.Stack()))
	return fmt.Errorf("f3 encountered an error")
}

func main() {
	err := f1()
	if err != nil {
		fmt.Printf("Error occurred: %s\n", err) // implement it to get complete stack trace
		return
	}
	fmt.Println("Ran successfully")
}
