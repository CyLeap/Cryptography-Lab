package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"cryptography-lab/utils/crack"
)

func main() {
	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"

	// Open verbose output file (overwrite existing)
	vfile, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating verbose output file:", err)
		return
	}
	defer vfile.Close()
	vwriter := bufio.NewWriter(vfile)
	defer vwriter.Flush()

	file, err := os.Open("nord_vpn.txt")
	if err != nil {
		msg := fmt.Sprintf("Error opening wordlist file: %v\n", err)
		fmt.Print(msg)
		vwriter.WriteString(msg)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	found := false
	for scanner.Scan() {
		password := scanner.Text()
		t := time.Now().Format(time.RFC3339)
		logLine := fmt.Sprintf("[%s] Checking: %s\n", t, password)
		// print to stdout and append to verbose file
		fmt.Print(logLine)
		vwriter.WriteString(logLine)

		if crack.CheckMD5(password, targetHash) {
			foundLine := fmt.Sprintf("[%s] Found: %s\n", t, password)
			fmt.Print(foundLine)
			vwriter.WriteString(foundLine)
			vwriter.Flush()
			found = true
			break
		}
	}

	if err := scanner.Err(); err != nil {
		msg := fmt.Sprintf("Error reading wordlist: %v\n", err)
		fmt.Print(msg)
		vwriter.WriteString(msg)
	}

	if !found {
		msg := fmt.Sprintf("[%s] Password not found in wordlist\n", time.Now().Format(time.RFC3339))
		fmt.Print(msg)
		vwriter.WriteString(msg)
	}
}
