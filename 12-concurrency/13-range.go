package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := new(sync.WaitGroup)
	wgWorker := new(sync.WaitGroup)
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 1; i <= 5; i++ {
			//
			wgWorker.Add(1)
			// fan out pattern, spinning up n number of goroutines, for n number of task
			go func(i int) {
				//x90
				defer wgWorker.Done()
				ch <- i
			}(i)

		}
		wgWorker.Wait()
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// range gives a guarantee everything would be received
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg.Wait()

}
