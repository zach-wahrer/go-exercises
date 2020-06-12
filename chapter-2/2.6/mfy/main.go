package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		for _, input := range os.Args[1:] {
			if number, err := strconv.ParseFloat(input, 32); err == nil {
				printConversions(number)
			} else {
				fmt.Println("Error: Arguments must only be numbers.")
				os.Exit(0)
			}
		}

	} else {
		reader := bufio.NewReader(os.Stdin)
		for true {
			fmt.Printf("Enter number (anything else exits): ")

			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error: %v", err)
				os.Exit(0)
			}

			input = strings.Replace(input, "\n", "", -1)

			if number, err := strconv.ParseFloat(input, 32); err == nil {
				printConversions(number)
			} else {
				os.Exit(0)
			}
		}
	}
}

func printConversions(number float64) {
	fmt.Printf("%v = %v\t %v = %v\n",
		Meters(number), MToF(Meters(number)), Feet(number), FToM(Feet(number)))
}
