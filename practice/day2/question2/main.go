
// q2.     second-proj-day-2/
//         ├── main.go
//         ├── go.mod
//         ├── auth/
//         │   └── auth.go
//         └── user/
//             └── user.go

//     In auth package create two functions
//     1. Authenticate
//     		Authenticate function simply prints a message, authenticating user
//     2. Name
//     		This function prints the Name of the user.

//     Note:- to print the name of the user,
//     use the user package to know who is the current user

//     In user package create one global variable, and one func named as AddToDb
//     1. AddToDb
//     		This function accepts database name as string
//     		It calls the Authenticate function from auth package
//     		At last it prints a msg, Adding to db DatabaseName [var which was accepted in the parameter]

//     Global Variable
//     1. CurrentUser = "any name here" // this would be fetched by auth package

//     Note:- Q2 should not work, it should give some import cycle issues, it is intended

//     **How to solve the import cycle**
//     		Extract the common piece of functionality in a separate package
//     		Import the new package where functionality is needed
//     ****************************************************

package main

import (
	"question2/user"
)

func main() {
	user.AddToDb("my_database")
}