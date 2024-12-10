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
)

// func convertStringToInt(ctx context.Context, str string) (int, error) {
// 	if ctx.Err() != nil {
// 		return 0, fmt.Errorf("timed out: %w", ctx.Err())
// 	}

// 	time.Sleep(2 * time.Second)

// 	if ctx.Err() != nil {
// 		return 0, fmt.Errorf("timed out: %w", ctx.Err())
// 	}

// 	result, err := strconv.Atoi(str)
// 	if err != nil {
// 		return 0, fmt.Errorf("conversion error: %w", err)
// 	}
// 	return result, nil
// }


func convertStringToInt(ctx context.Context, str string) (int, error) {
	result, err := strconv.Atoi(str)
	time.Sleep(2 * time.Millisecond) // time.Sleep(2 * time.Second) gives context deadline exceeded the millisecond one would pass
	select {
		case <-ctx.Done():
			return 0, fmt.Errorf("timed out in convertStringToInt: %w", ctx.Err())

		default:
			if err != nil {
				return 0, fmt.Errorf("conversion error: %w", err)
			}
			return result, nil
		}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// result, err := convertStringToInt(ctx, "123")
	result, err := convertStringToInt(ctx, "abc")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)
}
