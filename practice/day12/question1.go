/*

q1. Create 4 functions
    Add(int,int),Sub(int,int),Divide(int,int), CollectResults()
    Add,Sub,Divide do their operations and push value to an unbuffered channel

    CollectResult() -> It would receive the values from the channel and print it
*/

package main

import "fmt"

func Add(a, b int, ch chan int) {
	result := a + b
	ch <- result
}

func Sub(a, b int, ch chan int) {
	result := a - b
	ch <- result
}

func Divide(a, b int, ch chan int) {
	if b == 0 {
		ch <- 0
		return
	}
	result := a / b
	ch <- result
}

func CollectResults(ch chan int, numResults int) {
	result1 := <-ch
	result2 := <-ch
	result3 := <-ch
	fmt.Println("Result of addition:", result1)
	fmt.Println("Result of substraction:", result2)
	fmt.Println("Result of division:", result3)
}

func main() {
	ch := make(chan int)

	go Add(10, 5, ch)
	go Sub(10, 5, ch)
	go Divide(10, 5, ch)

	CollectResults(ch, 3)

	close(ch)
}
