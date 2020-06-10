// Echo2 is a program that prints out the command line args passed to it

package main

import (
	"fmt"
	"os"
)

func main() {
	commands, seperator := "", ""
	for _, command := range os.Args[1:] {
		commands += command + seperator
		seperator = " "
	}

	fmt.Println(commands)
}
