package models

import (
	"strings"
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
	if len(u.Name) < 2 || len(u.Name) > 50 {
		return &ValidationError{"name must be between 2 and 50 characters"}
	}
	if !strings.HasSuffix(strings.ToLower(u.Email), "@gmail.com") {
		return &ValidationError{"email must be a valid Gmail address ending with @gmail.com"}
	}
	if u.Phone != "" && (len(u.Phone) < 10 || len(u.Phone) > 15 || !isNumeric(u.Phone)) {
		return &ValidationError{"phone must be 10-15 digits only (no letters)"}
	}
	if len(strings.TrimSpace(u.Address)) == 0 {
		return &ValidationError{"address cannot be empty"}
	}
	return nil
}

// isNumeric checks if a string contains only digits
func isNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
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
