/*
q3. Create a new custom type based on float64 for handling temperatures in Celsius.
    Implement the Following Methods (not functions):
    Method 1: ToFahrenheit
    Description: Converts the Celsius temperature to Fahrenheit.
    Signature: ToFahrenheit() float64
    Method 2: IsFreezing
    Description: Checks if the temperature is at or below the freezing point (0°C).
    Signature: IsFreezing() bool

*/

package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64((c * 1.8) + 32)
}

func (c Celsius) IsFreezing() bool {
	return c <= 0
}

func main() {
	var temp Celsius = 25
	// var temp Celsius = 0
	
	fmt.Printf("Temperature in Celsius: %.2f°C\n", temp)
	fmt.Printf("Temperature in Fahrenheit: %.2f°F\n", temp.ToFahrenheit())

	if temp.IsFreezing() {
		fmt.Println("The temperature is freezing.")
		return
	} 
	fmt.Println("The temperature is not freezing.")
}
	