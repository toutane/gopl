package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		words[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Printf("word\tcount\n")
	for w, n := range words {
		fmt.Printf("%s\t%d\n", w, n)
	}
}
