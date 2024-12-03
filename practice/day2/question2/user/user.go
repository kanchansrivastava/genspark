package user

import (
	"fmt"
	"question2/auth"
	"question2/shared"
)

// AddToDb accepts a database name and prints a message after authenticating
func AddToDb(dbName string) {
	auth.Authenticate() // Calls Authenticate from the auth package
	fmt.Printf("Adding to db: %s\n", dbName)
	fmt.Printf("Current User is: %s\n", shared.CurrentUser)
}
