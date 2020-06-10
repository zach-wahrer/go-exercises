// Echo prints the command-line arguments given to it.
package main

import (
	"fmt"
	"os"
)

func main() {
	var string, seperator string
	for i := 0; i < len(os.Args); i++ {
		string += seperator + os.Args[i]
		seperator = " "
	}
	fmt.Println(string)
}
