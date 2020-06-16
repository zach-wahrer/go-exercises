// Wordfreq counts the frequency of words within an input text file
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"unicode"
)

func main() {
	if len(os.Args) != 2 {
		errorOut(errors.New("A single file name argument is required."))
	}

	wordCounts := count(os.Args[1])
	wordFreq := convertToFreq(wordCounts)
	printWordsByFreq(wordFreq)

}

func count(fileName string) map[string]int {
	file, err := os.Open(fileName)

	if err != nil {
		errorOut(err)
	}

	words := make(map[string]int)
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words[input.Text()]++
	}
	file.Close()
	return words
}

func convertToFreq(words map[string]int) map[int][]string {
	wordsByFreq := make(map[int][]string)
	for word, count := range words {
		if unicode.IsPunct(rune(word[len(word)-1])) && len(word) > 1 {
			word = word[:len(word)-2]
		}
		if len(word) > 0 && unicode.IsUpper(rune(word[0])) {
			word = string(unicode.ToLower(rune(word[0]))) + word[1:]
		}
		wordsByFreq[count] = append(wordsByFreq[count], word)
	}
	return wordsByFreq
}

func createSortKey(freq map[int][]string) []int {
	list := make([]int, 0, len(freq))
	for key := range freq {
		list = append(list, key)
	}
	sort.Ints(list)
	return list
}

func printWordsByFreq(words map[int][]string) {
	fmt.Printf("\nWord\t\t\tCount\n")
	for _, key := range createSortKey(words) {
		for _, word := range words[key] {
			fmt.Printf("%v", word)
			if len(word) < 8 {
				fmt.Printf("\t\t\t")
			} else {
				fmt.Printf("\t\t")
			}
			fmt.Printf("%d\n", key)
		}
	}
}

func errorOut(err error) {
	fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
	os.Exit(1)
}
