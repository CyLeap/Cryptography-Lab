package main

import (
	"bufio"
	"fmt"
	"os"

	"cryptography-lab/utils/crack"
)

func main() {
	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"

	file, err := os.Open("nord_vpn.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		password := scanner.Text()
		fmt.Printf("Checking: %s\n", password)
		if crack.CheckMD5(password, targetHash) {
			fmt.Printf("Found: %s\n", password)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
