package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

// Encrypt encrypts plaintext using AES-GCM with a key derived from password
func Encrypt(plaintext, password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	// Derive key from password using PBKDF2
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	key := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Generate nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt and authenticate
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// Combine salt + ciphertext
	combined := append(salt, ciphertext...)

	return hex.EncodeToString(combined), nil
}

// Decrypt decrypts ciphertext using AES-GCM with a key derived from password
func Decrypt(ciphertextHex, password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	// Decode hex
	combined, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return "", err
	}

	if len(combined) < 32 {
		return "", errors.New("invalid ciphertext")
	}

	// Extract salt and ciphertext
	salt := combined[:32]
	ciphertext := combined[32:]

	// Derive key from password
	key := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Extract nonce
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce := ciphertext[:nonceSize]
	ciphertext = ciphertext[nonceSize:]

	// Decrypt and verify
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
