package main

import "fmt"

func main() {
	var a, b int

	// Take input from user
	fmt.Print("Enter value for a: ")
	fmt.Scanln(&a)
	fmt.Print("Enter value for b: ")
	fmt.Scanln(&b)

	fmt.Println("\nTask1 Operators Demonstration:")

	// = operator
	x := a
	fmt.Printf("x = a -> x = %d\n", x)

	// += operator
	x += b
	fmt.Printf("x += b -> x = %d\n", x)

	// -= operator
	x -= b
	fmt.Printf("x -= b -> x = %d\n", x)

	// *= operator
	x *= b
	fmt.Printf("x *= b -> x = %d\n", x)

	// /= operator (avoid division by zero)
	if b != 0 {
		x /= b
		fmt.Printf("x /= b -> x = %d\n", x)
	} else {
		fmt.Println("Division by zero skipped for x /= b")
	}

	// %= operator (avoid modulo by zero)
	if b != 0 {
		x %= b
		fmt.Printf("x %%= b -> x = %d\n", x)
	} else {
		fmt.Println("Modulo by zero skipped for x %= b")
	}
}