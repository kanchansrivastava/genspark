/*
q1. Create a function that converts string to int
    use time.Sleep to make this function slow
    Pass context to this function with a certain timeout
    If timeout happens this function should report an error back to main
    If timeout didn't happen, this function should return the actual result of Atoi

*/

package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"sync"
)

func convertStringToInt1(ctx context.Context, str string, wg *sync.WaitGroup) (int, error) {
	resultCh := make(chan int)
	errorCh := make(chan error)

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)

		result, err := strconv.Atoi(str)
		if err != nil {
			errorCh <- err
			return
		}
		resultCh <- result
	}()

	select {
		case <-ctx.Done():
			return 0, fmt.Errorf("operation timed out: %w", ctx.Err())
		case err := <-errorCh:
			return 0, fmt.Errorf("conversion error: %w", err)
		case result := <-resultCh:
			return result, nil
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	wg := &sync.WaitGroup{}
	defer cancel()

	str := "123"

	result, err := convertStringToInt1(ctx, str, wg)
	wg.Wait()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
