package main

import "fmt"


func myAND(a, b int) {
	fmt.Printf("%d & %d = %d\n", a, b, a&b)
}

func myOR(a, b int){
	fmt.Printf("%d | %d = %d\n", a, b, a|b)
}

func myXOR(a, b int){
	fmt.Printf("%d ^ %d = %d\n", a, b, a^b)
}

func myNOT(a, b int){
	fmt.Printf("~%d = %d\n", a, ^a)
	fmt.Printf("~%d = %d\n", b, ^b)
}

func leftshift(a int, n uint){
	fmt.Printf("%d << %d = %d\n", a, n, a<<n)
}

func rightshift(a int, n uint){
	fmt.Printf("%d >> %d = %d\n", a, n, a>>n)
}

func showoff(a, b int){
	
	fmt.Println("\nBitwise and Assignment Opr.:")

	x := a
	fmt.Printf("x = %d\n", x)
	x += b
	fmt.Printf("x += %d -> %d\n", b, x)
	x -= b
	fmt.Printf("x -= %d -> %d\n", b, x)
	x *= b
	fmt.Printf("x *= %d -> %d\n", b, x)
	if b != 0 {
		x /= b
		fmt.Printf("x /= %d -> %d\n", b, x)
		x %= b
		fmt.Printf("x %%= %d -> %d\n", b, x)
	}
}

func main() {
	var a, b int

	// Input two integers
	fmt.Print("Enter value for a: ")
	fmt.Scanln(&a)
	fmt.Print("Enter value for b: ")
	fmt.Scanln(&b)

	fmt.Println("\nBitwise Operations:")
	myAND(a, b)
	myOR(a, b)
	myXOR(a, b)
	myNOT(a, b)

	leftshift(a, 1)
	rightshift(a, 1)

	showoff(a, b)
}