package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// xorEncrypt performs XOR encryption or decryption
// since XOR operation is reversible.
func xorEncrypt(text string, key byte) string {
	var result []byte
	for i := 0; i < len(text); i++ {
		// XOR each byte of text with the key
		result = append(result, text[i]^key)
	}
	return string(result)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Step 1: Input plaintext message
	fmt.Print("Enter plaintext: ")
	plaintext, _ := reader.ReadString('\n')
	plaintext = strings.TrimSpace(plaintext)

	// Step 2: Input key (single character)
	fmt.Print("Enter a single character key: ")
	keyInput, _ := reader.ReadString('\n')
	keyInput = strings.TrimSpace(keyInput)

	// Validate that key is only one character
	if len(keyInput) != 1 {
		fmt.Println("Error: Key must be a single character.")
		return
	}
	key := keyInput[0]

	// Step 3: Encrypt plaintext
	ciphertext := xorEncrypt(plaintext, key)
	fmt.Println("\n===== XOR Encryption =====")
	fmt.Println("Ciphertext (Encrypted):", ciphertext)

	// Step 4: Decrypt ciphertext (same function)
	decrypted := xorEncrypt(ciphertext, key)
	fmt.Println("Decrypted Message:", decrypted)
}