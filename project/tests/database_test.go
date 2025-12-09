package tests

import (
	"os"
	"testing"

	"encrypted-db/internal/database"
	"encrypted-db/internal/models"
)

const testPassword = "testpassword123"

func TestDatabaseOperations(t *testing.T) {
	// Create temporary database file
	tmpFile, err := os.CreateTemp("", "encrypted_db_test_*.db")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Initialize database
	db, err := database.NewDB(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to create database: %v", err)
	}
	defer db.Close()

	// Initialize schema
	if err := db.Init(); err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	// Test CreateUser
	user := models.NewUser("John Doe", "john@example.com", "1234567890", "123 Main St")
	if err := db.CreateUser(user, testPassword); err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	if user.ID == 0 {
		t.Error("User ID should be set after creation")
	}

	// Test GetUser
	retrievedUser, err := db.GetUser(user.ID, testPassword)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if retrievedUser.Name != user.Name || retrievedUser.Email != user.Email {
		t.Error("Retrieved user data does not match")
	}

	// Test UpdateUser
	user.Name = "Jane Doe"
	user.Email = "jane@example.com"
	if err := db.UpdateUser(user, testPassword); err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	updatedUser, err := db.GetUser(user.ID, testPassword)
	if err != nil {
		t.Fatalf("Failed to get updated user: %v", err)
	}

	if updatedUser.Name != "Jane Doe" || updatedUser.Email != "jane@example.com" {
		t.Error("User was not updated correctly")
	}

	// Test ListUsers
	users, err := db.ListUsers(testPassword)
	if err != nil {
		t.Fatalf("Failed to list users: %v", err)
	}

	if len(users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(users))
	}

	// Test DeleteUser
	if err := db.DeleteUser(user.ID); err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	// Verify deletion
	_, err = db.GetUser(user.ID, testPassword)
	if err == nil {
		t.Error("User should not exist after deletion")
	}
}

func TestEncryptionDecryption(t *testing.T) {
	// Create temporary database file
	tmpFile, err := os.CreateTemp("", "encrypted_db_test_*.db")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Initialize database
	db, err := database.NewDB(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to create database: %v", err)
	}
	defer db.Close()

	if err := db.Init(); err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	// Create user with sensitive data
	user := models.NewUser("Test User", "test@example.com", "0987654321", "456 Elm St")
	if err := db.CreateUser(user, testPassword); err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Retrieve and verify data is decrypted correctly
	retrievedUser, err := db.GetUser(user.ID, testPassword)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if retrievedUser.Email != "test@example.com" || retrievedUser.Phone != "0987654321" || retrievedUser.Address != "456 Elm St" {
		t.Error("Sensitive data was not decrypted correctly")
	}

	// Test with wrong password
	_, err = db.GetUser(user.ID, "wrongpassword")
	if err == nil {
		t.Error("Should not be able to decrypt with wrong password")
	}
}