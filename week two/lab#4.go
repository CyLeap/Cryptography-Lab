package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/fatih/color"
)

// Functions for each operation
func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("🚫 Division by zero error!")
	}
	return a / b, nil
}

func mod(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("🚫 Modulo by zero error!")
	}
	return a % b, nil
}

// Function to get integer input safely
func getIntInput(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("❌ Invalid input! Please enter a valid integer.")
			continue
		}
		return num
	}
}

// Main calculator menu
func main() {
	// Define a colorful header
	color.Set(color.FgCyan, color.Bold)
	fmt.Println("===== ⚡️ Mini Calculator ⚡️ =====")
	color.Unset()

	reader := bufio.NewReader(os.Stdin)

	// Loop for menu-driven calculator
	for {
		// Print menu options
		color.Set(color.FgYellow)
		fmt.Println("\n1) Add ")
		fmt.Println("2) Subtract ")
		fmt.Println("3) Multiply ")
		fmt.Println("4) Divide")
		fmt.Println("5) Modulo")
		fmt.Println("6) Exit 🚪")
		color.Unset()

		// Read user input for menu choice
		fmt.Print("Choose an option: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil || choice < 1 || choice > 6 {
			color.Set(color.FgRed)
			fmt.Println("❌ Invalid choice! Please enter a number between 1-6.")
			color.Unset()
			continue
		}

		// Exit condition
		if choice == 6 {
			color.Set(color.FgGreen)
			fmt.Println("💀 Exiting... Goodbye! 💀")
			color.Unset()
			break
		}

		// Get numbers from the user
		a := getIntInput("Enter a: ")
		b := getIntInput("Enter b: ")

		// Perform the operation based on the choice
		switch choice {
		case 1:
			color.Set(color.FgMagenta)
			fmt.Printf("🔵 %d + %d = %d\n", a, b, add(a, b))
			color.Unset()
		case 2:
			color.Set(color.FgMagenta)
			fmt.Printf("🔴 %d - %d = %d\n", a, b, sub(a, b))
			color.Unset()
		case 3:
			color.Set(color.FgMagenta)
			fmt.Printf("🟢 %d * %d = %d\n", a, b, mul(a, b))
			color.Unset()
		case 4:
			result, err := div(a, b)
			if err != nil {
				color.Set(color.FgRed)
				fmt.Println(err)
				color.Unset()
			} else {
				color.Set(color.FgBlue)
				fmt.Printf("🔽 %d / %d = %d\n", a, b, result)
				color.Unset()
			}
		case 5:
			result, err := mod(a, b)
			if err != nil {
				color.Set(color.FgRed)
				fmt.Println(err)
				color.Unset()
			} else {
				color.Set(color.FgBlue)
				fmt.Printf("💥 %d %% %d = %d\n", a, b, result)
				color.Unset()
			}
		}
	}
}