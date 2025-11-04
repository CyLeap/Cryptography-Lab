package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Decode hex string to ASCII
func hexToASCII(hexStr string) string {
	decoded := ""
	for i := 0; i < len(hexStr); i += 2 {
		s, _ := strconv.ParseUint(hexStr[i:i+2], 16, 64)
		decoded += string(s)
	}
	return decoded
}

func main() {
	// The hex values from the regex pattern
	hexValues := "6d656f77" // \x6d\x65\x6f\x77
	decoded := hexToASCII(hexValues)
	
	fmt.Printf("Decoding process:\n")
	fmt.Printf("1. Hex value from regex: \\x6d\\x65\\x6f\\x77\n")
	fmt.Printf("2. Hex without \\x: %s\n", hexValues)
	fmt.Printf("3. Decoded to ASCII: %s\n", decoded)
	fmt.Printf("4. Repeated twice (because of {2}): %s%s\n\n", decoded, decoded)
	
	// Build the flag
	flag := "cryptoCTF{" + decoded + decoded + "}"

	// The regex pattern from the challenge
	pattern := `^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$`

	// Create a new regexp
	re := regexp.MustCompile(pattern)

	// Check if the flag matches the pattern
	if re.MatchString(flag) {
		fmt.Println("Flag validation successful! ✅")
		fmt.Printf("Flag: %s\n", flag)
		
		// Explain the components
		fmt.Println("\nFlag breakdown:")
		fmt.Println("1. Prefix: cryptoCTF{")
		fmt.Println("2. Content: meow (repeated twice)")
		fmt.Println("3. Suffix: }")
		fmt.Println("\nHex decoded values:")
		fmt.Println("\\x6d\\x65\\x6f\\x77 = 'meow'")
	} else {
		fmt.Println("❌ Flag validation failed!")
	}
}