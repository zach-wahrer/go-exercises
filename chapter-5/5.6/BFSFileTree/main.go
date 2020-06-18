// BFSFileTree prints out the local file system using Breadth First Search
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	count := breadthFirst(traverse, os.Args[1:])
	fmt.Printf("Total directories: %d\n", count)
}

func Drill(currentPath string) (paths []string, err error) {
	listing, err := ioutil.ReadDir(currentPath)
	if err != nil {
		return nil, err
	}
	for _, path := range listing {
		if path.IsDir() {
			paths = append(paths, fmt.Sprintf("%s%s/", currentPath, path.Name()))
		}
	}
	return paths, err
}

func traverse(path string) []string {
	fmt.Println(path)
	list, err := Drill(path)
	if err != nil {
		log.Print(err)
	}
	return list
}

func breadthFirst(f func(item string) []string, worklist []string) (count int) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				count++
				worklist = append(worklist, f(item)...)
			}
		}
	}
	return count
}
