package domain

import (
	"github.com/google/uuid"
	"time"
)

// User represents a user entity
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
