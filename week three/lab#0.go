package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"

	"golang.org/x/crypto/sha3"
)

func main() {
	
	fmt.Println("=============== Name + Hashing Program =================")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please input value 1: ")
	scanner.Scan()
	val1 := scanner.Text()

	fmt.Print("Please input value 2: ")
	scanner.Scan()
	val2 := scanner.Text()

	// MD5
	h1 := md5.Sum([]byte(val1))
	h2 := md5.Sum([]byte(val2))
	fmt.Printf("Hash (MD5):\nOutput A = %x\nOutput B = %x\n", h1, h2)
	if fmt.Sprintf("%x", h1) == fmt.Sprintf("%x", h2) {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}

	// SHA1
	h1_sha1 := sha1.Sum([]byte(val1))
	h2_sha1 := sha1.Sum([]byte(val2))
	fmt.Printf("Hash (SHA1):\nOutput A = %x\nOutput B = %x\n", h1_sha1, h2_sha1)
	if fmt.Sprintf("%x", h1_sha1) == fmt.Sprintf("%x", h2_sha1) {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}

	// SHA256
	h1_sha256 := sha256.Sum256([]byte(val1))
	h2_sha256 := sha256.Sum256([]byte(val2))
	fmt.Printf("Hash (SHA256):\nOutput A = %x\nOutput B = %x\n", h1_sha256, h2_sha256)
	if fmt.Sprintf("%x", h1_sha256) == fmt.Sprintf("%x", h2_sha256) {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}

	// SHA512
	h1_sha512 := sha512.Sum512([]byte(val1))
	h2_sha512 := sha512.Sum512([]byte(val2))
	fmt.Printf("Hash (SHA512):\nOutput A = %x\nOutput B = %x\n", h1_sha512, h2_sha512)
	if fmt.Sprintf("%x", h1_sha512) == fmt.Sprintf("%x", h2_sha512) {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}

	// SHA3-256
	h1_sha3 := sha3.Sum256([]byte(val1))
	h2_sha3 := sha3.Sum256([]byte(val2))
	fmt.Printf("Hash (SHA3-256):\nOutput A = %x\nOutput B = %x\n", h1_sha3, h2_sha3)
	if fmt.Sprintf("%x", h1_sha3) == fmt.Sprintf("%x", h2_sha3) {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}
}
