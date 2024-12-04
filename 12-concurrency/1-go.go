package main

import (
	"fmt"
	"time"
)

func main() {
	//panic("panic") // reveals goroutine id and name
	go hello()

	fmt.Println("end of the main")

	// worst case to wait for goroutines
	time.Sleep(time.Second) // sleep would put the main in the blocking state
	// go scheduler would pick up another goroutine if available

}

func hello() {
	fmt.Println("Hello, world!")
}
