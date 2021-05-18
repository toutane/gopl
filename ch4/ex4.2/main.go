// Exercise 4.2 prints the SHA256 hash of its standard input by default but
// support a command-line flag to print the SHA384 or SHA512 hash instead.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var str string
	flag.StringVar(&str, "s", "256", "Choose hash algorithm between SHA256, SHA384 or SHA512")
	flag.Parse()

	args := flag.Args()
	for _, arg := range args {
		switch str {
		case "384":
			fmt.Println(Sha384(arg))
		case "512":
			fmt.Println(Sha512(arg))
		default:
			fmt.Println(Sha256(arg))
		}
	}
}

func Sha256(input string) string {
	sha := sha256.Sum256([]byte(input))
	return fmt.Sprintf("SHA256 hash of %v: %x", input, sha)
}

func Sha384(input string) string {
	sha := sha512.Sum384([]byte(input))
	return fmt.Sprintf("SHA384 hash of %v: %x", input, sha)
}

func Sha512(input string) string {
	sha := sha512.Sum512([]byte(input))
	return fmt.Sprintf("SHA512 hash of %v: %x", input, sha)
}
