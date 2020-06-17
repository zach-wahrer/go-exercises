// MovieData searches for information about a given movie from the website OMDb
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	movie, err := Search(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(movie)
}
