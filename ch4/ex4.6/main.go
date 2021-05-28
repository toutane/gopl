package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string = "hello\n\t\t\nworld!"
	fmt.Printf("%s, %v\n", str, []byte(str))
	res := squashe([]byte(str))
	fmt.Printf("%s, %v \n", res, res)
}

// Squashe shashes each run of adjacent Unicode spaces into a single ASCII
// space.
func squashe(s []byte) []byte {
	var n int
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
			s[i] = 32
			for j := i + 1; j < len(s)-1; j++ {
				s[j] = s[j+1]
			}
			n = n + 1
			i = 0
		}
	}
	return s[:len(s)-n]
}
