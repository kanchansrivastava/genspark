// q4. Print default values and Type names of variables from question 2 using printf
// // Quick Tip, Use %v if not sure about what verb should be used,
// // but don't use it in this question :)
// // but generally using %v should be fine


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

	fmt.Printf("Project Name: %s, Variable Type:%T \n", projectName, projectName)
	fmt.Printf("Code Lines Written: %d, Variable Type:%T \n", codeLines, codeLines)
	fmt.Printf("Bugs Found: %d, Variable Type:%T \n", bugsFound, bugsFound)
	fmt.Printf("Is Project Complete? %t, Variable Type:%T \n", isProjectComplete, isProjectComplete)
	fmt.Printf("Average Lines of Code per Hour: %.2f , Variable Type:%T \n", avgLinesPerHour, avgLinesPerHour)
	fmt.Printf("Team Lead Name: %s , Variable Type:%T \n", teamLeadName, teamLeadName)
	fmt.Printf("Project Deadline in Days: %d , Variable Type:%T \n", daysRemaining, daysRemaining)
}