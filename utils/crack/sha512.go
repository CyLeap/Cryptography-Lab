package crack

import (
	"crypto/sha512"
	"fmt"
)

// CheckSHA512 computes SHA512(password) and compares it to targetHash (hex lowercase)
func CheckSHA512(password string, targetHash string) bool {
	sum := sha512.Sum512([]byte(password))
	hashStr := fmt.Sprintf("%x", sum)
	return hashStr == targetHash
}