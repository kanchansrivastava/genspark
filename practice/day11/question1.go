	/*q1. create a function work that takes a work id and print work {id} is going on
	In the main function run a loop to run work function 10 times
	make the work function call concurrent
	Make sure your program waits for work function to finish gracefully
*/

package main

import (
	"fmt"
	"sync"
)

func main(){
	wg := new(sync.WaitGroup)

	for i:=1;i<=10;i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			work(id)
		}(i)
	}
	wg.Wait()
	fmt.Println("All work completed!")
}

func work(wordId int) {
	fmt.Println("work id: ", wordId, "is going on")
}
