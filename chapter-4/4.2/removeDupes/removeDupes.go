// removeDupes removes adjacent duplicates in a []string slice.
package main

import "fmt"

func main() {
	strings := []string{"test", "test", "keep", "these", "but", "not", "extra", "this", "this"}
	strings = removeDupes(strings)
	fmt.Println(strings)
}

func removeDupes(strings []string) []string {
	for i := 0; i < len(strings)-1; i++ {
		if strings[i] == strings[i+1] {
			strings = append(strings[:i], strings[i+1:]...)
		}
	}
	return strings
}
