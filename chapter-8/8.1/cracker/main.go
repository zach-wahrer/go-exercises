package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz1234567890"

func main() {
	go waiting(100 * time.Millisecond)
	testPass := encrypter("700a0a")
	targetLen := 6

	chunkSize := int(len(alphabet) / 6)
	start, end := 0, chunkSize
	for i := 0; i < chunkSize; i++ {
		if i == chunkSize-1 {
			// go cracker(start, len(alphabet), length testPass)
		} else {
			go cracker(start, end, targetLen, testPass)
		}
		start, end = end, end+chunkSize
	}
	cracker(30, len(alphabet), targetLen, testPass)
}

func cracker(start, finish, targetLen int, target [16]byte) {
	for _, char1 := range alphabet[start:finish] {
		if passEqual(string(char1), target) {
			fmt.Printf("\tPassword found: %s", string(char1))
		} else {
			recursiveCrack(string(char1), 2, targetLen, target)
		}
	}
}

func recursiveCrack(chars string, currLen, targetLen int, target [16]byte) {
	if currLen <= targetLen {
		for _, char := range alphabet {
			if passEqual(chars+string(char), target) {
				fmt.Printf("\tPassword found: %s", chars+string(char))
				return
			}

			recursiveCrack(chars+string(char), currLen+1, targetLen, target)

		}
	}
	return

}

func passEqual(pass string, target [16]byte) bool {
	return encrypter(pass) == target
}

func encrypter(s string) [16]byte {
	data := []byte(s)
	return md5.Sum(data)
}

func waiting(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
