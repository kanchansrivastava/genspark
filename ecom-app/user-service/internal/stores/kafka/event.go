package kafka

import "time"

const (
	TopicAccountCreated = `user-service.account-created`
	ConsumerGroup       = `user-service`
)

type MSGUserServiceAccountCreated struct {
	ID        string    `json:"id"` // UUID
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"` // Timestamp of creation
	UpdatedAt time.Time `json:"updated_at"` // Timestamp of last update
}
