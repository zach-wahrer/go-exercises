package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*,./1234567890"

func main() {
	go waiting(100 * time.Millisecond)
	testPass := encrypter("A/3@a0")
	crackedPass, err := cracker(0, len(alphabet), testPass)
	if err != nil {
		log.Printf("cracker: %v", err)
	}
	if crackedPass == "" {
		fmt.Println("Password not cracked")
		return
	}
	fmt.Printf("Password is: %s\n", crackedPass)
}

func cracker(start, finish int, target [16]byte) (result string, err error) {
	for _, char1 := range alphabet[start:finish] {
		pass := string(char1)
		if encrypter(pass) == target {
			return pass, nil
		}
		for _, char2 := range alphabet {
			pass = string(char1) + string(char2)
			if encrypter(pass) == target {
				return pass, nil
			}
			for _, char3 := range alphabet {
				pass = string(char1) + string(char2) + string(char3)
				if encrypter(pass) == target {
					return pass, nil
				}
				for _, char4 := range alphabet {
					pass = string(char1) + string(char2) + string(char3) + string(char4)
					if encrypter(pass) == target {
						return pass, nil
					}
					for _, char5 := range alphabet {
						pass = string(char1) + string(char2) + string(char3) + string(char4) + string(char5)
						if encrypter(pass) == target {
							return pass, nil
						}
						for _, char6 := range alphabet {
							pass = string(char1) + string(char2) + string(char3) + string(char4) + string(char5) + string(char6)
							if encrypter(pass) == target {
								return pass, nil
							}
						}
					}
				}

			}
		}
	}
	return "", nil
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
