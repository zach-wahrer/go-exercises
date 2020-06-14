// Anagram checks if two strings are anagrams of each other
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	word1, word2 := "Golang", "Langgo"
	fmt.Println(word1+" is an anagram of "+word2+": ", anagram(word1, word2))

}

func anagram(word1, word2 string) bool {
	w1 := strings.Split(strings.ToLower(word1), "")
	w2 := strings.Split(strings.ToLower(word2), "")

	if len(w1) != len(w2) {
		return false
	}

	sort.Strings(w1)
	sort.Strings(w2)

	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			return false
		}
	}

	return true

}
