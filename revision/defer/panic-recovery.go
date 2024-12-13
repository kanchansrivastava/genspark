package main

import (
	"fmt"
	"runtime/debug"
)

func recoverFunc() {
	msg := recover()
	if msg != nil {
		fmt.Println("panic recovered:\n", msg)
		fmt.Printf("%v", string(debug.Stack()))
	}
}

func updateSlice(s []int, index int, val int) {

	s[index] = val
}
func doWork() {

	updateSlice(nil, 1, 10)
	fmt.Println("end of doWork")
}

func main() {
	defer recoverFunc()
	doWork()
	fmt.Println()
	fmt.Println("end of main")
}
