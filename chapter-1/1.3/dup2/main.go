// Dup2 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	locations := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin", locations)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(file, counts, arg, locations)
			file.Close()
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\t%s\n", count, line, locations[line])
		}
	}
}

func countLines(file *os.File, counts map[string]int, fileName string, locations map[string][]string) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
		locations[input.Text()] = append(locations[input.Text()], fileName)
	}
}
