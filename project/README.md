# Encrypted Database with AES Encryption

A secure command-line application that implements encrypted data storage using AES-GCM encryption with PBKDF2 key derivation. Perfect for learning cryptography concepts and building secure database applications.

## 📋 Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Building the Application](#building-the-application)
- [Usage](#usage)
- [Security Features](#security-features)
- [Project Structure](#project-structure)
- [Testing](#testing)
- [Validation Rules](#validation-rules)
- [Contributing](#contributing)
- [License](#license)

## ✨ Features

- 🔐 **AES-256-GCM Encryption**: Authenticated encryption for sensitive data
- 🗄️ **SQLite Database**: Lightweight, file-based database with encrypted storage
- 🔑 **PBKDF2 Key Derivation**: Secure password-based key generation with salt
- 🛡️ **Data-at-Rest Security**: All sensitive fields encrypted on disk
- 🖥️ **CLI Interface**: User-friendly command-line operations
- ✅ **Data Validation**: Comprehensive input validation and error handling
- 🔒 **Password Protection**: Master password required for all operations

## 📋 Prerequisites

### System Requirements

- Go 1.19 or later
- Git (for cloning the repository)

### Go Installation

#### Windows

1. Download the MSI installer from [golang.org/dl](https://golang.org/dl)
2. Run the installer and follow the setup wizard
3. Verify installation: Open Command Prompt and run `go version`

#### macOS

1. Download the macOS package from [golang.org/dl](https://golang.org/dl)
2. Open the downloaded .pkg file and follow the installation wizard
3. Verify installation: Open Terminal and run `go version`

#### Linux (Ubuntu/Debian)

```bash
# Update package list
sudo apt update

# Install Go
sudo apt install golang-go

# Verify installation
go version
```

#### Linux (CentOS/RHEL/Fedora)

```bash
# Install Go
sudo dnf install golang

# Or for older systems
sudo yum install golang

# Verify installation
go version
```

## 🚀 Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/CyLeap/Cryptography-Lab.git
   cd project
   ```

2. **Download dependencies:**

   ```bash
   go mod download
   ```

3. **Verify dependencies:**
   ```bash
   go mod tidy
   ```

## 🔨 Building the Application

### Windows

```bash
cd project
go build ./cmd/encrypted-db
```

This creates `encrypted-db.exe`

### macOS/Linux

```bash
cd project
go build ./cmd/encrypted-db
```

This creates `encrypted-db` executable

### Cross-Platform Building

```bash
# Build for Windows (from Linux/Mac)
GOOS=windows GOARCH=amd64 go build ./cmd/encrypted-db

# Build for macOS (from Linux/Windows)
GOOS=darwin GOARCH=amd64 go build ./cmd/encrypted-db

# Build for Linux (from Windows/Mac)
GOOS=linux GOARCH=amd64 go build ./cmd/encrypted-db
```

## 📖 Usage

### Basic Commands

1. **Initialize Database:**

   ```bash
   ./encrypted-db -action=init
   ```

   Creates the SQLite database with required tables.

2. **Create User:**

   ```bash
   ./encrypted-db -action=create -name="John Doe" -email="john@gmail.com" -phone="1234567890" -address="123 Main St"
   ```

3. **List All Users:**

   ```bash
   ./encrypted-db -action=list
   ```

4. **Get Specific User:**

   ```bash
   ./encrypted-db -action=get -id=1
   ```

5. **Update User:**

   ```bash
   ./encrypted-db -action=update -id=1 -email="john.doe@gmail.com"
   ```

6. **Delete User:**
   ```bash
   ./encrypted-db -action=delete -id=1
   ```

### Command Line Options

| Flag        | Description                                                 | Required                               |
| ----------- | ----------------------------------------------------------- | -------------------------------------- |
| `-action`   | Action to perform (init, create, list, get, update, delete) | Yes                                    |
| `-name`     | User name                                                   | For create/update                      |
| `-email`    | User email (must be @gmail.com)                             | For create/update                      |
| `-phone`    | User phone (10-15 digits only)                              | Optional                               |
| `-address`  | User address (non-empty)                                    | Optional                               |
| `-id`       | User ID                                                     | For get/update/delete                  |
| `-password` | Master password                                             | Optional (will prompt if not provided) |
| `-db`       | Database file path                                          | Optional (default: encrypted.db)       |

### Example Workflow

```bash
# Initialize database
./encrypted-db -action=init

# Create users
./encrypted-db -action=create -name="Alice Johnson" -email="alice@gmail.com" -phone="1234567890" -address="123 Main St"
./encrypted-db -action=create -name="Bob Smith" -email="bob@gmail.com" -phone="0987654321" -address="456 Oak Ave"

# List all users
./encrypted-db -action=list

# Update user
./encrypted-db -action=update -id=1 -phone="1112223333"

# Delete user
./encrypted-db -action=delete -id=2
```

## 🔒 Security Features

- **AES-256-GCM**: Authenticated encryption providing confidentiality and integrity
- **PBKDF2**: Password-based key derivation with 100,000 iterations
- **Random Salts**: 32-byte salts prevent rainbow table attacks
- **Unique Nonces**: Per-encryption nonces prevent pattern analysis
- **Master Password**: Required for all data access operations
- **Encrypted Fields**: Email, phone, and address are encrypted at rest
- **Plain Fields**: ID, name, and timestamps remain unencrypted for indexing

## 📁 Project Structure

```
project/
├── cmd/
│   └── encrypted-db/
│       └── main.go              # CLI entry point and command parsing
├── internal/
│   ├── database/
│   │   └── db.go                # Database operations and CRUD functions
│   ├── encryption/
│   │   └── aes.go               # AES-GCM encryption/decryption logic
│   ├── models/
│   │   └── user.go              # User data model and validation
│   └── utils/
│       └── helpers.go           # Utility functions and password handling
├── tests/
│   └── database_test.go         # Unit tests for database operations
├── go.mod                       # Go module dependencies
├── go.sum                       # Dependency checksums
└── README.md                    # This documentation
```

## 🧪 Testing

### Run Unit Tests

```bash
cd project
go test ./tests/...
```

### Run with Verbose Output

```bash
go test -v ./tests/...
```

### Run Specific Test

```bash
go test -run TestCreateUser ./tests/
```

## ✅ Validation Rules

### Email

- Must end with `@gmail.com` (case-insensitive)
- Must be unique across all users
- Cannot be empty

### Phone

- Must contain only digits (0-9)
- Length: 10-15 characters
- Can be empty

### Address

- Cannot be empty or whitespace-only
- No specific format requirements

### Name

- Length: 2-50 characters
- Cannot be empty

### Password

- Length: 8-12 characters
- Must contain:
  - At least one uppercase letter (A-Z)
  - At least one lowercase letter (a-z)
  - At least one number (0-9)
  - At least one special character (!@#$%^&\*(),.?":{}|<>)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

If you encounter any issues or have questions:

1. Check the [Issues] (https://github.com/CyLeap/Cryptography-Lab.git) page
2. Create a new issue with detailed information (telegram: @bunleappp, gmail: bunleapthay@gmail.com)
3. Include your operating system, Go version, and error messages

## 🎯 Learning Outcomes

This project demonstrates:

- Cryptographic principles in practice
- Secure software development practices
- Database design with encryption
- Command-line interface design
- Input validation and error handling
- Go programming best practices
