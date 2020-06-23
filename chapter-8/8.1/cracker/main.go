package main

import (
	"crypto/md5"
	"fmt"
	"log"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	testPass := encrypter("ZZZZ")
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
