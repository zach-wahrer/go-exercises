package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz1234567890"

func main() {
	go waiting(100 * time.Millisecond)

	testPass := encrypter("1aezxm")
	targetLen := 6

	out := make(chan string)

	chunkSize := int(len(alphabet) / 6)
	start, end := 0, chunkSize
	for i := 0; i < chunkSize; i++ {
		if i == chunkSize-1 {
			go cracker(start, len(alphabet), targetLen, testPass, out)
		} else {
			go cracker(start, end, targetLen, testPass, out)
		}
		start, end = end, end+chunkSize
	}

	for reply := range out {
		fmt.Printf("\tPassword found: %s\n", reply)
		break
	}

}

func cracker(start, finish, targetLen int, target [16]byte, out chan<- string) {
	for _, char1 := range alphabet[start:finish] {
		if passEqual(string(char1), target) {
			out <- string(char1)
			break
		} else {
			recursiveCrack(string(char1), 2, targetLen, target, out)
		}
	}
	return
}

func recursiveCrack(chars string, currLen, targetLen int, target [16]byte, out chan<- string) {
	if currLen <= targetLen {
		for _, char := range alphabet {
			if passEqual(chars+string(char), target) {
				out <- chars + string(char)
				return
			}

			recursiveCrack(chars+string(char), currLen+1, targetLen, target, out)

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
