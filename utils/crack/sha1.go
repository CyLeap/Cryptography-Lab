package crack

import (
    "crypto/sha1"
    "fmt"
)

// CheckSHA1 computes SHA1(password) and compares it to targetHash (hex lowercase)
func CheckSHA1(password string, targetHash string) bool {
    sum := sha1.Sum([]byte(password))
    hashStr := fmt.Sprintf("%x", sum)
    return hashStr == targetHash
}
