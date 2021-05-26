package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := strings.Split(input.Text(), "")
		res := strings.Join(format(s), "")
		fmt.Println(res)
	}
}

// Format is a in-place function which eliminate adjacent duplicates in a []string slice.
func format(s []string) []string {
	var n int
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			for j := 0; j < len(s)-i-2; j++ {
				s[i+j] = s[i+j+2]
			}
			n = n + 1
			i = 0
		}
	}
	return s[:len(s)-2*n]
}
