package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	//b := []byte("Hello, 世界")
	//fmt.Printf("%s\n", runeReverse(b))
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Printf("%s\n", runeReverse([]byte(input.Text())))
	}
}

// Reverse reverse a slice of bytes in-place.
func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// RuneReverse allow to reverse a []byte slice that represents a UTF-8-encoded
// string, in place.
func runeReverse(s []byte) []byte {
	for i := 0; i < len(s); {
		_, size := utf8.DecodeRune(s[i:])
		// Reverse each rune.
		reverse(s[i : i+size])
		i = i + size
	}
	reverse(s)
	return s
}
