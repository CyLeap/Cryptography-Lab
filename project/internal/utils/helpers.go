package utils

import (
	"encrypted-db/internal/models"
	"fmt"
	"os"
	"strings"
)

// GetEnv gets an environment variable or returns a default value
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// PromptPassword prompts the user for a password (basic implementation)
func PromptPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	var password string
	fmt.Scanln(&password)
	return strings.TrimSpace(password), nil
}

// ValidatePassword performs basic password validation
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	return nil
}

// FormatUser displays a user in a readable format
func FormatUser(user *models.User) string {
	return fmt.Sprintf("ID: %d\nName: %s\nEmail: %s\nPhone: %s\nAddress: %s\nCreated: %s\nUpdated: %s\n",
		user.ID, user.Name, user.Email, user.Phone, user.Address,
		user.CreatedAt.Format("2006-01-02 15:04:05"),
		user.UpdatedAt.Format("2006-01-02 15:04:05"))
}
