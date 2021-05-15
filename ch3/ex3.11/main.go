// Exercise 3.11 format "-1234.5678" to "-1,234.567,8".
// It is a enhane version of exercise 3.10 (one with floating-point and optional
// sign support).
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	if s[:1] == "-" || s[:1] == "+" {
		pref := s[:1]
		buf.WriteString(pref)
		s = s[1:]
	}
	dot := strings.LastIndex(s, ".")
	if dot == -1 {
		dot = len(s)
	}
	s1 := s[:dot]
	mod := len(s1) % 3
	if mod == 0 {
		mod = 3
		buf.WriteString(s1[:mod])
	}
	for i := mod; i < len(s1); i += 3 {
		buf.WriteString(",")
		buf.WriteString(s1[i : i+3])
	}
	buf.WriteString(".")
	s2 := s[dot+1:]
	for i := 0; i < len(s2); i++ {
		if i%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(s2[i : i+1])
	}
	return buf.String()
}
