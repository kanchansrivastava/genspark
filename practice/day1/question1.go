// q1. Declare a variable to represent temperature in Celsius.
// Convert this temperature to Fahrenheit using the formula

// Use fmt.Printf instead of Println
// // fmt.Printf verbs could be found on top of fmt documentation // https://pkg.go.dev/fmt#hdr-Printing
// Use type casting in Go

package main

import (
	"fmt"
)

func main()  {
	var celsiusTemperature float32 = 25
	fahrenheitTemperature := (celsiusTemperature * 1.8) + 32
	fmt.Println(fahrenheitTemperature)
	fmt.Printf("Temperature in Fahrenheit: %.2f\n", fahrenheitTemperature) 
}

