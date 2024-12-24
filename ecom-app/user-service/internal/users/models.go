package users

import "time"

// User struct represents the users table in the stores
type User struct {
	ID               string    `json:"id"` // UUID
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	PasswordHash     string    `json:"-"`          // Password hash (not exposed in JSON)
	StripeCustomerID string    `json:"-"`          // Not part of json output
	CreatedAt        time.Time `json:"created_at"` // Timestamp of creation
	UpdatedAt        time.Time `json:"updated_at"` // Timestamp of last update
}
