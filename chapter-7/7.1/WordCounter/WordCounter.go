// WordCounter counts the words from Stdin and reports back to the user
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Input text for word count: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Printf("\nTotal words: %d\n", WordCounter(text))
}

func WordCounter(data string) (count int) {
	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Print("reading input:", err)
	}
	return count
}
