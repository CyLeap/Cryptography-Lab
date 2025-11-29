# Encrypted Database with AES Encryption

## Project Overview

This project implements a secure SQLite database where sensitive fields are encrypted using AES encryption. It provides a simple command-line interface for performing CRUD (Create, Read, Update, Delete) operations on encrypted data, focusing on data-at-rest security.

## Key Features

- **AES Encryption**: Sensitive data fields are encrypted before storage
- **SQLite Database**: Lightweight, file-based database
- **CRUD Operations**: Full create, read, update, delete functionality
- **Data-at-Rest Security**: Data remains encrypted when stored on disk
- **CLI Interface**: Easy-to-use command-line tool

## Technical Stack

- **Language**: Go
- **Database**: SQLite3
- **Encryption**: AES-GCM (Galois/Counter Mode)
- **Key Derivation**: PBKDF2 for secure key generation

## Project Structure

```
project/
├── cmd/
│   └── encrypted-db/
│       └── main.go              # CLI entry point
├── internal/
│   ├── database/
│   │   ├── db.go                # Database connection and operations
│   │   └── migrations.go        # Database schema setup
│   ├── encryption/
│   │   └── aes.go               # AES encryption/decryption functions
│   ├── models/
│   │   └── user.go              # Data models (User struct)
│   └── utils/
│       └── helpers.go           # Utility functions
├── tests/
│   └── database_test.go         # Unit tests
├── go.mod                       # Go module file
└── README.md                    # This file
```

## Security Considerations

- Master password for key derivation
- AES-GCM for authenticated encryption
- PBKDF2 with salt for key stretching
- Encrypted fields: email, phone, address
- Plain fields: id, name, created_at

## Usage Examples

```bash
# Build the application
go build ./cmd/encrypted-db

# Initialize database
./encrypted-db init --password mymasterpassword

# Create a user
./encrypted-db create --name "John Doe" --email "john@example.com" --phone "1234567890" --address "123 Main St"

# List all users
./encrypted-db list

# Update a user
./encrypted-db update --id 1 --email "john.doe@example.com"

# Delete a user
./encrypted-db delete --id 1
```

## Next Steps

1. Set up project structure and Go module
2. Implement AES encryption functions
3. Create database layer with SQLite
4. Build data models
5. Implement CRUD operations
6. Add CLI interface
7. Write unit tests
8. Add error handling and validation
