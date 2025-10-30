package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Function to convert string to binary representation
func toBinary(input string) string {
	var result string
	for _, c := range input {
		result += fmt.Sprintf("%08b ", c) // 8-bit binary for each character
	}
	return result
}

// Function to convert string to hexadecimal representation
func toHex(input string) string {
	return hex.EncodeToString([]byte(input))
}

// Function to convert string to Base64 representation
func toBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func main() {
	var input string

	// Get input from the user
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	// Clean input to remove unwanted newlines/spaces
	input = strings.TrimSpace(input)

	// Convert and print results
	fmt.Println("\n===== Conversions =====")
	fmt.Printf("Binary: %s\n", toBinary(input))
	fmt.Printf("Hexadecimal: %s\n", toHex(input))
	fmt.Printf("Base64: %s\n", toBase64(input))

	// Optional: Save to file
	fmt.Print("\nWould you like to save the results to a file? (y/n): ")
	var saveFile string
	fmt.Scanln(&saveFile)
	if strings.ToLower(saveFile) == "y" {
		file, err := os.Create("encoding_results.txt")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("Input: %s\n\nBinary:\n%s\nHexadecimal:\n%s\nBase64:\n%s\n", input, toBinary(input), toHex(input), toBase64(input)))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("Results saved to encoding_results.txt")
	} else {
		fmt.Println("Results not saved.")
	}
}
