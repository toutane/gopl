// Charcount computes couts of leters, digits and so on in their
// Unicode categories.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	cats := make(map[string]int)
	var utlen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for catName, rangeTable := range unicode.Properties {
			if unicode.In(r, rangeTable) {
				cats[catName]++
			}
		}
		counts[r]++
		utlen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\ncats\tcount\n")
	for k, n := range cats {
		fmt.Printf("%q\t%d\n", k, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utlen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d incalid UTF-8 characters\n", invalid)
	}
}
