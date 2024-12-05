package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()

		fmt.Println("sender picked up")
		ch <- 1
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		fmt.Println("recv picked up")
		fmt.Println(<-ch)
	}()
	wg.Wait()
}
