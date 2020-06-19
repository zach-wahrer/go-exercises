// Join is a reimplementation of strings.Join
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(Join(".", "test", "this", "joiner"))
	fmt.Println(Join("-", "single"))
	fmt.Println(Join("-"))
}

func Join(sep string, strings ...string) string {
	if len(strings) == 0 {
		return ""
	}
	if len(strings) == 1 {
		return strings[0]
	}

	var buffer bytes.Buffer
	for i, value := range strings {
		buffer.Write([]byte(value))
		if i != len(strings)-1 {
			buffer.Write([]byte(sep))
		}
	}

	return buffer.String()
}
