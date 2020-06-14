// comma inserts commas in a non-negative decimal integer string.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567890000"))
}

func comma(s string) string {
	reversed := reverseString(s)
	var buffer bytes.Buffer
	for i := len(reversed) - 1; i >= 0; i-- {
		buffer.WriteByte(byte(reversed[i]))
		if i > 0 && i%3 == 0 {
			buffer.WriteByte(',')
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
