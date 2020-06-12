package main

import (
	"fmt"
	"tempconv"
)

func main() {
	current := tempconv.Fahrenheit(65)
	fmt.Println(current)
	fmt.Println(tempconv.FToK(current))
	fmt.Println(tempconv.FToC(current))
}
