// Converter converts numbers to different bases and prints them
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(converter(10, 8))
}

func converter(number, base int) string {
	return strconv.FormatInt(int64(number), base)
}
