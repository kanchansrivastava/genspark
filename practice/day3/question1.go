// Q1. Write a Go program that:
//     Creates an empty slice with an initial capacity of 1.
//     Appends integers from 1 to 1,000,000 to the slice.
//     Tracks and prints the capacity change every time the slice's capacity increases.
//     Prints the total number of capacity changes at the end.

//     Formula:= (currentCap-lastCap) / lastCap * 100
//     // Hint :- use type casting

package main

import "fmt"

func main() {
	capMySlice := 10000
	myslice := make([]int, 0, capMySlice) // predefined cap == lessser allocation; creating & dropping expensive ops
	maxNum := 1000000
	prevCap := capMySlice
	changeCounter := 0

	for i := 1; i <= maxNum; i++ {
		myslice = append(myslice, i)
		currentCap := cap(myslice)
		if currentCap != prevCap {
			changeCounter++
			rateOfChange := (float64(currentCap-prevCap) / float64(prevCap)) * 100

			fmt.Printf("Capacity changed to: %d\n", currentCap)
			fmt.Printf("Rate of capacity change is: %.2f%%\n", rateOfChange)

			prevCap = currentCap
		}
	}

	fmt.Println("Total number of capacity changes in the program:", changeCounter)

}
