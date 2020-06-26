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

type comms struct {
	found    chan string
	notFound chan struct{}
}

func main() {
	go waiting(100 * time.Millisecond)

	targetString := "aslm0"
	targetPass := target{encrypter(targetString), len(targetString)}

	channels := comms{make(chan string), make(chan struct{})}

	chunkSize := int(math.Max(float64(len(alphabet)/6), 1))
	start, end := 0, chunkSize
	for i := 0; i < chunkSize; i++ {
		if i == chunkSize-1 {
			go crackChunk(start, len(alphabet), targetPass, channels)
		} else {
			go crackChunk(start, end, targetPass, channels)
		}
		start, end = end, end+chunkSize
	}

	running := chunkSize
	for {
		select {
		case reply := <-channels.found:
			fmt.Printf("Password found: %s\n", reply)
			return
		case <-channels.notFound:
			running--
			if running == 0 {
				fmt.Println("Password not found.")
				return
			}
		}
	}
}

func crackChunk(start, finish int, targetPass target, channels comms) {
	for _, char1 := range alphabet[start:finish] {
		if passEqual(string(char1), targetPass.hashedPass) {
			channels.found <- string(char1)
			return
		}
		recursiveCrack(string(char1), 2, targetPass, channels)
	}
	channels.notFound <- struct{}{}
	return
}

func recursiveCrack(chars string, currLen int, targetPass target, channels comms) {
	if currLen <= targetPass.maxLen {
		for _, char := range alphabet {
			if passEqual(chars+string(char), targetPass.hashedPass) {
				channels.found <- chars + string(char)
				return
			}
			recursiveCrack(chars+string(char), currLen+1, targetPass, channels)
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
