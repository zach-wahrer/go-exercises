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

	if err := output.Execute(os.Stdout, movie); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Download poster (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadRune()

	if err != nil {
		log.Fatal(err)
	}

	if string(input) == "y" {
		if err := Download(movie.Title, movie.Poster); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Poster downloaded successfully.")
	}
}
