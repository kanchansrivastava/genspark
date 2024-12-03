package auth

import (
	"fmt"
)

// Authenticate function prints an authentication message.
func Authenticate(userName string) {
	fmt.Printf("Authenticating user... %s\n", userName)
}
