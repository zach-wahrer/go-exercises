// HTMLMapper parses an HTML file given as an argument and counts the number of
// each type of element it contains.
package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		errorOut(errors.New("You must provide a single HTML file via argument."))
	}
	file := openFile(os.Args[1])
	head := parseHTML(file)
	file.Close()
	fmt.Println(head.FirstChild.NextSibling.Type)
}

func parseHTML(file *os.File) *html.Node {
	node, err := html.Parse(file)
	if err != nil {
		errorOut(err)
	}
	return node
}

func openFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		errorOut(err)
	}
	return f
}

func errorOut(err error) {
	log.Fatalf("HTMLMapper error: %v", err)
}
