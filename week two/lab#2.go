package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a, b int

	// Function to ensure valid integer input
	getIntInput := func(prompt string) int {
		for {
			// Prompt the user
			fmt.Print(prompt)
			
			// Read input as a string to handle non-integer cases
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			// Try to convert the string to an integer
			parsedValue, err := strconv.Atoi(input)
			if err != nil {
				// If conversion fails, print an error message
				fmt.Println("Invalid input. Please enter a valid integer.")
			} else {
				// Return the parsed integer
				return parsedValue
			}
		}
	}

	// Get valid inputs for a and b
	a = getIntInput("Enter value for a: ")
	b = getIntInput("Enter value for b: ")

	// 1. Check if both numbers are positive
	if a > 0 && b > 0 {
		fmt.Println("Both numbers input are positive")
	} else if a > 0 && b < 0 {
		fmt.Println("b value is negative")
	} else if a < 0 && b > 0 {
		fmt.Println("a value is negative")
	} else {
		fmt.Println("Both numbers are negative")
	}

	// 2. Check to find the greater number
	if a > b {
		fmt.Println("A is greater than B:", a)
	} else if a < b {
		fmt.Println("B is greater than A:", b)
	} else {
		fmt.Println("Both numbers are equal:", a)
	}

	// 3. Check if the numbers are equal or not
	if a != b {
		fmt.Println("Both numbers are equal")
	} else {
		fmt.Println("Both numbers are not equal")
	}
}