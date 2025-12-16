package utils

import (
	"encrypted-db/internal/models"
	"fmt"
	"os"
	"regexp"

	"github.com/howeyc/gopass"
)

const (
	MinPasswordLength = 8
	MaxPasswordLength = 12
)

// GetEnv gets an environment variable or returns a default value
func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// PromptPassword prompts the user for a password with masked input
func PromptPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	password, err := gopass.GetPasswdMasked()
	if err != nil {
		return "", fmt.Errorf("failed to read password: %v", err) // Wrapping the error for better clarity
	}
	return string(password), nil
}

// ValidatePassword performs comprehensive password validation
func ValidatePassword(password string) error {
	if len(password) < MinPasswordLength || len(password) > MaxPasswordLength {
		return fmt.Errorf("password must be between %d and %d characters long", MinPasswordLength, MaxPasswordLength)
	}

	// Regex for checking upper, lower, digit, and symbol
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSymbol := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	// Checking conditions
	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}
	if !hasSymbol {
		return fmt.Errorf("password must contain at least one symbol")
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
