// HTMLMapper parses an HTML file given as an argument and counts the number of
// each type of element it contains.
package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		errorOut(errors.New("You must provide a single HTML file via argument."))
	}
	file := openFile(os.Args[1])
	head := parseHTML(file)
	file.Close()
	elements := make(map[string]int)
	traverseNode(elements, head)
	printOutput(elements)
}

func traverseNode(elements map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseNode(elements, c)
	}
}

func parseHTML(file *os.File) *html.Node {
	node, err := html.Parse(file)
	if err != nil {
		errorOut(err)
	}
	return node
}

func printOutput(elements map[string]int) {
	longestLen := getLongestLen(elements)
	totalElements := 0
	fmt.Printf("Element Count:\n")
	for _, name := range sortedKeys(elements) {
		fmt.Printf("%s%*s %d\n", name, longestLen-len(name), "", elements[name])
		totalElements += elements[name]
	}
	fmt.Printf("Total Elements: %*s%d\n", int(math.Max(float64(longestLen-15), 0)), "", totalElements)

}

func sortedKeys(elements map[string]int) (sortedKeys []string) {
	for element := range elements {
		sortedKeys = append(sortedKeys, element)
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}

func getLongestLen(elements map[string]int) (length int) {
	for element := range elements {
		if len(element) > length {
			length = len(element)
		}
	}
	return length
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
