package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if anagrams(args[0], args[1]) {
		fmt.Println("These two strings are anagrams")
	} else {
		fmt.Println("These two strings are not anagrams")
	}
}

func anagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if strings.Contains(s2, s1[i:i+1]) {
			return strings.Count(s1, s1[i:i+1]) == strings.Count(s2, s1[i:i+1])
		}
		return false
	}
	return false
}
