package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz1234567890"

type target struct {
	hashedPass [16]byte
	maxLen     int
}

func main() {
	go waiting(100 * time.Millisecond)

	targetString := "gh0"
	targetPass := target{encrypter(targetString), len(targetString)}

	out := make(chan string)

	chunkSize := int(math.Max(float64(len(alphabet)/6), 1))
	start, end := 0, chunkSize
	for i := 0; i < chunkSize; i++ {
		if i == chunkSize-1 {
			go crackChunk(start, len(alphabet), targetPass, out)
		} else {
			go crackChunk(start, end, targetPass, out)
		}
		start, end = end, end+chunkSize
	}

	running := chunkSize
	for reply := range out {
		if reply == "" {
			running--
		} else {
			fmt.Printf("Password found: %s\n", reply)
			break
		}

		if running == 0 {
			fmt.Println("Password not found.")
			break
		}
	}

}

func crackChunk(start, finish int, targetPass target, out chan<- string) {
	for _, char1 := range alphabet[start:finish] {
		if passEqual(string(char1), targetPass.hashedPass) {
			out <- string(char1)
			return
		}
		recursiveCrack(string(char1), 2, targetPass, out)
	}
	out <- ""
	return
}

func recursiveCrack(chars string, currLen int, targetPass target, out chan<- string) {
	if currLen <= targetPass.maxLen {
		for _, char := range alphabet {
			if passEqual(chars+string(char), targetPass.hashedPass) {
				out <- chars + string(char)
				return
			}
			recursiveCrack(chars+string(char), currLen+1, targetPass, out)
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
