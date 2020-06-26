package main

import (
	"crypto/md5"
	"fmt"
	"math"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz1234567890@$."

func main() {
	go waiting(100 * time.Millisecond)

	targetString := "@.ax."
	testPass := encrypter(targetString)
	targetLen := len(targetString)

	found := make(chan string)
	notFound := make(chan struct{})

	chunkSize := int(math.Max(float64(len(alphabet)/6), 1))
	start, end := 0, chunkSize
	for i := 0; i < chunkSize; i++ {
		if i == chunkSize-1 {
			go cracker(start, len(alphabet), targetLen, testPass, found, notFound)
		} else {
			go cracker(start, end, targetLen, testPass, found, notFound)
		}
		start, end = end, end+chunkSize
	}

	running := chunkSize
	select {
	case reply := <-found:
		fmt.Printf("Password found: %s\n", reply)
		break
	case <-notFound:
		running--
		if running == 0 {
			fmt.Println("Password not found.")
			break
		}

	}

}

func cracker(start, finish, targetLen int, target [16]byte, found chan<- string, notFound chan<- struct{}) {
	for _, char1 := range alphabet[start:finish] {
		if passEqual(string(char1), target) {
			found <- string(char1)
			return
		}
		recursiveCrack(string(char1), 2, targetLen, target, found)
	}
	notFound <- struct{}{}
	return
}

func recursiveCrack(chars string, currLen, targetLen int, target [16]byte, found chan<- string) {
	if currLen <= targetLen {
		for _, char := range alphabet {
			if passEqual(chars+string(char), target) {
				found <- chars + string(char)
				return
			}
			recursiveCrack(chars+string(char), currLen+1, targetLen, target, found)
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
