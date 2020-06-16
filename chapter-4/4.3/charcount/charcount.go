// Charcount computes counts of character catagories
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if unicode.IsLetter(r) {
			counts["Letter"]++
		} else if unicode.IsDigit(r) {
			counts["Digit"]++
		} else if unicode.IsPunct(r) {
			counts["Punct"]++
		} else {
			counts["Other"]++
		}
	}

	fmt.Printf("\nCounts\n\nType:\tCount:\n")
	for key, value := range counts {
		fmt.Printf("%s\t%d\n", key, value)
	}
}
