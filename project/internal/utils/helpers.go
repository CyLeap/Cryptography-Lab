package utils

import (
	"encrypted-db/internal/models"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
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
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Println() // Print a newline after password input
	return string(password), nil
}

// ValidatePassword performs comprehensive password validation
func ValidatePassword(password string) error {
	if len(password) < 8 || len(password) > 12 {
		return fmt.Errorf("password must be between 8 and 12 characters long")
	}
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false
	for _, char := range password {
		switch {
		case char >= 'A' && char <= 'Z':
			hasUpper = true
		case char >= 'a' && char <= 'z':
			hasLower = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case (char >= '!' && char <= '/') || (char >= ':' && char <= '@') || (char >= '[' && char <= '`') || (char >= '{' && char <= '~'):
			hasSymbol = true
		}
	}
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
