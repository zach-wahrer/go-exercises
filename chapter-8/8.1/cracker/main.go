package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz1234567890"

func main() {
	go waiting(100 * time.Millisecond)
	testPass := encrypter("o12zab")
	go cracker(0, 6, testPass)
	go cracker(6, 12, testPass)
	go cracker(12, 18, testPass)
	go cracker(18, 24, testPass)
	go cracker(24, 30, testPass)
	cracker(30, len(alphabet), testPass)

}

func cracker(start, finish int, target [16]byte) {
	for _, char1 := range alphabet[start:finish] {
		pass := string(char1)
		if encrypter(pass) == target {
			fmt.Printf("Password is: %s\n", pass)
			return
		}
		for _, char2 := range alphabet {
			pass = string(char1) + string(char2)
			if encrypter(pass) == target {
				fmt.Printf("Password is: %s\n", pass)
				return
			}
			for _, char3 := range alphabet {
				pass = string(char1) + string(char2) + string(char3)
				if encrypter(pass) == target {
					fmt.Printf("Password is: %s\n", pass)
					return
				}
				for _, char4 := range alphabet {
					pass = string(char1) + string(char2) + string(char3) + string(char4)
					if encrypter(pass) == target {
						fmt.Printf("Password is: %s\n", pass)
						return
					}
					for _, char5 := range alphabet {
						pass = string(char1) + string(char2) + string(char3) + string(char4) + string(char5)
						if encrypter(pass) == target {
							fmt.Printf("Password is: %s\n", pass)
							return
						}
						for _, char6 := range alphabet {
							pass = string(char1) + string(char2) + string(char3) + string(char4) + string(char5) + string(char6)
							if encrypter(pass) == target {
								fmt.Printf("Password is: %s\n", pass)
								return
							}
						}
					}
				}

			}
		}
	}
	fmt.Printf("Password not found\n")
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
