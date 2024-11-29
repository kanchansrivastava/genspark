// q2. Create a program to store and print a person's and their project's details. Declare and initialize variables for the following details,
//     Project name (string)
//     Code lines written (uint8)
//     Bugs found (int)
//     Is the project complete? (bool)
//     Average lines of code written per hour (float64)
//     Team lead name (string)
//     Project deadline in days (int)
//     Additionally, demonstrate a uint overflow by initializing the largest possible value for uint and then adding 1 to it


package main

import (
	"fmt"
)

func main(){
	var (
		projectName       string  = "Genspark"
		codeLines         uint8   = 100
		bugsFound         int     = 2
		isProjectComplete bool    = false
		avgLinesPerHour   float64 = 10
		teamLeadName      string  = "Robert"
		daysRemaining     int     = 3
	)

	fmt.Printf("Project Name: %s\n", projectName)
	fmt.Printf("Code Lines Written: %d\n", codeLines)
	fmt.Printf("Bugs Found: %d\n", bugsFound)
	fmt.Printf("Is Project Complete? %t\n", isProjectComplete)
	fmt.Printf("Average Lines of Code per Hour: %.2f\n", avgLinesPerHour)
	fmt.Printf("Team Lead Name: %s\n", teamLeadName)
	fmt.Printf("Project Deadline in Days: %d\n", daysRemaining)

	var largestUint uint8 = 255
	largestUint = largestUint + 1
	fmt.Println(largestUint)
}