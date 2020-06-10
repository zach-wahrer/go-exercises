// A program that prints the index and value of each of the arguments passed in
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args[1:] {
		fmt.Println(index+1, value)
	}
}
