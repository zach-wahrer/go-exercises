// comma inserts commas in a non-negative decimal integer string.
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-1346711.89812"))
}

func comma(s string) string {
	whole := s
	var fraction string
	var buffer bytes.Buffer

	if strings.Contains(s, ".") {
		sides := strings.Split(s, ".")
		whole, fraction = sides[0], sides[1]
	}

	if strings.Contains(whole, "-") {
		whole = whole[1:]
		buffer.WriteByte('-')
	}

	reversedWhole := reverseString(whole)

	for i := len(reversedWhole) - 1; i >= 0; i-- {
		buffer.WriteByte(byte(reversedWhole[i]))
		if i > 0 && i%3 == 0 {
			buffer.WriteByte(',')
		}
	}

	if fraction != "" {
		buffer.WriteByte('.')
		for i := 0; i < len(fraction); i++ {
			buffer.WriteByte(byte(fraction[i]))
		}
	}

	return buffer.String()
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
