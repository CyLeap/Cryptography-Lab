package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"os"
	"time"

	"cryptography-lab/utils/crack"
)

func main() {
	// CLI flags
	algo := flag.String("algo", "md5", "hash algorithm to use: md5 or sha1")
	target := flag.String("hash", "6a85dfd77d9cb35770c9dc6728d73d3f", "target hash to crack")
	wordlist := flag.String("wordlist", "nord_vpn.txt", "path to wordlist file")
	verbose := flag.String("verbose", "verbose.txt", "path to verbose output file")
	flag.Parse()

	// Create/overwrite verbose output file
	vfile, err := os.Create(*verbose)
	if err != nil {
		fmt.Println("Error creating verbose output file:", err)
		return
	}
	defer vfile.Close()
	vwriter := bufio.NewWriter(vfile)
	defer vwriter.Flush()

	file, err := os.Open(*wordlist)
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
		fmt.Print(logLine)
		vwriter.WriteString(logLine)

		match := false
		switch *algo {
		case "md5":
			// reuse md5 quick compute for speed
			sum := md5.Sum([]byte(password))
			if fmt.Sprintf("%x", sum) == *target {
				match = true
			}
		case "sha1":
			if crack.CheckSHA1(password, *target) {
				match = true
			}
		default:
			msg := fmt.Sprintf("Unsupported algorithm: %s\n", *algo)
			fmt.Print(msg)
			vwriter.WriteString(msg)
			return
		}

		if match {
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
