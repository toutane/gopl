package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise2-4: %v\n", err)
			os.Exit(0)
		}
		fmt.Printf("%v: %v\n", n, PopCount(uint64(n)))
	}
}

func PopCount(x uint64) int {
	n := x
	var res int
	for i := 0; i < 64; i++ {
		if byte(n&1) == 1 {
			res++
			// fmt.Printf("(%b, %b) ", n, n>>1)
		}
		n = uint64(n >> 1)
	}
	return res
}
