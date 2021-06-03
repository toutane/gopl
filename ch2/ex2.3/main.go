package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise2-3: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%v: %v\n", n, PopCount(uint64(n)))
	}
}

func PopCount(x uint64) int {
	var res byte
	for i := 0; i < 8; i++ {
		res += pc[byte(x>>(i*8))]
	}
	return int(res)
}
