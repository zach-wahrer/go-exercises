// MovieData searches for information about a given movie from the website OMDb
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	movie, err := Search(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\t| %s | %s | %s | %s\n",
		movie.Title, movie.Year, movie.Rated, movie.Runtime, movie.Genre)

	fmt.Printf("Download poster (y/n): ")

	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadRune()

	if err != nil {
		log.Fatal(err)
	}

	if string(input) == "y" {

	}
}
