// hasher is a utility that prints the SHA256, SHA384, or SHA512 hash of an input
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input value to hash:")
	text, err := reader.ReadString('\n')
	fmt.Println(text)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "512" {
			fmt.Printf("Hash: %x\n", sha512.Sum512([]byte(text)))
		} else if os.Args[1] == "384" {
			fmt.Printf("Hash: %x\n", sha512.Sum384([]byte(text)))
		} else {
			fmt.Printf("Hash: %x\n", sha256.Sum256([]byte(text)))
		}
	} else {
		fmt.Printf("Hash: %x\n", sha256.Sum256([]byte(text)))
	}
}
