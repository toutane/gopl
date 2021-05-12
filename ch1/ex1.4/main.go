// Exercise1-4 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
// It also prints the file where there are duplicated lines.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countsForFiles := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countsForFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "exercise1-4: %v\n", err)
				continue
			}
			countLines(f, counts, countsForFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d \t%s\n", n, line)
		}
	}
	for filename, k := range countsForFiles {
		if k > 1 {
			fmt.Printf("%d \t%s\n", k, filename)
		}
	}
}

func countLines(f *os.File, counts, countsForFiles map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			countsForFiles[f.Name()]++
		}
	}
}
