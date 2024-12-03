package auth

import (
	"fmt"
	"question2/shared"
)

// Authenticate function prints an authentication message.
func Authenticate() {
	fmt.Println("Authenticating user...")
}

// Name function prints the current user's name.
func Name() {
	fmt.Printf("User Name: %s\n", shared.CurrentUser) // Accessing shared global variable
}
