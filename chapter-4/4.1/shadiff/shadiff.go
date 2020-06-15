// shaDiff computes the number of bits that are different in two SHA256 hashes
package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
)

func main() {
	value1 := "zach"
	value2 := "cat"

	fmt.Printf("There are %v bit diffs between the SHA256 of %s and %s\n",
		shaDiff(hash(value1), hash(value2)), value1, value2)
}

func hash(value string) [32]byte {
	return sha256.Sum256([]byte(value))
}

func shaDiff(hash1, hash2 [32]byte) int {
	diff := 0
	for i := 0; i < 32; i++ {
		if bits.OnesCount(uint(hash1[i])) != bits.OnesCount(uint(hash2[i])) {
			diff++
		}
	}
	return diff
}
