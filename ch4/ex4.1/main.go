// Exercise 4.1 counts de the number of bits that are different in two SHA256
// hashes.
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	if len(args) > 1 {
		c1 = sha256.Sum256([]byte(args[0]))
		c2 = sha256.Sum256([]byte(args[1]))
	}
	fmt.Printf("%d bits are different between the two SHA256 hashes.", diff(c1, c2))
}

// Diff send to popcount function a byte where different bits are set
// bits and return to main the number of bits that are different.
func diff(a1, a2 [32]byte) int {
	var result int
	for i := 0; i < 32; i++ {
		diff := byte(a1[i] ^ a2[i])
		result += popcount(diff)
	}
	return result
}

// Popcount counts the number of set bits.
func popcount(diff byte) int {
	var ret int
	input := diff
	for k := 0; k < 64; k++ {
		if byte(input&1) == 1 {
			ret++
		}
		input = input >> 1
	}
	return ret
}
