package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"encrypted-db/internal/encryption"
	"encrypted-db/internal/models"

	_ "modernc.org/sqlite"
)

// DB represents the database connection
type DB struct {
	conn   *sql.DB
	dbPath string
}

// NewDB creates a new database connection
func NewDB(dbPath string) (*DB, error) {
	// Ensure directory exists
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db := &DB{
		conn:   conn,
		dbPath: dbPath,
	}

	return db, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.conn.Close()
}

// Init initializes the database schema
func (db *DB) Init() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email_encrypted TEXT NOT NULL,
		phone_encrypted TEXT,
		address_encrypted TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.conn.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	return nil
}

// ListEncryptedUsers retrieves all users with encrypted data
func (db *DB) ListEncryptedUsers() ([]map[string]interface{}, error) {
	query := `
	SELECT id, name, email_encrypted, phone_encrypted, address_encrypted, created_at, updated_at
	FROM users ORDER BY created_at DESC
	`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []map[string]interface{}

	for rows.Next() {
		var id int
		var name, emailEncrypted, phoneEncrypted, addressEncrypted string
		var createdAt, updatedAt time.Time

		err := rows.Scan(&id, &name, &emailEncrypted, &phoneEncrypted, &addressEncrypted, &createdAt, &updatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		user := map[string]interface{}{
			"id":                 id,
			"name":               name,
			"email_encrypted":    emailEncrypted,
			"phone_encrypted":    phoneEncrypted,
			"address_encrypted":  addressEncrypted,
			"created_at":         createdAt,
			"updated_at":         updatedAt,
		}

		users = append(users, user)
	}

	return users, nil
}

// CreateUser creates a new user in the database
func (db *DB) CreateUser(user *models.User, masterPassword string) error {
	// Check if email already exists by comparing with all existing users
	existingUsers, err := db.ListUsers(masterPassword)
	if err != nil {
		return fmt.Errorf("failed to check email uniqueness: %w", err)
	}

	for _, existingUser := range existingUsers {
		if existingUser.Email == user.Email {
			return fmt.Errorf("email already exists")
		}
	}

	// Encrypt sensitive fields
	emailEncrypted, err := encryption.Encrypt(user.Email, masterPassword)
	if err != nil {
		return fmt.Errorf("failed to encrypt email: %w", err)
	}

	phoneEncrypted, err := encryption.Encrypt(user.Phone, masterPassword)
	if err != nil {
		return fmt.Errorf("failed to encrypt phone: %w", err)
	}

	addressEncrypted, err := encryption.Encrypt(user.Address, masterPassword)
	if err != nil {
		return fmt.Errorf("failed to encrypt address: %w", err)
	}

	query := `
	INSERT INTO users (name, email_encrypted, phone_encrypted, address_encrypted, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := db.conn.Exec(query, user.Name, emailEncrypted, phoneEncrypted, addressEncrypted, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get inserted user ID: %w", err)
	}

	user.ID = int(id)
	return nil
}

// GetUser retrieves a user by ID
func (db *DB) GetUser(id int, masterPassword string) (*models.User, error) {
	query := `
	SELECT id, name, email_encrypted, phone_encrypted, address_encrypted, created_at, updated_at
	FROM users WHERE id = ?
	`

	row := db.conn.QueryRow(query, id)

	var user models.User
	var emailEncrypted, phoneEncrypted, addressEncrypted string

	err := row.Scan(&user.ID, &user.Name, &emailEncrypted, &phoneEncrypted, &addressEncrypted, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to scan user: %w", err)
	}

	// Decrypt sensitive fields
	user.Email, err = encryption.Decrypt(emailEncrypted, masterPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt email: %w", err)
	}

	user.Phone, err = encryption.Decrypt(phoneEncrypted, masterPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt phone: %w", err)
	}

	user.Address, err = encryption.Decrypt(addressEncrypted, masterPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt address: %w", err)
	}

	return &user, nil
}

// ListUsers retrieves all users
func (db *DB) ListUsers(masterPassword string) ([]*models.User, error) {
	query := `
	SELECT id, name, email_encrypted, phone_encrypted, address_encrypted, created_at, updated_at
	FROM users ORDER BY created_at DESC
	`

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		var user models.User
		var emailEncrypted, phoneEncrypted, addressEncrypted string

		err := rows.Scan(&user.ID, &user.Name, &emailEncrypted, &phoneEncrypted, &addressEncrypted, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		// Decrypt sensitive fields
		user.Email, err = encryption.Decrypt(emailEncrypted, masterPassword)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt email: %w", err)
		}

		user.Phone, err = encryption.Decrypt(phoneEncrypted, masterPassword)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt phone: %w", err)
		}

		user.Address, err = encryption.Decrypt(addressEncrypted, masterPassword)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt address: %w", err)
		}

		users = append(users, &user)
	}

	return users, nil
}

// UpdateUser updates a user
func (db *DB) UpdateUser(user *models.User, masterPassword string) error {
	// Check if email already exists for another user
	var count int
	emailCheckQuery := `SELECT COUNT(*) FROM users WHERE email_encrypted = ? AND id != ?`
	emailEncrypted, err := encryption.Encrypt(user.Email, masterPassword)
	if err != nil {
		return fmt.Errorf("failed to encrypt email for check: %w", err)
	}

	err = db.conn.QueryRow(emailCheckQuery, emailEncrypted, user.ID).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check email uniqueness: %w", err)
	}
	if count > 0 {
		return fmt.Errorf("email already exists for another user")
	}

	// Encrypt sensitive fields
	phoneEncrypted, err := encryption.Encrypt(user.Phone, masterPassword)
	if err != nil {
		return fmt.Errorf("failed to encrypt phone: %w", err)
	}

	addressEncrypted, err := encryption.Encrypt(user.Address, masterPassword)
	if err != nil {
		return fmt.Errorf("failed to encrypt address: %w", err)
	}

	query := `
	UPDATE users
	SET name = ?, email_encrypted = ?, phone_encrypted = ?, address_encrypted = ?, updated_at = ?
	WHERE id = ?
	`

	user.UpdatedAt = time.Now()
	_, err = db.conn.Exec(query, user.Name, emailEncrypted, phoneEncrypted, addressEncrypted, user.UpdatedAt, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// DeleteUser deletes a user
func (db *DB) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := db.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", id)
	}

	return nil
}