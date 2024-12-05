	/*q1. create a function work that takes a work id and print work {id} is going on
	In the main function run a loop to run work function 10 times
	make the work function call concurrent
	Make sure your program waits for work function to finish gracefully
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	for i:=1;i<=10;i++ {
		wg.Add(1)
		go work1(i, &wg)
	}
	wg.Wait()
	fmt.Println("All work completed!")
}

func work1(wordId int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("anonymous goroutine started", wordId)
		time.Sleep(100 * time.Millisecond) 
		fmt.Println("anonymous goroutine finished", wordId)
	}()
	fmt.Println("main work id: ", wordId, "is going on")
}