package models

import (
	"time"
)

// User represents a user in the encrypted database
type User struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email_encrypted"`
	Phone     string    `json:"phone" db:"phone_encrypted"`
	Address   string    `json:"address" db:"address_encrypted"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// NewUser creates a new user with timestamps
func NewUser(name, email, phone, address string) *User {
	now := time.Now()
	return &User{
		Name:      name,
		Email:     email,
		Phone:     phone,
		Address:   address,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
	if u.Name == "" {
		return ErrInvalidName
	}
	if u.Email == "" {
		return ErrInvalidEmail
	}
	return nil
}

// Errors
var (
	ErrInvalidName  = &ValidationError{"name cannot be empty"}
	ErrInvalidEmail = &ValidationError{"email cannot be empty"}
)

// ValidationError represents a validation error
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
